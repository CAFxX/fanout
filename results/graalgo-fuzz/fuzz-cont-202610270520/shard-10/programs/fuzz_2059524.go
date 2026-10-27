package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }

func main() {
    var v1 uint8 = 127
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    switch 42 {
    case 973:
        v4 := GenS1{F0: 582, F1: 32768, F2: 42}
        _ = v4
    case 0:
        v5 := GenS1{F0: 32767, F1: 32767, F2: 0}
        _ = v5
    case -2147483648:
        fmt.Println(v2)
    }
    switch "hello" {
    case "\n\t":
        var v6 float64 = (float64(3.14159) * (-2.25 / (-2.25 / 1.0)))
        _ = v6
    }
    var v7 uint64 = 18446744073709551615
    var v8 string = "\n\t"
    if true {
        v2 += v2
        switch int(v7) {
        case -550:
            v1 = (^(v1))
        }
    } else {
        switch 607 {
        case 128:
            var v9 string = v8
            _ = v9
        case 4294967295:
            v8 = "\n\t"
        }
        v1 = (v1 | (v1 % ((v1) | 1)))
        fmt.Println(v8)
    }
    v2 = v2
    var v10 bool = false
    var v11 uint16 = 1
    var v12 string = v8
    var v13 uint64 = (v7 & v7)
    fmt.Println(v1, v11)
    if (v10 && false) {
        var v14 uint16 = (v11 - uint16(v1))
        var v15 uint8 = v1
        v16 := []int{-895, 0, 2147483647, 0}
        _ = v14
        _ = v15
        _ = v16
    }
    v10 = (((v12 + v8) + ((v12 + (("\n\t" + "hello") + "a")) + v12)) > "")
    switch ("hello" + (v8 + (v12 + "a"))) {
    case "":
        var v17 uint64 = (v13 / (((v7 * v13)) | 1))
        _ = v17
    case "hello":
        v11 = v11
    default:
        v11 += 304
    }
    var v18 uint64 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
