package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int8; F1 int64; F2 int16 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint8; F1 string; F2 uint }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint64; F2 string }
type GenS11 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 0
    var v3 float32 = (0.1 + 3.14159)
    _ = len(v1)
    if false {
        if (-128 > -128) {
            v4 := GenS1{F0: 1, F1: 1}
            v5 := []int{128, 4294967295, 42, -450}
            _ = v4
            _ = v5
        } else {
            if true {
                if (true || true) {
                    switch (("foo bar" + ("hello" + "\n\t")) + ("foo bar" + (("hello" + ("a" + "foo bar")) + ""))) {
                    case "":
                        var v6 float64 = (3.14159 * 1.5)
                        _ = v6
                    case "a":
                        var v7 bool = false
                        _ = v7
                    case "hello":
                        if (42 >= 42) {
                            if ((v2 + 2147483647) == v2) {
                                var v8 int32 = 1
                                for fi9 := 0; fi9 < 14; fi9++ {
                                    for fi10 := 0; fi10 < 5; fi10++ {
                                        if true {
                                            var v11 uint64 = 4294967295
                                            if (float64((v3 - 1e10)) > 1e10) {
                                                delete(v1, 587)
                                                v3 -= float32(0.1)
                                                v12 := v1[int(v2)]
                                                _ = v12
                                            } else {
                                                fmt.Println(fi9, fi10)
                                                var v13 int = fi10
                                                _ = v13
                                            }
                                            delete(v1, -503)
                                            _ = v11
                                        } else {
                                            var v14 float32 = v3
                                            v15 := GenS1{F0: 0, F1: 107}
                                            fmt.Println(v3, fi10, fi9)
                                            _ = v14
                                            _ = v15
                                        }
                                        v16 := GenS1{F0: 951, F1: uint8(v2)}
                                        _ = fi10
                                        _ = v16
                                    }
                                    _ = fi9
                                }
                                _ = v8
                            }
                            for fi17 := 0; fi17 < 14; fi17++ {
                                if true {
                                    fmt.Println(v2)
                                } else {
                                    for fi18 := 0; fi18 < 7; fi18++ {
                                        switch ("a" + "a") {
                                        case "hello":
                                            switch "\n\t" {
                                            case "hello":
                                                switch "a" {
                                                case "\n\t":
                                                    if ((3.14159 + v3) > v3) {
                                                        fmt.Println(v3)
                                                    }
                                                case "hello":
                                                    var v19 uint64 = 1
                                                    _ = v19
                                                }
                                            case "foo bar":
                                                fmt.Println(v3)
                                            default:
                                                if true {
                                                    if ((true || (231 <= uint8(v2))) && (899 <= 42)) {
                                                        v3 -= ((0.0 - v3) / 1.5)
                                                        v20 := []int{}
                                                        if len(v20) > 0 {
                                                            v20[0] = (fi17 + 32768)
                                                        }
                                                        if len(v20) >= 2 {
                                                            v20 = v20[1:len(v20)]
                                                        }
                                                        v2 = (v2 / ((v2) | 1))
                                                        switch fi17 {
                                                        case 32768:
                                                            fmt.Println(v3)
                                                            v2 = 1
                                                            var v21 float64 = 3.14159
                                                            v22 := GenS1{F0: 39, F1: 0}
                                                            v21 += 1e10
                                                            v21 = v21
                                                            v23 := GenS2{F0: 2147483647, F1: -1, F2: 0}
                                                            _ = v21
                                                            _ = v22
                                                            _ = v23
                                                        case -351:
                                                            var v24 int8 = 1
                                                            _ = v24
                                                        case -96:
                                                            if false {
                                                                v25 := GenS2{F0: (fi17 % (((fi18 & 32767)) | 1)), F1: -1, F2: 499}
                                                                if false {
                                                                    var v26 uint8 = 205
                                                                    _ = v26
                                                                }
                                                                fmt.Println(v2, v2)
                                                                _ = v25
                                                            }
                                                        default:
                                                            var v27 uint = 0
                                                            _ = v27
                                                        }
                                                        _ = v20
                                                    }
                                                }
                                            }
                                        case "a":
                                            v2 = (v2 % ((680) | 1))
                                        }
                                        _ = fi18
                                    }
                                    v2 += v2
                                }
                                switch "foo bar" {
                                case "":
                                    v28 := GenS3{F0: 100}
                                    _ = v28
                                case "foo bar":
                                    var v29 int16 = 1
                                    _ = v29
                                }
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    v2 = (^(v2))
                                    _ = fi30
                                }
                                _ = fi17
                            }
                        } else {
                            switch -1 {
                            case -279:
                                v31 := []int{32767}
                                _ = v31
                            case 924:
                                v32 := []int{-9223372036854775808}
                                _ = v32
                            default:
                                switch (("foo bar" + "\n\t") + "hello") {
                                case "":
                                    v3 = ((v3 - 1e10) - 0.1)
                                case "hello":
                                    var v33 int8 = 42
                                    _ = v33
                                default:
                                    switch 42 {
                                    case 128:
                                        fmt.Println(v2, v2)
                                    case 32767:
                                        v3 += 0.0
                                    default:
                                        var v34 uint64 = 1
                                        _ = v34
                                    }
                                }
                            }
                            v3 = float32(float64(v3))
                            v35 := GenS2{F0: -9223372036854775808, F1: -553, F2: 65535}
                            _ = v35
                        }
                    default:
                        var v36 int16 = 1
                        _ = v36
                    }
                    v37 := GenS4{F0: -733, F1: (v2 % (((^(v2))) | 1))}
                    var v38 float64 = -2.25
                    _ = v37
                    _ = v38
                }
                switch -9223372036854775808 {
                case 128:
                    if (1 != 4294967295) {
                        fmt.Println(v3)
                        fmt.Println(v3, v3)
                        v2 += v2
                    }
                case 127:
                    v39 := GenS2{F0: -536, F1: 0, F2: 50}
                    _ = v39
                case 344:
                    if (false || (1 > -1)) {
                        var v40 int64 = 9223372036854775807
                        fmt.Println(v3, v2, v40)
                        _ = v40
                    }
                default:
                    fmt.Println(v2)
                }
                v41 := GenS5{F0: 100, F1: -9223372036854775808, F2: -1}
                _ = v41
            } else {
                for fi42 := 0; fi42 < 9; fi42++ {
                    for fi43 := 0; fi43 < 15; fi43++ {
                        fmt.Println(v2, v3, fi42)
                        _ = fi43
                    }
                    fmt.Println(v2, fi42)
                    _ = fi42
                }
                for fi44 := 0; fi44 < 10; fi44++ {
                    var v45 bool = true
                    var v46 int = (fi44 % ((9223372036854775807) | 1))
                    fmt.Println(v45)
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            v3 += v3
            if true {
                if false {
                    var v47 int = 127
                    for fi48 := 0; fi48 < 18; fi48++ {
                        if (-1 < 1) {
                            v49 := []int{v47, fi48}
                            _ = len(v49)
                            _ = v49
                        } else {
                            var v50 uint64 = 628
                            _ = v50
                        }
                        var v51 uint16 = 65535
                        v1[(-(fi48))] = 2
                        _ = fi48
                        _ = v51
                    }
                    fmt.Println(v47, v2, v47)
                    _ = v47
                }
                for fi52 := 0; fi52 < 10; fi52++ {
                    v53 := GenS2{F0: fi52, F1: -1, F2: 244}
                    _ = fi52
                    _ = v53
                }
            } else {
                v2 += (v2 >> ((255) & 63))
            }
        }
        var v54 uint = 717
        v2 += (^(v2))
        _ = v54
    } else {
        for fi55 := 0; fi55 < 16; fi55++ {
            v56 := GenS6{F0: 99, F1: 218}
            v57 := []int{(fi55 + (fi55 & (fi55 - 28))), fi55, (fi55 << ((uint(v2)) & 63)), (-(fi55))}
            _ = fi55
            _ = v56
            _ = v57
        }
        switch ("hello" + "hello") {
        case "a":
            for fi58 := 0; fi58 < 18; fi58++ {
                if true {
                    v59 := GenS2{F0: 9223372036854775807, F1: 0, F2: 1}
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 3; fi60++ {
                        var v61 int8 = -27
                        v62 := []int{fi60, fi60, 32768}
                        if true {
                            if true {
                                v63 := GenS7{F0: -1, F1: ("foo bar" + "hello")}
                                v61 -= v61
                                if (42 <= 1000) {
                                    if false {
                                        var v64 float64 = (((1e10 - -2.25) / 3.14159) - 0.1)
                                        v65 := GenS4{F0: 183, F1: v2}
                                        var v66 int64 = 42
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    } else {
                                        v67 := GenS8{F0: 0, F1: "", F2: 18446744073709551615}
                                        v3 += v3
                                        _ = v67
                                    }
                                }
                                _ = v63
                            }
                        } else {
                            v68 := GenS9{F0: 42, F1: 83}
                            var v69 int16 = int16(v2)
                            _ = v68
                            _ = v69
                        }
                        fmt.Println(fi58)
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                }
                var v70 int = (fi58 ^ (fi58 | fi58))
                v71 := []int{v70}
                _ = fi58
                _ = v70
                _ = v71
            }
        case "foo bar":
            v1[-433] = 2147483647
        }
    }
    var v72 uint32 = (^(v2))
    v3 = (1e10 / ((v3 * (3.14159 / 1.0)) + v3))
    v73 := []int{200, 42, -9223372036854775808}
    v74 := GenS10{F0: ((3.14159 - float64((v3 / (v3 + v3)))) >= 3.14159), F1: 882, F2: "foo bar"}
    switch ("\n\t" + "a") {
    case "foo bar":
        v75 := GenS6{F0: 9223372036854775807, F1: 27}
        _ = v75
    case "\n\t":
        if true {
            fmt.Println(v2)
            var v76 int = 32767
            for fi77 := 0; fi77 < 9; fi77++ {
                var v78 uint = 0
                var v79 bool = false
                _ = fi77
                _ = v78
                _ = v79
            }
            _ = v76
        } else {
            v3 = v3
            var v80 int8 = 100
            v81 := GenS11{F0: 0, F1: v72}
            _ = v80
            _ = v81
        }
    }
    for fi82 := 0; fi82 < 17; fi82++ {
        v83 := GenS10{F0: (1 == 312), F1: 18446744073709551615, F2: "foo bar"}
        fmt.Println(v72, v72, v2)
        _ = fi82
        _ = v83
    }
    var v84 int = -173
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v72:", v72)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v72
    _ = v73
    _ = v74
    _ = v84
}
