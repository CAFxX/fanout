package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 uint16, p1 uint32) int8 {
    v6 := GenS1{F0: -1}
    v7 := GenS2{F0: ((((((("" + ("" + "")) + "foo bar") + "hello") + "foo bar") + "foo bar") + "\n\t") + ("\n\t" + "hello")), F1: uint16(p1), F2: -9223372036854775808}
    v8 := GenS2{F0: "a", F1: p0, F2: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{1}
    var v2 uint64 = 42
    var v3 bool = true
    var v4 float32 = 3.14159
    v5 := make(map[int]uint32)
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v3)
        var v10 bool = (-1 >= -479)
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: "a", F1: 702, F2: -769}
    v12 := GenS1{F0: -1}
    fmt.Println(v3)
    v2 = (^(v2))
    var v13 float64 = (0.1 - 1.5)
    var v14 string = ""
    fmt.Println(v3, v14)
    v4 = (float32(((0.0 - v13) + float64(v4))) + v4)
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 = (v2 + v2)
        fmt.Println(v3)
        _ = fi15
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v16 uint8 = 128
        if v3 {
            _ = len(v1)
            if false {
                v17 := GenS2{F0: ("foo bar" + v14), F1: 1, F2: 9223372036854775807}
                _ = v17
            }
            v18 := GenS1{F0: 42}
            _ = v18
        } else {
            switch v14 {
            case "foo bar":
                v13 = 0.1
            default:
                fmt.Println(v2)
            }
            v14 = v14
        }
        v14 = v14
        _ = v16
    }
    v5[9223372036854775807] = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
