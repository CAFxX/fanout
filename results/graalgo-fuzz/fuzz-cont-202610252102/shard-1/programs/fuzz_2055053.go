package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 uint8; F1 string }
func genFn0() string {
    v6 := []int{}
    var v7 int64 = 642
    var v8 float64 = (float64(1.5) - 0.1)
    _ = v6
    _ = v7
    _ = v8
    return ""
}
func genFn1(p0 uint, p1 uint8, p2 bool) string {
    for fi9 := 0; fi9 < 14; fi9++ {
        p0 = p0
        p2 = (!p2)
        for fi10 := 0; fi10 < 16; fi10++ {
            p0 = 2
            var v11 float32 = (1e10 / (float32(0.1) * (1.5 / (((0.0 - 1e10) / -2.25) / 1.0))))
            v12 := []int{(fi9 >> ((p0) & 63))}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 int = -82
    v2 := []uint32{}
    v3 := make(map[int]bool)
    v4 := []string{"a"}
    v5 := []uint32{0, 755, 4294967295, 374}
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 float32 = (1.5 * ((3.14159 + (float32(0.0) / float32(0.1))) / (float32(1.5) + 0.0)))
        if len(v2) > 0 {
            v2[0] = 52
        }
        _ = fi13
        _ = v14
    }
    switch 9223372036854775807 {
    case 1:
        v3[v1] = (!(658 == 448))
    case 463:
        _ = len(v3)
    default:
        v1 = v1
    }
    fmt.Println(v1)
    if len(v5) > 0 {
        v5[0] = 436
    }
    switch 42 {
    case -984:
        v15 := GenS1{F0: 42}
        _ = v15
    case -2147483648:
        switch "foo bar" {
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = "a"
            }
        case "":
            fmt.Println(v1)
        case "hello":
            var v16 uint8 = 127
            _ = v16
        }
    default:
        v1 = 32767
    }
    var v17 string = "hello"
    _ = len(v5)
    if ((false || (0.0 > 1e10)) && (!true)) {
        if false {
            if ((v1 >> ((530) & 63)) >= v1) {
                switch v1 {
                case 42:
                    v17 = v17
                case 32768:
                    var v18 uint8 = 255
                    _ = v18
                }
                if true {
                    if true {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            v17 = "hello"
                            if (666 < 32767) {
                                var v20 int64 = -9223372036854775808
                                v1 += 2147483647
                                _ = v20
                            } else {
                                var v21 uint8 = 255
                                for fi22 := 0; fi22 < 20; fi22++ {
                                    if true {
                                        v23 := GenS2{F0: 697, F1: 160, F2: -8}
                                        fmt.Println(fi22, v1, fi19)
                                        fmt.Println(v21)
                                        var v24 uint = 65535
                                        v24 = 4294967295
                                        _ = v23
                                        _ = v24
                                    }
                                    switch (fi22 >> ((0) & 63)) {
                                    case -1:
                                        switch (v1 / (((fi19 << ((18446744073709551615) & 63))) | 1)) {
                                        case 127:
                                            v1 = 32768
                                        default:
                                            if (65535 == 2) {
                                                var v25 int8 = 1
                                                _ = len(v2)
                                                if true {
                                                    v17 = v17
                                                    var v26 int16 = -1
                                                    _ = v26
                                                } else {
                                                    switch (((v17 + v17) + v17) + "foo bar") {
                                                    case "a":
                                                        v27 := GenS1{F0: 2}
                                                        _ = v27
                                                    case "":
                                                        var v28 bool = false
                                                        _ = v28
                                                    default:
                                                        if true {
                                                            if false {
                                                                var v29 int32 = -595
                                                                _ = v29
                                                            }
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v17 = ""
                                                            }
                                                            v1 += fi22
                                                        }
                                                    }
                                                }
                                                _ = v25
                                            }
                                        }
                                    case -95:
                                        switch fi19 {
                                        case -9223372036854775808:
                                            v21 = v21
                                        case 127:
                                            v3[(-(v1))] = false
                                        case 0:
                                            v21 = 1
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                v30 := GenS3{F0: 0, F1: 4294967295, F2: -821}
                                                _ = len(v4)
                                                var v31 uint16 = 32767
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                    case -9223372036854775808:
                                        _ = len(v5)
                                    default:
                                        fmt.Println(v17, v17)
                                    }
                                    _ = fi22
                                }
                                _ = v21
                            }
                            if (!(!(false && true))) {
                                switch (^(fi19)) {
                                case 42:
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        v17 = v17
                                        v17 = v17
                                        switch (fi32 + 128) {
                                        case 128:
                                            v1 = (v1 - (fi19 >> ((255) & 63)))
                                        default:
                                            v17 = v17
                                        }
                                        _ = fi32
                                    }
                                case 128:
                                    v33 := GenS4{F0: 208}
                                    _ = v33
                                case 2147483647:
                                    v17 = v17
                                default:
                                    if false {
                                        var v34 uint32 = 0
                                        var v35 int8 = 1
                                        _ = v34
                                        _ = v35
                                    } else {
                                        for fi36 := 0; fi36 < 7; fi36++ {
                                            fmt.Println(v1, fi36, v17)
                                            _ = fi36
                                        }
                                        v37 := GenS3{F0: 42, F1: 386, F2: 2147483647}
                                        _ = v37
                                    }
                                }
                            }
                            _ = fi19
                        }
                    }
                    if (255 <= 255) {
                        if (false && true) {
                            fmt.Println(v17)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            var v38 uint64 = 905
                            _ = v38
                        }
                        v17 = (v17 + "a")
                        var v39 float64 = 1e10
                        _ = v39
                    }
                } else {
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v17 = ("hello" + ("a" + v17))
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v1, fi40, fi40)
                        } else {
                            if (4294967295 != 2) {
                                var v41 string = "\n\t"
                                v17 = "\n\t"
                                v42 := GenS5{F0: 255}
                                _ = v41
                                _ = v42
                            } else {
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    v44 := GenS6{F0: -198, F1: 676}
                                    _ = fi43
                                    _ = v44
                                }
                            }
                            _ = len(v3)
                            var v45 uint32 = 2147483647
                            _ = v45
                        }
                        v46 := GenS7{F0: v1, F1: 42, F2: 494}
                        _ = fi40
                        _ = v46
                    }
                }
            }
            if (811 == -9223372036854775808) {
                for fi47 := 0; fi47 < 12; fi47++ {
                    v48 := GenS7{F0: 831, F1: 18446744073709551615, F2: 255}
                    var v49 int32 = -2147483648
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            } else {
                v50 := GenS5{F0: 2}
                _ = v50
            }
            fmt.Println(v1, v1)
        } else {
            var v51 int32 = 42
            v51 = v51
            if (42 <= 42) {
                v17 = (v17 + (v17 + (v17 + ("hello" + "\n\t"))))
            }
            _ = v51
        }
        if false {
            for fi52 := 0; fi52 < 19; fi52++ {
                if ((((4294967295 > 1) && false) || (-1 == -1)) || (-424 == 32768)) {
                    v53 := GenS6{F0: 1, F1: 1}
                    if false {
                        switch v17 {
                        case "\n\t":
                            v54 := GenS7{F0: (fi52 & v1), F1: 18446744073709551615, F2: 255}
                            _ = v54
                        case "hello":
                            _ = len(v5)
                        case "foo bar":
                            v1 = v1
                        default:
                            v55 := GenS8{F0: 42, F1: 127, F2: 42}
                            _ = v55
                        }
                        var v56 uint16 = 255
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 6; fi57++ {
                        var v58 uint = 42
                        v59 := GenS9{F0: 42, F1: false}
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    _ = v53
                } else {
                    v3[(v1 ^ -1)] = (2 <= 114)
                }
                v17 = "hello"
                _ = fi52
            }
            for fi60 := 0; fi60 < 8; fi60++ {
                fmt.Println(fi60, v1, fi60)
                var v61 int16 = 1000
                _ = fi60
                _ = v61
            }
            switch v17 {
            case "hello":
                var v62 int32 = -202
                _ = v62
            default:
                v63 := GenS10{F0: 42, F1: ""}
                _ = v63
            }
        }
        var v64 uint32 = 847
        _ = v64
    }
    v1 = 0
    v1 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3[v1] = (1 <= 122)
    if (!true) {
        _ = len(v3)
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v17, v1)
    v17 = v17
    v3[4294967295] = (-128 <= 42)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
