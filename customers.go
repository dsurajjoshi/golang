package customer_list

type Customer struct {
	Customer_id    string `json:"customer_id"`
	Activated_pack string `json:"activated_pack"`
	Days_remaining int    `json:"days_remaining"`
}
