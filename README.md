# MalmaBlog (Back End)

MalmaBlog is a minimalist personal blog application that doesn't have many features, making it easy to use.

This repository contains the backend (Rest API) of my application, MalmaBlog.

The repository uses Go as the main programming language, and several third-party applications such as:
- [Visual Studio Code](https://https://code.visualstudio.com/) as the text editor application to create the application.
- [Postman](https://postman.com) as the application for API testing.
- [PostgreSQL](https://postgresql.org) as the database for the application.
- [Fiber](https://gofiber.io) as the Go web framework for Rest API.
- [Gorm](https://gorm.io) as the Go library for the database (ORM).
- [Viper](https://github.com/spf13/viper) as the Go library for configuration in the application.
- [Gomal](https://github.com/ItsMalma/gomal) as the Go library for validation.

## Instalation and Run

To clone this repository using Git, run the following command:
```
git clone https://github.com/ItsMalma/malma-blog-be.git
```

Navigate to the cloned folder:
```
cd malma-blog-be
```

After that, set the environment for our application using the following command:
```
# Untuk linux, gunakan perintah berikut:
export ENV=ENVIRONMENT_NAME

# Untuk windows, gunakan perintah berikut:
set ENV=ENVIRONMENT_NAME
```

Replace `ENVIRONMENT_NAME` from the command above with either `DEV` or `PROD`.

Each `ENVIRONMENT_NAME` has its own uniqueness, such as:
- `DEV` indicates that the application is running on development, and will retrieve the config from the `configs/development.json` file.
- `PROD` indicates that the application is running on production, and will retrieve the config from the `configs/production.json` file.

After setting the environment, we create a config file according to the `ENVIRONMENT_NAME` that has been set.

Here is the JSON schema of the config file:
```json
{
    "server": {
        "host": "string;required",
        "port": "number;required"
    },
    "database": "string;required"
}
```

Then, run the application. Make sure that Make and Go are installed on your computer, and then run the following command:
```
make run
```

Or, if you don't have Make and only have Go, you can run the following command:
```
go run main.go
```

If you don't have both, please install Go or Go with Make on your computer!

