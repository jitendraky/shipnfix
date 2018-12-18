package user

import (
	"../../app"
	templates ".."
	"net/http"
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	dynamosession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/aws/request"
	//request2 "github.com/dgrijalva/jwt-go/request"
)




type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}



type Page struct {
Title, Content string
}
func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	



//	tmpl.Execute(w, data)






	templates.RenderTemplate(w, "user", session.Values["profile"])
}
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	billing := os.Getenv("BILLING")
	MQ :=os.Getenv("MQ")
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		//err2 := r.ParseForm()

		fmt.Println(r.Form["task"])

		if billing == "TRUE" {

			sess, err := dynamosession.NewSession(&aws.Config{
				Region: aws.String("us-east-1")},
			)

			// Create DynamoDB client
			svc := dynamodb.New(sess)

			// Create item in table Movies
			input := &dynamodb.UpdateItemInput{
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":r": {
						N: aws.String("0.5"),
					},
				},
				TableName: aws.String("shipfixbilling"),
				Key: map[string]*dynamodb.AttributeValue{
					"user": {
						N: aws.String(session.ID),
					},
					"task": {
						S: aws.String(string(r.Form["Task"][0])),
					},
				},
				ReturnValues:     aws.String("UPDATED_NEW"),
				UpdateExpression: aws.String("set info.rating = :r"),
			}

			_, err = svc.UpdateItem(input)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			//fmt.Println("Successfully updated 'The Big New Movie' (2015) rating to 0.5")

		}


		if MQ == "TRUE" {

		}
		for key, value := range r.Form{
			fmt.Println(key,value)
			//fmt.Println(err2["Task"])




			}

	}
	templates.RenderTemplate(w, "user", session.Values["profile"])
	}

func MytaskHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(session.Values)
	//
	// templates.RenderTemplate(w, "user", session.Values["profile"])
}
