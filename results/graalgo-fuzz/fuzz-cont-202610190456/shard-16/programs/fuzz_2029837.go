package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
func genFn0(p0 bool, p1 bool, p2 int16) int16 {
    fmt.Println(p1, p0)
    p0 = false
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{-439}
    v3 := make(map[uint32]string)
    var v4 uint8 = 42
    v5 := []uint32{142, 42, uint32(v4)}
    v6 := GenS1{F0: -69, F1: 32768, F2: v4}
    var v7 float64 = 1e10
    v3[644] = "\n\t"
    v1 = v1
    fmt.Println(v4)
    if (32768 > 4294967295) {
        v4 -= v4
    } else {
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 bool = (42 >= -128)
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: -1, F1: fi8, F2: 177}
            _ = fi8
            _ = v11
        }
    }
    if (42 != -128) {
        for fi12 := 0; fi12 < 19; fi12++ {
            for fi13 := 0; fi13 < 10; fi13++ {
                switch ("foo bar" + (("\n\t" + ("hello" + "foo bar")) + "\n\t")) {
                case "hello":
                    var v14 int = (fi12 / ((676) | 1))
                    _ = v14
                case "\n\t":
                    if false {
                        v15 := v3[65535]
                        _ = v15
                    } else {
                        fmt.Println(fi12, v1, fi12)
                    }
                default:
                    v4 = v4
                }
                v7 = 3.14159
                _ = fi13
            }
            _ = fi12
        }
    }
    if (255 < 42) {
        for fi16 := 0; fi16 < 8; fi16++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v17 uint8 = (v4 >> ((18446744073709551615) & 63))
            _ = fi16
            _ = v17
        }
    }
    if len(v2) > 0 {
        v2[0] = -814
    }
    fmt.Println(v7, v1)
    var v18 string = (("\n\t" + "foo bar") + "")
    v7 -= float64(0.0)
    if ((1.5 + -2.25) < (1e10 + -2.25)) {
        var v19 uint64 = 1
        switch 32768 {
        case 32767:
            if false {
                fmt.Println(v7, v19, v7)
            }
        case 4294967295:
            switch 49 {
            case 848:
                v20 := GenS1{F0: 127, F1: 12, F2: 1}
                _ = v20
            case -8:
                v21 := GenS1{F0: 124, F1: 42, F2: v4}
                _ = v21
            default:
                switch 294 {
                case 65:
                    var v22 bool = false
                    _ = v22
                default:
                    var v23 bool = true
                    _ = v23
                }
            }
        default:
            var v24 float64 = v7
            _ = v24
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
