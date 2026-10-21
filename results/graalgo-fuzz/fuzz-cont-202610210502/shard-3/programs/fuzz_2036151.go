package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ""
    v3 := []string{}
    var v4 int16 = -152
    fmt.Println(v4, v4)
    v5 := v1[0]
    var v6 uint64 = 357
    switch 32768 {
    case 682:
        fmt.Println(v4, v5)
    case 9223372036854775807:
        _ = len(v1)
    case 127:
        var v7 int32 = 490
        _ = v7
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: (v5 << ((4294967295) & 63)), F1: 2, F2: 42}
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: v5, F1: 127, F2: (v6 - (v6 * (v6 << ((1) & 63))))}
    v11 := GenS1{F0: v5, F1: 127, F2: v6}
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v6, v6, v4)
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: 451, F1: 1, F2: (v6 | 4294967295)}
            if (!true) {
                fmt.Println(v6, fi13, fi13)
                v5 -= (^(v5))
            }
            _ = fi13
            _ = v14
        }
        var v15 int64 = 0
        _ = fi12
        _ = v15
    }
    if false {
        var v16 float64 = (0.0 / -2.25)
        for fi17 := 0; fi17 < 11; fi17++ {
            v5 += 729
            _ = fi17
        }
        _ = v16
    }
    switch "" {
    case "":
        if (468 != 154) {
            v18 := GenS1{F0: v5, F1: 255, F2: (v6 - v6)}
            _ = v18
        }
    case "foo bar":
        for fi19 := 0; fi19 < 2; fi19++ {
            v4 = v4
            switch (fi19 & fi19) {
            case -977:
                if false {
                    v5 += v5
                    v20 := GenS1{F0: (v5 ^ v5), F1: 0, F2: 1}
                    var v21 int64 = -1
                    _ = v20
                    _ = v21
                } else {
                    v1[uint32(v6)] = (v5 + 65535)
                    v6 = v6
                }
            case 127:
                var v22 uint8 = 67
                _ = v22
            case -2147483648:
                if true {
                    v23 := GenS1{F0: (v5 << ((255) & 63)), F1: 2, F2: v6}
                    _ = v23
                }
            default:
                var v24 float32 = ((-2.25 / 0.1) / float32(-2.25))
                _ = v24
            }
            var v25 float32 = -2.25
            _ = fi19
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v6 = 18446744073709551615
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
