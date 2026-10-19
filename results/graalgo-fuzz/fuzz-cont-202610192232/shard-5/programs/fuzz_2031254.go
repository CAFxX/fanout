package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }
func genFn0(p0 int16, p1 int64, p2 int64) string {
    p0 = (p0 << ((476) & 63))
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 14; fi8++ {
        p1 = -591
        var v9 uint8 = 127
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1(p0 uint64) int64 {
    p0 -= p0
    v10 := GenS1{F0: 783}
    if true {
        v11 := GenS1{F0: 255}
        _ = v11
    }
    var v12 uint64 = p0
    switch 0 {
    case 69:
        if true {
            var v13 int64 = 422
            for fi14 := 0; fi14 < 8; fi14++ {
                if ((42 <= 65535) || true) {
                    switch "\n\t" {
                    case "\n\t":
                        v15 := []int{fi14, (fi14 * (-(fi14)))}
                        _ = v15
                    case "foo bar":
                        v12 -= p0
                    }
                }
                var v16 int = (-(fi14))
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 int32 = 0
                    v13 += v13
                    switch fi17 {
                    case -466:
                        if (1 != 42) {
                            if (!(568 >= 733)) {
                                v19 := GenS1{F0: 65535}
                                _ = v19
                            } else {
                                var v20 uint32 = 2
                                _ = v20
                            }
                            fmt.Println(fi14, v12, v12)
                        }
                    case -9223372036854775808:
                        fmt.Println(v12)
                    default:
                        v12 -= (v12 + 18446744073709551615)
                    }
                    v16 = (fi14 / ((fi14) | 1))
                    _ = fi17
                    _ = v18
                }
                _ = fi14
                _ = v16
            }
            var v21 float32 = (-2.25 / 1.0)
            _ = v13
            _ = v21
        }
    case 2147483647:
        for fi22 := 0; fi22 < 9; fi22++ {
            p0 -= p0
            _ = fi22
        }
    case 1:
        v23 := GenS1{F0: 32767}
        _ = v23
    default:
        v24 := GenS1{F0: uint16(v12)}
        _ = v24
    }
    _ = p0
    _ = v10
    _ = v12
    return int64(p0)
}

func main() {
    var v1 int = 622
    v2 := make(map[uint32]uint32)
    var v3 uint = 1
    var v4 int = 9223372036854775807
    var v5 int8 = -1
    var v6 bool = false
    var v25 uint = 255
    var v26 int64 = 42
    v2[189] = 93
    if (false && v6) {
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                if (!false) {
                    fmt.Println(v1, v26, v4)
                    var v29 uint8 = 42
                    fmt.Println(v6, fi27)
                    _ = v29
                } else {
                    var v30 uint16 = 870
                    v31 := GenS1{F0: v30}
                    v32 := []int{}
                    _ = v30
                    _ = v31
                    _ = v32
                }
                if ((!(v6 || true)) || ((v26 & 9223372036854775807) > v26)) {
                    v33 := []int{}
                    var v34 uint16 = 1
                    if (429 > -277) {
                        v35 := GenS1{F0: 894}
                        switch ("\n\t" + "hello") {
                        case "\n\t":
                            if (!true) {
                                var v36 float64 = (0.1 * 0.1)
                                _ = v36
                            }
                        case "foo bar":
                            if true {
                                if (v6 || v6) {
                                    if len(v33) >= 2 {
                                        v33 = v33[1:len(v33)]
                                    }
                                    v25 += v25
                                    fmt.Println(v1, fi27)
                                }
                                var v37 string = "\n\t"
                                if false {
                                    var v38 int64 = 1
                                    _ = v38
                                }
                                var v39 float64 = 0.1
                                fmt.Println(v39, v1, v37)
                                _ = v37
                                _ = v39
                            } else {
                                var v40 uint16 = (v34 | (v34 % (((^(v34))) | 1)))
                                v6 = (3.14159 == (1.5 * 0.0))
                                v41 := GenS2{F0: (-(v5)), F1: ((v3 >> ((255) & 63)) != (^(v3)))}
                                v42 := GenS2{F0: 61, F1: true}
                                delete(v2, 2147483647)
                                if v6 {
                                    if v6 {
                                        v43 := GenS3{F0: v26}
                                        fmt.Println(v40)
                                        if true {
                                            switch (-(v4)) {
                                            case 0:
                                                var v44 uint32 = 42
                                                _ = v44
                                            case 4294967295:
                                                switch (v4 << ((v3) & 63)) {
                                                case -416:
                                                    var v45 bool = v6
                                                    _ = v45
                                                case 375:
                                                    var v46 uint8 = 1
                                                    _ = v46
                                                }
                                            case 127:
                                                var v47 int16 = 0
                                                _ = v47
                                            default:
                                                _ = len(v2)
                                            }
                                            v26 = (^(v26))
                                        }
                                        _ = v43
                                    }
                                }
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        case "hello":
                            v48 := GenS2{F0: v5, F1: ((2 != 0) || v6)}
                            _ = v48
                        default:
                            v49 := v2[2]
                            _ = v49
                        }
                        switch "a" {
                        case "\n\t":
                            v3 += (v3 * (v25 * v3))
                        default:
                            var v50 uint16 = v34
                            _ = v50
                        }
                        _ = v35
                    } else {
                        if (v25 == v25) {
                            v51 := GenS4{F0: 689}
                            v52 := GenS3{F0: (v26 >> ((42) & 63))}
                            v53 := GenS1{F0: (v34 / ((v34) | 1))}
                            v25 -= (^(v3))
                            fmt.Println(v25, v1)
                            if false {
                                fmt.Println(v25)
                                if false {
                                    v54 := GenS5{F0: false}
                                    var v55 uint8 = 0
                                    var v56 uint32 = 65535
                                    v26 -= int64(v3)
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                } else {
                                    var v57 uint32 = 251
                                    v58 := GenS2{F0: -1, F1: v6}
                                    _ = v57
                                    _ = v58
                                }
                                var v59 int64 = (v26 + v26)
                                _ = v59
                            } else {
                                v26 = v26
                            }
                            var v60 int16 = 1
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v60
                        } else {
                            var v61 string = ("a" + "")
                            if v6 {
                                v25 += v25
                                var v62 int16 = 1
                                _ = v62
                            }
                            _ = v61
                        }
                    }
                    v33 = append(v33, (v1 - fi27))
                    _ = v33
                    _ = v34
                } else {
                    v63 := GenS3{F0: genFn1(1)}
                    v25 = v25
                    var v64 bool = (42 < 255)
                    v4 -= (fi28 | (v1 >> ((v25) & 63)))
                    _ = v63
                    _ = v64
                }
                fmt.Println(v3, v5, v4)
                _ = fi28
            }
            var v65 float32 = (-2.25 / float32((0.0 + 1e10)))
            _ = fi27
            _ = v65
        }
    }
    var v66 uint16 = 1
    switch "foo bar" {
    case "\n\t":
        v3 += v25
    case "hello":
        var v67 int = (v1 / ((-1) | 1))
        _ = v67
    }
    v68 := GenS2{F0: -1, F1: (1e10 <= float32(3.14159))}
    v69 := GenS4{F0: 435}
    v3 = (v25 | v25)
    if (!false) {
        fmt.Println(v5, v26, v25)
        for fi70 := 0; fi70 < 9; fi70++ {
            delete(v2, 0)
            _ = fi70
        }
    }
    fmt.Println(v4)
    v71 := GenS6{F0: 255}
    for fi72 := 0; fi72 < 17; fi72++ {
        for fi73 := 0; fi73 < 4; fi73++ {
            switch (("a" + "\n\t") + "foo bar") {
            case "hello":
                v74 := v2[4294967295]
                _ = v74
            }
            var v75 uint = (^(v3))
            v1 = v4
            var v76 int16 = 1
            _ = fi73
            _ = v75
            _ = v76
        }
        _ = fi72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v68
    _ = v69
    _ = v71
}
