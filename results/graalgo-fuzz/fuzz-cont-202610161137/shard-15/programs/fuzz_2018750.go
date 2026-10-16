package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 int16) int32 {
    if true {
        switch "" {
        case "hello":
            v6 := []int{32767, 9223372036854775807, 640, 2147483647}
            _ = v6
        }
        v7 := []int{42, -709, -432, 2147483647}
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
        _ = v7
    } else {
        v8 := GenS1{F0: -716}
        p0 = 0
        if false {
            for fi9 := 0; fi9 < 10; fi9++ {
                p0 = p0
                _ = fi9
            }
            v10 := GenS2{F0: "hello"}
            _ = v10
        }
        _ = v8
    }
    v11 := GenS3{F0: 744, F1: 173}
    _ = p0
    _ = v11
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 int = 92
    var v4 int8 = 84
    var v5 int32 = -2147483648
    _ = len(v1)
    v12 := GenS4{F0: 2, F1: (v5 | v5), F2: int32(v3)}
    fmt.Println(v5)
    if true {
        var v13 int = (v3 | 1)
        _ = v13
    }
    var v14 float64 = (0.1 + float64(((1e10 + (3.14159 + 1e10)) + (3.14159 / 1.0))))
    var v15 bool = ("a" >= "hello")
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v4)
        fmt.Println(v4, v14, fi16)
        var v17 int8 = v4
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v3, v4)
    switch -2147483648 {
    case 127:
        var v18 int64 = -1
        _ = v18
    case 1:
        v19 := []int{-9223372036854775808, v3}
        _ = v19
    case 128:
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 int16 = 32767
            v22 := GenS3{F0: 4294967295, F1: 42}
            v2[-809] = 810
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        v23 := GenS3{F0: 2147483648, F1: -9223372036854775808}
        _ = v23
    }
    v1[65535] = "foo bar"
    var v24 uint16 = 65535
    for fi25 := 0; fi25 < 6; fi25++ {
        fmt.Println(v3, v15)
        _ = fi25
    }
    var v26 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
