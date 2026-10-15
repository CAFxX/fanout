package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS10 struct { F0 int16; F1 uint8; F2 int }
func genFn0() uint64 {
    var v5 string = ("foo bar" + "a")
    var v6 uint16 = 0
    var v7 int32 = -1
    v8 := []int{9223372036854775807}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint64 {
    if (0.0 != float32(3.14159)) {
        switch 127 {
        case 32767:
            var v9 int16 = 32767
            _ = v9
        case 536:
            var v10 float32 = ((float32(3.14159) - ((1e10 + -2.25) - 0.1)) - 1e10)
            _ = v10
        case 127:
            var v11 int = 128
            _ = v11
        default:
            if false {
                v12 := GenS1{F0: 34}
                var v13 uint = 0
                _ = v12
                _ = v13
            } else {
                var v14 int32 = -266
                v15 := GenS1{F0: 6}
                var v16 uint64 = 2
                _ = v14
                _ = v15
                _ = v16
            }
        }
        var v17 uint8 = 127
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS2{F0: 127, F1: 42}
            _ = fi19
            _ = v20
        }
        switch (("a" + "\n\t") + "a") {
        case "foo bar":
            v21 := GenS1{F0: -1}
            _ = v21
        case "\n\t":
            v22 := GenS3{F0: ("hello" + ("" + "foo bar")), F1: 2}
            _ = v22
        default:
            v23 := GenS2{F0: 42, F1: 2147483647}
            _ = v23
        }
        _ = fi18
    }
    var v24 uint8 = 137
    var v25 uint64 = 534
    var v26 int64 = 42
    _ = v24
    _ = v25
    _ = v26
    return v25
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{9223372036854775807, 229, 128}
    v3 := make(map[int]int)
    var v4 float64 = ((1e10 * 0.1) * float64((1.5 + -2.25)))
    v3[2147483647] = -423
    var v27 int64 = -227
    var v28 float32 = (((1e10 - (0.0 * (((1.5 - 0.0) + (1.5 + float32(((v4 / (float64(3.14159) - 0.0)) - v4)))) / float32(-2.25)))) - (-2.25 * (1.5 / 1.0))) / 1.0)
    for fi29 := 0; fi29 < 13; fi29++ {
        v4 = v4
        _ = fi29
    }
    v30 := GenS4{F0: "", F1: 32767}
    var v31 float32 = v28
    var v32 uint32 = 42
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            switch (fi33 >> ((1) & 63)) {
            case 128:
                if (!(((2 > 42) && true) && (2 <= 42))) {
                    var v35 float32 = v31
                    v36 := GenS3{F0: "\n\t", F1: 65535}
                    var v37 int16 = -32768
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    v32 += v32
                    fmt.Println(v28, fi34)
                }
            case -917:
                v31 = v28
            case 2147483647:
                var v38 uint64 = 235
                _ = v38
            }
            _ = fi34
        }
        if (int32(v27) == 2147483647) {
            delete(v3, (fi33 & -847))
            var v39 int = (fi33 - (fi33 >> ((42) & 63)))
            for fi40 := 0; fi40 < 5; fi40++ {
                v31 = (3.14159 - 1.5)
                v41 := GenS5{F0: true}
                if ((false || (!false)) || false) {
                    var v42 uint8 = 1
                    fmt.Println(v27, fi33, v28)
                    var v43 int = int(v32)
                    _ = v42
                    _ = v43
                }
                _ = fi40
                _ = v41
            }
            _ = v39
        }
        v27 -= v27
        _ = fi33
    }
    if true {
        fmt.Println(v4, v27)
    }
    delete(v1, -239)
    fmt.Println(v32, v31, v28)
    var v44 string = ""
    switch v44 {
    case "hello":
        var v45 uint8 = 171
        _ = v45
    default:
        v31 += float32(v4)
    }
    v32 = (^(v32))
    v27 = 0
    if false {
        switch v44 {
        case "\n\t":
            var v46 int32 = 1
            _ = v46
        case "a":
            if true {
                if (true || true) {
                    switch 32768 {
                    case 9223372036854775807:
                        fmt.Println(v27)
                    case 1:
                        v31 -= (-2.25 * v31)
                    case 32768:
                        var v47 string = v44
                        _ = v47
                    default:
                        for fi48 := 0; fi48 < 16; fi48++ {
                            v49 := GenS1{F0: -128}
                            var v50 bool = ((42 <= -274) && ((float64(v31) == v4) && (v4 <= float64(((v31 / (1.5 + (v28 + (1e10 - (v28 / float32(((v4 - (3.14159 - (0.1 / 3.14159))) / v4))))))) / v28)))))
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                    }
                    v27 = (v27 * 220)
                    var v51 uint8 = 23
                    _ = v51
                } else {
                    var v52 uint = 4294967295
                    delete(v3, 9223372036854775807)
                    v53 := GenS6{F0: 1}
                    _ = v52
                    _ = v53
                }
            }
        case "":
            var v54 int32 = 2147483647
            _ = v54
        }
        v28 = 0.1
        switch 200 {
        case 4294967295:
            v55 := GenS2{F0: 128, F1: 1}
            _ = v55
        case 128:
            if (!true) {
                var v56 int8 = 100
                if false {
                    var v57 uint8 = 128
                    v57 = 179
                    _ = v57
                } else {
                    var v58 int8 = v56
                    v59 := GenS7{F0: 775, F1: 733}
                    _ = v58
                    _ = v59
                }
                _ = v56
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v60 := GenS2{F0: 0, F1: 2147483647}
                _ = v60
            }
        default:
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
        }
    }
    for fi61 := 0; fi61 < 5; fi61++ {
        if false {
            if (fi61 >= (^(fi61))) {
                switch (fi61 ^ (fi61 | (fi61 - (fi61 | (fi61 - fi61))))) {
                case 1:
                    v44 = ""
                case 32767:
                    if false {
                        v62 := GenS8{F0: 255, F1: 4294967295, F2: (v27 / ((42) | 1))}
                        v28 += -2.25
                        _ = v62
                    } else {
                        v63 := GenS5{F0: (4294967295 < genFn1())}
                        v64 := v1[fi61]
                        if ((!false) && (!true)) {
                            for fi65 := 0; fi65 < 9; fi65++ {
                                var v66 int16 = 808
                                _ = fi65
                                _ = v66
                            }
                            if true {
                                fmt.Println(v4, fi61)
                                if (127 < 139) {
                                    if (982 >= 32767) {
                                        for fi67 := 0; fi67 < 9; fi67++ {
                                            v31 += 3.14159
                                            _ = fi67
                                        }
                                    } else {
                                        v31 -= (0.1 * ((v31 * 0.1) + 1e10))
                                        var v68 float64 = v4
                                        _ = v68
                                    }
                                }
                                v69 := GenS9{F0: 296, F1: 1, F2: 42}
                                _ = v69
                            } else {
                                var v70 uint64 = uint64(v64)
                                _ = v70
                            }
                            var v71 bool = true
                            _ = v71
                        }
                        _ = v63
                        _ = v64
                    }
                }
                v72 := GenS10{F0: int16(v27), F1: 127, F2: (fi61 + fi61)}
                for fi73 := 0; fi73 < 10; fi73++ {
                    v28 += 0.0
                    v3[4294967295] = 843
                    var v74 uint64 = 0
                    _ = fi73
                    _ = v74
                }
                _ = v72
            }
            v75 := GenS5{F0: (fi61 < 2147483647)}
            var v76 uint8 = 38
            _ = v75
            _ = v76
        } else {
            v2 = append(v2, (fi61 | fi61))
        }
        v32 = v32
        _ = fi61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v44
}
