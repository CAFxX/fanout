package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS3 struct { F0 bool; F1 int8; F2 uint16 }
type GenS4 struct { F0 int; F1 string; F2 int32 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint; F2 int8 }
type GenS8 struct { F0 uint; F1 bool; F2 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64; F1 int8; F2 uint8 }
func genFn0() bool {
    if (((0.1 + 1.5) + 1e10) < 0.1) {
        var v4 uint32 = 2
        switch ("hello" + ("\n\t" + "\n\t")) {
        case "hello":
            var v5 int64 = 1
            _ = v5
        case "foo bar":
            v6 := GenS1{F0: -958}
            _ = v6
        }
        fmt.Println(v4)
        _ = v4
    }
    switch -1 {
    case 953:
        var v7 uint32 = 2147483647
        _ = v7
    case 2147483647:
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 float32 = 1.5
            if (true && true) {
                if true {
                    var v10 uint16 = 255
                    for fi11 := 0; fi11 < 9; fi11++ {
                        if false {
                            var v12 uint8 = 255
                            if true {
                                if true {
                                    v13 := []int{-743}
                                    v9 = v9
                                    v10 += uint16(v12)
                                    _ = v13
                                }
                                v14 := []int{(fi11 + (fi8 / (((fi11 ^ fi8)) | 1))), 1}
                                _ = v14
                            } else {
                                v15 := GenS1{F0: 0}
                                var v16 float64 = 1.5
                                v17 := GenS2{F0: (^(v10)), F1: v10, F2: -1}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            _ = v12
                        } else {
                            var v18 uint8 = 2
                            v19 := GenS3{F0: false, F1: -128, F2: v10}
                            _ = v18
                            _ = v19
                        }
                        _ = fi11
                    }
                    _ = v10
                }
            } else {
                for fi20 := 0; fi20 < 11; fi20++ {
                    if true {
                        v21 := GenS4{F0: (fi20 % ((-183) | 1)), F1: "a", F2: -2147483648}
                        v22 := []int{fi8, 127}
                        v9 = v9
                        v9 = (1.5 * v9)
                        if (42 <= -1) {
                            fmt.Println(fi8, fi20)
                        } else {
                            var v23 uint8 = 255
                            _ = v23
                        }
                        fmt.Println(fi20)
                        _ = v21
                        _ = v22
                    }
                    _ = fi20
                }
            }
            _ = fi8
            _ = v9
        }
    default:
        var v24 uint64 = 42
        _ = v24
    }
    var v25 uint32 = 2147483647
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := []int{}
        var v28 bool = (true || true)
        v28 = false
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = v25
    return false
}
func genFn1(p0 int) int16 {
    v29 := GenS3{F0: true, F1: 1, F2: 2}
    var v30 float64 = (1e10 - -2.25)
    v31 := GenS3{F0: (1 <= 258), F1: -1, F2: 65535}
    v32 := []int{p0, 2147483647, p0}
    _ = p0
    _ = v29
    _ = v30
    _ = v31
    _ = v32
    return 0
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = 0.1
    var v3 string = ("" + ("hello" + "a"))
    v33 := []int{}
    for fi34 := 0; fi34 < 9; fi34++ {
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        for fi35 := 0; fi35 < 11; fi35++ {
            switch ("hello" + "foo bar") {
            case "":
                v1 -= (v1 / (((^(v1))) | 1))
            default:
                v1 -= (^(v1))
            }
            v36 := GenS4{F0: fi34, F1: (("foo bar" + v3) + v3), F2: 2147483647}
            _ = fi35
            _ = v36
        }
        if false {
            v33 = append(v33, fi34)
            _ = len(v33)
            v1 = (v1 >> ((255) & 63))
        }
        _ = fi34
    }
    v3 = "foo bar"
    var v37 int16 = -32768
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    if len(v33) > 0 {
        v33[0] = 127
    }
    switch 32767 {
    case 128:
        v38 := GenS2{F0: 239, F1: 65535, F2: -1}
        _ = v38
    case 816:
        if len(v33) > 0 {
            v33[0] = 749
        }
    case 32767:
        for fi39 := 0; fi39 < 19; fi39++ {
            if (!((!(false || ((42 == v37) || (true || ((v1 <= (v1 - (v1 << ((255) & 63)))) && true))))) || false)) {
                var v40 uint = 873
                switch 1 {
                case -721:
                    var v41 uint16 = 32767
                    _ = v41
                case -1:
                    var v42 bool = (false || false)
                    _ = v42
                default:
                    v37 = (v37 >> ((v40) & 63))
                }
                v33 = append(v33, 42)
                _ = v40
            } else {
                v43 := GenS1{F0: 561}
                for fi44 := 0; fi44 < 4; fi44++ {
                    if len(v33) > 0 {
                        v33[0] = (fi44 >> ((160) & 63))
                    }
                    _ = fi44
                }
                for fi45 := 0; fi45 < 14; fi45++ {
                    fmt.Println(v37, fi45, v37)
                    _ = fi45
                }
                _ = v43
            }
            if len(v33) > 0 {
                v33[0] = fi39
            }
            switch v3 {
            case "":
                fmt.Println(v3, v37, fi39)
            case "foo bar":
                v2 += 1.5
            case "a":
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
            }
            _ = fi39
        }
    }
    if false {
        v46 := GenS3{F0: ((71 == 42) && (!true)), F1: -128, F2: 32767}
        _ = v46
    } else {
        v37 += (v37 * v37)
        v2 = float64(1.5)
        var v47 uint8 = 2
        _ = v47
    }
    v48 := GenS5{F0: v1, F1: true}
    var v49 uint32 = 2147483647
    fmt.Println(v2, v49, v49)
    fmt.Println(v49, v49)
    if false {
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        v37 = 42
        if true {
            v2 += float64(float32(float64(3.14159)))
            for fi50 := 0; fi50 < 11; fi50++ {
                for fi51 := 0; fi51 < 6; fi51++ {
                    fmt.Println(fi50)
                    _ = fi51
                }
                if (451 <= 156) {
                    fmt.Println(v2, v37)
                    v1 -= 415
                    v52 := GenS6{F0: uint8(v1), F1: 2, F2: 1}
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v2 -= float64((float32(((v2 - v2) / 1.0)) * 0.0))
                        v1 += 0
                        v54 := GenS7{F0: 2147483647, F1: 255, F2: 42}
                        _ = fi53
                        _ = v54
                    }
                }
                _ = fi50
            }
            v1 -= (v1 - v1)
        }
    } else {
        var v55 int = 9223372036854775807
        var v56 int64 = 0
        v3 = v3
        _ = v55
        _ = v56
    }
    switch v3 {
    case "a":
        switch v3 {
        case "hello":
            fmt.Println(v2, v2)
        default:
            if (478 == 9223372036854775807) {
                var v57 bool = ((255 >= 128) && (42 != -2147483648))
                for fi58 := 0; fi58 < 9; fi58++ {
                    switch (fi58 - (fi58 % ((0) | 1))) {
                    case -333:
                        var v59 uint64 = v1
                        _ = v59
                    case 42:
                        v60 := GenS8{F0: 1, F1: v57, F2: v57}
                        _ = v60
                    default:
                        if len(v33) > 0 {
                            v33[0] = 2147483647
                        }
                    }
                    _ = fi58
                }
                _ = v57
            }
        }
    case "\n\t":
        for fi61 := 0; fi61 < 1; fi61++ {
            fmt.Println(v49)
            fmt.Println(v2, v3)
            switch fi61 {
            case 128:
                for fi62 := 0; fi62 < 8; fi62++ {
                    var v63 int64 = 42
                    _ = fi62
                    _ = v63
                }
            case -117:
                var v64 int8 = 0
                _ = v64
            case -2147483648:
                if ("foo bar" != v3) {
                    v65 := GenS7{F0: -1, F1: 358, F2: 100}
                    _ = v65
                } else {
                    switch (("foo bar" + (v3 + v3)) + v3) {
                    case "foo bar":
                        if false {
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            if (!(v37 < (v37 / ((v37) | 1)))) {
                                for fi66 := 0; fi66 < 20; fi66++ {
                                    v67 := GenS9{F0: (134 > 159)}
                                    _ = fi66
                                    _ = v67
                                }
                            }
                        } else {
                            var v68 string = (v3 + "\n\t")
                            _ = v68
                        }
                    case "\n\t":
                        if (((v3 + v3) != v3) || (!false)) {
                            fmt.Println(v37, v49)
                            v69 := GenS10{F0: -571}
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            _ = v69
                        } else {
                            v49 = 981
                        }
                    case "hello":
                        var v70 int16 = -1
                        _ = v70
                    }
                    fmt.Println(v49, v37, v37)
                    if len(v33) > 0 {
                        v33[0] = 2147483647
                    }
                }
            default:
                switch (fi61 ^ 127) {
                case -173:
                    fmt.Println(v49, v37, v1)
                case 32768:
                    var v71 uint = 0
                    _ = v71
                default:
                    v72 := GenS11{F0: 0, F1: -128, F2: 128}
                    _ = v72
                }
            }
            _ = fi61
        }
    case "hello":
        var v73 uint16 = 1
        _ = v73
    default:
        _ = len(v33)
    }
    v49 = (v49 - 1)
    var v74 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v37:", v37)
    fmt.Println("v49:", v49)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v48
    _ = v49
    _ = v74
}
