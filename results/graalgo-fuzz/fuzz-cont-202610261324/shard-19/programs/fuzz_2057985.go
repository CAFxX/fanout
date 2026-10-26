package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 string }
type GenS2 struct { F0 int64; F1 int; F2 uint32 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 bool) string {
    v7 := []int{}
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = 122
        switch "" {
        case "a":
            v10 := GenS1{F0: ("hello" + ""), F1: 992, F2: "foo bar"}
            _ = v10
        case "foo bar":
            p0 = p0
        default:
            var v11 uint16 = 65535
            _ = v11
        }
        if ("hello" > "foo bar") {
            var v12 int64 = v9
            _ = v12
        } else {
            if (65535 > 65535) {
                var v13 float32 = 3.14159
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v14 string = (((("hello" + "") + "") + ("" + (("a" + "\n\t") + "a"))) + ("" + ""))
                _ = v13
                _ = v14
            } else {
                p0 = true
                for fi15 := 0; fi15 < 11; fi15++ {
                    if (-32768 == 1) {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(fi15, fi15)
                            v9 += v9
                        case "foo bar":
                            var v16 string = ("" + "hello")
                            _ = v16
                        }
                        switch 42 {
                        case 565:
                            var v17 uint32 = 0
                            _ = v17
                        case -396:
                            var v18 uint16 = 998
                            _ = v18
                        case 128:
                            p0 = p0
                        default:
                            var v19 uint = 18446744073709551615
                            _ = v19
                        }
                    }
                    if ((!(!(fi15 > fi8))) && p0) {
                        p0 = true
                        if p0 {
                            v20 := GenS2{F0: v9, F1: 32768, F2: 791}
                            _ = v20
                        }
                    }
                    _ = fi15
                }
                v9 = (^(v9))
            }
        }
        _ = fi8
        _ = v9
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        p0 = ((0.1 - (((-2.25 / ((3.14159 + 1.5) - (float32(-2.25) * ((((float32(0.0) / 1e10) + 3.14159) - (3.14159 - 3.14159)) * float32(0.1))))) + ((0.0 + ((0.0 - 1e10) - 3.14159)) / (0.1 + -2.25))) / 3.14159)) != -2.25)
        _ = fi21
    }
    _ = p0
    _ = v7
    return "hello"
}
func genFn1(p0 int64) string {
    var v22 uint8 = 127
    var v23 int = 639
    _ = p0
    _ = v22
    _ = v23
    return "hello"
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]string)
    var v3 uint16 = 1
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    var v24 uint = 2
    var v25 uint8 = 255
    var v26 int = 42
    if ((true && false) || false) {
        v25 += (^(v25))
    }
    v26 = v26
    var v27 int32 = -2147483648
    v25 += v25
    v28 := GenS2{F0: 1, F1: v26, F2: 65535}
    v27 = (v27 * -2147483648)
    v27 += 0
    fmt.Println(v27)
    v29 := GenS3{F0: "a", F1: (v3 / ((v3) | 1))}
    fmt.Println(v26)
    v30 := GenS3{F0: "foo bar", F1: (v3 << (((v24 << (((^(v24))) & 63))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
