package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint; F1 int; F2 bool }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32; F1 bool }
func genFn0() uint32 {
    v5 := []int{558, 862}
    var v6 float32 = (3.14159 * 1e10)
    if false {
        if false {
            v6 = ((v6 - v6) * 0.0)
            v7 := GenS1{F0: 0, F1: 25}
            if false {
                v8 := GenS1{F0: 2, F1: 461}
                v6 -= 1.5
                switch "foo bar" {
                case "a":
                    var v9 int8 = 0
                    _ = v9
                case "hello":
                    var v10 bool = ((((v6 - (((3.14159 / v6) - v6) * ((0.0 / float32(float64(v6))) + float32(-2.25)))) + v6) * v6) <= v6)
                    _ = v10
                case "\n\t":
                    if false {
                        switch "a" {
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = -1
                            }
                        case "hello":
                            v11 := GenS2{F0: 236}
                            _ = v11
                        case "foo bar":
                            fmt.Println(v6)
                        }
                        switch "foo bar" {
                        case "":
                            if ((0.1 >= v6) && false) {
                                v6 = v6
                            } else {
                                v6 += ((float32(((((0.0 / float64(((3.14159 - 1e10) - (3.14159 * float32(1e10))))) * 3.14159) * 1e10) - 0.1)) / (1.5 + v6)) + float32(1e10))
                                var v12 uint16 = 32767
                                _ = v12
                            }
                        case "hello":
                            if len(v5) > 0 {
                                v5[0] = 493
                            }
                        case "a":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    } else {
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    }
                default:
                    if false {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            var v14 uint16 = 1
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        switch "hello" {
                        case "a":
                            for fi15 := 0; fi15 < 15; fi15++ {
                                v6 += (v6 * ((v6 * v6) / (v6 / (3.14159 + 3.14159))))
                                v6 -= v6
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    fmt.Println(fi15, fi15)
                                    var v17 uint = 465
                                    v18 := GenS2{F0: 58}
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi15
                            }
                        case "\n\t":
                            _ = len(v5)
                        case "foo bar":
                            for fi19 := 0; fi19 < 19; fi19++ {
                                if ((true || (("foo bar" + ("" + "a")) <= "foo bar")) || ("hello" <= ("\n\t" + "foo bar"))) {
                                    var v20 uint16 = 42
                                    v21 := GenS3{F0: 439}
                                    v22 := GenS4{F0: (("a" + "hello") + ("\n\t" + "")), F1: 0, F2: 0}
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi19
                            }
                        default:
                            if false {
                                _ = len(v5)
                                var v23 string = ""
                                _ = v23
                            }
                        }
                        var v24 uint8 = 128
                        _ = v24
                    }
                }
                _ = v8
            }
            _ = v7
        }
        v5 = append(v5, 181)
    }
    if (127 > 100) {
        var v25 bool = (!(true && (42 > 0)))
        if (2147483647 <= 32767) {
            v26 := GenS2{F0: 2}
            fmt.Println(v6, v25)
            fmt.Println(v6)
            _ = v26
        } else {
            if true {
                v6 -= v6
                v6 += ((1.5 / (3.14159 / 3.14159)) + ((((1e10 / v6) / (v6 - (0.0 * (v6 / (0.1 / 3.14159))))) * v6) / v6))
            }
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := GenS1{F0: 127, F1: 255}
                var v29 uint16 = 242
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = v25
    }
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483647
    var v3 uint16 = 65535
    v4 := make(map[uint32]string)
    v3 = v3
    v2 += (v2 << ((255) & 63))
    for fi30 := 0; fi30 < 3; fi30++ {
        v2 = genFn0()
        v2 = v2
        v31 := GenS5{F0: 255, F1: int(v3), F2: true}
        _ = fi30
        _ = v31
    }
    var v32 string = "foo bar"
    switch v32 {
    case "\n\t":
        v33 := GenS6{F0: v32, F1: -75}
        _ = v33
    default:
        var v34 bool = (v3 != v3)
        _ = v34
    }
    switch -1 {
    case -473:
        fmt.Println(v32)
    case 702:
        for fi35 := 0; fi35 < 20; fi35++ {
            for fi36 := 0; fi36 < 4; fi36++ {
                if (0 <= 0) {
                    var v37 int = fi35
                    _ = v37
                } else {
                    v38 := v4[v2]
                    v3 = v3
                    _ = v38
                }
                v39 := GenS7{F0: (v3 & 65535), F1: 2}
                var v40 string = (v32 + v32)
                _ = fi36
                _ = v39
                _ = v40
            }
            if (v2 > (v2 << ((713) & 63))) {
                if ((42 != -2147483648) && (!(!(v2 >= v2)))) {
                    var v41 int8 = 76
                    var v42 bool = false
                    _ = v41
                    _ = v42
                } else {
                    v4[genFn0()] = ""
                    var v43 uint16 = (v3 << ((255) & 63))
                    _ = v43
                }
                switch (-(fi35)) {
                case 127:
                    v3 = (v3 >> ((255) & 63))
                default:
                    v44 := []int{fi35, fi35, 128}
                    _ = v44
                }
            } else {
                v45 := []int{}
                _ = v45
            }
            switch fi35 {
            case -642:
                var v46 uint32 = v2
                _ = v46
            default:
                v47 := GenS1{F0: 255, F1: 18446744073709551615}
                _ = v47
            }
            _ = fi35
        }
    case 42:
        _ = len(v1)
    }
    var v48 uint64 = 1
    var v49 bool = true
    fmt.Println(v49, v48)
    v50 := GenS1{F0: 6, F1: uint(v48)}
    fmt.Println(v2, v48)
    fmt.Println(v2)
    for fi51 := 0; fi51 < 10; fi51++ {
        v3 -= uint16(v2)
        v52 := GenS8{F0: 888, F1: (-1 == 0)}
        var v53 int32 = -2147483648
        _ = fi51
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 16; fi54++ {
        for fi55 := 0; fi55 < 17; fi55++ {
            var v56 float32 = (-2.25 / 1.0)
            var v57 int16 = int16(v3)
            _ = fi55
            _ = v56
            _ = v57
        }
        v32 = "\n\t"
        _ = fi54
    }
    if false {
        var v58 bool = (!(-168 != 42))
        if v49 {
            var v59 int16 = 0
            for fi60 := 0; fi60 < 13; fi60++ {
                v61 := []int{}
                for fi62 := 0; fi62 < 18; fi62++ {
                    switch 128 {
                    case 128:
                        if false {
                            if v58 {
                                var v63 uint = uint(v3)
                                fmt.Println(v2, fi60, v58)
                                var v64 bool = true
                                _ = v63
                                _ = v64
                            }
                            if (-1 < 1) {
                                if len(v61) > 0 {
                                    v61[0] = (-(fi60))
                                }
                            } else {
                                v58 = true
                                v2 = 234
                                if (-32768 >= v59) {
                                    var v65 uint = 881
                                    var v66 bool = (!false)
                                    v65 = (v65 % (((v65 + 4294967295)) | 1))
                                    switch fi60 {
                                    case -595:
                                        fmt.Println(v48)
                                    case 4294967295:
                                        v48 -= v48
                                    }
                                    _ = v65
                                    _ = v66
                                }
                                _ = len(v1)
                            }
                            fmt.Println(v49, fi62, v48)
                        } else {
                            fmt.Println(v32, v3)
                            switch 0 {
                            case 128:
                                v67 := GenS9{F0: (-267 == (v59 & v59))}
                                _ = v67
                            case -9223372036854775808:
                                fmt.Println(v48)
                            case 488:
                                var v68 int32 = int32(fi60)
                                _ = v68
                            }
                            v69 := GenS5{F0: 921, F1: fi60, F2: (-2.25 > ((-2.25 - 0.0) + float32(0.0)))}
                            _ = v69
                        }
                    default:
                        switch "a" {
                        case "a":
                            var v70 uint64 = v48
                            _ = v70
                        case "":
                            fmt.Println(fi62, v48, fi60)
                        default:
                            fmt.Println(v59, v58, v2)
                        }
                    }
                    switch "foo bar" {
                    case "hello":
                        switch -1 {
                        case 32767:
                            var v71 int32 = -2147483648
                            _ = v71
                        case -988:
                            var v72 float64 = (0.1 * -2.25)
                            _ = v72
                        case 430:
                            if (!v58) {
                                fmt.Println(fi60, fi60)
                                var v73 int32 = -393
                                v74 := GenS9{F0: false}
                                _ = v73
                                _ = v74
                            } else {
                                if v49 {
                                    _ = len(v4)
                                    if v49 {
                                        v2 = v2
                                        var v75 uint8 = 255
                                        _ = v75
                                    } else {
                                        var v76 bool = (9223372036854775807 != -9223372036854775808)
                                        var v77 string = (("hello" + (((v32 + v32) + ("\n\t" + "hello")) + v32)) + (v32 + v32))
                                        _ = v76
                                        _ = v77
                                    }
                                    if (!(((v32 + (v32 + "\n\t")) + "a") >= "hello")) {
                                        var v78 uint16 = 255
                                        var v79 uint16 = v3
                                        _ = v78
                                        _ = v79
                                    }
                                }
                            }
                        default:
                            var v80 uint64 = v48
                            _ = v80
                        }
                    case "a":
                        switch fi62 {
                        case -2147483648:
                            if (265 >= v3) {
                                var v81 string = "hello"
                                _ = v81
                            }
                        case -1:
                            var v82 uint32 = uint32(v3)
                            _ = v82
                        case 262:
                            v83 := GenS9{F0: false}
                            _ = v83
                        default:
                            v3 += 2
                        }
                        var v84 int = fi60
                        v58 = (2147483647 != -1)
                        _ = v84
                    case "":
                        v85 := GenS1{F0: 255, F1: 65535}
                        _ = v85
                    }
                    if (("a" > v32) || true) {
                        switch -434 {
                        case 65:
                            fmt.Println(fi62, v59, fi60)
                        }
                        switch 69 {
                        case 9223372036854775807:
                            fmt.Println(v49, v2, v59)
                        case 32768:
                            var v86 int16 = 32767
                            _ = v86
                        case 285:
                            var v87 float64 = 0.1
                            _ = v87
                        default:
                            v88 := v1[(-(fi60))]
                            _ = v88
                        }
                    }
                    fmt.Println(fi60, fi60)
                    _ = fi62
                }
                _ = fi60
                _ = v61
            }
            if (int8(v3) >= 100) {
                v89 := GenS10{F0: 42, F1: false}
                v90 := GenS4{F0: v32, F1: v48, F2: (v48 / ((18446744073709551615) | 1))}
                _ = v89
                _ = v90
            } else {
                var v91 uint64 = v48
                _ = v91
            }
            _ = v59
        } else {
            fmt.Println(v49)
        }
        var v92 int16 = 1000
        _ = v58
        _ = v92
    } else {
        var v93 bool = v49
        v48 = (v48 * (v48 & (v48 | v48)))
        v49 = (!false)
        _ = v93
    }
    if v49 {
        for fi94 := 0; fi94 < 4; fi94++ {
            v2 = v2
            var v95 int8 = 61
            _ = fi94
            _ = v95
        }
    } else {
        var v96 int64 = 42
        var v97 uint = 65535
        _ = v96
        _ = v97
    }
    for fi98 := 0; fi98 < 1; fi98++ {
        fmt.Println(v3, fi98, fi98)
        _ = fi98
    }
    var v99 int32 = -721
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v99
}
