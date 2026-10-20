package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 int }
func genFn0(p0 int8, p1 int8, p2 uint64) bool {
    fmt.Println(p0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return false
}
func genFn1(p0 uint8, p1 uint8) uint16 {
    fmt.Println(p0, p1)
    var v6 int8 = -1
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 564
    var v2 int32 = 2147483647
    var v3 bool = true
    var v4 float64 = 1e10
    v5 := []int{32768, -9223372036854775808, -1}
    for fi7 := 0; fi7 < 17; fi7++ {
        switch "hello" {
        case "":
            var v8 int64 = -725
            _ = v8
        default:
            if len(v5) > 0 {
                v5[0] = (fi7 + fi7)
            }
        }
        v9 := GenS1{F0: 100, F1: 42}
        var v10 int64 = int64(v2)
        _ = fi7
        _ = v10
        _ = v9
    }
    if (!false) {
        v3 = (-428 >= 128)
    }
    v2 -= (v2 + (v2 * v2))
    var v11 int16 = int16(v2)
    for fi12 := 0; fi12 < 3; fi12++ {
        if (!v3) {
            var v13 uint16 = 477
            _ = v13
        }
        _ = fi12
    }
    if v3 {
        v3 = (!true)
        v14 := GenS2{F0: 18446744073709551615, F1: 100}
        _ = v14
    } else {
        var v15 int8 = 42
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v15
    }
    if (!((!v3) && (!v3))) {
        v16 := GenS3{F0: 82}
        _ = v16
    } else {
        v1 = v1
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint8 = 128
            v11 = v11
            _ = fi17
            _ = v18
        }
        switch ("hello" + ("" + ("hello" + ("" + "")))) {
        case "hello":
            _ = len(v5)
        case "a":
            v19 := GenS2{F0: 65535, F1: 1}
            _ = v19
        case "\n\t":
            fmt.Println(v3, v3)
        }
    }
    var v20 float64 = (((v4 + v4) - 1e10) - 0.1)
    v11 -= (v11 ^ v11)
    switch ("foo bar" + "foo bar") {
    case "":
        v3 = true
    case "foo bar":
        if false {
            for fi21 := 0; fi21 < 17; fi21++ {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 float64 = (1e10 / v4)
                    switch fi21 {
                    case 0:
                        v5 = append(v5, 4294967295)
                    case -9223372036854775808:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 uint8 = 138
                    v2 = 1
                    _ = fi24
                    _ = v25
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi21
            }
            switch 2147483647 {
            case 4294967295:
                fmt.Println(v4, v4, v3)
            case 9223372036854775807:
                v5 = append(v5, 611)
            case 127:
                var v26 uint16 = 1
                _ = v26
            default:
                v27 := GenS3{F0: 2}
                _ = v27
            }
            v2 += 2147483647
        }
    default:
        v20 += (v4 * v4)
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v3 = false
        v3 = false
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi28
    }
    if (42 > 32767) {
        v29 := GenS1{F0: -128, F1: v11}
        _ = v29
    } else {
        v20 -= float64(3.14159)
    }
    v30 := GenS4{F0: (v11 + (v11 & -731)), F1: 374, F2: 42}
    v20 = float64(((0.0 - -2.25) / 1e10))
    var v31 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
