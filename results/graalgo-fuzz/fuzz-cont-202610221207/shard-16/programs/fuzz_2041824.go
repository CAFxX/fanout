package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 int32, p1 bool, p2 uint16) bool {
    var v7 int32 = 1
    var v8 float32 = ((1.5 - 1e10) * 1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (!true)
}

func main() {
    v1 := []string{"hello", ((("foo bar" + ("a" + "foo bar")) + ("" + "a")) + "hello")}
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint8 = 0
    var v5 bool = (float64(0.1) >= -2.25)
    v6 := make(map[uint32]int)
    switch 0 {
    case 158:
        var v9 int8 = -47
        _ = v9
    case 93:
        v10 := GenS1{F0: -703, F1: 9223372036854775807}
        _ = v10
    case -147:
        v5 = v5
    default:
        v5 = (-236 >= 0)
    }
    var v11 string = "\n\t"
    var v12 int = -1
    var v13 int16 = -585
    switch v12 {
    case 981:
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 string = ("hello" + (v11 + ("a" + (v11 + "a"))))
            var v16 uint = 255
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint = (v16 & (v16 >> (((v16 | 43)) & 63)))
                _ = fi17
                _ = v18
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v12 = -36
    var v19 uint32 = 593
    v20 := GenS2{F0: v19, F1: v12}
    var v21 int = v12
    if false {
        var v22 int32 = 0
        var v23 int = v12
        v12 -= 127
        _ = v22
        _ = v23
    } else {
        v12 = v12
        if (-1 < 0) {
            switch 2147483647 {
            case -885:
                var v24 int8 = 127
                _ = v24
            case 910:
                _ = len(v1)
            case -2147483648:
                v12 -= v12
            default:
                v25 := GenS1{F0: (v13 + v13), F1: 9223372036854775807}
                _ = v25
            }
        }
        if len(v1) > 0 {
            v1[0] = v11
        }
    }
    var v26 float64 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
