package helpers

type SomeType struct {
	TypeName    string
	TypeNumber  int
	privateType string
}

func (s *SomeType) PrivateType(pt string) {
	s.privateType = pt
}
