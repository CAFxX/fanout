package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 uint16; F1 int8 }
func genFn0(p0 int8, p1 uint8) int16 {
    var v7 uint64 = 1
    p0 = 42
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: 0}
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 bool = (-1 != 920)
    switch 32767 {
    case -1:
        v12 := GenS1{F0: 32767}
        _ = v12
    default:
        var v13 int = -318
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 int) uint16 {
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint8 = 2
        _ = fi14
        _ = v15
    }
    p1 = -9223372036854775808
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{177, 2147483648}
    var v3 int = 1
    v4 := []int{(v3 ^ v3), (v3 >> ((334) & 63)), 32768}
    var v5 uint32 = 84
    var v6 int64 = 1
    if len(v4) > 0 {
        v4[0] = (v3 | 1)
    }
    v16 := GenS1{F0: 947}
    v17 := v1["foo bar"]
    fmt.Println(v17, v17, v17)
    for fi18 := 0; fi18 < 14; fi18++ {
        v6 += v6
        v3 += (v3 + (fi18 ^ 9223372036854775807))
        _ = fi18
    }
    if ((0 >= 255) && (0.0 == 0.0)) {
        _ = len(v1)
        for fi19 := 0; fi19 < 11; fi19++ {
            v6 += v6
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS2{F0: v17, F1: true}
                v22 := GenS3{F0: 648, F1: 0, F2: 0}
                fmt.Println(v3)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    } else {
        v3 = v3
        v5 = (v5 >> ((11) & 63))
    }
    v6 -= (v6 >> ((0) & 63))
    v5 += v5
    fmt.Println(v6)
    fmt.Println(v3, v3, v6)
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 float32 = 1.5
        for fi25 := 0; fi25 < 18; fi25++ {
            v6 = v6
            v5 = (v5 % ((996) | 1))
            var v26 int8 = int8(v6)
            _ = fi25
            _ = v26
        }
        var v27 uint16 = 255
        _ = fi23
        _ = v24
        _ = v27
    }
    var v28 int32 = 0
    v28 += 2147483647
    v5 = 226
    if true {
        var v29 int8 = 114
        for fi30 := 0; fi30 < 2; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                v32 := GenS4{F0: (v29 & (v29 >> ((uint(v5)) & 63))), F1: false}
                v3 += 42
                var v33 bool = false
                if (677 != 698) {
                    _ = len(v4)
                    v34 := GenS5{F0: (v28 << ((1) & 63))}
                    var v35 bool = (!(1 >= 2))
                    v5 = v5
                    if v33 {
                        var v36 uint8 = 1
                        _ = v36
                    } else {
                        v1[v17] = v17
                        v2 = append(v2, 0)
                        v6 = 1
                    }
                    v37 := GenS4{F0: v29, F1: v33}
                    if (!(fi31 < fi31)) {
                        v38 := GenS2{F0: (v17 + v17), F1: v35}
                        if v35 {
                            v3 = (-(v3))
                            var v39 float64 = 3.14159
                            if v33 {
                                delete(v1, ((("" + v17) + ((v17 + v17) + "")) + ""))
                            } else {
                                switch (v17 + v17) {
                                case "":
                                    v1["\n\t"] = ("foo bar" + v17)
                                default:
                                    v35 = (0 < v28)
                                }
                                v3 += fi31
                                if v35 {
                                    switch ("\n\t" + ("" + v17)) {
                                    case "a":
                                        v40 := GenS3{F0: 65535, F1: 32768, F2: -351}
                                        _ = v40
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = (v5 << ((754) & 63))
                                        }
                                    }
                                    if v33 {
                                        _ = len(v1)
                                        v39 += float64((float32(1e10) + 0.1))
                                        v4 = append(v4, v3)
                                    } else {
                                        var v41 int8 = 100
                                        _ = v41
                                    }
                                    var v42 int32 = v28
                                    fmt.Println(v3, v6, v6)
                                    var v43 int32 = 42
                                    _ = v42
                                    _ = v43
                                } else {
                                    v6 = v6
                                }
                            }
                            _ = v39
                        } else {
                            v44 := GenS1{F0: genFn0(100, 255)}
                            var v45 string = (v17 + "\n\t")
                            var v46 uint64 = 4294967295
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = v38
                    } else {
                        v33 = v33
                    }
                    var v47 uint32 = v5
                    _ = v34
                    _ = v35
                    _ = v37
                    _ = v47
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            v28 = 1
            if false {
                if true {
                    for fi48 := 0; fi48 < 14; fi48++ {
                        if (127 <= 1) {
                            _ = len(v1)
                            var v49 uint = 948
                            v50 := GenS6{F0: (("\n\t" + v17) + (v17 + v17)), F1: int(v6)}
                            _ = v49
                            _ = v50
                        } else {
                            v5 = v5
                        }
                        var v51 uint64 = 2
                        v17 = v17
                        _ = fi48
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 10; fi52++ {
                        var v53 string = ((v17 + "foo bar") + ((v17 + (v17 + ("foo bar" + v17))) + (v17 + ("hello" + (v17 + "\n\t")))))
                        var v54 int = -9223372036854775808
                        switch v53 {
                        case "":
                            v55 := GenS6{F0: v53, F1: 4294967295}
                            if true {
                                v17 = v53
                                var v56 int = 1
                                v17 = ("a" + "")
                                _ = v56
                            }
                            _ = v55
                        case "foo bar":
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(v28, v53)
                            } else {
                                fmt.Println(fi30, v29, v29)
                            }
                        }
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    v57 := GenS6{F0: v17, F1: (fi30 >> ((42) & 63))}
                    _ = v57
                } else {
                    switch 439 {
                    case 0:
                        v58 := GenS7{F0: (v6 % ((-79) | 1)), F1: 2, F2: (v5 ^ (v5 * v5))}
                        _ = v58
                    case 849:
                        v59 := GenS8{F0: 396, F1: 901, F2: 621}
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 11; fi60++ {
                        if true {
                            v5 += (^(v5))
                            var v61 uint = 65535
                            v6 = v6
                            _ = v61
                        } else {
                            fmt.Println(v5, v5, v6)
                            v62 := GenS9{F0: v28, F1: v17}
                            v17 = v17
                            _ = v62
                        }
                        switch v17 {
                        case "\n\t":
                            v63 := GenS10{F0: 255, F1: (v29 & (v29 / ((v29) | 1)))}
                            if ((v5 << ((18446744073709551615) & 63)) == 2) {
                                v2 = append(v2, v5)
                                v64 := v1[v17]
                                _ = v64
                            }
                            _ = v63
                        default:
                            var v65 int32 = v28
                            _ = v65
                        }
                        v29 += int8(v6)
                        _ = fi60
                    }
                }
            }
            _ = fi30
        }
        v17 = "a"
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
