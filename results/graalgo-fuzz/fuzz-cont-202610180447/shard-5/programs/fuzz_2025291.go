package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint64) int {
    p0 -= (p0 - p0)
    if true {
        var v5 int = 9223372036854775807
        if ((!false) && true) {
            v6 := []int{(^(v5)), (v5 & 42), v5}
            fmt.Println(p0, p0, p0)
            v6 = append(v6, (v5 + -9223372036854775808))
            _ = v6
        }
        _ = v5
    }
    var v7 uint = 1
    switch "\n\t" {
    case "a":
        var v8 int = 0
        _ = v8
    default:
        fmt.Println(p1)
    }
    fmt.Println(v7, p1, v7)
    _ = p0
    _ = p1
    _ = v7
    return 741
}

func main() {
    var v1 float32 = (3.14159 / 0.1)
    var v2 string = ""
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    if false {
        v1 += 0.1
        if (933 >= -97) {
            v2 = v2
        } else {
            var v9 int16 = 0
            _ = v9
        }
        if (!false) {
            var v10 int32 = -863
            fmt.Println(v1, v10, v10)
            _ = v10
        }
    }
    fmt.Println(v2, v2)
    v11 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: 4294967295}
    if (0 > 65535) {
        switch 32767 {
        case 127:
            if ((float64((1e10 - 0.1)) * (0.0 / float64(v1))) > 0.0) {
                var v12 uint16 = 65535
                _ = v12
            } else {
                var v13 uint32 = 725
                v2 = v2
                v14 := v4[-597]
                _ = v13
                _ = v14
            }
        case 275:
            switch "\n\t" {
            case "a":
                fmt.Println(v2)
            }
        case 614:
            v15 := []int{127, -987}
            _ = v15
        default:
            for fi16 := 0; fi16 < 14; fi16++ {
                v2 = (("a" + (v2 + (v2 + ((("a" + "") + v2) + v2)))) + "a")
                if ((v1 - 0.1) >= (float32(3.14159) * v1)) {
                    delete(v3, 2147483647)
                    for fi17 := 0; fi17 < 14; fi17++ {
                        var v18 int16 = -32768
                        v2 = (("foo bar" + (v2 + ((v2 + ("hello" + (v2 + v2))) + "hello"))) + ("a" + ("a" + (v2 + ((v2 + v2) + v2)))))
                        v2 = (v2 + (v2 + v2))
                        _ = fi17
                        _ = v18
                    }
                } else {
                    fmt.Println(v1, v2, v1)
                }
                fmt.Println(v2, v2)
                _ = fi16
            }
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v2 = ""
            v1 = ((v1 - v1) - -2.25)
            _ = fi19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            if (true && false) {
                v21 := GenS2{F0: 32767}
                fmt.Println(v2, v1, v1)
                for fi22 := 0; fi22 < 6; fi22++ {
                    v23 := []int{9223372036854775807}
                    v23 = append(v23, (fi22 | (^(fi22))))
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
            _ = fi20
        }
    }
    v2 = v2
    if (1 == 24) {
        v1 -= float32(0.1)
        v2 = v2
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 uint32 = 2147483647
            fmt.Println(v25, v25, fi24)
            v25 -= (v25 / (((v25 >> ((2) & 63))) | 1))
            _ = fi24
            _ = v25
        }
    } else {
        v26 := []int{genFn0(275, 0), genFn0(1, 179), -167, genFn0(42, 18446744073709551615)}
        var v27 uint64 = 887
        _ = v26
        _ = v27
    }
    var v28 int8 = 127
    v2 = v2
    v2 = v2
    var v29 uint64 = 2
    v1 = 1e10
    for fi30 := 0; fi30 < 2; fi30++ {
        v29 -= (v29 * (v29 | (v29 << ((212) & 63))))
        v29 += (v29 ^ (^(v29)))
        v31 := GenS2{F0: 867}
        _ = fi30
        _ = v31
    }
    switch genFn0(32767, (v29 << ((255) & 63))) {
    case 114:
        var v32 uint8 = 28
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
