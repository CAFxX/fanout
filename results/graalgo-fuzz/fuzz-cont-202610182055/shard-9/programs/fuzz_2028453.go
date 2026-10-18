package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int8; F1 int; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint64; F2 bool }
type GenS6 struct { F0 uint; F1 int32; F2 uint32 }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint32) int32 {
    var v7 float64 = 1.5
    v7 = (-2.25 / float64(1.5))
    var v8 int8 = -128
    fmt.Println(v7, v8, v7)
    var v9 int16 = 93
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    v3 := []int{-2147483648, 42}
    v4 := make(map[uint32]bool)
    var v5 bool = false
    var v6 int32 = 42
    v1[545] = ("hello" + "foo bar")
    v5 = ((!(1 >= 0)) && v5)
    var v10 float32 = 1e10
    switch -583 {
    case 1:
        switch -851 {
        case 331:
            v10 = 1.5
        }
    case 143:
        for fi11 := 0; fi11 < 6; fi11++ {
            fmt.Println(fi11, v5, v10)
            _ = fi11
        }
    case 631:
        if v5 {
            var v12 int8 = -128
            _ = v12
        }
    }
    if false {
        var v13 string = "hello"
        _ = v13
    }
    v10 += v10
    delete(v4, 915)
    v6 = (v6 - v6)
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int8 = 127
        _ = fi14
        _ = v15
    }
    var v16 bool = (409 != 330)
    for fi17 := 0; fi17 < 7; fi17++ {
        v6 = v6
        v18 := GenS1{F0: 750}
        _ = fi17
        _ = v18
    }
    v10 = 1e10
    var v19 uint = 1
    var v20 uint16 = 0
    var v21 float64 = 3.14159
    if true {
        if ("hello" >= "a") {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v22 := GenS2{F0: 32767, F1: "\n\t"}
            _ = v22
        } else {
            var v23 uint16 = 983
            _ = v23
        }
        if false {
            var v24 int64 = 696
            if v5 {
                fmt.Println(v19)
                if (777 == 4294967295) {
                    switch "" {
                    case "\n\t":
                        var v25 uint = (v19 | (v19 << (((v19 - v19)) & 63)))
                        _ = v25
                    case "foo bar":
                        v20 = 65535
                    case "hello":
                        if v5 {
                            v3 = append(v3, -113)
                            v26 := v1[2147483648]
                            _ = v26
                        }
                    }
                    if v5 {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            v28 := v4[413]
                            _ = fi27
                            _ = v28
                        }
                        var v29 uint16 = v20
                        var v30 uint16 = v20
                        _ = v29
                        _ = v30
                    }
                } else {
                    switch -2147483648 {
                    case 32767:
                        v31 := GenS2{F0: 776, F1: (("a" + "hello") + ("\n\t" + "a"))}
                        _ = v31
                    case 42:
                        fmt.Println(v24, v16, v16)
                    }
                    var v32 int64 = -1
                    var v33 int = 954
                    _ = v32
                    _ = v33
                }
            } else {
                for fi34 := 0; fi34 < 13; fi34++ {
                    v35 := GenS3{F0: 1, F1: 1, F2: -1}
                    if v5 {
                        var v36 uint8 = 142
                        if (int32(fi34) != (v6 * (v6 - v6))) {
                            v36 = 4
                            _ = len(v4)
                        }
                        if v16 {
                            var v37 int16 = -1
                            _ = v37
                        } else {
                            for fi38 := 0; fi38 < 5; fi38++ {
                                fmt.Println(v6, v6, v16)
                                _ = fi38
                            }
                            for fi39 := 0; fi39 < 19; fi39++ {
                                var v40 string = ("\n\t" + "")
                                if ((0.0 > 0.0) || (1e10 >= v21)) {
                                    if v16 {
                                        var v41 uint = (v19 + v19)
                                        v36 += v36
                                        _ = v41
                                    } else {
                                        v42 := GenS3{F0: -63, F1: 1, F2: -1}
                                        _ = v42
                                    }
                                    var v43 int32 = 42
                                    var v44 uint16 = (v20 % (((v20 ^ (v20 | (v20 % (((^(v20))) | 1))))) | 1))
                                    _ = v43
                                    _ = v44
                                } else {
                                    if (true || true) {
                                        v4[1] = false
                                        v45 := GenS4{F0: uint8(v19)}
                                        _ = v45
                                    } else {
                                        v20 = (v20 / ((v20) | 1))
                                        var v46 int = (fi39 % ((-1) | 1))
                                        var v47 uint16 = (v20 ^ (v20 | (v20 - 0)))
                                        v48 := GenS5{F0: (v47 | (v20 - uint16(v19))), F1: 0, F2: (0 != 32767)}
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                    if true {
                                        v40 = "a"
                                        if ((v24 << (((v19 * (v19 | (v19 << ((v19) & 63))))) & 63)) != -1) {
                                            if (v24 > 644) {
                                                var v49 uint32 = 0
                                                _ = v49
                                            } else {
                                                v40 = (v40 + (v40 + v40))
                                            }
                                            v50 := GenS2{F0: (fi34 + 330), F1: (v40 + v40)}
                                            _ = v50
                                        } else {
                                            var v51 uint64 = 4294967295
                                            v52 := GenS3{F0: -1, F1: fi34, F2: 32767}
                                            var v53 int16 = 745
                                            if (uint(v51) != 1) {
                                                switch -815 {
                                                case -2147483648:
                                                    v6 += -2147483648
                                                case -282:
                                                    if true {
                                                        _ = len(v2)
                                                        var v54 uint32 = 112
                                                        _ = v54
                                                    } else {
                                                        fmt.Println(fi34, v36, v24)
                                                        switch (("\n\t" + v40) + v40) {
                                                        case "hello":
                                                            v55 := GenS1{F0: v19}
                                                            _ = v55
                                                        default:
                                                            if (-385 < 42) {
                                                                fmt.Println(v20, v53)
                                                                var v56 int64 = (v24 % ((v24) | 1))
                                                                v5 = true
                                                                _ = v56
                                                            } else {
                                                                v16 = (v36 > 2)
                                                                var v57 uint = v19
                                                                var v58 uint8 = v36
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                        }
                                                        v59 := GenS1{F0: v19}
                                                        _ = v59
                                                    }
                                                }
                                            } else {
                                                if ((v24 < -1) || (!((v6 << (((v19 * v19)) & 63)) > (v6 - v6)))) {
                                                    fmt.Println(v36)
                                                    var v60 float64 = 3.14159
                                                    _ = v60
                                                }
                                            }
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                        var v61 uint8 = 4
                                        _ = v61
                                    } else {
                                        var v62 int16 = 264
                                        var v63 uint = (v19 * (v19 + (^(v19))))
                                        _ = v62
                                        _ = v63
                                    }
                                    var v64 int32 = -145
                                    _ = v64
                                }
                                var v65 int = (fi34 << ((281) & 63))
                                _ = fi39
                                _ = v40
                                _ = v65
                            }
                        }
                        _ = v36
                    } else {
                        v24 = int64(fi34)
                    }
                    _ = fi34
                    _ = v35
                }
                v66 := v4[288]
                _ = v66
            }
            var v67 int16 = -1
            _ = v24
            _ = v67
        }
    } else {
        v68 := GenS4{F0: 138}
        if ((float64(1e10) <= v21) || (-1 > -9223372036854775808)) {
            v16 = (!(11 < 255))
        } else {
            var v69 uint64 = 4294967295
            if (0 >= -9223372036854775808) {
                if (2147483648 <= 2147483647) {
                    for fi70 := 0; fi70 < 14; fi70++ {
                        var v71 float64 = v21
                        fmt.Println(v19)
                        _ = fi70
                        _ = v71
                    }
                } else {
                    var v72 uint32 = 42
                    if (-357 != -2147483648) {
                        for fi73 := 0; fi73 < 17; fi73++ {
                            if v16 {
                                for fi74 := 0; fi74 < 10; fi74++ {
                                    v6 = 184
                                    _ = fi74
                                }
                                v75 := GenS3{F0: 22, F1: (fi73 ^ 42), F2: 119}
                                for fi76 := 0; fi76 < 10; fi76++ {
                                    var v77 uint16 = v20
                                    _ = fi76
                                    _ = v77
                                }
                                _ = v75
                            } else {
                                for fi78 := 0; fi78 < 11; fi78++ {
                                    switch (fi73 >> ((v19) & 63)) {
                                    case 9223372036854775807:
                                        v79 := GenS6{F0: (v19 * 2), F1: genFn0(2147483648), F2: 4294967295}
                                        _ = v79
                                    }
                                    v20 = 65535
                                    var v80 uint16 = v20
                                    _ = fi78
                                    _ = v80
                                }
                            }
                            _ = fi73
                        }
                        fmt.Println(v19, v6)
                        v81 := GenS5{F0: v20, F1: (^(v69)), F2: false}
                        _ = v81
                    } else {
                        var v82 int8 = 1
                        for fi83 := 0; fi83 < 11; fi83++ {
                            var v84 float32 = v10
                            v85 := GenS7{F0: 0}
                            _ = fi83
                            _ = v84
                            _ = v85
                        }
                        _ = v82
                    }
                    v5 = (false && v16)
                    _ = v72
                }
                v20 += (v20 << ((v19) & 63))
                var v86 uint16 = uint16(v69)
                _ = v86
            }
            _ = v69
        }
        if v16 {
            v10 = 1e10
            delete(v1, 65535)
        }
        _ = v68
    }
    if v16 {
        for fi87 := 0; fi87 < 19; fi87++ {
            v88 := GenS7{F0: 1}
            _ = fi87
            _ = v88
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
