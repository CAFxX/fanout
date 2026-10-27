package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int32; F1 int; F2 uint }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint; F2 bool }
func genFn0() uint16 {
    v5 := GenS1{F0: ("a" + (("a" + ("" + "foo bar")) + "a")), F1: 1}
    if (1 != 1) {
        var v6 int8 = 1
        if (449 < 42) {
            v7 := []int{}
            _ = v7
        } else {
            switch "" {
            case "a":
                v6 = 1
            case "\n\t":
                var v8 uint64 = 18446744073709551615
                _ = v8
            case "foo bar":
                v6 += -128
            default:
                v6 = (v6 + -1)
            }
        }
        _ = v6
    }
    var v9 uint8 = 2
    v10 := []int{32768, 392, -9223372036854775808}
    v9 += (v9 + 1)
    _ = v10
    _ = v5
    _ = v9
    return 0
}

func main() {
    var v1 int64 = -1
    var v2 float64 = 3.14159
    v3 := make(map[int]bool)
    var v4 float32 = (-2.25 / float32(v2))
    switch ("a" + "hello") {
    case "":
        if (1 == 840) {
            v4 -= (v4 / (v4 * v4))
        }
    default:
        fmt.Println(v2, v1)
    }
    if (true || (!true)) {
        var v11 int16 = -32768
        _ = v11
    }
    v4 -= (1e10 / v4)
    if ((v2 + v2) <= v2) {
        v12 := v3[128]
        _ = v12
    }
    if true {
        if true {
            var v13 int32 = 42
            switch 378 {
            case 128:
                if true {
                    v2 -= 0.0
                    v14 := GenS2{F0: -1, F1: 828, F2: 2}
                    var v15 uint = 1
                    _ = v14
                    _ = v15
                }
            case 2147483647:
                _ = len(v3)
            case 4294967295:
                for fi16 := 0; fi16 < 17; fi16++ {
                    if ((v1 != v1) && true) {
                        var v17 uint16 = 378
                        fmt.Println(fi16, v1, v13)
                        _ = v17
                    } else {
                        var v18 int32 = (v13 * v13)
                        _ = v18
                    }
                    v19 := GenS1{F0: "foo bar", F1: (v13 << ((42) & 63))}
                    for fi20 := 0; fi20 < 16; fi20++ {
                        fmt.Println(v4, v2, v4)
                        v21 := GenS3{F0: 318, F1: 128}
                        switch "hello" {
                        case "\n\t":
                            _ = len(v3)
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = fi16
                    _ = v19
                }
            default:
                v22 := []int{2147483647, -9223372036854775808}
                _ = v22
            }
            _ = v13
        } else {
            if ((!(721 >= 32768)) && (106 == 1)) {
                fmt.Println(v1, v1, v2)
                delete(v3, 943)
            } else {
                if (true && true) {
                    v2 = 1e10
                }
                v4 = v4
                v23 := GenS1{F0: "\n\t", F1: int32(v1)}
                _ = v23
            }
            v4 += ((1.5 / 0.1) * v4)
            v24 := GenS4{F0: "hello"}
            _ = v24
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            switch fi25 {
            case 32768:
                v26 := GenS5{F0: -119, F1: 42, F2: false}
                _ = v26
            case 128:
                if (false && ("a" != "")) {
                    v4 += v4
                    v2 += float64(v4)
                    fmt.Println(v1)
                } else {
                    if false {
                        v2 = (0.0 - v2)
                        var v27 bool = (136 == 42)
                        _ = v27
                    }
                    var v28 int = -2147483648
                    _ = v28
                }
            case -634:
                v29 := []int{}
                _ = v29
            }
            _ = fi25
        }
        fmt.Println(v4)
    }
    switch ("\n\t" + "hello") {
    case "hello":
        fmt.Println(v4)
    case "":
        if (-128 < -128) {
            v2 = v2
            v30 := GenS2{F0: 42, F1: 0, F2: 352}
            v4 += v4
            _ = v30
        }
    default:
        if true {
            v31 := []int{32767, 515, -9223372036854775808, -260}
            _ = v31
        }
    }
    var v32 uint16 = 2
    var v33 uint16 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
