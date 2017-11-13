package aclutil

type Acl map[string]map[string]bool

func Init() Acl {
	return map[string]map[string]bool{}
}

func (acl Acl) Allow(role string, get, update, delete bool) Acl {

	acl[role] = map[string]bool{}
	if get {
		acl[role]["get"] = get
	}
	if update {
		acl[role]["update"] = update
	}
	if delete {
		acl[role]["delete"] = delete
	}
	return acl
}
