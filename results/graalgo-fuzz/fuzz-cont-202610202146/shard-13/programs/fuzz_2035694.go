package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint8; F2 int64 }
func genFn0(p0 int16) uint64 {
    p0 -= (p0 % ((42) | 1))
    p0 -= -526
    if true {
        p0 -= p0
    } else {
        p0 = p0
    }
    _ = p0
    return 1
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = 0.0
    v3 := []uint32{v1, (v1 + (v1 ^ 65535))}
    var v4 uint16 = 0
    v5 := []string{"", ("\n\t" + "hello")}
    var v6 int32 = 42
    if (!true) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if (128 != 0) {
        for fi7 := 0; fi7 < 1; fi7++ {
            v4 += (v4 & v4)
            _ = fi7
        }
    }
    v8 := GenS1{F0: -128}
    var v9 uint = 511
    if true {
        var v10 uint32 = v1
        switch 9223372036854775807 {
        case -1:
            v11 := GenS1{F0: -128}
            _ = v11
        case 165:
            if true {
                for fi12 := 0; fi12 < 12; fi12++ {
                    switch fi12 {
                    case -612:
                        v10 += (^(v10))
                    case 4294967295:
                        v10 -= (v1 ^ v1)
                    }
                    v9 += v9
                    v4 += (v4 - 42)
                    _ = fi12
                }
                var v13 float64 = (((float64((0.1 + 3.14159)) - v2) * float64(0.1)) * 3.14159)
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 15; fi14++ {
                v2 += v2
                _ = fi14
            }
        }
        v15 := GenS2{F0: 42, F1: 118, F2: 1}
        _ = v10
        _ = v15
    }
    var v16 int32 = 42
    var v17 float32 = 3.14159
    v4 += v4
    var v18 bool = (!true)
    v9 += v9
    var v19 string = ("" + "\n\t")
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
