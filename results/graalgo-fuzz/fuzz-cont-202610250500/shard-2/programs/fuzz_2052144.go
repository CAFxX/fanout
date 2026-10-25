package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{0, 127}
    var v3 int16 = -32768
    var v4 float64 = (1e10 / 3.14159)
    var v5 int = 217
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: -7, F1: -9223372036854775808, F2: 0}
    if false {
        if len(v2) > 0 {
            v2[0] = (v5 | (v5 / (((v5 * 32768)) | 1)))
        }
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint32 = 571
            _ = fi8
            _ = v9
        }
        var v10 int32 = 685
        _ = v10
    } else {
        v4 = float64((0.1 - (1.5 * float32(v4))))
        var v11 uint16 = 1
        _ = v11
    }
    if (float32(float64((-2.25 * 0.1))) != float32(float64(0.1))) {
        var v12 uint32 = 378
        v13 := GenS1{F0: 0, F1: -849, F2: 726}
        var v14 string = "a"
        _ = v12
        _ = v13
        _ = v14
    }
    if true {
        if false {
            v15 := GenS1{F0: 0, F1: -1, F2: 4294967295}
            for fi16 := 0; fi16 < 19; fi16++ {
                delete(v6, (fi16 & -1))
                _ = fi16
            }
            _ = v15
        }
        var v17 int16 = 1000
        _ = v17
    }
    fmt.Println(v3)
    var v18 uint8 = 42
    v5 += (v5 << ((42) & 63))
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := v6[(v5 & fi19)]
        v21 := GenS1{F0: -1, F1: -112, F2: v20}
        v4 = v4
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 bool = (!true)
    var v23 int = v5
    switch "foo bar" {
    case "\n\t":
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS1{F0: 95, F1: 31, F2: 4294967295}
            for fi26 := 0; fi26 < 2; fi26++ {
                if true {
                    var v27 int64 = 42
                    _ = v27
                } else {
                    v28 := GenS1{F0: 100, F1: -9223372036854775808, F2: uint32(v18)}
                    v29 := GenS1{F0: -1, F1: -1, F2: 0}
                    var v30 uint8 = v18
                    v18 = (v18 / ((v30) | 1))
                    fmt.Println(v4, v30)
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = fi26
            }
            var v31 int = v23
            _ = fi24
            _ = v25
            _ = v31
        }
    case "":
        v32 := GenS2{F0: 1, F1: (v3 ^ (v3 % (((v3 / ((42) | 1))) | 1))), F2: v3}
        _ = v32
    default:
        v3 += (v3 % ((v3) | 1))
    }
    v4 = v4
    var v33 float32 = (0.0 - 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
