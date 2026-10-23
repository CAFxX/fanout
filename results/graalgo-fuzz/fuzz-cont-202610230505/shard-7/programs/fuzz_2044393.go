package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int64; F1 int64 }
func genFn0(p0 uint, p1 bool, p2 uint16) int32 {
    v7 := GenS1{F0: p0}
    if (2 != 2) {
        p2 -= 42
        if (((282 > 9223372036854775807) && (!(-9223372036854775808 < 120))) || true) {
            v8 := []int{42}
            p2 -= p2
            v9 := GenS2{F0: "\n\t", F1: p2}
            _ = v8
            _ = v9
        }
    } else {
        v10 := GenS2{F0: "\n\t", F1: (p2 - (p2 >> ((p0) & 63)))}
        _ = v10
    }
    var v11 uint32 = uint32(p0)
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int64 = 42
        if ((816 >= (fi12 & (fi12 ^ 4294967295))) || (255 <= 128)) {
            v14 := []int{fi12, (fi12 - fi12)}
            _ = v14
        }
        v13 = 1
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return -472
}
func genFn1(p0 bool) uint16 {
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float32 = ((1e10 / -2.25) - 0.0)
        if p0 {
            if p0 {
                v17 := GenS1{F0: 1}
                _ = v17
            } else {
                if ((float32(0.1) + (3.14159 * (v16 + v16))) > v16) {
                    v18 := GenS1{F0: 2}
                    var v19 int32 = -2147483648
                    _ = v18
                    _ = v19
                }
                v20 := GenS3{F0: 1, F1: (!(true && (3.14159 != 1e10)))}
                _ = v20
            }
            var v21 uint = 896
            var v22 uint64 = 2
            _ = v21
            _ = v22
        } else {
            if true {
                v23 := []int{fi15, (fi15 | -2147483648), fi15}
                _ = v23
            } else {
                v24 := []int{}
                v24 = append(v24, (fi15 * (fi15 % (((-(fi15))) | 1))))
                var v25 int = (-(fi15))
                _ = v24
                _ = v25
            }
            var v26 uint64 = 42
            var v27 uint = 2
            _ = v26
            _ = v27
        }
        p0 = p0
        _ = fi15
        _ = v16
    }
    p0 = p0
    var v28 int32 = 42
    var v29 float64 = float64(0.0)
    _ = p0
    _ = v28
    _ = v29
    return 32767
}

