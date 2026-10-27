package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int64 = 933
    var v2 uint64 = 2
    var v3 uint = 65535
    v4 := make(map[int]string)
    fmt.Println(v3)
    switch "a" {
    case "":
        v5 := GenS1{F0: (v2 | (^(v2))), F1: 125, F2: "a"}
        _ = v5
    default:
        v6 := GenS2{F0: -732}
        _ = v6
    }
    v3 -= v3
    var v7 float64 = 0.1
    v8 := GenS3{F0: 18446744073709551615, F1: 2}
    var v9 string = ("a" + (("" + "") + (("hello" + "\n\t") + ("foo bar" + "hello"))))
    var v10 int64 = -965
    var v11 int = -1
    var v12 float64 = -2.25
    if (1 == uint16(v3)) {
        fmt.Println(v9, v12, v7)
        v13 := GenS3{F0: 161, F1: 0}
        _ = v13
    } else {
        v1 = (v1 % ((v10) | 1))
        fmt.Println(v10, v11)
        fmt.Println(v7)
    }
    fmt.Println(v9, v7)
    var v14 uint64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
