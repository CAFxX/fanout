package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint8 }

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 0
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]uint32)
    var v5 int64 = (^(v3))
    var v6 int16 = 1
    for fi7 := 0; fi7 < 12; fi7++ {
        delete(v4, v2)
        var v8 int = fi7
        var v9 string = ""
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 42 {
    case -1:
        switch "\n\t" {
        case "foo bar":
            var v10 bool = false
            _ = v10
        case "a":
            v1 = v1
        default:
            if (127 != 1) {
                switch 1 {
                case 319:
                    for fi11 := 0; fi11 < 11; fi11++ {
                        v1 -= v1
                        _ = fi11
                    }
                case 42:
                    if (!(((-2.25 / float32(-2.25)) >= -2.25) && (!(((float32(float64(-2.25)) + 0.0) - float32((1e10 / 3.14159))) > (0.1 * 3.14159))))) {
                        v12 := GenS1{F0: -616, F1: ("hello" > ""), F2: 695}
                        _ = v12
                    }
                default:
                    fmt.Println(v3, v2, v3)
                }
                var v13 uint32 = (v2 * (v2 << ((4294967295) & 63)))
                v14 := v4[(v13 | 0)]
                _ = v13
                _ = v14
            }
        }
    case 42:
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := v4[v2]
                fmt.Println(fi15, fi16)
                _ = fi16
                _ = v17
            }
            v18 := GenS2{F0: ("hello" + "\n\t")}
            _ = fi15
            _ = v18
        }
    }
    switch "a" {
    case "hello":
        var v19 bool = ((v1 + v1) < 1)
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS1{F0: 2147483647, F1: true, F2: 0}
            var v22 float32 = -2.25
            _ = fi20
            _ = v21
            _ = v22
        }
    case "":
        v2 += (v2 % ((v2) | 1))
    default:
        v2 = v2
    }
    v23 := GenS1{F0: -800, F1: true, F2: 255}
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS2{F0: "hello"}
            v6 = 1000
            _ = fi25
            _ = v26
        }
        var v27 int = fi24
        _ = fi24
        _ = v27
    }
    v3 += v3
    var v28 uint16 = 255
    if false {
        var v29 uint = 822
        var v30 int32 = 2147483647
        v31 := GenS3{F0: 0, F1: 0}
        _ = v29
        _ = v30
        _ = v31
    }
    v32 := []int{-886}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
