package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 uint; F1 int16; F2 string }
func genFn0(p0 uint32) string {
    var v5 int32 = -2147483648
    v6 := GenS1{F0: ("\n\t" + "\n\t"), F1: (p0 - p0)}
    v7 := []int{-1, 2147483647, -9223372036854775808}
    if (-484 <= -9223372036854775808) {
        var v8 uint32 = 2147483647
        for fi9 := 0; fi9 < 1; fi9++ {
            v5 -= (^(v5))
            for fi10 := 0; fi10 < 9; fi10++ {
                v8 += (p0 - (^(p0)))
                v11 := GenS2{F0: 1, F1: fi10, F2: -1}
                p0 -= 1
                _ = fi10
                _ = v11
            }
            p0 -= v8
            _ = fi9
        }
        _ = v8
    } else {
        v7 = append(v7, 161)
    }
    if ((p0 == (p0 % (((p0 - (p0 << ((18446744073709551615) & 63)))) | 1))) || true) {
        p0 += (p0 & (p0 / (((p0 << ((255) & 63))) | 1)))
        p0 = (p0 * p0)
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return "hello"
}
func genFn1() uint16 {
    switch -843 {
    case 58:
        var v12 uint32 = 2
        _ = v12
    case 42:
        var v13 string = "foo bar"
        _ = v13
    default:
        var v14 bool = (!(42 <= 0))
        _ = v14
    }
    var v15 uint = 2
    _ = v15
    return 255
}

func main() {
    var v1 uint = 411
    v2 := make(map[string]bool)
    var v3 float64 = float64(0.1)
    var v4 uint64 = 332
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS1{F0: "\n\t", F1: 909}
        v18 := []int{(fi16 ^ 1)}
        if len(v18) > 0 {
            v18[0] = fi16
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float32 = ((((1e10 * 3.14159) / float32(((v3 * v3) * v3))) + 1.5) / (3.14159 / 0.1))
    v20 := GenS3{F0: (v1 << ((65535) & 63)), F1: 338, F2: (("" + (("foo bar" + "foo bar") + "a")) + "hello")}
    fmt.Println(v19, v3, v4)
    v3 -= ((v3 - ((float64((v19 + 0.1)) * ((v3 - ((v3 / (1e10 * float64(v19))) / (v3 - (v3 / (float64((v19 - v19)) * float64((v19 * float32(0.0)))))))) - (1.5 - float64(v19)))) - (float64(v19) * (1e10 + (((v3 * v3) - 1e10) * 0.1))))) - 1.5)
    fmt.Println(v19, v19, v3)
    var v21 int16 = -806
    fmt.Println(v1, v19, v19)
    fmt.Println(v4, v19)
    if (!true) {
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(v3, v3, v19)
            _ = fi22
        }
        var v23 int8 = -1
        if true {
            v3 -= (v3 / v3)
            v3 = (v3 + 3.14159)
        } else {
            switch "foo bar" {
            case "hello":
                fmt.Println(v4, v21)
            default:
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS1{F0: "a", F1: 845}
                    var v26 int64 = -63
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        var v27 float64 = 1e10
                        _ = v27
                    default:
                        v28 := []int{}
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            if (false || (!(!true))) {
                for fi29 := 0; fi29 < 18; fi29++ {
                    v19 = 1.5
                    _ = fi29
                }
            }
            if (true && (false || (v19 < (v19 / v19)))) {
                var v30 uint64 = v4
                _ = v30
            } else {
                fmt.Println(v1, v23, v19)
                v19 = (v19 / (0.1 - 3.14159))
            }
        }
        _ = v23
    } else {
        fmt.Println(v3)
    }
    v31 := []int{2147483647}
    switch "hello" {
    case "foo bar":
        for fi32 := 0; fi32 < 12; fi32++ {
            v31 = append(v31, 95)
            _ = fi32
        }
    case "\n\t":
        v31 = append(v31, 32768)
    }
    v2["\n\t"] = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v4
}
