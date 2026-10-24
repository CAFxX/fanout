package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
func genFn0(p0 int32) int16 {
    fmt.Println(p0)
    p0 = (^(p0))
    if true {
        v4 := []int{127}
        _ = v4
    }
    _ = p0
    return 251
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{4294967295, 65535, 1}
    var v3 float32 = 1e10
    if (!false) {
        var v5 int = 213
        v5 = v5
        _ = v5
    }
    v3 += 1e10
    var v6 int8 = 100
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    if ((-1 < 0) && true) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v1 += (v1 - 1)
            var v8 uint16 = 32767
            _ = fi7
            _ = v8
        }
    }
    v3 = (1e10 * -2.25)
    v6 = 42
    for fi9 := 0; fi9 < 14; fi9++ {
        v1 = 18446744073709551615
        _ = fi9
    }
    switch 1 {
    case 1:
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: 4294967295, F1: 401, F2: 649}
            v6 = 1
            switch (fi10 + (-(fi10))) {
            case 32768:
                switch "a" {
                case "hello":
                    switch fi10 {
                    case -956:
                        v12 := GenS1{F0: 2147483647, F1: 957, F2: 2147483648}
                        _ = v12
                    case 4294967295:
                        var v13 int16 = 372
                        _ = v13
                    case 9223372036854775807:
                        var v14 uint8 = 127
                        _ = v14
                    }
                }
            case -1:
                fmt.Println(v3)
            }
            _ = fi10
            _ = v11
        }
    default:
        v1 = (v1 << ((18446744073709551615) & 63))
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v6 = (v6 & v6)
        fmt.Println(v6, v6, v3)
        if false {
            v16 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 1}
            v3 -= (1e10 * (0.1 * v3))
            v1 -= 644
            _ = v16
        }
        _ = fi15
    }
    switch 582 {
    case 4294967295:
        var v17 int32 = 646
        _ = v17
    default:
        var v18 float32 = (v3 / -2.25)
        _ = v18
    }
    v3 = v3
    v1 += 0
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float64 = (float64(float32(0.1)) / (1e10 + (1e10 * 0.0)))
        v1 = (^(v1))
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
