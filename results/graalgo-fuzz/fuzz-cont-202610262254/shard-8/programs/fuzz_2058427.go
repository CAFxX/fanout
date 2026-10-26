package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint; F1 int8; F2 int32 }
type GenS6 struct { F0 int8; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 int }
func genFn0() uint32 {
    v6 := GenS1{F0: (!false), F1: 1000}
    var v7 int32 = -1
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int8 = -83
        for fi10 := 0; fi10 < 16; fi10++ {
            if false {
                if true {
                    var v11 string = ("a" + "")
                    if (false || (!(v9 <= 127))) {
                        var v12 int8 = v9
                        v13 := []int{127, fi8, (fi10 ^ 109), (^(fi8))}
                        v9 -= v12
                        _ = v12
                        _ = v13
                    } else {
                        var v14 string = v11
                        _ = v14
                    }
                    _ = v11
                }
                v9 = v9
            } else {
                v7 = v7
                fmt.Println(fi8)
                v7 = v7
                fmt.Println(v7, fi10, fi8)
                v15 := []int{}
                _ = v15
            }
            var v16 uint8 = 99
            v16 = (v16 * v16)
            var v17 int64 = int64(v7)
            v18 := GenS2{F0: int32(v17)}
            _ = fi10
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return 546
}
func genFn1(p0 int) uint16 {
    p0 = (p0 | (p0 & (p0 | -135)))
    for fi19 := 0; fi19 < 15; fi19++ {
        fmt.Println(p0, p0)
        fmt.Println(p0, fi19)
        _ = fi19
    }
    _ = p0
    return 48
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -855
    var v3 uint16 = 65535
    var v4 bool = (-1 > (v2 << ((65535) & 63)))
    v5 := make(map[int]uint32)
    if v4 {
        fmt.Println(v3, v3)
    } else {
        v20 := GenS3{F0: (v3 | v3), F1: -2147483648}
        switch 127 {
        case 32767:
            v21 := GenS2{F0: -898}
            _ = v21
        default:
            var v22 uint8 = 117
            _ = v22
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 float64 = ((0.1 + 1e10) + 3.14159)
            _ = fi23
            _ = v24
        }
        _ = v20
    }
    var v25 int = 32767
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        for fi26 := 0; fi26 < 20; fi26++ {
            if v4 {
                v27 := GenS2{F0: -2147483648}
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 int64 = (v2 + -961)
                    switch "hello" {
                    case "foo bar":
                        if (!((0 < 1) || false)) {
                            if (((3.14159 * 0.1) * (1.5 + 1e10)) >= 1.5) {
                                v29 = (v2 >> ((2) & 63))
                                v3 = genFn1(v25)
                            } else {
                                delete(v5, 119)
                                v25 = fi28
                                var v30 int32 = 42
                                if v4 {
                                    v30 = (v30 * v30)
                                    v31 := []int{}
                                    _ = v31
                                } else {
                                    delete(v5, (fi26 + 127))
                                    if (v4 || (!v4)) {
                                        var v32 int16 = int16(v30)
                                        v33 := GenS1{F0: (!v4), F1: int16(v3)}
                                        v34 := []int{(v25 + fi26)}
                                        v32 -= (^(v32))
                                        var v35 uint64 = 4294967295
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 bool = ((fi28 << ((42) & 63)) == 0)
                                        var v37 float64 = 0.1
                                        if v4 {
                                            var v38 string = "\n\t"
                                            var v39 string = v38
                                            switch (v38 + (v38 + v39)) {
                                            case "\n\t":
                                                if (true || (false && v36)) {
                                                    var v40 uint64 = 446
                                                    var v41 uint64 = uint64(v3)
                                                    v1[2] = 644
                                                    _ = v40
                                                    _ = v41
                                                }
                                            case "":
                                                if v36 {
                                                    var v42 uint = 255
                                                    _ = v42
                                                } else {
                                                    v43 := GenS4{F0: (^(v3)), F1: 42, F2: (v3 ^ (v3 & v3))}
                                                    v36 = ((v30 & (v30 * (v30 % (((v30 % ((0) | 1))) | 1)))) < (v30 & (v30 % (((v30 | v30)) | 1))))
                                                    switch (fi26 & fi26) {
                                                    case 1:
                                                        var v44 float32 = 1e10
                                                        _ = v44
                                                    case -1:
                                                        v45 := GenS4{F0: 32767, F1: v30, F2: (v3 >> ((418) & 63))}
                                                        _ = v45
                                                    }
                                                    _ = v43
                                                }
                                            }
                                            _ = v38
                                            _ = v39
                                        } else {
                                            switch ("foo bar" + "foo bar") {
                                            case "a":
                                                v1[42] = 64
                                            case "hello":
                                                var v46 int16 = -1
                                                _ = v46
                                            default:
                                                v47 := []int{(v25 >> ((704) & 63)), fi28, -1}
                                                _ = v47
                                            }
                                            v36 = v4
                                            if v36 {
                                                var v48 float32 = 0.0
                                                _ = v48
                                            } else {
                                                _ = len(v1)
                                                switch "hello" {
                                                case "":
                                                    var v49 float32 = float32(0.1)
                                                    if (-1 >= -128) {
                                                        if true {
                                                            v30 = v30
                                                        }
                                                        v49 = 0.0
                                                    } else {
                                                        fmt.Println(v25)
                                                        if v4 {
                                                            v50 := GenS5{F0: 1, F1: 25, F2: v30}
                                                            var v51 int = (-(v25))
                                                            v4 = v36
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        var v52 uint16 = 1
                                                        _ = v52
                                                    }
                                                    var v53 int = v25
                                                    var v54 int32 = v30
                                                    _ = v49
                                                    _ = v53
                                                    _ = v54
                                                case "hello":
                                                    v55 := []int{fi26}
                                                    _ = v55
                                                default:
                                                    switch 4294967295 {
                                                    case 847:
                                                        fmt.Println(v2, fi26)
                                                    case 420:
                                                        if (v37 <= v37) {
                                                            switch "hello" {
                                                            case "hello":
                                                                v1[2147483648] = 47
                                                            case "a":
                                                                var v56 int16 = 1000
                                                                _ = v56
                                                            case "\n\t":
                                                                var v57 uint = 18446744073709551615
                                                                _ = v57
                                                            }
                                                            v4 = v4
                                                            var v58 uint8 = 255
                                                            _ = v58
                                                        }
                                                    default:
                                                        v59 := GenS6{F0: 100, F1: 176, F2: (v29 - 482)}
                                                        _ = v59
                                                    }
                                                }
                                                var v60 uint64 = 42
                                                _ = v60
                                            }
                                            v29 = 9223372036854775807
                                            v61 := GenS6{F0: 100, F1: 42, F2: 42}
                                            _ = v61
                                        }
                                        _ = v36
                                        _ = v37
                                    }
                                    if v4 {
                                        v62 := []int{(fi28 & fi28)}
                                        v3 = genFn1((v25 * (fi26 % (((^(v25))) | 1))))
                                        _ = v62
                                    }
                                }
                                _ = v30
                            }
                        } else {
                            _ = len(v5)
                            fmt.Println(v4, v29)
                            v1[genFn0()] = 65535
                        }
                    case "hello":
                        v63 := []int{-647}
                        _ = v63
                    }
                    _ = fi28
                    _ = v29
                }
                _ = v27
            } else {
                v3 = genFn1(fi26)
            }
            v4 = v4
            _ = fi26
        }
    }
    var v64 int16 = 1000
    v65 := GenS7{F0: 4294967295}
    var v66 uint64 = 299
    v67 := GenS6{F0: -1, F1: 127, F2: -1}
    v68 := GenS1{F0: true, F1: v64}
    fmt.Println(v64, v3, v2)
    switch "a" {
    case "":
        for fi69 := 0; fi69 < 11; fi69++ {
            if v4 {
                v70 := GenS7{F0: v66}
                _ = v70
            } else {
                switch "foo bar" {
                case "a":
                    if (v4 || ((-226 >= 1) || (v3 <= (^(v3))))) {
                        if (uint(v66) <= 959) {
                            if v4 {
                                fmt.Println(v64)
                            } else {
                                v3 -= (^(v3))
                                var v71 uint8 = uint8(v66)
                                _ = v71
                            }
                        } else {
                            if v4 {
                                if v4 {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v72 := GenS8{F0: (v66 ^ v66)}
                                        _ = v72
                                    }
                                    v73 := GenS4{F0: genFn1(649), F1: -737, F2: genFn1(fi69)}
                                    _ = v73
                                } else {
                                    if v4 {
                                        v1[777] = 2
                                        v74 := GenS9{F0: ("foo bar" + "hello")}
                                        _ = v74
                                    }
                                    v75 := []int{fi69, (v25 * fi69)}
                                    v25 -= v25
                                    _ = v75
                                }
                            }
                            if true {
                                v66 -= (v66 + (v66 & (^(v66))))
                            }
                            var v76 uint16 = uint16(v66)
                            _ = v76
                        }
                    } else {
                        var v77 uint64 = v66
                        var v78 int8 = 42
                        _ = v77
                        _ = v78
                    }
                case "":
                    v2 = v2
                }
                v79 := GenS10{F0: uint8(v3), F1: fi69}
                _ = v79
            }
            fmt.Println(v64, v25)
            var v80 int64 = v2
            _ = fi69
            _ = v80
        }
    case "a":
        fmt.Println(v3, v4)
    case "foo bar":
        var v81 float64 = ((0.0 + -2.25) * 1e10)
        _ = v81
    default:
        v5[(v25 * (v25 >> ((1) & 63)))] = 487
    }
    for fi82 := 0; fi82 < 13; fi82++ {
        var v83 int64 = (v2 ^ v2)
        if (true || (!true)) {
            var v84 uint32 = 0
            _ = v84
        } else {
            for fi85 := 0; fi85 < 18; fi85++ {
                v25 = (v25 | v25)
                v3 = (v3 & v3)
                v83 = v2
                _ = fi85
            }
            switch 9223372036854775807 {
            case 23:
                v64 = -32768
            case -796:
                if true {
                    v1[47] = 444
                    if (!(!(18446744073709551615 >= (v66 & 18446744073709551615)))) {
                        if v4 {
                            if true {
                                if v4 {
                                    if true {
                                        delete(v1, 2147483648)
                                        v4 = ((-883 > (v83 - (v83 * (^(v2))))) && false)
                                        for fi86 := 0; fi86 < 8; fi86++ {
                                            v83 += v83
                                            var v87 string = "foo bar"
                                            v4 = v4
                                            var v88 int = 40
                                            _ = fi86
                                            _ = v87
                                            _ = v88
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            v4 = (-1 >= 1)
                                        case "\n\t":
                                            if v4 {
                                                v2 += -9223372036854775808
                                            } else {
                                                for fi89 := 0; fi89 < 10; fi89++ {
                                                    var v90 int16 = v64
                                                    _ = fi89
                                                    _ = v90
                                                }
                                                v66 = v66
                                                fmt.Println(v3)
                                            }
                                        case "foo bar":
                                            var v91 float64 = 0.1
                                            _ = v91
                                        }
                                    }
                                    for fi92 := 0; fi92 < 20; fi92++ {
                                        fmt.Println(v66, v83)
                                        _ = fi92
                                    }
                                    if v4 {
                                        v25 = 42
                                        v25 = (v25 - (v25 * v25))
                                        _ = len(v5)
                                    }
                                } else {
                                    if false {
                                        v93 := GenS4{F0: v3, F1: 1, F2: v3}
                                        delete(v1, 0)
                                        var v94 uint8 = 1
                                        _ = v93
                                        _ = v94
                                    }
                                    v3 = (v3 >> ((958) & 63))
                                }
                            }
                            if (18446744073709551615 != v66) {
                                var v95 int16 = -219
                                v2 += 42
                                _ = v95
                            }
                            switch ("hello" + ("" + "")) {
                            case "\n\t":
                                var v96 int8 = -15
                                _ = v96
                            }
                        }
                    } else {
                        v66 -= (v66 - v66)
                        v66 = (v66 / (((v66 + 18446744073709551615)) | 1))
                        var v97 uint64 = 836
                        _ = v97
                    }
                }
            case 0:
                var v98 bool = ((3.14159 * 3.14159) >= (-2.25 * 0.0))
                _ = v98
            }
        }
        for fi99 := 0; fi99 < 3; fi99++ {
            var v100 uint8 = 157
            var v101 uint32 = genFn0()
            _ = fi99
            _ = v100
            _ = v101
        }
        _ = fi82
        _ = v83
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
