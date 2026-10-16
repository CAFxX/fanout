package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 float32 = ((1.5 + 1.5) * 0.1)
    var v5 uint8 = 232
    v4 += v4
    switch "\n\t" {
    case "\n\t":
        v4 += 1e10
    }
    v5 -= 42
    switch ("foo bar" + "") {
    case "hello":
        var v6 string = ""
        _ = v6
    case "a":
        if (2147483647 == 4294967295) {
            _ = len(v3)
            v5 = 0
            if true {
                for fi7 := 0; fi7 < 3; fi7++ {
                    fmt.Println(v4, v5)
                    v5 -= 128
                    _ = fi7
                }
            } else {
                v4 = 1.5
                v4 -= 0.1
            }
        } else {
            if false {
                v8 := GenS1{F0: 524, F1: -128}
                _ = v8
            }
            var v9 int8 = int8(v5)
            v5 = v5
            _ = v9
        }
    default:
        var v10 uint = 0
        _ = v10
    }
    v4 += float32(1e10)
    v4 = 3.14159
    v5 = (v5 % ((v5) | 1))
    fmt.Println(v4, v5)
    _ = len(v3)
    v4 = v4
    var v11 float32 = (v4 - 1.5)
    v11 = v11
    var v12 int32 = 0
    v12 -= v12
    v12 = (v12 / (((v12 - v12)) | 1))
    v12 = (v12 >> ((42) & 63))
    for fi13 := 0; fi13 < 14; fi13++ {
        v11 += (3.14159 / (v4 - -2.25))
        var v14 string = ""
        v4 -= v4
        _ = fi13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
