package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 bool = ((true || (18446744073709551615 < 4294967295)) && (!(262 >= 1)))
    v2 := make(map[string]uint32)
    var v3 int64 = -1
    v4 := []int{32768, 128}
    v3 += (v3 | (v3 & v3))
    v1 = v1
    v5 := GenS1{F0: 42, F1: 172, F2: true}
    for fi6 := 0; fi6 < 15; fi6++ {
        _ = len(v4)
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 float64 = 1e10
            v1 = (840 == 42)
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            _ = len(v4)
            v4 = append(v4, (fi9 + fi6))
            _ = fi9
        }
        _ = fi6
    }
    v10 := GenS1{F0: 42, F1: 898, F2: v1}
    v1 = (("foo bar" + "hello") > ("a" + "\n\t"))
    v11 := GenS1{F0: 1, F1: 4294967295, F2: v1}
    v4 = append(v4, 0)
    var v12 uint8 = 84
    v3 -= (^(v3))
    v2[("" + "a")] = 65535
    if false {
        var v13 uint64 = 4294967295
        if v1 {
            if false {
                switch 973 {
                case 0:
                    v4 = append(v4, 42)
                case 2147483647:
                    for fi14 := 0; fi14 < 16; fi14++ {
                        var v15 int64 = (v3 >> ((0) & 63))
                        _ = fi14
                        _ = v15
                    }
                default:
                    v12 -= v12
                }
                var v16 string = "hello"
                _ = v16
            }
            if (v3 > 9223372036854775807) {
                var v17 float32 = ((0.0 * ((((float32(0.0) + 3.14159) * 1.5) + (1.5 * (-2.25 * (-2.25 * (3.14159 * 0.1))))) - 1.5)) - -2.25)
                fmt.Println(v12)
                if (!(18446744073709551615 == 398)) {
                    v12 += v12
                    for fi18 := 0; fi18 < 3; fi18++ {
                        v19 := GenS1{F0: 1, F1: (^(v13)), F2: v1}
                        var v20 bool = false
                        fmt.Println(fi18, v17)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v17 -= v17
                }
                _ = v17
            }
            var v21 int64 = 1
            _ = v21
        }
        _ = v13
    }
    v1 = v1
    v22 := GenS2{F0: uint64(v12)}
    if v1 {
        _ = len(v2)
        var v23 int16 = 1
        _ = v23
    }
    var v24 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
