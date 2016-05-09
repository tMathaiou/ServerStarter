package userModel


import (
    "golang.org/x/crypto/bcrypt"  
    "fmt"  
    "startup/database"
)


type Users struct {  
  Id uint `gorm:"primary_key"`
  Email string `gorm:"type:varchar(100);unique_index;not null"`
  Password string `gorm:"not null"`
  Role string 
}

type UsersArray struct { Users []Users }


func (u *Users) BeforeSave() (err error) {   

    // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    u.Password = string(hashedPassword)
    return
}

func ComparePass(password, hashedPassword string) (err error) {   
  err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
  return err
}

func Migrate(){	
	db.Db.AutoMigrate(&Users{})		
}

func Save(u *Users) (err error) {   
  if err != nil {
   fmt.Printf("%v", err)
  }

  db.Db.Create(&Users{Email: u.Email, Password: u.Password, Role: u.Role})
  return 
}

func FindById(id string) (Users) {
  user :=  Users{} 

  db.Db.Find(&user, id) 

  return user
}

func FindByEmail(email string) (Users) {
  user :=  Users{}   

  db.Db.Where("email = ?", email).First(&user)

  return user
}


func FindAll() (UsersArray) {  
  users :=  UsersArray{}
  db.Db.Find(&users.Users)
  return users
}

func Update(id string, data *Users) (err error) {   
  user :=  Users{} 

  db.Db.Find(&user, id) 
  
  user.Email = data.Email

  user.Role = data.Role


  if data.Password != "" {   
    user.Password = data.Password
  }

  db.Db.Save(&user) 
  return 
}


func Delete(id string) {
  user :=  Users{} 

  db.Db.Delete(&user, id)
}