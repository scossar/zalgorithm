---
date: "07/22/2024"
tags: ["development"]
---

# Liftoff

Fuck it. I've been piecing together various apps to make the "ultimate" writing app.
There's some logic to integrating with existing apps, but I'm tired of working around
limitations imposed by other people's design decisions. I'm going to make my own
thing - starting with this post. The goal is to make the application that
_I_ want to use. So the first design decision is Neovim as the editor.
(maybe I just want different limitations.)

```yaml
// db/db.go

package db

import (
 "database/sql"
 "log"

 _ "github.com/mattn/go-sqlite3"
)

func PrepareDB() {
 db, err := sql.Open("sqlite3", "notes.db")
 checkErr(err)
 defer db.Close()

 ensureTablesExist(db)
}

func ensureTablesExist(db *sql.DB) {
 createTableQuery := `
  CREATE TABLE IF NOT EXISTS notes (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT NOT NULL,
      slug TEXT NOT NULL,
      markdown TEXT NOT NULL
  );
  `
 _, err := db.Exec(createTableQuery)
 checkErr(err)
}

func checkErr(err error) {
 if err != nil {
  log.Fatalf("database error: %v", err)
 }
}
```

and this file's in the db:

```TEXT

projects/golang/zalgorithm [🐹 v1.22.5]
❯ ls
db/  go.mod  go.sum  main.go  markdown/  notes.db

projects/golang/zalgorithm [🐹 v1.22.5]
❯ ls markdown/
liftoff.md  non-ordinary-states.md
```

copying that made me realize the `notes.db` file is going in the wrong directory.

```TEXT

projects/golang/zalgorithm [🐹 v1.22.5]
❯ go run .
Non ordinary states⏎

projects/golang/zalgorithm [🐹 v1.22.5]
❯ ls
db/  go.mod  go.sum  main.go  markdown/

projects/golang/zalgorithm [🐹 v1.22.5]
❯ ls db
db.go  notes.db

```

that's better:

## Vim is the editor!

I'm not even good at it.

### Add file to db on save?
