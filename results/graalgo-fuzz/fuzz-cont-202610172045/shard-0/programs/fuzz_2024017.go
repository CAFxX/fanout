package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int32, p1 uint64, p2 uint64) int16 {
    v4 := []int{-2147483648, -2147483648, -636, -147}
    p2 -= (p1 % ((p2) | 1))
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 671, F1: 0, F2: (p0 ^ p0)}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (p0 + (p0 | (p0 << ((18446744073709551615) & 63)))), F1: 81, F2: p0}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 929
}

func main() {
    var v1 uint8 = 1
    var v2 bool = true
    var v3 int16 = 1
    fmt.Println(v2, v2, v3)
    v8 := []int{-409, -432}
    fmt.Println(v3, v1)
    v2 = ((255 > 1) && true)
    fmt.Println(v1, v2, v3)
    switch -520 {
    case 32768:
        var v9 uint16 = uint16(v1)
        _ = v9
    default:
        fmt.Println(v3, v2, v2)
    }
    var v10 string = ("\n\t" + "a")
    v11 := GenS1{F0: 0, F1: 0, F2: -2147483648}
    _ = len(v8)
    if (false || (v1 <= v1)) {
        if false {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(fi12)
                _ = fi12
            }
        } else {
            if (v2 || (2 < uint64(v1))) {
                for fi13 := 0; fi13 < 19; fi13++ {
                    fmt.Println(v3, v2, v2)
                    switch (fi13 % ((32768) | 1)) {
                    case -1:
                        var v14 uint8 = v1
                        _ = v14
                    case 32768:
                        if len(v8) > 0 {
                            v8[0] = fi13
                        }
                    case 4294967295:
                        v10 = ""
                    }
                    _ = fi13
                }
                fmt.Println(v1, v10)
            } else {
                switch -1 {
                case 127:
                    switch 393 {
                    case -915:
                        var v15 float32 = float32(float64((((0.0 / (3.14159 + 1e10)) / 1e10) * (0.0 - 0.0))))
                        _ = v15
                    }
                case 128:
                    switch 9223372036854775807 {
                    case -766:
                        var v16 int = 866
                        _ = v16
                    case 0:
                        var v17 int64 = 728
                        _ = v17
                    }
                }
                v3 += v3
            }
        }
        fmt.Println(v2, v2)
        var v18 uint16 = 42
        _ = v18
    }
    switch v10 {
    case "hello":
        fmt.Println(v10, v1, v1)
    case "a":
        v19 := GenS1{F0: 42, F1: -9, F2: 264}
        _ = v19
    case "":
        v20 := GenS1{F0: -2147483648, F1: -1, F2: -1}
        _ = v20
    }
    fmt.Println(v3, v1, v1)
    v1 -= 2
    v2 = false
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v3 -= genFn0(-215, 957, 0)
    for fi21 := 0; fi21 < 14; fi21++ {
        if false {
            if false {
                for fi22 := 0; fi22 < 16; fi22++ {
                    if ((1.5 + 1.5) >= 3.14159) {
                        var v23 string = v10
                        _ = v23
                    } else {
                        if true {
                            v1 += (^(v1))
                        }
                        v1 = (v1 % ((255) | 1))
                        var v24 int64 = -468
                        _ = v24
                    }
                    if false {
                        var v25 int16 = v3
                        _ = v25
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi22
                }
                v26 := GenS2{F0: 18446744073709551615}
                _ = v26
            }
        } else {
            fmt.Println(v10, v3, v2)
        }
        _ = fi21
    }
    v27 := GenS1{F0: -1, F1: -20, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v8
}
