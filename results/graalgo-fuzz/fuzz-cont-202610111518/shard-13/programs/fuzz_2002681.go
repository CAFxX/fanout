package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 string; F1 uint16; F2 int16 }
type GenS6 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint32, p1 bool, p2 uint16) uint {
    v4 := GenS1{F0: 1, F1: ("" + "a"), F2: 42}
    fmt.Println(p2, p2, p1)
    if (0 == 2) {
        p2 += p2
        v5 := []int{}
        _ = v5
    } else {
        if (!(42 == 116)) {
            v6 := GenS2{F0: 65535, F1: ("foo bar" + "hello")}
            var v7 uint32 = 4294967295
            _ = v6
            _ = v7
        } else {
            var v8 float64 = 0.0
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := GenS3{F0: p2, F1: 148}
                _ = fi9
                _ = v10
            }
            var v11 int16 = 1
            _ = v11
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]string)
    var v3 uint32 = 2147483647
    if false {
        var v12 uint32 = (v3 - (v3 >> ((18446744073709551615) & 63)))
        var v13 uint8 = 255
        if false {
            var v14 uint8 = (v13 & (^(v13)))
            _ = v14
        } else {
            switch "" {
            case "hello":
                if (-100 < (v1 ^ -114)) {
                    v2[-974] = ""
                }
            default:
                var v15 int64 = 9223372036854775807
                _ = v15
            }
        }
        _ = v12
        _ = v13
    }
    switch 1 {
    case -9223372036854775808:
        if false {
            var v16 int = 0
            var v17 string = "\n\t"
            var v18 float64 = 3.14159
            _ = v16
            _ = v17
            _ = v18
        } else {
            switch "\n\t" {
            case "\n\t":
                v1 += (v1 >> ((255) & 63))
            default:
                var v19 int8 = 0
                _ = v19
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                v3 = 1
                _ = fi20
            }
            v3 += v3
        }
    default:
        if (2147483647 >= 299) {
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 int8 = -1
                _ = fi21
                _ = v22
            }
            var v23 uint16 = 32767
            var v24 uint = genFn0(v3, (42 != 2), 1)
            _ = v23
            _ = v24
        } else {
            if true {
                if false {
                    if false {
                        v2[1] = "foo bar"
                    }
                    v25 := GenS2{F0: 18446744073709551615, F1: ("a" + "hello")}
                    fmt.Println(v3, v3)
                    _ = v25
                }
                v26 := GenS4{F0: 0, F1: false}
                _ = v26
            } else {
                v1 = 1
                if (!((234 < 0) || false)) {
                    if (42 < -2147483648) {
                        v27 := []int{2147483647}
                        fmt.Println(v3, v1)
                        _ = v27
                    } else {
                        v28 := []int{429, -2147483648, -9223372036854775808}
                        if (32767 != -32768) {
                            v29 := GenS5{F0: "\n\t", F1: 65535, F2: -722}
                            v30 := GenS4{F0: 382, F1: ((((("\n\t" != "") && (183 >= 255)) && true) && true) && true)}
                            _ = v29
                            _ = v30
                        }
                        fmt.Println(v3, v3)
                        _ = v28
                    }
                } else {
                    fmt.Println(v3, v3)
                    var v31 int = -272
                    if (!true) {
                        v32 := GenS1{F0: -1, F1: "hello", F2: 345}
                        _ = v32
                    }
                    _ = v31
                }
                var v33 uint16 = 661
                _ = v33
            }
            v3 = (v3 / (((^(v3))) | 1))
            v1 = -128
        }
    }
    v34 := v2[42]
    for fi35 := 0; fi35 < 20; fi35++ {
        delete(v2, 1)
        var v36 bool = false
        _ = fi35
        _ = v36
    }
    v34 = ""
    if false {
        fmt.Println(v34, v34)
        if false {
            var v37 uint32 = v3
            _ = v37
        }
    }
    var v38 float64 = -2.25
    if (false && false) {
        var v39 bool = false
        _ = v39
    }
    for fi40 := 0; fi40 < 10; fi40++ {
        var v41 uint16 = 177
        _ = fi40
        _ = v41
    }
    var v42 float64 = 3.14159
    v43 := GenS6{F0: 9223372036854775807, F1: 2}
    v44 := []int{1, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v42
    _ = v43
    _ = v44
}
