package models

type Link struct{
	ID string `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`

}

type UpdateLink struct{
	Name string `json:"name,omitempty"`
	URL string `json:"url,omitempty"`
}