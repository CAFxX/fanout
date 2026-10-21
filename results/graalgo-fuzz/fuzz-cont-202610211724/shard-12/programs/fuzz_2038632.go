package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 uint; F1 uint; F2 string }
func genFn0() uint8 {
    var v6 int64 = -9223372036854775808
    v6 = (v6 - v6)
    var v7 uint64 = 703
    var v8 int64 = (-(v6))
    _ = v6
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 bool = true
    v2 := []int{4294967295}
    v3 := make(map[uint32]int)
    var v4 string = "a"
    var v5 float32 = float32(1e10)
    var v9 uint32 = 2147483647
    switch -949 {
    case 32767:
        v2 = append(v2, -9223372036854775808)
    case 32768:
        for fi10 := 0; fi10 < 17; fi10++ {
            v4 = (v4 + ("foo bar" + ("\n\t" + ((v4 + v4) + v4))))
            switch (fi10 - fi10) {
            case 807:
                v11 := GenS1{F0: 361, F1: v9}
                _ = v11
            }
            _ = fi10
        }
    }
    var v12 uint8 = 0
    v13 := GenS2{F0: 49, F1: 0, F2: 65535}
    var v14 int8 = 127
    v15 := GenS3{F0: 0, F1: -32768}
    var v16 float64 = 0.0
    var v17 uint16 = 0
    v16 = ((v16 + (v16 + 3.14159)) - -2.25)
    switch -178 {
    case 32767:
        _ = len(v2)
    case 2147483647:
        var v18 string = (v4 + "\n\t")
        _ = v18
    case 32768:
        switch 756 {
        case -256:
            v9 += 4294967295
        case -9223372036854775808:
            for fi19 := 0; fi19 < 8; fi19++ {
                v2 = append(v2, fi19)
                var v20 uint64 = 18446744073709551615
                var v21 uint = 625
                _ = fi19
                _ = v20
                _ = v21
            }
        case 973:
            if len(v2) > 0 {
                v2[0] = -556
            }
        default:
            v22 := GenS3{F0: 42, F1: 296}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        fmt.Println(v14)
        v17 = v17
        _ = fi23
    }
    if (0 == 1) {
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
        switch (v4 + v4) {
        case "hello":
            fmt.Println(v14)
        case "\n\t":
            var v24 uint64 = 42
            _ = v24
        case "":
            for fi25 := 0; fi25 < 9; fi25++ {
                if v1 {
                    switch (fi25 / ((fi25) | 1)) {
                    case 347:
                        v26 := GenS2{F0: v14, F1: 42, F2: v17}
                        _ = v26
                    case -435:
                        if (32767 >= 0) {
                            for fi27 := 0; fi27 < 1; fi27++ {
                                v28 := GenS3{F0: 42, F1: 32767}
                                var v29 string = (v4 + v4)
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            var v30 int16 = -986
                            _ = v30
                        } else {
                            fmt.Println(v12, v4)
                        }
                    case -9223372036854775808:
                        v17 -= v17
                    }
                    v31 := GenS3{F0: 1, F1: 0}
                    if (!((v17 + v17) <= v17)) {
                        v32 := GenS3{F0: 0, F1: 32767}
                        if (false && true) {
                            v33 := GenS4{F0: (^(v9)), F1: -1, F2: 4294967295}
                            if (614 > 32767) {
                                v2 = append(v2, (fi25 % ((fi25) | 1)))
                                fmt.Println(v17, v5, v9)
                            } else {
                                fmt.Println(v4, v17)
                            }
                            v34 := GenS5{F0: (v12 * v12)}
                            _ = v33
                            _ = v34
                        } else {
                            if ((v12 & v12) == (^(v12))) {
                                fmt.Println(v16, v14, v1)
                                for fi35 := 0; fi35 < 13; fi35++ {
                                    v17 = (v17 % ((0) | 1))
                                    if true {
                                        if (uint64(v9) <= 2) {
                                            delete(v3, (^(v9)))
                                            if ((v9 % ((v9) | 1)) <= (v9 >> ((0) & 63))) {
                                                v36 := GenS6{F0: 32767, F1: 1, F2: int64(v17)}
                                                var v37 string = v4
                                                _ = v36
                                                _ = v37
                                            } else {
                                                _ = len(v3)
                                                if true {
                                                    if v1 {
                                                        v38 := GenS7{F0: v12}
                                                        var v39 uint8 = (v12 + uint8(v9))
                                                        _ = v38
                                                        _ = v39
                                                    } else {
                                                        var v40 uint32 = uint32(v12)
                                                        v16 = v16
                                                        fmt.Println(v9, v16)
                                                        _ = v40
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = fi35
                                                    }
                                                } else {
                                                    v41 := GenS8{F0: 1}
                                                    var v42 int64 = -807
                                                    v14 += v14
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                            v43 := GenS5{F0: (v12 + 0)}
                                            _ = v43
                                        } else {
                                            var v44 int = int(v9)
                                            v14 = (v14 << ((127) & 63))
                                            _ = v44
                                        }
                                        if v1 {
                                            switch (fi35 + (fi25 % ((0) | 1))) {
                                            case 9223372036854775807:
                                                if (((^(fi25)) == fi25) || v1) {
                                                    v16 -= 3.14159
                                                    switch v4 {
                                                    case "":
                                                        v17 = (v17 << ((1) & 63))
                                                    case "foo bar":
                                                        if v1 {
                                                            var v45 int64 = 1
                                                            _ = v45
                                                        } else {
                                                            if true {
                                                                fmt.Println(fi35, v4)
                                                                if (!false) {
                                                                    if v1 {
                                                                        if (!(42 > 501)) {
                                                                            if v1 {
                                                                                v17 = (v17 / ((v17) | 1))
                                                                                v14 -= 127
                                                                            }
                                                                            if ((!v1) || (v17 <= uint16(v9))) {
                                                                                var v46 int16 = 1
                                                                                _ = v46
                                                                            }
                                                                        } else {
                                                                            var v47 float32 = (v5 * (3.14159 / 3.14159))
                                                                            var v48 int8 = 100
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                        switch int(v12) {
                                                                        case -189:
                                                                            v49 := v3[(v9 / ((2147483648) | 1))]
                                                                            _ = v49
                                                                        case 127:
                                                                            fmt.Println(fi35)
                                                                        case 254:
                                                                            v3[v9] = (fi25 - fi35)
                                                                        default:
                                                                            v16 = v16
                                                                        }
                                                                        v17 = 42
                                                                    }
                                                                } else {
                                                                    var v50 int32 = -2147483648
                                                                    var v51 uint16 = 1
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                                v52 := GenS3{F0: 9223372036854775807, F1: -32768}
                                                                _ = v52
                                                            } else {
                                                                var v53 int32 = 2147483647
                                                                v53 += v53
                                                                v54 := GenS8{F0: v9}
                                                                var v55 int8 = (v14 + 0)
                                                                _ = v53
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            v56 := GenS1{F0: (v17 | 2), F1: (v9 >> ((65535) & 63))}
                                                            _ = v56
                                                        }
                                                    case "hello":
                                                        v5 -= float32(0.0)
                                                    default:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                            case 996:
                                                var v57 uint32 = 0
                                                _ = v57
                                            case -376:
                                                var v58 int16 = int16(v17)
                                                _ = v58
                                            }
                                            var v59 bool = (!(int(v17) >= 32767))
                                            _ = v59
                                        }
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = fi25
                                        }
                                        v60 := GenS8{F0: 65535}
                                        if v1 {
                                            _ = len(v2)
                                            var v61 int = (fi35 << ((0) & 63))
                                            _ = v61
                                        } else {
                                            if (-9223372036854775808 > 0) {
                                                switch (fi35 & fi35) {
                                                case 1:
                                                    if (0.1 >= float64(v5)) {
                                                        v12 += 22
                                                    } else {
                                                        v3[v9] = (fi25 % (((fi35 & int(v17))) | 1))
                                                    }
                                                case 32767:
                                                    _ = len(v2)
                                                case 343:
                                                    var v62 uint = 65535
                                                    _ = v62
                                                }
                                                if v1 {
                                                    if v1 {
                                                        fmt.Println(v5)
                                                        v12 = v12
                                                    } else {
                                                        v63 := GenS7{F0: v12}
                                                        v4 = v4
                                                        var v64 uint8 = v12
                                                        var v65 uint = 255
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    v17 -= 0
                                                }
                                            } else {
                                                v9 = v9
                                            }
                                            if (!true) {
                                                if v1 {
                                                    v1 = false
                                                    v4 = ""
                                                }
                                            } else {
                                                v3[(^(v9))] = fi25
                                                if true {
                                                    var v66 int32 = 0
                                                    if (("hello" <= v4) || v1) {
                                                        _ = len(v3)
                                                        if len(v2) > 0 {
                                                            v2[0] = fi25
                                                        }
                                                        v16 += ((v16 - 0.0) / (1.5 - (v16 - v16)))
                                                        var v67 uint16 = v17
                                                        v68 := GenS2{F0: 94, F1: 108, F2: 255}
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    v17 += (v17 % ((v17) | 1))
                                                    v3[18] = 911
                                                    _ = v66
                                                } else {
                                                    v14 -= v14
                                                    v9 = v9
                                                    if v1 {
                                                        fmt.Println(v5)
                                                    }
                                                    if (0 == 1) {
                                                        v69 := GenS2{F0: 48, F1: -1, F2: v17}
                                                        delete(v3, v9)
                                                        switch (v4 + v4) {
                                                        case "foo bar":
                                                            _ = len(v3)
                                                        default:
                                                            var v70 bool = ((v5 / v5) <= v5)
                                                            _ = v70
                                                        }
                                                        fmt.Println(v16, fi35, fi35)
                                                        _ = v69
                                                    }
                                                }
                                            }
                                            var v71 int8 = v14
                                            _ = v71
                                        }
                                        _ = v60
                                    }
                                    if (994 == 42) {
                                        v5 = v5
                                    }
                                    v1 = (genFn0() > v12)
                                    _ = fi35
                                }
                                fmt.Println(v4, v4)
                            }
                            v16 -= v16
                        }
                        switch ("\n\t" + v4) {
                        case "\n\t":
                            fmt.Println(v4)
                        case "a":
                            var v72 string = v4
                            _ = v72
                        default:
                            if v1 {
                                fmt.Println(v5, v16)
                                v2 = append(v2, fi25)
                            }
                        }
                        _ = v32
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v73 uint = 4294967295
                        v74 := GenS9{F0: v9, F1: 42}
                        _ = v73
                        _ = v74
                    }
                    _ = v31
                }
                _ = fi25
            }
        }
        var v75 int32 = int32(v9)
        _ = v75
    } else {
        v4 = v4
        v4 = ((v4 + "foo bar") + (("\n\t" + (v4 + v4)) + v4))
    }
    var v76 bool = v1
    v77 := v3[uint32(v17)]
    if false {
        switch v77 {
        case -2147483648:
            for fi78 := 0; fi78 < 12; fi78++ {
                fmt.Println(v5, v77, v12)
                _ = fi78
            }
        case 32767:
            v79 := v3[v9]
            _ = v79
        default:
            if (!(56 == 1)) {
                v77 += 128
                v80 := GenS10{F0: 166, F1: 4294967295, F2: ("a" + v4)}
                _ = v80
            }
        }
        var v81 uint8 = (v12 + genFn0())
        for fi82 := 0; fi82 < 9; fi82++ {
            var v83 uint64 = 1
            var v84 float32 = v5
            var v85 int = -384
            _ = fi82
            _ = v83
            _ = v84
            _ = v85
        }
        _ = v81
    } else {
        var v86 int8 = v14
        fmt.Println(v4)
        var v87 uint8 = (^(v12))
        _ = v86
        _ = v87
    }
    v76 = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
    _ = v9
}
