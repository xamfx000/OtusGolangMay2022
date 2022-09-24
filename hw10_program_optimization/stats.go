package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}
type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

func getUsers(r io.Reader) (result [100_000]User, err error) {
	scanner := bufio.NewScanner(r)

	const maxCapacity int = 1000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	var user User
	i := 0
	for scanner.Scan() {
		if err = easyjson.Unmarshal(scanner.Bytes(), &user); err != nil {
			return
		}
		result[i] = user
		i++
	}
	return
}

func countDomains(u [100000]User, domain string) (DomainStat, error) {
	result := make(DomainStat)

	var rootDomain string
	for _, user := range u {
		if strings.Contains(user.Email, "."+domain) {
			rootDomain = strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[rootDomain]++
		}
	}
	return result, nil
}
