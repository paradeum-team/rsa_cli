package ini



import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type IniParser struct {
	FileName string
	Data     map[string]string
}

func NewIniParser(iniFullPath string) *IniParser {
	return &IniParser{
		FileName: iniFullPath,
		Data:     map[string]string{},
	}
}
func (c *IniParser) Parse() error {
	//check file exists
	isExists,err := c.pathExists(c.FileName)
	if !isExists{
		fmt.Printf("The file not found. error=%s \n",err.Error())
		return fmt.Errorf("Parse fail. The file not found ")
	}

	return c.readLine(c.FileName,c.handler)
}

func (c *IniParser) Get(key string) string {
	return c.Data[key]
}
func (c *IniParser) handler(raw string) {
	if raw != "" && !strings.HasPrefix(raw, "//") && strings.Contains(raw, "=") { // 空行或者行注释，去掉
		raws := strings.Split(raw, "=")
		if len(raws) == 2 {
			key :=strings.Trim(raws[0]," ")
			if key !=""{
				c.Data[key] = strings.Trim(raws[1]," ") //去掉空格
			}

			//key 为空 忽略
		} else {
			//忽略
		}

	}
}

func (c *IniParser) readLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func (c *IniParser)pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
