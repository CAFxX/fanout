package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool) int16 {
    v6 := []int{296}
    if p0 {
        var v7 uint64 = 509
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch ("\n\t" + "hello") {
        case "":
            if len(v6) > 0 {
                v6[0] = fi8
            }
        default:
            p0 = (42 > -745)
        }
        _ = fi8
    }
    if (float32(0.1) <= 1e10) {
        p0 = p0
    } else {
        v6 = append(v6, 32767)
        for fi9 := 0; fi9 < 9; fi9++ {
            if len(v6) > 0 {
                v6[0] = (fi9 & 32767)
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: 0, F1: 831}
                var v12 uint16 = 691
                p0 = (!p0)
                _ = fi10
                _ = v11
                _ = v12
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi9
        }
    }
    if (!p0) {
        var v13 int8 = 100
        v14 := GenS2{F0: 1}
        if (4294967295 >= 18446744073709551615) {
            if (1 < 42) {
                if p0 {
                    fmt.Println(v13)
                    p0 = (4294967295 <= 2)
                }
                v6 = append(v6, -1)
            }
            fmt.Println(p0, v13)
            var v15 int = -9223372036854775808
            _ = v15
        } else {
            v16 := GenS1{F0: 916, F1: 566}
            v13 -= 37
            var v17 string = "hello"
            _ = v16
            _ = v17
        }
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v6
    return 1000
}
func genFn1(p0 uint, p1 string) int {
    p0 += p0
    var v18 float64 = (3.14159 + (float64(1.5) / 1e10))
    _ = p0
    _ = p1
    _ = v18
    return -2147483648
}

func main() {
    v1 := []string{((("hello" + ("\n\t" + "")) + (("hello" + ("" + "\n\t")) + (("" + (("a" + ("\n\t" + "a")) + "a")) + "foo bar"))) + ("" + (("hello" + "foo bar") + ("foo bar" + ("foo bar" + "\n\t"))))), "", "hello", (("hello" + "hello") + "")}
    var v2 int64 = 9223372036854775807
    var v3 int8 = 1
    v4 := []int{-981}
    var v5 int64 = 42
    var v19 int64 = -579
    var v20 uint16 = 479
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v21 int32 = 856
    v19 += (-(v5))
    v20 += (^(v20))
    _ = len(v1)
    fmt.Println(v2)
    switch -1 {
    case 32768:
        fmt.Println(v3)
    default:
        switch "hello" {
        case "":
            v2 = 1
        case "a":
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 float64 = (3.14159 + 0.1)
                _ = fi22
                _ = v23
            }
        }
    }
    var v24 uint64 = 42
    v21 += 42
    v21 -= -2147483648
    v19 -= (v5 | v2)
    _ = len(v1)
    var v25 uint16 = 42
    var v26 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
