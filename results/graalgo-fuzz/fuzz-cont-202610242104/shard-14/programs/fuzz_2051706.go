package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int32; F1 string; F2 uint16 }
type GenS3 struct { F0 string; F1 int64; F2 int }
type GenS4 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS6 struct { F0 string }

func main() {
    var v1 bool = false
    var v2 uint = 0
    var v3 int = 814
    v4 := GenS1{F0: v2, F1: false}
    if (!(1.5 == 0.0)) {
        v5 := []int{int(v2), -753, (v3 * v3), 872}
        v6 := GenS1{F0: v2, F1: v1}
        _ = v5
        _ = v6
    }
    var v7 int16 = -1
    var v8 string = "a"
    if v1 {
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v8, v2, fi9)
            v3 += fi9
            fmt.Println(v8)
            _ = fi9
        }
        fmt.Println(v3, v1)
        var v10 uint64 = 4294967295
        _ = v10
    }
    v11 := GenS1{F0: v2, F1: v1}
    if (44 < 75) {
        v12 := GenS2{F0: -2147483648, F1: ((v8 + "hello") + (v8 + (("foo bar" + v8) + v8))), F2: 42}
        _ = v12
    } else {
        v7 = (v7 << ((v2) & 63))
        if v1 {
            if (((v2 != v2) || false) || (v1 && (!(-379 != 42)))) {
                v13 := GenS2{F0: -1, F1: v8, F2: 2}
                if v1 {
                    var v14 bool = (1000 == v7)
                    _ = v14
                }
                _ = v13
            }
        } else {
            v15 := []int{}
            v1 = false
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint64 = 4294967295
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            v2 = v2
            v8 = v8
            _ = fi19
        }
        if v1 {
            v20 := []int{(fi18 << ((0) & 63)), fi18, fi18, (v3 * 127)}
            _ = v20
        } else {
            if v1 {
                switch "" {
                case "\n\t":
                    switch fi18 {
                    case 42:
                        if v1 {
                            var v21 int8 = 1
                            v22 := GenS1{F0: (v2 % ((v2) | 1)), F1: v1}
                            _ = v21
                            _ = v22
                        } else {
                            v23 := []int{}
                            switch v8 {
                            case "":
                                if (!(false || false)) {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                } else {
                                    var v24 uint = v2
                                    _ = v24
                                }
                            case "a":
                                _ = len(v23)
                            default:
                                v25 := GenS2{F0: 1, F1: (v8 + v8), F2: 940}
                                _ = v25
                            }
                            _ = v23
                        }
                    }
                default:
                    for fi26 := 0; fi26 < 14; fi26++ {
                        v7 -= v7
                        if false {
                            if (234 >= 0) {
                                var v27 uint = (v2 / ((v2) | 1))
                                switch fi18 {
                                case 127:
                                    v2 += v2
                                default:
                                    if (!((fi26 / ((fi26) | 1)) <= fi26)) {
                                        v1 = v1
                                        v28 := []int{fi26, 4294967295, fi26}
                                        v29 := GenS2{F0: -730, F1: "\n\t", F2: 255}
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                if v1 {
                                    v27 += v27
                                    fmt.Println(v3)
                                } else {
                                    var v30 int = (fi18 & -448)
                                    if (!(v7 < 0)) {
                                        if v1 {
                                            var v31 uint = 42
                                            v32 := []int{int(v31), (fi18 * v3), fi18, (v3 * fi18)}
                                            _ = v31
                                            _ = v32
                                        } else {
                                            v7 = (v7 % (((v7 | (-(v7)))) | 1))
                                        }
                                        if v1 {
                                            switch ((v8 + v8) + v8) {
                                            case "":
                                                var v33 int16 = v7
                                                _ = v33
                                            case "hello":
                                                fmt.Println(fi18)
                                            case "a":
                                                v34 := []int{127, fi26, v30}
                                                _ = v34
                                            }
                                            fmt.Println(v7, fi26, fi26)
                                        }
                                        switch fi26 {
                                        case 32767:
                                            if (4294967295 != 4294967295) {
                                                v7 = v7
                                                v35 := []int{(fi18 / ((fi26) | 1)), (fi26 * fi26), 32768}
                                                _ = v35
                                            } else {
                                                v27 = (v27 ^ (v27 % ((v27) | 1)))
                                            }
                                        case 568:
                                            v36 := []int{(-(v30))}
                                            _ = v36
                                        case -1:
                                            fmt.Println(v7, v2)
                                        default:
                                            var v37 uint32 = 861
                                            _ = v37
                                        }
                                        var v38 int64 = 968
                                        _ = v38
                                    }
                                    fmt.Println(v3, v1, v3)
                                    _ = v30
                                }
                                v39 := GenS3{F0: (v8 + v8), F1: 9223372036854775807, F2: v3}
                                _ = v27
                                _ = v39
                            }
                            v1 = false
                            var v40 int64 = 97
                            _ = v40
                        }
                        var v41 int = fi26
                        _ = fi26
                        _ = v41
                    }
                }
                switch "" {
                case "\n\t":
                    var v42 int64 = 0
                    _ = v42
                }
            }
            v43 := GenS4{F0: 1, F1: 42, F2: uint8(v2)}
            for fi44 := 0; fi44 < 13; fi44++ {
                if v1 {
                    v45 := GenS3{F0: "\n\t", F1: 42, F2: (fi44 & fi18)}
                    v2 -= v2
                    _ = v45
                }
                _ = fi44
            }
            _ = v43
        }
        v46 := []int{fi18, (^(v3))}
        _ = fi18
        _ = v46
    }
    v47 := GenS1{F0: 681, F1: true}
    v7 = v7
    var v48 int8 = 0
    for fi49 := 0; fi49 < 9; fi49++ {
        v8 = "\n\t"
        var v50 int16 = 1
        v51 := GenS5{F0: 42, F1: 165, F2: (v7 % ((v50) | 1))}
        _ = fi49
        _ = v50
        _ = v51
    }
    var v52 int32 = 399
    switch (((v8 + v8) + ("a" + v8)) + v8) {
    case "":
        v53 := GenS3{F0: v8, F1: -879, F2: -2147483648}
        _ = v53
    case "hello":
        v7 -= v7
    case "\n\t":
        v54 := GenS5{F0: 250, F1: 1, F2: v7}
        _ = v54
    default:
        v55 := GenS6{F0: v8}
        _ = v55
    }
    for fi56 := 0; fi56 < 2; fi56++ {
        v57 := []int{(fi56 - fi56), fi56, v3, v3}
        v48 -= 1
        var v58 uint = 420
        _ = fi56
        _ = v57
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v52:", v52)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v52
    _ = v7
    _ = v8
}
