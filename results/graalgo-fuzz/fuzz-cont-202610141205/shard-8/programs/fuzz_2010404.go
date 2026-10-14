package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint; F1 int64; F2 uint16 }
type GenS4 struct { F0 int; F1 int8; F2 int16 }
type GenS5 struct { F0 int8; F1 uint64; F2 int64 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32) int8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            switch (fi7 - 4294967295) {
            case 204:
                v9 := GenS1{F0: -32768, F1: 0, F2: false}
                _ = v9
            default:
                if true {
                    fmt.Println(p0, p0)
                }
                var v10 int16 = 831
                _ = v10
            }
            var v11 bool = (255 != 0)
            v12 := []int{fi7, (fi7 % ((fi7) | 1))}
            _ = fi8
            _ = v11
            _ = v12
        }
        fmt.Println(fi7)
        _ = fi7
    }
    switch 0 {
    case 422:
        switch 402 {
        case -9223372036854775808:
            v13 := []int{-570, 127, -760}
            _ = v13
        case 4294967295:
            var v14 int = -2147483648
            _ = v14
        default:
            v15 := GenS1{F0: -32768, F1: 100, F2: (32767 == 65535)}
            _ = v15
        }
    }
    var v16 int = 2147483647
    v17 := []int{-846, (v16 ^ v16), (v16 << ((341) & 63)), (^(v16))}
    _ = p0
    _ = v16
    _ = v17
    return int8(p0)
}
func genFn1(p0 uint32) bool {
    var v18 float32 = 0.1
    var v19 float64 = 0.0
    _ = p0
    _ = v18
    _ = v19
    return false
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]uint32)
    var v3 uint8 = 255
    v4 := []uint32{2, (v1 + 0), (v1 % ((v1) | 1)), (v1 + 2)}
    var v5 float32 = 3.14159
    var v6 int8 = 0
    fmt.Println(v6)
    if (true || true) {
        var v20 uint = 0
        _ = v20
    }
    var v21 int16 = 1
    v22 := v2[-2147483648]
    v5 = float32(3.14159)
    v5 = 0.1
    v22 = 42
    switch 2147483647 {
    case 488:
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 bool = (!(127 <= (v3 * 38)))
            _ = fi23
            _ = v24
        }
    case -9223372036854775808:
        v5 = (v5 * 3.14159)
    case 4294967295:
        for fi25 := 0; fi25 < 18; fi25++ {
            v21 = (v21 % ((v21) | 1))
            _ = fi25
        }
    default:
        v26 := GenS2{F0: 122, F1: "hello"}
        _ = v26
    }
    fmt.Println(v1, v21, v6)
    switch "foo bar" {
    case "foo bar":
        v4 = append(v4, 4294967295)
    case "\n\t":
        v22 = (v22 >> ((42) & 63))
    }
    switch "a" {
    case "":
        v27 := GenS3{F0: 377, F1: -473, F2: 2}
        _ = v27
    case "hello":
        v5 -= 1.5
    case "\n\t":
        v28 := GenS2{F0: -1, F1: (("a" + "a") + "a")}
        _ = v28
    default:
        v3 += (v3 + v3)
    }
    v5 = float32(1.5)
    if true {
        delete(v2, 860)
        v1 -= v1
        v1 = v22
    } else {
        v3 += (v3 + (v3 | (v3 ^ (v3 + (v3 & v3)))))
    }
    fmt.Println(v3, v5, v3)
    for fi29 := 0; fi29 < 15; fi29++ {
        _ = len(v4)
        _ = fi29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        if true {
            fmt.Println(v3, v5)
            fmt.Println(v1, v3, fi30)
            v31 := v2[fi30]
            _ = v31
        } else {
            switch (fi30 % ((int(v3)) | 1)) {
            case 32768:
                var v32 int16 = 0
                _ = v32
            case 127:
                v1 += v22
            case 4294967295:
                var v33 int = fi30
                _ = v33
            default:
                v34 := GenS4{F0: fi30, F1: 42, F2: v21}
                _ = v34
            }
            if false {
                for fi35 := 0; fi35 < 16; fi35++ {
                    var v36 int = fi30
                    v6 -= (v6 / ((1) | 1))
                    _ = fi35
                    _ = v36
                }
            } else {
                fmt.Println(v1, v22, v5)
                v37 := GenS1{F0: 4, F1: v6, F2: false}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v37
            }
            v22 = 65
        }
        _ = fi30
    }
    if true {
        for fi38 := 0; fi38 < 17; fi38++ {
            v22 += (v1 | v22)
            var v39 float32 = v5
            var v40 uint64 = 18446744073709551615
            _ = fi38
            _ = v39
            _ = v40
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            if true {
                if ((!((514 != 18446744073709551615) || (!false))) || (true && true)) {
                    var v42 int16 = (v21 | v21)
                    switch fi41 {
                    case -1:
                        var v43 int64 = 0
                        _ = v43
                    case -958:
                        if (2 != 509) {
                            for fi44 := 0; fi44 < 12; fi44++ {
                                switch (fi44 + 2147483647) {
                                case 32767:
                                    v45 := GenS3{F0: uint(v1), F1: -889, F2: 65535}
                                    _ = v45
                                case -9223372036854775808:
                                    v46 := GenS5{F0: (^(v6)), F1: 2, F2: 0}
                                    _ = v46
                                case 42:
                                    v21 -= (v21 ^ v42)
                                default:
                                    if true {
                                        v5 = ((float32((1.5 - (float64(float32((((float64(v5) * 1.5) * 0.0) / 1.5))) * -2.25))) / v5) + 0.0)
                                    }
                                }
                                var v47 uint32 = (^(v22))
                                _ = fi44
                                _ = v47
                            }
                            var v48 int64 = 1
                            v49 := GenS1{F0: int16(v3), F1: v6, F2: (1 > 42)}
                            _ = v48
                            _ = v49
                        } else {
                            v50 := GenS4{F0: (^(fi41)), F1: int8(v22), F2: (v42 / ((v21) | 1))}
                            _ = v50
                        }
                    default:
                        var v51 uint = 359
                        _ = v51
                    }
                    _ = len(v4)
                    _ = v42
                } else {
                    for fi52 := 0; fi52 < 17; fi52++ {
                        if ((v21 + (v21 - (v21 << ((uint(v22)) & 63)))) == (v21 / (((v21 & (v21 % ((-1) | 1)))) | 1))) {
                            if false {
                                var v53 float32 = ((v5 + v5) - (v5 - v5))
                                _ = v53
                            }
                        } else {
                            var v54 int = -9223372036854775808
                            var v55 int32 = -715
                            if (273 != 694) {
                                v4 = append(v4, 4294967295)
                                if true {
                                    v5 -= v5
                                    if true {
                                        _ = len(v2)
                                    } else {
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            var v56 int64 = -1
                                            _ = v56
                                        }
                                        v6 = v6
                                        v21 = (-(v21))
                                    }
                                    var v57 int32 = (v55 ^ (v55 - (^(v55))))
                                    _ = v57
                                } else {
                                    if (!true) {
                                        switch 0 {
                                        case 42:
                                            if (!true) {
                                                v5 += v5
                                                fmt.Println(v6, fi52)
                                            } else {
                                                if false {
                                                    v5 -= -2.25
                                                    v5 += v5
                                                }
                                                var v58 int16 = (v21 | v21)
                                                _ = v58
                                            }
                                        case 1:
                                            fmt.Println(v22, v22, v54)
                                        case 32768:
                                            var v59 int64 = -9223372036854775808
                                            _ = v59
                                        }
                                        v21 -= (v21 % ((v21) | 1))
                                    } else {
                                        v3 = (v3 - v3)
                                        v60 := GenS3{F0: 65535, F1: 756, F2: 1}
                                        _ = v60
                                    }
                                    var v61 uint16 = 255
                                    _ = v61
                                }
                                var v62 uint64 = 1
                                _ = v62
                            } else {
                                v63 := GenS1{F0: v21, F1: v6, F2: (!false)}
                                v1 += v22
                                _ = v63
                            }
                            v4 = append(v4, (^(v22)))
                            _ = v54
                            _ = v55
                        }
                        var v64 uint8 = (v3 % ((uint8(v22)) | 1))
                        _ = fi52
                        _ = v64
                    }
                }
                var v65 int32 = -2147483648
                switch "a" {
                case "":
                    var v66 int64 = 1
                    _ = v66
                default:
                    for fi67 := 0; fi67 < 15; fi67++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v68 int = (fi67 << ((42) & 63))
                        switch (fi41 | (v68 >> ((850) & 63))) {
                        case 32768:
                            var v69 uint8 = 146
                            _ = v69
                        default:
                            switch ("foo bar" + "hello") {
                            case "\n\t":
                                if ((v3 >> ((uint(v3)) & 63)) == 127) {
                                    v70 := GenS6{F0: (^(v6))}
                                    if (float64(v5) == ((float64(v5) * ((3.14159 - ((((((((0.0 - ((3.14159 + (float64(v5) / (-2.25 * 1.5))) + 0.1)) + 1.5) - 3.14159) * ((1e10 * (1e10 + -2.25)) + 3.14159)) - (1e10 + 3.14159)) - (0.1 + float64(float32(1.5)))) / 1e10) - 1.5)) / 0.1)) - (1e10 / 3.14159))) {
                                        v71 := GenS6{F0: v6}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v21 = (v21 / ((v21) | 1))
                                        _ = v71
                                    }
                                    _ = v70
                                }
                            case "a":
                                var v72 int16 = (v21 >> ((0) & 63))
                                _ = v72
                            case "hello":
                                var v73 float32 = (v5 / (v5 - v5))
                                _ = v73
                            }
                        }
                        if (2 != 1) {
                            var v74 uint8 = v3
                            _ = v74
                        } else {
                            v6 += (v6 | genFn0(42))
                        }
                        v5 = (v5 + v5)
                        _ = fi67
                        _ = v68
                    }
                }
                _ = v65
            }
            _ = fi41
        }
    } else {
        _ = len(v2)
        v5 = 0.0
        delete(v2, 4294967295)
    }
    v22 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
