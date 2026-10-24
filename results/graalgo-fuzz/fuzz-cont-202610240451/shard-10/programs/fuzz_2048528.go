package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint64; F2 bool }
type GenS6 struct { F0 int8; F1 int16; F2 bool }
type GenS7 struct { F0 string; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 int32; F1 int64; F2 string }
func genFn0(p0 string, p1 int16, p2 int8) uint {
    fmt.Println(p2, p1, p2)
    v6 := GenS1{F0: 510, F1: 42, F2: -74}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 939
}

func main() {
    var v1 uint = 1
    var v2 int8 = 117
    var v3 int8 = v2
    var v4 uint8 = 42
    var v5 int32 = -502
    v3 = v2
    var v7 int16 = -699
    var v8 uint16 = 28
    var v9 int16 = v7
    if true {
        if (-2147483648 <= 42) {
            v10 := GenS2{F0: v5, F1: 42}
            v11 := GenS2{F0: 897, F1: 1}
            _ = v10
            _ = v11
        }
        if (2147483648 != 332) {
            v12 := GenS1{F0: genFn0(("" + (("\n\t" + "foo bar") + "a")), 569, 97), F1: -508, F2: 1}
            v13 := []int{-399, -35, 32767, 290}
            _ = v12
            _ = v13
        }
    }
    v14 := GenS2{F0: 369, F1: 65535}
    if false {
        if false {
            var v15 uint64 = uint64(v1)
            fmt.Println(v8)
            if (!(-128 == int8(v15))) {
                var v16 int32 = (v5 % ((v5) | 1))
                v17 := GenS1{F0: (v1 / (((v1 * genFn0("foo bar", v9, v2))) | 1)), F1: 1, F2: 0}
                _ = v16
                _ = v17
            }
            _ = v15
        } else {
            var v18 float64 = ((((0.0 + 0.0) + (-2.25 * 0.0)) - (0.0 * 0.1)) - ((1.5 - ((float64(1.5) - (1.5 / 1.5)) / 0.1)) * -2.25))
            _ = v18
        }
        v19 := GenS1{F0: genFn0("hello", 42, (v3 / ((v3) | 1))), F1: 42, F2: -910}
        _ = v19
    } else {
        v2 = v3
        if (4294967295 >= 2147483647) {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 float64 = (((0.1 * 0.0) + (-2.25 - -2.25)) + -2.25)
                v22 := GenS3{F0: "hello", F1: (v4 > (v4 - v4)), F2: "\n\t"}
                _ = fi20
                _ = v21
                _ = v22
            }
            if false {
                var v23 uint64 = 2
                if true {
                    if ("a" < "\n\t") {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            var v25 uint64 = v23
                            _ = fi24
                            _ = v25
                        }
                    }
                }
                switch ("a" + "a") {
                case "a":
                    if (int8(v1) <= v3) {
                        v26 := []int{32768, 1}
                        v7 = v7
                        _ = v26
                    }
                case "hello":
                    v1 -= genFn0("foo bar", 1000, v3)
                case "foo bar":
                    fmt.Println(v4, v9, v2)
                default:
                    v1 += v1
                }
                _ = v23
            }
            v2 = (v2 >> ((v1) & 63))
        }
    }
    if (v3 < (v3 + int8(v7))) {
        v4 = (^(v4))
    } else {
        var v27 int = 414
        switch "a" {
        case "foo bar":
            v28 := []int{(v27 ^ (v27 - (v27 / ((279) | 1)))), v27}
            _ = v28
        case "hello":
            fmt.Println(v27)
        case "\n\t":
            v29 := []int{v27}
            _ = v29
        default:
            var v30 uint8 = (v4 / (((v4 / ((v4) | 1))) | 1))
            _ = v30
        }
        _ = v27
    }
    var v31 float64 = 1.5
    v1 += (v1 | v1)
    if false {
        for fi32 := 0; fi32 < 16; fi32++ {
            v4 += (v4 - 127)
            v33 := GenS2{F0: v5, F1: 1}
            v5 = (v5 + -2147483648)
            _ = fi32
            _ = v33
        }
    }
    v34 := []int{368, 2147483647, -2147483648}
    if ((false && false) || true) {
        if false {
            switch -2147483648 {
            case 910:
                v1 = (^(v1))
            case 4294967295:
                v3 = v3
            default:
                var v35 uint = v1
                _ = v35
            }
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 uint32 = 2147483648
                for fi38 := 0; fi38 < 18; fi38++ {
                    if false {
                        var v39 float64 = -2.25
                        v34 = append(v34, 0)
                        _ = v39
                    } else {
                        var v40 int64 = int64(fi36)
                        if false {
                            var v41 uint = genFn0("hello", v7, v3)
                            v42 := GenS4{F0: (v41 & (v1 >> ((genFn0((("\n\t" + ("\n\t" + ("a" + ("foo bar" + "hello")))) + "foo bar"), 32767, v2)) & 63)))}
                            v40 += (^(v40))
                            _ = v41
                            _ = v42
                        }
                        v5 = -679
                        v43 := GenS1{F0: (v1 * v1), F1: (fi36 >> ((42) & 63)), F2: -465}
                        _ = v40
                        _ = v43
                    }
                    _ = len(v34)
                    if ((-413 >= (v5 >> ((v1) & 63))) || false) {
                        var v44 uint64 = 593
                        v45 := GenS3{F0: (("foo bar" + "") + "foo bar"), F1: true, F2: "a"}
                        _ = v44
                        _ = v45
                    }
                    switch "a" {
                    case "\n\t":
                        var v46 float32 = float32(v31)
                        _ = v46
                    case "foo bar":
                        v9 -= 1000
                    }
                    v2 = 0
                    _ = fi38
                }
                _ = fi36
                _ = v37
            }
            for fi47 := 0; fi47 < 1; fi47++ {
                switch (fi47 % ((fi47) | 1)) {
                case 1:
                    switch (-(fi47)) {
                    case 0:
                        for fi48 := 0; fi48 < 14; fi48++ {
                            switch fi48 {
                            case -976:
                                fmt.Println(fi47)
                            }
                            _ = fi48
                        }
                    case 42:
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                    case -2147483648:
                        var v49 bool = (!((v2 + (v3 + v2)) > (v3 * (v2 - 1))))
                        _ = v49
                    }
                case 127:
                    _ = len(v34)
                default:
                    v50 := GenS5{F0: uint(v8), F1: 553, F2: (0 == uint32(v4))}
                    _ = v50
                }
                switch -9223372036854775808 {
                case 32768:
                    v51 := GenS6{F0: v3, F1: (v7 & -584), F2: (1 > (v8 / ((42) | 1)))}
                    _ = v51
                case -1:
                    if true {
                        fmt.Println(v7, v31, v9)
                        v52 := GenS7{F0: "", F1: (v4 ^ v4), F2: v4}
                        if false {
                            v53 := GenS8{F0: -97, F1: ("foo bar" <= "\n\t")}
                            _ = v53
                        } else {
                            fmt.Println(v8)
                        }
                        _ = v52
                    }
                }
                v54 := GenS9{F0: (v5 + int32(v7)), F1: 0, F2: "foo bar"}
                _ = fi47
                _ = v54
            }
        }
    }
    fmt.Println(v5, v3)
    v55 := GenS9{F0: (v5 % ((v5) | 1)), F1: 0, F2: "\n\t"}
    v9 -= v9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v55
    _ = v7
    _ = v8
    _ = v9
}
