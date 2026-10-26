package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 uint; F1 int16 }
type GenS10 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0(p0 uint8, p1 int) bool {
    var v4 int8 = -128
    for fi5 := 0; fi5 < 14; fi5++ {
        if (!(255 != 42)) {
            for fi6 := 0; fi6 < 18; fi6++ {
                fmt.Println(p1)
                v7 := []int{fi6}
                var v8 uint32 = 925
                var v9 uint64 = 1
                v9 += v9
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 1, F1: 0}
            for fi11 := 0; fi11 < 11; fi11++ {
                p0 -= p0
                v12 := []int{1}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        var v13 bool = true
        _ = fi5
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v4
    return true
}
func genFn1(p0 uint16, p1 int) uint {
    p1 -= -9223372036854775808
    switch (p1 + 4294967295) {
    case -2147483648:
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(p0)
            _ = fi14
        }
    default:
        fmt.Println(p0)
    }
    switch "foo bar" {
    case "a":
        fmt.Println(p0)
    default:
        p0 = 1
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint32 = 2147483648
        fmt.Println(fi15, p0, fi15)
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    switch -343 {
    case 4294967295:
        v17 := GenS2{F0: (4294967295 < 523)}
        _ = v17
    case 2147483647:
        var v18 int8 = 127
        _ = v18
    case 32767:
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := GenS1{F0: 697, F1: 65535}
            var v21 uint32 = 998
            _ = fi19
            _ = v20
            _ = v21
        }
    default:
        var v22 uint8 = 255
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 float64 = (0.1 - 1e10)
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS3{F0: 0, F1: true}
            if true {
                switch fi25 {
                case 127:
                    if (((false && false) || true) && (!true)) {
                        v27 := []int{}
                        v24 = 3.14159
                        v28 := GenS4{F0: 982, F1: 0, F2: -784}
                        switch fi23 {
                        case 9223372036854775807:
                            switch "" {
                            case "foo bar":
                                v24 += v24
                            case "hello":
                                fmt.Println(fi25, fi23, fi23)
                            default:
                                if true {
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    var v29 int16 = 0
                                    _ = v29
                                }
                            }
                            v30 := GenS5{F0: 47, F1: -1}
                            _ = v30
                        case 128:
                            fmt.Println(v24)
                        }
                        var v31 float32 = (3.14159 + 1.5)
                        _ = len(v27)
                        _ = v27
                        _ = v28
                        _ = v31
                    }
                case 2147483647:
                    var v32 uint16 = 255
                    _ = v32
                }
                v24 -= v24
                if (true && (1.5 >= (1.5 * 3.14159))) {
                    v24 = ((((-2.25 - 0.0) * v24) - 3.14159) - v24)
                    var v33 uint32 = 4294967295
                    if (1000 == 42) {
                        v34 := GenS6{F0: 0, F1: genFn1(32767, (fi23 + (fi25 % ((fi23) | 1))))}
                        fmt.Println(v33, v24)
                        _ = v34
                    } else {
                        switch (-(fi25)) {
                        case -859:
                            v35 := []int{fi23}
                            _ = v35
                        case 0:
                            var v36 int16 = 193
                            _ = v36
                        case -712:
                            v24 = v24
                        }
                    }
                    _ = v33
                } else {
                    v24 -= v24
                    v3[-53] = (-1 > 1)
                }
                if false {
                    fmt.Println(fi25, fi25)
                } else {
                    var v37 uint32 = 610
                    var v38 bool = false
                    v39 := GenS7{F0: (^(fi23)), F1: 223}
                    _ = v37
                    _ = v38
                    _ = v39
                }
                fmt.Println(fi23, fi23, fi25)
                v24 = ((v24 - float64(((1e10 / 1.0) + 3.14159))) + ((1.5 * (v24 / 1.0)) + 0.0))
                var v40 bool = false
                _ = v40
            }
            fmt.Println(fi23, fi25)
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    v41 := GenS2{F0: (-171 > -1)}
    var v42 int16 = 0
    fmt.Println(v42)
    for fi43 := 0; fi43 < 5; fi43++ {
        v42 = v42
        if (-1 > -2147483648) {
            for fi44 := 0; fi44 < 16; fi44++ {
                v45 := GenS1{F0: 318, F1: 42}
                if true {
                    delete(v1, fi44)
                } else {
                    v46 := GenS2{F0: (392 > -1)}
                    _ = v46
                }
                _ = fi44
                _ = v45
            }
            var v47 uint64 = 1
            v42 = 42
            _ = v47
        } else {
            v48 := []int{(fi43 + -756), fi43}
            _ = v48
        }
        v49 := []int{fi43, fi43, (fi43 / ((fi43) | 1))}
        _ = fi43
        _ = v49
    }
    switch 2147483647 {
    case 211:
        v42 = (v42 * -1)
    case -1:
        v50 := v1[-968]
        _ = v50
    }
    if true {
        switch ("\n\t" + "") {
        case "hello":
            var v51 uint16 = 255
            _ = v51
        default:
            switch 127 {
            case -853:
                if (0 != 967) {
                    fmt.Println(v42)
                    switch (("\n\t" + (("a" + "foo bar") + "foo bar")) + "a") {
                    case "":
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            var v52 uint = 42
                            _ = v52
                        }
                    default:
                        var v53 float64 = 0.1
                        _ = v53
                    }
                    if (255 >= 255) {
                        v54 := GenS6{F0: -1, F1: 42}
                        _ = v54
                    }
                } else {
                    switch ("foo bar" + "") {
                    case "a":
                        for fi55 := 0; fi55 < 3; fi55++ {
                            for fi56 := 0; fi56 < 16; fi56++ {
                                v42 -= (v42 - (v42 ^ (v42 & (v42 * v42))))
                                var v57 uint = 65535
                                if true {
                                    v42 -= (v42 & (v42 % (((v42 ^ int16(fi55))) | 1)))
                                    v58 := GenS8{F0: -116, F1: 42}
                                    v59 := GenS9{F0: v57, F1: (v42 - (v42 / ((-1) | 1)))}
                                    _ = v58
                                    _ = v59
                                } else {
                                    fmt.Println(v57, fi56, fi56)
                                    delete(v2, ((("hello" + "foo bar") + "foo bar") + ("" + ("a" + ""))))
                                }
                                _ = fi56
                                _ = v57
                            }
                            for fi60 := 0; fi60 < 1; fi60++ {
                                if ((166 != 558) && (255 < 405)) {
                                    v42 = v42
                                    v61 := []int{}
                                    _ = v61
                                }
                                var v62 int8 = 0
                                if (false || (true && (!(1 != 1)))) {
                                    v63 := []int{-332, (-(fi55)), 2147483647}
                                    v64 := GenS6{F0: -128, F1: genFn1(32767, (^(fi55)))}
                                    _ = v63
                                    _ = v64
                                }
                                _ = fi60
                                _ = v62
                            }
                            _ = fi55
                        }
                    }
                }
            default:
                v65 := GenS5{F0: 85, F1: (v42 | 1000)}
                _ = v65
            }
        }
    }
    v66 := GenS5{F0: 42, F1: 1}
    v42 = (v42 | v42)
    var v67 int = 4294967295
    for fi68 := 0; fi68 < 5; fi68++ {
        if true {
            fmt.Println(v42)
        }
        var v69 bool = false
        _ = fi68
        _ = v69
    }
    for fi70 := 0; fi70 < 19; fi70++ {
        var v71 uint64 = 963
        v42 = (v42 & (v42 * (v42 >> ((18446744073709551615) & 63))))
        if true {
            for fi72 := 0; fi72 < 2; fi72++ {
                switch (v67 << ((4294967295) & 63)) {
                case -935:
                    var v73 uint32 = 65535
                    _ = v73
                case -9223372036854775808:
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v71, v42)
                    default:
                        v74 := []int{fi70, (v67 ^ fi70), -850, (fi70 << ((255) & 63))}
                        var v75 int16 = v42
                        _ = v74
                        _ = v75
                    }
                    var v76 float64 = -2.25
                    _ = v76
                case 0:
                    v77 := []int{(v67 >> ((4294967295) & 63))}
                    fmt.Println(fi72, v67)
                    if (-1 <= -128) {
                        var v78 bool = (!(-824 >= 42))
                        v67 = -754
                        if len(v77) >= 2 {
                            v77 = v77[1:len(v77)]
                        }
                        _ = v78
                    }
                    _ = v77
                default:
                    v79 := []int{(fi72 % ((int(v42)) | 1)), fi70, v67, fi72}
                    _ = v79
                }
                v80 := v2["foo bar"]
                _ = fi72
                _ = v80
            }
            v81 := GenS10{F0: 1, F1: 4294967295, F2: uint8(v71)}
            _ = v81
        }
        _ = fi70
        _ = v71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v66
    _ = v67
}
