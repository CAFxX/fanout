package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483647
    var v3 uint32 = (v2 | (v2 | v2))
    if false {
        if false {
            v4 := []int{128, 191, -1}
            if false {
                var v5 int16 = -275
                v4 = append(v4, 351)
                _ = len(v4)
                _ = v5
            }
            if (0 <= 0) {
                if false {
                    if (1 != 2) {
                        v6 := GenS1{F0: 0}
                        if false {
                            v3 += v3
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = v6
                    } else {
                        var v7 uint32 = (^(v2))
                        v3 += 65535
                        fmt.Println(v3)
                        _ = v7
                    }
                    for fi8 := 0; fi8 < 4; fi8++ {
                        v2 -= 2
                        fmt.Println(v2, v3, v2)
                        _ = fi8
                    }
                } else {
                    v2 -= (v2 - 65535)
                    fmt.Println(v3)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v9 uint8 = 0
                v2 -= (v3 & 42)
                _ = v9
            }
            _ = v4
        }
        delete(v1, -1)
    }
    var v10 string = "a"
    v11 := GenS2{F0: 0, F1: 42, F2: 97}
    if (false || (!((true && true) || (true && false)))) {
        v10 = v10
        for fi12 := 0; fi12 < 8; fi12++ {
            switch v10 {
            case "foo bar":
                var v13 int16 = -32768
                _ = v13
            case "\n\t":
                v10 = v10
            default:
                v10 = (v10 + (("hello" + ("foo bar" + (v10 + v10))) + v10))
            }
            if (42 >= 0) {
                switch 1 {
                case -2147483648:
                    v10 = (v10 + ((v10 + v10) + v10))
                case 32767:
                    fmt.Println(v3)
                case -827:
                    v14 := GenS1{F0: 255}
                    _ = v14
                }
            } else {
                switch fi12 {
                case 938:
                    for fi15 := 0; fi15 < 5; fi15++ {
                        var v16 float32 = (((0.0 + (-2.25 + 1.5)) * (float32((1e10 + (1.5 - 3.14159))) * float32(1.5))) + -2.25)
                        switch int(v2) {
                        case 180:
                            fmt.Println(v10, v16)
                        default:
                            var v17 uint32 = v2
                            _ = v17
                        }
                        _ = fi15
                        _ = v16
                    }
                }
                switch v10 {
                case "foo bar":
                    if false {
                        for fi18 := 0; fi18 < 15; fi18++ {
                            if (true || false) {
                                if (true && false) {
                                    var v19 int32 = -2147483648
                                    if false {
                                        if (1000 <= -965) {
                                            v19 -= 42
                                            v19 = (v19 % ((v19) | 1))
                                            v20 := []int{fi18, fi18}
                                            _ = v20
                                        }
                                        if (fi18 != (fi12 & fi18)) {
                                            var v21 int = fi12
                                            if (!((0.1 + float32(1e10)) < 0.0)) {
                                                var v22 int32 = (v19 + v19)
                                                v23 := v1[4294967295]
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = v21
                                        } else {
                                            var v24 int = (^(fi18))
                                            if false {
                                                fmt.Println(v3, fi18)
                                                var v25 int64 = -702
                                                _ = v25
                                            } else {
                                                v26 := GenS3{F0: 600}
                                                switch fi18 {
                                                case -416:
                                                    v27 := GenS1{F0: 399}
                                                    _ = v27
                                                case 1:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if (true || (-642 > 737)) {
                                                            if (-1 > 0) {
                                                                var v28 uint32 = 1000
                                                                v2 += v28
                                                                var v29 float32 = 1.5
                                                                switch v24 {
                                                                case 9223372036854775807:
                                                                    fmt.Println(v3)
                                                                    v30 := []int{(fi12 + (fi12 + int(v2))), -9223372036854775808}
                                                                    _ = v30
                                                                case -1:
                                                                    if (false && false) {
                                                                        if false {
                                                                            var v31 uint8 = 128
                                                                            var v32 float64 = float64(((0.0 - v29) / 1.0))
                                                                            if false {
                                                                                v33 := GenS3{F0: 2}
                                                                                _ = v33
                                                                            }
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                        fmt.Println(v24)
                                                                    }
                                                                default:
                                                                    fmt.Println(fi18)
                                                                }
                                                                v1[fi12] = v28
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v34 uint16 = 887
                                                        _ = v34
                                                    case "":
                                                        v2 += 544
                                                    }
                                                default:
                                                    v10 = v10
                                                }
                                                switch v10 {
                                                case "foo bar":
                                                    v24 -= 32767
                                                default:
                                                    switch v10 {
                                                    case "foo bar":
                                                        var v35 uint32 = 102
                                                        _ = v35
                                                    case "\n\t":
                                                        fmt.Println(v24, v10, v2)
                                                    case "hello":
                                                        var v36 int32 = (-(v19))
                                                        _ = v36
                                                    default:
                                                        v37 := []int{(fi18 & (-(v24)))}
                                                        _ = v37
                                                    }
                                                }
                                                _ = v26
                                            }
                                            _ = v24
                                        }
                                        v3 = (v3 * v3)
                                        if true {
                                            v3 -= v3
                                            var v38 bool = false
                                            v38 = true
                                            switch 218 {
                                            case 1:
                                                if (true && (v38 || v38)) {
                                                    if v38 {
                                                        fmt.Println(fi18, v3, v2)
                                                    }
                                                } else {
                                                    if (!true) {
                                                        v39 := GenS1{F0: 32767}
                                                        v1[fi18] = v2
                                                        _ = v39
                                                    }
                                                }
                                            case -9223372036854775808:
                                                v40 := []int{(fi18 - -9223372036854775808), (-(fi18))}
                                                _ = v40
                                            case 9223372036854775807:
                                                v41 := GenS1{F0: 65535}
                                                _ = v41
                                            }
                                            _ = v38
                                        }
                                    } else {
                                        v3 = 65535
                                        v3 -= (v3 % ((2) | 1))
                                    }
                                    v42 := GenS4{F0: 42, F1: 42}
                                    v43 := []int{fi12, (^(fi18)), fi18}
                                    v44 := GenS4{F0: 0, F1: 1}
                                    _ = v19
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                                if true {
                                    if ((fi18 >= 2147483647) || (1 <= 0)) {
                                        fmt.Println(fi12, fi12)
                                    } else {
                                        if (0 != 255) {
                                            v10 = (("hello" + "") + (v10 + v10))
                                            var v45 float32 = 1e10
                                            v46 := GenS4{F0: 0, F1: 398}
                                            _ = v45
                                            _ = v46
                                        } else {
                                            switch fi12 {
                                            case 128:
                                                var v47 uint8 = 127
                                                _ = v47
                                            case 839:
                                                fmt.Println(fi12, v2, v3)
                                                _ = len(v1)
                                                var v48 int64 = -9223372036854775808
                                                var v49 int8 = 127
                                                v50 := GenS4{F0: 0, F1: 42}
                                                v48 = (v48 % (((v48 - v48)) | 1))
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            case -2147483648:
                                                fmt.Println(fi12)
                                            default:
                                                switch -313 {
                                                case 0:
                                                    v51 := GenS1{F0: 753}
                                                    _ = v51
                                                case 159:
                                                    v10 = v10
                                                }
                                            }
                                            if true {
                                                v52 := GenS2{F0: 42, F1: -86, F2: 100}
                                                if (-2.25 != 3.14159) {
                                                    v2 -= v3
                                                    v10 = (("" + v10) + (((((((("hello" + v10) + "foo bar") + (("" + v10) + v10)) + v10) + v10) + (v10 + v10)) + ((v10 + v10) + "a")) + v10))
                                                } else {
                                                    v1[fi18] = (^(v2))
                                                    var v53 int8 = -1
                                                    _ = v53
                                                }
                                                _ = v52
                                            }
                                        }
                                        var v54 uint8 = 255
                                        var v55 uint32 = (^(v2))
                                        _ = v54
                                        _ = v55
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v56 bool = (false && ((fi18 | fi12) != 984))
                                        _ = v56
                                    case "hello":
                                        v10 = ""
                                    case "a":
                                        if true {
                                            v57 := GenS5{F0: 1}
                                            v58 := []int{}
                                            v2 -= v2
                                            _ = v57
                                            _ = v58
                                        }
                                    }
                                } else {
                                    fmt.Println(v2)
                                    v59 := []int{}
                                    _ = v59
                                }
                                var v60 float32 = (((0.1 + 1.5) / 0.1) + 1.5)
                                _ = v60
                            }
                            v61 := GenS6{F0: "a", F1: 81, F2: 746}
                            if true {
                                if (!(1 > 1000)) {
                                    v3 = v3
                                    v3 = 891
                                    fmt.Println(v3, v10)
                                    switch fi12 {
                                    case 380:
                                        v62 := GenS6{F0: v10, F1: 227, F2: 42}
                                        _ = v62
                                    case 32768:
                                        if ((65535 <= 0) && true) {
                                            var v63 string = v10
                                            v64 := v1[(fi18 & 739)]
                                            var v65 int32 = 1
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                } else {
                                    var v66 int32 = 42
                                    v67 := []int{(^(fi18)), 0, (-(fi12))}
                                    _ = v66
                                    _ = v67
                                }
                                v3 = v2
                                switch (v10 + ((v10 + v10) + v10)) {
                                case "":
                                    v3 += (^(v2))
                                case "\n\t":
                                    var v68 int16 = 1
                                    _ = v68
                                case "foo bar":
                                    if ((32767 <= 22) || (1 == -1)) {
                                        var v69 float32 = 0.0
                                        var v70 bool = (532 <= 32767)
                                        var v71 uint64 = 1
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    } else {
                                        v2 -= (v2 & (v3 - (^(v2))))
                                        v72 := GenS5{F0: 1}
                                        v73 := GenS2{F0: 1, F1: 127, F2: 42}
                                        _ = v72
                                        _ = v73
                                    }
                                }
                            }
                            _ = fi18
                            _ = v61
                        }
                    }
                default:
                    fmt.Println(fi12, v2)
                }
                var v74 uint32 = (v3 ^ (v3 >> ((2) & 63)))
                _ = v74
            }
            v3 = v3
            _ = fi12
        }
        v10 = v10
    } else {
        var v75 uint64 = 18446744073709551615
        v3 -= v2
        var v76 int16 = 479
        _ = v75
        _ = v76
    }
    v77 := GenS5{F0: -88}
    fmt.Println(v2, v3)
    v10 = (v10 + ("\n\t" + v10))
    if (true && (false && false)) {
        if (!false) {
            var v78 string = ((("\n\t" + v10) + "") + (v10 + (v10 + (v10 + v10))))
            v3 += v2
            _ = v78
        }
        v79 := []int{4294967295, 629, -772}
        _ = v79
    } else {
        switch -2147483648 {
        case 32768:
            if (true && true) {
                v1[32768] = (v3 + v3)
            }
        case 933:
            switch 32768 {
            case 937:
                var v80 uint64 = 18446744073709551615
                _ = v80
            }
        }
        v2 = v3
    }
    if (float64(0.1) > float64((3.14159 / (float32(1e10) * float32(1.5))))) {
        v81 := []int{32767, 0}
        var v82 int16 = 1
        _ = v81
        _ = v82
    } else {
        var v83 float32 = 1e10
        _ = v83
    }
    v10 = v10
    var v84 int64 = 14
    delete(v1, 2147483647)
    var v85 int8 = int8(v84)
    fmt.Println(v10)
    if false {
        fmt.Println(v84, v3, v3)
    } else {
        for fi86 := 0; fi86 < 7; fi86++ {
            if (-2.25 == 3.14159) {
                var v87 string = v10
                v88 := []int{fi86}
                _ = v87
                _ = v88
            } else {
                v89 := GenS1{F0: 42}
                _ = v89
            }
            var v90 uint8 = 255
            _ = fi86
            _ = v90
        }
        v84 = v84
    }
    fmt.Println(v10, v10)
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v77
    _ = v84
    _ = v85
}
