package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int; F1 uint32; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 uint; F2 int64 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 int32; F1 int64 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{(fi6 >> ((4294967295) & 63)), fi6}
        fmt.Println(fi6)
        var v8 uint = 4294967295
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        switch 128 {
        case -1:
            for fi10 := 0; fi10 < 9; fi10++ {
                switch "hello" {
                case "a":
                    var v11 float32 = 0.0
                    _ = v11
                case "hello":
                    if false {
                        v12 := []int{(fi9 | fi10), fi9}
                        switch fi9 {
                        case 0:
                            var v13 bool = false
                            _ = v13
                        case 127:
                            if true {
                                if (1 != 698) {
                                    fmt.Println(fi9, fi9)
                                } else {
                                    if false {
                                        var v14 int = (^(fi10))
                                        _ = v14
                                    } else {
                                        var v15 int32 = 42
                                        _ = v15
                                    }
                                    if (65535 > 42) {
                                        var v16 int = fi10
                                        _ = v16
                                    } else {
                                        var v17 uint8 = 86
                                        if len(v12) > 0 {
                                            v12[0] = -717
                                        }
                                        _ = v17
                                    }
                                }
                                switch (-(fi10)) {
                                case 2147483647:
                                    v18 := GenS1{F0: (fi9 - (fi10 * 127)), F1: (!true)}
                                    var v19 uint32 = 2
                                    _ = v18
                                    _ = v19
                                }
                                if (1.5 <= 3.14159) {
                                    var v20 int64 = 0
                                    v21 := GenS1{F0: fi9, F1: false}
                                    if len(v12) > 0 {
                                        v12[0] = (fi10 | fi10)
                                    }
                                    fmt.Println(fi10, fi10)
                                    _ = v20
                                    _ = v21
                                } else {
                                    v22 := GenS1{F0: (fi10 + -9223372036854775808), F1: (42 == 42)}
                                    if (1 <= 933) {
                                        fmt.Println(fi10, fi10)
                                    }
                                    _ = v22
                                }
                            }
                        default:
                            v23 := GenS2{F0: fi9, F1: 752, F2: (^(fi9))}
                            _ = v23
                        }
                        _ = v12
                    }
                }
                var v24 uint32 = 113
                _ = fi10
                _ = v24
            }
        case 97:
            if true {
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int8 = 0
                    v27 := GenS3{F0: 65535}
                    v28 := []int{}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            } else {
                if (!true) {
                    v29 := []int{}
                    _ = v29
                }
                switch fi9 {
                case 162:
                    v30 := []int{128, (^(fi9)), fi9}
                    _ = v30
                case 4294967295:
                    for fi31 := 0; fi31 < 2; fi31++ {
                        var v32 float64 = 0.1
                        _ = fi31
                        _ = v32
                    }
                }
                var v33 int16 = -331
                _ = v33
            }
        default:
            var v34 bool = ((1e10 * float32(float64(0.1))) == 1.5)
            _ = v34
        }
        v35 := []int{(fi9 / ((127) | 1)), (fi9 << ((4294967295) & 63)), 9223372036854775807}
        v35 = append(v35, (fi9 >> ((4294967295) & 63)))
        _ = fi9
        _ = v35
    }
    var v36 int64 = -731
    _ = v36
    return 42
}
func genFn1() string {
    v37 := GenS4{F0: -1}
    var v38 int32 = 1
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 uint = 1
        _ = fi39
        _ = v40
    }
    v41 := []int{0, -866, 0}
    v38 += (v38 | v38)
    _ = v37
    _ = v38
    _ = v41
    return "foo bar"
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]int)
    v3 := []uint32{}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    if (65535 <= 1) {
        for fi42 := 0; fi42 < 2; fi42++ {
            if (871 <= -143) {
                v43 := GenS3{F0: 732}
                _ = v43
            }
            v1 = (^(v1))
            _ = len(v5)
            _ = fi42
        }
    } else {
        if (-9223372036854775808 == 127) {
            for fi44 := 0; fi44 < 17; fi44++ {
                v1 -= v1
                _ = fi44
            }
        } else {
            v45 := GenS3{F0: 0}
            _ = len(v4)
            switch "hello" {
            case "hello":
                var v46 uint8 = 127
                _ = v46
            default:
                var v47 float32 = 0.0
                _ = v47
            }
            _ = v45
        }
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        if (57 == 42) {
            v49 := GenS4{F0: 298}
            _ = v49
        }
        _ = fi48
    }
    switch 42 {
    case 9223372036854775807:
        fmt.Println(v1)
    case 4294967295:
        for fi50 := 0; fi50 < 14; fi50++ {
            if true {
                var v51 uint64 = 973
                var v52 uint64 = 1
                _ = v51
                _ = v52
            } else {
                fmt.Println(fi50, v1)
            }
            _ = fi50
        }
    case 0:
        if false {
            v2[7] = 127
            if (617 > 1) {
                for fi53 := 0; fi53 < 13; fi53++ {
                    if (true && (true && false)) {
                        _ = len(v4)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v1 = 0
                        v54 := GenS5{F0: 124, F1: v1}
                        _ = v54
                    }
                    _ = fi53
                }
                v2[32768] = -206
                v1 = (v1 / (((v1 * v1)) | 1))
            } else {
                var v55 uint64 = 2
                _ = v55
            }
        }
    }
    for fi56 := 0; fi56 < 13; fi56++ {
        fmt.Println(fi56)
        _ = fi56
    }
    switch "foo bar" {
    case "":
        var v57 uint16 = 65535
        _ = v57
    case "foo bar":
        for fi58 := 0; fi58 < 16; fi58++ {
            var v59 bool = (((!(1 != 4294967295)) && (128 <= 213)) || true)
            v60 := GenS3{F0: genFn0()}
            _ = fi58
            _ = v59
            _ = v60
        }
    }
    v61 := GenS6{F0: 2, F1: ""}
    v62 := GenS7{F0: (v1 ^ (v1 << ((0) & 63)))}
    fmt.Println(v1)
    var v63 uint64 = 2
    if false {
        if true {
            var v64 uint32 = 656
            _ = v64
        } else {
            var v65 float64 = -2.25
            switch "a" {
            case "foo bar":
                v5[genFn0()] = (false || true)
            }
            _ = v65
        }
    }
    if (!false) {
        v66 := GenS8{F0: 1, F1: 0, F2: -92}
        if true {
            delete(v4, -392)
            var v67 float32 = 0.1
            v68 := GenS1{F0: 32767, F1: false}
            _ = v67
            _ = v68
        } else {
            var v69 uint8 = 255
            _ = v69
        }
        _ = v66
    } else {
        if (446 <= 256) {
            v70 := v2[197]
            v2[v70] = (v70 * (v70 | v70))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v70
        } else {
            fmt.Println(v63)
        }
        var v71 int32 = 734
        var v72 int16 = v1
        _ = v71
        _ = v72
    }
    v73 := GenS9{F0: 2147483647, F1: 32767}
    switch "" {
    case "\n\t":
        var v74 uint16 = 0
        _ = v74
    case "foo bar":
        var v75 int64 = -1
        _ = v75
    case "hello":
        v1 -= (v1 % ((v1) | 1))
    }
    var v76 uint64 = 0
    delete(v4, -2147483648)
    v76 += v76
    v77 := GenS10{F0: -2147483648, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v63:", v63)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v73
    _ = v76
    _ = v77
}
