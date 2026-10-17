package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint; F1 int64; F2 uint }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int64; F1 bool; F2 uint32 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int32; F1 int; F2 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 uint = 2
    var v4 int64 = 1
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 32767, F1: 2}
            _ = fi6
            _ = v7
        }
        var v8 string = ("hello" + ("hello" + ("foo bar" + "foo bar")))
        for fi9 := 0; fi9 < 16; fi9++ {
            switch "" {
            case "\n\t":
                v8 = ("a" + v8)
            case "foo bar":
                v10 := GenS1{F0: 42, F1: 18446744073709551615}
                _ = v10
            default:
                fmt.Println(fi9, fi5)
            }
            v8 = ((v8 + v8) + "\n\t")
            _ = fi9
        }
        _ = fi5
        _ = v8
    }
    v11 := GenS2{F0: v3, F1: -1, F2: v3}
    v4 = -403
    switch ("a" + "hello") {
    case "hello":
        var v12 float32 = 0.1
        _ = v12
    case "a":
        v4 = (v4 ^ 42)
    case "\n\t":
        switch ((("\n\t" + "") + "\n\t") + "\n\t") {
        case "":
            fmt.Println(v3)
        case "foo bar":
            if true {
                v2 = 877
                if false {
                    v13 := GenS3{F0: 0, F1: 973}
                    v14 := GenS4{F0: -1}
                    _ = v13
                    _ = v14
                }
            } else {
                switch 2147483647 {
                case -9223372036854775808:
                    delete(v1, (("" + "") + "hello"))
                case 32768:
                    for fi15 := 0; fi15 < 1; fi15++ {
                        var v16 int32 = -2147483648
                        v2 -= 0
                        for fi17 := 0; fi17 < 14; fi17++ {
                            var v18 int = (fi17 & fi15)
                            switch "foo bar" {
                            case "":
                                var v19 string = ""
                                _ = v19
                            }
                            v20 := GenS5{F0: 1}
                            _ = fi17
                            _ = v18
                            _ = v20
                        }
                        _ = fi15
                        _ = v16
                    }
                case 42:
                    v4 = v4
                }
                if (!("\n\t" < "\n\t")) {
                    var v21 uint16 = 563
                    switch 32768 {
                    case 1:
                        for fi22 := 0; fi22 < 17; fi22++ {
                            for fi23 := 0; fi23 < 13; fi23++ {
                                if (!(false && true)) {
                                    v24 := GenS6{F0: 2, F1: fi23}
                                    var v25 float64 = float64((1.5 / 3.14159))
                                    _ = v24
                                    _ = v25
                                }
                                fmt.Println(v4)
                                _ = fi23
                            }
                            fmt.Println(v2, v2)
                            v4 = (-(v4))
                            _ = fi22
                        }
                    case -2147483648:
                        var v26 uint = 1
                        _ = v26
                    case 642:
                        for fi27 := 0; fi27 < 13; fi27++ {
                            if true {
                                if true {
                                    for fi28 := 0; fi28 < 2; fi28++ {
                                        fmt.Println(fi27, fi28, v4)
                                        v21 = v21
                                        _ = fi28
                                    }
                                    v4 = 0
                                } else {
                                    var v29 uint8 = 253
                                    v4 = 0
                                    if false {
                                        v2 = uint32(v29)
                                        v21 = v21
                                    } else {
                                        fmt.Println(fi27, v2)
                                        switch 260 {
                                        case 9223372036854775807:
                                            if false {
                                                fmt.Println(fi27, v3, v3)
                                                v4 -= (v4 ^ v4)
                                            }
                                        default:
                                            for fi30 := 0; fi30 < 18; fi30++ {
                                                fmt.Println(v2, v3, v21)
                                                var v31 int32 = 1
                                                _ = fi30
                                                _ = v31
                                            }
                                        }
                                    }
                                    _ = v29
                                }
                                if true {
                                    v32 := GenS7{F0: -1, F1: false, F2: (v2 | (v2 * v2))}
                                    if ("foo bar" < ("" + "\n\t")) {
                                        v21 = (v21 ^ v21)
                                        var v33 uint32 = (v2 >> (((^(v3))) & 63))
                                        v34 := []int{2147483647, fi27, (fi27 | fi27), (fi27 ^ int(v2))}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        delete(v1, "\n\t")
                                    }
                                    if (((0.1 - 3.14159) != (float32(0.1) / (1.5 + 1.5))) && (!false)) {
                                        for fi35 := 0; fi35 < 5; fi35++ {
                                            if ((1e10 * 0.0) <= 1.5) {
                                                var v36 bool = false
                                                switch (fi35 - (fi27 << ((v3) & 63))) {
                                                case -515:
                                                    switch fi27 {
                                                    case 0:
                                                        var v37 int64 = 566
                                                        _ = v37
                                                    default:
                                                        var v38 float64 = float64(-2.25)
                                                        _ = v38
                                                    }
                                                case -294:
                                                    var v39 uint64 = 1
                                                    _ = v39
                                                }
                                                delete(v1, "hello")
                                                _ = v36
                                            } else {
                                                var v40 int64 = (v4 * v4)
                                                _ = v40
                                            }
                                            _ = fi35
                                        }
                                        var v41 uint8 = 42
                                        v42 := GenS2{F0: v3, F1: (v4 << (((v3 + 18446744073709551615)) & 63)), F2: (v3 >> ((18446744073709551615) & 63))}
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v21 -= (^(v21))
                                    }
                                    _ = v32
                                }
                            } else {
                                for fi43 := 0; fi43 < 14; fi43++ {
                                    fmt.Println(v2, v21)
                                    fmt.Println(fi43)
                                    _ = fi43
                                }
                                _ = len(v1)
                                fmt.Println(v2)
                            }
                            for fi44 := 0; fi44 < 20; fi44++ {
                                fmt.Println(v2, fi44, fi44)
                                v45 := GenS2{F0: v3, F1: (v4 + (^(v4))), F2: 255}
                                _ = fi44
                                _ = v45
                            }
                            v46 := []int{(fi27 % (((^(fi27))) | 1)), (fi27 - (fi27 % ((fi27) | 1))), (^(fi27))}
                            _ = fi27
                            _ = v46
                        }
                    default:
                        v2 += v2
                    }
                    var v47 string = ""
                    _ = v21
                    _ = v47
                }
                switch "" {
                case "foo bar":
                    var v48 uint32 = v2
                    _ = v48
                case "":
                    fmt.Println(v3)
                case "hello":
                    v2 = v2
                }
            }
        case "a":
            var v49 int16 = 930
            _ = v49
        default:
            var v50 int64 = v4
            _ = v50
        }
    default:
        v51 := GenS8{F0: 42, F1: true}
        _ = v51
    }
    var v52 float32 = (1.5 * 0.1)
    if (true || false) {
        for fi53 := 0; fi53 < 19; fi53++ {
            delete(v1, "hello")
            v2 += (v2 * v2)
            var v54 string = (("\n\t" + ("\n\t" + "")) + (("" + ("hello" + "\n\t")) + "foo bar"))
            _ = fi53
            _ = v54
        }
        switch "hello" {
        case "foo bar":
            if false {
                v4 = (v4 - 1)
            }
        case "hello":
            for fi55 := 0; fi55 < 2; fi55++ {
                v2 += (^(v2))
                v56 := GenS4{F0: (v4 << ((v3) & 63))}
                _ = fi55
                _ = v56
            }
        case "":
            v2 += 2
        }
    } else {
        v52 = v52
        switch 128 {
        case 2147483647:
            v4 = v4
        case -144:
            v57 := GenS7{F0: 1, F1: false, F2: (v2 * (v2 & v2))}
            _ = v57
        case 879:
            for fi58 := 0; fi58 < 15; fi58++ {
                v59 := []int{fi58, (fi58 - (fi58 / ((321) | 1))), fi58, fi58}
                v52 = v52
                for fi60 := 0; fi60 < 11; fi60++ {
                    var v61 uint8 = 255
                    v62 := GenS3{F0: (v61 << ((v3) & 63)), F1: v3}
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                _ = fi58
                _ = v59
            }
        }
    }
    for fi63 := 0; fi63 < 1; fi63++ {
        fmt.Println(v52, v52, v4)
        _ = fi63
    }
    var v64 int16 = -995
    for fi65 := 0; fi65 < 4; fi65++ {
        v66 := GenS4{F0: v4}
        v67 := GenS9{F0: (v2 << (((v3 & uint(v2))) & 63))}
        _ = fi65
        _ = v66
        _ = v67
    }
    v2 = (v2 * (v2 << ((v3) & 63)))
    v68 := GenS9{F0: v2}
    v69 := []int{1, -9223372036854775808}
    v70 := GenS1{F0: (^(v64)), F1: uint64(v3)}
    for fi71 := 0; fi71 < 6; fi71++ {
        for fi72 := 0; fi72 < 20; fi72++ {
            v3 -= v3
            switch (fi72 / ((-102) | 1)) {
            case 127:
                switch "\n\t" {
                case "a":
                    var v73 float64 = 3.14159
                    _ = v73
                }
            case 379:
                if false {
                    if false {
                        v3 -= v3
                        fmt.Println(v4, v4, fi71)
                        delete(v1, ("a" + "a"))
                        if true {
                            switch 9223372036854775807 {
                            case 2147483647:
                                if (-1 <= 1) {
                                    v74 := GenS5{F0: -1}
                                    var v75 uint32 = v2
                                    _ = v74
                                    _ = v75
                                } else {
                                    fmt.Println(v4, v4, fi71)
                                }
                            case 32768:
                                v76 := GenS4{F0: 720}
                                _ = v76
                            }
                            fmt.Println(fi71)
                            v77 := GenS10{F0: -744, F1: (fi72 & (fi72 / (((-(fi72))) | 1))), F2: ("\n\t" + "foo bar")}
                            if (v4 > (-(v4))) {
                                var v78 float32 = v52
                                v79 := v1["foo bar"]
                                _ = v78
                                _ = v79
                            }
                            v80 := GenS11{F0: 0}
                            var v81 int16 = 32767
                            _ = v77
                            _ = v80
                            _ = v81
                        } else {
                            v64 += v64
                            if len(v69) >= 2 {
                                v69 = v69[1:len(v69)]
                            }
                            if false {
                                v82 := v1[""]
                                v1["a"] = v82
                                if (65535 > 42) {
                                    var v83 string = "hello"
                                    switch (fi71 & (fi71 / (((fi72 % ((42) | 1))) | 1))) {
                                    case -1:
                                        if false {
                                            v84 := GenS12{F0: fi72}
                                            _ = v84
                                        }
                                    case 42:
                                        var v85 int = fi72
                                        _ = v85
                                    }
                                    v1["foo bar"] = 2
                                    _ = v83
                                }
                                _ = v82
                            }
                        }
                    }
                    v4 += v4
                } else {
                    var v86 int16 = v64
                    _ = v86
                }
            case -9223372036854775808:
                var v87 float32 = 0.1
                _ = v87
            }
            var v88 int8 = 42
            var v89 string = ""
            var v90 int8 = -1
            _ = fi72
            _ = v88
            _ = v89
            _ = v90
        }
        var v91 uint32 = 681
        var v92 int = (fi71 | -1)
        _ = fi71
        _ = v91
        _ = v92
    }
    v93 := GenS13{F0: 478, F1: ("a" + "foo bar")}
    fmt.Println(v4, v4)
    var v94 uint8 = 190
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v64
    _ = v68
    _ = v69
    _ = v70
    _ = v93
    _ = v94
}
