package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint16, p2 int8) uint8 {
    v5 := GenS1{F0: "foo bar", F1: 0}
    if false {
        p1 -= 276
        var v6 uint = 4294967295
        _ = v6
    }
    var v7 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 2
}
func genFn1() uint8 {
    v8 := GenS1{F0: ("hello" + ("foo bar" + ("\n\t" + "\n\t"))), F1: 18446744073709551615}
    var v9 uint64 = 0
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := []int{496, 2147483647, -510}
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 int8 = 0
    v12 := GenS1{F0: "foo bar", F1: 4294967295}
    if true {
        v4 = v4
        var v13 uint64 = 2
        _ = v13
    } else {
        fmt.Println(v4)
        var v14 uint32 = 42
        _ = v14
    }
    v4 -= v4
    for fi15 := 0; fi15 < 13; fi15++ {
        if ((true && (!false)) || true) {
            fmt.Println(v4, fi15)
            v4 -= (v4 / ((v4) | 1))
            if (!((!(32 <= 197)) || ((-21 >= 25) && false))) {
                var v16 string = ""
                _ = v16
            } else {
                fmt.Println(fi15, fi15)
            }
        }
        v4 += v4
        _ = fi15
    }
    v4 -= (v4 + v4)
    var v17 int8 = v4
    var v18 int16 = -25
    v17 = 0
    v19 := GenS2{F0: (^(v18))}
    switch "foo bar" {
    case "foo bar":
        if false {
            v4 = -44
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 uint16 = 152
                v4 = -70
                v18 = v18
                _ = fi20
                _ = v21
            }
        }
    case "":
        switch 32767 {
        case 2147483647:
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 int = fi22
                v23 += (fi22 * (v23 + (v23 << ((18446744073709551615) & 63))))
                _ = fi22
                _ = v23
            }
        default:
            switch -1 {
            case 4294967295:
                if false {
                    v17 += v17
                    if len(v1) > 0 {
                        v1[0] = 32767
                    }
                }
            default:
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
            }
        }
    case "a":
        v24 := GenS1{F0: ("" + "a"), F1: 1}
        _ = v24
    default:
        v17 = v4
    }
    var v25 uint64 = 42
    var v26 int16 = (v18 << ((904) & 63))
    if false {
        v25 += 514
        if (0.0 != 3.14159) {
            _ = len(v1)
            v27 := v3[351]
            v28 := GenS2{F0: (^(v26))}
            _ = v27
            _ = v28
        } else {
            var v29 float64 = (1e10 / float64((float32(3.14159) - -2.25)))
            _ = v29
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            v17 = -118
            var v31 bool = true
            _ = fi30
            _ = v31
        }
    }
    v17 += v4
    switch 584 {
    case -639:
        var v32 uint32 = 609
        _ = v32
    default:
        var v33 uint8 = 1
        _ = v33
    }
    v1 = append(v1, 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
