package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int16; F1 int8; F2 int8 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint; F1 uint32; F2 int16 }
func genFn0(p0 uint) int32 {
    var v5 int32 = 0
    v6 := GenS1{F0: true}
    for fi7 := 0; fi7 < 12; fi7++ {
        if (0.0 < 0.1) {
            v5 = (v5 % (((v5 ^ v5)) | 1))
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(fi8, fi8)
                p0 = (p0 >> ((455) & 63))
                _ = fi8
            }
            v9 := GenS1{F0: (fi7 <= (fi7 & (fi7 * (fi7 >> (((p0 << ((p0) & 63))) & 63)))))}
            _ = v9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            if ((448 < p0) || (true && true)) {
                if (false || ((!true) && false)) {
                    fmt.Println(v5)
                    var v11 uint32 = 94
                    _ = v11
                }
            } else {
                if (!(37 >= 1)) {
                    v12 := GenS2{F0: 0, F1: (!(-1 != 1))}
                    if true {
                        v13 := []int{fi7, fi10, (fi10 ^ (-(fi10))), fi10}
                        _ = v13
                    }
                    _ = v12
                }
                var v14 int = 141
                _ = v14
            }
            v15 := []int{42}
            if false {
                p0 = (p0 | p0)
            } else {
                switch ("" + "") {
                case "a":
                    if ((142 <= 42) || (!("" > "\n\t"))) {
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            v16 := GenS2{F0: -73, F1: false}
                            _ = v16
                        case "a":
                            v17 := GenS3{F0: "foo bar"}
                            _ = v17
                        case "":
                            if true {
                                p0 = 18446744073709551615
                                v5 -= v5
                                var v18 uint16 = 42
                                _ = v18
                            } else {
                                v19 := GenS4{F0: "foo bar", F1: 42}
                                switch "" {
                                case "hello":
                                    v20 := GenS4{F0: "hello", F1: 1}
                                    _ = v20
                                default:
                                    p0 = 1
                                }
                                _ = v19
                            }
                        }
                        var v21 uint = p0
                        _ = v21
                    }
                }
                switch ("a" + "\n\t") {
                case "a":
                    p0 -= p0
                default:
                    p0 = 18446744073709551615
                }
                var v22 uint32 = 2
                if (!((true || true) && (true && (42 > -9223372036854775808)))) {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                } else {
                    var v23 float64 = (1e10 / (3.14159 + ((0.0 / (((((((0.0 - 1e10) + (1e10 * 0.1)) / 1.5) / 1e10) / ((0.0 - 0.0) - 0.1)) - (1.5 * float64((1.5 / ((float32((1e10 * (1e10 - float64(0.1)))) / ((float32(0.1) * (0.1 / 1.5)) + 1.5)) - -2.25))))) / 3.14159)) / -2.25)))
                    _ = v23
                }
                v24 := GenS3{F0: ("hello" + "foo bar")}
                _ = v22
                _ = v24
            }
            _ = fi10
            _ = v15
        }
        fmt.Println(v5)
        _ = fi7
    }
    var v25 int32 = (v5 << (((p0 * 0)) & 63))
    var v26 int8 = 100
    _ = p0
    _ = v25
    _ = v26
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 uint16) int {
    fmt.Println(p0)
    switch "foo bar" {
    case "":
        v27 := []int{-9223372036854775808, -2147483648}
        _ = v27
    case "a":
        var v28 float32 = 1e10
        _ = v28
    case "hello":
        if (1e10 > 3.14159) {
            v29 := []int{876, 282, 1, -614}
            switch ("a" + "\n\t") {
            case "a":
                for fi30 := 0; fi30 < 20; fi30++ {
                    var v31 uint = 206
                    if (!false) {
                        v29 = append(v29, fi30)
                        fmt.Println(p0)
                        switch ("\n\t" + "foo bar") {
                        case "foo bar":
                            v32 := GenS4{F0: "a", F1: -9223372036854775808}
                            _ = v32
                        default:
                            switch "" {
                            case "a":
                                for fi33 := 0; fi33 < 2; fi33++ {
                                    var v34 int = fi33
                                    _ = fi33
                                    _ = v34
                                }
                            case "hello":
                                for fi35 := 0; fi35 < 13; fi35++ {
                                    if (!(v31 < v31)) {
                                        var v36 int = fi30
                                        var v37 string = ""
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS5{F0: -1, F1: -83, F2: 100}
                                    v31 = v31
                                    _ = fi35
                                    _ = v38
                                }
                            case "":
                                v39 := GenS6{F0: 2147483647, F1: 122}
                                _ = v39
                            default:
                                switch fi30 {
                                case 318:
                                    for fi40 := 0; fi40 < 8; fi40++ {
                                        v41 := GenS1{F0: (-1 >= -1)}
                                        _ = fi40
                                        _ = v41
                                    }
                                case -2147483648:
                                    p0 = p0
                                case 4294967295:
                                    fmt.Println(v31, p0, fi30)
                                }
                            }
                        }
                    }
                    _ = fi30
                    _ = v31
                }
            case "hello":
                v29 = append(v29, 4294967295)
            case "\n\t":
                _ = len(v29)
            }
            v42 := GenS7{F0: (!false), F1: 0, F2: (^(p0))}
            _ = v29
            _ = v42
        }
    default:
        if (!false) {
            for fi43 := 0; fi43 < 13; fi43++ {
                var v44 int64 = 42
                v45 := GenS8{F0: p0, F1: 2147483648}
                for fi46 := 0; fi46 < 18; fi46++ {
                    v47 := GenS7{F0: false, F1: 836, F2: p0}
                    _ = fi46
                    _ = v47
                }
                _ = fi43
                _ = v44
                _ = v45
            }
        }
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 uint64 = 4294967295
        _ = fi48
        _ = v49
    }
    p0 -= p0
    _ = p0
    return 4294967295
}

func main() {
    var v1 uint = 677
    var v2 float64 = (3.14159 + 3.14159)
    v3 := make(map[int]string)
    v4 := []int{}
    var v50 bool = false
    var v51 uint32 = 2147483647
    v2 += (((v2 - (0.1 / (1.5 - (1e10 + (-2.25 + v2))))) / v2) * (v2 + 3.14159))
    if (!false) {
        for fi52 := 0; fi52 < 18; fi52++ {
            for fi53 := 0; fi53 < 13; fi53++ {
                v50 = true
                var v54 uint8 = 127
                _ = fi53
                _ = v54
            }
            if v50 {
                var v55 uint8 = 127
                v50 = v50
                v55 = (v55 << (((v1 >> ((v1) & 63))) & 63))
                _ = v55
            }
            if len(v4) > 0 {
                v4[0] = 128
            }
            _ = fi52
        }
        v50 = true
        v50 = v50
    } else {
        var v56 float64 = ((((float64((((1.5 - 1.5) - 0.0) - (float32(((v2 + (v2 - float64(0.0))) + v2)) - 3.14159))) * v2) + 1e10) / v2) + float64(1e10))
        _ = v56
    }
    var v57 int = genFn1(624)
    v51 -= 42
    v58 := GenS2{F0: 0, F1: true}
    var v59 int = (v57 >> (((^(v1))) & 63))
    fmt.Println(v51)
    if v50 {
        var v60 uint64 = uint64(v51)
        v57 = v59
        _ = v60
    } else {
        v59 = v59
        if len(v4) > 0 {
            v4[0] = -1
        }
        v59 = (v59 ^ v57)
    }
    v61 := GenS9{F0: 1}
    for fi62 := 0; fi62 < 17; fi62++ {
        v57 -= (v57 + (v57 % ((genFn1(42)) | 1)))
        if true {
            v59 = v57
            if (!v50) {
                if ((!(((42 < 17) && (0 > 0)) || ((uint16(v51) <= 32767) || true))) && v50) {
                    fmt.Println(v57)
                } else {
                    if (100 <= 106) {
                        var v63 string = "a"
                        _ = v63
                    }
                    if v50 {
                        var v64 float32 = -2.25
                        _ = v64
                    }
                }
            } else {
                fmt.Println(v50, v57)
            }
        }
        if (!(!v50)) {
            var v65 int = fi62
            _ = v65
        } else {
            v4 = append(v4, (v59 * 1))
            v66 := GenS8{F0: 65535, F1: (^(v51))}
            v67 := GenS10{F0: (v1 & 65535), F1: v51, F2: -32768}
            _ = v66
            _ = v67
        }
        _ = fi62
    }
    var v68 uint8 = 2
    v68 = (v68 & 0)
    if len(v4) > 0 {
        v4[0] = (v57 / ((-127) | 1))
    }
    for fi69 := 0; fi69 < 10; fi69++ {
        v4 = append(v4, 0)
        if (uint8(v1) <= v68) {
            v57 += 219
            v68 -= v68
        } else {
            for fi70 := 0; fi70 < 1; fi70++ {
                var v71 uint16 = 65535
                _ = fi70
                _ = v71
            }
            var v72 uint = (^(v1))
            switch (v59 & genFn1(653)) {
            case 792:
                if v50 {
                    v73 := GenS5{F0: -134, F1: 0, F2: 1}
                    _ = v73
                }
            case 32767:
                v74 := GenS1{F0: v50}
                _ = v74
            case 9223372036854775807:
                if (42 >= 9223372036854775807) {
                    switch "foo bar" {
                    case "":
                        _ = len(v3)
                    case "a":
                        if v50 {
                            var v75 uint32 = v51
                            _ = v75
                        } else {
                            fmt.Println(v68)
                        }
                    case "\n\t":
                        fmt.Println(fi69, v68)
                    default:
                        v76 := GenS6{F0: 1, F1: v68}
                        _ = v76
                    }
                    switch (v57 << ((v1) & 63)) {
                    case 0:
                        delete(v3, v59)
                    default:
                        if (v50 && true) {
                            fmt.Println(v68, fi69)
                            v77 := GenS8{F0: 2, F1: uint32(v68)}
                            _ = v77
                        }
                    }
                } else {
                    fmt.Println(fi69, v68)
                    v59 = (-(v57))
                }
            }
            _ = v72
        }
        switch v59 {
        case 128:
            var v78 uint16 = uint16(v68)
            _ = v78
        case 4294967295:
            v79 := GenS9{F0: 250}
            _ = v79
        default:
            if (!false) {
                fmt.Println(v2, v68, v59)
                var v80 int = 1
                var v81 uint64 = 411
                _ = v80
                _ = v81
            }
        }
        _ = fi69
    }
    v1 += (v1 & (v1 * 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v57
    _ = v58
    _ = v59
    _ = v61
    _ = v68
}
