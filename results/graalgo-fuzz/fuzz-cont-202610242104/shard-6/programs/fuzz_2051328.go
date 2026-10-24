package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint; F2 int8 }
type GenS6 struct { F0 string }
func genFn0(p0 uint, p1 uint64, p2 int64) uint8 {
    if ((3.14159 - 0.1) >= (float64(3.14159) + 1e10)) {
        var v5 int64 = (p2 | (-(p2)))
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: 634, F1: int64(p1), F2: -614}
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    fmt.Println(p2, p1, p1)
    switch -744 {
    case -9223372036854775808:
        if true {
            v8 := GenS2{F0: (^(p0)), F1: (p2 / ((p2) | 1)), F2: 42}
            _ = v8
        } else {
            var v9 int8 = 1
            _ = v9
        }
    case 2147483647:
        if true {
            for fi10 := 0; fi10 < 4; fi10++ {
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 string = "hello"
                    _ = fi11
                    _ = v12
                }
                if (false || true) {
                    var v13 float32 = 0.1
                    _ = v13
                }
                _ = fi10
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 153
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]bool)
    var v3 int16 = 42
    var v4 float64 = 3.14159
    var v14 uint8 = 255
    switch 127 {
    case -9223372036854775808:
        v15 := GenS3{F0: false, F1: ("\n\t" + ("a" + "a"))}
        _ = v15
    case 32768:
        if false {
            _ = len(v2)
            for fi16 := 0; fi16 < 11; fi16++ {
                if false {
                    var v17 float32 = (1.5 + 1.5)
                    _ = v17
                }
                v18 := GenS4{F0: v3, F1: 757}
                _ = fi16
                _ = v18
            }
            v3 -= v3
        } else {
            v19 := GenS5{F0: true, F1: (v1 ^ (^(v1))), F2: int8(v1)}
            _ = v19
        }
    case -2147483648:
        var v20 string = (("" + "a") + "")
        _ = v20
    default:
        fmt.Println(v3, v1)
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := []int{fi22, (fi21 << ((65535) & 63)), fi21}
            var v24 uint = (v1 - (v1 >> (((v1 >> ((v1) & 63))) & 63)))
            v24 += (v1 | (v24 | v1))
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    var v25 bool = false
    var v26 int64 = 42
    switch "a" {
    case "hello":
        if (!(!false)) {
            v25 = (v25 || v25)
            v3 = v3
        }
    case "":
        fmt.Println(v3, v1, v26)
    default:
        if v25 {
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := GenS3{F0: v25, F1: "a"}
                v29 := []int{715, 4294967295}
                _ = len(v29)
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    fmt.Println(v25, v3, v4)
    v14 = genFn0(v1, 18446744073709551615, v26)
    v26 = -1
    v25 = v25
    fmt.Println(v1, v26, v3)
    v30 := GenS5{F0: v25, F1: 42, F2: 58}
    v31 := GenS6{F0: "a"}
    var v32 uint64 = 1
    if (((((v4 / (((1.5 + 1e10) + 0.1) + 0.1)) - (0.0 - (v4 - v4))) * v4) - (1.5 / v4)) >= float64((1.5 / 1.0))) {
        v33 := GenS5{F0: true, F1: (v1 + (v1 ^ (v1 + v1))), F2: 74}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
