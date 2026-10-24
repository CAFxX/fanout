package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 uint }
func genFn0(p0 uint64, p1 int64) uint64 {
    v4 := GenS1{F0: 4294967295}
    p0 += p0
    _ = p0
    _ = p1
    _ = v4
    return 4294967295
}
func genFn1(p0 int64, p1 uint32) uint64 {
    v5 := GenS1{F0: p1}
    for fi6 := 0; fi6 < 17; fi6++ {
        p1 = (p1 | (p1 * p1))
        p0 = (p0 % (((p0 - (p0 << ((228) & 63)))) | 1))
        if (0.1 < -2.25) {
            if true {
                v7 := GenS2{F0: (0.1 <= 0.0), F1: "a"}
                fmt.Println(p0, fi6, fi6)
                _ = v7
            } else {
                for fi8 := 0; fi8 < 15; fi8++ {
                    if false {
                        var v9 int = 127
                        if (1 <= 2) {
                            var v10 string = "\n\t"
                            v11 := []int{fi8}
                            v11 = append(v11, int(p0))
                            _ = v10
                            _ = v11
                        }
                        if true {
                            var v12 uint32 = p1
                            _ = v12
                        } else {
                            var v13 uint = 2
                            v14 := []int{}
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS3{F0: (p0 - 638), F1: 1}
                        var v16 int32 = -622
                        _ = v15
                        _ = v16
                        _ = v9
                    } else {
                        v17 := GenS4{F0: (1000 < 1), F1: -9223372036854775808, F2: 2}
                        v18 := GenS2{F0: true, F1: "foo bar"}
                        _ = v17
                        _ = v18
                    }
                    var v19 uint32 = (p1 * 65535)
                    _ = fi8
                    _ = v19
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    p1 = (p1 - p1)
                    fmt.Println(fi6, fi20, fi6)
                    _ = fi20
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 uint32 = (p1 * p1)
                    p1 = 2
                    switch (("foo bar" + "\n\t") + (("hello" + ("foo bar" + (("foo bar" + (("hello" + (("\n\t" + "hello") + "hello")) + ("a" + "\n\t"))) + (("foo bar" + (("a" + "hello") + "foo bar")) + "foo bar")))) + "hello")) {
                    case "foo bar":
                        p0 -= p0
                    case "hello":
                        v23 := []int{(fi21 & fi6), fi6}
                        _ = v23
                    default:
                        v24 := []int{2147483647, 127, fi21, (fi21 | (fi21 << ((18446744073709551615) & 63)))}
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
            }
            v25 := []int{}
            var v26 int16 = -943
            _ = v25
            _ = v26
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 859
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (^(v1))
    v3 := []int{-9223372036854775808, -938, 1}
    v1 = (v1 & v2)
    fmt.Println(v1, v2)
    if false {
        v1 = (v1 - (v1 | 610))
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 int8 = -128
            _ = fi27
            _ = v28
        }
        v1 = (^(v1))
    } else {
        v1 = v2
        for fi29 := 0; fi29 < 9; fi29++ {
            if len(v3) > 0 {
                v3[0] = 32768
            }
            _ = fi29
        }
        fmt.Println(v1)
    }
    v30 := GenS5{F0: 2147483647, F1: "a"}
    if true {
        v31 := GenS6{F0: -32768}
        _ = v31
    }
    v2 = v2
    v1 -= genFn1(0, 65535)
    v1 += 450
    fmt.Println(v1, v1)
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := GenS7{F0: 9223372036854775807, F1: 963}
        switch "hello" {
        case "foo bar":
            var v34 int64 = -21
            _ = v34
        case "hello":
            for fi35 := 0; fi35 < 4; fi35++ {
                if ((^(v2)) > (v1 & (v2 & (v1 / (((v2 | (v2 / (((v2 - v2)) | 1)))) | 1))))) {
                    v3 = append(v3, fi32)
                }
                _ = fi35
            }
        case "\n\t":
            var v36 uint16 = 715
            _ = v36
        default:
            v2 += v2
        }
        fmt.Println(v1, v1, v1)
        _ = fi32
        _ = v33
    }
    v37 := GenS2{F0: false, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
}
