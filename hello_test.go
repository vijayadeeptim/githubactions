package main
import "testing" 
func TestHello(t *testing.T) {
	got := HelloWelcome()
	want := "Hello Vijaya, Welcome to the first session"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}