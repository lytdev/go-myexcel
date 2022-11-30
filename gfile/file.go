package gfile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type byName []os.FileInfo

// Less 文件名倒序
func (f byName) Less(i, j int) bool { return f[i].Name() > f[j].Name() }
func (f byName) Len() int           { return len(f) }
func (f byName) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// ExtName 获取文件的扩展名
func ExtName(fp string) string {
	suffix := filepath.Ext(fp)
	if suffix != "" {
		return suffix[1:]
	}
	return suffix
}

// MainName 获取文件的名称,不带后缀
func MainName(fp string) string {
	return strings.TrimSuffix(fp, filepath.Ext(fp))
}

// IsExist 判断文件或文件夹是否存在
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	if err != nil {
		return false
	}
	return true
}

// IsFile 是否是文件
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// ReplaceFileStr 替换某个文件中的字符串
func ReplaceFileStr(fp, src, target string) {
	in, err := os.Open(fp)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	out, err := os.OpenFile(fp+".mdf", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()

	br := bufio.NewReader(in)
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), src, target, -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
		index++
	}
}

// ListFileName 列出目录下文件名列表
//  @Description: 根据后缀名查找文件列表
//  @param dirPath 目录地址
//  @param needDir 是否需要拼接目录
//  @param isDescend 是否按照倒序排列
//  @return []string 文件列表
//
func ListFileName(dirPath string, needDir, isDescend bool) ([]string, error) {
	files := make([]string, 0)
	fis, err := ListFile(dirPath)
	if err != nil {
		return nil, err
	}
	if isDescend {
		sort.Sort(byName(fis))
	}
	for _, fi := range fis {
		if needDir {
			files = append(files, filepath.Join(dirPath, fi.Name()))
		} else {
			files = append(files, fi.Name())
		}
	}
	return files, nil

}

// ListFile 列出目录下文件列表(不递归)
func ListFile(dirPath string) ([]os.FileInfo, error) {
	files := make([]os.FileInfo, 0)
	if !IsExist(dirPath) {
		return nil, fmt.Errorf("given path does not exist: %s", dirPath)
	} else if IsFile(dirPath) {
		return files, nil
	}

	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, fi := range fis {
		files = append(files, fi)
	}
	return files, nil
}

// RecurveListDirFile 递归获取目录下所有的文件
func RecurveListDirFile(dirname string) ([]string, error) {
	dirname = strings.TrimSuffix(dirname, string(os.PathSeparator))
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	resultList := make([]string, 0, len(files))
	for _, fi := range files {
		path := dirname + string(os.PathSeparator) + fi.Name()
		if fi.IsDir() {
			tmp, err := RecurveListDirFile(path)
			if err != nil {
				return nil, err
			}
			resultList = append(resultList, tmp...)
			continue
		}
		resultList = append(resultList, path)
	}
	return resultList, nil
}

// FileListBySuffix
//  @Description: 根据后缀名查找文件列表
//  @param dirPath 目录地址
//  @param suffix 后缀名
//  @param needDir 是否需要拼接目录
//  @param isDescend 是否按照倒序排列
//  @param num 获取的个数
//  @return []string 文件列表
//
func FileListBySuffix(dirPath, suffix string, needDir bool, isDescend bool, num int) ([]string, error) {
	if !IsExist(dirPath) {
		return nil, fmt.Errorf("given path does not exist: %s", dirPath)
	} else if IsFile(dirPath) {
		return []string{dirPath}, nil
	}

	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	if isDescend {
		sort.Sort(byName(fis))
	}

	if num == 0 {
		num = len(fis)
	}
	files := make([]string, 0, num)
	for i := 0; i < num; i++ {
		fi := fis[i]
		if strings.HasSuffix(fi.Name(), suffix) {
			if needDir {
				files = append(files, filepath.Join(dirPath, fi.Name()))
			} else {
				files = append(files, fi.Name())
			}
		}
	}

	return files, nil
}

// FileListByPrefix
//  @Description: 根据前缀名查找文件列表
//  @param dirPath 目录地址
//  @param suffix 后缀名
//  @param needDir 是否需要拼接目录
//  @param isDescend 是否按照倒序排列
//  @param num 获取的个数
//  @return []string 文件列表
//
func FileListByPrefix(dirPath, suffix string, needDir bool, isDescend bool, num int) ([]string, error) {
	if !IsExist(dirPath) {
		return nil, fmt.Errorf("目录不存在错误: %s", dirPath)
	} else if IsFile(dirPath) {
		return []string{dirPath}, nil
	}

	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	if isDescend {
		sort.Sort(byName(fis))
	}

	if num == 0 {
		num = len(fis)
	}
	files := make([]string, 0, num)
	for i := 0; i < num; i++ {
		fi := fis[i]
		if strings.HasPrefix(fi.Name(), suffix) {
			if needDir {
				files = append(files, filepath.Join(dirPath, fi.Name()))
			} else {
				files = append(files, fi.Name())
			}
		}
	}

	return files, nil
}

// FileListByKey
//  @Description: 根据文件名关键字模糊查找文件列表
//  @param dirPath 目录地址
//  @param key 文件名关键字
//  @param needDir 是否需要拼接目录
//  @param isDescend 是否按照倒序排列
//  @param num 获取的个数
//  @return []string 文件列表
//
func FileListByKey(dirPath, key string, needDir bool, isDescend bool, num int) ([]string, error) {
	if !IsExist(dirPath) {
		return nil, fmt.Errorf("given path does not exist: %s", dirPath)
	} else if IsFile(dirPath) {
		return []string{dirPath}, nil
	}

	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	if isDescend {
		sort.Sort(byName(fis))
	}

	if num == 0 {
		num = len(fis)
	}
	files := make([]string, 0, num)
	for i := 0; i < num; i++ {
		fi := fis[i]
		if strings.Contains(fi.Name(), key) {
			if needDir {
				files = append(files, filepath.Join(dirPath, fi.Name()))
			} else {
				files = append(files, fi.Name())
			}
		}
	}

	return files, nil
}