package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint }
type GenS6 struct { F0 uint64; F1 string }
func genFn0(p0 uint32, p1 uint, p2 uint32) int {
    v7 := []int{}
    v8 := GenS1{F0: (^(p0)), F1: 4294967295}
    v7 = append(v7, 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 bool = (v1 == 170)
    var v5 uint64 = 2
    var v6 string = ""
    if v4 {
        var v9 int16 = 1
        var v10 int16 = v9
        var v11 float64 = 1e10
        _ = v10
        _ = v11
        _ = v9
    } else {
        v12 := GenS1{F0: 0, F1: 879}
        _ = v12
    }
    if (!(18446744073709551615 > 18446744073709551615)) {
        fmt.Println(v6)
    }
    var v13 int16 = 1
    v5 = 1
    if false {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 bool = true
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := []int{(fi14 | 825), (-(fi14)), (fi16 << ((293) & 63))}
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    } else {
        v18 := v2[2147483647]
        _ = v18
    }
    var v19 string = v6
    v20 := GenS1{F0: 2, F1: (v5 | v5)}
    var v21 uint = 568
    var v22 float32 = ((0.1 / -2.25) + 3.14159)
    switch -2147483648 {
    case -542:
        var v23 uint64 = (v5 >> (((v21 / ((v21) | 1))) & 63))
        _ = v23
    }
    _ = len(v3)
    v13 = -1
    v24 := GenS1{F0: uint32(v1), F1: (v5 << ((42) & 63))}
    if (66 >= 127) {
        delete(v3, 2147483647)
    } else {
        v5 += 0
        if false {
            v5 += uint64(v1)
            v25 := GenS1{F0: 1, F1: v5}
            _ = v25
        } else {
            var v26 float64 = 0.0
            _ = v26
        }
        if (!v4) {
            switch 971 {
            case 907:
                v5 = (v5 * v5)
            case -672:
                for fi27 := 0; fi27 < 19; fi27++ {
                    v28 := []int{115, genFn0(729, v21, uint32(v21)), (^(fi27)), -1}
                    var v29 uint = (v21 % ((v21) | 1))
                    v3[4294967295] = 65535
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case 32768:
                v30 := v2[-9223372036854775808]
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 7; fi31++ {
                v32 := []int{int(v1), fi31, genFn0(0, (^(v21)), 694), fi31}
                switch -328 {
                case 2147483647:
                    if true {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            var v34 uint16 = (v1 << ((196) & 63))
                            v5 = (v5 % ((v5) | 1))
                            if ((0.0 + 1.5) != 1.5) {
                                switch "a" {
                                case "hello":
                                    var v35 int64 = -9223372036854775808
                                    _ = v35
                                default:
                                    var v36 int = (fi31 & 2147483647)
                                    _ = v36
                                }
                                if (v4 || (!true)) {
                                    var v37 int64 = -733
                                    var v38 uint = v21
                                    _ = v37
                                    _ = v38
                                }
                                _ = len(v2)
                            }
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        var v39 string = "a"
                        _ = v39
                    }
                default:
                    for fi40 := 0; fi40 < 20; fi40++ {
                        v41 := GenS1{F0: 0, F1: (v5 & v5)}
                        _ = fi40
                        _ = v41
                    }
                }
                var v42 string = v6
                _ = fi31
                _ = v32
                _ = v42
            }
            switch (v19 + ((v19 + v6) + v19)) {
            case "\n\t":
                for fi43 := 0; fi43 < 9; fi43++ {
                    if (1 != -1) {
                        v19 = v6
                        if ((fi43 * fi43) != genFn0(491, v21, 594)) {
                            v1 = (v1 ^ v1)
                            v44 := GenS2{F0: v13, F1: 42, F2: 255}
                            var v45 int32 = -1
                            _ = v44
                            _ = v45
                        }
                    } else {
                        v13 += v13
                        v46 := GenS3{F0: v5}
                        _ = v46
                    }
                    if ((v21 ^ uint(v1)) == (v21 % (((v21 + v21)) | 1))) {
                        v19 = (v6 + v19)
                        var v47 int64 = 397
                        for fi48 := 0; fi48 < 18; fi48++ {
                            if ((v13 << ((v21) & 63)) <= 0) {
                                if v4 {
                                    var v49 float32 = v22
                                    _ = v49
                                } else {
                                    var v50 int64 = v47
                                    v22 = v22
                                    v3[4294967295] = uint32(v5)
                                    if ((v5 << ((v21) & 63)) <= (v5 - 42)) {
                                        var v51 uint16 = 1
                                        var v52 int = fi48
                                        v1 = (v51 + 1)
                                        _ = v51
                                        _ = v52
                                    } else {
                                        switch v6 {
                                        case "":
                                            v5 = (v5 - v5)
                                        case "a":
                                            if v4 {
                                                if (0.1 <= -2.25) {
                                                    _ = len(v2)
                                                    if (((v6 + (((v6 + v6) + v19) + (v19 + (v19 + (v6 + v6))))) + "foo bar") < v6) {
                                                        v53 := GenS4{F0: v13}
                                                        delete(v2, genFn0(2, v21, 4294967295))
                                                        v1 = (v1 / ((32767) | 1))
                                                        _ = v53
                                                    }
                                                    if v4 {
                                                        var v54 int64 = v50
                                                        fmt.Println(v54, v5, v1)
                                                        v19 = "hello"
                                                        _ = v54
                                                    }
                                                } else {
                                                    if (!((v50 - (v50 / ((v50) | 1))) >= (v50 / ((v50) | 1)))) {
                                                        var v55 string = v19
                                                        v1 += 0
                                                        _ = v55
                                                    }
                                                }
                                                v13 = v13
                                                switch (fi43 & 128) {
                                                case 32767:
                                                    v19 = "hello"
                                                case 0:
                                                    v21 = (v21 ^ uint(v1))
                                                case 127:
                                                    v4 = v4
                                                }
                                            } else {
                                                var v56 int8 = -1
                                                _ = len(v3)
                                                _ = v56
                                            }
                                        case "\n\t":
                                            v57 := GenS3{F0: 125}
                                            _ = v57
                                        default:
                                            var v58 bool = v4
                                            _ = v58
                                        }
                                    }
                                    fmt.Println(fi48, v5, v13)
                                    _ = v50
                                }
                                delete(v2, fi48)
                            } else {
                                v59 := GenS5{F0: 0, F1: 236, F2: v21}
                                _ = v59
                            }
                            var v60 uint32 = uint32(v21)
                            _ = fi48
                            _ = v60
                        }
                        _ = v47
                    }
                    for fi61 := 0; fi61 < 6; fi61++ {
                        v5 -= (v5 >> (((^(v21))) & 63))
                        switch genFn0(206, (v21 << ((v21) & 63)), 265) {
                        case 32768:
                            fmt.Println(v4, v4)
                        case 1:
                            v21 -= (v21 % ((18446744073709551615) | 1))
                        case -763:
                            v19 = "hello"
                        default:
                            switch (fi43 ^ fi61) {
                            case -479:
                                v13 = v13
                            default:
                                if (18446744073709551615 <= (v5 / (((v5 & v5)) | 1))) {
                                    switch (v6 + ((v6 + (v19 + v6)) + v6)) {
                                    case "a":
                                        v62 := GenS6{F0: v5, F1: "foo bar"}
                                        switch (fi61 ^ fi61) {
                                        case 787:
                                            v4 = (42 <= 42)
                                        case 42:
                                            var v63 int16 = (v13 - v13)
                                            _ = v63
                                        case 4294967295:
                                            _ = len(v2)
                                        }
                                        _ = v62
                                    case "\n\t":
                                        v5 = (v5 << (((v21 << (((^(v21))) & 63))) & 63))
                                    }
                                    switch (fi43 + fi43) {
                                    case 9223372036854775807:
                                        v19 = v19
                                    case 127:
                                        var v64 int = (fi43 & genFn0(1, (v21 / ((v21) | 1)), 204))
                                        _ = v64
                                    }
                                }
                            }
                        }
                        var v65 uint16 = (v1 - v1)
                        var v66 int64 = -1
                        _ = fi61
                        _ = v65
                        _ = v66
                    }
                    _ = fi43
                }
            case "":
                if v4 {
                    v6 = v19
                } else {
                    var v67 int16 = int16(v1)
                    var v68 uint32 = 2
                    v19 = v6
                    _ = v67
                    _ = v68
                }
            case "foo bar":
                for fi69 := 0; fi69 < 11; fi69++ {
                    var v70 uint16 = (v1 & 0)
                    var v71 int = (fi69 | (fi69 ^ (^(fi69))))
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            }
        }
    }
    var v72 string = v19
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
}
