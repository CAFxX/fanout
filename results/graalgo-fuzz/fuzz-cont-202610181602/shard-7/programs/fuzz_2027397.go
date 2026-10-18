package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int16; F1 int; F2 int32 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64; F1 uint16; F2 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int64; F1 string }
func genFn0(p0 uint16) uint64 {
    p0 = (p0 * p0)
    if (2 != 357) {
        for fi4 := 0; fi4 < 10; fi4++ {
            v5 := GenS1{F0: 848, F1: 989, F2: 1}
            var v6 uint64 = 1
            _ = fi4
            _ = v5
            _ = v6
        }
    } else {
        v7 := []int{-1, 4294967295, 999, 803}
        _ = v7
    }
    _ = p0
    return 42
}
func genFn1() string {
    var v8 float64 = 0.0
    if true {
        fmt.Println(v8)
    }
    _ = v8
    return ("" + "hello")
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 int32 = -547
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float32 = -2.25
        var v11 int = fi9
        fmt.Println(fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = len(v2)
    if (32767 >= 42) {
        v12 := GenS2{F0: v1, F1: (v3 / (((v3 >> ((2) & 63))) | 1))}
        for fi13 := 0; fi13 < 5; fi13++ {
            if false {
                var v14 int = (fi13 & int(v3))
                _ = v14
            } else {
                var v15 int64 = 1
                v16 := GenS1{F0: 566, F1: 0, F2: v3}
                var v17 uint8 = 2
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := []int{128, 32767}
            _ = fi13
            _ = v18
        }
        _ = v12
    }
    if (2 != 549) {
        fmt.Println(v3)
        v19 := GenS1{F0: 42, F1: 571, F2: -2147483648}
        var v20 int16 = 32767
        _ = v19
        _ = v20
    }
    v21 := []int{424, 127, 433, 127}
    fmt.Println(v3, v3)
    var v22 uint16 = 32767
    if (false && false) {
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 float64 = 0.1
            switch (fi23 + 0) {
            case 4294967295:
                v24 = 3.14159
            case 0:
                var v25 int16 = 1
                _ = v25
            }
            fmt.Println(v24, v24)
            _ = fi23
            _ = v24
        }
    } else {
        if (1 > 2147483647) {
            var v26 bool = false
            _ = v26
        } else {
            v27 := GenS2{F0: true, F1: (-(v3))}
            _ = v27
        }
        var v28 int = 32768
        v22 -= (v22 % (((^(v22))) | 1))
        _ = v28
    }
    if (127 != 40) {
        v29 := GenS3{F0: 687, F1: 32767, F2: v3}
        var v30 uint64 = 373
        _ = v29
        _ = v30
    } else {
        v3 = (v3 & (^(v3)))
        v3 = v3
        v1 = v1
    }
    if (!false) {
        switch ("" + "foo bar") {
        case "a":
            for fi31 := 0; fi31 < 1; fi31++ {
                v1 = ((108 >= 128) && v1)
                _ = fi31
            }
        case "foo bar":
            var v32 uint = 42
            _ = v32
        case "\n\t":
            fmt.Println(v1)
        default:
            v1 = true
        }
        v3 -= (v3 % ((-822) | 1))
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        if false {
            for fi34 := 0; fi34 < 12; fi34++ {
                var v35 bool = (2 > 21)
                if v35 {
                    fmt.Println(v3, v1)
                }
                if v1 {
                    var v36 uint = 42
                    _ = v36
                } else {
                    switch (fi34 ^ (-(fi33))) {
                    case 4294967295:
                        var v37 uint16 = (^(v22))
                        _ = v37
                    }
                }
                _ = fi34
                _ = v35
            }
        } else {
            fmt.Println(fi33, fi33, v1)
            v38 := GenS4{F0: 42, F1: 42, F2: true}
            _ = v38
        }
        if (0 == -128) {
            if false {
                var v39 uint = 42
                v39 += (^(v39))
                _ = v39
            }
            v40 := GenS5{F0: (v3 >> ((42) & 63))}
            for fi41 := 0; fi41 < 5; fi41++ {
                if true {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    fmt.Println(fi33, v22, v3)
                    if true {
                        fmt.Println(v1, fi33)
                        v42 := GenS3{F0: 32767, F1: fi41, F2: int32(v22)}
                        var v43 float64 = (3.14159 + 1.5)
                        var v44 string = ""
                        v2[(fi41 ^ (fi33 - (fi41 - -9223372036854775808)))] = fi33
                        _ = v42
                        _ = v43
                        _ = v44
                    } else {
                        v1 = (42 != 42)
                        v2[fi33] = fi41
                    }
                } else {
                    var v45 uint8 = 1
                    var v46 uint64 = 42
                    var v47 bool = false
                    _ = v45
                    _ = v46
                    _ = v47
                }
                if false {
                    v48 := GenS6{F0: v1, F1: 149}
                    var v49 uint8 = 255
                    _ = v48
                    _ = v49
                } else {
                    v50 := GenS7{F0: ""}
                    v1 = v1
                    if (fi41 <= (fi33 ^ 9223372036854775807)) {
                        switch -928 {
                        case 1:
                            _ = len(v21)
                        case -9223372036854775808:
                            switch fi41 {
                            case 4294967295:
                                if ((v1 && v1) || v1) {
                                    v22 = v22
                                } else {
                                    switch int(v22) {
                                    case 42:
                                        v51 := GenS6{F0: v1, F1: 18446744073709551615}
                                        _ = v51
                                    case 4294967295:
                                        var v52 uint8 = 0
                                        _ = v52
                                    default:
                                        switch "hello" {
                                        case "\n\t":
                                            var v53 int32 = (v3 + (v3 / (((v3 & v3)) | 1)))
                                            _ = v53
                                        case "":
                                            v21 = append(v21, 466)
                                        default:
                                            v3 -= v3
                                            if v1 {
                                                if v1 {
                                                    if (((v22 == 1) && v1) || v1) {
                                                        v54 := v2[(fi41 | -2147483648)]
                                                        v3 = (v3 & (^(v3)))
                                                        var v55 uint64 = 2
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        fmt.Println(v1, fi41)
                                                    }
                                                    v56 := GenS8{F0: fi33}
                                                    if ((false && ("" <= ("a" + "\n\t"))) || (345 != 1)) {
                                                        var v57 float32 = 3.14159
                                                        if (!(v3 > (^(v3)))) {
                                                            fmt.Println(v22, v1, v57)
                                                            switch "" {
                                                            case "\n\t":
                                                                var v58 uint64 = 0
                                                                _ = v58
                                                            default:
                                                                if v1 {
                                                                    v1 = (v3 == v3)
                                                                } else {
                                                                    var v59 int8 = -86
                                                                    v1 = v1
                                                                    var v60 float64 = 1e10
                                                                    if v1 {
                                                                        fmt.Println(v1)
                                                                        switch (fi41 - int(v3)) {
                                                                        case -933:
                                                                            v61 := GenS9{F0: 42, F1: v22, F2: (32767 > -440)}
                                                                            _ = v61
                                                                        default:
                                                                            v60 -= 3.14159
                                                                        }
                                                                        v22 = v22
                                                                    }
                                                                    v62 := GenS9{F0: 250, F1: 32767, F2: true}
                                                                    var v63 int = (^(fi41))
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                            }
                                                            if (!v1) {
                                                                v3 = -2147483648
                                                                if len(v21) > 0 {
                                                                    v21[0] = fi33
                                                                }
                                                                v2[(fi33 - 32768)] = 42
                                                            } else {
                                                                v64 := GenS10{F0: 225}
                                                                v65 := GenS4{F0: -1, F1: 0, F2: v1}
                                                                if v1 {
                                                                    v3 += v3
                                                                }
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                        }
                                                        v57 -= v57
                                                        _ = v57
                                                    }
                                                    _ = v56
                                                } else {
                                                    v1 = v1
                                                    v22 = v22
                                                    if (((-2.25 / 3.14159) != 0.1) && (-32768 != -303)) {
                                                        v22 += v22
                                                        if v1 {
                                                            var v66 string = "a"
                                                            v1 = (0 != 0)
                                                            v21 = append(v21, (fi33 >> ((23) & 63)))
                                                            _ = v66
                                                        }
                                                    }
                                                }
                                                if (float64(-2.25) > 0.0) {
                                                    if v1 {
                                                        fmt.Println(v22)
                                                        v67 := v2[(fi41 % ((9223372036854775807) | 1))]
                                                        _ = v67
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if ((v22 & 1) <= v22) {
                                                                v3 = v3
                                                                v3 = (v3 >> ((65535) & 63))
                                                            } else {
                                                                v68 := GenS11{F0: -273}
                                                                v2[32768] = fi41
                                                                _ = v68
                                                            }
                                                        case "":
                                                            v69 := GenS6{F0: (42 <= 42), F1: 771}
                                                            _ = v69
                                                        case "foo bar":
                                                            fmt.Println(fi33, v1, v3)
                                                        }
                                                    }
                                                } else {
                                                    var v70 int = -36
                                                    if v1 {
                                                        if (false && (113 == 1)) {
                                                            var v71 uint = 1
                                                            _ = v71
                                                        }
                                                        if (!true) {
                                                            if len(v21) > 0 {
                                                                v21[0] = 611
                                                            }
                                                            _ = len(v21)
                                                        }
                                                    }
                                                    _ = v70
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        case 127:
                            v1 = (false && v1)
                        }
                        delete(v2, fi33)
                        var v72 float64 = (0.1 / 1.5)
                        _ = v72
                    }
                    var v73 int8 = 100
                    _ = v50
                    _ = v73
                }
                var v74 int8 = 0
                if v1 {
                    if ((v22 % ((42) | 1)) < v22) {
                        var v75 uint32 = 65535
                        delete(v2, -603)
                        _ = v75
                    }
                    v76 := GenS12{F0: -9223372036854775808, F1: "foo bar"}
                    v77 := GenS12{F0: int64(fi33), F1: "\n\t"}
                    _ = v76
                    _ = v77
                } else {
                    v78 := v2[(^(fi33))]
                    fmt.Println(fi41, fi41)
                    _ = v78
                }
                v2[(-(fi41))] = (^(fi33))
                _ = fi41
                _ = v74
            }
            _ = v40
        }
        v1 = (42 >= 4294967295)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
