package integer
import(
  "testing"
)
func TestAddInteger(t *testing.T){
 t.Run("normal case",func(t *testing.T){
    want:=10
    got:=AddInteger(5,5)
    if want!=got{
     t.Errorf("got %d want %d",got,want)
    }
  }) 
}

