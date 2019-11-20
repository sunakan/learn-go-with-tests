package hello

import "testing"

// ファイル名 xxx_test.go
// 関数名 Testxxx
// 引数 t *testing.T
// 無理やりFailさせるなら t.Fail()
// %qというプレースホルダーはダブルクォーテーション付きでエラー所つ力する
// godoc -http :8000
// http://localhost:8000/pkg
// 余談
// go mod init github.com/sunakan/tdd
// これで現在いる場所をgithub.com/sunakan/tddと誤認させることが可能
func TestHello(t *testing.T) {
	// t.Helper()とは...?
	//
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// t.Runとは
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
