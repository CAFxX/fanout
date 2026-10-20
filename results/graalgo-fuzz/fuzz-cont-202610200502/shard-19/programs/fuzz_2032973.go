package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint16, p1 int64, p2 string) uint16 {
    p1 -= p1
    fmt.Println(p0, p0, p0)
    v7 := GenS1{F0: true, F1: false, F2: 92}
    var v8 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p0 >> ((2) & 63))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{909, 879, 2147483647}
    v3 := []string{"", "hello", "foo bar", "a"}
    v4 := make(map[uint32]int)
    v5 := []int{}
    v6 := []string{"\n\t", "\n\t", (((("a" + "hello") + "") + "foo bar") + "a"), ""}
    switch -2147483648 {
    case 42:
        for fi9 := 0; fi9 < 10; fi9++ {
            v4[99] = fi9
            var v10 float32 = 1.5
            _ = fi9
            _ = v10
        }
    case 107:
        var v11 float64 = float64(1.5)
        _ = v11
    default:
        var v12 int32 = 42
        _ = v12
    }
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    if (!(952 <= 870)) {
        switch "a" {
        case "\n\t":
            if false {
                var v13 uint8 = 0
                if true {
                    var v14 int = 2147483647
                    switch "" {
                    case "hello":
                        fmt.Println(v13, v13)
                    }
                    v15 := GenS2{F0: 1, F1: true}
                    _ = v14
                    _ = v15
                } else {
                    var v16 uint16 = 0
                    _ = v16
                }
                _ = v13
            } else {
                for fi17 := 0; fi17 < 10; fi17++ {
                    fmt.Println(fi17)
                    _ = fi17
                }
                var v18 int32 = -896
                v18 = 710
                _ = v18
            }
        case "hello":
            var v19 float64 = 0.1
            _ = v19
        case "":
            if false {
                if (!true) {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        fmt.Println(fi20)
                        v21 := GenS1{F0: (!((fi20 & fi20) >= fi20)), F1: (!false), F2: 100}
                        _ = fi20
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 18; fi22++ {
                    switch fi22 {
                    case 32768:
                        var v23 int8 = 87
                        _ = v23
                    case -542:
                        switch "foo bar" {
                        case "foo bar":
                            var v24 uint64 = 2
                            _ = v24
                        case "\n\t":
                            var v25 uint64 = 1
                            _ = v25
                        case "hello":
                            if false {
                                fmt.Println(fi22)
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    v4[4294967295] = (^(fi26))
                                    if (18446744073709551615 <= 2) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v27 := GenS3{F0: 1}
                                            var v28 float64 = 1.5
                                            _ = v27
                                            _ = v28
                                        default:
                                            v3 = append(v3, "hello")
                                        }
                                    } else {
                                        switch "" {
                                        case "\n\t":
                                            var v29 uint64 = 726
                                            _ = v29
                                        case "a":
                                            var v30 int64 = 1
                                            _ = v30
                                        default:
                                            var v31 int = 127
                                            _ = v31
                                        }
                                        var v32 float64 = (0.1 / (1.5 / 3.14159))
                                        _ = v32
                                    }
                                    _ = fi26
                                }
                            }
                        }
                    default:
                        v4[199] = fi22
                    }
                    _ = fi22
                }
                var v33 uint8 = 128
                _ = v33
            } else {
                switch "hello" {
                case "hello":
                    if false {
                        var v34 uint = 0
                        var v35 uint16 = 65535
                        _ = v34
                        _ = v35
                    }
                case "a":
                    var v36 uint64 = 42
                    _ = v36
                case "":
                    switch "foo bar" {
                    case "foo bar":
                        var v37 uint = 363
                        _ = v37
                    }
                }
                var v38 float64 = (-2.25 / 1.0)
                _ = v38
            }
        default:
            v39 := GenS4{F0: -128, F1: 32767}
            _ = v39
        }
        switch "" {
        case "\n\t":
            v1[""] = 65535
        case "a":
            var v40 int = -9223372036854775808
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 20; fi41++ {
            for fi42 := 0; fi42 < 7; fi42++ {
                v43 := GenS4{F0: 0, F1: genFn0(156, 1, ("a" + "hello"))}
                v44 := GenS3{F0: 2}
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = fi41
        }
    }
    v45 := GenS2{F0: 2147483647, F1: ((0.0 == 0.0) || true)}
    var v46 string = ("hello" + ("a" + "foo bar"))
    v47 := GenS5{F0: 654, F1: 0}
    fmt.Println(v46)
    v46 = v46
    v48 := GenS6{F0: (!false)}
    v46 = v46
    var v49 uint64 = 1
    switch 0 {
    case -1:
        v46 = ("hello" + (v46 + v46))
    case 32768:
        if true {
            v46 = ("hello" + "\n\t")
            var v50 uint64 = v49
            _ = v50
        } else {
            v46 = (("foo bar" + v46) + v46)
        }
    default:
        var v51 float64 = 0.1
        _ = v51
    }
    if true {
        v49 -= 2
    } else {
        if (19 == -128) {
            v52 := v1[v46]
            fmt.Println(v49)
            _ = v52
        }
        v46 = (v46 + (v46 + v46))
        v1["a"] = 2147483648
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
