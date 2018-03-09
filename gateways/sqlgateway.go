package sqlgateway

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"

  "github.com/sumeetv/dndcharacters/entities"
)

func GetClasses() {
  db, err := sql.Open("mysql", "root@/dnd")
  if (err != nil) {
    fmt.Println("Error connecting to db")
    return
  }

  classRows, err := db.Query("SELECT * FROM Classes")
  if (err != nil) {
    fmt.Println("Error fetching rows from table Classes")
  }
  defer classRows.Close()

  fmt.Println("Columns:")
  columns, err := classRows.Columns()
  for _, column := range(columns) {
    fmt.Println(column)
  }

  var playableClasses []entities.PlayableClass
  for classRows.Next() {
    var id int
    var name string
    err = classRows.Scan(&id, &name)
    playableClasses = append(playableClasses, entities.PlayableClass{ID: id, Name: name})
    fmt.Println("We have a class:", "|", id, "|", name, "|")
  }

  return
}
