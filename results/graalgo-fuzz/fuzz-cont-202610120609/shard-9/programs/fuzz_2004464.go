package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int64; F1 string; F2 string }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int32; F1 uint; F2 int }
func genFn0(p0 uint32) int {
    var v7 uint = 2
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{(fi8 | -9223372036854775808)}
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    v10 := []int{4294967295, int(p0), 128, 128}
    var v11 uint = 42
    if false {
        switch 42 {
        case 9223372036854775807:
            var v12 int = 975
            _ = v12
        case 32768:
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    fmt.Println(v11, p0, v11)
                    _ = fi14
                }
                _ = fi13
            }
        default:
            v11 = v7
        }
        v15 := GenS1{F0: 42, F1: 0, F2: p0}
        fmt.Println(p0, v11, v7)
        _ = v15
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return -258
}
func genFn1(p0 uint8, p1 int32, p2 int64) uint32 {
    v16 := GenS1{F0: 255, F1: 2, F2: 541}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v4 string = ""
    var v5 int16 = 1
    v6 := make(map[uint32]uint32)
    var v17 uint8 = 0
    v18 := GenS1{F0: 0, F1: (v17 | v17), F2: 0}
    if (v4 <= v4) {
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := GenS1{F0: 0, F1: v17, F2: genFn1((v17 >> ((1) & 63)), 2147483647, -9223372036854775808)}
            _ = fi19
            _ = v20
        }
        v5 -= v5
        v21 := GenS1{F0: 0, F1: v17, F2: 65535}
        _ = v21
    } else {
        switch (((v4 + v4) + v4) + (((v4 + (v4 + "hello")) + "") + v4)) {
        case "\n\t":
            v5 -= v5
        }
        v22 := GenS1{F0: 255, F1: v17, F2: 4294967295}
        _ = v22
    }
    if (32 == 42) {
        v17 += v17
        var v23 int = -307
        v5 -= 42
        _ = v23
    }
    var v24 int64 = 340
    fmt.Println(v4, v4, v17)
    v5 = -373
    if false {
        if (42 < 586) {
            v4 = v4
        } else {
            if true {
                for fi25 := 0; fi25 < 19; fi25++ {
                    v26 := []int{(fi25 & genFn0(42)), genFn0(2), fi25, (fi25 << ((65535) & 63))}
                    if true {
                        var v27 int32 = 816
                        for fi28 := 0; fi28 < 11; fi28++ {
                            v29 := GenS1{F0: 42, F1: 217, F2: 42}
                            var v30 int8 = 100
                            var v31 uint = 2
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
                var v32 uint32 = 1
                v33 := GenS2{F0: (v17 - (v17 - 42)), F1: 4294967295}
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                if (true && (1.5 < 3.14159)) {
                    v5 -= v5
                    if false {
                        v35 := GenS3{F0: v4, F1: (156 < 717)}
                        for fi36 := 0; fi36 < 6; fi36++ {
                            v37 := GenS3{F0: "hello", F1: false}
                            var v38 int32 = -2147483648
                            var v39 bool = (((-2.25 + (1e10 - 0.0)) - 1e10) <= 1e10)
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        delete(v2, 42)
                        _ = v35
                    } else {
                        for fi40 := 0; fi40 < 6; fi40++ {
                            v4 = ((v4 + v4) + ("" + v4))
                            var v41 uint = 255
                            v42 := GenS4{F0: 65535, F1: 32767, F2: 2147483647}
                            if false {
                                if true {
                                    fmt.Println(fi40, fi40, v24)
                                    v43 := GenS2{F0: v17, F1: 42}
                                    v44 := []int{-741, -245}
                                    var v45 uint16 = 2
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                            } else {
                                v46 := []int{0, (fi40 % (((fi40 - fi40)) | 1)), genFn0(1)}
                                if true {
                                    var v47 uint8 = 229
                                    v41 -= v41
                                    switch (fi34 ^ (fi40 / ((0) | 1))) {
                                    case -131:
                                        v4 = "\n\t"
                                    case 42:
                                        _ = len(v6)
                                        v6[2147483648] = 472
                                        var v48 bool = (2 <= 2147483648)
                                        _ = v48
                                    default:
                                        fmt.Println(v24)
                                    }
                                    _ = v47
                                }
                                var v49 uint16 = 32767
                                _ = v46
                                _ = v49
                            }
                            v5 += v5
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        v50 := v1[genFn0(0)]
                        var v51 uint64 = 42
                        _ = v50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 15; fi52++ {
                        if true {
                            fmt.Println(fi34, fi34)
                            var v53 int32 = 342
                            _ = v53
                        } else {
                            if false {
                                if false {
                                    if (((v4 + (v4 + v4)) + v4) < (v4 + "hello")) {
                                        if true {
                                            if (-128 != 42) {
                                                v54 := GenS5{F0: fi34, F1: 42}
                                                v55 := []int{genFn0(2147483647), fi52, -2147483648}
                                                v56 := GenS5{F0: fi34, F1: 18446744073709551615}
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            } else {
                                                v57 := GenS5{F0: genFn0(genFn1(2, -2147483648, v24)), F1: 1}
                                                _ = v57
                                            }
                                            var v58 uint64 = 4294967295
                                            v59 := GenS4{F0: 2, F1: 32767, F2: 42}
                                            _ = v58
                                            _ = v59
                                        } else {
                                            v60 := GenS6{F0: v24, F1: ((v4 + "") + v4), F2: "foo bar"}
                                            var v61 int32 = 2147483647
                                            _ = v60
                                            _ = v61
                                        }
                                        v4 = v4
                                    } else {
                                        switch (fi52 >> ((1) & 63)) {
                                        case 0:
                                            v62 := []int{genFn0(genFn1(v17, 42, 42)), (fi52 % ((42) | 1)), (fi34 & (^(fi52)))}
                                            _ = v62
                                        default:
                                            v17 += (v17 << ((700) & 63))
                                        }
                                        switch (-(fi34)) {
                                        case -226:
                                            fmt.Println(v24, v5)
                                        case -778:
                                            if ((true && (true && (0 == 9223372036854775807))) && (!(v17 >= 165))) {
                                                v63 := GenS7{F0: -1, F1: 127}
                                                var v64 uint32 = 708
                                                var v65 uint16 = 699
                                                v24 = 1
                                                if false {
                                                    v66 := GenS1{F0: (v65 / (((v65 * 32767)) | 1)), F1: v17, F2: 4294967295}
                                                    _ = v66
                                                }
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            }
                                        default:
                                            fmt.Println(v5)
                                        }
                                        v5 -= v5
                                        var v67 int16 = 32767
                                        _ = v67
                                    }
                                    var v68 bool = false
                                    if v68 {
                                        if v68 {
                                            var v69 int = (fi52 | fi34)
                                            _ = v69
                                        } else {
                                            v70 := GenS8{F0: 42, F1: 1, F2: fi52}
                                            _ = v70
                                        }
                                        var v71 int64 = 42
                                        delete(v6, 2147483648)
                                        var v72 bool = ((0 <= 1) && false)
                                        v68 = v68
                                        switch v4 {
                                        case "hello":
                                            if (genFn1(v17, -1, (v71 ^ v71)) <= 2147483648) {
                                                var v73 string = (v4 + ("a" + "foo bar"))
                                                _ = v73
                                            }
                                            v24 -= (v24 + (v71 << ((0) & 63)))
                                        case "\n\t":
                                            var v74 int64 = v24
                                            _ = v74
                                        default:
                                            if v72 {
                                                v75 := []int{(^(fi52))}
                                                _ = v75
                                            } else {
                                                var v76 int16 = v5
                                                _ = v76
                                            }
                                        }
                                        var v77 int16 = 580
                                        delete(v1, fi34)
                                        _ = v71
                                        _ = v72
                                        _ = v77
                                    }
                                    _ = v68
                                }
                                switch (v4 + v4) {
                                case "foo bar":
                                    var v78 float32 = ((1e10 + 0.0) / 1e10)
                                    _ = v78
                                default:
                                    var v79 int16 = (v5 | v5)
                                    _ = v79
                                }
                                switch (fi52 ^ fi52) {
                                case -2147483648:
                                    v24 -= v24
                                }
                            }
                            v5 = 0
                        }
                        _ = fi52
                    }
                }
                if false {
                    for fi80 := 0; fi80 < 8; fi80++ {
                        v4 = "\n\t"
                        _ = fi80
                    }
                    if (4294967295 == 353) {
                        fmt.Println(v24)
                        v1[fi34] = "\n\t"
                        var v81 int8 = -23
                        _ = v81
                    }
                    v17 = (v17 - v17)
                }
                if (v4 < (v4 + (v4 + ("hello" + ((v4 + "foo bar") + (v4 + v4)))))) {
                    v82 := GenS3{F0: (v4 + ""), F1: false}
                    _ = v82
                } else {
                    v24 = v24
                    v83 := v3["a"]
                    for fi84 := 0; fi84 < 20; fi84++ {
                        var v85 float32 = 1e10
                        _ = fi84
                        _ = v85
                    }
                    _ = v83
                }
                _ = fi34
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
