package server

// ctx := context.Background()

// conn, err := sql.Open("oracle", fmt.Sprintf("oracle://SYSTEM:SuperPassword@%s:%s/FREEPDB1", host, port.Port()))
// if err != nil {
// 	panic(err)
// }

// if err = conn.Ping(); err != nil {
// 	fmt.Println("Can't connect to the database: ", err)
// }

// migration, err := migrate.NewMigrateOracle(conn, "file://../database/migrations", "FREEPDB1")
// if err != nil {
// 	panic(err)
// }

// if err = migration.Execute(); err != nil {
// 	panic(err)
// }

// defer func() {
// 	err = conn.Close()
// 	if err != nil {
// 		fmt.Println("Can't close connection: ", err)
// 	}
// }()

// users, err := getUsers(context.Background(), conn)
// if err != nil {
// 	fmt.Println("Can't get users: ", err)
// 	return
// }

// for _, user := range users {
// 	fmt.Printf("%+v\n", user)
// }
// }

// type user struct {
// ID        int
// Name      NullableString
// Email     NullableString
// BirthDate NullableTime
// Active    NullableBool
// }

// func getUsers(ctx context.Context, conn *sql.DB) ([]user, error) {
// query := "SELECT ID, NAME, EMAIL, BIRTH_DATE, ACTIVE FROM JJ.USERS u"

// rows, err := conn.QueryContext(ctx, query)
// if err != nil {
// 	return nil, err
// }
// defer rows.Close()

// var users []user
// for rows.Next() {
// 	var user user
// 	err := rows.Scan(
// 		&user.ID,
// 		&user.Name.String,
// 		&user.Email.String,
// 		&user.BirthDate.Time,
// 		&user.Active.Bool,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	users = append(users, user)
// }
// return users, nil
// }

// type NullableTime struct {
// Time  *time.Time
// Valid bool
// }

// func NewNullableTime(t time.Time) NullableTime {
// return NullableTime{Time: &t, Valid: true}
// }

// type NullableString struct {
// String *string
// Valid  bool
// }

// func NewNullableString(s string) NullableString {
// return NullableString{String: &s, Valid: true}
// }

// type NullableBool struct {
// Bool  *bool
// Valid bool
// }

// func NewNullableBool(b bool) NullableBool {
// return NullableBool{Bool: &b, Valid: true}
// }

// type NullableInt struct {
// Int   *int
// Valid bool
// }

// func NewNullableInt(i int) NullableInt {
// return NullableInt{Int: &i, Valid: true}
// }

// func (n NullableTime) Value() (time.Time, error) {
// if n.Valid {
// 	return *n.Time, nil
// }
// return time.Time{}, nil
// }

// func (n NullableString) Value() (string, error) {
// if n.Valid {
// 	return *n.String, nil
// }
// return "", nil
// }

// func (n NullableBool) Value() (bool, error) {
// if n.Valid {
// 	return *n.Bool, nil
// }
// return false, nil
// }

// func (n NullableInt) Value() (int, error) {
// if n.Valid {
// 	return *n.Int, nil
// }
// return 0, nil
// }
