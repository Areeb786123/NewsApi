package entitiy

type TokenEntity struct {
	UserId    string `json:"userId,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	Exp       int64  `json:"exp,omitempty"`
}
