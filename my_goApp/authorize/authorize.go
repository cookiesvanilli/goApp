package authorize
import
	"github.com/cookiesvanilli/goApp/my_goApp/authorize/token"
func Auth() *token.Token {
tok := token.New()
//token to base
return tok
}