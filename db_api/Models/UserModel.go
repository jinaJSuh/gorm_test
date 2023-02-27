// Models/UserModel.go
package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// var User = []User{
// 	{Id: 1, Name: "Blue Train", Email: "John@gmail.com", Phone: "01022223333", Address: "Boston"},
// 	{Id: 2, Name: "Jeru", Email: "Gerry@gmail.com", Phone: "01022223333", Address: "New York"},
// 	{Id: 3, Name: "Sarah Vaughan and Clifford Brown", Email: "Gerry@gmail.com", Phone: "01022223333", Address: "Seoul"},
// }

func (b *User) TableName() string {
	return "user"
}
