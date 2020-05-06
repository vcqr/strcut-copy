package struct_copy

import (
	"reflect"
	"strings"
)

type StructUtil interface {
	CopyProperties(target, source interface{}) error
}

type StructUtils struct {
	assert *Assert
}

func NewStructUtils() *StructUtils {
	assert := NewAssert()

	return &StructUtils{
		assert: assert,
	}
}

type ReflectInfo struct {
	t reflect.Type
	v reflect.Value
}

// 拷贝结构体相关属性
// 拷贝条件：
//      a. 字段名字一样
//      b. 字段类型一样
//      c. 包不同，但类型名称一样
func (st *StructUtils) CopyProperties(target, source interface{}) error {
	srcV, srcT, err := st.getSourceReflectInfo(source)
	if err != nil {
		return err
	}

	destV, destT, err := st.getTargetReflectInfo(target)
	if err != nil {
		return err
	}

	srcMap := make(map[string]*ReflectInfo)

	for i := 0; i < srcT.NumField(); i++ {
		rf := &ReflectInfo{
			t: srcT.Field(i).Type,
			v: srcV.Field(i),
		}

		srcMap[srcT.Field(i).Name] = rf
	}

	for i := 0; i < destT.NumField(); i++ {
		destFieldV := destV.Field(i)
		if !destFieldV.CanSet() {
			continue
		}

		destFieldT := destT.Field(i)

		if ri, ok := srcMap[destFieldT.Name]; ok {
			// 类型一致，直接赋值
			if ri.t == destFieldT.Type {
				destV.Field(i).Set(ri.v)
			} else {
				riEndName := st.getTypeEndName(ri.t.String())
				dtEndName := st.getTypeEndName(destFieldT.Type.String())

				// 解析不同包，同类型名
				if riEndName == dtEndName {
					if destFieldV.Kind() == reflect.Ptr {
						if ri.v.IsNil() {
							continue
						}

						// 目标结构是指针，则需要初始化
						if destFieldV.IsNil() {
							destFieldV.Set(reflect.New(destFieldT.Type.Elem()))
						}

						if err := st.CopyProperties(destFieldV.Interface(), ri.v.Interface()); err != nil {
							continue
						}
					} else if destFieldV.Kind() == reflect.Slice && ri.t.Kind() == reflect.Slice {
						di := &ReflectInfo{
							t: destFieldT.Type,
							v: destFieldV,
						}
						st.SliceCopyProperties(di, ri)
					} else if destFieldV.Kind() == reflect.Map && ri.t.Kind() == reflect.Map {
						// 校验两个Map，Key的数据类型是否一致
						if destFieldT.Type.Key().Kind() != ri.t.Key().Kind() {
							continue
						}

						di := &ReflectInfo{
							t: destFieldT.Type,
							v: destFieldV,
						}

						st.MapCopyProperties(di, ri)
					} else if destFieldV.Kind() == reflect.Array && ri.t.Kind() == reflect.Array {
						di := &ReflectInfo{
							t: destFieldT.Type,
							v: destFieldV,
						}

						st.ArrayCopyProperties(di, ri)
					} else {
						if err := st.CopyProperties(destFieldV.Addr().Interface(), ri.v.Interface()); err != nil {
							continue
						}
					}
				}
			}
		}
	}

	return nil
}

func (st *StructUtils) ArrayCopyProperties(dest, src *ReflectInfo) error {
	iL := st.Min(dest.t.Len(), src.t.Len())
	var idxT reflect.Type

	for i := 0; i < iL; i++ {
		item := dest.v.Index(i)
		if idxT == nil {
			idxT = item.Type()
		}

		if idxT.Kind() == reflect.Ptr {
			if src.v.Index(i).IsNil() {
				continue
			}

			// 指针类型需要初始化数据
			item.Set(reflect.New(idxT.Elem()))
			st.CopyProperties(item.Interface(), src.v.Index(i).Interface())
		} else {
			st.CopyProperties(item.Addr().Interface(), src.v.Index(i).Interface())
		}
	}

	return nil
}

func (st *StructUtils) SliceCopyProperties(dest, src *ReflectInfo) error {
	retErr := &StructUtilsError{}

	iL := src.v.Len()
	arr := reflect.MakeSlice(dest.t, iL, iL)

	var idxT reflect.Type
	for i := 0; i < iL; i++ {
		item := arr.Index(i)
		if idxT == nil {
			idxT = item.Type()
		}

		if idxT.Kind() == reflect.Ptr {
			if src.v.Index(i).IsNil() {
				continue
			}

			// 指针类型需要初始化数据
			item.Set(reflect.New(idxT.Elem()))
			st.CopyProperties(item.Interface(), src.v.Index(i).Interface())
		} else {
			st.CopyProperties(item.Addr().Interface(), src.v.Index(i).Interface())
		}
	}

	if !arr.IsNil() {
		dest.v.Set(arr)
	}

	return retErr
}

func (st *StructUtils) MapCopyProperties(dest, src *ReflectInfo) error {
	iL := src.v.Len()
	mp := reflect.MakeMapWithSize(dest.t, iL)

	for _, k := range src.v.MapKeys() {
		if src.v.MapIndex(k).IsNil() {
			continue
		}

		if dest.t.Elem().Kind() == reflect.Ptr {
			// 指针类型需要初始化数据
			item := reflect.New(dest.t.Elem().Elem())
			st.CopyProperties(item.Interface(), src.v.MapIndex(k).Interface())

			mp.SetMapIndex(k, item)
		} else {
			// 初始化数据
			item := reflect.New(dest.t.Elem())
			st.CopyProperties(item.Interface(), src.v.MapIndex(k).Interface())

			mp.SetMapIndex(k, item.Elem())
		}
	}

	if !mp.IsNil() {
		dest.v.Set(mp)
	}

	return nil
}

func (st *StructUtils) copyProperties(dest, src interface{}) error {
	return st.CopyProperties(dest, src)
}

// 获取不同包下的同名类型的名字
func (st *StructUtils) getTypeEndName(typeName string) string {
	strArr := strings.Split(typeName, ".")
	return strArr[len(strArr)-1]
}

// 反射出源结构体信息
func (st *StructUtils) getSourceReflectInfo(obj interface{}) (reflect.Value, reflect.Type, error) {
	st.assert.NotNil(obj, "CopyProperties: Source must not be nil")

	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), t, &StructUtilsError{"Source are not struct"}
	}

	return v, t, nil
}

// 反射出目标结构体信息
func (st *StructUtils) getTargetReflectInfo(obj interface{}) (reflect.Value, reflect.Type, error) {
	st.assert.NotNil(obj, "CopyProperties: Target must not be nil")

	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if t.Kind() != reflect.Ptr {
		return reflect.Zero(v.Type()), t, &StructUtilsError{"Target are not ptr"}
	} else {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), t, &StructUtilsError{"Target are not struct"}
	}

	return v, t, nil
}

// 获取两个值的最小值
func (st *StructUtils) Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
