package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 int16) int16 {
    p1 -= 624
    var v7 uint32 = 2
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1(p0 uint16) uint {
    var v8 bool = ((1 != 4294967295) && true)
    var v9 uint64 = 4294967295
    v10 := GenS1{F0: 0, F1: 255}
    for fi11 := 0; fi11 < 8; fi11++ {
        v9 = (v9 << ((65535) & 63))
        _ = fi11
    }
    v12 := GenS2{F0: -486}
    _ = p0
    _ = v10
    _ = v12
    _ = v8
    _ = v9
    return 886
}

func main() {
    var v1 float32 = (0.1 / 0.1)
    var v2 uint32 = 4294967295
    v3 := make(map[string]int)
    var v4 bool = false
    var v5 int16 = -515
    var v6 uint16 = 340
    var v13 int = 0
    if (9223372036854775807 != 42) {
        v4 = v4
        var v14 uint16 = (v6 >> ((65535) & 63))
        _ = v14
    }
    v15 := GenS1{F0: 93, F1: 435}
    _ = len(v3)
    if ((!v4) || v4) {
        v4 = v4
        v3["\n\t"] = -99
        for fi16 := 0; fi16 < 5; fi16++ {
            v1 = v1
            _ = fi16
        }
    } else {
        fmt.Println(v4, v5)
        v1 -= v1
        if ((v5 / ((-1) | 1)) <= (-(v5))) {
            if ((v2 / ((v2) | 1)) >= v2) {
                var v17 float64 = 0.0
                switch "" {
                case "a":
                    delete(v3, "\n\t")
                case "\n\t":
                    v18 := []int{319}
                    _ = v18
                case "":
                    var v19 float32 = (v1 - v1)
                    _ = v19
                }
                _ = v17
            }
            v4 = (!true)
            if v4 {
                fmt.Println(v4)
            } else {
                if v4 {
                    var v20 int32 = 119
                    fmt.Println(v6, v1, v5)
                    _ = v20
                } else {
                    var v21 uint = 264
                    var v22 uint32 = v2
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: 0, F1: 102}
                v2 -= (v2 / ((1) | 1))
                _ = v23
            }
        }
    }
    switch 275 {
    case 1:
        v2 = (v2 & 1)
    }
    switch "foo bar" {
    case "a":
        v24 := GenS1{F0: 127, F1: genFn1((v6 / (((v6 | 255)) | 1)))}
        _ = v24
    case "":
        var v25 float64 = (float64((float32(-2.25) * (v1 * v1))) / (1e10 / 1e10))
        _ = v25
    case "hello":
        v26 := GenS2{F0: -2147483648}
        _ = v26
    default:
        var v27 uint64 = 0
        _ = v27
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
