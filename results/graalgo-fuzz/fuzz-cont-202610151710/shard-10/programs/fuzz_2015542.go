package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int32; F2 int64 }
type GenS8 struct { F0 int8; F1 uint16; F2 uint16 }
func genFn0(p0 int, p1 int64, p2 int16) bool {
    if (true || true) {
        v7 := GenS1{F0: -4}
        var v8 float64 = -2.25
        p2 = p2
        _ = v7
        _ = v8
    }
    p1 -= (^(p1))
    if ((false || (100 != 127)) || (true && ((true || (31 >= 14)) || false))) {
        p1 += p1
        var v9 bool = true
        var v10 uint8 = 70
        _ = v10
        _ = v9
    } else {
        if (!true) {
            for fi11 := 0; fi11 < 5; fi11++ {
                switch -361 {
                case -882:
                    v12 := GenS1{F0: (p2 / (((^(p2))) | 1))}
                    _ = v12
                case 42:
                    v13 := []int{p0, 2147483647}
                    _ = v13
                }
                _ = fi11
            }
        } else {
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := GenS2{F0: 1, F1: 255}
                p1 = (p1 * (-(p1)))
                _ = fi14
                _ = v15
            }
            p2 = (p2 | -174)
        }
        var v16 float32 = 1.5
        _ = v16
    }
    if true {
        v17 := []int{0, (^(p0))}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return (!true)
}
func genFn1(p0 uint64, p1 bool, p2 bool) bool {
    v18 := GenS3{F0: (float32(-2.25) >= (0.0 + 1.5)), F1: 42}
    p2 = (42 > 255)
    p1 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (1.5 <= 1e10)
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "foo bar"
    v3 := []int{128, 128}
    var v4 int64 = 759
    var v5 string = "hello"
    v6 := []int{-104, -310}
    if (4294967295 < 4294967295) {
        v3 = append(v3, 65)
    } else {
        v19 := GenS1{F0: 32767}
        _ = v19
    }
    switch -1 {
    case -9223372036854775808:
        if false {
            v2 = "\n\t"
            var v20 int64 = v4
            _ = v20
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v21 := GenS4{F0: (v2 + ("" + "a"))}
            _ = v21
        }
    case 531:
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 int8 = 0
            if (false || (!false)) {
                var v24 float64 = 0.1
                for fi25 := 0; fi25 < 8; fi25++ {
                    fmt.Println(fi25, v5, v2)
                    if false {
                        v2 = v5
                        var v26 string = "foo bar"
                        v27 := GenS5{F0: 42, F1: 127}
                        _ = v26
                        _ = v27
                    }
                    _ = len(v6)
                    v2 = "hello"
                    var v28 int32 = -1
                    _ = fi25
                    _ = v28
                }
                if ((fi22 % ((-2147483648) | 1)) >= fi22) {
                    var v29 int64 = int64(fi22)
                    fmt.Println(fi22, v4, v2)
                    _ = v29
                } else {
                    if len(v6) > 0 {
                        v6[0] = (^(fi22))
                    }
                    v30 := GenS2{F0: v23, F1: 2}
                    _ = v30
                }
                _ = v24
            } else {
                var v31 int64 = (v4 | 869)
                _ = v31
            }
            if true {
                if (650 == 4294967295) {
                    for fi32 := 0; fi32 < 13; fi32++ {
                        v23 = -1
                        switch (("\n\t" + ("hello" + v5)) + "") {
                        case "":
                            if true {
                                v3 = append(v3, fi22)
                            } else {
                                v2 = v2
                            }
                        case "a":
                            switch v5 {
                            case "hello":
                                v2 = ((v5 + v2) + "hello")
                            case "\n\t":
                                _ = len(v1)
                            case "":
                                switch v2 {
                                case "a":
                                    if false {
                                        if false {
                                            v6 = append(v6, fi22)
                                            if (v5 == v5) {
                                                v2 = ("foo bar" + v2)
                                                v33 := GenS1{F0: -801}
                                                _ = v33
                                            }
                                        }
                                        var v34 int8 = (v23 / (((v23 & (v23 % ((v23) | 1)))) | 1))
                                        _ = v34
                                    }
                                    v1[fi22] = (true || false)
                                case "foo bar":
                                    v2 = v2
                                default:
                                    v35 := GenS6{F0: false}
                                    _ = v35
                                }
                            default:
                                fmt.Println(v23, fi22, fi22)
                                var v36 float32 = 0.0
                                var v37 int32 = -603
                                _ = v36
                                _ = v37
                            }
                        case "hello":
                            if false {
                                v38 := GenS4{F0: v5}
                                v39 := GenS2{F0: v23, F1: 127}
                                v2 = ((v5 + v5) + v5)
                                switch 128 {
                                case 32768:
                                    _ = len(v3)
                                }
                                fmt.Println(fi32, fi22)
                                _ = v38
                                _ = v39
                            } else {
                                v40 := GenS1{F0: 0}
                                v41 := GenS4{F0: v2}
                                _ = v40
                                _ = v41
                            }
                        default:
                            fmt.Println(v4)
                            if false {
                                var v42 int32 = 42
                                var v43 uint8 = 127
                                v42 -= (^(v42))
                                v44 := GenS3{F0: (4294967295 < (-(fi22))), F1: 4294967295}
                                _ = v42
                                _ = v43
                                _ = v44
                            } else {
                                var v45 int64 = (v4 >> ((0) & 63))
                                v46 := GenS3{F0: true, F1: 809}
                                _ = v45
                                _ = v46
                            }
                        }
                        switch "foo bar" {
                        case "hello":
                            var v47 float32 = 1e10
                            _ = v47
                        default:
                            _ = len(v1)
                        }
                        var v48 int64 = v4
                        var v49 int64 = (v4 & (v48 | (^(v48))))
                        _ = fi32
                        _ = v48
                        _ = v49
                    }
                    v2 = ("\n\t" + v5)
                } else {
                    v5 = ""
                }
            }
            _ = fi22
            _ = v23
        }
    default:
        for fi50 := 0; fi50 < 10; fi50++ {
            var v51 uint16 = 42
            _ = fi50
            _ = v51
        }
    }
    v52 := GenS7{F0: (42 == 69), F1: 1, F2: 42}
    v5 = v5
    if false {
        for fi53 := 0; fi53 < 15; fi53++ {
            if true {
                v4 = v4
            }
            _ = fi53
        }
        switch 128 {
        case 4294967295:
            fmt.Println(v2)
        case 1:
            v5 = "hello"
        default:
            v54 := GenS7{F0: (true || (!(!true))), F1: -1, F2: 0}
            _ = v54
        }
        v55 := GenS8{F0: 6, F1: 65535, F2: 42}
        _ = v55
    } else {
        for fi56 := 0; fi56 < 18; fi56++ {
            v57 := GenS3{F0: (false && true), F1: 898}
            _ = fi56
            _ = v57
        }
    }
    v2 = "a"
    var v58 uint32 = 2147483647
    for fi59 := 0; fi59 < 3; fi59++ {
        var v60 int8 = 42
        if true {
            v60 = v60
        } else {
            v61 := GenS7{F0: true, F1: 2147483647, F2: (v4 << ((65535) & 63))}
            var v62 float32 = ((float32(0.1) / -2.25) / 1.0)
            _ = v61
            _ = v62
        }
        if (!false) {
            if (!true) {
                var v63 uint64 = 42
                for fi64 := 0; fi64 < 7; fi64++ {
                    v3 = append(v3, 4294967295)
                    _ = fi64
                }
                v65 := GenS1{F0: -1}
                _ = v63
                _ = v65
            }
            v66 := GenS4{F0: (v5 + (v2 + v5))}
            _ = v66
        }
        _ = fi59
        _ = v60
    }
    switch 101 {
    case 32768:
        var v67 int = -2147483648
        _ = v67
    case 42:
        var v68 float64 = (1e10 - 1.5)
        _ = v68
    }
    v2 = ("hello" + v5)
    v2 = v2
    v5 = v5
    var v69 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v58
    _ = v6
    _ = v69
}
