package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 int; F2 int8 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint32 }
func genFn0(p0 int16, p1 uint8, p2 int16) uint {
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 -= (p0 | -349)
        p2 = (p2 ^ (p2 - p0))
        var v8 int8 = -5
        _ = fi7
        _ = v8
    }
    p1 = p1
    var v9 uint8 = 127
    v10 := []int{32768, -2147483648, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 2
    var v2 uint = 802
    var v3 bool = ("\n\t" < "\n\t")
    var v4 float64 = (0.1 - 3.14159)
    v5 := []uint32{2147483648, 185, 2147483648, 1}
    v6 := make(map[int]bool)
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := v6[fi11]
        v13 := GenS1{F0: 100, F1: uint16(v1), F2: 791}
        for fi14 := 0; fi14 < 14; fi14++ {
            v2 = v2
            _ = len(v6)
            v4 = v4
            _ = fi14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    v15 := GenS2{F0: -128}
    if ((0 > 0) || (-1 == 42)) {
        var v16 uint = (v2 | v2)
        var v17 uint64 = (v1 * v1)
        _ = len(v5)
        _ = v16
        _ = v17
    } else {
        v18 := GenS2{F0: 42}
        fmt.Println(v3, v2, v2)
        _ = v18
    }
    v2 -= v2
    v3 = ((!false) || v3)
    if (18446744073709551615 > v2) {
        var v19 bool = false
        _ = v19
    } else {
        v1 = v1
    }
    var v20 uint = uint(v1)
    delete(v6, -2147483648)
    var v21 uint16 = 42
    fmt.Println(v20, v21, v2)
    v6[128] = v3
    if (-674 <= -2147483648) {
        for fi22 := 0; fi22 < 14; fi22++ {
            if (1 > 1) {
                fmt.Println(v1)
            } else {
                var v23 uint16 = v21
                _ = v23
            }
            v24 := v6[-264]
            for fi25 := 0; fi25 < 4; fi25++ {
                v4 = 3.14159
                if len(v5) > 0 {
                    v5[0] = 1
                }
                _ = fi25
            }
            _ = fi22
            _ = v24
        }
        if (42 == 447) {
            var v26 bool = v3
            _ = v26
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            v21 += 42
            _ = fi27
        }
    } else {
        if (!v3) {
            v1 -= (v1 * (v1 | (v1 % (((v1 + v1)) | 1))))
            for fi28 := 0; fi28 < 15; fi28++ {
                if v3 {
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                }
                switch (fi28 + fi28) {
                case 1:
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v6[-2147483648] = (22 <= 180)
                        fmt.Println(v4)
                        _ = fi29
                    }
                case -409:
                    if true {
                        for fi30 := 0; fi30 < 2; fi30++ {
                            v2 = (v2 % ((v20) | 1))
                            switch "a" {
                            case "foo bar":
                                var v31 int64 = 397
                                _ = v31
                            default:
                                switch "\n\t" {
                                case "a":
                                    v20 += (v2 - (v20 - 4294967295))
                                    v1 += uint64(v20)
                                    if v3 {
                                        v2 += 111
                                        if (v3 || v3) {
                                            v2 += genFn0(519, 42, 0)
                                            if v3 {
                                                _ = len(v5)
                                                var v32 uint = (^(v2))
                                                _ = v32
                                            } else {
                                                v5 = append(v5, 204)
                                                v33 := GenS2{F0: 1}
                                                fmt.Println(fi30)
                                                _ = v33
                                            }
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if v3 {
                                            var v34 uint16 = v21
                                            var v35 uint32 = uint32(v20)
                                            _ = v34
                                            _ = v35
                                        } else {
                                            var v36 int16 = 357
                                            if len(v5) > 0 {
                                                v5[0] = 300
                                            }
                                            _ = v36
                                        }
                                        var v37 int = fi28
                                        _ = v37
                                    } else {
                                        if (int(v1) > (fi30 << ((v20) & 63))) {
                                            v38 := GenS3{F0: v3}
                                            _ = v38
                                        }
                                        if v3 {
                                            v39 := GenS1{F0: -128, F1: (v21 & v21), F2: 435}
                                            var v40 int32 = -2147483648
                                            delete(v6, (fi28 % (((-(fi28))) | 1)))
                                            _ = v39
                                            _ = v40
                                        } else {
                                            var v41 uint = 659
                                            switch "a" {
                                            case "\n\t":
                                                var v42 uint32 = 2147483648
                                                v43 := GenS2{F0: -1}
                                                _ = v42
                                                _ = v43
                                            default:
                                                v21 = v21
                                            }
                                            delete(v6, -9223372036854775808)
                                            _ = v41
                                        }
                                        fmt.Println(fi30, v21, v2)
                                        if (1 >= 4294967295) {
                                            if len(v5) > 0 {
                                                v5[0] = 65535
                                            }
                                            _ = len(v6)
                                        }
                                        var v44 uint = (v2 - v2)
                                        v45 := GenS3{F0: v3}
                                        _ = v44
                                        _ = v45
                                    }
                                case "\n\t":
                                    switch (fi28 ^ int(v21)) {
                                    case -1:
                                        if v3 {
                                            if v3 {
                                                if (true || (0 >= int64(v20))) {
                                                    fmt.Println(v21)
                                                    if true {
                                                        v46 := GenS3{F0: v3}
                                                        _ = v46
                                                    }
                                                    switch fi30 {
                                                    case 32768:
                                                        switch fi30 {
                                                        case 615:
                                                            var v47 uint32 = 593
                                                            _ = v47
                                                        case -2147483648:
                                                            switch (fi30 & (fi28 >> (((v20 ^ v20)) & 63))) {
                                                            case 127:
                                                                _ = len(v5)
                                                            default:
                                                                v3 = v3
                                                            }
                                                        case -1:
                                                            fmt.Println(v1, v4)
                                                        }
                                                    case 42:
                                                        v3 = (!true)
                                                    case 127:
                                                        v3 = true
                                                    default:
                                                        if ((fi30 | -1) == fi28) {
                                                            fmt.Println(fi30)
                                                            v48 := GenS4{F0: -54}
                                                            var v49 int8 = 49
                                                            var v50 int32 = 1
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    }
                                                    fmt.Println(fi28, fi28)
                                                    v51 := GenS1{F0: 33, F1: (v21 / (((v21 | v21)) | 1)), F2: -814}
                                                    var v52 uint8 = 255
                                                    fmt.Println(fi28, v20, v2)
                                                    var v53 uint = v2
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                }
                                                if (1.5 < 1.5) {
                                                    var v54 float64 = ((1e10 - v4) + (((v4 / float64(3.14159)) * ((((v4 + (v4 / 0.1)) * v4) * 1e10) + -2.25)) * v4))
                                                    v20 += v20
                                                    v3 = ((v2 ^ v20) != genFn0(-589, 0, -32768))
                                                    var v55 float64 = float64((float32(float64(1.5)) / 0.1))
                                                    var v56 bool = v3
                                                    var v57 uint8 = 128
                                                    var v58 uint32 = 80
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                }
                                                if ((2 > 0) || (699 != 2147483648)) {
                                                    v59 := GenS1{F0: 118, F1: v21, F2: -785}
                                                    var v60 bool = false
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    if ((!v3) || v3) {
                                                        if (4294967295 > 65535) {
                                                            var v61 float64 = ((v4 + float64(((3.14159 + 1.5) + float32(float64(3.14159))))) * float64((1e10 * 1.5)))
                                                            _ = v61
                                                        }
                                                    } else {
                                                        var v62 uint64 = v1
                                                        var v63 int64 = 180
                                                        if true {
                                                            if (-1 == -19) {
                                                                var v64 int16 = 1000
                                                                _ = v64
                                                            }
                                                            v4 -= 0.1
                                                            v21 -= 942
                                                        }
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    var v65 float32 = -2.25
                                                    _ = v65
                                                }
                                            } else {
                                                v66 := GenS2{F0: -69}
                                                _ = v66
                                            }
                                            if (!true) {
                                                fmt.Println(v21)
                                            } else {
                                                v67 := GenS3{F0: true}
                                                _ = v67
                                            }
                                            v68 := GenS5{F0: 2, F1: 128, F2: 1}
                                            _ = v68
                                        }
                                    case 1:
                                        v69 := GenS4{F0: -60}
                                        _ = v69
                                    case 32767:
                                        var v70 uint32 = 2147483648
                                        _ = v70
                                    }
                                }
                            }
                            _ = fi30
                        }
                        v4 = 1e10
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi28
            }
        }
        v1 -= (v1 * (v1 % (((v1 - v1)) | 1)))
        if ((false && ((v21 / ((348) | 1)) != (v21 + v21))) && v3) {
            v21 = v21
            if ((0.0 * 1.5) >= v4) {
                if v3 {
                    var v71 int8 = -1
                    _ = v71
                } else {
                    v5 = append(v5, 2147483647)
                    v3 = (2 != 32)
                    v21 = (v21 - 65535)
                }
            } else {
                v72 := GenS6{F0: v3, F1: 1, F2: 4294967295}
                _ = v72
            }
            v4 = v4
        }
    }
    v73 := GenS3{F0: ((^(v21)) == (v21 << ((644) & 63)))}
    fmt.Println(v4, v3, v3)
    v3 = true
    v74 := GenS5{F0: 2, F1: 128, F2: -13}
    v1 -= 18446744073709551615
    v1 = (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
