package people

import (
	"testing"
	"time"
)

func TestPeople_Append(t *testing.T) {
	var pl People

	pl = pl.Append(Person{
		"Ivan",
		"Ivanov",
		time.Now(),
	})

	if l := len(pl); l != 1 {
		t.Error("len(People.Append(People.Person)) != 1. Returned: ", l)
	}
}

func TestPeople_Sort(t *testing.T) {
	var ps People

	dt1 := time.Date(2005, 8, 10, 0, 0, 0, 0, time.UTC)
	ps = append(ps, Person{
		"Ivan",
		"Ivanov",
		dt1,
	})

	dt2 := time.Date(2003, 8, 05, 0, 0, 0, 0, time.UTC)
	ps = append(ps, Person{
		"Ivan",
		"Ivanov",
		dt2,
	})

	dt3 := time.Date(2005, 8, 10, 0, 0, 0, 0, time.UTC)
	ps = append(ps, Person{
		"Artiom",
		"Ivanov",
		dt3,
	})

	ps.Sort()

	if !ps[0].birthDay.Equal(dt3) || ps[0].firstName+ps[0].lastName != "ArtiomIvanov" {
		t.Error("People.Sort() incorrect sorting by name. Returned:\n", ps)
	}

	if !ps[1].birthDay.Equal(dt1) || !ps[2].birthDay.Equal(dt2) {
		t.Error("People.Sort() incorrect sorting by date: ", ps)
	}

}
