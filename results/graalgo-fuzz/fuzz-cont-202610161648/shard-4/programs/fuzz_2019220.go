package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 bool; F1 uint; F2 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 bool }
type GenS5 struct { F0 bool; F1 uint8 }
func genFn0(p0 int8, p1 int16) int {
    fmt.Println(p1)
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int = 4294967295
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (v7 * v7)}
    v11 := GenS2{F0: (^(v7)), F1: 255}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    return 128
}
func genFn1(p0 int, p1 int32) uint64 {
    fmt.Println(p1)
    switch p0 {
    case 409:
        for fi12 := 0; fi12 < 9; fi12++ {
            p0 -= 42
            _ = fi12
        }
    case 66:
        p1 -= 1
    case 2147483647:
        var v13 bool = false
        _ = v13
    default:
        if false {
            var v14 string = (("" + "foo bar") + ("foo bar" + (("\n\t" + (("a" + ("a" + "\n\t")) + "")) + "")))
            var v15 uint64 = 18446744073709551615
            _ = v14
            _ = v15
        }
    }
    if false {
        var v16 uint32 = 2
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int8 = 100
            var v19 uint64 = 482
            var v20 int = (p0 << ((2) & 63))
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v16
    }
    v21 := GenS2{F0: 546, F1: 0}
    _ = p0
    _ = p1
    _ = v21
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 4294967295
    v3 := []string{("" + "\n\t"), ("" + ""), "hello", "\n\t"}
    var v4 int32 = 42
    var v5 int8 = -1
    var v6 uint64 = 42
    v22 := GenS2{F0: 32767, F1: 65535}
    if true {
        v6 -= (v6 & genFn1(42, -1))
    } else {
        v23 := GenS3{F0: false, F1: 18446744073709551615, F2: (0 < 2)}
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 bool = false
            _ = fi24
            _ = v25
        }
        if len(v3) > 0 {
            v3[0] = (("a" + ("foo bar" + ("hello" + (("foo bar" + (("" + "foo bar") + "hello")) + "a")))) + "a")
        }
        _ = v23
    }
    v26 := GenS3{F0: true, F1: 790, F2: false}
    for fi27 := 0; fi27 < 9; fi27++ {
        switch "\n\t" {
        case "hello":
            fmt.Println(fi27, fi27, v4)
        case "\n\t":
            var v28 uint16 = 2
            _ = v28
        }
        for fi29 := 0; fi29 < 6; fi29++ {
            delete(v1, 2)
            v2 = v2
            _ = fi29
        }
        v5 -= 0
        _ = fi27
    }
    var v30 int = (-(v2))
    switch "a" {
    case "":
        if (65535 != 32767) {
            for fi31 := 0; fi31 < 10; fi31++ {
                var v32 float32 = (-2.25 + float32(1.5))
                _ = fi31
                _ = v32
            }
            var v33 bool = true
            fmt.Println(v30, v33)
            _ = v33
        } else {
            switch "a" {
            case "hello":
                for fi34 := 0; fi34 < 10; fi34++ {
                    var v35 uint8 = 2
                    switch "a" {
                    case "\n\t":
                        v35 = v35
                    }
                    v35 = (v35 - (v35 - v35))
                    _ = fi34
                    _ = v35
                }
            }
        }
    case "a":
        var v36 uint8 = 1
        _ = v36
    }
    fmt.Println(v6, v2, v4)
    if (-9223372036854775808 < -242) {
        fmt.Println(v5, v4, v6)
        for fi37 := 0; fi37 < 14; fi37++ {
            for fi38 := 0; fi38 < 2; fi38++ {
                v4 -= (v4 * (v4 + 601))
                _ = fi38
            }
            v39 := GenS4{F0: 1, F1: 127, F2: false}
            v4 -= v4
            _ = fi37
            _ = v39
        }
    } else {
        var v40 string = ("foo bar" + "a")
        _ = v40
    }
    var v41 int8 = v5
    if false {
        _ = len(v1)
    }
    var v42 bool = (("foo bar" + "hello") == "hello")
    var v43 float64 = (0.1 - (3.14159 * 0.1))
    switch (v2 | (v2 / (((v30 % (((v30 ^ v30)) | 1))) | 1))) {
    case -2147483648:
        v5 = v41
    case 4294967295:
        v6 = v6
    case 2147483647:
        for fi44 := 0; fi44 < 20; fi44++ {
            for fi45 := 0; fi45 < 11; fi45++ {
                fmt.Println(v43, v6, v43)
                switch "\n\t" {
                case "\n\t":
                    v46 := GenS5{F0: v42, F1: 42}
                    _ = v46
                case "foo bar":
                    v47 := GenS5{F0: (!false), F1: 2}
                    _ = v47
                case "":
                    v42 = true
                default:
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v42, v41, fi44)
                    case "hello":
                        if (v42 && (0 <= 18446744073709551615)) {
                            v48 := GenS5{F0: (!false), F1: 128}
                            _ = v48
                        } else {
                            v5 += -63
                        }
                    default:
                        v49 := GenS2{F0: 42, F1: 65535}
                        _ = v49
                    }
                }
                fmt.Println(v4, v4, v41)
                var v50 uint64 = (v6 + v6)
                var v51 string = ("foo bar" + "a")
                _ = fi45
                _ = v50
                _ = v51
            }
            v52 := GenS3{F0: v42, F1: 4294967295, F2: v42}
            _ = fi44
            _ = v52
        }
    default:
        v1[2147483647] = ("\n\t" + "hello")
    }
    var v53 int16 = -723
    if (!(!(false && true))) {
        var v54 uint = 1
        var v55 float32 = (3.14159 * (0.1 * (3.14159 + 0.1)))
        _ = v54
        _ = v55
    } else {
        v3 = append(v3, "hello")
        v5 += v41
        switch v2 {
        case -12:
            v41 -= (v41 - v5)
        case -210:
            delete(v1, 184)
        }
    }
    v4 = (-(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v53
    _ = v6
}
