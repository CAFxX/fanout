package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint64, p1 uint) uint8 {
    v6 := []int{-1}
    var v7 int8 = 1
    p1 += uint(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 209
}
func genFn1(p0 uint, p1 int64) uint32 {
    v8 := []int{}
    if false {
        v9 := GenS1{F0: (797 <= 128), F1: 127}
        var v10 int = 128
        _ = v10
        _ = v9
    }
    switch "hello" {
    case "hello":
        var v11 bool = false
        _ = v11
    case "\n\t":
        v12 := GenS2{F0: "\n\t"}
        _ = v12
    default:
        var v13 int32 = -816
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        if (1.5 == float32(3.14159)) {
            if true {
                for fi15 := 0; fi15 < 12; fi15++ {
                    if (!(!(3.14159 == 0.0))) {
                        v16 := GenS2{F0: ("foo bar" + "hello")}
                        p0 -= (^(p0))
                        switch "hello" {
                        case "hello":
                            v8 = append(v8, (fi15 << ((42) & 63)))
                        }
                        p0 += 255
                        if (float64(((1.5 + 0.0) - -2.25)) > -2.25) {
                            var v17 int32 = 222
                            switch (fi14 & fi14) {
                            case 839:
                                p1 += (p1 & (p1 >> (((p0 & 2)) & 63)))
                            default:
                                var v18 string = "foo bar"
                                _ = v18
                            }
                            _ = v17
                        }
                        _ = v16
                    } else {
                        if (false || (genFn0(2, (p0 ^ (p0 / ((18446744073709551615) | 1)))) > genFn0(228, p0))) {
                            p0 += (p0 ^ p0)
                        } else {
                            v19 := GenS3{F0: 0, F1: 335}
                            _ = v19
                        }
                    }
                    p0 += 440
                    _ = fi15
                }
                p0 = (^(p0))
            }
            v20 := GenS3{F0: -255, F1: 510}
            _ = v20
        }
        _ = fi14
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = v8
    return 65535
}

func main() {
    var v1 int64 = 42
    var v2 int8 = -102
    v3 := []string{("foo bar" + "\n\t"), "a", "foo bar"}
    var v4 bool = false
    v5 := make(map[string]string)
    fmt.Println(v1)
    var v21 float32 = 3.14159
    v22 := GenS2{F0: (("foo bar" + "\n\t") + "hello")}
    v23 := GenS2{F0: "hello"}
    _ = len(v5)
    for fi24 := 0; fi24 < 7; fi24++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi24
    }
    var v25 int64 = v1
    v5["foo bar"] = ("a" + "foo bar")
    v26 := GenS3{F0: -2147483648, F1: 405}
    var v27 uint8 = 0
    switch 279 {
    case 42:
        v1 = -9223372036854775808
    default:
        fmt.Println(v2, v27)
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = fi28
    }
    var v29 int32 = -693
    var v30 uint16 = 32767
    var v31 float64 = (1.5 - (0.1 + (3.14159 / -2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
