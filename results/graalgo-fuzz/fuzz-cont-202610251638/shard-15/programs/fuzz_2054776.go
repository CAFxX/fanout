package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
func genFn0() uint32 {
    var v5 int = 32767
    var v6 float64 = 0.0
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    var v4 uint = (v1 | v1)
    fmt.Println(v4, v1)
    _ = len(v2)
    var v7 uint = (v4 | (v1 ^ 0))
    delete(v2, 65535)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: v4}
        v2[42] = 2
        var v10 int16 = -686
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float64 = 0.0
    v7 = v7
    if ((true && (-289 <= -647)) && ((4294967295 >= 18446744073709551615) && ((("hello" + "hello") + "\n\t") >= (("\n\t" + ("foo bar" + "")) + "foo bar")))) {
        v2[2] = 2147483648
        var v12 int32 = 258
        _ = v12
    }
    v1 = v1
    if false {
        for fi13 := 0; fi13 < 20; fi13++ {
            v1 += (v7 / (((v4 ^ (v7 | (v1 % ((v1) | 1))))) | 1))
            v14 := GenS2{F0: 127, F1: -4, F2: 1}
            _ = fi13
            _ = v14
        }
        switch 32767 {
        case 42:
            fmt.Println(v4, v7, v4)
        case 127:
            v15 := GenS1{F0: v7}
            _ = v15
        default:
            switch 4294967295 {
            case 32768:
                fmt.Println(v7, v1)
            case 2147483647:
                var v16 uint16 = 42
                _ = v16
            case -70:
                var v17 int8 = 100
                _ = v17
            default:
                for fi18 := 0; fi18 < 11; fi18++ {
                    v7 = v4
                    var v19 int8 = -1
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v1 = (^(v4))
                        _ = fi20
                    }
                    _ = fi18
                    _ = v19
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
