package services_test

import (
	mock_db "mxshs/movieLibrary/src/adapters/repositories/mock"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/services"
	"mxshs/movieLibrary/src/utils"
	"testing"
	"time"
)

var as *services.ActorService

func init() {
    db := mock_db.NewDB()
    as = services.NewActorService(db, db)
}

func dateHelper(ds string) utils.Date {
    date, _ := time.Parse("02.01.2006", ds)

    return utils.Date{Time: date}
}


func TestCreateActor(t *testing.T) {
    tests := []struct{
        name string
        gender string
        birthdate utils.Date
        err bool
        expected domain.Actor
    }{
        {
            "Leonardo DiCaprio",
            "male",
            dateHelper("11.11.1974"),
            false,
            domain.Actor{
                Id: 1,
                Name: "Leonardo DiCaprio",
                Gender: "male",
                Birthdate: dateHelper("11.11.1974"),
            },
        },
        {
            "Kate Winslet",
            "female",
            dateHelper("05.10.1975"),
            false,
            domain.Actor{
                Id: 2,
                Name: "Kate Winslet",
                Gender: "female",
                Birthdate: dateHelper("05.10.1975"),
            },
        },
    }

    for _, test := range tests {
        res, err := as.CreateActor(test.name, test.gender, test.birthdate)
        if err != nil && !test.err {
            t.Fail()
        }

        if res.Id != test.expected.Id {
            t.Fail()
        }

        if res.Name != test.expected.Name {
            t.Fail()
        }

        if res.Gender != test.expected.Gender {
            t.Fail()
        }

        if res.Birthdate != test.expected.Birthdate {
            t.Fail()
        }
    }
}
