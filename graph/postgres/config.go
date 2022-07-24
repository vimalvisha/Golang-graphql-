package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jackc/pgx/v4"
	"github.com/vishal/gqlgen-todos/graph/model"
)

func AddStudent(id int, full_name string) {
	if pool == nil {
		pool = GetPool()
	}

	querystring := `insert into student (id,full_name) values($1,$2)`
	err := pool.QueryRow(context.Background(), querystring, &id, &full_name)
	if err != nil {

		log.Println("err", err)

	}
}
func Addperson(id int, first_name string, last_name string, gender string, car_id int) {
	if pool == nil {
		pool = GetPool()
	}
	querystring := `insert into person(id ,first_name,last_name,gender,car_id)Values($1,$2,$3,$4,$5)`
	err := pool.QueryRow(context.Background(), querystring, &id, &first_name, &last_name, &gender, &car_id)
	if err != nil {
		log.Println("err", err)
	}
}

func CreateExcel() {
	if pool == nil {
		pool = GetPool()
	}
	var personArr []*model.Iperson
	var id int
	var full_name string
	query := `select * from student`
	rows, err1 := pool.Query(context.Background(), query)

	if err1 != nil {
		log.Println(err1)
	} else {

		for rows.Next() {
			err := rows.Scan(&id, &full_name)
			if err != nil {
				log.Fatal(err)
			}
			person := model.Iperson{
				ID:       id,
				FullName: full_name,
			}
			personArr = append(personArr, &person)
		}
	}
	f := excelize.NewFile()
	for i := 0; i < len(personArr); i++ {

		s := i + 1
		A := "A" + strconv.Itoa(s)
		B := "B" + strconv.Itoa(s)

		f.SetCellValue("Sheet1", A, personArr[i].ID)
		f.SetCellValue("Sheet1", B, personArr[i].FullName)
	}
	if err := f.SaveAs("simple3.xlsx"); err != nil {
		log.Fatal(err)
	}
}
func SetFromExcel() {
	if pool == nil {
		pool = GetPool()
	}
	f, err := excelize.OpenFile("simple1.xlsx")

	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.Rows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		row := rows.Columns()
		querystring := `insert into student (id,full_name) values (DEFAULT, $1)`
		err := pool.QueryRow(context.Background(), querystring, row[1])
		if err != nil {

			log.Println("err", err)
		}
	}
}
func Retrievedataforchart() []string {
	if pool == nil {
		pool = GetPool()
	}
	var id int
	var full_name string
	query := `SELECT * from student`
	rows, err1 := pool.Query(context.Background(), query)
	if err1 != nil {
		log.Println(err1)
	} else {
		f := excelize.NewFile()
		var arr []*model.Iperson
		for rows.Next() {
			err := rows.Scan(&id, &full_name)
			if err != nil {
				log.Fatal(err)
			}
			x := model.Iperson{
				ID:       id,
				FullName: full_name,
			}

			arr = append(arr, &x)
		}
		log.Println(arr)
		for i := 0; i <= len(arr)-1; i++ {
			j := i + 1
			A := "A" + strconv.Itoa(j+1)
			B := "B" + strconv.Itoa(j+1)
			f.SetCellValue("Sheet1", "A1", "id")
			f.SetCellValue("Sheet1", "B1", "full_name")
			f.SetCellValue("Sheet1", A, arr[i].ID)
			f.SetCellValue("Sheet1", B, arr[i].FullName)
		}
		// pie chart for data
		if err := f.AddChart("Sheet1", "E1", `{"type":"pie","series":[{"name":"pie data","name":"Sheet1!$B$2:$B$13","values":"Sheet1!$A$2:$A$13"}],"title":{"name":"Student Chart"}}`); err != nil {
			println(err.Error())
		}
		// bar chart for data
		// if err := f.AddChart("Sheet1", "E1", `{"type":"bar","series":[{"name":"Bar chart data","name":"Sheet1!$A$2:$A$4","values":"Sheet1!$B$2:$B$4"}],"title":{"name":"Fruit Bar Chart"}}`); err != nil {
		//  println(err.Error())
		// }

		// Save xlsx file by the given path.
		if err := f.SaveAs("studentchart.xlsx"); err != nil {
			println(err.Error())
		}

	}
	arr := []string{"error"}
	return arr
}
func Update(id *int, full_name *string) {
	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	querystring := "update student set full_name = $1 where id = $2"
	commandTag, err := tx.Exec(context.Background(), querystring, &full_name, &id)
	if err != nil {
		log.Println(err, "1")
		return
	}
	if commandTag.RowsAffected() != 1 {

		log.Println("2")
		return
	} else {

		log.Println("updated")
	}

	txErr := tx.Commit(context.Background())
	if txErr != nil {
		log.Println(txErr, "4")
	}
}
func Registeration(email_id string, username string, pasword string, phone_no string) {
	if pool == nil {
		pool = GetPool()
	}

	query := `insert into userlogin (email_id,username,pasword,phone_no) values($1,$2,$3,$4)`
	// data := util.Encrypt(pasword, util.Passphrase)
	fmt.Println(pasword)
	row := pool.QueryRow(context.Background(), query, &email_id, &username, &pasword, &phone_no)
	var count int
	err := row.Scan(&count)
	if row != nil {
		log.Println("err", err.Error())
	}
	log.Println(count)
}

func UserResponse(emailid string, pasword string) ([]*model.Loginop, error) {
	if pool == nil {
		pool = GetPool()
	}

	var username string

	var phone_no string

	query := ` SELECT * from userlogin where email_id=$1 and pasword=$2`
	var userarr []*model.Loginop
	count := 0
	// data := util.Decode(pasword)
	rows, err1 := pool.Query(context.Background(), query, &emailid, &pasword)
	if err1 != nil {
		if err1.Error() == pgx.ErrNoRows.Error() {
			log.Println("err1", err1.Error())
			return nil, errors.New("incorrect emailid or password")
		} else {
			log.Println("err1", err1.Error())
			return nil, err1
		}
	} else {
		for rows.Next() {
			err := rows.Scan(&emailid, &username, &pasword, &phone_no)

			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			count++

			user := model.Loginop{
				EmailID:  emailid,
				Username: username,
				PhoneNo:  phone_no,
			}
			userarr = append(userarr, &user)
		}
		if count < 1 {
			return nil, errors.New("incorrect emailid or password")
		}

	}
	return userarr, nil
}

func FetchResponse(first *int, after *int) []*model.Loginop {
	if pool == nil {
		pool = GetPool()
	}

	var inputargs []interface{}
	var email_id string
	var username string
	var pasword string
	var phone_no string

	query := `SELECT * from userlogin `

	var getuser []*model.Loginop

	if first != nil && after != nil {
		query = query + ` LIMIT $1 OFFSET $2`
		inputargs = append(inputargs, first, after)

	}
	log.Println(inputargs...)
	rows, err2 := pool.Query(context.Background(), query, first, after)
	if err2 != nil {
		log.Println(err2)
	} else {
		for rows.Next() {
			err := rows.Scan(&email_id, &username, &pasword, &phone_no)
			if err != nil {
				log.Fatal(err)

			}
			person := model.Loginop{
				EmailID:  email_id,
				Username: username,
				PhoneNo:  phone_no,
			}
			getuser = append(getuser, &person)
		}
	}
	return getuser
}
