package core

import (
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Login(username string, password string, db DB_Handler) (string, error) {
	//u := User{}
	//passwordHash := hash(password)
	//query, args := u.Verify(username, strconv.Itoa(int(passwordHash)))
	//db.QueryRow(query, args...)
	return "", nil
}
