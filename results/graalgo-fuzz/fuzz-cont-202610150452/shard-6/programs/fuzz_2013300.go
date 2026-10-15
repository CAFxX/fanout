package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int8; F1 uint8 }

func main() {
    var v1 float64 = (float64(3.14159) * 3.14159)
    v2 := make(map[uint32]bool)
    v3 := []uint32{4294967295, 2147483647, 1}
    v4 := []string{("foo bar" + ("hello" + ("foo bar" + ("a" + "a")))), ("" + "hello")}
    var v5 int = 128
    var v6 uint8 = 0
    if (!(true || (3.14159 < ((-2.25 - float32(v1)) * 1.5)))) {
        v1 -= ((v1 + float64(3.14159)) - (-2.25 + v1))
    }
    v7 := GenS1{F0: 406}
    for fi8 := 0; fi8 < 10; fi8++ {
        v6 = (v6 ^ (v6 - v6))
        var v9 int64 = 1
        _ = fi8
        _ = v9
    }
    v5 = (v5 - 526)
    v10 := GenS2{F0: (v6 | 255), F1: 1}
    v5 = (v5 | v5)
    if true {
        fmt.Println(v6)
        var v11 uint = 0
        _ = v11
    }
    var v12 int8 = -128
    var v13 float32 = 0.1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS1{F0: 253}
            _ = fi15
            _ = v16
        }
        v5 = (fi14 | (v5 - (v5 >> ((42) & 63))))
        _ = fi14
    }
    v17 := GenS2{F0: v6, F1: int64(v12)}
    switch (^(v5)) {
    case -9223372036854775808:
        var v18 float64 = 3.14159
        _ = v18
    case 489:
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS3{F0: (v12 & -128), F1: (v6 % ((0) | 1))}
            _ = fi19
            _ = v20
        }
    case -2147483648:
        var v21 int8 = (v12 | (v12 - (^(v12))))
        _ = v21
    default:
        var v22 uint8 = v6
        _ = v22
    }
    var v23 bool = (!true)
    switch -396 {
    case 4294967295:
        v4 = append(v4, (("a" + "") + "a"))
    case 127:
        var v24 string = ("" + "")
        _ = v24
    case 32767:
        v5 += 1
    }
    v13 -= v13
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
