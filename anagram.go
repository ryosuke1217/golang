package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    filename := "text.txt"

    // ファイルオープン
    fp, err := os.Open(filename)
    if err != nil {
        // エラー処理
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)
    str_list := []string{}

    for scanner.Scan() {
        // ここで一行ずつ処理
        str_list = append(str_list, scanner.Text())
    }

    if err = scanner.Err(); err != nil {
        // エラー処理
    }
    fmt.Println(del_list(str_list, 0))
    //arrangement(str_list)
}

func anagram(str string) []string{
    if(len(str) == 1){
        return []string{str}
    }

    ret := []string{}
    for i := 0; i < len(str); i++ {
        anagrams := anagram(str[0:i] + str[i+1:len(str)])
        for _, anagram := range anagrams {
             ret = append(ret, str[i:i+1] + anagram)
        }
    }
    return ret
}

func arrangement(strs []string) {
    ret := []string{}
    for {
        ret = del_list(str_list, 0)
        



    }
    for _, str := range strs {
        words := anagram(str)
        ret = make_list(words, strs)
        fmt.Println(ret)
    }
}

func make_list(words []string, strs []string) []string{
    ret := []string{}
    for _, word := range words {
        for _, str := range strs {
            if word == str {
                ret = append(ret, word)
            }
        }
    }
    return ret
}

func del_list(words []string, index int) []string{
    ret := []string{}
    ret = append(ret, words[0:index]...)
    ret = append(ret, words[index + 1:]...)
    return ret
}