package full_outer_join

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	f1Con, errf1 := os.ReadFile(f1Path)
	f2Con, errf2 := os.ReadFile(f2Path)

	if errf1 != nil {
		fmt.Println("not found f1")
	}

	if errf2 != nil {
		fmt.Println("not found f2")
	}

	arrF1 := strings.Split(string(f1Con), "\n")
	arrF2 := strings.Split(string(f2Con), "\n")

	mpf1 := make(map[string]int, len(arrF1))
	mpf2 := make(map[string]int, len(arrF2))

	for i, val := range arrF1 {
		mpf1[val] = i
	}

	for i, val := range arrF2 {
		mpf2[val] = i
	}

	res := []string{}

	for key, _ := range mpf1 {
		_, ok := mpf2[key]

		if !ok {
			res = append(res, key)
		} else {
			delete(mpf2, key)
		}
	}

	for key, _ := range mpf2 {
		res = append(res, key)
	}

	sort.Strings(res)

	s := ""

	for i, val := range res {
		s += val

		if i != len(res) - 1 {
			s += "\n"
		}
	}

	if err := os.WriteFile(resultPath, []byte(s), os.ModePerm); err != nil {
		fmt.Println("not write")
	}
}
