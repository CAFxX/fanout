package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32) string {
    switch ("hello" + (("foo bar" + "a") + (("\n\t" + (("" + "a") + "a")) + ""))) {
    case "\n\t":
        p0 = p0
    case "":
        fmt.Println(p0)
    case "hello":
        if ("a" < ("a" + "a")) {
            p0 += 369
            p0 = p0
        } else {
            if (928 > 911) {
                v4 := GenS1{F0: 756, F1: 1}
                p0 += (p0 - p0)
                _ = v4
            } else {
                p0 = p0
            }
            v5 := []int{}
            fmt.Println(p0)
            _ = v5
        }
    }
    v6 := []int{127, -977, 128}
    _ = p0
    _ = v6
    return (("hello" + "foo bar") + "\n\t")
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]bool)
    var v3 int = -1
    v1 -= v1
    fmt.Println(v1)
    for fi7 := 0; fi7 < 12; fi7++ {
        if true {
            v1 -= 0
            v1 -= v1
            if (false && ((9223372036854775807 == 1) && ((!false) || false))) {
                v8 := []int{32767}
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: 32767, F1: 18446744073709551615}
                    _ = fi9
                    _ = v10
                }
                if ("hello" == ("a" + "foo bar")) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    v8 = append(v8, 128)
                } else {
                    if (((float64(1e10) / 1.0) / (1.5 / float64(-2.25))) < ((((3.14159 - 1e10) - -2.25) + 0.1) / 3.14159)) {
                        v11 := GenS1{F0: 465, F1: 557}
                        for fi12 := 0; fi12 < 10; fi12++ {
                            if ((!false) && false) {
                                switch "\n\t" {
                                case "hello":
                                    var v13 uint = 2
                                    _ = v13
                                }
                            }
                            _ = fi12
                        }
                        v14 := GenS1{F0: 568, F1: 65535}
                        _ = v11
                        _ = v14
                    }
                }
                _ = v8
            } else {
                var v15 uint8 = 59
                _ = v15
            }
        }
        var v16 uint8 = 1
        _ = fi7
        _ = v16
    }
    v17 := GenS2{F0: (v3 ^ (-(v3))), F1: 0}
    var v18 uint16 = 32767
    switch "foo bar" {
    case "":
        v19 := GenS3{F0: 63}
        _ = v19
    case "foo bar":
        v18 += (v18 * (v18 - v18))
    default:
        var v20 int16 = 0
        _ = v20
    }
    var v21 uint = 18446744073709551615
    v18 = v18
    v22 := GenS1{F0: v18, F1: (v21 - v21)}
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 int16 = 393
        switch 774 {
        case -627:
            v2[367] = (!((-9223372036854775808 >= 0) && false))
        case -9223372036854775808:
            v25 := GenS3{F0: 142}
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    var v26 uint64 = 2
    var v27 int64 = 0
    v26 = v26
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 float32 = -2.25
        var v30 uint = v21
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = len(v2)
    var v31 uint = v21
    v32 := []int{(v3 / ((v3) | 1)), (v3 - (v3 >> ((v31) & 63))), (-(v3)), v3}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
}
