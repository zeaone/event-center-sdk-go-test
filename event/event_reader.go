package event

var _ EventReader = (*IcEvent)(nil)

func (i *IcEvent) Source() string {
	return i.Event.Source()
}

func (i *IcEvent) Type() string {
	return i.Event.Type()
}

func (i *IcEvent) Subject() string {
	return i.Event.Subject()
}

func (i *IcEvent) Body() []byte {
	return i.Event.Data()
}
func (i *IcEvent) Action() string {
	extensions := i.Extensions()
	if _, ok := extensions["action"].(string) ; ok {
		return extensions["action"].(string)
	}
	return ""
}

func (i *IcEvent) Region() string {
	extensions := i.Extensions()
	if _, ok := extensions["region"].(string) ; ok {
		return extensions["region"].(string)
	}
	return ""
}

func (i *IcEvent) Accountid() string {
	extensions := i.Extensions()
	if _, ok := extensions["accountid"].(string) ; ok {
		return extensions["accountid"].(string)
	}
	return ""
}

func (i *IcEvent) Resource() string {
	extensions := i.Extensions()
	if _, ok := extensions["resource"].(string) ; ok {
		return extensions["resource"].(string)
	}
	return ""
}

func (i *IcEvent) Extensions() map[string]interface{} {
	return i.Event.Extensions()
}