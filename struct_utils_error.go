package structutils

type StructUtilsError struct {
	err    string
	target interface{}
	source interface{}
}

func (ste *StructUtilsError) Error() string {
	return ste.err
}

func (ste *StructUtilsError) TargetNil() bool {
	return ste.target == nil
}

func (ste *StructUtilsError) SourceNil() bool {
	return ste.source == nil
}
