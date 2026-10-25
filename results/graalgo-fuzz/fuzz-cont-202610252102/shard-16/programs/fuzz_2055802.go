package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int8; F1 uint16 }
func genFn0(p0 int32, p1 int8, p2 int32) string {
    v4 := []int{}
    fmt.Println(p1, p1)
    if true {
        if true {
            p2 += (p2 | -2147483648)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    var v5 int16 = 424
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "hello"
}

func main() {
    v1 := []int{-943, -2147483648}
    v2 := make(map[string]bool)
    var v3 uint64 = 18446744073709551615
    v3 -= v3
    for fi6 := 0; fi6 < 13; fi6++ {
        if true {
            for fi7 := 0; fi7 < 10; fi7++ {
                v3 = 805
                var v8 float32 = float32(1e10)
                _ = fi7
                _ = v8
            }
        } else {
            v2["hello"] = true
        }
        switch ("\n\t" + "hello") {
        case "foo bar":
            v3 -= 463
        case "":
            if ((v3 >= (^(v3))) && ((float32(-2.25) * ((0.1 + 1.5) * 3.14159)) >= 1.5)) {
                var v9 float32 = -2.25
                var v10 uint = uint(v3)
                _ = v10
                _ = v9
            }
        default:
            fmt.Println(fi6, v3)
        }
        var v11 int32 = -1
        _ = fi6
        _ = v11
    }
    fmt.Println(v3)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 string = ((((("foo bar" + "hello") + ("\n\t" + "foo bar")) + ("a" + "a")) + ("a" + "hello")) + "")
        var v14 int8 = 127
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        v3 = (v3 / ((v3) | 1))
        switch 9223372036854775807 {
        case 1:
            if (0.1 >= 3.14159) {
                v3 = (v3 / (((v3 << ((4294967295) & 63))) | 1))
                if ((-2.25 >= (float64(3.14159) / 1.5)) || true) {
                    var v15 uint16 = 42
                    _ = v15
                }
            } else {
                v3 = (v3 * 42)
                v16 := v2[("a" + ("hello" + ""))]
                fmt.Println(v3)
                _ = v16
            }
        case 9223372036854775807:
            for fi17 := 0; fi17 < 19; fi17++ {
                v3 = 42
                _ = fi17
            }
        default:
            if false {
                v2[""] = (!(42 != -2147483648))
                for fi18 := 0; fi18 < 16; fi18++ {
                    fmt.Println(v3, v3)
                    v3 = v3
                    v19 := GenS1{F0: 1, F1: 0}
                    _ = fi18
                    _ = v19
                }
            }
        }
    }
    _ = len(v1)
    v20 := GenS2{F0: 441}
    v21 := GenS2{F0: 99}
    v22 := v2["hello"]
    var v23 int = 32767
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 int8 = -101
        switch int(v3) {
        case -1:
            if v22 {
                fmt.Println(fi24, v23)
                switch "hello" {
                case "hello":
                    v23 += v23
                case "\n\t":
                    var v26 int16 = 1
                    _ = v26
                }
                var v27 uint32 = uint32(v3)
                _ = v27
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if (-2.25 > -2.25) {
            v23 = (fi24 & (v23 * (fi24 / (((-(fi24))) | 1))))
            v22 = (v3 < (v3 | (^(v3))))
        } else {
            for fi28 := 0; fi28 < 16; fi28++ {
                v22 = ((^(v25)) != (v25 % (((-(v25))) | 1)))
                if v22 {
                    v29 := GenS3{F0: "", F1: 461}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v22, v25, v23)
                    v30 := GenS1{F0: -565, F1: -377}
                    switch "a" {
                    case "hello":
                        v3 -= v3
                    }
                    _ = v29
                    _ = v30
                }
                _ = fi28
            }
            v31 := GenS1{F0: 2147483647, F1: 42}
            v3 = (v3 % (((v3 - (v3 | v3))) | 1))
            _ = v31
        }
        _ = fi24
        _ = v25
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        if v22 {
            if len(v1) > 0 {
                v1[0] = fi32
            }
            v3 = (v3 * (v3 << ((927) & 63)))
        }
        fmt.Println(v23, v22)
        var v33 uint64 = (^(v3))
        _ = fi32
        _ = v33
    }
    v34 := GenS4{F0: 1, F1: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
}
