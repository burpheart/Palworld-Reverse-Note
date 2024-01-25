package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	"os"
)

func main() {

	file, err := os.Open("Level.sav.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data interface{}
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	//printLeafKeys(data, "")
	//fmt.Println(GetValueByKey(&data, "root.properties.worldSaveData.Struct.value.Struct.CharacterSaveParameterMap.Map.value[7].key.Struct.Struct.PlayerUId.Struct.value.Guid"))
	obj, err := GetValueByKey(data, "root.properties.worldSaveData.Struct.value.Struct.CharacterSaveParameterMap.Map.value")
	if err != nil {
		panic(err)
	}
	value := reflect.ValueOf(obj)

	fmt.Println(value.Kind())
	if value.Kind() != reflect.Slice {
		fmt.Println(value.Elem().Kind())
	}

	for i := 0; i < value.Len(); i++ {
		value1 := value.Index(i)
		UUID, err := GetValueByKey(value1.Interface(), "key.Struct.Struct.PlayerUId.Struct.value.Guid")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if UUID == "00000000-0000-0000-0000-000000000000" {
			continue
		}

		RawData, err := GetValueByKey(value1.Interface(), "value.Struct.Struct.RawData.Array.value.Base.Byte.Byte")
		if err != nil {
			fmt.Println(err)
			continue
		}
		RawDataR := reflect.ValueOf(RawData)

		file, err := os.Create(UUID.(string) + ".bin")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		// 遍历整数数组并逐个写入文件

		for i := 0; i < RawDataR.Len(); i++ {
			err := binary.Write(file, binary.LittleEndian, uint8(RawDataR.Index(i).Interface().(float64)))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		cmd := exec.Command("pal-uesave.exe", "raw2json", UUID.(string)+".bin")
		// 执行命令并获取标准输出流
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}
		// 将输出内容转换为字符串并打印
		result := string(output)
		fmt.Println(result)
		var data interface{}
		err = json.Unmarshal(output, &data)
		if err != nil {
			panic(err)
		}

		NickName, err := GetValueByKey(data, "Struct.value.Struct.NickName.Str.value")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("NickName:", NickName.(string))

		Level, err := GetValueByKey(data, "Struct.value.Struct.Level.Int.value")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Level:", int64(Level.(float64)))

		Exp, err := GetValueByKey(data, "Struct.value.Struct.Exp.Int.value")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Exp:", int64(Exp.(float64)))

		MaxHP, err := GetValueByKey(data, "Struct.value.Struct.MaxHP.Struct.value.Struct.Value.Int64.value")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("MaxHP:", int64(MaxHP.(float64)))
		/*
			MaxHP, err := GetValueByKey(data, "Struct.value.Struct.MaxHP.Struct.value.Struct.Value.Int64.value")
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("MaxHP:", int64(MaxHP.(float64)))

		*/

		//printLeafKeys(data, "")
		//os.Exit(0)
	}

}

func GetValueByKey(obj any, key string) (interface{}, error) {
	keys := strings.Split(key, ".")
	value := reflect.ValueOf(obj)
	for _, k := range keys {
		// 处理数组访问
		if strings.Contains(k, "[") && strings.Contains(k, "]") {
			indexStart := strings.Index(k, "[")
			indexEnd := strings.Index(k, "]")

			arrayName := k[:indexStart]
			indexStr := k[indexStart+1 : indexEnd]
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return nil, err
			}
			if value.Kind() == reflect.Map {
				mapValue := value.MapIndex(reflect.ValueOf(arrayName))
				if !mapValue.IsValid() {
					return nil, fmt.Errorf("key not found: %s", arrayName)
				}
				value = mapValue
			} else if value.Kind() == reflect.Interface && value.Elem().Kind() == reflect.Map {
				// 处理接口类型中的map访问
				mapValue := value.Elem().MapIndex(reflect.ValueOf(arrayName))
				if !mapValue.IsValid() {
					return nil, fmt.Errorf("key not found: %s", arrayName)
				}
				value = mapValue
			} else {
				// 获取结构体字段
				field := value.FieldByName(arrayName)
				if !field.IsValid() {
					return nil, fmt.Errorf("key not found: %s", arrayName)
				}
				value = field
			}

			if value.Kind() == reflect.Slice || value.Elem().Kind() == reflect.Slice || value.Kind() == reflect.Array {
				if index < 0 || index >= value.Elem().Len() {
					return nil, fmt.Errorf("index out of range: %d", index)
				}
				value = value.Elem().Index(index)
			} else {

				return nil, fmt.Errorf("not a slice or array: %s %s", k, value.Kind())
			}
		} else {
			// 获取结构体字段
			if value.Kind() == reflect.Map {
				mapValue := value.MapIndex(reflect.ValueOf(k))
				if !mapValue.IsValid() {
					return nil, fmt.Errorf("key not found: %s", k)
				}
				value = mapValue
			} else if value.Kind() == reflect.Interface && value.Elem().Kind() == reflect.Map {
				// 处理接口类型中的map访问
				mapValue := value.Elem().MapIndex(reflect.ValueOf(k))
				if !mapValue.IsValid() {
					return nil, fmt.Errorf("key not found: %s", k)
				}
				value = mapValue
			} else {
				// 获取结构体字段
				field := value.FieldByName(k)
				if !field.IsValid() {
					return nil, fmt.Errorf("key not found: %s", k)
				}

				value = field
			}
		}
	}

	return value.Interface(), nil
}

func printLeafKeys(data interface{}, currentKey string) {
	switch value := data.(type) {
	case map[string]interface{}:
		// 如果是map类型，继续递归处理每个键值对
		if len(value) > 1000 {
			//fmt.Println(currentKey)
			return
		}
		for key, val := range value {
			newKey := key
			if currentKey != "" {
				newKey = currentKey + "." + key
			}
			if key == "Byte" {
				fmt.Println(currentKey)
				continue
			}

			if key == "MapObjectSaveData" {
				fmt.Println(currentKey)
				continue
			}
			if key == "ItemContainerSaveData" {
				fmt.Println(currentKey)
				continue
			}
			if key == "MapObjectSpawnerInStageSaveData" {
				fmt.Println(currentKey)
				continue
			}
			if key == "FoliageGridSaveDataMap" {
				fmt.Println(currentKey)
				continue
			}

			printLeafKeys(val, newKey)
		}
	case []interface{}:
		// 如果是数组类型，继续递归处理每个元素
		if len(value) > 1000 {
			//fmt.Println(currentKey)
			return
		}
		for i, val := range value {
			newKey := fmt.Sprintf("%s[%d]", currentKey, i)
			printLeafKeys(val, newKey)
		}
	case string:
	
		fmt.Println(currentKey)

	default:
		// 如果是基本类型（叶子节点），打印键
		fmt.Println(currentKey)
	}
}
