package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 string }
func genFn0() uint {
    var v7 uint8 = 127
    var v8 uint = 168
    var v9 float64 = (1.5 / ((0.0 + 0.0) + (1.5 - 0.1)))
    fmt.Println(v7, v7, v8)
    v10 := []int{int(v8), 0, -9223372036854775808}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 197
}
func genFn1(p0 bool, p1 int32, p2 uint64) string {
    var v11 uint64 = p2
    v12 := GenS1{F0: 65535, F1: (505 < 2147483647), F2: -379}
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return "foo bar"
}

func main() {
    var v1 int8 = -51
    var v2 int64 = -819
    var v3 float32 = ((1e10 * 1.5) / (1.5 * (float32((float64(1e10) + -2.25)) - float32(0.1))))
    v4 := []uint32{2, 0, 42}
    var v5 uint8 = 127
    v6 := make(map[int]string)
    fmt.Println(v2, v2, v3)
    v13 := GenS2{F0: v1, F1: v2}
    if false {
        fmt.Println(v3)
        var v14 float64 = (3.14159 * ((1.5 - (float64((-2.25 + v3)) / (1e10 * 1e10))) + (0.0 * 1e10)))
        _ = v14
    } else {
        v15 := GenS3{F0: false, F1: 1}
        _ = v15
    }
    fmt.Println(v5, v3, v1)
    v3 = v3
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := GenS4{F0: 2, F1: (v1 / ((127) | 1))}
            fmt.Println(v3, v2, v5)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    switch "" {
    case "a":
        v2 = v2
    case "\n\t":
        var v19 uint16 = 1
        _ = v19
    default:
        v1 += (v1 - 0)
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v6)
        _ = fi20
    }
    fmt.Println(v3, v1, v5)
    var v21 bool = true
    v6[-926] = "\n\t"
    if false {
        var v22 float32 = v3
        var v23 bool = (1000 == -565)
        v24 := GenS2{F0: -1, F1: 525}
        _ = v22
        _ = v23
        _ = v24
    }
    v5 -= (^(v5))
    v25 := GenS2{F0: v1, F1: (v2 + (v2 + (-(v2))))}
    if true {
        v26 := GenS5{F0: (("hello" + "foo bar") + ((("a" + "a") + "") + ("a" + "a")))}
        switch "a" {
        case "a":
            var v27 uint64 = 18446744073709551615
            _ = v27
        case "":
            for fi28 := 0; fi28 < 8; fi28++ {
                v29 := GenS1{F0: 0, F1: ((v5 / ((v5) | 1)) > (v5 >> ((genFn0()) & 63))), F2: 1000}
                _ = fi28
                _ = v29
            }
        }
        _ = v26
    } else {
        var v30 float32 = v3
        for fi31 := 0; fi31 < 1; fi31++ {
            if ((0.1 >= (1.5 / -2.25)) && v21) {
                switch "hello" {
                case "foo bar":
                    if true {
                        if (0.1 > -2.25) {
                            var v32 int32 = 1
                            _ = v32
                        }
                        if (!v21) {
                            v21 = (v2 <= v2)
                            for fi33 := 0; fi33 < 17; fi33++ {
                                v30 -= (v30 / v3)
                                fmt.Println(v5)
                                _ = fi33
                            }
                        }
                    }
                default:
                    var v34 int32 = 1
                    _ = v34
                }
                var v35 uint = 255
                var v36 uint = v35
                _ = v35
                _ = v36
            } else {
                var v37 int8 = (-(v1))
                v2 = (v2 % ((42) | 1))
                switch ("hello" + "a") {
                case "hello":
                    v38 := GenS5{F0: "foo bar"}
                    _ = v38
                case "a":
                    _ = len(v6)
                case "foo bar":
                    fmt.Println(v3, v37, v37)
                }
                _ = v37
            }
            v3 += (1e10 - float32(1.5))
            _ = fi31
        }
        _ = v30
    }
    delete(v6, 727)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
