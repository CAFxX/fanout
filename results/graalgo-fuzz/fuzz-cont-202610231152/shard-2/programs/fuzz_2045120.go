package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 456
    v3 := []string{"", "", "hello"}
    v4 := []int{819, 810, 128, 2147483647}
    var v5 bool = (113 == -127)
    if true {
        v6 := v1[2147483647]
        _ = v6
    }
    switch "" {
    case "":
        v5 = v5
    default:
        fmt.Println(v5)
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v5, v5, v5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v8 bool = (((-(fi7)) != fi7) && false)
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 189
    var v10 int16 = 42
    v10 = 32767
    var v11 int32 = 42
    delete(v1, -1)
    v9 = (v9 | v9)
    if (91 > uint64(v2)) {
        v12 := GenS1{F0: "hello", F1: -1}
        if (3.14159 >= 1e10) {
            v10 = v10
            v13 := GenS2{F0: v9, F1: (v2 - v2)}
            fmt.Println(v2)
            _ = v13
        }
        v3 = append(v3, (("\n\t" + ("" + "\n\t")) + "hello"))
        _ = v12
    }
    v2 += (v2 ^ (v2 | (v2 + (v2 * (v2 * uint16(v9))))))
    v2 += 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
