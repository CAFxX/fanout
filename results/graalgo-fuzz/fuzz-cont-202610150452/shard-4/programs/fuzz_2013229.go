package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 string; F2 int64 }
type GenS6 struct { F0 int8; F1 uint; F2 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint16; F1 int16; F2 int64 }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int16 }
func genFn0(p0 uint16, p1 string, p2 string) uint {
    v5 := []int{868, 2147483647, 128}
    var v6 int32 = -450
    v7 := GenS1{F0: 0, F1: 100}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int64, p1 int32, p2 string) uint8 {
    var v8 int32 = 2147483647
    v9 := GenS2{F0: 39}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 float32 = 3.14159
        if true {
            var v12 string = ""
            var v13 float64 = 1.5
            var v14 int8 = 127
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 float32 = v11
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v8
    _ = v9
    return 103
}

func main() {
    var v1 uint32 = 65535
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    var v4 uint16 = 32767
    if ((904 <= 0) || true) {
        var v17 float64 = 1.5
        if (0 >= 2147483647) {
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 int64 = 42
                var v20 int64 = v19
                var v21 int = (-(fi18))
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            v22 := GenS3{F0: v1}
            _ = v22
        }
        _ = v17
    } else {
        v2 = ((((("\n\t" + v2) + ("foo bar" + v2)) + "") + v2) + "")
        for fi23 := 0; fi23 < 20; fi23++ {
            v4 -= v4
            v1 = (v1 + (v1 << ((4294967295) & 63)))
            v24 := []int{0}
            _ = fi23
            _ = v24
        }
        fmt.Println(v2)
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            if true {
                var v27 uint16 = v4
                fmt.Println(fi25)
                v28 := []int{fi26, (fi26 ^ -537)}
                _ = v27
                _ = v28
            }
            v1 = (^(v1))
            _ = fi26
        }
        _ = fi25
    }
    v1 -= (v1 - v1)
    if false {
        for fi29 := 0; fi29 < 19; fi29++ {
            if false {
                if false {
                    var v30 int64 = 787
                    for fi31 := 0; fi31 < 14; fi31++ {
                        var v32 bool = (((109 != 42) || true) && (!(455 < (v1 + (v1 / (((v1 << ((genFn0((v4 / ((97) | 1)), v2, v2)) & 63))) | 1))))))
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
                v1 = v1
            }
            for fi33 := 0; fi33 < 3; fi33++ {
                _ = len(v3)
                _ = fi33
            }
            if (18446744073709551615 > 183) {
                switch "\n\t" {
                case "":
                    for fi34 := 0; fi34 < 12; fi34++ {
                        switch (v2 + (v2 + "")) {
                        case "foo bar":
                            if (!(-1 != 14)) {
                                v35 := []int{}
                                v2 = ""
                                _ = len(v3)
                                _ = v35
                            } else {
                                v1 = (^(v1))
                                delete(v3, v1)
                            }
                        case "\n\t":
                            switch (v2 + v2) {
                            case "\n\t":
                                if false {
                                    if false {
                                        v36 := v3[(v1 | (v1 / (((^(v1))) | 1)))]
                                        v37 := GenS3{F0: v1}
                                        var v38 uint32 = (v1 + v1)
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v39 := []int{(^(fi29))}
                                        var v40 float64 = 0.0
                                        v41 := GenS4{F0: v4}
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 int32 = 0
                                    v43 := GenS5{F0: ((v2 + v2) + ((v2 + "") + v2)), F1: v2, F2: 0}
                                    _ = v42
                                    _ = v43
                                } else {
                                    var v44 uint32 = v1
                                    v45 := GenS1{F0: 143, F1: 25}
                                    switch ((("\n\t" + v2) + v2) + ((v2 + v2) + v2)) {
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "hello":
                                            v2 = ((v2 + v2) + v2)
                                        case "\n\t":
                                            delete(v3, 2147483648)
                                        default:
                                            v46 := GenS4{F0: 254}
                                            v47 := GenS6{F0: 42, F1: 18446744073709551615, F2: 1}
                                            fmt.Println(fi29, fi29, v1)
                                            _ = v46
                                            _ = v47
                                        }
                                    case "a":
                                        switch 42 {
                                        case -617:
                                            var v48 uint32 = (v1 << ((uint(v1)) & 63))
                                            delete(v3, v44)
                                            v49 := GenS7{F0: -86}
                                            _ = v48
                                            _ = v49
                                        case -1:
                                            switch "a" {
                                            case "a":
                                                v2 = v2
                                            case "\n\t":
                                                if (!false) {
                                                    v50 := []int{127, fi34}
                                                    v1 = (v1 / (((v1 | (v1 & (v1 >> ((2) & 63))))) | 1))
                                                    _ = v50
                                                }
                                            }
                                        default:
                                            v51 := GenS8{F0: 0, F1: 32767}
                                            _ = v51
                                        }
                                    case "\n\t":
                                        var v52 float32 = float32((0.1 - 0.1))
                                        _ = v52
                                    default:
                                        var v53 uint64 = 4294967295
                                        _ = v53
                                    }
                                    v54 := v3[v44]
                                    v55 := GenS9{F0: (v4 ^ (v4 % (((^(v4))) | 1))), F1: 1, F2: -9223372036854775808}
                                    _ = v44
                                    _ = v45
                                    _ = v54
                                    _ = v55
                                }
                            case "":
                                if false {
                                    fmt.Println(fi34, v1, v1)
                                    if ((^(v1)) != (v1 >> ((42) & 63))) {
                                        v56 := GenS9{F0: v4, F1: -507, F2: 42}
                                        var v57 int8 = 127
                                        v58 := []int{fi34}
                                        v4 -= v4
                                        var v59 float64 = (((0.1 * 1.5) * 1e10) - ((3.14159 - float64(3.14159)) + 0.1))
                                        switch (^(fi29)) {
                                        case 2147483647:
                                            var v60 float64 = v59
                                            _ = v60
                                        default:
                                            delete(v3, (v1 & (v1 % ((v1) | 1))))
                                        }
                                        v1 += v1
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v61 int8 = -15
                                        v1 -= v1
                                        _ = v61
                                    }
                                } else {
                                    var v62 int = fi29
                                    fmt.Println(v62, v62)
                                    v63 := []int{fi29, v62}
                                    _ = v62
                                    _ = v63
                                }
                            }
                        default:
                            v64 := GenS10{F0: 0, F1: 935}
                            _ = v64
                        }
                        _ = fi34
                    }
                case "\n\t":
                    if false {
                        _ = len(v3)
                    }
                default:
                    var v65 int8 = 0
                    _ = v65
                }
                fmt.Println(v1, v1)
                v2 = v2
            }
            _ = fi29
        }
        var v66 uint32 = v1
        _ = v66
    }
    v67 := []int{-2147483648, 180, -2147483648}
    if (2 == 65535) {
        var v68 uint8 = 151
        v69 := v3[v1]
        var v70 int64 = 42
        _ = v68
        _ = v69
        _ = v70
    } else {
        var v71 float64 = 1.5
        if true {
            v3[v1] = (v1 <= v1)
            var v72 int8 = 1
            _ = v72
        }
        var v73 uint32 = (^(v1))
        _ = v71
        _ = v73
    }
    fmt.Println(v4, v1)
    switch -96 {
    case -744:
        v74 := GenS11{F0: -89}
        _ = v74
    case -2147483648:
        fmt.Println(v4, v2)
    case 527:
        var v75 int = 3
        _ = v75
    default:
        v76 := GenS12{F0: 437}
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
}
