package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint32; F1 int; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int16; F1 uint; F2 int32 }
type GenS8 struct { F0 uint32; F1 int16; F2 int8 }
type GenS9 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 uint64 }
type GenS12 struct { F0 int16; F1 uint64; F2 string }
func genFn0(p0 uint32, p1 uint64) uint8 {
    var v7 int64 = -295
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int16 = 32767
        v7 = 875
        p0 = (p0 - 1)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: "", F1: 100}
    v11 := GenS2{F0: 255, F1: 1, F2: 65535}
    v7 -= (v7 << ((42) & 63))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    return 127
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    var v4 string = "a"
    var v5 string = v4
    var v6 int64 = -1
    delete(v2, 318)
    v12 := GenS3{F0: (v1 - v1), F1: false}
    fmt.Println(v5)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int32 = 852
        fmt.Println(v6, v1)
        _ = fi13
        _ = v14
    }
    switch -636 {
    case 9223372036854775807:
        if false {
            v15 := GenS4{F0: 243, F1: -899, F2: 2}
            if false {
                v5 = v5
            }
            v1 += (v1 / ((428) | 1))
            _ = v15
        } else {
            v16 := GenS5{F0: -153}
            _ = v16
        }
    case 128:
        v17 := GenS6{F0: 0, F1: v6}
        _ = v17
    case 0:
        switch (v5 + v4) {
        case "hello":
            var v18 string = v5
            _ = v18
        default:
            for fi19 := 0; fi19 < 11; fi19++ {
                v5 = ((((v5 + v4) + v5) + "") + (("\n\t" + v4) + v4))
                for fi20 := 0; fi20 < 4; fi20++ {
                    v6 = (v6 ^ v6)
                    if (true && false) {
                        v1 = (^(v1))
                        if false {
                            v21 := []int{(fi19 << ((0) & 63)), (fi20 << ((1) & 63))}
                            _ = v21
                        } else {
                            if true {
                                fmt.Println(fi20)
                                if (true || (float32((3.14159 - -2.25)) <= float32(float64(((0.0 - 3.14159) - 0.0))))) {
                                    v22 := []int{}
                                    v4 = ((v4 + "a") + (v5 + v4))
                                    _ = v22
                                } else {
                                    fmt.Println(v4, v5)
                                }
                            } else {
                                var v23 bool = false
                                _ = v23
                            }
                            v1 = v1
                        }
                        var v24 int64 = -1
                        _ = v24
                    } else {
                        var v25 bool = (!false)
                        v26 := []int{-9223372036854775808, 4294967295, 1, fi20}
                        if true {
                            fmt.Println(v5, v4)
                        } else {
                            _ = len(v26)
                        }
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS7{F0: 173, F1: 0, F2: 0}
                    _ = fi20
                    _ = v27
                }
                _ = fi19
            }
        }
    default:
        var v28 int32 = 576
        _ = v28
    }
    if (!(false || (v1 >= v1))) {
        v5 = ("\n\t" + v4)
        if true {
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 int32 = 1
                if ((v1 * (v1 | (v1 * 42))) <= v1) {
                    var v31 float32 = ((3.14159 - 0.1) + 1.5)
                    for fi32 := 0; fi32 < 15; fi32++ {
                        if (1000 != 713) {
                            v33 := []int{389, fi32, (fi32 & fi29), (-(fi29))}
                            v30 += 0
                            v4 = (v4 + v4)
                            _ = v33
                        } else {
                            v1 = v1
                            v34 := GenS8{F0: 42, F1: -239, F2: 100}
                            var v35 int8 = 124
                            _ = v34
                            _ = v35
                        }
                        if (fi29 != (fi29 | (fi29 - fi32))) {
                            switch v5 {
                            case "foo bar":
                                if true {
                                    v5 = v5
                                    _ = len(v3)
                                    var v36 uint8 = 0
                                    _ = v36
                                }
                            default:
                                if (127 > 148) {
                                    if ((true || false) || true) {
                                        if true {
                                            v37 := GenS2{F0: uint(v1), F1: 255, F2: v1}
                                            switch ((v5 + v5) + "a") {
                                            case "":
                                                v30 = v30
                                            case "\n\t":
                                                if (0 >= (-(v6))) {
                                                    v1 -= 65535
                                                    switch (fi29 - (fi29 - fi29)) {
                                                    case -2147483648:
                                                        v31 = v31
                                                    default:
                                                        if true {
                                                            v2[(fi29 % ((fi29) | 1))] = 2
                                                            var v38 int32 = -1
                                                            if ((true || (((fi29 >= int(v38)) || false) && (!false))) || false) {
                                                                v6 += v6
                                                            }
                                                            _ = v38
                                                        }
                                                    }
                                                } else {
                                                    var v39 float32 = (v31 * 1e10)
                                                    _ = v39
                                                }
                                            }
                                            fmt.Println(fi32)
                                            _ = v37
                                        } else {
                                            v40 := []int{32767, 32767, -651, (fi32 / (((fi29 >> ((152) & 63))) | 1))}
                                            if ((!((false || false) && (!false))) && (42 >= 42)) {
                                                v2[(fi32 - fi29)] = 398
                                                _ = len(v3)
                                            }
                                            _ = v40
                                        }
                                    } else {
                                        if false {
                                            var v41 uint32 = 2147483647
                                            if true {
                                                v1 -= v1
                                                v42 := GenS9{F0: false, F1: v41, F2: 2}
                                                v43 := GenS3{F0: (v1 * v1), F1: (true && (0 <= 0))}
                                                _ = v42
                                                _ = v43
                                            }
                                            _ = v41
                                        } else {
                                            if true {
                                                if (!(true || (v6 > v6))) {
                                                    var v44 float64 = float64(0.1)
                                                    _ = v44
                                                }
                                                delete(v3, 745)
                                                if false {
                                                    var v45 int64 = 1
                                                    _ = v45
                                                } else {
                                                    fmt.Println(v1, v30, v4)
                                                }
                                            }
                                        }
                                        var v46 uint32 = 2147483648
                                        v47 := []int{fi32}
                                        if true {
                                            var v48 uint64 = 2
                                            _ = v48
                                        } else {
                                            v31 = 0.0
                                            if len(v47) > 0 {
                                                v47[0] = (fi29 - -183)
                                            }
                                            _ = len(v47)
                                        }
                                        _ = v46
                                        _ = v47
                                    }
                                } else {
                                    var v49 uint64 = 1
                                    _ = v49
                                }
                            }
                        }
                        _ = fi32
                    }
                    _ = v31
                } else {
                    v50 := GenS10{F0: "a"}
                    var v51 uint64 = 2
                    _ = v50
                    _ = v51
                }
                v5 = "foo bar"
                _ = fi29
                _ = v30
            }
            v1 -= v1
            var v52 bool = false
            _ = v52
        }
        var v53 uint64 = 173
        _ = v53
    } else {
        switch 128 {
        case -915:
            v54 := GenS7{F0: -186, F1: 18446744073709551615, F2: -270}
            _ = v54
        case 32767:
            var v55 int64 = -1
            _ = v55
        case 106:
            v56 := []int{42}
            _ = v56
        }
        fmt.Println(v5, v4, v6)
        if true {
            switch 126 {
            case 573:
                switch 42 {
                case 2147483647:
                    v6 -= v6
                case -903:
                    if (-128 <= -1) {
                        var v57 bool = true
                        switch 992 {
                        case 4294967295:
                            if (v6 < (v6 + (v6 >> ((823) & 63)))) {
                                v1 += 255
                            }
                        case 32767:
                            v6 = v6
                        case 128:
                            fmt.Println(v1)
                        default:
                            var v58 uint8 = 255
                            _ = v58
                        }
                        _ = v57
                    }
                }
            case -9223372036854775808:
                for fi59 := 0; fi59 < 6; fi59++ {
                    var v60 int32 = 1
                    var v61 int64 = 0
                    v62 := GenS8{F0: 65535, F1: 0, F2: 0}
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v62
                }
            case 2147483647:
                var v63 uint32 = 2
                _ = v63
            }
            v6 -= (v6 << ((255) & 63))
            var v64 int32 = 1
            _ = v64
        } else {
            v65 := []int{4294967295, -2147483648, 128}
            v5 = (v5 + v4)
            _ = v65
        }
    }
    v4 = v5
    fmt.Println(v6, v6, v1)
    switch 434 {
    case -1:
        v6 -= (v6 - (v6 * 0))
    case -282:
        for fi66 := 0; fi66 < 15; fi66++ {
            v67 := []int{209, fi66}
            _ = fi66
            _ = v67
        }
    case -9223372036854775808:
        fmt.Println(v6)
    default:
        for fi68 := 0; fi68 < 20; fi68++ {
            v69 := GenS11{F0: 0, F1: 5}
            v70 := []int{fi68, fi68, fi68}
            for fi71 := 0; fi71 < 1; fi71++ {
                v72 := GenS6{F0: 44, F1: v6}
                fmt.Println(v5)
                var v73 int8 = 1
                v74 := GenS4{F0: 2147483648, F1: 9223372036854775807, F2: 2}
                _ = fi71
                _ = v72
                _ = v73
                _ = v74
            }
            _ = fi68
            _ = v69
            _ = v70
        }
    }
    v6 = 0
    v3[4294967295] = 65535
    var v75 int8 = 127
    v4 = ((v4 + "a") + v5)
    if false {
        v75 += (v75 << ((18446744073709551615) & 63))
    } else {
        v76 := GenS12{F0: 130, F1: 907, F2: (v5 + v4)}
        var v77 float32 = ((((3.14159 + (((3.14159 / ((3.14159 - 3.14159) + -2.25)) * (1.5 - 1.5)) * 0.1)) / 1.0) / 1.5) + 0.0)
        _ = v76
        _ = v77
    }
    v78 := GenS8{F0: 4294967295, F1: 32767, F2: 123}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v78
}
