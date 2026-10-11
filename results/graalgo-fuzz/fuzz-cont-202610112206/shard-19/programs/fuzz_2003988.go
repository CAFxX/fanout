package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    var v4 uint16 = 440
    fmt.Println(v4)
    v4 = v4
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(v3, fi5)
        delete(v2, fi5)
        _ = fi5
    }
    if false {
        v4 -= 803
    }
    fmt.Println(v4)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 bool = (122 < 124)
        _ = fi6
        _ = v7
    }
    var v8 string = "a"
    switch ("a" + v8) {
    case "hello":
        v9 := []int{int(v4), -9223372036854775808, 32768}
        _ = v9
    }
    fmt.Println(v4, v8, v4)
    switch 4294967295 {
    case 4294967295:
        if true {
            fmt.Println(v4, v8, v3)
        }
    }
    v8 = (("foo bar" + "\n\t") + v8)
    v8 = v8
    fmt.Println(v3, v3, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
