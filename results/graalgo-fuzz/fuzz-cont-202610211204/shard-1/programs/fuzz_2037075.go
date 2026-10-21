package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int8; F2 string }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 string; F1 int32; F2 uint }
type GenS8 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS9 struct { F0 int64 }
func genFn0(p0 int32, p1 string, p2 bool) uint64 {
    p0 = (^(p0))
    if p2 {
        v5 := GenS1{F0: 0, F1: (94 < 666), F2: 252}
        var v6 uint32 = 704
        _ = v5
        _ = v6
    } else {
        v7 := GenS2{F0: ((740 > 668) || true)}
        v8 := []int{2147483647, -2147483648}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615, F1: (p2 || false), F2: 135}
    for fi10 := 0; fi10 < 14; fi10++ {
        p0 = p0
        _ = fi10
    }
    p0 = -327
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}
func genFn1(p0 int16, p1 int16, p2 uint32) uint8 {
    var v11 int16 = p1
    p2 = (^(p2))
    switch 9223372036854775807 {
    case 32768:
        switch (("" + "hello") + "hello") {
        case "hello":
            if (true || false) {
                p0 = 328
                if true {
                    fmt.Println(p0)
                    var v12 float32 = -2.25
                    if false {
                        v13 := []int{}
                        for fi14 := 0; fi14 < 4; fi14++ {
                            switch "a" {
                            case "\n\t":
                                p2 -= (p2 / ((2) | 1))
                            default:
                                p0 -= p0
                            }
                            v15 := GenS2{F0: (((false && (9223372036854775807 < 931)) && false) && (-753 <= 42))}
                            if (false && false) {
                                var v16 int32 = -2147483648
                                v17 := GenS2{F0: ((479 > genFn0(v16, "", (v12 <= (v12 + (v12 * v12))))) || false)}
                                v18 := GenS3{F0: p2}
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                v19 := GenS4{F0: 1}
                                v20 := GenS3{F0: 1}
                                p0 = (v11 + v11)
                                _ = v19
                                _ = v20
                            }
                            _ = fi14
                            _ = v15
                        }
                        fmt.Println(v12, v12, v11)
                        _ = v13
                    } else {
                        v21 := GenS5{F0: 1, F1: 100, F2: "foo bar"}
                        switch 764 {
                        case -2147483648:
                            v22 := GenS6{F0: 42, F1: (1 == 890)}
                            _ = v22
                        case 1:
                            v23 := []int{128}
                            _ = v23
                        case 32768:
                            v24 := []int{-438, 32767, -738, -853}
                            _ = v24
                        }
                        var v25 uint64 = 2
                        _ = v21
                        _ = v25
                    }
                    _ = v12
                }
                if false {
                    var v26 uint8 = 142
                    _ = v26
                }
            }
        case "foo bar":
            p0 = (p0 * -1)
        case "\n\t":
            v27 := GenS7{F0: "", F1: 688, F2: 888}
            _ = v27
        }
    case 500:
        var v28 uint64 = 4294967295
        _ = v28
    }
    v29 := GenS3{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v29
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v4 int8 = -1
    v4 += v4
    v30 := []int{1, -479, 9223372036854775807}
    v4 = v4
    _ = len(v30)
    fmt.Println(v4)
    var v31 int32 = 782
    fmt.Println(v31, v4)
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := GenS8{F0: 18446744073709551615, F1: 0, F2: 581}
        _ = fi32
        _ = v33
    }
    var v34 int16 = 32767
    if false {
        var v35 string = (("foo bar" + ("foo bar" + "a")) + "a")
        for fi36 := 0; fi36 < 16; fi36++ {
            v37 := GenS3{F0: 65535}
            fmt.Println(v34, v35)
            if true {
                v34 = 1
                v38 := GenS9{F0: 287}
                _ = v38
            } else {
                _ = len(v30)
            }
            _ = fi36
            _ = v37
        }
        switch "\n\t" {
        case "":
            v39 := GenS2{F0: false}
            _ = v39
        default:
            v30 = append(v30, 446)
        }
        _ = v35
    } else {
        v34 -= (v34 / ((v34) | 1))
    }
    v40 := GenS2{F0: true}
    v31 += -1
    _ = len(v3)
    switch 9223372036854775807 {
    case 400:
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
    case -963:
        var v41 bool = true
        _ = v41
    case -2147483648:
        for fi42 := 0; fi42 < 15; fi42++ {
            v43 := GenS3{F0: 42}
            _ = fi42
            _ = v43
        }
    default:
        var v44 float32 = 3.14159
        _ = v44
    }
    switch "\n\t" {
    case "":
        if len(v30) > 0 {
            v30[0] = 2147483647
        }
    default:
        fmt.Println(v31, v34, v31)
    }
    if (65535 != 4294967295) {
        switch (("" + ("foo bar" + "")) + "foo bar") {
        case "\n\t":
            _ = len(v30)
        }
    }
    var v45 bool = false
    var v46 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v40
    _ = v45
    _ = v46
}
