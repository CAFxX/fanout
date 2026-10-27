package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint32; F1 int; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int; F1 bool; F2 int16 }
func genFn0() bool {
    var v6 bool = (!false)
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{(fi8 % ((fi7) | 1)), fi8}
            fmt.Println(v6)
            _ = fi8
            _ = v9
        }
        var v10 int = (fi7 / ((fi7) | 1))
        v11 := []int{v10, (fi7 / ((v10) | 1)), (fi7 >> ((255) & 63)), (v10 & v10)}
        _ = fi7
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: -9223372036854775808}
    _ = v12
    _ = v6
    return v6
}

func main() {
    var v1 int8 = 1
    var v2 float64 = (1.5 + 0.1)
    var v3 int8 = (v1 * 100)
    var v4 uint8 = 0
    v5 := make(map[uint32]int)
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint32 = 890
            _ = fi14
            _ = v15
        }
        if false {
            switch 128 {
            case 0:
                _ = len(v5)
            case 2147483647:
                v16 := GenS2{F0: v4}
                _ = v16
            }
        }
        _ = fi13
    }
    if true {
        for fi17 := 0; fi17 < 18; fi17++ {
            if (!(1 < 153)) {
                fmt.Println(v1, fi17)
                v1 = 1
                if (false || (!false)) {
                    switch ("" + "") {
                    case "\n\t":
                        if (0.0 <= v2) {
                            v18 := GenS3{F0: 172, F1: v4, F2: v4}
                            v2 = (v2 + v2)
                            _ = v18
                        }
                    case "":
                        v4 = v4
                    default:
                        _ = len(v5)
                    }
                    for fi19 := 0; fi19 < 13; fi19++ {
                        fmt.Println(v1, v3)
                        if (32767 == -32768) {
                            var v20 int64 = 529
                            _ = v20
                        }
                        fmt.Println(v4)
                        fmt.Println(fi17)
                        _ = fi19
                    }
                    v21 := []int{}
                    _ = v21
                }
            } else {
                v1 = -70
                fmt.Println(v1, v1, fi17)
            }
            var v22 int32 = 0
            _ = fi17
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := v5[90]
        v3 -= v1
        _ = fi23
        _ = v24
    }
    if false {
        var v25 bool = false
        _ = v25
    }
    fmt.Println(v1, v3)
    v26 := GenS3{F0: (v4 | 2), F1: v4, F2: (v4 << ((0) & 63))}
    v3 += v3
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 float64 = 0.1
        _ = fi27
        _ = v28
    }
    if false {
        v29 := GenS4{F0: int64(v4), F1: 1}
        fmt.Println(v1)
        for fi30 := 0; fi30 < 7; fi30++ {
            v31 := []int{fi30, fi30}
            v32 := GenS1{F0: 1}
            v4 = (v4 - 31)
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    } else {
        delete(v5, 4294967295)
        v1 = (-(v3))
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        if false {
            v34 := []int{fi33}
            v35 := GenS5{F0: 673, F1: fi33, F2: (^(v4))}
            fmt.Println(v3)
            _ = v34
            _ = v35
        }
        _ = fi33
    }
    if false {
        fmt.Println(v1)
        var v36 uint = 0
        v3 = -1
        _ = v36
    } else {
        var v37 float32 = (1e10 + -2.25)
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 int = (fi38 << ((18446744073709551615) & 63))
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    v40 := GenS1{F0: -1}
    for fi41 := 0; fi41 < 11; fi41++ {
        if true {
            switch fi41 {
            case 4294967295:
                var v42 int64 = 9223372036854775807
                _ = v42
            case 0:
                var v43 uint = 2
                _ = v43
            case 1:
                switch "hello" {
                case "":
                    switch (^(fi41)) {
                    case 42:
                        v44 := []int{fi41}
                        _ = v44
                    default:
                        if false {
                            var v45 int8 = v3
                            if true {
                                if false {
                                    for fi46 := 0; fi46 < 12; fi46++ {
                                        var v47 float32 = float32(float64((1e10 - ((1.5 / ((((0.1 * 1e10) * (0.1 - (1.5 + (1.5 - 0.0)))) / -2.25) - 1e10)) + 1.5))))
                                        v48 := GenS4{F0: 0, F1: -259}
                                        v2 = v2
                                        if false {
                                            v4 = v4
                                            switch (fi41 * fi41) {
                                            case 0:
                                                v49 := GenS2{F0: v4}
                                                _ = v49
                                            default:
                                                v45 -= (v1 ^ v45)
                                            }
                                            v45 = (v45 & v45)
                                        }
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                    switch (fi41 ^ (fi41 / ((fi41) | 1))) {
                                    case 137:
                                        var v50 float64 = (v2 / (((v2 * v2) - -2.25) * (v2 - v2)))
                                        _ = v50
                                    case 9223372036854775807:
                                        var v51 int = fi41
                                        _ = v51
                                    }
                                    fmt.Println(v2, v3, v2)
                                } else {
                                    if true {
                                        fmt.Println(v4, v1, v3)
                                    } else {
                                        fmt.Println(v1, v3, v1)
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v52 uint8 = v4
                                        _ = v52
                                    case "hello":
                                        v45 = (v1 * (-(v3)))
                                    case "":
                                        v1 += (v45 << ((65535) & 63))
                                    default:
                                        v53 := []int{fi41, fi41, 128, 9223372036854775807}
                                        _ = v53
                                    }
                                }
                                v4 += v4
                            } else {
                                v3 = (v45 & (v1 % (((v3 | v45)) | 1)))
                                fmt.Println(fi41, v4)
                                fmt.Println(v2)
                            }
                            _ = v45
                        } else {
                            if false {
                                v54 := GenS6{F0: -1}
                                var v55 float32 = (3.14159 * 3.14159)
                                _ = v54
                                _ = v55
                            } else {
                                fmt.Println(v3, v4)
                            }
                            for fi56 := 0; fi56 < 17; fi56++ {
                                v57 := GenS7{F0: 894, F1: v4}
                                v58 := []int{-1, (-(fi41)), (fi56 - (fi56 % ((-984) | 1)))}
                                switch "" {
                                case "a":
                                    switch fi56 {
                                    case 2147483647:
                                        var v59 uint16 = 37
                                        _ = v59
                                    }
                                    if len(v58) >= 2 {
                                        v58 = v58[1:len(v58)]
                                    }
                                    v3 = -128
                                    v2 += (0.0 / (v2 - v2))
                                case "":
                                    if (((v3 - -128) > (v1 << ((65535) & 63))) && (1 == 0)) {
                                        v1 += v3
                                        v1 -= v3
                                    }
                                default:
                                    if false {
                                        v3 -= (v1 & v1)
                                        v2 = (v2 * v2)
                                    } else {
                                        if (127 < v4) {
                                            if len(v58) > 0 {
                                                v58[0] = 32768
                                            }
                                            v1 += v1
                                            _ = len(v5)
                                        }
                                    }
                                }
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                            var v60 float32 = 1e10
                            _ = v60
                        }
                    }
                case "\n\t":
                    var v61 uint64 = 848
                    _ = v61
                case "a":
                    for fi62 := 0; fi62 < 3; fi62++ {
                        v63 := []int{fi62, fi62, fi62, fi41}
                        fmt.Println(fi62, fi62)
                        _ = fi62
                        _ = v63
                    }
                }
            }
        }
        _ = fi41
    }
    v64 := []int{733, int(v4)}
    switch ("" + ("hello" + "a")) {
    case "hello":
        v65 := GenS8{F0: 204, F1: v1}
        _ = v65
    case "":
        switch "\n\t" {
        case "foo bar":
            for fi66 := 0; fi66 < 5; fi66++ {
                if len(v64) > 0 {
                    v64[0] = fi66
                }
                var v67 string = ("\n\t" + "")
                _ = fi66
                _ = v67
            }
        case "\n\t":
            v3 = v3
        default:
            for fi68 := 0; fi68 < 20; fi68++ {
                var v69 float64 = v2
                _ = fi68
                _ = v69
            }
        }
    case "\n\t":
        var v70 uint64 = uint64(v4)
        _ = v70
    default:
        v71 := GenS2{F0: v4}
        _ = v71
    }
    v1 -= (v1 ^ (v1 | v3))
    fmt.Println(v3)
    if (-1 != -840) {
        switch (("foo bar" + (("a" + "\n\t") + "foo bar")) + "") {
        case "hello":
            var v72 uint8 = v4
            _ = v72
        case "":
            if len(v64) > 0 {
                v64[0] = 931
            }
        case "foo bar":
            if len(v64) >= 2 {
                v64 = v64[1:len(v64)]
            }
        }
        v73 := GenS9{F0: -142, F1: (4294967295 > 18446744073709551615), F2: 42}
        for fi74 := 0; fi74 < 9; fi74++ {
            var v75 uint32 = uint32(v4)
            var v76 string = ""
            v3 = 103
            _ = fi74
            _ = v75
            _ = v76
        }
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v64
}
