package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32; F1 int; F2 uint8 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint64; F2 uint }
type GenS8 struct { F0 bool; F1 int8 }
func genFn0(p0 int16) uint64 {
    var v4 uint64 = 801
    if true {
        p0 += -377
    }
    var v5 int16 = int16(v4)
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    _ = v5
    return v4
}
func genFn1() int16 {
    v6 := GenS1{F0: 570, F1: 827}
    var v7 bool = (-9223372036854775808 > -1)
    if (false || (2 <= 128)) {
        switch 127 {
        case 32768:
            v7 = (1.5 >= (3.14159 * (0.1 / -2.25)))
        case 316:
            v8 := []int{128, 231, 720}
            _ = v8
        }
    } else {
        v9 := GenS1{F0: 41, F1: 418}
        if false {
            for fi10 := 0; fi10 < 19; fi10++ {
                fmt.Println(fi10, v7)
                for fi11 := 0; fi11 < 12; fi11++ {
                    v7 = v7
                    v12 := GenS2{F0: 42}
                    _ = fi11
                    _ = v12
                }
                var v13 uint64 = genFn0(-1)
                _ = fi10
                _ = v13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                if true {
                    v7 = v7
                } else {
                    v15 := GenS1{F0: 755, F1: 81}
                    v7 = v7
                    v16 := []int{fi14, (fi14 | (fi14 / ((42) | 1)))}
                    _ = v15
                    _ = v16
                }
                var v17 uint8 = 1
                _ = fi14
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 float64 = 3.14159
                for fi20 := 0; fi20 < 6; fi20++ {
                    v19 = (3.14159 / v19)
                    var v21 uint64 = 18446744073709551615
                    _ = fi20
                    _ = v21
                }
                v7 = (42 <= 9223372036854775807)
                _ = fi18
                _ = v19
            }
        } else {
            var v22 uint = 255
            var v23 int16 = 42
            switch "a" {
            case "":
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := []int{(fi24 * fi24)}
                    var v26 int16 = (v23 << ((1) & 63))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case "\n\t":
                v7 = v7
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    v27 := []int{}
                    _ = v27
                case "hello":
                    v23 += v23
                default:
                    var v28 bool = v7
                    _ = v28
                }
            default:
                var v29 uint64 = 4294967295
                _ = v29
            }
            _ = v22
            _ = v23
        }
        switch 1 {
        case 4294967295:
            switch "hello" {
            case "\n\t":
                if (v7 && v7) {
                    if true {
                        for fi30 := 0; fi30 < 9; fi30++ {
                            v7 = v7
                            v7 = v7
                            _ = fi30
                        }
                    } else {
                        if false {
                            var v31 uint32 = 901
                            var v32 uint64 = 668
                            _ = v31
                            _ = v32
                        } else {
                            v7 = (0.1 != ((0.1 / 3.14159) * 1.5))
                            v33 := []int{4294967295}
                            _ = v33
                        }
                        var v34 int = 0
                        _ = v34
                    }
                } else {
                    switch -661 {
                    case 4294967295:
                        var v35 uint8 = 0
                        _ = v35
                    case -908:
                        var v36 bool = true
                        _ = v36
                    case 2147483647:
                        v37 := []int{-37, 32767, 127}
                        _ = v37
                    }
                }
            case "":
                v7 = v7
            case "foo bar":
                var v38 uint16 = 857
                _ = v38
            default:
                v39 := []int{42}
                _ = v39
            }
        default:
            v40 := []int{394, 2147483647, 1}
            _ = v40
        }
        _ = v9
    }
    fmt.Println(v7)
    _ = v6
    _ = v7
    return -110
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (-80 >= 42)
    var v3 uint64 = 554
    v41 := GenS3{F0: 1, F1: 714, F2: 206}
    if (-2147483648 < 2147483647) {
        var v42 uint16 = 730
        var v43 int16 = 1
        _ = v42
        _ = v43
    } else {
        var v44 int8 = -33
        for fi45 := 0; fi45 < 8; fi45++ {
            for fi46 := 0; fi46 < 7; fi46++ {
                var v47 string = "a"
                _ = fi46
                _ = v47
            }
            _ = fi45
        }
        _ = v44
    }
    if (255 >= 255) {
        v3 = (^(v3))
        for fi48 := 0; fi48 < 1; fi48++ {
            if (-1 >= 100) {
                v49 := []int{int(v3), fi48, fi48, fi48}
                fmt.Println(v3)
                v50 := v1[42]
                _ = v49
                _ = v50
            } else {
                fmt.Println(fi48)
                fmt.Println(fi48)
            }
            _ = fi48
        }
    }
    for fi51 := 0; fi51 < 9; fi51++ {
        v3 += (v3 << ((255) & 63))
        v52 := GenS2{F0: v3}
        _ = fi51
        _ = v52
    }
    var v53 uint64 = (v3 / ((v3) | 1))
    v54 := GenS2{F0: (v3 % ((v3) | 1))}
    _ = len(v1)
    var v55 uint16 = 454
    var v56 uint64 = uint64(v55)
    v57 := GenS4{F0: 511, F1: 126, F2: 1}
    if ((!(!false)) || (!(410 >= 2))) {
        for fi58 := 0; fi58 < 17; fi58++ {
            fmt.Println(v2)
            v3 -= (^(v3))
            var v59 int64 = -914
            _ = fi58
            _ = v59
        }
        _ = len(v1)
    } else {
        if v2 {
            v60 := GenS5{F0: 0, F1: 33}
            _ = v60
        }
        if (false || (((9223372036854775807 > 0) && (v2 && true)) || (42 == 32767))) {
            v61 := GenS6{F0: 0}
            for fi62 := 0; fi62 < 4; fi62++ {
                switch (fi62 / (((fi62 - fi62)) | 1)) {
                case 4294967295:
                    for fi63 := 0; fi63 < 20; fi63++ {
                        fmt.Println(fi62, v53, v53)
                        var v64 uint16 = 0
                        _ = fi63
                        _ = v64
                    }
                case -2147483648:
                    v65 := []int{1, fi62, fi62, -9223372036854775808}
                    _ = v65
                case 988:
                    v56 -= uint64(v55)
                }
                v3 = (v53 / (((^(v53))) | 1))
                _ = fi62
            }
            v53 -= (v53 ^ genFn0(0))
            _ = v61
        } else {
            if (-508 < 9223372036854775807) {
                var v66 float64 = -2.25
                v53 = (v56 ^ 2)
                _ = v66
            }
            v55 = (v55 % (((^(v55))) | 1))
        }
        for fi67 := 0; fi67 < 7; fi67++ {
            v3 -= v3
            for fi68 := 0; fi68 < 8; fi68++ {
                v69 := GenS4{F0: 42, F1: -46, F2: 45}
                switch "a" {
                case "a":
                    v70 := GenS2{F0: v53}
                    _ = v70
                default:
                    fmt.Println(fi68)
                }
                v71 := GenS5{F0: 1, F1: 0}
                _ = fi68
                _ = v69
                _ = v71
            }
            _ = fi67
        }
    }
    fmt.Println(v3, v2, v53)
    switch "a" {
    case "":
        switch 260 {
        case 1:
            fmt.Println(v55, v53)
        case 2147483647:
            for fi72 := 0; fi72 < 9; fi72++ {
                switch ("" + "") {
                case "\n\t":
                    var v73 int16 = 1
                    _ = v73
                case "":
                    v53 = v3
                default:
                    v2 = (false || (uint32(v56) < 0))
                }
                fmt.Println(v53, v56)
                for fi74 := 0; fi74 < 6; fi74++ {
                    switch ("" + "\n\t") {
                    case "a":
                        v75 := GenS3{F0: 255, F1: (^(v55)), F2: 127}
                        _ = v75
                    }
                    _ = fi74
                }
                _ = fi72
            }
        }
    }
    v76 := GenS4{F0: -2147483648, F1: -1, F2: 238}
    for fi77 := 0; fi77 < 18; fi77++ {
        v78 := []int{fi77, fi77, fi77}
        v79 := GenS4{F0: -2147483648, F1: fi77, F2: 2}
        _ = fi77
        _ = v78
        _ = v79
    }
    v80 := GenS7{F0: 2147483647, F1: (v53 % ((0) | 1)), F2: 2}
    v81 := GenS8{F0: v2, F1: -90}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v76
    _ = v80
    _ = v81
}
