package test

import "os"

type EnvSetter struct {
	Env string
}

func NewEnvSetter(env string) *EnvSetter {
	return &EnvSetter{
		Env: env,
	}
}

func (e *EnvSetter) SetEnv() {
	os.Setenv("PORT", "8080")
	os.Setenv("SENS_HOST", "https://sens.apigw.ntruss.com")
	os.Setenv("SENS_SERVICE_ID", "ncp:sms:kr:290881020329:marketbill-project")
	os.Setenv("SENS_ACCESS_KEY_ID", "2aJkrtHdUtk5NP4oG8yh")
	os.Setenv("SENS_SECRET_KEY", "x2A3OXOz0P1qmaLDnTiqo2dQ7if6BzOElQEPNg6b")

	switch e.Env {
	case "local":
		os.Setenv("PROFILE", "local")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PW", "postgres")
		os.Setenv("DB_NET", "tcp")
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "marketbill-test")
	case "dev":
		os.Setenv("PROFILE", "dev")
		os.Setenv("DB_USER", "marketbill")
		os.Setenv("DB_PW", "marketbill1234!")
		os.Setenv("DB_NET", "tcp")
		os.Setenv("DB_HOST", "marketbill-db.ciegftzvpg1l.ap-northeast-2.rds.amazonaws.com")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "dev-db")
	case "prod":
		os.Setenv("PROFILE", "prod")
		os.Setenv("DB_USER", "marketbill")
		os.Setenv("DB_PW", "marketbill1234!")
		os.Setenv("DB_NET", "tcp")
		os.Setenv("DB_HOST", "marketbill-db.ciegftzvpg1l.ap-northeast-2.rds.amazonaws.com")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "prod-db")
	}

}
