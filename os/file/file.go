package file

import (
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strings"
	. "wu"
	"wu/os/user"
)

type ByteSize int64

const (
	_           = iota // ignore 0
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const (
	Byte  = 1
	KByte = Byte * 1024
	MByte = KByte * 1024
	GByte = MByte * 1024
	TByte = GByte * 1024
	PByte = TByte * 1024
	EByte = PByte * 1024
)

func Create(name string) (*os.File, error) {
	return os.Create(name)
}

func CreateHomeDotDir(dirName string) string {
	userHomeDir, err := user.Home()
	ErrFatal(err)
	dotConfigDir := userHomeDir + "/." + dirName // $HOME/.dirName
	if IsNotExist(dotConfigDir) {
		_ = os.Mkdir(dotConfigDir, 0755)
	}
	return dotConfigDir
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsNotExist(path string) bool {
	return !IsExist(path)
}

func IsFileExist(filename string) bool {
	return !IsFileNotExist(filename)
}
func IsFileNotExist(filename string) bool {
	f, err := os.Open(filename)
	if err != nil {
		return os.IsNotExist(err)
	}
	defer f.Close()
	return os.IsNotExist(err)
}

func Read(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
func ReadStr(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	return string(bytes), err
}

func ContainsStr(filename string, keyword string) bool {
	body, _ := Read(filename)
	return strings.Contains(string(body), keyword)
}

func Save(filename string, body string) (int, error) {
	f, err := os.Create(filename)
	if Err(err) {
		return 0, err
	}
	defer f.Close()
	return f.WriteString(body)
}

func Remove(name string) error {
	return os.Remove(name)
}

func RemoveAll(name string) error {
	return os.Remove(name)
}

func Del(name string) error {
	return os.Remove(name)
}

func DelAll(name string) error {
	return os.RemoveAll(name)
}

func Mkdir(dir string) error {
	return os.Mkdir(dir, 0777)
}

func MkdirAll(dir string) error {
	return os.MkdirAll(dir, 0777)
}

func RelpaceContent(filename string, o string, n string) error {
	body, err := Read(filename)
	if Err(err) {
		return err
	}
	_, err = Save(filename, strings.Replace(string(body), o, n, -1))
	return err
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return Sprintf(f+"%s", val, suffix)
}

// HumaneFileSize calculates the file size and generate user-friendly string.
func HumaneFileSize(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1024, sizes)
}

// FileMTime returns file modified time and possible error.
func FileMTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// FileSize returns file size in bytes and possible error.
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

// Copy copies file from source to target path.
func Copy(src, dest string) error {
	// Gather file information to set back later.
	si, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Handle symbolic link.
	if si.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(src)
		if err != nil {
			return err
		}
		// NOTE: os.Chmod and os.Chtimes don't recoganize symbolic link,
		// which will lead "no such file or directory" error.
		return os.Symlink(target, dest)
	}

	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	if _, err = io.Copy(dw, sr); err != nil {
		return err
	}

	// Set back file information.
	if err = os.Chtimes(dest, si.ModTime(), si.ModTime()); err != nil {
		return err
	}
	return os.Chmod(dest, si.Mode())
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it
// and its upper level paths.
func WriteFile(filename string, data []byte) error {
	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if Err(err) {
		return err
	}
	return ioutil.WriteFile(filename, data, 0655)
}

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}
