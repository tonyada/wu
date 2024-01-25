package p7zip

import (
	. "wu"
	"wu/os/path"
	"wu/os/shell"
)

func Zip(filename, passwd string) (output string, err error) {
	if passwd != "" {
		output, err = shell.Run(Sprintf("7z a -t7z -m0=lzma -mx=9 %v.7z %v -p%v", filename, filename, passwd))
	} else {
		output, err = shell.Run(Sprintf("7z a -t7z -m0=lzma -mx=9 %v.7z %v", filename, filename))
	}
	return
}

func Unzip(filename, passwd string, isOverwrite bool) (string, error) {
	outputDir, err := path.FilePath(filename)
	if Err(err) {
		ErrExit("Extract output dir err")
	}
	if isOverwrite {
		return shell.Run(Sprintf("7z x -y -o%v %v -p%v", outputDir, filename, passwd))
	} else {
		return shell.Run(Sprintf("7z x -o%v %v -p%v", outputDir, filename, passwd))
	}
}

func UnzipHere(filename, passwd string, isOverwrite bool) (string, error) {
	if isOverwrite {
		return shell.Run(Sprintf("7z x -y -p%v", filename, passwd))
	} else {
		return shell.Run(Sprintf("7z x %v -p%v", filename, passwd))
	}
}

func UnzipTo(filename, passwd, outputPath string, isOverwrite bool) (string, error) {
	if outputPath == "" {
		var err error
		outputPath, err = path.FilePath(filename)
		if Err(err) {
			ErrExit("Extract output path err")
		}
	}

	if isOverwrite {
		return shell.Run(Sprintf("7z e -y -o%v %v -p%v", outputPath, filename, passwd))
	} else {
		return shell.Run(Sprintf("7z e -o%v %v -p%v", outputPath, filename, passwd))
	}
}
