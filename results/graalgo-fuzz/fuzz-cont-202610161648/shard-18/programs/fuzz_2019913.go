package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint8, p2 int64) uint64 {
    var v6 float32 = ((-2.25 * 0.0) + float32((0.0 + 3.14159)))
    p2 = (p2 / (((p2 * (p2 % ((0) | 1)))) | 1))
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int64) uint8 {
    v8 := []int{32767, 354, 1, 1}
    v9 := GenS1{F0: 42}
    var v10 int = -2147483648
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 uint = 340
    var v3 int16 = 0
    v4 := make(map[string]bool)
    v5 := []string{(("\n\t" + ("\n\t" + "a")) + ""), "hello"}
    v3 = 42
    if (!false) {
        if true {
            v1 += 9223372036854775807
            for fi11 := 0; fi11 < 19; fi11++ {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                _ = fi11
            }
        }
        switch -340 {
        case 972:
            v12 := GenS2{F0: 42}
            _ = v12
        }
        v13 := GenS3{F0: -396, F1: ("" + "foo bar"), F2: v1}
        _ = v13
    } else {
        if ("" < "") {
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS4{F0: fi14, F1: 69}
                _ = fi14
                _ = v15
            }
        }
        v16 := GenS5{F0: 516}
        _ = v16
    }
    fmt.Println(v1, v2, v2)
    var v17 uint = v2
    var v18 int = -756
    v2 = (^(v2))
    v3 = v3
    v2 += (v17 ^ v17)
    if true {
        switch ("" + ("hello" + "a")) {
        case "foo bar":
            v19 := GenS3{F0: (-(v1)), F1: "a", F2: (v1 ^ (v1 & -646))}
            _ = v19
        }
    }
    switch 4294967295 {
    case 1:
        v3 = (v3 | v3)
    case 4294967295:
        var v20 uint = v2
        _ = v20
    case -1:
        v17 = (v2 & 4294967295)
    default:
        for fi21 := 0; fi21 < 18; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 uint8 = 128
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    v2 -= v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
