package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 int; F1 bool }
func genFn0() int {
    var v5 int32 = -42
    v6 := GenS1{F0: 0, F1: -648, F2: 32767}
    v7 := GenS1{F0: -114, F1: 225, F2: 65535}
    _ = v5
    _ = v6
    _ = v7
    return 128
}
func genFn1() int64 {
    var v8 int16 = 42
    var v9 float64 = -2.25
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483647, 42}
    var v3 int16 = 1000
    v4 := make(map[string]uint32)
    var v10 bool = false
    var v11 uint32 = 244
    v3 += (v3 | 42)
    var v12 string = "foo bar"
    _ = len(v4)
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v10 = (957 <= 0)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v10 = ((0 > -1) || (!v10))
            _ = fi14
        }
        fmt.Println(v10, v3)
        _ = fi13
    }
    switch v12 {
    case "\n\t":
        var v15 bool = (255 >= 255)
        _ = v15
    case "":
        fmt.Println(v11, v3, v3)
    case "foo bar":
        v11 += (v11 | v11)
    default:
        var v16 uint16 = 649
        _ = v16
    }
    fmt.Println(v12, v12)
    switch 4294967295 {
    case 1:
        fmt.Println(v11, v3, v11)
    case 9223372036854775807:
        v10 = v10
    case -524:
        if (1 == 65535) {
            v11 = v11
        } else {
            v12 = v12
            v11 -= 917
        }
    default:
        if true {
            if len(v2) > 0 {
                v2[0] = (v11 % ((1) | 1))
            }
            v17 := GenS1{F0: -12, F1: (v3 ^ (v3 << ((65535) & 63))), F2: 32767}
            _ = v17
        } else {
            delete(v1, v11)
            v18 := GenS2{F0: -251, F1: (356 > 38)}
            for fi19 := 0; fi19 < 14; fi19++ {
                v11 = 4294967295
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 int = fi19
                    _ = fi20
                    _ = v21
                }
                var v22 float64 = 0.0
                _ = fi19
                _ = v22
            }
            _ = v18
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
