package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 int32 }
type GenS3 struct { F0 int8; F1 int16; F2 int }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint; F1 int8; F2 bool }
type GenS8 struct { F0 bool; F1 uint; F2 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int8; F1 int8 }
func genFn0(p0 int8) int64 {
    v7 := []int{}
    if (0 == -32768) {
        var v8 uint = 4294967295
        _ = v8
    }
    _ = p0
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := []string{"\n\t", ""}
    var v4 string = "foo bar"
    v5 := []int{}
    v6 := []int{42, -9223372036854775808, -756, 42}
    var v9 uint = 0
    v10 := GenS1{F0: 2, F1: 1}
    var v11 float64 = (float64((1.5 - (1.5 + 1e10))) + 0.1)
    v11 = v11
    v4 = "foo bar"
    fmt.Println(v4, v4)
    if (!(2147483647 > 127)) {
        var v12 float64 = v11
        var v13 uint16 = 0
        v14 := GenS1{F0: 0, F1: 18446744073709551615}
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 bool = false
        var v16 uint8 = 248
        _ = v15
        _ = v16
    }
    var v17 uint64 = 944
    v9 += (v9 % ((0) | 1))
    v11 = (v11 / v11)
    fmt.Println(v9, v17, v9)
    var v18 float32 = float32(3.14159)
    for fi19 := 0; fi19 < 20; fi19++ {
        v18 += v18
        fmt.Println(v9, v4, fi19)
        _ = fi19
    }
    var v20 uint = (^(v9))
    var v21 int64 = 3
    if true {
        if (65535 > 2147483648) {
            var v22 bool = true
            _ = v22
        } else {
            v4 = v4
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            if ((-(fi23)) >= fi23) {
                if (false && ((!(v11 == v11)) && (2147483647 <= 2147483647))) {
                    fmt.Println(v11, v4)
                    v24 := GenS2{F0: -128, F1: 0, F2: 42}
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v4 = ""
                        var v26 bool = (0 < -32768)
                        var v27 float32 = v18
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                _ = len(v5)
            } else {
                v28 := GenS1{F0: 200, F1: v17}
                v29 := GenS3{F0: 127, F1: int16(v9), F2: (fi23 | fi23)}
                var v30 string = (v4 + ((("hello" + "") + (v4 + (("a" + "\n\t") + "hello"))) + "\n\t"))
                _ = v28
                _ = v29
                _ = v30
            }
            var v31 float64 = 3.14159
            v32 := v1[2]
            _ = fi23
            _ = v31
            _ = v32
        }
        var v33 string = "foo bar"
        _ = v33
    } else {
        v34 := v1[151]
        v17 = 42
        for fi35 := 0; fi35 < 3; fi35++ {
            for fi36 := 0; fi36 < 3; fi36++ {
                if (v21 > v21) {
                    var v37 uint32 = 2
                    if true {
                        v18 += 1e10
                        var v38 uint32 = (v34 ^ (v37 ^ v37))
                        _ = v38
                    }
                    _ = v37
                } else {
                    if false {
                        fmt.Println(fi36, v11)
                        var v39 float64 = v11
                        v40 := GenS4{F0: (fi36 % (((fi36 + (fi36 % ((fi35) | 1)))) | 1)), F1: 0, F2: true}
                        _ = v39
                        _ = v40
                    } else {
                        if true {
                            var v41 float64 = 0.1
                            var v42 int16 = 1
                            v43 := GenS5{F0: v42, F1: (v34 | v34), F2: 2}
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                    }
                    v44 := GenS1{F0: (^(v17)), F1: 55}
                    var v45 float64 = (3.14159 * float64(1e10))
                    v46 := GenS6{F0: (fi35 << ((v20) & 63)), F1: int8(v34)}
                    v9 += v20
                    if (!true) {
                        var v47 uint64 = 2
                        v1[v34] = v34
                        var v48 uint16 = 42
                        v49 := GenS3{F0: 100, F1: 32767, F2: (fi36 >> (((^(v20))) & 63))}
                        _ = v47
                        _ = v48
                        _ = v49
                    } else {
                        var v50 int32 = 2147483647
                        var v51 float32 = ((3.14159 / v18) - 0.1)
                        switch (fi36 * fi35) {
                        case 1:
                            v6 = append(v6, (fi36 + 128))
                        case -1:
                            if (!(false || false)) {
                                var v52 uint16 = 65535
                                v53 := v1[v34]
                                var v54 int16 = 42
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        default:
                            var v55 int = (fi35 % ((fi35) | 1))
                            _ = v55
                        }
                        _ = v50
                        _ = v51
                    }
                    _ = v44
                    _ = v45
                    _ = v46
                }
                if false {
                    fmt.Println(v21, fi36)
                    v56 := v1[(v34 * (^(v34)))]
                    var v57 float32 = v18
                    _ = v56
                    _ = v57
                } else {
                    if false {
                        if ((!(898 <= 255)) && true) {
                            v58 := GenS7{F0: v9, F1: 127, F2: false}
                            _ = v58
                        } else {
                            var v59 uint32 = 762
                            v60 := GenS1{F0: (v17 * (v17 * v17)), F1: v17}
                            if false {
                                if ((((v4 + ("hello" + v4)) + (v4 + (v4 + v4))) + "hello") != "\n\t") {
                                    if false {
                                        switch int(v21) {
                                        case -51:
                                            var v61 bool = (v21 != -1)
                                            _ = v61
                                        case 414:
                                            v17 = 2
                                        case -9223372036854775808:
                                            v17 += v17
                                        }
                                        fmt.Println(v11, fi35)
                                        var v62 float32 = v18
                                        _ = v62
                                    }
                                }
                                if true {
                                    if false {
                                        v4 = v4
                                        if true {
                                            v4 = v4
                                            switch v4 {
                                            case "a":
                                                v63 := GenS8{F0: false, F1: (v20 % ((uint(v17)) | 1)), F2: uint64(v20)}
                                                _ = v63
                                            case "foo bar":
                                                var v64 uint16 = 32767
                                                _ = v64
                                            default:
                                                v21 = v21
                                                switch (fi36 >> ((65535) & 63)) {
                                                case -771:
                                                    v11 -= 3.14159
                                                case -1:
                                                    fmt.Println(v59, v21, v4)
                                                default:
                                                    if (-372 > -2147483648) {
                                                        v21 += 544
                                                        if (0 <= 42) {
                                                            switch fi35 {
                                                            case -1:
                                                                var v65 string = v4
                                                                switch v65 {
                                                                case "hello":
                                                                    fmt.Println(v21)
                                                                }
                                                                v34 = v34
                                                                _ = v65
                                                            }
                                                            v66 := GenS9{F0: 127}
                                                            var v67 int8 = int8(v21)
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        v9 += (v9 * (v20 * 0))
                                                    } else {
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        if true {
                                            v68 := GenS10{F0: int8(v17), F1: fi36}
                                            _ = v68
                                        } else {
                                            if true {
                                                switch v4 {
                                                case "a":
                                                    v3 = append(v3, v4)
                                                }
                                                v69 := GenS11{F0: (v17 - uint64(v9))}
                                                if (!false) {
                                                    var v70 uint16 = 65535
                                                    v9 = (v20 * v20)
                                                    _ = v70
                                                }
                                                if false {
                                                    v71 := GenS12{F0: -1, F1: int8(v9)}
                                                    fmt.Println(fi35, v17)
                                                    _ = v71
                                                }
                                                _ = v69
                                            }
                                        }
                                        delete(v1, v34)
                                    }
                                    var v72 uint64 = (v17 * (v17 * 135))
                                    var v73 int32 = -1
                                    _ = v72
                                    _ = v73
                                }
                            }
                            _ = v59
                            _ = v60
                        }
                    } else {
                        v4 = (v4 + (v4 + (v4 + v4)))
                        fmt.Println(v18, fi36)
                        v34 -= (v34 ^ 0)
                        v74 := GenS3{F0: -96, F1: int16(fi36), F2: (fi36 << ((uint(v17)) & 63))}
                        _ = v74
                    }
                }
                v17 = v17
                _ = fi36
            }
            v4 = v4
            _ = fi35
        }
        _ = v34
    }
    v75 := GenS1{F0: 2, F1: (v17 + v17)}
    var v76 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v76:", v76)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v9
}
