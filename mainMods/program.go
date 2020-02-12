package mainMods

type Program struct {
	ID   string `json:"id" description:"identifier of the program"`
	TITLE string `json:"title" description:"title of the program" `
	DESC  string    `json:"des" description:"description of the program"`
}