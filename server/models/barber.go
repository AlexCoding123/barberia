package models

type Barber struct {
    ID         uint   `json:"id"`
    Name       string `json:"name"`
    LastName   string `json:"last_name"`
    BarberName string `json:"barber_name"`
    PhoneNumber int    `json:"phone_number"`
    Password   string `json:"password"`
}

