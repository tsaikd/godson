package payload

type Push struct {
	Actor      User       `json:"actor,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	Push       PushDetail `json:"push,omitempty"`
}

type PushDetail struct {
	Changes []PushChange `json:"changes,omitempty"`
}

type PushChange struct {
	New     PushChangeState `json:"new,omitempty"`
	Old     PushChangeState `json:"old,omitempty"`
	Created bool            `json:"created,omitempty"`
	Forced  bool            `json:"forced,omitempty"`
	Closed  bool            `json:"closed,omitempty"`
}

type PushChangeState struct {
	Type   string                `json:"type,omitempty"`
	Name   string                `json:"name,omitempty"`
	Target PushChangeStateTarget `json:"target,omitempty"`
}

type PushChangeStateTarget struct {
	Type string `json:"type,omitempty"`
	Hash string `json:"hash,omitempty"`
}
