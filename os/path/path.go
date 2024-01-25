package path

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	. "wu"
	. "wu/os/file"
)

func CurrentExePath() (string, error) {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Abs(filepath.Dir(filename))
}

func CurrentPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func FilePath(filename string) (string, error) {
	return filepath.Abs(filepath.Dir(filename))
}

// IsDir returns true if given path is a directory,
// or returns false when it's a file or does not exist.
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

func statDir(dirPath, recPath string, includeDir, isDirOnly bool) ([]string, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fis, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	statList := make([]string, 0)
	for _, fi := range fis {
		if strings.Contains(fi.Name(), ".DS_Store") {
			continue
		}

		relPath := path.Join(recPath, fi.Name())
		curPath := path.Join(dirPath, fi.Name())
		if fi.IsDir() {
			if includeDir {
				statList = append(statList, relPath+"/")
			}
			s, err := statDir(curPath, relPath, includeDir, isDirOnly)
			if err != nil {
				return nil, err
			}
			statList = append(statList, s...)
		} else if !isDirOnly {
			statList = append(statList, relPath)
		}
	}
	return statList, nil
}

// StatDir gathers information of given directory by depth-first.
// It returns slice of file list and includes subdirectories if enabled;
// it returns error and nil slice when error occurs in underlying functions,
// or given path is not a directory or does not exist.
//
// Slice does not include given path itself.
// If subdirectories is enabled, they will have suffix '/'.
func StatDir(rootPath string, includeDir ...bool) ([]string, error) {
	if !IsDir(rootPath) {
		return nil, ErrNew("not a directory or does not exist: " + rootPath)
	}

	isIncludeDir := false
	if len(includeDir) >= 1 {
		isIncludeDir = includeDir[0]
	}
	return statDir(rootPath, "", isIncludeDir, false)
}

// GetAllSubDirs returns all subdirectories of given root path.
// Slice does not include given path itself.
func GetAllSubDirs(rootPath string) ([]string, error) {
	if !IsDir(rootPath) {
		return nil, ErrNew("not a directory or does not exist: " + rootPath)
	}
	return statDir(rootPath, "", true, true)
}

// GetFileListBySuffix returns an ordered list of file paths.
// It recognize if given path is a file, and don't do recursive find.
func GetFileListBySuffix(dirPath, suffix string) ([]string, error) {
	if !IsExist(dirPath) {
		return nil, Errf("given path does not exist: %s", dirPath)
	} else if IsFile(dirPath) {
		return []string{dirPath}, nil
	}

	// Given path is a directory.
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fis, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0, len(fis))
	for _, fi := range fis {
		if strings.HasSuffix(fi.Name(), suffix) {
			files = append(files, path.Join(dirPath, fi.Name()))
		}
	}

	return files, nil
}

// CopyDir copy files recursively from source to target directory.
//
// The filter accepts a function that process the path info.
// and should return true for need to filter.
//
// It returns error when error occurs in underlying functions.
func CopyDir(srcPath, destPath string, filters ...func(filePath string) bool) error {
	// Check if target directory exists.
	if IsExist(destPath) {
		return ErrNew("file or directory alreay exists: " + destPath)
	}

	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Gather directory info.
	infos, err := StatDir(srcPath, true)
	if err != nil {
		return err
	}

	var filter func(filePath string) bool
	if len(filters) > 0 {
		filter = filters[0]
	}

	for _, info := range infos {
		if filter != nil && filter(info) {
			continue
		}

		curPath := path.Join(destPath, info)
		if strings.HasSuffix(info, "/") {
			err = os.MkdirAll(curPath, os.ModePerm)
		} else {
			err = Copy(path.Join(srcPath, info), curPath)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetGOPATHs returns all paths in GOPATH variable.
func GetGOPATHs() []string {
	gopath := os.Getenv("GOPATH")
	var paths []string
	if runtime.GOOS == "windows" {
		gopath = strings.Replace(gopath, "\\", "/", -1)
		paths = strings.Split(gopath, ";")
	} else {
		paths = strings.Split(gopath, ":")
	}
	return paths
}

// GetSrcPath returns app. source code path.
// It only works when you have src. folder in GOPATH,
// it returns error not able to locate source folder path.
func GetSrcPath(importPath string) (appPath string, err error) {
	paths := GetGOPATHs()
	for _, p := range paths {
		if IsExist(p + "/src/" + importPath + "/") {
			appPath = p + "/src/" + importPath + "/"
			break
		}
	}

	if len(appPath) == 0 {
		return "", ErrNew("Unable to locate source folder path")
	}

	appPath = filepath.Dir(appPath) + "/"
	if runtime.GOOS == "windows" {
		// Replace all '\' to '/'.
		appPath = strings.Replace(appPath, "\\", "/", -1)
	}

	return appPath, nil
}

// HomeDir returns path of '~'(in Linux) on Windows,
// it returns error when the variable does not exist.
func HomeDir() (home string, err error) {
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
		if len(home) == 0 {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		}
	} else {
		home = os.Getenv("HOME")
	}

	if len(home) == 0 {
		return "", ErrNew("Cannot specify home directory because it's empty")
	}

	return home, nil
}
