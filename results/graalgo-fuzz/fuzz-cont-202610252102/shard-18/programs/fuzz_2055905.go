package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 string; F1 uint; F2 int }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 int32, p1 int, p2 int) string {
    v7 := []int{(p2 % ((p1) | 1))}
    _ = len(v7)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (119 > int8(p1)) {
        for fi8 := 0; fi8 < 13; fi8++ {
            switch "foo bar" {
            case "":
                fmt.Println(p0, p0)
            default:
                p2 = 128
            }
            fmt.Println(fi8, p1)
            fmt.Println(p2, p1)
            _ = fi8
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            switch "hello" {
            case "\n\t":
                var v10 uint32 = 42
                _ = v10
            case "a":
                if (42 <= 407) {
                    v11 := GenS1{F0: -32768, F1: ((!(4294967295 >= 2)) || false), F2: (("\n\t" + "") + "foo bar")}
                    v12 := GenS1{F0: 0, F1: true, F2: "foo bar"}
                    if (9223372036854775807 < -1) {
                        var v13 float32 = float32((float64(0.1) - 1e10))
                        var v14 uint16 = 159
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS1{F0: -32768, F1: (255 > 255), F2: "\n\t"}
                        _ = v15
                    }
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                fmt.Println(p2, p2, p1)
            }
            var v16 uint64 = 117
            _ = fi9
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ""
}
func genFn1() uint16 {
    if true {
        if true {
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 int = fi17
                v19 := []int{(v18 + (v18 >> ((65535) & 63))), v18}
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: ("a" + ("" + (("" + "foo bar") + ("a" + "\n\t")))), F1: 566, F2: 1}
            _ = v20
        }
        v21 := []int{127, 9223372036854775807}
        _ = v21
    }
    var v22 int8 = -48
    _ = v22
    return 969
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42, 2147483647, 158, 979}
    v3 := make(map[int]int)
    v4 := []string{("a" + ("foo bar" + ("" + "a"))), ("" + "foo bar"), "", ""}
    var v5 uint8 = 127
    var v6 float32 = 0.1
    _ = len(v2)
    v6 = float32(((3.14159 * 1e10) * -2.25))
    v6 = (float32((3.14159 / 1e10)) + ((((v6 / 3.14159) * v6) * float32(0.1)) / v6))
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 uint = 762
        _ = fi23
        _ = v24
    }
    v6 = v6
    if (377 == 42) {
        var v25 int64 = -9223372036854775808
        v25 = v25
        switch "" {
        case "a":
            fmt.Println(v25, v25)
        case "hello":
            v26 := GenS1{F0: -295, F1: (690 >= 2147483648), F2: "\n\t"}
            _ = v26
        }
        _ = v25
    } else {
        v27 := GenS3{F0: v5, F1: 0}
        var v28 float64 = (((0.0 * 0.1) / 1.5) - 3.14159)
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int32 = -1
            switch "\n\t" {
            case "":
                v31 := v3[32767]
                _ = v31
            case "hello":
                v5 = v5
            case "\n\t":
                v5 += (^(v5))
            default:
                for fi32 := 0; fi32 < 9; fi32++ {
                    v6 += v6
                    _ = fi32
                }
            }
            _ = fi29
            _ = v30
        }
        _ = v27
        _ = v28
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 8; fi34++ {
            switch (fi33 - fi33) {
            case 128:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v35 uint16 = 199
                _ = v35
            }
            v36 := GenS4{F0: 42, F1: 100}
            v37 := v3[4294967295]
            _ = fi34
            _ = v36
            _ = v37
        }
        var v38 uint8 = (v5 * (v5 + 0))
        fmt.Println(v5, v6)
        _ = fi33
        _ = v38
    }
    var v39 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
