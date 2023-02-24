package config

var Config = []byte(`
{
	"app":{
		"port":"8081"
	},
	"mysql":{
		"dbUser":"root",
		"dbPassword":"123456",
		"dbName":"ttt",
		"dbHost":"127.0.0.1",
		"dbPort":"3304"
	},
	"redis":{
		"addr": "127.0.0.1:6379",
		"password":"",
		"db":"0"
	},
	"jwt":{
		"secret":"gouzixihuanchibaozi"
	},
	"sign":{
		"salt":"***!@#!@#89~#BCV***!@#!@#89~**",
		"open":"0"
	},
	"rsa_request":{
		"V3":{
			"public_key":"MIIBCgKCAQEAxP77GseiMAEcJePuh4upDGVgjBGlQY/0K1aiOud6x+3178xpsQhb+yZ3X5O5L9jzTCXsJVW6TvvM7btRwBaW6JmSAd+1/svtvmaRE+SNIK5R+ftVXYWYM//pI4Z4NVWo1Hm6sxAy/YOJpNZEc3Lyfvkl48hYUZ0cTij/vbvWgDlp7In0j1GoomgWil/+ajmUy2jpsPA8/K1xJIwl4QcJlNx2pONVNAyDV5jcQg+uCnWwyZi5myaOjfAPvHdBIXpDVRnhyQlRNElFgGiJ+DeZld90x2/4AOW1I8tkRD1tbUvyNvnNohGjkJ9Mduop1eNDWmgA90FjIe41Gmonx1CwkwIDAQAB",
			"private_key":"MIIEpAIBAAKCAQEAxP77GseiMAEcJePuh4upDGVgjBGlQY/0K1aiOud6x+3178xpsQhb+yZ3X5O5L9jzTCXsJVW6TvvM7btRwBaW6JmSAd+1/svtvmaRE+SNIK5R+ftVXYWYM//pI4Z4NVWo1Hm6sxAy/YOJpNZEc3Lyfvkl48hYUZ0cTij/vbvWgDlp7In0j1GoomgWil/+ajmUy2jpsPA8/K1xJIwl4QcJlNx2pONVNAyDV5jcQg+uCnWwyZi5myaOjfAPvHdBIXpDVRnhyQlRNElFgGiJ+DeZld90x2/4AOW1I8tkRD1tbUvyNvnNohGjkJ9Mduop1eNDWmgA90FjIe41Gmonx1CwkwIDAQABAoIBADPYZ33EzIqVwDK4qi5CM1tv0tkKS9PVbw1433lhMo8rsW5K/gX5jTusA/7Dghl1n5KK6Htq9H2VB3oJspo7V7wfrq2PkvZb9VYG6Gez3vCa8Dg6TV0vq78Duvf+i+R9htFkuB4oRUMzOyaGvb0ko22HZNFuCNBx6OEpKSVmKi3UeDEyCw2FgncAlyy4Mi3FKZ3qU/uXlgZ4D0Jz+yfp70VN6Z8HBMg72Lhd28LG6n6TH35Zocga9FuLYFNkOib0yhgAiY671IbjT+CnL6ab6xGUX2fqRTuE+CO9TkHvH3/EjhzBFJioR0tTBbW4dAHDBLTnK/8FW3n8zzI81AUBFQECgYEA5Tj9FiaL+tOzr3Q+ZPNXezUvcdAoKn7XmLiVivoiRWs4yFjliGg9RtuOuhwHCdnoGA8FbrJxQrat0tpe3DGkNUbZpO8g4XJ9PGvmM/RGrZF81is3citOhLkeHd19TlQvbpQ3JIQP+Jot/wXmuH8TOOCuNySEiykqAy6RGHv/zgcCgYEA3AI9m3KAGTLf41OlpyZFBKSm6rqdEJ7ffzY+1WHUv6lxltUBiwaxxmVhs3mV+wZRTmBvq0x/NDvz7oi8hvewR89mNfmmXp5L71+vNvBL2cS/7u64FxWO0ri51OYTsSmssa/VZrd2gdNhYE9RqmteblXti4MpsI9w8RqwmIdSZhUCgYBs35yWIMjr4hG+jF/2Yv5yLtjSHiR1yrlseH/O+u+8OSlHP7IGEDzKow4vp5BQRYfPpZLW+TqCjXISbFYSECIGjBhkFpBvhImnYTh+BNBtMGUy9GPpflhVAfgkB0e1JJdC1nM6jN9pHRIrxwnHJ9Rhd5ZH6rhpwKuLx1pzXrIpkQKBgQCV1kjTB6ZKgyWMEznsntm+kczEbJfv6+PXJcChXuGgLuyXhzZn0wVcU0u5CgOlVOYm9PLYvV6c2oAFYsbTPPmnzo6UEBTIuBKLUqVoSzu9cowtVBO4VQxSY1DcuF2M+Q8ObOxv6l8hiC73gs3hvQIxLwB7Vbcc8nQ/IKhJGVB7CQKBgQCzQ5aPTVOy5C5R6Hff3NxqlqOLH+8t324CBAhDAKC5ywd4+x550zlgckztqiDedeAnxk+//Sn/2VFKLpjBy7Ktdvwwj8WlTYxksFxn+jPJ/x3YwvNONi9IDkv4hnbcQpyMI8xF3tjAHtta9WuwODWxjYZWOe5axToCy8GEKJkCVA=="
			}
	}
}
`)
