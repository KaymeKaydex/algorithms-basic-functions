package conversion

import "strconv"

func SliceAtoi(a []string) ([]int, error) {
	si := make([]int, 0, len(a))
	for _, a := range a {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
