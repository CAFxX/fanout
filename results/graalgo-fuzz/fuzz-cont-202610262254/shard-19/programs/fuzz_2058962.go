package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint64; F2 int64 }
type GenS3 struct { F0 string; F1 uint; F2 int }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int16; F2 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 int64; F2 int8 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := []int{32767, 32768}
        for fi6 := 0; fi6 < 8; fi6++ {
            v5 = append(v5, fi4)
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    var v7 int32 = 601
    _ = v7
    return v7
}

func main() {
    var v1 int16 = 32767
    var v2 float32 = 3.14159
    var v3 float64 = 0.1
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint64 = 42
        v10 := []int{fi8, fi8}
        switch "a" {
        case "":
            var v11 uint = 481
            _ = v11
        case "foo bar":
            fmt.Println(v1, v3)
        case "hello":
            var v12 string = "foo bar"
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 -= (v3 / (((0.0 + v3) + v3) + (v3 + float64(v2))))
    var v13 uint32 = 2147483648
    v14 := []int{-490, 128}
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v13)
        _ = len(v14)
        _ = fi15
    }
    if false {
        if (229 >= 1) {
            if false {
                var v16 float64 = v3
                _ = v16
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 int32 = -1
                    _ = fi17
                    _ = v18
                }
                if len(v14) > 0 {
                    v14[0] = 42
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    if len(v14) > 0 {
                        v14[0] = fi19
                    }
                    if (-22 == 42) {
                        v3 += v3
                    } else {
                        v20 := GenS1{F0: (v13 >> ((18446744073709551615) & 63))}
                        if (42 > 1) {
                            var v21 float64 = v3
                            v22 := GenS1{F0: (v13 + (v13 >> ((65535) & 63)))}
                            v1 = (v1 | -1)
                            _ = v21
                            _ = v22
                        } else {
                            if false {
                                v1 -= (v1 ^ (v1 * (^(v1))))
                            } else {
                                _ = len(v14)
                                if false {
                                    for fi23 := 0; fi23 < 19; fi23++ {
                                        var v24 string = "\n\t"
                                        fmt.Println(v24, v2)
                                        var v25 uint64 = 2
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                    }
                                }
                            }
                        }
                        if false {
                            v13 = (v13 * 2)
                            v26 := GenS1{F0: v13}
                            v27 := GenS2{F0: true, F1: 18446744073709551615, F2: -9223372036854775808}
                            _ = v26
                            _ = v27
                        } else {
                            if true {
                                var v28 int16 = v1
                                v28 = v28
                                var v29 int8 = 0
                                _ = v28
                                _ = v29
                            }
                            fmt.Println(fi19, v3)
                            if (false && ((!false) && (!(684 <= 42)))) {
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    if false {
                                        fmt.Println(v3, fi30, fi19)
                                    } else {
                                        v13 += v13
                                        if (!true) {
                                            v3 -= 0.1
                                            if false {
                                                var v31 bool = true
                                                fmt.Println(v3)
                                                _ = v31
                                            }
                                            var v32 int32 = 0
                                            _ = v32
                                        }
                                    }
                                    switch fi30 {
                                    case 0:
                                        if (true && (float32(float64((v2 - (v2 - ((float32(float64(float32(0.1))) - 1.5) * 1e10))))) > v2)) {
                                            v14 = append(v14, 16)
                                            v2 -= v2
                                        }
                                    case 127:
                                        v13 = 77
                                        v33 := GenS1{F0: v13}
                                        _ = v33
                                    case 539:
                                        var v34 uint32 = v13
                                        _ = v34
                                    default:
                                        v35 := GenS3{F0: "\n\t", F1: 842, F2: fi30}
                                        _ = v35
                                    }
                                    _ = fi30
                                }
                                fmt.Println(v1, v2)
                            } else {
                                if false {
                                    if (true && (true && ((127 < 0) && (2 == 18446744073709551615)))) {
                                        v36 := GenS4{F0: 447, F1: (v13 - (^(v13)))}
                                        if len(v14) > 0 {
                                            v14[0] = -1
                                        }
                                        _ = v36
                                    } else {
                                        v2 -= 1e10
                                        var v37 uint8 = 0
                                        switch "foo bar" {
                                        case "hello":
                                            switch int(v1) {
                                            case 9223372036854775807:
                                                v38 := GenS2{F0: true, F1: 18446744073709551615, F2: -234}
                                                _ = v38
                                            case 1:
                                                v39 := GenS2{F0: false, F1: 736, F2: -9223372036854775808}
                                                _ = v39
                                            case -79:
                                                if (true || false) {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if len(v14) > 0 {
                                                            v14[0] = (fi19 * fi19)
                                                        }
                                                    case "":
                                                        switch "hello" {
                                                        case "a":
                                                            switch ("hello" + "a") {
                                                            case "foo bar":
                                                                v37 = v37
                                                            case "hello":
                                                                var v40 uint = 1
                                                                _ = v40
                                                            default:
                                                                for fi41 := 0; fi41 < 7; fi41++ {
                                                                    var v42 int16 = (v1 * v1)
                                                                    v43 := GenS5{F0: 568, F1: v1, F2: (("foo bar" + "\n\t") + (("" + "foo bar") + (("\n\t" + ("hello" + "\n\t")) + "foo bar")))}
                                                                    switch (("\n\t" + ("\n\t" + "foo bar")) + (((("\n\t" + "\n\t") + "foo bar") + "") + "")) {
                                                                    case "a":
                                                                        var v44 int16 = (v42 & (-(v1)))
                                                                        switch "a" {
                                                                        case "a":
                                                                            switch -1 {
                                                                            case 127:
                                                                                fmt.Println(v44, v44)
                                                                            case 2147483647:
                                                                                v44 -= -529
                                                                            case 181:
                                                                                v1 += (v42 % ((32767) | 1))
                                                                            default:
                                                                                v45 := GenS2{F0: false, F1: 2, F2: 42}
                                                                                _ = v45
                                                                            }
                                                                        }
                                                                        if len(v14) >= 2 {
                                                                            v14 = v14[1:len(v14)]
                                                                        }
                                                                        _ = v44
                                                                    case "\n\t":
                                                                        switch -989 {
                                                                        case 128:
                                                                            if ((fi19 == (fi41 / ((fi19) | 1))) && ((v13 ^ v13) != 901)) {
                                                                                if true {
                                                                                    v46 := GenS4{F0: v13, F1: v13}
                                                                                    v3 += 0.0
                                                                                    switch fi19 {
                                                                                    case -2147483648:
                                                                                        if (!(664 <= 4294967295)) {
                                                                                            v2 = v2
                                                                                        } else {
                                                                                            var v47 uint8 = 182
                                                                                            v37 = v37
                                                                                            v48 := GenS2{F0: (65535 != 373), F1: 113, F2: 9223372036854775807}
                                                                                            _ = v47
                                                                                            _ = v48
                                                                                        }
                                                                                    default:
                                                                                        var v49 float64 = 1e10
                                                                                        _ = v49
                                                                                    }
                                                                                    _ = v46
                                                                                } else {
                                                                                    var v50 int32 = 917
                                                                                    v1 = 0
                                                                                    v51 := GenS2{F0: false, F1: 18446744073709551615, F2: -9223372036854775808}
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                }
                                                                                v1 = -1
                                                                            }
                                                                        default:
                                                                            var v52 bool = true
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                    v3 += v3
                                                                    _ = fi41
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                            }
                                                        case "\n\t":
                                                            switch "a" {
                                                            case "a":
                                                                var v53 bool = true
                                                                _ = v53
                                                            }
                                                        case "hello":
                                                            var v54 int64 = 1
                                                            _ = v54
                                                        default:
                                                            v55 := GenS3{F0: "", F1: 65535, F2: fi19}
                                                            _ = v55
                                                        }
                                                    case "hello":
                                                        fmt.Println(v13, v13, fi19)
                                                    default:
                                                        if (true && ((!(v37 >= v37)) || (42 < 0))) {
                                                            for fi56 := 0; fi56 < 5; fi56++ {
                                                                if (false || true) {
                                                                    var v57 float32 = (float32(-2.25) + (0.0 - v2))
                                                                    var v58 int8 = 1
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    var v59 string = "foo bar"
                                                                    switch fi56 {
                                                                    case 1:
                                                                        v60 := GenS5{F0: uint16(v37), F1: -32768, F2: (v59 + "")}
                                                                        _ = v60
                                                                    case 9223372036854775807:
                                                                        if len(v14) >= 2 {
                                                                            v14 = v14[1:len(v14)]
                                                                        }
                                                                    }
                                                                    _ = v59
                                                                }
                                                                fmt.Println(fi19, v13)
                                                                _ = fi56
                                                            }
                                                            v3 = (v3 / 0.1)
                                                        }
                                                    }
                                                }
                                            default:
                                                v61 := GenS6{F0: 1}
                                                _ = v61
                                            }
                                        case "a":
                                            v13 += (v13 << ((42) & 63))
                                        }
                                        _ = v37
                                    }
                                }
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                v2 = v2
                            }
                        }
                        _ = v20
                    }
                    if len(v14) > 0 {
                        v14[0] = (fi19 ^ 127)
                    }
                    _ = fi19
                }
            }
        } else {
            fmt.Println(v1, v3, v13)
        }
    }
    v62 := GenS2{F0: false, F1: 77, F2: 0}
    if false {
        var v63 int8 = -1
        _ = v63
    } else {
        v64 := GenS1{F0: v13}
        var v65 int16 = (v1 >> ((42) & 63))
        _ = v64
        _ = v65
    }
    var v66 uint = 255
    v67 := GenS7{F0: -2147483648, F1: 9223372036854775807, F2: 42}
    var v68 uint32 = 645
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v62
    _ = v66
    _ = v67
    _ = v68
}
