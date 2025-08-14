package internal

import (
	"cmp"
	"fmt"
	"io/fs"
	"os"
	"path"
	"slices"
)

func ScanDir(startPath string, levels int, sort, perm, fullPath, printDir bool) string {
	st, dc, fc := searchPath(startPath, "", 0, 0, levels, sort, fullPath, perm, printDir)
	res := startPath + "\n" + st

	if printDir {
		res += fmt.Sprintf("\n%d directories,", dc)
	}else {
		res += fmt.Sprintf("\n%d directories, %d files", dc, fc)
	}

	return res
}

func searchPath(dirPath, prefix string, dirCount, fileCount, levels int, sort, fullPath, perm, printDir bool) (string, int, int) {
	if levels == 0 {
		return "", dirCount, fileCount
	}

	res := ""
	found, err := os.ReadDir(dirPath)
	if err != nil {
		return err.Error(), dirCount, fileCount
	}

	if sort {
		slices.SortFunc(found, compareByModtime)
	}

	for i, f := range found {
		
		isLast := i == len(found)-1
		branch := "├──"
		newPrefix := prefix + "│   "
		if isLast {
			branch = "└──"
			newPrefix = prefix + "    "
		}
		if f.IsDir() {
			var str string 
			dirCount++
			p := ""
			if perm {
				p = getPermissions(f)
			}
			name := f.Name()
			if fullPath{
				name = path.Join(dirPath, f.Name())
			}
			maxLevel := levels - 1

			res += fmt.Sprintf("%s%s%s %s\n", prefix, branch, p, name)
			str, dirCount, fileCount = searchPath(path.Join(dirPath, f.Name()), newPrefix, dirCount, fileCount, maxLevel, sort, fullPath, perm, printDir)
			res += str
		}else if !printDir{	
			fileCount++
			p := ""
			if perm {
				p = getPermissions(f)
			}
			name := f.Name()
			if fullPath{
				name = path.Join(dirPath, f.Name())
			}
			res += fmt.Sprintf("%s%s %s %s\n", prefix, branch, p, name)
		}
	}

	return res, dirCount, fileCount
}

func compareByModtime(a, b fs.DirEntry) int {
	infoA, _ := a.Info()
	infoB, _ := b.Info()

	timeA := infoA.ModTime()
	timeB := infoB.ModTime()

	switch {
	case timeA.After(timeB):
		return -1
	case timeA.Before(timeB):
		return 1
	default:
		return cmp.Compare(a.Name(), b.Name())
	}
}

func getPermissions(f fs.DirEntry) string {
	x, err := f.Info()
	if err != nil {
		return "[error getting permissions]"
	} else {
		return "[" + x.Mode().Perm().String() + "]"
	}
}