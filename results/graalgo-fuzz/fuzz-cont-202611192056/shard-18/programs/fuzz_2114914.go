package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint }
type GenS4 struct { F0 int64; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int64 }

func main() {
    var v1 int = 42
    var v2 uint64 = 879
    var v3 int32 = 0
    var v4 string = (("a" + "foo bar") + "foo bar")
    var v5 int8 = -128
    if (false && true) {
        v6 := GenS1{F0: 0, F1: -610, F2: -9223372036854775808}
        v7 := []int{(-(v1)), v1}
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 931, F1: 1000, F2: -1}
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 12; fi10++ {
                v4 = v4
                _ = fi10
            }
            var v11 float32 = 3.14159
            _ = fi9
            _ = v11
        }
        switch (v1 + v1) {
        case -265:
            v2 -= v2
        case 0:
            switch v1 {
            case 32768:
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 bool = false
                    if true {
                        if true {
                            for fi14 := 0; fi14 < 5; fi14++ {
                                var v15 int64 = 0
                                v3 += -2147483648
                                _ = fi14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 8; fi16++ {
                                v17 := []int{(v1 % (((v1 + -9223372036854775808)) | 1)), fi12}
                                var v18 uint32 = 449
                                if len(v17) > 0 {
                                    v17[0] = fi16
                                }
                                var v19 float32 = -2.25
                                fmt.Println(v4, v1, v1)
                                _ = fi16
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        }
                        var v20 int32 = (v3 << ((606) & 63))
                        _ = v20
                    } else {
                        var v21 uint32 = 42
                        var v22 float32 = 0.1
                        v23 := []int{fi12}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    v24 := []int{128}
                    _ = fi12
                    _ = v13
                    _ = v24
                }
            default:
                v25 := GenS1{F0: 4294967295, F1: 32767, F2: -526}
                _ = v25
            }
        case 177:
            v26 := []int{}
            _ = v26
        default:
            var v27 uint8 = 42
            _ = v27
        }
        _ = v8
    }
    if ((0.0 < 1e10) && ((v2 > 0) && false)) {
        for fi28 := 0; fi28 < 12; fi28++ {
            v29 := GenS2{F0: -1, F1: 100}
            if true {
                v4 = ("a" + "hello")
                for fi30 := 0; fi30 < 7; fi30++ {
                    fmt.Println(v2)
                    switch v4 {
                    case "foo bar":
                        if false {
                            if true {
                                if (42 >= 0) {
                                    v31 := []int{}
                                    v1 -= (v1 >> ((79) & 63))
                                    v3 += (v3 % (((v3 | 473)) | 1))
                                    _ = v31
                                }
                            } else {
                                var v32 bool = (v2 <= 1)
                                if v32 {
                                    var v33 bool = true
                                    v34 := GenS2{F0: -1, F1: (v5 & v5)}
                                    v35 := []int{v1, -1, fi28, (v1 - fi28)}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := GenS3{F0: 0, F1: 42, F2: 2}
                                    var v37 int64 = 9223372036854775807
                                    switch (fi30 ^ (-(fi30))) {
                                    case 0:
                                        var v38 uint16 = 32767
                                        _ = v38
                                    case -886:
                                        if ((v37 >> ((65535) & 63)) != 216) {
                                            v39 := GenS2{F0: (v37 >> ((1) & 63)), F1: 0}
                                            v40 := []int{(^(fi30))}
                                            fmt.Println(v5)
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v1 -= v1
                                            var v41 float32 = 0.0
                                            _ = v41
                                        }
                                    default:
                                        if false {
                                            if false {
                                                v42 := []int{883}
                                                var v43 uint16 = 689
                                                v42 = append(v42, 127)
                                                v32 = (!(v1 < 505))
                                                v1 = fi30
                                                _ = v42
                                                _ = v43
                                            } else {
                                                switch fi30 {
                                                case 128:
                                                    v44 := GenS4{F0: 453, F1: v5, F2: 0}
                                                    _ = v44
                                                case -1:
                                                    v2 -= 0
                                                }
                                                switch 76 {
                                                case 612:
                                                    switch v4 {
                                                    case "a":
                                                        var v45 bool = ((1 > 714) || v32)
                                                        _ = v45
                                                    }
                                                case -993:
                                                    v46 := []int{fi30}
                                                    _ = v46
                                                default:
                                                    var v47 uint16 = 0
                                                    _ = v47
                                                }
                                                switch -2147483648 {
                                                case 2147483647:
                                                    var v48 int64 = (v37 << ((400) & 63))
                                                    _ = v48
                                                case -31:
                                                    v49 := GenS1{F0: 0, F1: 32767, F2: -981}
                                                    _ = v49
                                                default:
                                                    var v50 uint8 = 2
                                                    _ = v50
                                                }
                                                var v51 uint32 = 1
                                                _ = v51
                                            }
                                            v52 := GenS3{F0: 111, F1: 9223372036854775807, F2: 4294967295}
                                            switch v4 {
                                            case "foo bar":
                                                v37 = v37
                                            default:
                                                v53 := GenS1{F0: 3, F1: 32767, F2: 614}
                                                _ = v53
                                            }
                                            _ = v52
                                        }
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                _ = v32
                            }
                            var v54 float32 = 1e10
                            if true {
                                if (!true) {
                                    if false {
                                        var v55 uint8 = 54
                                        _ = v55
                                    }
                                }
                                v4 = v4
                                var v56 bool = false
                                _ = v56
                            }
                            _ = v54
                        } else {
                            var v57 uint32 = 2147483648
                            _ = v57
                        }
                    default:
                        v1 = fi28
                    }
                    var v58 float32 = ((1.5 / 1.0) - 1e10)
                    var v59 int16 = 1
                    _ = fi30
                    _ = v58
                    _ = v59
                }
            }
            var v60 uint16 = 42
            _ = fi28
            _ = v29
            _ = v60
        }
        v61 := GenS2{F0: -9223372036854775808, F1: (v5 ^ v5)}
        _ = v61
    }
    if true {
        var v62 float32 = 3.14159
        _ = v62
    } else {
        v63 := []int{v1, (v1 * (v1 << ((414) & 63))), v1, v1}
        fmt.Println(v1, v5, v4)
        v1 += v1
        _ = v63
    }
    v64 := GenS2{F0: -74, F1: v5}
    if true {
        var v65 uint = 2
        for fi66 := 0; fi66 < 3; fi66++ {
            for fi67 := 0; fi67 < 19; fi67++ {
                if false {
                    switch (v4 + v4) {
                    case "a":
                        if true {
                            var v68 int8 = (v5 - v5)
                            v65 += (v65 + v65)
                            _ = v68
                        } else {
                            var v69 uint16 = 729
                            var v70 uint32 = 2147483648
                            _ = v69
                            _ = v70
                        }
                    case "foo bar":
                        v3 = (v3 % (((v3 << ((251) & 63))) | 1))
                    case "\n\t":
                        v3 = v3
                    }
                    v65 = 1
                    v1 = (fi66 % (((v1 / (((v1 * fi67)) | 1))) | 1))
                    var v71 uint8 = 251
                    _ = v71
                }
                v72 := []int{(-(v1)), fi67}
                _ = fi67
                _ = v72
            }
            _ = fi66
        }
        fmt.Println(v4, v1)
        _ = v65
    }
    var v73 string = (v4 + v4)
    for fi74 := 0; fi74 < 2; fi74++ {
        var v75 string = "a"
        for fi76 := 0; fi76 < 16; fi76++ {
            var v77 uint8 = 127
            _ = fi76
            _ = v77
        }
        fmt.Println(v3)
        _ = fi74
        _ = v75
    }
    v3 += v3
    var v78 bool = false
    fmt.Println(v73, v4)
    var v79 int16 = int16(v2)
    for fi80 := 0; fi80 < 1; fi80++ {
        v79 = v79
        v81 := GenS5{F0: 1}
        var v82 float32 = float32((0.1 - (0.1 * ((1e10 + ((0.0 - 0.1) * 3.14159)) / 3.14159))))
        _ = fi80
        _ = v81
        _ = v82
    }
    var v83 uint = 855
    v1 = 127
    v84 := GenS6{F0: v78, F1: 42}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v73
    _ = v78
    _ = v79
    _ = v83
    _ = v84
}