func main() {
    var v1 uint64 = 2
    var v2 int64 = -9223372036854775808
    var v3 int8 = 100
    var v4 uint8 = 42
    var v5 int64 = v2
    v6 := make(map[uint32]string)
    v30 := []int{0, 2147483647, 116}
    var v31 float32 = 0.0
    v3 += (v3 + 42)
    v32 := GenS3{F0: 127, F1: false}
    v33 := GenS3{F0: (v3 % ((v3) | 1)), F1: (!(uint64(v4) >= v1))}
    var v34 float64 = 0.1
    for fi35 := 0; fi35 < 19; fi35++ {
        fmt.Println(v1, fi35, v3)
        _ = fi35
    }
    if (float32(1.5) <= (0.1 * 0.1)) {
        var v36 bool = true
        var v37 uint32 = uint32(v4)
        v6[(v37 - (v37 ^ (v37 << ((255) & 63))))] = "a"
        _ = v36
        _ = v37
    } else {
        fmt.Println(v1)
        switch "\n\t" {
        case "":
            if (v4 == v4) {
                var v38 float64 = (float64(v31) * (((v34 * (-2.25 + v34)) + (3.14159 / float64(3.14159))) - (v34 * v34)))
                _ = v38
            } else {
                var v39 uint8 = (^(v4))
                _ = v39
            }
        }
    }
    v34 = (0.1 - -2.25)
    for fi40 := 0; fi40 < 13; fi40++ {
        v5 -= 502
        v6[0] = ""
        for fi41 := 0; fi41 < 18; fi41++ {
            v42 := GenS2{F0: "a", F1: 65535}
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    if (921 != 838) {
        if true {
            switch 0 {
            case -2147483648:
                _ = len(v30)
            case 1:
                switch "foo bar" {
                case "\n\t":
                    v43 := GenS2{F0: "", F1: 1}
                    _ = v43
                default:
                    fmt.Println(v34, v31, v4)
                }
            default:
                switch 32767 {
                case 9223372036854775807:
                    var v44 bool = (!(genFn0(18446744073709551615, (true || (2147483647 > 513)), 431) > -2147483648))
                    _ = v44
                case 32767:
                    if len(v30) > 0 {
                        v30[0] = -410
                    }
                case 32768:
                    switch ("\n\t" + "a") {
                    case "a":
                        var v45 int32 = 1
                        _ = v45
                    case "":
                        v34 = 0.1
                    case "\n\t":
                        if (true && false) {
                            var v46 int64 = 856
                            _ = v46
                        } else {
                            fmt.Println(v3, v3, v4)
                            v4 = (v4 - (v4 + v4))
                        }
                    }
                }
            }
        } else {
            if (false && (751 >= -786)) {
                v30 = append(v30, 127)
                for fi47 := 0; fi47 < 8; fi47++ {
                    if (2147483647 >= 42) {
                        var v48 uint = 55
                        _ = v48
                    }
                    v49 := GenS1{F0: 523}
                    v50 := GenS3{F0: int8(v4), F1: true}
                    _ = fi47
                    _ = v49
                    _ = v50
                }
            }
        }
        if (0 > 65535) {
            v1 -= 956
        } else {
            var v51 uint32 = 2147483647
            _ = v51
        }
        switch "\n\t" {
        case "":
            for fi52 := 0; fi52 < 5; fi52++ {
                fmt.Println(v4)
                for fi53 := 0; fi53 < 6; fi53++ {
                    if ((((fi53 / ((fi52) | 1)) != (-(fi52))) && true) || false) {
                        var v54 bool = (-32768 != 42)
                        v54 = (768 == 42)
                        v55 := GenS1{F0: 65535}
                        _ = v54
                        _ = v55
                    }
                    if true {
                        var v56 float64 = (v34 * v34)
                        var v57 int16 = -1
                        v1 = 4294967295
                        _ = v56
                        _ = v57
                    } else {
                        var v58 uint16 = 255
                        var v59 int8 = -1
                        switch ("a" + "\n\t") {
                        case "hello":
                            var v60 int = -1
                            _ = v60
                        default:
                            v1 -= (v1 << ((65535) & 63))
                        }
                        _ = v58
                        _ = v59
                    }
                    fmt.Println(v31, v34, v2)
                    v61 := GenS4{F0: false, F1: (v4 ^ v4)}
                    _ = fi53
                    _ = v61
                }
                _ = fi52
            }
        case "foo bar":
            v4 -= 195
        default:
            var v62 uint32 = 2147483647
            _ = v62
        }
    } else {
        for fi63 := 0; fi63 < 17; fi63++ {
            v64 := GenS2{F0: "a", F1: 42}
            _ = len(v30)
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                for fi65 := 0; fi65 < 19; fi65++ {
                    v1 += v1
                    fmt.Println(v31, v34, v1)
                    if true {
                        _ = len(v30)
                        v1 = 4294967295
                        var v66 float64 = 1.5
                        _ = v66
                    }
                    v67 := v6[65535]
                    _ = fi65
                    _ = v67
                }
            case "\n\t":
                v68 := GenS3{F0: v3, F1: (2147483647 > uint32(v1))}
                _ = v68
            case "a":
                switch (-(fi63)) {
                case -582:
                    v2 -= v5
                case -9223372036854775808:
                    var v69 bool = true
                    _ = v69
                case -871:
                    v70 := GenS5{F0: 919, F1: (v4 - v4)}
                    _ = v70
                }
            default:
                v71 := GenS6{F0: (-(v2)), F1: v5}
                _ = v71
            }
            _ = fi63
            _ = v64
        }
        var v72 int = 127
        _ = v72
    }
    var v73 bool = (731 <= 1)
    if len(v30) > 0 {
        v30[0] = -806
    }
    v30 = append(v30, 1)
    var v74 int16 = -32768
    var v75 int64 = (v5 >> ((42) & 63))
    var v76 int = 1
    switch -869 {
    case 2147483647:
        v2 += (-(v75))
    case 0:
        v4 = 187
    default:
        v5 += v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
