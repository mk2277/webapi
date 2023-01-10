package main

type account struct {
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	AccountNumber  int    `json:"accountnumber"`
	Address        string `json:"address"`
	PhoneNumber    int    `json:"phonenumber"`
	CurrentBalance int    `json:"currentbalance"`
}
