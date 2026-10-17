package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint64; F2 int }
type GenS7 struct { F0 int16 }
func genFn0(p0 bool) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float32 = 1.5
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    p0 = p0
    p0 = p0
    v10 := []int{-757, 323, 689}
    _ = p0
    _ = v10
    return 1
}
func genFn1() int16 {
    var v11 uint = 42
    var v12 uint8 = 127
    _ = v11
    _ = v12
    return 941
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 int16 = -32768
    var v6 uint64 = 18446744073709551615
    if len(v1) > 0 {
        v1[0] = 32768
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v5 = v5
        if ((v5 * (v5 | v5)) > -32768) {
            var v14 float32 = ((1e10 - float32(float64(3.14159))) / (1.5 / 1.0))
            if (1.5 <= 0.1) {
                var v15 uint = 520
                if len(v1) > 0 {
                    v1[0] = 32768
                }
                _ = v15
            }
            v6 += (v6 >> ((152) & 63))
            _ = v14
        } else {
            _ = len(v4)
            var v16 int = int(v6)
            _ = v16
        }
        if true {
            var v17 int16 = -944
            _ = v17
        } else {
            v5 = (-(v5))
            v18 := GenS1{F0: (-(v5)), F1: v6, F2: 1}
            _ = v18
        }
        _ = fi13
    }
    var v19 int64 = -242
    if (false && (255 > 744)) {
        v6 = v6
        v5 -= v5
        v20 := GenS1{F0: (v5 * genFn1()), F1: v6, F2: 2}
        _ = v20
    }
    v6 = v6
    var v21 float64 = 1.5
    v22 := GenS2{F0: 18446744073709551615}
    v6 = (v6 + (v6 << ((682) & 63)))
    if (-226 < -368) {
        var v23 float64 = 0.1
        _ = v23
    } else {
        v6 += 0
    }
    var v24 int16 = -353
    v5 = v24
    if true {
        var v25 float32 = 1.5
        _ = v25
    } else {
        var v26 bool = (false || ("a" > ("hello" + "")))
        var v27 uint8 = 69
        if false {
            if true {
                _ = len(v1)
            }
            switch 2147483647 {
            case 2147483647:
                fmt.Println(v19, v21, v21)
            case 32767:
                fmt.Println(v5, v5)
            case -9223372036854775808:
                v24 = genFn1()
            }
            switch -1 {
            case 4294967295:
                var v28 uint8 = 128
                _ = v28
            case 1:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 128:
                _ = len(v1)
            default:
                switch 0 {
                case 127:
                    for fi29 := 0; fi29 < 9; fi29++ {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            if (!(!v26)) {
                                var v31 uint32 = 2
                                _ = v31
                            }
                            if (!v26) {
                                v32 := GenS3{F0: 726, F1: 4294967295, F2: 32767}
                                v19 = (v19 ^ (-(v19)))
                                _ = v32
                            } else {
                                delete(v3, 42)
                            }
                            var v33 int8 = -128
                            _ = fi30
                            _ = v33
                        }
                        v34 := GenS4{F0: 2, F1: 2, F2: 42}
                        v35 := GenS5{F0: (fi29 / ((fi29) | 1))}
                        _ = fi29
                        _ = v34
                        _ = v35
                    }
                }
            }
        } else {
            v36 := GenS5{F0: -2147483648}
            var v37 uint = 65535
            if v26 {
                v1 = append(v1, 127)
                for fi38 := 0; fi38 < 20; fi38++ {
                    fmt.Println(v19)
                    for fi39 := 0; fi39 < 12; fi39++ {
                        v40 := GenS6{F0: (v5 & genFn1()), F1: (^(v6)), F2: (fi38 + -946)}
                        if ((v19 - 42) < 42) {
                            var v41 uint8 = (v27 - v27)
                            if (v26 || (v6 >= 2)) {
                                if v26 {
                                    v19 += (v19 & v19)
                                } else {
                                    fmt.Println(v27)
                                }
                                if false {
                                    v42 := GenS6{F0: (^(v24)), F1: (^(v6)), F2: (fi38 / ((int(v6)) | 1))}
                                    v1 = append(v1, 42)
                                    v19 = v19
                                    _ = v42
                                } else {
                                    if (false || (-2147483648 > 42)) {
                                        var v43 int16 = (-(v24))
                                        delete(v2, 42)
                                        _ = len(v1)
                                        _ = v43
                                    } else {
                                        fmt.Println(v37, fi38, v27)
                                        var v44 bool = false
                                        fmt.Println(fi38, v21, v41)
                                        _ = v44
                                    }
                                    if (42 != 255) {
                                        var v45 int = fi38
                                        if false {
                                            if v26 {
                                                v19 += v19
                                                v46 := GenS4{F0: 2147483647, F1: uint32(v6), F2: -275}
                                                switch ("hello" + "") {
                                                case "":
                                                    var v47 int64 = v19
                                                    _ = v47
                                                case "\n\t":
                                                    var v48 uint8 = v41
                                                    _ = v48
                                                case "hello":
                                                    v19 -= v19
                                                }
                                                _ = v46
                                            }
                                        }
                                        var v49 float32 = 1e10
                                        switch "hello" {
                                        case "":
                                            var v50 string = ""
                                            _ = v50
                                        default:
                                            if (!(3.14159 <= ((v49 * v49) + v49))) {
                                                var v51 int32 = 1
                                                switch "hello" {
                                                case "":
                                                    if v26 {
                                                        v52 := GenS4{F0: 42, F1: 182, F2: (v51 + v51)}
                                                        v53 := v3[2147483647]
                                                        var v54 float64 = v21
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        var v55 bool = (!(v51 > (v51 >> ((uint(v6)) & 63))))
                                                        v55 = v55
                                                        if (int8(v51) > -1) {
                                                            v45 -= (-(v45))
                                                            var v56 float64 = 1.5
                                                            _ = v56
                                                        }
                                                        _ = v55
                                                    }
                                                case "hello":
                                                    fmt.Println(v26)
                                                }
                                                var v57 int64 = v19
                                                _ = v51
                                                _ = v57
                                            }
                                        }
                                        var v58 uint8 = 2
                                        _ = v45
                                        _ = v49
                                        _ = v58
                                    }
                                }
                                v5 = 78
                                var v59 string = "hello"
                                _ = v59
                            } else {
                                v5 = 628
                            }
                            var v60 int64 = (v19 << ((307) & 63))
                            _ = v41
                            _ = v60
                        }
                        _ = fi39
                        _ = v40
                    }
                    _ = fi38
                }
                v37 += (v37 / ((v37) | 1))
            }
            _ = v36
            _ = v37
        }
        _ = v26
        _ = v27
    }
    if len(v1) > 0 {
        v1[0] = 127
    }
    if (574 != 905) {
        if (v24 > 522) {
            delete(v2, 1)
        } else {
            switch 385 {
            case -2147483648:
                v61 := GenS7{F0: v5}
                _ = v61
            default:
                var v62 uint32 = 568
                _ = v62
            }
        }
        var v63 int32 = 2147483647
        _ = v63
    }
    var v64 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
