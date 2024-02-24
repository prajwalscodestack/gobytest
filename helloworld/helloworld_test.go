package helloworld

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello World"
	got := SayHello()
	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestSayHelloToPerson(t *testing.T) {
	t.Run("testing hello to person with passing argument", func(t *testing.T) {
		person := "PrajwalP"
		want := "Hello " + person
		got := SayHelloToPerson(person)
		assertCorrectMessage(t, got, want)
	})
	t.Run("testing hello to person without passing empty argument", func(t *testing.T) {
		want := "Hello World"
		got := SayHelloToPerson("")
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}
