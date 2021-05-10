package app_test

import (
	"database/app"
	"database/database"
	"testing"
)

func TestApp_AddUser(t *testing.T) {
	type args struct {
		username string
		email    string
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *app.App
		inspect func(r *app.App, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		{
			name: "Insert de chave que ñ existe no banco deve ser sucesso",
			init: func(*testing.T) *app.App {
				return app.NewApp(database.NewMemoryDatabase())
			},
			inspect: func(*app.App, *testing.T) {},
			args: func(*testing.T) args {
				return args{
					username: "nicolascb",
					email:    "nicolas@email.com",
				}
			},
			wantErr:    false,
			inspectErr: func(error, *testing.T) {},
		},
		// {
		// 	name: "Insert de chave duplicada deve retornar erro",
		// 	init: func(*testing.T) *app.App {

		// 		db := database.NewMemoryDatabase()
		// 		db.Insert("duplicado@email", "nicolas")
		// 		return app.NewApp(db)
		// 	},
		// 	inspect: func(*app.App, *testing.T) {},
		// 	args: func(*testing.T) args {
		// 		return args{
		// 			username: "nicolascb",
		// 			email:    "duplicado@email",
		// 		}
		// 	},
		// 	wantErr: true,
		// 	inspectErr: func(err error, t *testing.T) {
		// 		if err.Error() != "email duplicado@email já existe" {
		// 			t.Fatalf("Deu ruim aqui")
		// 		}
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			err := receiver.AddUser(tArgs.username, tArgs.email)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("App.AddUser error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
