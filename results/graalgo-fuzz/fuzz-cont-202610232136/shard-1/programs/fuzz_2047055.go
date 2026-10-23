package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int16; F1 uint8; F2 int }
type GenS5 struct { F0 int8; F1 int8; F2 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int }
func genFn0() int64 {
    if true {
        for fi4 := 0; fi4 < 10; fi4++ {
            var v5 uint16 = 2
            _ = fi4
            _ = v5
        }
        v6 := GenS1{F0: 350, F1: 53, F2: 0}
        v7 := []int{128, 42, 2147483647, 2147483647}
        _ = v6
        _ = v7
    } else {
        var v8 int = -274
        var v9 int16 = -32768
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 127, F2: 496}
    _ = v10
    return -455
}

func main() {
    var v1 uint32 = 1
    var v2 int32 = 42
    var v3 int8 = int8(v2)
    v11 := GenS1{F0: 42, F1: 42, F2: 42}
    var v12 int8 = v3
    v1 += (v1 << ((1) & 63))
    switch 2147483647 {
    case 9223372036854775807:
        v13 := []int{2147483647, 1, 42}
        _ = v13
    case 4294967295:
        if ((974 >= 1) && false) {
            v3 = v3
            v14 := GenS2{F0: 0, F1: -9223372036854775808, F2: (false || (4 != 188))}
            var v15 int64 = int64(v12)
            _ = v14
            _ = v15
        } else {
            v16 := GenS3{F0: 1, F1: (v2 == 1)}
            if false {
                v12 = (v12 / ((71) | 1))
                var v17 uint16 = 821
                fmt.Println(v17, v17)
                _ = v17
            } else {
                if ((float64((((0.0 * (-2.25 - ((0.0 / float32(((3.14159 * 0.0) - 1e10))) - ((0.1 + 0.0) - -2.25)))) - ((float32(3.14159) - 0.1) + (0.0 - (0.1 + float32(0.0))))) * (0.0 * 3.14159))) / 0.1) != -2.25) {
                    v18 := []int{-589, 1}
                    if (!true) {
                        v12 = 100
                    } else {
                        if false {
                            if len(v18) > 0 {
                                v18[0] = 128
                            }
                        } else {
                            for fi19 := 0; fi19 < 3; fi19++ {
                                for fi20 := 0; fi20 < 5; fi20++ {
                                    v21 := GenS4{F0: 1000, F1: 127, F2: (fi19 + (fi19 - 42))}
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    if true {
                                        switch (fi19 >> ((255) & 63)) {
                                        case 42:
                                            v12 = (^(v12))
                                        case 32768:
                                            var v22 uint64 = 951
                                            _ = v22
                                        }
                                        v2 = (v2 & 1)
                                    } else {
                                        v1 -= (v1 + (v1 + (v1 * 2)))
                                        v23 := GenS1{F0: -9223372036854775808, F1: 42, F2: 2}
                                        _ = v23
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                                _ = fi19
                            }
                            v12 = v3
                            for fi24 := 0; fi24 < 6; fi24++ {
                                if false {
                                    v25 := GenS2{F0: -1, F1: 42, F2: false}
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    _ = v25
                                } else {
                                    v2 = -1
                                    if false {
                                        var v26 bool = true
                                        for fi27 := 0; fi27 < 9; fi27++ {
                                            if len(v18) > 0 {
                                                v18[0] = (fi27 & fi24)
                                            }
                                            _ = fi27
                                        }
                                        if (18446744073709551615 == 886) {
                                            if (-1 > 482) {
                                                var v28 uint16 = 639
                                                v12 += v3
                                                for fi29 := 0; fi29 < 6; fi29++ {
                                                    if len(v18) > 0 {
                                                        v18[0] = fi29
                                                    }
                                                    if len(v18) >= 2 {
                                                        v18 = v18[1:len(v18)]
                                                    }
                                                    _ = fi29
                                                }
                                                _ = v28
                                            }
                                            v18 = append(v18, fi24)
                                        }
                                        _ = v26
                                    } else {
                                        var v30 float64 = 0.1
                                        var v31 int32 = v2
                                        var v32 int64 = genFn0()
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                v33 := GenS5{F0: (v3 + (-(v3))), F1: v3, F2: 386}
                                _ = fi24
                                _ = v33
                            }
                        }
                    }
                    _ = v18
                }
                switch 128 {
                case 863:
                    v1 = v1
                case -418:
                    v3 = (v3 & (v3 << ((366) & 63)))
                }
            }
            _ = v16
        }
    default:
        for fi34 := 0; fi34 < 6; fi34++ {
            v1 -= (v1 >> ((65535) & 63))
            var v35 int16 = 1000
            _ = fi34
            _ = v35
        }
    }
    switch -506 {
    case 32767:
        var v36 int16 = 684
        _ = v36
    case 2147483647:
        var v37 int16 = -32768
        _ = v37
    case 132:
        for fi38 := 0; fi38 < 20; fi38++ {
            v39 := GenS6{F0: v2}
            if (true || (897 < 42)) {
                for fi40 := 0; fi40 < 6; fi40++ {
                    v2 = v2
                    _ = fi40
                }
                if false {
                    switch ("\n\t" + "a") {
                    case "a":
                        v41 := []int{fi38, (fi38 << ((163) & 63)), fi38}
                        _ = v41
                    default:
                        for fi42 := 0; fi42 < 3; fi42++ {
                            v12 -= v12
                            v43 := []int{(fi42 ^ 0), (fi38 | fi38)}
                            v44 := GenS5{F0: v12, F1: v12, F2: 574}
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    }
                    for fi45 := 0; fi45 < 13; fi45++ {
                        v12 = v12
                        v46 := []int{int(v2), (-(fi38)), 0, -2147483648}
                        _ = fi45
                        _ = v46
                    }
                } else {
                    for fi47 := 0; fi47 < 18; fi47++ {
                        if true {
                            var v48 float64 = ((3.14159 / 1.0) + 3.14159)
                            v49 := GenS7{F0: 2147483648}
                            _ = v48
                            _ = v49
                        } else {
                            if ((1000 <= 42) && false) {
                                switch (fi38 - (fi38 % (((fi38 * (fi47 | (-(fi47))))) | 1))) {
                                case 2147483647:
                                    v50 := []int{fi47}
                                    _ = v50
                                }
                            }
                            switch int(v3) {
                            case 2147483647:
                                var v51 uint32 = v1
                                _ = v51
                            case 9223372036854775807:
                                v1 -= (^(v1))
                            case 32767:
                                if (216 < 18446744073709551615) {
                                    var v52 uint = 1
                                    var v53 int = -2147483648
                                    var v54 int64 = -1
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                            default:
                                v55 := GenS1{F0: genFn0(), F1: 0, F2: 42}
                                _ = v55
                            }
                            var v56 int8 = v12
                            _ = v56
                        }
                        var v57 uint = 833
                        var v58 string = "a"
                        _ = fi47
                        _ = v57
                        _ = v58
                    }
                }
            }
            _ = fi38
            _ = v39
        }
    }
    if (1.5 < 3.14159) {
        if true {
            v59 := GenS8{F0: -1}
            _ = v59
        } else {
            v60 := []int{9223372036854775807, 127, -748}
            if len(v60) > 0 {
                v60[0] = 9223372036854775807
            }
            _ = v60
        }
        v12 = v3
    } else {
        var v61 uint8 = 69
        v62 := []int{}
        _ = v61
        _ = v62
    }
    switch -2147483648 {
    case 32767:
        fmt.Println(v1, v3, v1)
    case -2147483648:
        var v63 int = 127
        _ = v63
    }
    if (2 > 0) {
        var v64 int = -2147483648
        fmt.Println(v3, v64, v12)
        _ = v64
    }
    var v65 uint64 = 73
    for fi66 := 0; fi66 < 16; fi66++ {
        v67 := []int{fi66, -696, 2147483647}
        var v68 uint8 = 255
        switch fi66 {
        case 586:
            for fi69 := 0; fi69 < 15; fi69++ {
                if len(v67) > 0 {
                    v67[0] = -65
                }
                _ = fi69
            }
        case 1:
            for fi70 := 0; fi70 < 14; fi70++ {
                var v71 int32 = v2
                if false {
                    v72 := GenS5{F0: 100, F1: 42, F2: 0}
                    _ = v72
                }
                _ = fi70
                _ = v71
            }
        default:
            v12 = (v12 % (((v3 / ((v3) | 1))) | 1))
        }
        _ = fi66
        _ = v67
        _ = v68
    }
    v2 = (v2 / (((v2 - v2)) | 1))
    fmt.Println(v12, v65)
    var v73 int = 0
    var v74 uint8 = 30
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v65
    _ = v73
    _ = v74
}
