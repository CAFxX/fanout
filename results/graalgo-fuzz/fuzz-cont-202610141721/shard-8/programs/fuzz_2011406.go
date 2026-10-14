package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 uint, p1 int) uint16 {
    var v5 int8 = 42
    var v6 float64 = -2.25
    v7 := []int{(p1 | p1)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []int{}
    v2 := []uint32{0, 163}
    var v3 int32 = 1
    var v4 float32 = 0.1
    if (161 > 212) {
        v1 = append(v1, 9223372036854775807)
        fmt.Println(v3)
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v3)
            _ = fi8
        }
    }
    if (65535 < 255) {
        if (float32(3.14159) != v4) {
            v9 := GenS1{F0: "\n\t", F1: 128, F2: (v3 | v3)}
            v4 += (3.14159 + (v4 / 1.0))
            _ = v9
        } else {
            v3 -= (v3 * (-(v3)))
        }
        switch 4294967295 {
        case 2147483647:
            v10 := GenS1{F0: "a", F1: 1, F2: 42}
            _ = v10
        case 9223372036854775807:
            fmt.Println(v3)
        case 1:
            if (!false) {
                if (!(65535 != 445)) {
                    var v11 int = -913
                    var v12 uint32 = 1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v11
                    _ = v12
                }
                var v13 int32 = (v3 >> ((42) & 63))
                v13 += v13
                _ = v13
            }
        }
    } else {
        v4 += 1.5
    }
    if (!(65535 < 65535)) {
        var v14 float64 = (-2.25 * 1.5)
        v15 := GenS1{F0: ("\n\t" + ""), F1: 42, F2: -225}
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint16 = 850
                fmt.Println(v3, fi17)
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        var v19 uint64 = 42
        _ = v19
    }
    if false {
        v20 := GenS1{F0: ("a" + "hello"), F1: 16, F2: v3}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            if len(v2) > 0 {
                v2[0] = 551
            }
            var v22 int64 = -9223372036854775808
            _ = fi21
            _ = v22
        }
    }
    v4 += v4
    var v23 int64 = 0
    var v24 int = -9223372036854775808
    var v25 bool = (!false)
    if false {
        v26 := GenS2{F0: 255, F1: "hello"}
        _ = v26
    } else {
        var v27 int16 = 32767
        if (false && v25) {
            _ = len(v1)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
