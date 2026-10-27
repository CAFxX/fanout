package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 bool }
func genFn0(p0 string, p1 int64, p2 int32) bool {
    var v6 int16 = -392
    v7 := GenS1{F0: 4294967295}
    v8 := GenS1{F0: 428}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (!false)
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    var v3 uint8 = 0
    var v4 int32 = 950
    var v5 float32 = (0.1 / 1.0)
    var v9 uint64 = 701
    v10 := GenS2{F0: 1, F1: 32768, F2: -1}
    var v11 float32 = float32(((3.14159 + float64(1e10)) + (0.1 - 1e10)))
    var v12 uint32 = 196
    v13 := GenS1{F0: 65535}
    v14 := v1[("hello" + ((("hello" + "a") + ("a" + "foo bar")) + ("foo bar" + "")))]
    v15 := GenS1{F0: v12}
    var v16 uint = 4294967295
    v17 := GenS3{F0: 2, F1: 2, F2: (false && (true && (((((3.14159 + ((3.14159 * 3.14159) / 3.14159)) - -2.25) - float64(v5)) / 1.0) <= 1.5)))}
    if false {
        fmt.Println(v9, v14)
    }
    fmt.Println(v3)
    v18 := GenS1{F0: (v12 / ((4294967295) | 1))}
    if false {
        v12 = (v12 % (((v12 << (((v16 + v16)) & 63))) | 1))
        v4 -= (v4 + -818)
        var v19 uint64 = (v9 / ((v9) | 1))
        _ = v19
    }
    var v20 int16 = -32768
    switch v14 {
    case 512:
        fmt.Println(v4, v16)
    case 926:
        var v21 int8 = (v2 | (v2 & 0))
        _ = v21
    default:
        var v22 string = "hello"
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
