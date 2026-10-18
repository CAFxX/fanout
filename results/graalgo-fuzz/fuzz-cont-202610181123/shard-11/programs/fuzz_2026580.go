package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 string) uint8 {
    var v7 float32 = 0.0
    var v8 uint8 = 1
    _ = p0
    _ = v7
    _ = v8
    return 50
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 uint64 = 4294967295
    var v4 uint8 = 169
    v5 := []uint32{426}
    v6 := make(map[int]string)
    switch -675 {
    case 11:
        var v9 uint16 = 42
        _ = v9
    case 2147483647:
        var v10 int8 = 0
        _ = v10
    case 9223372036854775807:
        if true {
            var v11 int32 = 426
            fmt.Println(v3, v3)
            _ = v11
        } else {
            v4 += v4
        }
    default:
        var v12 int8 = -128
        _ = v12
    }
    delete(v1, -178)
    var v13 int16 = -1
    switch "" {
    case "foo bar":
        switch -343 {
        case 1:
            var v14 int8 = 100
            _ = v14
        case 128:
            var v15 string = ("hello" + "\n\t")
            _ = v15
        case 2147483647:
            for fi16 := 0; fi16 < 2; fi16++ {
                v4 -= v4
                if (908 <= uint(v3)) {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        var v18 int32 = -501
                        v13 -= v13
                        v19 := GenS1{F0: (v13 ^ (v13 | (v13 << ((329) & 63)))), F1: 42, F2: 42}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi16
            }
        }
    default:
        if ("a" == (("" + "hello") + "hello")) {
            v20 := GenS2{F0: 1}
            _ = v20
        } else {
            v3 = 4294967295
        }
    }
    switch 152 {
    case 458:
        v2 = (0 < 2)
    default:
        var v21 uint8 = (^(v4))
        _ = v21
    }
    v5 = append(v5, 945)
    var v22 float32 = (0.0 - 0.1)
    v23 := GenS2{F0: -226}
    fmt.Println(v4, v13, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
