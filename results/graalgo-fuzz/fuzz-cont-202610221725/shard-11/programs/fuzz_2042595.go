package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int16; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }

func main() {
    var v1 string = "a"
    var v2 bool = (980 <= 1)
    var v3 int = 4294967295
    v1 = v1
    switch 128 {
    case 42:
        var v4 uint8 = 0
        _ = v4
    case 340:
        fmt.Println(v1, v1, v1)
    case -342:
        if (42 == 65535) {
            v5 := GenS1{F0: 0, F1: 1, F2: "foo bar"}
            for fi6 := 0; fi6 < 15; fi6++ {
                v2 = (fi6 != v3)
                if (((v3 << ((65535) & 63)) <= 0) || ("\n\t" > v1)) {
                    fmt.Println(fi6, v2, fi6)
                }
                if (!v2) {
                    var v7 float64 = -2.25
                    v8 := []int{325, v3, 128}
                    var v9 uint = 65535
                    _ = v7
                    _ = v8
                    _ = v9
                } else {
                    fmt.Println(v1)
                    switch (v1 + "hello") {
                    case "hello":
                        var v10 uint8 = 0
                        _ = v10
                    case "a":
                        v11 := []int{v3}
                        _ = v11
                    case "":
                        v12 := GenS1{F0: 1, F1: 1, F2: ("foo bar" + v1)}
                        _ = v12
                    default:
                        if false {
                            for fi13 := 0; fi13 < 9; fi13++ {
                                if v2 {
                                    v14 := GenS1{F0: 42, F1: 703, F2: v1}
                                    var v15 int32 = 2147483647
                                    var v16 int64 = -1
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    if true {
                                        v1 = ("a" + "a")
                                    } else {
                                        v17 := []int{152, (fi6 % (((v3 * v3)) | 1))}
                                        if false {
                                            var v18 float32 = -2.25
                                            _ = v18
                                        }
                                        _ = v17
                                    }
                                    var v19 int8 = -128
                                    switch "a" {
                                    case "hello":
                                        var v20 float32 = 0.1
                                        _ = v20
                                    default:
                                        v21 := GenS2{F0: 42}
                                        _ = v21
                                    }
                                    _ = v19
                                }
                                _ = fi13
                            }
                            var v22 int = (fi6 ^ fi6)
                            if v2 {
                                v3 += 1
                                v23 := []int{(^(fi6))}
                                _ = v23
                            } else {
                                if (-601 < -2147483648) {
                                    if true {
                                        var v24 uint8 = 128
                                        v22 = 0
                                        if (!(1 > 362)) {
                                            v25 := GenS3{F0: (fi6 % ((111) | 1)), F1: 48, F2: 1}
                                            v26 := GenS4{F0: 42}
                                            v27 := []int{(v3 | int(v24)), (-(v22)), fi6, (^(fi6))}
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v2 = (0 <= 1)
                                            v22 += v22
                                        }
                                        _ = v24
                                    }
                                    if v2 {
                                        fmt.Println(v1, v3, v1)
                                        var v28 uint8 = 42
                                        _ = v28
                                    } else {
                                        var v29 int8 = -24
                                        if (-1 < 1) {
                                            fmt.Println(v2, v22, v29)
                                            if (2 > 786) {
                                                var v30 uint = 706
                                                _ = v30
                                            } else {
                                                for fi31 := 0; fi31 < 13; fi31++ {
                                                    v32 := []int{(v3 ^ fi31), 127, 92, v22}
                                                    if len(v32) >= 2 {
                                                        v32 = v32[1:len(v32)]
                                                    }
                                                    var v33 uint16 = 32767
                                                    v34 := GenS5{F0: 0, F1: 42, F2: int(v33)}
                                                    var v35 uint8 = 2
                                                    _ = fi31
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                }
                                                for fi36 := 0; fi36 < 13; fi36++ {
                                                    v2 = false
                                                    var v37 int16 = 42
                                                    _ = fi36
                                                    _ = v37
                                                }
                                            }
                                            for fi38 := 0; fi38 < 17; fi38++ {
                                                if true {
                                                    var v39 uint64 = 1
                                                    _ = v39
                                                } else {
                                                    var v40 int = 85
                                                    v22 -= v3
                                                    _ = v40
                                                }
                                                v22 = (fi6 - (v22 >> ((255) & 63)))
                                                v41 := []int{(fi38 / ((248) | 1))}
                                                _ = fi38
                                                _ = v41
                                            }
                                        } else {
                                            var v42 uint = 255
                                            _ = v42
                                        }
                                        _ = v29
                                    }
                                } else {
                                    v43 := GenS6{F0: (2147483647 < 2147483647)}
                                    var v44 bool = v2
                                    _ = v43
                                    _ = v44
                                }
                                v3 = 1
                                var v45 uint = 855
                                _ = v45
                            }
                            _ = v22
                        } else {
                            switch v1 {
                            case "foo bar":
                                fmt.Println(v1, v3, fi6)
                            case "\n\t":
                                var v46 int8 = 42
                                _ = v46
                            case "":
                                switch "\n\t" {
                                case "a":
                                    var v47 float64 = 0.1
                                    _ = v47
                                default:
                                    if false {
                                        fmt.Println(fi6)
                                    } else {
                                        if (481 > 1000) {
                                            if (18446744073709551615 == 2) {
                                                v48 := []int{}
                                                for fi49 := 0; fi49 < 3; fi49++ {
                                                    switch (v1 + "foo bar") {
                                                    case "foo bar":
                                                        switch 2147483647 {
                                                        case 9223372036854775807:
                                                            switch -1 {
                                                            case 424:
                                                                var v50 int32 = 0
                                                                _ = v50
                                                            case 4294967295:
                                                                v2 = (2147483647 >= -861)
                                                            default:
                                                                if len(v48) > 0 {
                                                                    v48[0] = (fi49 * fi49)
                                                                }
                                                                if len(v48) >= 2 {
                                                                    v48 = v48[1:len(v48)]
                                                                }
                                                                var v51 int8 = 100
                                                                _ = v51
                                                            }
                                                        case 32768:
                                                            if true {
                                                                v48 = append(v48, (fi6 << ((102) & 63)))
                                                                fmt.Println(fi6, v3)
                                                                v3 = 128
                                                            } else {
                                                                if ((-170 >= -65) || false) {
                                                                    v1 = v1
                                                                    if ((false || v2) || (v2 && (0.0 >= float32(float64((((3.14159 / 1.0) - 1e10) + (0.1 * float32(-2.25)))))))) {
                                                                        v1 = ("a" + (v1 + (v1 + v1)))
                                                                        var v52 uint8 = 0
                                                                        _ = v52
                                                                    }
                                                                    v1 = (v1 + "hello")
                                                                }
                                                                switch v1 {
                                                                case "hello":
                                                                    v3 = (^(fi49))
                                                                case "":
                                                                    var v53 string = ("foo bar" + v1)
                                                                    _ = v53
                                                                }
                                                            }
                                                            if len(v48) >= 2 {
                                                                v48 = v48[1:len(v48)]
                                                            }
                                                        case -434:
                                                            v2 = false
                                                        default:
                                                            v3 -= v3
                                                        }
                                                    case "hello":
                                                        v3 += 4294967295
                                                    case "\n\t":
                                                        v54 := GenS4{F0: 1}
                                                        _ = v54
                                                    }
                                                    if len(v48) > 0 {
                                                        v48[0] = 127
                                                    }
                                                    _ = fi49
                                                }
                                                for fi55 := 0; fi55 < 12; fi55++ {
                                                    var v56 int32 = 0
                                                    _ = fi55
                                                    _ = v56
                                                }
                                                _ = v48
                                            }
                                        } else {
                                            for fi57 := 0; fi57 < 16; fi57++ {
                                                v58 := GenS3{F0: (fi57 + (v3 >> ((2) & 63))), F1: 100, F2: 36}
                                                v59 := GenS4{F0: 65535}
                                                v60 := GenS3{F0: (^(v3)), F1: 1, F2: -1}
                                                _ = fi57
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                            for fi61 := 0; fi61 < 14; fi61++ {
                                                var v62 uint32 = 65535
                                                var v63 int8 = -1
                                                _ = fi61
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                        var v64 uint = 490
                                        _ = v64
                                    }
                                }
                            default:
                                v1 = v1
                            }
                            if v2 {
                                v2 = (2 > 932)
                            } else {
                                v2 = (1 != 2)
                                var v65 float32 = 3.14159
                                if (2 != 312) {
                                    if v2 {
                                        var v66 int64 = 42
                                        _ = v66
                                    } else {
                                        if v2 {
                                            fmt.Println(v3)
                                            if false {
                                                var v67 int = 574
                                                if (true || (!((100 < 100) || (32767 < 96)))) {
                                                    var v68 int32 = -2147483648
                                                    var v69 uint32 = 704
                                                    fmt.Println(v65)
                                                    _ = v68
                                                    _ = v69
                                                }
                                                _ = v67
                                            } else {
                                                v70 := GenS6{F0: false}
                                                var v71 uint16 = 65535
                                                _ = v70
                                                _ = v71
                                            }
                                        } else {
                                            for fi72 := 0; fi72 < 6; fi72++ {
                                                v73 := []int{-1, (v3 / (((fi6 * fi72)) | 1)), fi6}
                                                v65 -= v65
                                                _ = fi72
                                                _ = v73
                                            }
                                        }
                                        if true {
                                            v65 += ((1e10 / 0.1) * v65)
                                            v74 := []int{(v3 + v3)}
                                            _ = v74
                                        }
                                    }
                                } else {
                                    fmt.Println(v2, v3, v3)
                                    var v75 uint8 = 127
                                    v76 := GenS7{F0: 9223372036854775807}
                                    _ = v75
                                    _ = v76
                                }
                                _ = v65
                            }
                        }
                    }
                }
                _ = fi6
            }
            v77 := []int{93, v3, v3}
            _ = v5
            _ = v77
        } else {
            v78 := []int{v3, -343}
            switch ((v1 + "\n\t") + ("hello" + (v1 + (v1 + (("\n\t" + "a") + "foo bar"))))) {
            case "foo bar":
                for fi79 := 0; fi79 < 13; fi79++ {
                    if (v2 || (2 < 236)) {
                        v2 = v2
                        if v2 {
                            v80 := GenS5{F0: 42, F1: 0, F2: -525}
                            _ = v80
                        }
                    }
                    _ = fi79
                }
            case "hello":
                _ = len(v78)
            case "\n\t":
                if v2 {
                    if true {
                        v2 = (!(!false))
                    }
                    var v81 string = "hello"
                    var v82 float32 = 3.14159
                    _ = v81
                    _ = v82
                } else {
                    switch v3 {
                    case 128:
                        for fi83 := 0; fi83 < 12; fi83++ {
                            var v84 int32 = -2147483648
                            var v85 int64 = 42
                            _ = fi83
                            _ = v84
                            _ = v85
                        }
                    default:
                        if len(v78) > 0 {
                            v78[0] = (v3 >> ((2) & 63))
                        }
                    }
                }
            default:
                v3 = v3
            }
            var v86 uint16 = 42
            _ = v78
            _ = v86
        }
    }
    var v87 uint = 255
    v2 = (!(-145 != 0))
    v1 = v1
    for fi88 := 0; fi88 < 5; fi88++ {
        var v89 uint32 = 2
        _ = fi88
        _ = v89
    }
    var v90 int16 = 1000
    v90 += (v90 + (v90 - v90))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v87:", v87)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v87
    _ = v90
}
