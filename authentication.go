package main

import(
	_"bytes"
	"crypto/rand"
	"fmt"
	"time"
	"strings"
	"io"
	"crypto/sha512"
	"database/sql"
	
)

type User struct{
	UserId uint
	AccessToken []byte
}


/*
	Struct returned via JSON for Authentication
*/
type AccessCode struct{
	//The AccessToken bytes. This will be cross-referenced with a database
	AccessToken []byte
	//When the Token Expires
	Expires time.Time

	UserId uint;
}

const ACCESS_TOKEN_LENGTH = 50

func CreateSessionId() ([]byte,error) {
	b := make([]byte, ACCESS_TOKEN_LENGTH)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return nil, fmt.Errorf("Could not make access token")
	}

	return b, nil
	// The slice should now contain random bytes instead of only zeroes.
}

func AuthenticateUser(user string, pass string, con Configuration, db * sql.DB) (AccessCode, error) {
	var pw string

	err3 := db.QueryRow("SELECT HEX(password) FROM users WHERE name=?",user).Scan(&pw)

	//SHA hashing
	h512 := sha512.New()
	io.WriteString(h512, pass)



	if err3 == nil{
		if strings.EqualFold(fmt.Sprintf("%s",pw),fmt.Sprintf("%x",h512.Sum(nil))){
			b, err := CreateSessionId()
			if b!=nil{
				a := AccessCode{AccessToken: b, Expires: time.Now().UTC().Add(24 * time.Hour) }
				err4 := addToDatabase(a, db)
				if err4!=nil{
					AccessCode{AccessToken: nil, Expires: time.Now().UTC() }, err4 
				}else{
					return a, nil
				}
				
			}else{
				fmt.Println("Auth Error: 1")
				return AccessCode{AccessToken: nil, Expires: time.Now().UTC() }, err
			}
		}else{
			fmt.Println("Auth Error: 2")
			return AccessCode{AccessToken: nil, Expires: time.Now().UTC() }, fmt.Errorf("Could not authenticate user")
		}
	}else{
		fmt.Println("Auth Error: 3")
		return AccessCode{AccessToken: nil, Expires: time.Now().UTC() }, err3
	}
	//Shouldn't Get Here
	fmt.Println("Auth Error: 4")
	return AccessCode{AccessToken: nil, Expires: time.Now().UTC() }, fmt.Errorf("Could not authenticate user")
}
//Add's an Access Token to the Database
func addToDatabase(a AccessCode, db * sql.DB) error {
	
	_, err := db.Exec("INSERT INTO access_token VALUES(?,?,NOW())",fmt.Sprintf("%x",a.AccessToken),1)

	return err
}