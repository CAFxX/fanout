package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint8 }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 int32; F1 string; F2 int32 }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS12 struct { F0 int; F1 int16 }
func genFn0(p0 uint32) uint8 {
    v5 := GenS1{F0: (22 == 127), F1: 0, F2: 82}
    var v6 int = -194
    p0 += (^(p0))
    if true {
        v7 := []int{}
        _ = v7
    }
    v8 := GenS2{F0: v6, F1: ("\n\t" + ("\n\t" + "")), F2: 1}
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return 81
}
func genFn1(p0 string) int32 {
    v9 := GenS3{F0: (2147483648 >= 1)}
    v10 := GenS4{F0: 100}
    p0 = (p0 + ((p0 + ("foo bar" + "a")) + ("" + (p0 + (p0 + "hello")))))
    p0 = ("a" + ("\n\t" + p0))
    _ = p0
    _ = v10
    _ = v9
    return -561
}

func main() {
    v1 := []uint32{2147483648, 42, 2147483648, 4294967295}
    v2 := []uint32{4294967295, 4294967295, 2}
    v3 := []uint32{1, 2147483648, 293, 2}
    v4 := []int{32767, 9223372036854775807, -19, 128}
    switch ("a" + (("" + ((((("hello" + "foo bar") + ((("hello" + "foo bar") + "") + "hello")) + "hello") + ("hello" + "foo bar")) + "a")) + "")) {
    case "\n\t":
        switch "\n\t" {
        case "":
            var v11 int32 = 2147483647
            _ = v11
        }
    case "foo bar":
        var v12 uint8 = 255
        _ = v12
    case "hello":
        v13 := GenS5{F0: -1, F1: false}
        _ = v13
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v14 := GenS1{F0: false, F1: 2, F2: genFn0(2147483647)}
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS6{F0: genFn0(4294967295), F1: 85}
        var v17 int8 = 42
        v18 := GenS7{F0: -2147483648, F1: ("" + (("foo bar" + ("a" + "a")) + ((("a" + (("\n\t" + "hello") + "\n\t")) + "") + "hello"))), F2: 42}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint8 = 1
    var v20 int64 = -597
    if true {
        var v21 int = 2147483647
        _ = v21
    } else {
        v19 -= v19
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v20, v20)
    switch ("a" + "") {
    case "a":
        var v22 uint = 735
        _ = v22
    case "\n\t":
        var v23 int8 = -1
        _ = v23
    case "hello":
        var v24 uint = 0
        _ = v24
    default:
        fmt.Println(v20)
    }
    if (-1 > -1) {
        v19 += v19
        fmt.Println(v19)
        if ((v19 % ((v19) | 1)) != v19) {
            var v25 bool = (v20 > -975)
            var v26 uint = 2
            var v27 int64 = 42
            _ = v25
            _ = v26
            _ = v27
        }
    } else {
        v28 := GenS6{F0: genFn0(961), F1: (v19 % ((v19) | 1))}
        if ((2147483647 < 1000) && false) {
            for fi29 := 0; fi29 < 7; fi29++ {
                if true {
                    fmt.Println(fi29, fi29, fi29)
                    if (!false) {
                        if false {
                            fmt.Println(v20)
                        }
                    } else {
                        var v30 uint16 = 65535
                        v19 += v19
                        _ = v30
                    }
                } else {
                    v31 := GenS7{F0: -2147483648, F1: "", F2: -2147483648}
                    switch ("foo bar" + ("a" + ("a" + "\n\t"))) {
                    case "hello":
                        var v32 bool = false
                        _ = v32
                    case "a":
                        v33 := GenS1{F0: (1 == 947), F1: 255, F2: genFn0(293)}
                        _ = v33
                    default:
                        fmt.Println(fi29, v20, fi29)
                    }
                    v20 = (v20 / ((0) | 1))
                    _ = v31
                }
                v19 -= (v19 << ((602) & 63))
                v34 := GenS1{F0: false, F1: 2, F2: (v19 / ((v19) | 1))}
                _ = fi29
                _ = v34
            }
        } else {
            if (true && true) {
                v35 := GenS8{F0: 1, F1: (v20 ^ (-(v20)))}
                v3 = append(v3, 42)
                if false {
                    v36 := GenS9{F0: (v20 & 893)}
                    _ = v36
                }
                _ = v35
            }
            v37 := GenS10{F0: 0}
            fmt.Println(v19)
            _ = v37
        }
        v38 := GenS11{F0: v19, F1: 2, F2: 255}
        _ = v28
        _ = v38
    }
    var v39 int8 = -42
    var v40 float32 = (((float32((-2.25 / 0.1)) * (3.14159 / 1.0)) + ((float32(((((3.14159 - 0.1) + -2.25) * ((-2.25 * -2.25) / 1e10)) - -2.25)) / 3.14159) / 1e10)) * 0.0)
    switch -296 {
    case 32768:
        if len(v4) > 0 {
            v4[0] = 0
        }
    default:
        var v41 int = 4294967295
        _ = v41
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        v43 := GenS12{F0: 128, F1: -809}
        var v44 uint16 = 1
        _ = fi42
        _ = v43
        _ = v44
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
