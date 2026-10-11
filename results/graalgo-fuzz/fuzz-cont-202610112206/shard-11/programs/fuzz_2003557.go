package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint32; F1 int16; F2 string }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int16; F1 int8; F2 int64 }
type GenS7 struct { F0 uint64; F1 int32; F2 string }
func genFn0(p0 bool, p1 bool, p2 bool) int64 {
    p0 = (42 >= 42)
    switch "a" {
    case "a":
        var v4 uint16 = 1
        _ = v4
    case "foo bar":
        var v5 int32 = 2147483647
        _ = v5
    case "\n\t":
        var v6 int = 617
        _ = v6
    }
    if false {
        p1 = (!p2)
    }
    v7 := []int{-1, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []uint32{2147483647, 4294967295}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 32767
    if (-1 != -9223372036854775808) {
        if false {
            var v8 int16 = (v3 << ((255) & 63))
            _ = v8
        }
    } else {
        var v9 string = "\n\t"
        _ = v9
    }
    var v10 int16 = v3
    switch 32767 {
    case -515:
        fmt.Println(v10)
    case -746:
        for fi11 := 0; fi11 < 2; fi11++ {
            switch "foo bar" {
            case "hello":
                switch ((((((((("foo bar" + "a") + ("foo bar" + ("a" + "hello"))) + "") + ("hello" + "")) + "") + "") + "foo bar") + ("\n\t" + "hello")) + "foo bar") {
                case "":
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v10 -= 0
                        _ = fi12
                    }
                case "foo bar":
                    if false {
                        fmt.Println(v3, v3, v3)
                    }
                case "a":
                    if (64 > 81) {
                        v13 := GenS1{F0: 100, F1: 117}
                        var v14 float32 = 3.14159
                        v15 := GenS1{F0: 1, F1: 0}
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                default:
                    fmt.Println(v10, v3)
                }
            case "":
                fmt.Println(v10)
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 2
                }
            default:
                v10 = 32767
            }
            if false {
                if (-87 >= 42) {
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                    var v16 float32 = 1e10
                    var v17 int64 = -782
                    _ = v16
                    _ = v17
                } else {
                    var v18 uint32 = 556
                    _ = v18
                }
            }
            _ = fi11
        }
    case 127:
        v19 := GenS2{F0: 0, F1: -9223372036854775808}
        _ = v19
    default:
        switch 1 {
        case 32767:
            for fi20 := 0; fi20 < 18; fi20++ {
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := v2[2]
                    _ = len(v2)
                    v23 := GenS1{F0: 127, F1: -2147483648}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 19; fi24++ {
                    fmt.Println(v3)
                    v3 = (v3 - (v3 * 32767))
                    if true {
                        switch ("" + (("hello" + "\n\t") + ("\n\t" + ""))) {
                        case "":
                            v10 += v3
                        case "a":
                            if false {
                                if true {
                                    fmt.Println(v3, fi24, fi20)
                                    var v25 uint64 = 430
                                    _ = v25
                                } else {
                                    if true {
                                        var v26 uint = 255
                                        v3 = (v10 - v10)
                                        _ = v26
                                    } else {
                                        v27 := GenS3{F0: 2, F1: false}
                                        _ = v27
                                    }
                                    v28 := GenS3{F0: 32767, F1: false}
                                    var v29 int16 = (v10 >> ((4294967295) & 63))
                                    _ = v28
                                    _ = v29
                                }
                                if (v3 > 858) {
                                    var v30 string = ""
                                    var v31 int64 = -9223372036854775808
                                    switch 717 {
                                    case 289:
                                        v10 += 0
                                    case -211:
                                        v32 := GenS4{F0: 0, F1: v10, F2: ((v30 + (v30 + ("hello" + ((v30 + (((v30 + v30) + v30) + "hello")) + v30)))) + ("" + ""))}
                                        _ = v32
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = v30
                                    _ = v31
                                } else {
                                    var v33 uint = 42
                                    fmt.Println(fi20, v10)
                                    _ = v33
                                }
                            } else {
                                fmt.Println(v3, v10)
                                v3 -= (^(v3))
                            }
                        }
                        v34 := GenS5{F0: 1, F1: 18446744073709551615}
                        v35 := GenS6{F0: (-(v3)), F1: 42, F2: -1}
                        _ = v34
                        _ = v35
                    } else {
                        var v36 uint32 = 4294967295
                        var v37 int16 = v10
                        v36 += (^(v36))
                        _ = v36
                        _ = v37
                    }
                    _ = fi24
                }
                _ = fi20
            }
        case 128:
            if true {
                if false {
                    switch -1 {
                    case -275:
                        v10 = 0
                    case -1:
                        v3 = v3
                    case 4294967295:
                        v38 := GenS5{F0: 845, F1: 2}
                        _ = v38
                    }
                    v3 -= (v10 + v10)
                } else {
                    v3 = -172
                    switch 127 {
                    case 0:
                        switch "\n\t" {
                        case "foo bar":
                            v10 -= -440
                        case "\n\t":
                            for fi39 := 0; fi39 < 7; fi39++ {
                                for fi40 := 0; fi40 < 2; fi40++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch ((((("a" + "") + ("a" + "")) + ("a" + "hello")) + "hello") + "") {
                                        case "\n\t":
                                            switch -1 {
                                            case -2147483648:
                                                switch "hello" {
                                                case "\n\t":
                                                    v41 := v2[614]
                                                    _ = v41
                                                }
                                            case 467:
                                                v42 := GenS6{F0: (v10 + v3), F1: 127, F2: 42}
                                                _ = v42
                                            case -9223372036854775808:
                                                if true {
                                                    if (false && (-9223372036854775808 == genFn0(false, false, (1e10 != 0.1)))) {
                                                        v10 += int16(fi40)
                                                    } else {
                                                        var v43 float64 = 1.5
                                                        _ = v43
                                                    }
                                                }
                                            }
                                        case "foo bar":
                                            fmt.Println(v3, fi39, v10)
                                        }
                                    case "a":
                                        fmt.Println(fi40, v10)
                                    case "\n\t":
                                        v1 = append(v1, 351)
                                    default:
                                        var v44 bool = false
                                        _ = v44
                                    }
                                    var v45 float32 = (((1.5 - (0.0 / 1.0)) / (1.5 / 0.1)) / 0.1)
                                    _ = fi40
                                    _ = v45
                                }
                                var v46 int8 = 42
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    v46 -= v46
                                    _ = fi47
                                }
                                _ = fi39
                                _ = v46
                            }
                        default:
                            v3 -= (v3 ^ (v3 ^ (v10 - (-(v10)))))
                        }
                    default:
                        var v48 int = -2147483648
                        _ = v48
                    }
                }
                v49 := GenS2{F0: 100, F1: -713}
                _ = v49
            } else {
                v50 := v2[2]
                v10 = (v3 << ((4294967295) & 63))
                _ = v50
            }
        case 9223372036854775807:
            v51 := GenS5{F0: 42, F1: 42}
            _ = v51
        }
    }
    v10 = (-(v10))
    var v52 int = 619
    for fi53 := 0; fi53 < 3; fi53++ {
        var v54 int64 = 0
        v54 += v54
        v55 := GenS7{F0: 18446744073709551615, F1: 2147483647, F2: "hello"}
        _ = fi53
        _ = v54
        _ = v55
    }
    fmt.Println(v3, v52)
    var v56 uint16 = 1
    var v57 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v52
    _ = v56
    _ = v57
}
