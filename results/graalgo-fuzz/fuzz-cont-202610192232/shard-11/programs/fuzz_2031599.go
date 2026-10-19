package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint64; F1 string; F2 int }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int64; F1 int8; F2 int16 }
type GenS9 struct { F0 int64 }
func genFn0(p0 int16, p1 int16) int32 {
    p0 += (p1 & (p0 >> ((30) & 63)))
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            p0 += p0
            _ = fi5
        }
        _ = fi4
    }
    var v6 uint16 = 2
    v7 := []int{843, 2147483647, 1, -612}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 int8, p2 uint) string {
    v8 := []int{329, -2147483648, 0}
    var v9 uint16 = 840
    switch 2147483647 {
    case 1:
        v9 -= v9
    case 128:
        for fi10 := 0; fi10 < 16; fi10++ {
            v8 = append(v8, (fi10 | (fi10 * (fi10 ^ 0))))
            _ = fi10
        }
    case 9223372036854775807:
        for fi11 := 0; fi11 < 1; fi11++ {
            if false {
                v12 := GenS1{F0: false, F1: 42, F2: -1}
                for fi13 := 0; fi13 < 17; fi13++ {
                    switch "\n\t" {
                    case "":
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    case "a":
                        switch (("a" + "\n\t") + "\n\t") {
                        case "":
                            if (-1 != 833) {
                                var v14 float64 = (0.0 / 1.0)
                                if false {
                                    v9 = 942
                                    v15 := GenS1{F0: (false && ((!true) || false)), F1: 872, F2: 42}
                                    _ = v15
                                } else {
                                    v16 := GenS2{F0: -1, F1: 4294967295}
                                    if false {
                                        p2 += (^(p0))
                                    } else {
                                        p2 = (p0 * (p2 >> ((p2) & 63)))
                                        if ((fi13 / (((fi11 + fi13)) | 1)) < fi11) {
                                            v14 -= v14
                                            p1 += p1
                                            var v17 uint8 = 2
                                            _ = v17
                                        }
                                        v18 := GenS2{F0: genFn0(952, 32767), F1: uint32(v9)}
                                        _ = v18
                                    }
                                    _ = v16
                                }
                                var v19 bool = false
                                _ = v14
                                _ = v19
                            }
                            v20 := GenS3{F0: 2, F1: genFn0(-269, 1)}
                            _ = v20
                        case "a":
                            if false {
                                if (false && true) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v21 uint32 = 2147483647
                                        _ = v21
                                    default:
                                        if true {
                                            p0 = 397
                                        }
                                    }
                                    var v22 float64 = (1e10 - 1.5)
                                    v23 := GenS2{F0: 42, F1: 2147483647}
                                    _ = v22
                                    _ = v23
                                }
                                var v24 uint64 = uint64(p2)
                                var v25 int16 = 32767
                                _ = v24
                                _ = v25
                            }
                        case "\n\t":
                            v26 := GenS3{F0: 0, F1: 285}
                            _ = v26
                        }
                    case "hello":
                        var v27 float32 = 1e10
                        _ = v27
                    default:
                        v9 = v9
                    }
                    var v28 float64 = 0.0
                    _ = fi13
                    _ = v28
                }
                _ = v12
            }
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 uint = p0
                var v31 int8 = (p1 ^ (p1 % ((p1) | 1)))
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 uint16 = v9
            _ = fi11
            _ = v32
        }
    default:
        if true {
            for fi33 := 0; fi33 < 3; fi33++ {
                if (!true) {
                    if true {
                        switch ("hello" + ("\n\t" + ("" + "a"))) {
                        case "\n\t":
                            p1 -= p1
                        case "hello":
                            if (1e10 <= (1e10 - (1e10 + (3.14159 * 1.5)))) {
                                if false {
                                    if (!false) {
                                        for fi34 := 0; fi34 < 4; fi34++ {
                                            if ((-9223372036854775808 != 9223372036854775807) && false) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch ("foo bar" + "a") {
                                                    case "hello":
                                                        var v35 int = (fi34 - (fi33 * (fi33 >> ((uint(v9)) & 63))))
                                                        _ = v35
                                                    case "":
                                                        v36 := GenS4{F0: 263, F1: 141, F2: -974}
                                                        _ = v36
                                                    default:
                                                        v8 = append(v8, (fi33 - (fi34 % ((fi33) | 1))))
                                                    }
                                                }
                                                if true {
                                                    if (true || true) {
                                                        p1 = (^(p1))
                                                        if ((p2 / (((p2 * (p2 | p2))) | 1)) == (^(p0))) {
                                                            fmt.Println(v9, p0)
                                                        }
                                                        p1 = -128
                                                    } else {
                                                        switch ("hello" + "") {
                                                        case "hello":
                                                            fmt.Println(fi34, fi34, p1)
                                                        case "":
                                                            v8 = append(v8, fi34)
                                                        default:
                                                            var v37 uint64 = 42
                                                            _ = v37
                                                        }
                                                    }
                                                    v38 := GenS5{F0: -151, F1: 0}
                                                    _ = v38
                                                }
                                            }
                                            v39 := GenS1{F0: (255 > 1), F1: 66, F2: -32768}
                                            var v40 uint32 = 65535
                                            _ = fi34
                                            _ = v39
                                            _ = v40
                                        }
                                        for fi41 := 0; fi41 < 11; fi41++ {
                                            var v42 uint = p0
                                            _ = fi41
                                            _ = v42
                                        }
                                    }
                                    for fi43 := 0; fi43 < 16; fi43++ {
                                        var v44 uint8 = 128
                                        if (false && false) {
                                            _ = len(v8)
                                        } else {
                                            p1 += p1
                                            var v45 uint16 = v9
                                            if (true || false) {
                                                p2 -= 1
                                                v46 := GenS6{F0: 16, F1: ("foo bar" + "foo bar"), F2: (fi33 ^ -993)}
                                                v47 := GenS7{F0: 916, F1: fi33}
                                                _ = v46
                                                _ = v47
                                            } else {
                                                if len(v8) > 0 {
                                                    v8[0] = (fi43 % ((fi43) | 1))
                                                }
                                            }
                                            _ = v45
                                        }
                                        _ = fi43
                                        _ = v44
                                    }
                                    var v48 int16 = -32768
                                    _ = v48
                                }
                            }
                        default:
                            var v49 int = fi33
                            _ = v49
                        }
                        p2 += 0
                        if ("foo bar" == "") {
                            fmt.Println(fi33, p2)
                        }
                    }
                } else {
                    for fi50 := 0; fi50 < 11; fi50++ {
                        v51 := GenS8{F0: 0, F1: p1, F2: -32768}
                        v52 := GenS8{F0: -1, F1: p1, F2: 0}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    var v53 int64 = 1
                    _ = v53
                }
                _ = fi33
            }
            var v54 int8 = 42
            _ = v54
        } else {
            for fi55 := 0; fi55 < 20; fi55++ {
                var v56 int = (fi55 >> ((p0) & 63))
                v57 := GenS9{F0: -1}
                _ = fi55
                _ = v56
                _ = v57
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return ("hello" + "a")
}

func main() {
    var v1 int16 = 32767
    var v2 float32 = (3.14159 - 0.0)
    v3 := make(map[int]int)
    v58 := GenS9{F0: 9223372036854775807}
    var v59 int16 = (v1 & v1)
    v2 -= (float32(float64(v2)) - v2)
    delete(v3, 1)
    if ((2 > 11) && true) {
        if (false && (!(397 >= 208))) {
            if true {
                var v60 uint64 = 42
                _ = v60
            } else {
                v1 = (-(v1))
                var v61 int8 = 127
                for fi62 := 0; fi62 < 15; fi62++ {
                    var v63 int = (-(fi62))
                    _ = fi62
                    _ = v63
                }
                _ = v61
            }
        }
        var v64 uint64 = 0
        _ = v64
    } else {
        v65 := []int{4294967295}
        var v66 int64 = -1
        _ = v65
        _ = v66
    }
    v59 = (v1 * v1)
    var v67 uint16 = 118
    var v68 int64 = -1
    v3[867] = -367
    v68 = v68
    v68 -= v68
    var v69 uint = 220
    if ((42 == 0) || (2 >= 42)) {
        var v70 int = -2147483648
        _ = v70
    }
    v71 := GenS4{F0: v67, F1: 1, F2: v59}
    var v72 int = -792
    for fi73 := 0; fi73 < 7; fi73++ {
        var v74 uint64 = uint64(v69)
        v75 := GenS8{F0: v68, F1: -86, F2: (v1 % (((v1 - -475)) | 1))}
        _ = fi73
        _ = v74
        _ = v75
    }
    v76 := []int{}
    switch (("a" + "foo bar") + "hello") {
    case "foo bar":
        v59 -= int16(v72)
    case "":
        var v77 int = 128
        _ = v77
    case "a":
        v78 := GenS8{F0: v68, F1: 100, F2: (v1 << (((v69 + 42)) & 63))}
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v59:", v59)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v72:", v72)
    fmt.Println("slicelen v76:", len(v76), v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
    _ = v59
    _ = v67
    _ = v68
    _ = v69
    _ = v71
    _ = v72
    _ = v76
}
