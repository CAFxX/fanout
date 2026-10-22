package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
func genFn0(p0 int16, p1 uint16) int32 {
    var v6 int8 = 0
    v7 := GenS1{F0: v6, F1: 2147483647}
    switch 305 {
    case 128:
        if true {
            v8 := []int{584}
            _ = v8
        }
    default:
        p1 = 629
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 uint32) uint {
    p1 = 65535
    p1 = p1
    var v9 int16 = 42
    if ((p0 << ((0) & 63)) <= (^(p0))) {
        p1 += p1
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = ((((0.1 * 1.5) + (((-2.25 - 1.5) * 0.0) * (0.1 * -2.25))) + 0.1) - 0.1)
    var v3 uint32 = 4294967295
    v4 := []string{"\n\t"}
    var v5 uint = 2
    v10 := GenS1{F0: 70, F1: -2147483648}
    fmt.Println(v5, v3, v2)
    if false {
        v11 := GenS1{F0: -59, F1: 0}
        _ = v11
    }
    var v12 string = "\n\t"
    if ((^(v5)) != (v5 + (v5 * (v5 & 1)))) {
        switch "" {
        case "a":
            if (3.14159 == (-2.25 - 0.1)) {
                v13 := GenS1{F0: -128, F1: -2147483648}
                fmt.Println(v12)
                var v14 uint8 = 83
                _ = v13
                _ = v14
            } else {
                var v15 uint16 = 42
                v2 = v2
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint16 = v15
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        default:
            v18 := GenS1{F0: 91, F1: 2147483647}
            _ = v18
        }
        if ((2 < 725) || (!true)) {
            fmt.Println(v5, v5, v3)
            delete(v1, (v3 + v3))
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            switch fi19 {
            case -9223372036854775808:
                v12 = (("a" + (("a" + (("\n\t" + v12) + "a")) + v12)) + "\n\t")
            case 4294967295:
                v12 = v12
            }
            v20 := GenS2{F0: 128, F1: (fi19 ^ fi19)}
            var v21 uint8 = 29
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        v2 += ((v2 - v2) + v2)
    }
    v2 = v2
    v22 := GenS3{F0: (v12 + (v12 + (v12 + "a"))), F1: 128}
    v23 := GenS4{F0: "", F1: 0, F2: 52}
    var v24 uint64 = 1
    v2 = 0.0
    v12 = ((("a" + "a") + "\n\t") + v12)
    switch 128 {
    case -524:
        v12 = ""
    case -307:
        fmt.Println(v24, v3)
    default:
        v25 := GenS2{F0: 42, F1: -903}
        _ = v25
    }
    var v26 uint16 = 42
    var v27 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
