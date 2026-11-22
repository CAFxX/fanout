package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint, p1 int32, p2 uint64) uint16 {
    var v7 int8 = 127
    v8 := GenS1{F0: 1, F1: p2}
    v9 := []int{9223372036854775807}
    p0 = 4294967295
    if (!false) {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint = (p0 & (p0 / (((^(p0))) | 1)))
            _ = fi10
            _ = v11
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v12 bool = true
        _ = v12
    } else {
        p2 = p2
        if false {
            if (1 != -823) {
                fmt.Println(p2, p0)
                p0 = (p0 << ((1) & 63))
            }
            if len(v9) > 0 {
                v9[0] = 32768
            }
        } else {
            if (uint(p2) > 625) {
                var v13 uint = 0
                if ((!(!(((p1 / ((0) | 1)) < p1) && (3.14159 != float64(-2.25))))) && (false || (float64(0.0) == ((-2.25 / 1.0) - 1e10)))) {
                    v14 := GenS1{F0: 128, F1: (p2 + 242)}
                    v9 = append(v9, -715)
                    _ = v14
                } else {
                    v9 = append(v9, 32768)
                    var v15 uint32 = 2147483648
                    p2 += p2
                    _ = v15
                }
                _ = v13
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v16 uint16 = 0
            _ = v16
        }
        switch 32767 {
        case -2147483648:
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(p1, p1, p0)
                if (float64(1e10) == 0.1) {
                    if (!false) {
                        for fi18 := 0; fi18 < 7; fi18++ {
                            v7 += (v7 - 63)
                            fmt.Println(fi17, p2)
                            _ = fi18
                        }
                    }
                    v9 = append(v9, (fi17 ^ (-(fi17))))
                    var v19 bool = false
                    _ = v19
                }
                v20 := GenS1{F0: 183, F1: (p2 / (((p2 << ((42) & 63))) | 1))}
                _ = fi17
                _ = v20
            }
        default:
            v21 := GenS1{F0: 128, F1: p2}
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 163
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 2
    var v3 uint8 = uint8(v2)
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    var v6 int16 = 756
    v4[""] = (((float64((0.0 * float32(0.1))) * 1e10) < ((1.5 + 0.0) * 1e10)) && (658 > -2147483648))
    if true {
        v22 := []int{42, 4294967295, -852, 4294967295}
        _ = v22
    }
    var v23 string = ""
    var v24 uint8 = 183
    var v25 uint16 = 124
    v2 = 65535
    var v26 int16 = (v6 - (v6 - v6))
    v27 := GenS2{F0: 1}
    var v28 float64 = 3.14159
    v29 := []int{1, -2147483648}
    v30 := GenS2{F0: 727}
    for fi31 := 0; fi31 < 2; fi31++ {
        fmt.Println(v25, v6, v3)
        var v32 uint16 = (^(v25))
        _ = fi31
        _ = v32
    }
    switch v23 {
    case "a":
        var v33 uint8 = (v24 % ((128) | 1))
        _ = v33
    default:
        for fi34 := 0; fi34 < 20; fi34++ {
            _ = len(v29)
            _ = fi34
        }
    }
    if (315 < 2) {
        v24 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
