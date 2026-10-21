package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint; F2 int8 }
func genFn0() int64 {
    var v6 uint8 = 127
    v7 := []int{503, 646, -591, 9223372036854775807}
    if true {
        if (868 > 0) {
            v8 := GenS1{F0: 1000, F1: 255}
            if (!false) {
                v6 = 1
                v6 += v6
                fmt.Println(v6)
            }
            _ = v8
        } else {
            for fi9 := 0; fi9 < 14; fi9++ {
                fmt.Println(v6, fi9)
                _ = fi9
            }
            switch 2147483647 {
            case 9223372036854775807:
                var v10 uint32 = 2147483647
                _ = v10
            case 127:
                v7 = append(v7, -689)
            case 0:
                v11 := GenS1{F0: -124, F1: 42}
                _ = v11
            }
        }
    }
    _ = v6
    _ = v7
    return 516
}
func genFn1(p0 int8, p1 uint) int32 {
    var v12 uint = (p1 & 2)
    var v13 float32 = ((((float32((float64(-2.25) - 1.5)) + 1.5) * 3.14159) + -2.25) / (1e10 - 0.0))
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []uint32{2, 2147483647, 2, 65535}
    var v2 bool = false
    var v3 uint64 = 42
    var v4 int8 = 1
    v5 := []int{}
    for fi14 := 0; fi14 < 15; fi14++ {
        if v2 {
            fmt.Println(v4)
            v15 := GenS2{F0: -515, F1: 42, F2: (^(v3))}
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                var v16 int16 = 32767
                _ = v16
            }
            _ = v15
        }
        v5 = append(v5, fi14)
        _ = fi14
    }
    switch 9223372036854775807 {
    case 4294967295:
        _ = len(v1)
    case -64:
        if len(v5) > 0 {
            v5[0] = 32768
        }
    case 42:
        v2 = (2 > 1)
    }
    switch ("\n\t" + (("a" + "\n\t") + "\n\t")) {
    case "a":
        var v17 int8 = (v4 / ((v4) | 1))
        _ = v17
    }
    if v2 {
        v4 -= (v4 * -21)
    }
    var v18 uint64 = v3
    for fi19 := 0; fi19 < 10; fi19++ {
        v4 += 42
        _ = len(v5)
        if (-2.25 >= 1e10) {
            v18 = v18
            fmt.Println(v4, v2)
            if (53 > 2) {
                v20 := GenS1{F0: -1, F1: 0}
                _ = v20
            } else {
                v18 = (v3 ^ v18)
            }
        } else {
            fmt.Println(fi19)
        }
        _ = fi19
    }
    if (("a" != "a") && (1 != 436)) {
        var v21 float64 = float64(0.0)
        for fi22 := 0; fi22 < 2; fi22++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v23 int32 = -1
            _ = fi22
            _ = v23
        }
        v4 = (v4 << ((18446744073709551615) & 63))
        _ = v21
    } else {
        var v24 uint64 = 42
        v4 = -1
        _ = v24
    }
    v4 = 42
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 uint32 = 65535
            var v28 bool = true
            v29 := GenS1{F0: 1, F1: 57}
            var v30 bool = (false || (!v2))
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = len(v1)
        _ = fi25
    }
    switch "hello" {
    case "hello":
        for fi31 := 0; fi31 < 10; fi31++ {
            if len(v1) > 0 {
                v1[0] = 65535
            }
            v2 = (!(0 >= 2))
            var v32 int64 = 9223372036854775807
            _ = fi31
            _ = v32
        }
    case "foo bar":
        v18 = v3
    default:
        v18 -= v18
    }
    var v33 int8 = (v4 * (v4 | v4))
    v34 := GenS3{F0: 0, F1: 828, F2: (-(v33))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
