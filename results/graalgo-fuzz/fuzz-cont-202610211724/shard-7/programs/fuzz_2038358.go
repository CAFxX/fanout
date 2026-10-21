package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint; F1 uint16; F2 uint }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int8; F1 int16 }
type GenS12 struct { F0 uint; F1 bool }
type GenS13 struct { F0 string; F1 int8 }

func main() {
    var v1 bool = (255 != 1)
    var v2 uint = 794
    var v3 uint = (^(v2))
    var v4 float32 = 0.1
    var v5 string = "a"
    if (1 >= 42) {
        var v6 string = (v5 + "foo bar")
        _ = v6
    }
    var v7 string = v5
    v4 += (0.1 + v4)
    fmt.Println(v2, v5)
    v4 = float32(1e10)
    if (-9223372036854775808 > 1) {
        switch "" {
        case "\n\t":
            if false {
                v3 += 0
                v8 := GenS1{F0: 42, F1: 32767}
                if ("" == v5) {
                    if (-32768 >= 690) {
                        for fi9 := 0; fi9 < 13; fi9++ {
                            var v10 int8 = -1
                            _ = fi9
                            _ = v10
                        }
                    }
                    if (-43 <= 119) {
                        if false {
                            var v11 uint32 = 2
                            v11 = 426
                            var v12 uint16 = 615
                            _ = v11
                            _ = v12
                        } else {
                            v13 := []int{-579, -9223372036854775808, -590, 42}
                            fmt.Println(v2, v2, v1)
                            if (430 < 1) {
                                v14 := GenS2{F0: 1}
                                _ = v14
                            }
                            _ = v13
                        }
                        v15 := GenS3{F0: -310, F1: 86, F2: v7}
                        _ = v15
                    } else {
                        v16 := []int{9223372036854775807, 42, 32767, 127}
                        v3 -= 65535
                        _ = v16
                    }
                } else {
                    v3 = v2
                }
                _ = v8
            }
        default:
            if (v3 == 399) {
                v17 := []int{-709, -244}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 int16 = 42
                    _ = fi18
                    _ = v19
                }
            }
        }
        fmt.Println(v4, v7)
    }
    var v20 int32 = -1
    if ((((v7 + "hello") + "a") + v7) <= ((("" + v7) + v5) + (((("a" + "foo bar") + v5) + v5) + v7))) {
        v21 := []int{127, -9223372036854775808, 4294967295}
        v22 := GenS4{F0: -9223372036854775808}
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 bool = (v1 || v1)
            v25 := GenS5{F0: 127, F1: (((((("" + (v5 + ("hello" + (v7 + "\n\t")))) + ("" + v7)) + (("" + v5) + (((v7 + (((v7 + v7) + v7) + "hello")) + v5) + v7))) + (v5 + v7)) + (v7 + v5)) + (v5 + ((v7 + v7) + v5)))}
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v21
        _ = v22
    } else {
        if (!(321 >= 388)) {
            var v26 uint = (v3 * v3)
            v27 := GenS2{F0: 492}
            v5 = v7
            _ = v26
            _ = v27
        }
        var v28 uint8 = 42
        _ = v28
    }
    v29 := GenS6{F0: ((v5 + v5) + "foo bar"), F1: 599}
    switch v5 {
    case "\n\t":
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 uint = v3
                fmt.Println(fi30)
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
    case "":
        switch 4294967295 {
        case 4294967295:
            v33 := GenS7{F0: 417}
            _ = v33
        case 2147483647:
            if (!(42 < 100)) {
                switch 407 {
                case -1:
                    var v34 int8 = 91
                    _ = v34
                case 9223372036854775807:
                    for fi35 := 0; fi35 < 1; fi35++ {
                        for fi36 := 0; fi36 < 6; fi36++ {
                            if (!v1) {
                                v37 := GenS7{F0: 255}
                                switch v5 {
                                case "a":
                                    v1 = false
                                case "\n\t":
                                    if (v1 || v1) {
                                        v38 := []int{0, (fi35 % (((fi35 - fi36)) | 1)), (fi35 & (fi36 % (((-(fi36))) | 1))), (fi35 << ((1) & 63))}
                                        v38 = append(v38, (fi36 % (((fi36 + (fi35 ^ 4294967295))) | 1)))
                                        switch v7 {
                                        case "foo bar":
                                            if v1 {
                                                if (!(true || v1)) {
                                                    v4 += (1.5 - v4)
                                                } else {
                                                    v39 := GenS5{F0: 128, F1: v5}
                                                    v40 := GenS7{F0: 65535}
                                                    var v41 uint16 = 32767
                                                    switch fi36 {
                                                    case -1:
                                                        v20 = v20
                                                    case 128:
                                                        fmt.Println(fi35, v7, fi36)
                                                    default:
                                                        switch int(v3) {
                                                        case -962:
                                                            v42 := GenS5{F0: 0, F1: v7}
                                                            _ = v42
                                                        case 128:
                                                            if (((0.0 * float64(3.14159)) == 0.1) || v1) {
                                                                v43 := GenS8{F0: v20}
                                                                switch ((v5 + ("\n\t" + (v5 + v7))) + "foo bar") {
                                                                case "hello":
                                                                    v3 = v2
                                                                }
                                                                if true {
                                                                    v38 = append(v38, 127)
                                                                }
                                                                _ = v43
                                                            } else {
                                                                var v44 int8 = 22
                                                                v3 += (^(v2))
                                                                fmt.Println(v41)
                                                                _ = v44
                                                            }
                                                        default:
                                                            fmt.Println(v1)
                                                        }
                                                    }
                                                    v38 = append(v38, (-(fi36)))
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                }
                                                if v1 {
                                                    var v45 uint8 = 193
                                                    _ = v45
                                                } else {
                                                    v46 := GenS8{F0: v20}
                                                    switch (-(fi36)) {
                                                    case -9223372036854775808:
                                                        _ = len(v38)
                                                    case 32768:
                                                        v47 := GenS9{F0: 4294967295, F1: 65535, F2: v2}
                                                        _ = v47
                                                    case -876:
                                                        fmt.Println(v20, v5)
                                                        v20 -= v20
                                                    default:
                                                        var v48 uint8 = 133
                                                        _ = v48
                                                    }
                                                    _ = v46
                                                }
                                                var v49 int32 = 0
                                                var v50 int64 = 42
                                                v51 := GenS5{F0: 42, F1: v5}
                                                if len(v38) >= 2 {
                                                    v38 = v38[1:len(v38)]
                                                }
                                                if (!(!v1)) {
                                                    v50 = (v50 - (v50 & (v50 + v50)))
                                                    if v1 {
                                                        if true {
                                                            v52 := GenS10{F0: (v5 + ((("a" + "") + (v5 + (((v5 + v5) + v7) + v7))) + "hello"))}
                                                            switch 1 {
                                                            case -55:
                                                                var v53 uint16 = 264
                                                                _ = v53
                                                            case -595:
                                                                var v54 uint8 = 42
                                                                _ = v54
                                                            case -1:
                                                                switch fi35 {
                                                                case 32767:
                                                                    v49 = v49
                                                                default:
                                                                    if len(v38) >= 2 {
                                                                        v38 = v38[1:len(v38)]
                                                                    }
                                                                }
                                                            }
                                                            _ = v52
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            var v55 int = (^(fi35))
                                                            _ = v55
                                                        case "\n\t":
                                                            var v56 string = v5
                                                            _ = v56
                                                        default:
                                                            fmt.Println(v5, v3, v5)
                                                        }
                                                        var v57 int8 = 0
                                                        _ = v57
                                                    } else {
                                                        v20 = v20
                                                        switch v5 {
                                                        case "\n\t":
                                                            fmt.Println(v50)
                                                        default:
                                                            switch fi36 {
                                                            case 4294967295:
                                                                var v58 int16 = 836
                                                                _ = v58
                                                            case 42:
                                                                fmt.Println(fi35, v4, fi35)
                                                                if (v1 && false) {
                                                                    fmt.Println(v4)
                                                                } else {
                                                                    fmt.Println(v20)
                                                                    v49 = v20
                                                                    if len(v38) >= 2 {
                                                                        v38 = v38[1:len(v38)]
                                                                    }
                                                                    v5 = "a"
                                                                }
                                                            case -1:
                                                                v59 := GenS11{F0: -1, F1: 0}
                                                                _ = v59
                                                            default:
                                                                fmt.Println(fi36)
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            } else {
                                                if false {
                                                    v38 = append(v38, fi36)
                                                    v4 += v4
                                                    var v60 float64 = 0.1
                                                    _ = v60
                                                } else {
                                                    var v61 uint16 = 661
                                                    _ = v61
                                                }
                                                switch ("foo bar" + v7) {
                                                case "\n\t":
                                                    v62 := GenS12{F0: 635, F1: v1}
                                                    var v63 float64 = 1e10
                                                    if (80 > 127) {
                                                        v63 = v63
                                                    }
                                                    _ = v62
                                                    _ = v63
                                                case "foo bar":
                                                    switch v5 {
                                                    case "a":
                                                        var v64 int64 = 631
                                                        _ = v64
                                                    case "hello":
                                                        if v1 {
                                                            v65 := GenS8{F0: v20}
                                                            _ = len(v38)
                                                            v20 += v20
                                                            _ = v65
                                                        } else {
                                                            v66 := GenS9{F0: v3, F1: 65, F2: (v3 * v2)}
                                                            v1 = (32767 < fi35)
                                                            v5 = v5
                                                            _ = v66
                                                        }
                                                    }
                                                case "hello":
                                                    var v67 uint = 2
                                                    _ = v67
                                                default:
                                                    v7 = v7
                                                }
                                            }
                                        case "hello":
                                            v68 := GenS9{F0: 255, F1: 42, F2: v3}
                                            v38 = append(v38, -1)
                                            var v69 uint32 = 2
                                            v1 = (!((v3 % ((4294967295) | 1)) != v3))
                                            _ = v68
                                            _ = v69
                                        default:
                                            if len(v38) >= 2 {
                                                v38 = v38[1:len(v38)]
                                            }
                                        }
                                        _ = v38
                                    }
                                default:
                                    v70 := []int{}
                                    _ = v70
                                }
                                _ = v37
                            } else {
                                if (v5 != v5) {
                                    v71 := []int{fi36}
                                    _ = v71
                                } else {
                                    v72 := []int{(^(fi35)), fi36, (fi36 ^ fi35)}
                                    _ = v72
                                }
                                v73 := GenS1{F0: 59, F1: 578}
                                v1 = v1
                                v1 = v1
                                fmt.Println(fi36)
                                v2 = (v3 ^ 709)
                                _ = v73
                            }
                            if ((v20 | (v20 << (((v3 / ((4294967295) | 1))) & 63))) < v20) {
                                var v74 bool = (2 >= 1)
                                v4 = (v4 * v4)
                                _ = v74
                            }
                            _ = fi36
                        }
                        fmt.Println(v3, v2)
                        v75 := []int{}
                        _ = fi35
                        _ = v75
                    }
                default:
                    var v76 int16 = -1
                    _ = v76
                }
            }
        default:
            switch 127 {
            case -789:
                v5 = (v5 + "")
            }
        }
    }
    fmt.Println(v5, v7)
    var v77 int = -1
    fmt.Println(v1, v1, v77)
    var v78 string = "foo bar"
    var v79 float32 = v4
    v3 = 255
    for fi80 := 0; fi80 < 9; fi80++ {
        if (-9223372036854775808 == -524) {
            var v81 int8 = -2
            _ = v81
        }
        for fi82 := 0; fi82 < 1; fi82++ {
            switch (fi80 ^ 9223372036854775807) {
            case 9223372036854775807:
                switch (fi82 << ((v2) & 63)) {
                case 128:
                    v1 = (!(367 <= 32767))
                case 504:
                    if (127 > 0) {
                        v20 -= (-(v20))
                        v83 := GenS4{F0: fi80}
                        v20 = v20
                        v2 = 1
                        _ = v83
                    } else {
                        fmt.Println(v20, fi82, fi80)
                    }
                default:
                    fmt.Println(v20)
                }
            }
            var v84 int64 = -9223372036854775808
            v85 := GenS3{F0: (fi82 & fi82), F1: 100, F2: ((v78 + v78) + v5)}
            var v86 int = 929
            _ = fi82
            _ = v84
            _ = v85
            _ = v86
        }
        v87 := GenS13{F0: ("foo bar" + "foo bar"), F1: 127}
        _ = fi80
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v77
    _ = v78
    _ = v79
}
