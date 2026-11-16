package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string; F2 int16 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    var v3 float64 = 1.5
    var v4 string = (("foo bar" + ("foo bar" + "foo bar")) + "hello")
    var v5 float64 = float64(0.1)
    switch "\n\t" {
    case "foo bar":
        v6 := []int{1, -329, -2147483648}
        _ = v6
    case "":
        v7 := GenS1{F0: true}
        _ = v7
    case "a":
        switch v4 {
        case "foo bar":
            v1[2] = "hello"
        }
    }
    fmt.Println(v3, v3)
    var v8 int = 2147483647
    v4 = ""
    if true {
        var v9 int64 = 110
        _ = v9
    }
    var v10 int64 = -66
    delete(v1, 0)
    var v11 int32 = -2147483648
    v12 := GenS2{F0: (v10 + (v10 | (^(v10)))), F1: (("hello" + "\n\t") + v4), F2: -32768}
    fmt.Println(v3, v8, v3)
    v10 = v10
    if true {
        var v13 int32 = (v11 % (((v11 & 0)) | 1))
        var v14 int16 = 0
        v15 := GenS3{F0: v14}
        _ = v13
        _ = v14
        _ = v15
    } else {
        v8 += v8
        v10 += v10
        if ((((v5 * (3.14159 - ((float64(3.14159) + v3) / ((v3 / (v3 / v5)) * (v3 - 1.5))))) - 1.5) - float64(1.5)) != v5) {
            v16 := []int{(v8 & 194), (v8 % ((v8) | 1)), 437, -730}
            if ((^(v11)) != 2147483647) {
                var v17 bool = (v11 > v11)
                _ = v17
            } else {
                if len(v16) > 0 {
                    v16[0] = (v8 - (v8 & v8))
                }
            }
            _ = v16
        } else {
            var v18 int8 = (^(v2))
            var v19 int64 = (v10 + v10)
            _ = v18
            _ = v19
        }
    }
    fmt.Println(v3, v4, v10)
    switch (v8 & (v8 >> ((283) & 63))) {
    case 1:
        v8 = (v8 ^ v8)
    case 9223372036854775807:
        fmt.Println(v10, v3)
    default:
        var v20 int = (v8 | (v8 + (^(v8))))
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
