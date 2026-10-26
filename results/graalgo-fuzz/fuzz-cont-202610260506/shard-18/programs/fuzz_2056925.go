package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 string; F2 bool }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int32; F1 uint64; F2 bool }
func genFn0(p0 uint8, p1 int32, p2 uint) int64 {
    var v4 uint = p2
    if (!false) {
        p0 += 2
    }
    switch 128 {
    case -185:
        v5 := GenS1{F0: (p2 >> (((^(p2))) & 63)), F1: 4294967295, F2: 9223372036854775807}
        _ = v5
    case 128:
        p2 += 131
    }
    p2 = p2
    v6 := []int{1, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 int32, p1 bool, p2 int32) uint16 {
    if true {
        var v7 int = -585
        if (2 > 18446744073709551615) {
            var v8 uint16 = 2
            v9 := []int{4294967295}
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(p1)
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(p0, fi11, v7)
                switch ("" + "hello") {
                case "foo bar":
                    switch "foo bar" {
                    case "foo bar":
                        switch "" {
                        case "\n\t":
                            var v12 int = v7
                            _ = v12
                        default:
                            v13 := GenS1{F0: 0, F1: 353, F2: -9223372036854775808}
                            _ = v13
                        }
                    case "a":
                        var v14 uint8 = 163
                        _ = v14
                    default:
                        fmt.Println(fi10)
                    }
                case "hello":
                    fmt.Println(p1)
                case "":
                    var v15 uint64 = 2
                    _ = v15
                default:
                    var v16 string = ("foo bar" + ("hello" + "foo bar"))
                    _ = v16
                }
                p0 += 68
                _ = fi11
            }
            if (623 <= 0) {
                switch -9223372036854775808 {
                case 552:
                    var v17 float64 = 1.5
                    _ = v17
                default:
                    p1 = p1
                }
                var v18 string = "hello"
                _ = v18
            } else {
                v19 := GenS1{F0: 939, F1: 2147483648, F2: 1}
                v7 = (fi10 - fi10)
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := []int{fi10, fi20}
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
            _ = fi10
        }
        _ = v7
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 float64 = float64(-2.25)
        _ = fi22
        _ = v23
    }
    var v24 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 2
}

func main() {
    v1 := []string{"\n\t", "a"}
    var v2 int = -249
    var v3 uint8 = 255
    v2 -= (v2 + -297)
    v2 = (v2 % ((v2) | 1))
    v2 = v2
    if false {
        v25 := GenS2{F0: -765}
        _ = v25
    } else {
        v26 := GenS2{F0: -1}
        fmt.Println(v2, v2)
        v27 := GenS3{F0: (^(v2)), F1: ("a" + (("hello" + ("\n\t" + (("foo bar" + "a") + "\n\t"))) + "hello")), F2: true}
        _ = v26
        _ = v27
    }
    v2 -= v2
    fmt.Println(v3)
    for fi28 := 0; fi28 < 12; fi28++ {
        v3 -= (^(v3))
        var v29 int32 = -280
        v2 = fi28
        _ = fi28
        _ = v29
    }
    var v30 uint = 255
    v31 := GenS4{F0: v30, F1: false}
    for fi32 := 0; fi32 < 9; fi32++ {
        for fi33 := 0; fi33 < 14; fi33++ {
            v3 += v3
            var v34 int32 = 43
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    if false {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    fmt.Println(v30, v2)
    v30 = (v30 + (v30 ^ 255))
    if (false && true) {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        v2 -= v2
        v3 = v3
    } else {
        v30 -= (v30 / ((255) | 1))
    }
    _ = len(v1)
    if (true && ((float64(-2.25) > (-2.25 - 1.5)) && true)) {
        var v35 float32 = (((0.0 + 1e10) / 0.1) + (1.5 - float32(-2.25)))
        v1 = append(v1, ("a" + "a"))
        _ = v35
    } else {
        var v36 int64 = 1
        if false {
            v37 := GenS2{F0: -1}
            var v38 int32 = 42
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            v36 = (v36 * (v36 % ((v36) | 1)))
            _ = fi39
        }
        _ = v36
    }
    v40 := GenS5{F0: 42, F1: 0, F2: (v30 != 1)}
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v40
}
