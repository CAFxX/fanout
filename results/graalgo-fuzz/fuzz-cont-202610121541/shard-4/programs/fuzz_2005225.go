package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 int32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS7 struct { F0 bool; F1 bool; F2 int }
func genFn0(p0 uint32, p1 uint32) uint16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint32 = p1
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 += (p1 >> ((1) & 63))
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 int8) int64 {
    v10 := []int{128, int(p0)}
    var v11 string = "\n\t"
    v11 = v11
    v11 = "\n\t"
    fmt.Println(v11, p0)
    _ = p0
    _ = v10
    _ = v11
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2
    var v3 uint64 = 356
    var v4 int64 = 366
    var v5 int32 = -2147483648
    var v6 float32 = 1e10
    v5 += v5
    var v12 uint64 = v3
    switch int(v2) {
    case 1:
        v13 := GenS1{F0: -32768, F1: "\n\t", F2: genFn0((v2 / (((^(v2))) | 1)), 4294967295)}
        _ = v13
    case 11:
        switch -1 {
        case 2147483647:
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int64 = (-(v4))
                _ = fi14
                _ = v15
            }
        case 868:
            var v16 uint64 = (v12 % ((v3) | 1))
            _ = v16
        case 42:
            var v17 uint64 = v12
            _ = v17
        }
    }
    if (!(!(((1e10 / (((0.1 * 1.5) / ((((-2.25 + float64(0.0)) + 0.0) * float64((-2.25 + float32((1e10 - 0.1))))) / 0.1)) / ((float64(v6) * 1.5) - (1e10 / (1e10 * 1.5))))) - (0.0 * (1e10 * (float64(v6) - 0.0)))) != (1.5 - 1e10)))) {
        var v18 int16 = int16(v4)
        _ = v18
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v5, v6, v12)
    case "\n\t":
        switch 32767 {
        case 9223372036854775807:
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 float64 = 1e10
                _ = fi19
                _ = v20
            }
        case 32768:
            fmt.Println(v12, v12, v3)
        case 889:
            var v21 uint64 = (v3 / ((747) | 1))
            _ = v21
        }
    case "":
        var v22 int8 = 1
        _ = v22
    default:
        v5 -= -1
    }
    var v23 string = (("\n\t" + (("" + "a") + ("foo bar" + "\n\t"))) + ("\n\t" + (("\n\t" + ("\n\t" + "")) + ("foo bar" + (((("foo bar" + ("a" + "a")) + "") + "") + "a")))))
    v24 := []int{}
    fmt.Println(v2, v3, v23)
    var v25 uint64 = (v12 << ((918) & 63))
    var v26 float32 = ((v6 * float32(-2.25)) / v6)
    v23 = v23
    v26 -= -2.25
    for fi27 := 0; fi27 < 19; fi27++ {
        if true {
            v28 := GenS1{F0: 32767, F1: "foo bar", F2: 257}
            _ = v28
        } else {
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := v1[v2]
                v2 = 65535
                if (v12 != v3) {
                    switch (fi29 << ((4294967295) & 63)) {
                    case 128:
                        v2 = v2
                    case 4294967295:
                        var v31 bool = false
                        _ = v31
                    default:
                        v32 := GenS2{F0: 0}
                        _ = v32
                    }
                    switch v23 {
                    case "a":
                        if (919 <= 255) {
                            if v30 {
                                v33 := v1[v2]
                                if (-168 != 32767) {
                                    fmt.Println(v30, v23, v3)
                                    delete(v1, v2)
                                    var v34 float64 = 0.1
                                    _ = v34
                                } else {
                                    if false {
                                        v35 := v1[(v2 >> ((2) & 63))]
                                        if v33 {
                                            if len(v24) > 0 {
                                                v24[0] = int(v12)
                                            }
                                            v5 = (v5 | (v5 / ((v5) | 1)))
                                        } else {
                                            var v36 uint8 = 92
                                            _ = v36
                                        }
                                        if true {
                                            if len(v24) > 0 {
                                                v24[0] = (^(fi29))
                                            }
                                            v37 := GenS1{F0: int16(v5), F1: "foo bar", F2: genFn0(1, 361)}
                                            if v30 {
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                                switch 127 {
                                                case 360:
                                                    fmt.Println(v12)
                                                }
                                                var v38 int32 = v5
                                                v24 = append(v24, fi29)
                                                _ = v38
                                            } else {
                                                var v39 uint32 = v2
                                                v40 := GenS2{F0: -126}
                                                _ = v39
                                                _ = v40
                                            }
                                            v12 = v3
                                            _ = v37
                                        } else {
                                            var v41 uint64 = v3
                                            if false {
                                                if v35 {
                                                    if true {
                                                        v42 := GenS1{F0: int16(fi27), F1: "a", F2: 42}
                                                        v43 := v1[v2]
                                                        v44 := GenS2{F0: 100}
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    if len(v24) >= 2 {
                                                        v24 = v24[1:len(v24)]
                                                    }
                                                } else {
                                                    v41 = (^(v3))
                                                    v25 += (v12 >> ((750) & 63))
                                                    v45 := GenS3{F0: genFn0((^(v2)), (^(v2))), F1: v23, F2: -44}
                                                    var v46 int8 = 42
                                                    _ = v45
                                                    _ = v46
                                                }
                                                if (109 > 42) {
                                                    if true {
                                                        v47 := GenS4{F0: (v5 + v5), F1: (-(v5))}
                                                        var v48 string = v23
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    v49 := GenS4{F0: v5, F1: (v5 - (^(v5)))}
                                                    if (v30 || (v4 >= v4)) {
                                                        v6 = float32((float64(v26) * 1e10))
                                                        switch v23 {
                                                        case "foo bar":
                                                            var v50 int16 = 32767
                                                            _ = v50
                                                        case "a":
                                                            var v51 uint8 = 42
                                                            v25 += 18446744073709551615
                                                            fmt.Println(v6, v35)
                                                            _ = v51
                                                        }
                                                    } else {
                                                        if v33 {
                                                            v52 := GenS5{F0: fi27, F1: -556}
                                                            if (120 > 255) {
                                                                var v53 uint64 = (v12 ^ (v3 & v25))
                                                                v54 := v1[2147483648]
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            _ = v52
                                                        } else {
                                                            v55 := GenS1{F0: 477, F1: ("a" + v23), F2: 682}
                                                            v56 := GenS4{F0: v5, F1: v5}
                                                            if (v23 >= "foo bar") {
                                                                var v57 int16 = -1
                                                                _ = v57
                                                            }
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                    _ = v49
                                                }
                                                if v35 {
                                                    if len(v24) >= 2 {
                                                        v24 = v24[1:len(v24)]
                                                    }
                                                    switch (fi27 >> ((4294967295) & 63)) {
                                                    case 127:
                                                        v23 = ""
                                                    case -9223372036854775808:
                                                        var v58 float32 = 1e10
                                                        _ = v58
                                                    default:
                                                        v25 = (v12 / ((v3) | 1))
                                                    }
                                                } else {
                                                    fmt.Println(v26, v30, v41)
                                                }
                                            }
                                            v59 := GenS6{F0: v3, F1: v2, F2: -1}
                                            if (v3 == v3) {
                                                _ = len(v24)
                                            } else {
                                                v12 = v3
                                                fmt.Println(v33, v6, v25)
                                            }
                                            _ = v41
                                            _ = v59
                                        }
                                        _ = v35
                                    } else {
                                        v60 := GenS7{F0: ((v3 + (v3 ^ (v3 ^ 142))) <= 1), F1: false, F2: -570}
                                        var v61 bool = v30
                                        fmt.Println(v3, v23)
                                        _ = v60
                                        _ = v61
                                    }
                                    if ((v4 << ((18446744073709551615) & 63)) < genFn1(int8(v5))) {
                                        var v62 bool = v33
                                        _ = v62
                                    } else {
                                        delete(v1, v2)
                                    }
                                }
                                _ = v33
                            } else {
                                fmt.Println(v2, v3)
                                fmt.Println(v3, v26, v2)
                            }
                            fmt.Println(v12, v2)
                        } else {
                            if len(v24) > 0 {
                                v24[0] = int(v2)
                            }
                            var v63 bool = true
                            if len(v24) > 0 {
                                v24[0] = (fi29 / ((fi29) | 1))
                            }
                            _ = v63
                        }
                    case "":
                        v6 += v6
                    case "hello":
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                    default:
                        var v64 int32 = v5
                        _ = v64
                    }
                } else {
                    var v65 bool = (127 <= -1)
                    switch v23 {
                    case "hello":
                        v5 = (v5 % ((v5) | 1))
                    case "":
                        v65 = (((v3 != 4294967295) || (((1.5 / 1.0) + ((1e10 * (float32(0.0) / v6)) / 0.1)) < v26)) || v30)
                    case "\n\t":
                        v23 = (((v23 + v23) + v23) + "a")
                    }
                    _ = v65
                }
                var v66 uint16 = 588
                v67 := GenS6{F0: v12, F1: 1, F2: -1}
                _ = fi29
                _ = v30
                _ = v66
                _ = v67
            }
        }
        if (true && (0 == 660)) {
            v6 += v6
            fmt.Println(v2, v2, v12)
        }
        _ = fi27
    }
    if (!false) {
        var v68 int8 = 42
        _ = v68
    }
    v69 := GenS1{F0: int16(v4), F1: "", F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
