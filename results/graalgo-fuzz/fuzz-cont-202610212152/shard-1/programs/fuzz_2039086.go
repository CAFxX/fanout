package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int16; F2 int }
type GenS6 struct { F0 uint8; F1 string; F2 uint16 }
type GenS7 struct { F0 uint8; F1 uint8; F2 string }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int8; F1 bool; F2 bool }
type GenS11 struct { F0 uint8; F1 int; F2 int }
type GenS12 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint32) uint64 {
    var v4 float64 = 1e10
    v5 := GenS1{F0: 702, F1: 375}
    v6 := GenS2{F0: 0}
    if true {
        switch "hello" {
        case "\n\t":
            v7 := []int{456, 930}
            _ = v7
        case "hello":
            if true {
                v8 := []int{}
                _ = v8
            }
        case "foo bar":
            if false {
                fmt.Println(p0, v4)
            } else {
                switch 32768 {
                case -9223372036854775808:
                    for fi9 := 0; fi9 < 19; fi9++ {
                        if false {
                            fmt.Println(v4, fi9, p0)
                            v4 = (0.1 + ((1.5 - 0.0) * 1e10))
                        } else {
                            v10 := []int{fi9, (fi9 | fi9)}
                            v11 := GenS3{F0: 1, F1: 909}
                            _ = v10
                            _ = v11
                        }
                        fmt.Println(v4, fi9, p0)
                        for fi12 := 0; fi12 < 10; fi12++ {
                            switch fi9 {
                            case 1:
                                var v13 uint = 0
                                _ = v13
                            case -1:
                                var v14 int32 = -680
                                _ = v14
                            case 128:
                                var v15 uint = 2
                                _ = v15
                            }
                            switch (fi9 % ((823) | 1)) {
                            case 9223372036854775807:
                                if (!true) {
                                    var v16 int64 = 42
                                    _ = v16
                                } else {
                                    if true {
                                        v17 := GenS4{F0: fi12}
                                        if true {
                                            p0 = 2147483648
                                            if (false && false) {
                                                var v18 uint = 170
                                                v19 := []int{(-(fi12)), fi12}
                                                v20 := GenS5{F0: 0, F1: -32768, F2: fi9}
                                                var v21 uint = (v18 << (((^(v18))) & 63))
                                                v21 += (v18 - (v18 & v21))
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            }
                                            var v22 uint32 = p0
                                            _ = v22
                                        } else {
                                            v4 += float64((1e10 + 1e10))
                                            v23 := GenS6{F0: 42, F1: "a", F2: 255}
                                            var v24 bool = true
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = v17
                                    } else {
                                        v25 := []int{fi12, (fi12 >> ((562) & 63)), fi9, fi12}
                                        _ = v25
                                    }
                                }
                            case 127:
                                switch fi9 {
                                case 32768:
                                    var v26 uint8 = 255
                                    _ = v26
                                }
                            case -2147483648:
                                v27 := []int{(^(fi12))}
                                _ = v27
                            }
                            _ = fi12
                        }
                        _ = fi9
                    }
                }
                v28 := []int{32767, 2147483647, -972}
                if false {
                    fmt.Println(p0, v4)
                } else {
                    var v29 bool = ((float32(3.14159) / float32(v4)) == -2.25)
                    _ = v29
                }
                _ = v28
            }
        default:
            var v30 int16 = 1000
            _ = v30
        }
        switch 9223372036854775807 {
        case 43:
            for fi31 := 0; fi31 < 15; fi31++ {
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 string = ("hello" + (("\n\t" + "hello") + ("foo bar" + ("" + ("foo bar" + "\n\t")))))
                    v34 := GenS7{F0: 240, F1: 255, F2: v33}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = fi31
            }
        default:
            fmt.Println(p0)
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            if (185 > 0) {
                v36 := GenS8{F0: 0, F1: false}
                if (!(42 == 179)) {
                    fmt.Println(fi35)
                }
                _ = v36
            } else {
                if ((42 >= 655) || (false && (((!true) || true) && (-2.25 < (1.5 + v4))))) {
                    var v37 int16 = 42
                    _ = v37
                } else {
                    v4 -= ((float64(-2.25) * (1.5 + float64(1e10))) + v4)
                    if (-89 >= 127) {
                        p0 = p0
                        if (!((fi35 % (((fi35 | fi35)) | 1)) > 32767)) {
                            for fi38 := 0; fi38 < 13; fi38++ {
                                v4 = 0.0
                                v39 := GenS4{F0: (fi35 / ((531) | 1))}
                                _ = fi38
                                _ = v39
                            }
                        }
                        var v40 float32 = ((3.14159 + 0.0) + (-2.25 - 1e10))
                        _ = v40
                    }
                }
                v41 := GenS2{F0: 127}
                _ = v41
            }
            v4 -= v4
            _ = fi35
        }
    } else {
        v42 := []int{-270}
        _ = v42
    }
    var v43 int64 = 42
    _ = p0
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int16, p1 uint32, p2 uint16) int32 {
    fmt.Println(p0, p1, p2)
    var v44 string = ((("hello" + (("hello" + (("hello" + "") + "a")) + "foo bar")) + "") + "hello")
    v45 := GenS9{F0: 2, F1: 42}
    var v46 int32 = 407
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    _ = v45
    _ = v46
    return (v46 | 0)
}

func main() {
    var v1 uint64 = 2
    var v2 uint = 354
    v3 := make(map[int]string)
    for fi47 := 0; fi47 < 12; fi47++ {
        for fi48 := 0; fi48 < 4; fi48++ {
            v49 := []int{-680, (fi48 + (fi48 / ((fi47) | 1))), (fi48 | (fi47 + 2147483647))}
            v3[42] = "a"
            _ = fi48
            _ = v49
        }
        v50 := v3[(fi47 + (fi47 >> ((2) & 63)))]
        _ = fi47
        _ = v50
    }
    v2 += v2
    v1 -= (^(v1))
    v1 = v1
    if true {
        v51 := []int{32768, 4294967295}
        v52 := GenS10{F0: -1, F1: false, F2: (2147483648 != 65535)}
        v1 = v1
        _ = v51
        _ = v52
    } else {
        fmt.Println(v1)
        var v53 uint = 0
        _ = v53
    }
    v1 -= genFn0(65535)
    v54 := GenS11{F0: 1, F1: 0, F2: -289}
    v3[-1] = "hello"
    v1 = genFn0(871)
    v55 := GenS12{F0: 32767, F1: 100}
    for fi56 := 0; fi56 < 16; fi56++ {
        for fi57 := 0; fi57 < 11; fi57++ {
            v2 += (v2 | 447)
            switch (fi56 % ((-1) | 1)) {
            case -85:
                v58 := GenS10{F0: -1, F1: ((!true) || false), F2: true}
                _ = v58
            }
            v2 -= v2
            _ = fi57
        }
        v59 := GenS12{F0: 355, F1: -128}
        _ = fi56
        _ = v59
    }
    switch ("hello" + "a") {
    case "a":
        fmt.Println(v2, v2)
    case "\n\t":
        fmt.Println(v1, v1)
    case "hello":
        var v60 bool = true
        _ = v60
    default:
        var v61 uint16 = 658
        _ = v61
    }
    switch 32767 {
    case 973:
        if true {
            v2 -= v2
            v62 := []int{}
            switch -9223372036854775808 {
            case -1:
                fmt.Println(v1, v1)
            default:
                if (!false) {
                    for fi63 := 0; fi63 < 17; fi63++ {
                        _ = len(v62)
                        _ = fi63
                    }
                    if false {
                        _ = len(v3)
                        v62 = append(v62, 0)
                    }
                } else {
                    var v64 uint = 42
                    _ = v64
                }
            }
            _ = v62
        }
    case -800:
        for fi65 := 0; fi65 < 12; fi65++ {
            v1 -= (v1 + (v1 + v1))
            v1 -= 951
            _ = fi65
        }
    default:
        v66 := []int{9223372036854775807, 837, 2147483647}
        _ = v66
    }
    _ = len(v3)
    var v67 int32 = -240
    var v68 float64 = (3.14159 + 1e10)
    var v69 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v67
    _ = v68
    _ = v69
}
