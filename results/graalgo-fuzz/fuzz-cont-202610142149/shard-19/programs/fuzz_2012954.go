package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint8; F1 uint64; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 bool }
func genFn0(p0 int32, p1 uint, p2 int8) uint32 {
    p0 = 42
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 692, F1: "foo bar", F2: (p2 / (((p2 & p2)) | 1))}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS2{F0: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 303
}
func genFn1(p0 uint64, p1 int64) int {
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: 1000, F1: "foo bar", F2: 24}
        switch "hello" {
        case "\n\t":
            fmt.Println(fi11, p0, p0)
        case "a":
            var v13 float32 = float32((float64(((1.5 + 0.1) * 3.14159)) + (-2.25 / ((float64(1.5) / (1.5 * (((0.1 - -2.25) / float64((-2.25 - float32(1.5)))) * -2.25))) * 0.1))))
            _ = v13
        }
        fmt.Println(p0)
        _ = fi11
        _ = v12
    }
    p0 += p0
    v14 := GenS3{F0: 0, F1: (p0 * (p0 >> ((386) & 63)))}
    _ = p0
    _ = p1
    _ = v14
    return -2147483648
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 uint8 = 128
    var v5 float64 = 0.1
    v6 := []int{-2147483648, -9223372036854775808, 2147483647, -2147483648}
    var v15 int8 = -1
    v4 = (v4 & (v4 + 255))
    switch 372 {
    case 793:
        var v16 uint8 = (v4 - 74)
        _ = v16
    case -554:
        var v17 float64 = v5
        _ = v17
    default:
        v18 := GenS1{F0: 42, F1: "", F2: (v15 ^ v15)}
        _ = v18
    }
    v15 = int8(v1)
    v1 -= v1
    var v19 bool = false
    _ = len(v2)
    if (2147483647 >= 32767) {
        switch "foo bar" {
        case "\n\t":
            switch -2147483648 {
            case 2147483647:
                v20 := GenS2{F0: v19}
                _ = v20
            default:
                var v21 uint = 65535
                _ = v21
            }
        }
    } else {
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 int8 = v15
            _ = fi22
            _ = v23
        }
        v24 := v3["foo bar"]
        var v25 uint64 = 18446744073709551615
        _ = v24
        _ = v25
    }
    switch (("hello" + "hello") + "\n\t") {
    case "hello":
        v4 -= (v4 | (v4 - (v4 - v4)))
    case "\n\t":
        var v26 string = ("\n\t" + "")
        _ = v26
    }
    if (255 <= 352) {
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 float64 = (-2.25 + v5)
            _ = fi27
            _ = v28
        }
    }
    _ = len(v2)
    v5 -= (v5 - (float64(1.5) * v5))
    if v19 {
        if ((2 != uint64(v1)) || (float32(v5) >= 1e10)) {
            v29 := GenS2{F0: (!v19)}
            _ = v29
        }
        switch -715 {
        case 0:
            for fi30 := 0; fi30 < 4; fi30++ {
                fmt.Println(v4, v19)
                for fi31 := 0; fi31 < 13; fi31++ {
                    var v32 int8 = v15
                    switch "hello" {
                    case "foo bar":
                        var v33 int64 = 42
                        _ = v33
                    case "":
                        var v34 uint64 = 606
                        var v35 uint16 = uint16(v4)
                        _ = v34
                        _ = v35
                    case "hello":
                        var v36 float32 = 0.1
                        _ = v36
                    default:
                        switch (-(fi31)) {
                        case 572:
                            switch "" {
                            case "foo bar":
                                fmt.Println(v4, v15)
                            case "hello":
                                v6 = append(v6, 32767)
                            }
                        case 287:
                            switch fi30 {
                            case 32768:
                                if len(v6) > 0 {
                                    v6[0] = (fi30 & (fi30 & genFn1(42, 42)))
                                }
                            case 128:
                                v19 = (v19 || v19)
                            case 4294967295:
                                fmt.Println(v15)
                            }
                        case 32767:
                            var v37 string = "foo bar"
                            _ = v37
                        }
                    }
                    _ = fi31
                    _ = v32
                }
                _ = fi30
            }
        case 1:
            v38 := GenS2{F0: ((v4 / (((v4 * v4)) | 1)) < v4)}
            _ = v38
        default:
            if (42 > 2147483647) {
                for fi39 := 0; fi39 < 9; fi39++ {
                    v40 := GenS2{F0: v19}
                    _ = fi39
                    _ = v40
                }
                if v19 {
                    var v41 uint8 = (^(v4))
                    _ = v41
                } else {
                    var v42 uint8 = v4
                    v42 = v4
                    _ = len(v6)
                    _ = v42
                }
                var v43 uint64 = uint64(v4)
                _ = v43
            } else {
                v15 = (^(v15))
                var v44 int = 42
                _ = v44
            }
        }
    } else {
        v45 := GenS4{F0: -2147483648, F1: "a"}
        v46 := GenS5{F0: v4, F1: 675, F2: genFn1(18446744073709551615, 1)}
        _ = v45
        _ = v46
    }
    switch 702 {
    case 128:
        if (!(1 >= 1)) {
            switch "a" {
            case "hello":
                v19 = v19
            case "":
                v47 := GenS6{F0: (v4 / ((v4) | 1))}
                _ = v47
            default:
                for fi48 := 0; fi48 < 3; fi48++ {
                    v49 := GenS1{F0: 1, F1: ("\n\t" + "\n\t"), F2: (-(v15))}
                    fmt.Println(v4)
                    _ = fi48
                    _ = v49
                }
            }
        } else {
            v6 = append(v6, 32767)
            v2[-1] = 2
            v3["hello"] = 588
        }
    case 1:
        if len(v6) > 0 {
            v6[0] = -338
        }
    }
    if ((v19 || v19) || (v15 > (v15 - (-(v15))))) {
        switch "\n\t" {
        case "":
            var v50 uint32 = 65535
            _ = v50
        default:
            if (v4 != v4) {
                fmt.Println(v15)
                for fi51 := 0; fi51 < 2; fi51++ {
                    fmt.Println(v1, v19)
                    var v52 string = "hello"
                    _ = fi51
                    _ = v52
                }
            } else {
                v4 = 2
                if v19 {
                    v4 += 255
                    v6 = append(v6, genFn1(738, -9223372036854775808))
                } else {
                    v4 += v4
                    _ = len(v3)
                    if v19 {
                        v3[""] = 251
                        v53 := GenS2{F0: (!false)}
                        fmt.Println(v5, v5)
                        _ = v53
                    }
                }
            }
        }
        fmt.Println(v5, v1, v4)
        if ((136 < v4) || v19) {
            fmt.Println(v19, v5, v1)
            if false {
                if false {
                    fmt.Println(v1)
                    v54 := GenS1{F0: 0, F1: ("hello" + "hello"), F2: v15}
                    var v55 int32 = 1
                    _ = v54
                    _ = v55
                } else {
                    if len(v6) > 0 {
                        v6[0] = -199
                    }
                    for fi56 := 0; fi56 < 17; fi56++ {
                        for fi57 := 0; fi57 < 8; fi57++ {
                            var v58 int16 = -1
                            _ = fi57
                            _ = v58
                        }
                        _ = fi56
                    }
                    if v19 {
                        for fi59 := 0; fi59 < 11; fi59++ {
                            v60 := GenS7{F0: 682, F1: 105}
                            v61 := GenS8{F0: v19}
                            v4 = v4
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        if (576 <= 2) {
                            fmt.Println(v5)
                            if true {
                                v5 = v5
                                fmt.Println(v5, v1)
                            } else {
                                var v62 uint = 18446744073709551615
                                switch -902 {
                                case 2147483647:
                                    fmt.Println(v15, v15)
                                default:
                                    v6 = append(v6, -9223372036854775808)
                                }
                                _ = v62
                            }
                            var v63 int = 128
                            _ = v63
                        } else {
                            v4 = 1
                        }
                        v4 += (v4 | 2)
                    }
                }
            }
            if (false || v19) {
                if v19 {
                    v1 = v1
                    v19 = ("\n\t" == "foo bar")
                }
                v2[42] = 794
                for fi64 := 0; fi64 < 5; fi64++ {
                    v65 := v2[(fi64 % (((fi64 * (fi64 * (fi64 - -1)))) | 1))]
                    var v66 uint8 = uint8(v1)
                    v65 = (v65 ^ v65)
                    _ = fi64
                    _ = v65
                    _ = v66
                }
            } else {
                switch "foo bar" {
                case "a":
                    var v67 int64 = 9223372036854775807
                    _ = v67
                case "\n\t":
                    var v68 string = ""
                    _ = v68
                }
            }
        }
    }
    fmt.Println(v5)
    switch 1 {
    case 127:
        fmt.Println(v5, v15, v5)
    default:
        var v69 uint32 = 65535
        _ = v69
    }
    fmt.Println(v5, v15, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
