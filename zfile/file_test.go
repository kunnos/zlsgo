package zfile

import (
	"os"
	"sync"
	"testing"
	"time"

	. "github.com/sohaha/zlsgo"
)

func TestFile(T *testing.T) {
	t := NewTest(T)

	filePath := "../doc.go"
	tIsFile := FileExist(filePath)
	t.Equal(true, tIsFile)

	notPath := "zlsgo.php"
	status, _ := PathExist(notPath)
	t.Equal(0, status)

	size := FileSize("../doc.go")
	t.Equal("0 B" != size, true)

	size = FileSize("../_doc.go")
	t.Equal("0 B" == size, true)

	dirPath := RealPathMkdir("../zfile")
	tIsDir := DirExist(dirPath)
	t.Equal(true, tIsDir)

	dirPath = SafePath("../zfile/ok")
	t.Equal("ok", dirPath)

	tmpPath := TmpPath()
	t.EqualTrue(tmpPath != "")

	path := RealPathMkdir("../tmp")
	path2 := RealPathMkdir(path + "/ooo")
	T.Log(path, path2)
	t.Equal(true, Rmdir(path, true))
	t.Equal(true, Rmdir(path))
	ePath := ProgramPath(true)
	ProjectPath = ePath
	path = RealPathMkdir("../ppppp")
	testPath := ePath + "../ppppp"
	T.Log(path, testPath)
	t.EqualTrue(DirExist(path))
	t.EqualTrue(DirExist(testPath))
	ok := Rmdir(testPath)

	T.Log(path, testPath, ok)
	var g sync.WaitGroup
	g.Add(1)
	// g.Wait()

}

func TestPut(t *testing.T) {
	var err error
	tt := NewTest(t)
	defer os.Remove("./text.txt")
	err = PutOffset("./text.txt", []byte(time.Now().String()+"\n"), 0)
	tt.EqualNil(err)
	err = PutAppend("./text.txt", []byte(time.Now().String()+"\n"))
	tt.EqualNil(err)
	_ = os.Remove("./text.txt")
	err = PutAppend("./put/text.txt", []byte(time.Now().String()+"\n"))
	tt.EqualNil(err)
	_ = os.Remove("./put/text.txt")
	err = PutOffset("./text.txt", []byte("\n(ok)\n"), 5)
	tt.EqualNil(err)
}
