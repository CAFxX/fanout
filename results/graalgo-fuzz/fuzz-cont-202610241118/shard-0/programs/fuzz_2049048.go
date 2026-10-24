package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint8; F1 int16 }
func genFn0() uint {
    var v7 bool = true
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 = v7
        _ = fi8
    }
    _ = v7
    return 255
}
func genFn1() uint64 {
    for fi9 := 0; fi9 < 15; fi9++ {
        switch (("\n\t" + "a") + "\n\t") {
        case "a":
            switch "" {
            case "a":
                if (1 > 18446744073709551615) {
                    v10 := []int{(fi9 | fi9), fi9}
                    var v11 float64 = 1.5
                    _ = v10
                    _ = v11
                }
            case "\n\t":
                if (9 != 1) {
                    switch "" {
                    case "\n\t":
                        var v12 int = fi9
                        _ = v12
                    case "foo bar":
                        var v13 int32 = 2147483647
                        _ = v13
                    }
                    v14 := GenS1{F0: 2147483647}
                    switch "foo bar" {
                    case "foo bar":
                        var v15 uint16 = 608
                        _ = v15
                    case "\n\t":
                        var v16 int64 = 0
                        _ = v16
                    }
                    _ = v14
                } else {
                    var v17 uint64 = 4294967295
                    _ = v17
                }
            case "foo bar":
                var v18 bool = ((true && true) && (("foo bar" + ("" + "hello")) < ("hello" + "hello")))
                _ = v18
            default:
                var v19 float64 = 0.0
                _ = v19
            }
        case "foo bar":
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := GenS2{F0: -32768}
                v22 := GenS1{F0: 2147483648}
                fmt.Println(fi20, fi20)
                var v23 bool = true
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        case "":
            v24 := GenS3{F0: 42, F1: 1, F2: (-2.25 < 0.1)}
            _ = v24
        default:
            switch 562 {
            case -9223372036854775808:
                var v25 bool = false
                _ = v25
            case 9223372036854775807:
                var v26 uint8 = 42
                _ = v26
            }
        }
        _ = fi9
    }
    v27 := []int{42, 678}
    _ = v27
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(fi28, fi28)
        v29 := GenS2{F0: 32767}
        _ = fi28
        _ = v29
    }
    _ = len(v1)
    var v30 int8 = -1
    for fi31 := 0; fi31 < 13; fi31++ {
        v2 = (v2 / ((v2) | 1))
        if (false && false) {
            v2 -= (^(v2))
            if true {
                if false {
                    fmt.Println(fi31, v2, fi31)
                }
                fmt.Println(v30)
            }
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        }
        if true {
            var v32 bool = false
            v2 = v2
            var v33 uint64 = 2
            _ = v32
            _ = v33
        } else {
            for fi34 := 0; fi34 < 19; fi34++ {
                v2 -= (v2 >> ((2) & 63))
                var v35 int = (-(fi34))
                if true {
                    v30 = (v30 >> ((1) & 63))
                }
                switch (("\n\t" + "foo bar") + "foo bar") {
                case "foo bar":
                    var v36 int16 = -32768
                    _ = v36
                case "hello":
                    if true {
                        v35 = v35
                    } else {
                        v30 = v30
                        switch "" {
                        case "":
                            _ = len(v4)
                        }
                        v35 = (fi31 << ((766) & 63))
                    }
                case "\n\t":
                    if ((genFn1() != 106) && (v2 > v2)) {
                        var v37 int16 = 42
                        if (0 >= 2) {
                            var v38 uint8 = 73
                            var v39 uint32 = 2147483648
                            _ = v38
                            _ = v39
                        }
                        v40 := GenS4{F0: (-(v30)), F1: v2}
                        _ = v37
                        _ = v40
                    }
                default:
                    v4 = append(v4, 2)
                }
                _ = len(v5)
                var v41 int = fi34
                v2 -= 42
                switch "hello" {
                case "a":
                    v1[(v41 & 32767)] = 2147483647
                case "\n\t":
                    switch (("a" + (("hello" + "a") + "a")) + "a") {
                    case "a":
                        var v42 bool = false
                        _ = v42
                    case "\n\t":
                        if ((true || (!(false || ((-(v2)) == (v2 ^ v2))))) && true) {
                            v30 = (^(v30))
                        }
                    }
                case "":
                    var v43 uint32 = 65535
                    _ = v43
                }
                v44 := GenS5{F0: 174, F1: -179}
                _ = fi34
                _ = v35
                _ = v41
                _ = v44
            }
            switch "\n\t" {
            case "a":
                if (1 >= 1) {
                    v45 := GenS5{F0: 255, F1: 1}
                    for fi46 := 0; fi46 < 7; fi46++ {
                        fmt.Println(fi46)
                        switch "" {
                        case "a":
                            v47 := GenS4{F0: v30, F1: v2}
                            _ = v47
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v48 float32 = 0.1
                            _ = v48
                        case "a":
                            fmt.Println(v2, fi31)
                        case "hello":
                            v30 -= v30
                        default:
                            v30 = (v30 * -111)
                        }
                        fmt.Println(v2, fi31, fi31)
                        fmt.Println(fi31, v30)
                        switch "hello" {
                        case "a":
                            v49 := GenS2{F0: 0}
                            _ = v49
                        }
                        _ = fi46
                    }
                    _ = v45
                }
            }
        }
        _ = fi31
    }
    v4 = append(v4, 0)
    var v50 int16 = 32767
    v50 = 1
    v51 := GenS3{F0: v50, F1: -113, F2: ((4294967295 == 536) && true)}
    var v52 int8 = v30
    fmt.Println(v30)
    _ = len(v6)
    v52 = v30
    fmt.Println(v30, v30)
    v2 = (v2 - (v2 % ((1) | 1)))
    delete(v3, 720)
    v53 := GenS5{F0: 42, F1: v50}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
