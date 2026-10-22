package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS8 struct { F0 int; F1 uint16; F2 uint64 }

func main() {
    var v1 uint = 685
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 127
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: v3, F1: (2147483647 >= 65535), F2: 2}
        _ = fi6
        _ = v7
    }
    if ((false || (!(((751 == 0) || true) || ((-1 <= 1) && false)))) && ((-10 > 29) && true)) {
        var v8 int = -2147483648
        _ = v8
    }
    if false {
        if (!(32768 <= 128)) {
            if ((true && (100 == 1)) || (!(!(4294967295 <= 32768)))) {
                v3 = (v3 >> (((v1 % ((v1) | 1))) & 63))
                if false {
                    v1 = v1
                    if false {
                        v3 = 42
                    } else {
                        v3 = v3
                        var v9 uint = v1
                        delete(v5, 4294967295)
                        _ = v9
                    }
                    fmt.Println(v1)
                } else {
                    fmt.Println(v3, v3)
                    var v10 int64 = 0
                    _ = v10
                }
                fmt.Println(v3, v1)
            }
            if (42 < 0) {
                fmt.Println(v1, v1)
            }
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                fmt.Println(v1, v3)
                _ = fi11
            }
            var v12 uint8 = v3
            _ = v12
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(fi13)
            fmt.Println(v3, v1, v3)
            _ = fi13
        }
    } else {
        v3 = 251
    }
    var v14 int64 = 0
    var v15 uint = v1
    switch "\n\t" {
    case "a":
        switch -641 {
        case -667:
            v16 := GenS1{F0: 2, F1: true, F2: 18446744073709551615}
            _ = v16
        default:
            for fi17 := 0; fi17 < 17; fi17++ {
                if false {
                    v3 = (v3 * (v3 / ((v3) | 1)))
                    v18 := []int{}
                    _ = v18
                } else {
                    var v19 int32 = 2147483647
                    _ = v19
                }
                var v20 int8 = 45
                _ = fi17
                _ = v20
            }
        }
    case "":
        for fi21 := 0; fi21 < 10; fi21++ {
            if false {
                for fi22 := 0; fi22 < 12; fi22++ {
                    v1 += 1
                    _ = len(v4)
                    var v23 uint = v15
                    _ = fi22
                    _ = v23
                }
            }
            v3 = 127
            v24 := GenS1{F0: uint8(v15), F1: true, F2: 0}
            _ = fi21
            _ = v24
        }
    }
    v25 := GenS2{F0: (((("hello" + "") + "\n\t") + "hello") + ("hello" + "")), F1: v1}
    if false {
        var v26 int64 = int64(v3)
        if false {
            delete(v4, 128)
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            if (2147483647 > 157) {
                var v28 float64 = (1.5 + (0.1 + float64((0.1 * 3.14159))))
                var v29 bool = false
                _ = v28
                _ = v29
            } else {
                if (fi27 >= 42) {
                    v30 := []int{(fi27 * fi27), (fi27 & fi27)}
                    _ = v30
                }
                v31 := GenS3{F0: fi27}
                v32 := []int{fi27, fi27, (fi27 - (fi27 >> (((v1 << ((v1) & 63))) & 63)))}
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 5; fi33++ {
                if ((uint32(v15) < 451) || ((!false) || ((float32(-2.25) > ((float32(-2.25) + (0.0 / 1e10)) / (3.14159 / (float32(1.5) * -2.25)))) || true))) {
                    switch (("\n\t" + "\n\t") + "a") {
                    case "\n\t":
                        v34 := []int{(^(fi27)), (fi33 | (fi33 | (fi27 / (((fi27 >> ((v1) & 63))) | 1)))), fi27}
                        _ = v34
                    case "foo bar":
                        v35 := []int{(fi33 * fi27)}
                        _ = v35
                    case "a":
                        v36 := GenS4{F0: -128, F1: v3}
                        _ = v36
                    default:
                        var v37 uint8 = v3
                        _ = v37
                    }
                    var v38 uint8 = v3
                    v14 += -1
                    _ = v38
                } else {
                    var v39 float32 = 1e10
                    v40 := []int{fi33}
                    _ = v39
                    _ = v40
                }
                fmt.Println(v3, fi27)
                switch "" {
                case "a":
                    var v41 int = (fi33 | (fi33 % (((fi27 * (^(fi33)))) | 1)))
                    _ = v41
                case "foo bar":
                    v1 = v15
                }
                if false {
                    v15 -= v1
                    v15 = v1
                } else {
                    var v42 int32 = 1
                    fmt.Println(v26, v14)
                    if false {
                        v43 := GenS1{F0: (v3 + 98), F1: true, F2: 42}
                        v15 = v1
                        v14 = v26
                        _ = v43
                    } else {
                        v26 -= v26
                    }
                    _ = v42
                }
                v14 += v26
                if true {
                    fmt.Println(v15, v14)
                    var v44 uint = (^(v1))
                    _ = v44
                } else {
                    var v45 bool = ("foo bar" == "foo bar")
                    v15 -= (^(v1))
                    if v45 {
                        v46 := GenS5{F0: 1, F1: 42}
                        var v47 uint16 = 255
                        switch ("hello" + "hello") {
                        case "hello":
                            if true {
                                fmt.Println(fi33)
                                v48 := GenS6{F0: v47}
                                _ = v48
                            } else {
                                var v49 uint8 = (v3 ^ (v3 - (v3 | v3)))
                                _ = v49
                            }
                        }
                        _ = v46
                        _ = v47
                    } else {
                        switch ("" + "\n\t") {
                        case "a":
                            v50 := GenS1{F0: (v3 | 42), F1: false, F2: 2}
                            _ = v50
                        case "foo bar":
                            if v45 {
                                if v45 {
                                    v51 := GenS3{F0: fi27}
                                    _ = v51
                                } else {
                                    switch 2147483647 {
                                    case 0:
                                        var v52 uint = (v1 ^ (v15 / ((2) | 1)))
                                        v53 := []int{(^(fi33)), -708, fi27, (^(fi27))}
                                        _ = v52
                                        _ = v53
                                    case 9223372036854775807:
                                        v15 = v1
                                    default:
                                        v54 := []int{(-(fi33)), fi33, (fi27 ^ 127)}
                                        _ = v54
                                    }
                                    var v55 bool = (!false)
                                    switch (fi27 >> ((v1) & 63)) {
                                    case -2147483648:
                                        v45 = (!v45)
                                    default:
                                        var v56 int8 = -64
                                        _ = v56
                                    }
                                    _ = v55
                                }
                                var v57 float32 = 0.1
                                _ = v57
                            } else {
                                v15 = (v15 & uint(v3))
                                if (0.1 == ((0.1 * (1.5 - 0.0)) + (0.1 - -2.25))) {
                                    fmt.Println(v15, fi27)
                                    v58 := GenS7{F0: 4294967295, F1: -32768, F2: 145}
                                    var v59 float64 = float64(1e10)
                                    var v60 uint32 = 293
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                } else {
                                    delete(v4, (^(fi33)))
                                    var v61 int16 = int16(v14)
                                    fmt.Println(fi33)
                                    _ = v61
                                }
                            }
                        case "\n\t":
                            v15 = v1
                        }
                        var v62 uint32 = 65535
                        v63 := GenS8{F0: (^(fi33)), F1: 32767, F2: 2}
                        _ = v62
                        _ = v63
                    }
                    v14 = -1
                    v14 += int64(v1)
                    v14 = v14
                    var v64 uint32 = 0
                    switch 1 {
                    case 79:
                        var v65 int = (fi27 / ((fi27) | 1))
                        _ = v65
                    case 245:
                        var v66 float32 = (0.0 * 3.14159)
                        _ = v66
                    case 128:
                        v26 += v26
                    }
                    _ = v45
                    _ = v64
                }
                _ = fi33
            }
            v1 += v15
            _ = fi27
        }
        _ = v26
    } else {
        if false {
            v1 -= (^(v1))
            _ = len(v4)
            v15 = v15
        } else {
            if true {
                if false {
                    if true {
                        v67 := []int{127, 1, 42, 1}
                        _ = v67
                    } else {
                        var v68 bool = true
                        _ = v68
                    }
                }
                var v69 int32 = -1
                _ = v69
            }
            v70 := []int{463, 128, 32768}
            var v71 float32 = float32(0.1)
            _ = v70
            _ = v71
        }
        var v72 string = "hello"
        _ = v72
    }
    if true {
        v15 = 1
        var v73 uint8 = v3
        var v74 int8 = 1
        _ = v73
        _ = v74
    }
    if true {
        if true {
            if (683 > 32767) {
                v3 = (^(v3))
                v14 += (v14 >> (((v15 + 42)) & 63))
                v3 += (v3 & 255)
            } else {
                if (false || false) {
                    if (4294967295 < 18446744073709551615) {
                        fmt.Println(v14, v15)
                        var v75 float64 = float64(1e10)
                        _ = v75
                    } else {
                        var v76 int16 = 0
                        v14 = (-(v14))
                        fmt.Println(v1, v3, v3)
                        _ = v76
                    }
                    v1 = v15
                    v14 = (v14 + v14)
                }
            }
        } else {
            var v77 int64 = v14
            for fi78 := 0; fi78 < 11; fi78++ {
                switch "a" {
                case "":
                    var v79 int64 = -499
                    _ = v79
                case "\n\t":
                    if (v15 < 42) {
                        v1 -= 1
                        var v80 string = "a"
                        _ = v80
                    } else {
                        for fi81 := 0; fi81 < 18; fi81++ {
                            var v82 int16 = -749
                            fmt.Println(v82, v15)
                            var v83 uint64 = 852
                            v77 = -9223372036854775808
                            v1 -= (v15 & (v15 & v1))
                            if (42 > (v82 ^ (v82 + v82))) {
                                var v84 int32 = 1
                                _ = v84
                            } else {
                                v14 = (v77 * (v14 << ((v1) & 63)))
                            }
                            _ = fi81
                            _ = v82
                            _ = v83
                        }
                        if false {
                            switch 128 {
                            case -1:
                                v14 = 42
                            case 127:
                                fmt.Println(v15, v14)
                            case -749:
                                v85 := []int{fi78, (-(fi78))}
                                _ = v85
                            default:
                                for fi86 := 0; fi86 < 18; fi86++ {
                                    fmt.Println(v1, fi78)
                                    v87 := []int{fi78, fi86, int(v1), (fi78 << ((1) & 63))}
                                    _ = fi86
                                    _ = v87
                                }
                            }
                            var v88 float64 = ((((0.1 - -2.25) / (float64(3.14159) + (((0.0 / (-2.25 * 0.1)) / 1.0) - 1.5))) + 0.0) - 0.1)
                            _ = v88
                        }
                    }
                case "a":
                    if (true && false) {
                        for fi89 := 0; fi89 < 12; fi89++ {
                            _ = len(v5)
                            _ = fi89
                        }
                    }
                default:
                    v90 := GenS4{F0: -1, F1: v3}
                    _ = v90
                }
                v14 = -1
                _ = fi78
            }
            _ = v77
        }
        switch 32767 {
        case -2147483648:
            var v91 uint64 = 873
            _ = v91
        case 42:
            for fi92 := 0; fi92 < 1; fi92++ {
                fmt.Println(v1, v3)
                v1 -= (v15 % (((^(v1))) | 1))
                fmt.Println(v14)
                _ = fi92
            }
        }
        for fi93 := 0; fi93 < 1; fi93++ {
            v94 := GenS1{F0: (v3 * 86), F1: (42 == 65535), F2: 739}
            fmt.Println(v3)
            _ = fi93
            _ = v94
        }
    } else {
        fmt.Println(v3, v1)
    }
    var v95 int8 = -68
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v95:", v95)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v95
}
