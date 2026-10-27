package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 int64; F2 string }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint; F1 uint64; F2 int32 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 int8 }
func genFn0() int8 {
    var v7 float64 = 1.5
    v8 := GenS1{F0: 808, F1: 2147483647, F2: -2147483648}
    _ = v7
    _ = v8
    return -128
}
func genFn1() int64 {
    if true {
        var v9 string = "\n\t"
        fmt.Println(v9)
        _ = v9
    }
    v10 := GenS1{F0: 909, F1: 2147483647, F2: -949}
    switch "\n\t" {
    case "a":
        var v11 int16 = -55
        _ = v11
    case "":
        v12 := GenS1{F0: 1, F1: -1000, F2: 0}
        _ = v12
    case "hello":
        if (859 <= 42) {
            var v13 int64 = 1
            _ = v13
        }
    default:
        var v14 float32 = (0.1 * (3.14159 + 1e10))
        _ = v14
    }
    v15 := GenS1{F0: 1, F1: -1, F2: -1}
    if ("\n\t" < "a") {
        v16 := GenS2{F0: (!true), F1: 0, F2: ("foo bar" + "a")}
        var v17 int64 = -948
        if false {
            var v18 uint64 = 4294967295
            v19 := GenS2{F0: ((false && true) || true), F1: -9223372036854775808, F2: (("foo bar" + "\n\t") + ("a" + ("hello" + "")))}
            _ = v18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    _ = v10
    _ = v15
    return -1
}

func main() {
    var v1 int8 = -128
    var v2 int = -648
    v3 := make(map[int]uint32)
    var v4 int32 = 42
    var v5 int32 = v4
    v6 := []int{(v2 - -2147483648)}
    switch "" {
    case "hello":
        v20 := GenS2{F0: false, F1: -947, F2: "foo bar"}
        _ = v20
    }
    var v21 uint32 = 912
    for fi22 := 0; fi22 < 17; fi22++ {
        delete(v3, (v2 + 651))
        v23 := GenS2{F0: true, F1: 0, F2: "foo bar"}
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: "hello", F1: -1}
    var v25 uint32 = v21
    v26 := GenS1{F0: -616, F1: (v4 - 42), F2: v5}
    var v27 int = (v2 % ((v2) | 1))
    v3[v2] = v25
    v28 := GenS4{F0: v27, F1: "a"}
    v29 := GenS3{F0: "hello", F1: (v27 % ((433) | 1))}
    switch v2 {
    case 14:
        v1 = 118
    case 32767:
        fmt.Println(v21, v27)
    }
    fmt.Println(v25)
    _ = len(v6)
    if (255 >= 65535) {
        for fi30 := 0; fi30 < 7; fi30++ {
            if (int(v4) != (v27 * 1)) {
                var v31 int = (v27 << ((0) & 63))
                v21 = 65535
                v1 = v1
                _ = v31
            }
            if (!(!true)) {
                switch (fi30 >> ((42) & 63)) {
                case -587:
                    v5 -= v4
                case -2147483648:
                    if (1 > 32767) {
                        fmt.Println(v2, fi30)
                        v6 = append(v6, (v2 + (v27 >> ((395) & 63))))
                    } else {
                        var v32 uint8 = 127
                        _ = v32
                    }
                case 34:
                    v1 = v1
                }
            }
            var v33 uint8 = 144
            _ = fi30
            _ = v33
        }
    } else {
        if (v5 != (v4 & -480)) {
            switch "hello" {
            case "foo bar":
                var v34 int32 = 42
                _ = v34
            case "\n\t":
                v1 = (v1 ^ (v1 / ((0) | 1)))
            case "":
                if true {
                    v6 = append(v6, v2)
                    v27 += (v27 >> ((451) & 63))
                    v35 := GenS4{F0: v2, F1: "foo bar"}
                    _ = v35
                }
            default:
                if true {
                    var v36 string = "\n\t"
                    v5 -= v5
                    v37 := GenS5{F0: 0, F1: (v21 >> ((42) & 63))}
                    _ = v36
                    _ = v37
                } else {
                    if (-222 < 42) {
                        delete(v3, v27)
                        fmt.Println(v5)
                    }
                }
            }
            if (!((true && (0 != 42)) && false)) {
                v1 = (v1 / ((v1) | 1))
            } else {
                var v38 float32 = 3.14159
                for fi39 := 0; fi39 < 16; fi39++ {
                    switch fi39 {
                    case 9223372036854775807:
                        if (!false) {
                            v27 = -9223372036854775808
                            v40 := GenS6{F0: 2, F1: 42, F2: 344}
                            _ = v40
                        } else {
                            if true {
                                v2 += 2147483647
                                var v41 float32 = v38
                                for fi42 := 0; fi42 < 7; fi42++ {
                                    var v43 int32 = v4
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v41
                            }
                            if (108 != 127) {
                                v5 = (v5 / (((v5 / (((v4 | (v4 | (v5 / ((2147483647) | 1))))) | 1))) | 1))
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v44 float32 = 0.1
                                _ = v44
                            }
                            v38 += v38
                        }
                    case 4294967295:
                        for fi45 := 0; fi45 < 18; fi45++ {
                            var v46 uint64 = 1
                            _ = fi45
                            _ = v46
                        }
                    case -1:
                        fmt.Println(v1, fi39)
                    default:
                        if (!true) {
                            if (1 == 482) {
                                var v47 uint = uint(v21)
                                _ = v47
                            }
                            if false {
                                _ = len(v3)
                                if false {
                                    var v48 uint32 = v21
                                    v38 += v38
                                    if (genFn0() > -128) {
                                        v49 := GenS7{F0: v5, F1: (2 < 175)}
                                        _ = v49
                                    }
                                    _ = v48
                                } else {
                                    if true {
                                        v4 -= v4
                                        v50 := v3[v27]
                                        var v51 uint16 = 217
                                        _ = v50
                                        _ = v51
                                    } else {
                                        if false {
                                            v52 := GenS7{F0: v4, F1: (!false)}
                                            for fi53 := 0; fi53 < 17; fi53++ {
                                                v6 = append(v6, -1)
                                                if false {
                                                    var v54 uint64 = 492
                                                    switch v27 {
                                                    case -2147483648:
                                                        fmt.Println(v1)
                                                    case 127:
                                                        fmt.Println(v38)
                                                    default:
                                                        var v55 int8 = v1
                                                        _ = v55
                                                    }
                                                    _ = v54
                                                } else {
                                                    v6 = append(v6, fi39)
                                                    var v56 uint8 = 9
                                                    v57 := GenS2{F0: false, F1: 9223372036854775807, F2: "foo bar"}
                                                    var v58 bool = ((fi39 ^ (fi53 >> ((610) & 63))) >= v2)
                                                    var v59 int8 = 42
                                                    v60 := GenS8{F0: (v1 % (((v1 * genFn0())) | 1))}
                                                    v25 += 4294967295
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                }
                                                var v61 uint64 = 4294967295
                                                _ = fi53
                                                _ = v61
                                            }
                                            v62 := GenS1{F0: -32768, F1: (v5 >> ((255) & 63)), F2: v4}
                                            _ = v52
                                            _ = v62
                                        }
                                        fmt.Println(v38, v5, v38)
                                        v4 = (-(v4))
                                    }
                                }
                            }
                        }
                    }
                    var v63 uint32 = 406
                    fmt.Println(v25, v1, v25)
                    _ = fi39
                    _ = v63
                }
                fmt.Println(v38, v25)
                _ = v38
            }
        } else {
            _ = len(v3)
            var v64 uint32 = v21
            var v65 uint16 = 505
            _ = v64
            _ = v65
        }
        v21 += (v21 - v25)
    }
    var v66 float64 = (3.14159 - 3.14159)
    v25 += (^(v25))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
