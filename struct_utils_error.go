package struct_copy

type StructUtilsError struct {
	err string
}

func (ste *StructUtilsError) Error() string {
	return "CopyProperties: " + ste.err
}
