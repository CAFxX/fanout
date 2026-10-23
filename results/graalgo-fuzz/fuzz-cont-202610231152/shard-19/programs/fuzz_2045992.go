package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 uint32; F2 bool }
type GenS7 struct { F0 string; F1 int64; F2 uint8 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0(p0 uint, p1 int8, p2 int16) int16 {
    var v4 int8 = p1
    for fi5 := 0; fi5 < 16; fi5++ {
        p2 = 1
        fmt.Println(v4, fi5)
        _ = fi5
    }
    var v6 uint8 = 110
    for fi7 := 0; fi7 < 19; fi7++ {
        v4 = 68
        var v8 float64 = 0.0
        switch fi7 {
        case 32768:
            v9 := GenS1{F0: (-(fi7)), F1: "foo bar", F2: 0}
            _ = v9
        case 9223372036854775807:
            for fi10 := 0; fi10 < 5; fi10++ {
                p2 -= p2
                _ = fi10
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return (^(p2))
}
func genFn1(p0 uint32, p1 int16, p2 int16) int8 {
    p2 = (p2 >> ((1) & 63))
    fmt.Println(p1, p1, p0)
    if (0 >= 0) {
        var v11 bool = false
        v12 := []int{42, -2147483648, -950, 32768}
        v13 := GenS2{F0: uint16(p0), F1: 1, F2: 0}
        _ = v11
        _ = v12
        _ = v13
    } else {
        p1 -= 42
    }
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return -14
}

func main() {
    var v1 float64 = -2.25
    v2 := []int{4294967295, 9223372036854775807, 127}
    v3 := make(map[string]string)
    var v14 bool = false
    if (v1 >= float64((float32((1.5 - ((float64(1e10) - (1e10 / v1)) / 0.1))) + ((3.14159 * 1e10) * float32(v1))))) {
        var v15 uint16 = 114
        switch 9223372036854775807 {
        case 32768:
            var v16 bool = v14
            _ = v16
        default:
            var v17 uint64 = 113
            _ = v17
        }
        v18 := GenS3{F0: 65535, F1: 42, F2: ("a" + "")}
        _ = v15
        _ = v18
    } else {
        v14 = ((false || true) && false)
    }
    v1 += v1
    switch 32767 {
    case 130:
        if (((!(((3.14159 + 1e10) + -2.25) >= 1.5)) && true) && (!v14)) {
            fmt.Println(v1, v1)
            var v19 int32 = -1
            var v20 bool = v14
            _ = v19
            _ = v20
        }
    default:
        v14 = (((!v14) && (v14 && (808 <= 781))) && true)
    }
    v1 = v1
    var v21 int16 = 1
    switch ("" + ("\n\t" + "foo bar")) {
    case "":
        for fi22 := 0; fi22 < 2; fi22++ {
            fmt.Println(v14)
            switch "a" {
            case "\n\t":
                _ = len(v2)
            case "":
                fmt.Println(v1, fi22)
            case "a":
                _ = len(v3)
            default:
                switch (("a" + ("a" + "foo bar")) + "\n\t") {
                case "a":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        _ = len(v3)
                        var v24 int64 = 9223372036854775807
                        fmt.Println(v14, v21, v24)
                        _ = fi23
                        _ = v24
                    }
                case "":
                    v14 = false
                case "hello":
                    switch fi22 {
                    case 0:
                        if (9223372036854775807 <= -1) {
                            for fi25 := 0; fi25 < 6; fi25++ {
                                v21 += -858
                                var v26 int8 = -128
                                _ = fi25
                                _ = v26
                            }
                            if (-813 > v21) {
                                v27 := GenS4{F0: 1}
                                _ = v27
                            } else {
                                _ = len(v2)
                                var v28 int = 4294967295
                                v28 += fi22
                                _ = v28
                            }
                            v29 := GenS4{F0: 42}
                            _ = v29
                        } else {
                            if (true || ((0.0 * 1e10) > ((((3.14159 * (0.1 + 0.0)) - 1.5) / (0.1 - 0.0)) / ((((1e10 - 0.1) + 0.1) - (0.1 + 0.0)) - 0.0)))) {
                                for fi30 := 0; fi30 < 18; fi30++ {
                                    delete(v3, "foo bar")
                                    var v31 bool = (true || v14)
                                    fmt.Println(fi30)
                                    _ = fi30
                                    _ = v31
                                }
                                switch 9223372036854775807 {
                                case 128:
                                    fmt.Println(v14)
                                case 1:
                                    v32 := GenS3{F0: 4294967295, F1: -985, F2: "\n\t"}
                                    _ = v32
                                case 335:
                                    if v14 {
                                        v21 -= 952
                                        for fi33 := 0; fi33 < 9; fi33++ {
                                            v21 += v21
                                            if false {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            _ = fi33
                                        }
                                    }
                                default:
                                    for fi34 := 0; fi34 < 11; fi34++ {
                                        v35 := v3[((("hello" + "foo bar") + ("" + "a")) + ("\n\t" + (("" + ("\n\t" + ("foo bar" + ""))) + "\n\t")))]
                                        _ = fi34
                                        _ = v35
                                    }
                                }
                            } else {
                                v14 = v14
                            }
                            v1 = v1
                        }
                    }
                default:
                    v36 := GenS5{F0: 1}
                    _ = v36
                }
            }
            v37 := GenS6{F0: -635, F1: 0, F2: true}
            _ = fi22
            _ = v37
        }
    case "a":
        var v38 bool = true
        _ = v38
    case "\n\t":
        v14 = (v14 && (103 != 128))
    }
    switch "a" {
    case "a":
        v1 = v1
    case "\n\t":
        if (true && (856 > -201)) {
            for fi39 := 0; fi39 < 15; fi39++ {
                v40 := GenS3{F0: 4294967295, F1: -2147483648, F2: ("hello" + "")}
                _ = fi39
                _ = v40
            }
        }
    case "foo bar":
        var v41 float32 = 1e10
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        switch "hello" {
        case "hello":
            if (0 >= 193) {
                if v14 {
                    var v43 uint64 = 0
                    if v14 {
                        fmt.Println(v14, v43)
                    } else {
                        v44 := GenS7{F0: "", F1: 42, F2: 42}
                        v2 = append(v2, (fi42 + -9223372036854775808))
                        _ = v44
                    }
                    _ = v43
                } else {
                    for fi45 := 0; fi45 < 9; fi45++ {
                        fmt.Println(fi45, fi42, fi42)
                        _ = fi45
                    }
                }
            }
        case "foo bar":
            for fi46 := 0; fi46 < 12; fi46++ {
                v1 -= float64(0.1)
                var v47 uint = 0
                var v48 uint16 = 2
                _ = fi46
                _ = v47
                _ = v48
            }
        case "a":
            if (false || (0 < -468)) {
                v1 = (1.5 - float64(((1.5 + 1e10) + (0.1 + 0.1))))
            }
        default:
            v49 := GenS8{F0: 0, F1: 0, F2: 18446744073709551615}
            _ = v49
        }
        if false {
            if v14 {
                v14 = true
                v14 = (128 <= 71)
                v50 := GenS5{F0: 18446744073709551615}
                _ = v50
            }
            fmt.Println(v14, v1)
            v51 := v3[""]
            _ = v51
        }
        _ = fi42
    }
    _ = len(v3)
    fmt.Println(v14)
    fmt.Println(v21)
    v21 += (v21 % ((v21) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
}
