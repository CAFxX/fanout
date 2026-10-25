package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 int32 }
type GenS5 struct { F0 bool; F1 string; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 int64 }
func genFn0(p0 string) uint {
    var v7 uint = 18446744073709551615
    var v8 int8 = 127
    _ = p0
    _ = v7
    _ = v8
    return (v7 << (((v7 * 615)) & 63))
}

func main() {
    var v1 uint32 = 65535
    var v2 int32 = -2147483648
    v3 := []uint32{4294967295, (v1 % (((v1 | (v1 ^ (v1 ^ v1)))) | 1)), (^(v1))}
    v4 := make(map[string]uint32)
    var v5 int16 = 0
    v6 := make(map[string]uint32)
    var v9 float64 = ((1.5 / 1e10) * -2.25)
    var v10 uint32 = 40
    switch -658 {
    case 32768:
        v2 = v2
    case 128:
        var v11 int8 = -128
        _ = v11
    case 127:
        if (1 <= 0) {
            for fi12 := 0; fi12 < 15; fi12++ {
                switch (fi12 % (((^(fi12))) | 1)) {
                case -936:
                    var v13 uint8 = 42
                    _ = v13
                }
                if true {
                    var v14 int32 = (^(v2))
                    switch (fi12 % ((fi12) | 1)) {
                    case -357:
                        fmt.Println(v10)
                    case 2147483647:
                        v1 += v1
                    case 4294967295:
                        fmt.Println(v14)
                    default:
                        if true {
                            for fi15 := 0; fi15 < 4; fi15++ {
                                if ((((517 != (fi15 / ((-6) | 1))) || (1 >= -1)) && false) && (true || (!true))) {
                                    var v16 float32 = 0.0
                                    _ = v16
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if false {
                                    v10 -= v10
                                } else {
                                    var v17 int = (fi12 + 127)
                                    _ = v17
                                }
                                _ = fi15
                            }
                            if true {
                                var v18 uint32 = (v1 * (v1 / (((v1 ^ v1)) | 1)))
                                _ = v18
                            } else {
                                v19 := GenS1{F0: v1, F1: (fi12 ^ (^(fi12)))}
                                _ = v19
                            }
                            v1 = (v1 & 110)
                        }
                    }
                    if true {
                        var v20 uint = genFn0("")
                        _ = v20
                    }
                    _ = v14
                }
                _ = fi12
            }
            var v21 string = ("" + ("\n\t" + ("" + "hello")))
            var v22 uint = genFn0("a")
            _ = v21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        fmt.Println(v5, v10)
        _ = fi23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v10 = (v1 + v10)
        if false {
            fmt.Println(v2)
            v25 := GenS2{F0: (^(v1))}
            _ = len(v4)
            _ = v25
        }
        _ = fi24
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS2{F0: v1}
        delete(v4, "")
        _ = fi26
        _ = v27
    }
    var v28 int8 = 100
    v9 = (0.1 / v9)
    if len(v3) > 0 {
        v3[0] = (v10 | v10)
    }
    v9 = 3.14159
    var v29 string = ""
    if ((-1 < v5) && (128 != 255)) {
        for fi30 := 0; fi30 < 16; fi30++ {
            _ = len(v6)
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 uint64 = 2
                var v33 float32 = 1.5
                v28 -= v28
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = fi30
        }
        fmt.Println(v28)
    }
    var v34 uint = 0
    if true {
        v35 := v6[v29]
        var v36 int = -716
        if true {
            v36 = v36
            v5 = int16(v36)
        } else {
            v9 = v9
            fmt.Println(v28)
        }
        _ = v35
        _ = v36
    } else {
        for fi37 := 0; fi37 < 20; fi37++ {
            v34 = (v34 >> ((genFn0(v29)) & 63))
            var v38 int = -2147483648
            _ = fi37
            _ = v38
        }
    }
    var v39 int64 = -1
    switch 9223372036854775807 {
    case 32767:
        if true {
            v39 -= (v39 ^ (v39 % ((v39) | 1)))
            v2 -= (v2 - v2)
            for fi40 := 0; fi40 < 7; fi40++ {
                v9 += float64(3.14159)
                for fi41 := 0; fi41 < 1; fi41++ {
                    v42 := v6[v29]
                    v43 := GenS3{F0: (v5 * v5), F1: int8(v5), F2: v39}
                    var v44 int64 = v39
                    if (true || (int32(v39) < 0)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = fi40
            }
        } else {
            v45 := GenS4{F0: (v5 % ((v5) | 1)), F1: (^(v10)), F2: v2}
            for fi46 := 0; fi46 < 6; fi46++ {
                switch fi46 {
                case 32767:
                    if (2 == 801) {
                        v47 := GenS4{F0: 42, F1: v10, F2: (-(v2))}
                        _ = len(v3)
                        _ = v47
                    } else {
                        var v48 float32 = float32((v9 * -2.25))
                        for fi49 := 0; fi49 < 15; fi49++ {
                            v29 = v29
                            var v50 int16 = 917
                            _ = fi49
                            _ = v50
                        }
                        if ((^(v34)) <= (v34 << (((v34 + v34)) & 63))) {
                            switch v29 {
                            case "hello":
                                var v51 string = v29
                                _ = v51
                            case "":
                                if (!true) {
                                    fmt.Println(v10)
                                    if false {
                                        fmt.Println(v28, v2)
                                        fmt.Println(v5, v5, v2)
                                    }
                                }
                            default:
                                for fi52 := 0; fi52 < 6; fi52++ {
                                    v53 := GenS5{F0: (509 > 32767), F1: "\n\t", F2: 847}
                                    var v54 int16 = (v5 * 1000)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                            }
                            v55 := GenS5{F0: true, F1: ("" + "a"), F2: -282}
                            var v56 uint = v34
                            _ = v55
                            _ = v56
                        } else {
                            var v57 int16 = v5
                            var v58 uint16 = 10
                            _ = v57
                            _ = v58
                        }
                        _ = v48
                    }
                }
                var v59 bool = false
                _ = fi46
                _ = v59
            }
            v10 = v10
            _ = v45
        }
    case -651:
        if (false && ((v1 / (((v1 + (v1 + v1))) | 1)) != 1)) {
            fmt.Println(v39, v39)
            for fi60 := 0; fi60 < 3; fi60++ {
                var v61 int64 = v39
                v62 := GenS2{F0: v1}
                _ = len(v6)
                _ = fi60
                _ = v61
                _ = v62
            }
        } else {
            v28 = (v28 ^ v28)
            switch 9223372036854775807 {
            case 32767:
                switch -1 {
                case 32767:
                    v1 = v10
                }
            case -1:
                var v63 uint32 = v1
                _ = v63
            case 1:
                if true {
                    v10 = v10
                    for fi64 := 0; fi64 < 13; fi64++ {
                        var v65 int16 = 1000
                        _ = fi64
                        _ = v65
                    }
                } else {
                    for fi66 := 0; fi66 < 15; fi66++ {
                        v67 := GenS4{F0: 1000, F1: 0, F2: -279}
                        v2 += v2
                        _ = fi66
                        _ = v67
                    }
                }
            default:
                fmt.Println(v34)
            }
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v68 float64 = v9
                v69 := GenS6{F0: 255}
                _ = v68
                _ = v69
            }
        }
    case 963:
        switch (v29 + v29) {
        case "foo bar":
            var v70 float64 = v9
            _ = v70
        case "a":
            _ = len(v4)
        }
    default:
        if false {
            switch v29 {
            case "a":
                var v71 uint64 = 184
                _ = v71
            default:
                if false {
                    for fi72 := 0; fi72 < 2; fi72++ {
                        var v73 uint = (^(v34))
                        var v74 float64 = v9
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                    fmt.Println(v28, v9)
                    var v75 uint16 = 562
                    _ = v75
                }
            }
        }
    }
    if false {
        v76 := GenS7{F0: (-(v2)), F1: v39}
        if ((^(v2)) < v2) {
            if len(v3) > 0 {
                v3[0] = v10
            }
        }
        switch v29 {
        case "foo bar":
            v77 := GenS5{F0: false, F1: v29, F2: v5}
            _ = v77
        default:
            v78 := GenS7{F0: -1, F1: (v39 + v39)}
            _ = v78
        }
        _ = v76
    } else {
        v39 -= (-(v39))
        var v79 float32 = (3.14159 - float32(0.0))
        for fi80 := 0; fi80 < 14; fi80++ {
            v81 := GenS3{F0: 1000, F1: v28, F2: v39}
            _ = len(v6)
            _ = fi80
            _ = v81
        }
        _ = v79
    }
    for fi82 := 0; fi82 < 13; fi82++ {
        switch "a" {
        case "hello":
            var v83 float64 = v9
            _ = v83
        case "foo bar":
            var v84 uint32 = v10
            _ = v84
        case "":
            v5 = v5
        }
        if true {
            v1 += v1
            v85 := GenS6{F0: 128}
            var v86 bool = false
            _ = v85
            _ = v86
        } else {
            v28 = -122
        }
        if true {
            v9 -= (0.0 / (v9 + ((1.5 - v9) / v9)))
            for fi87 := 0; fi87 < 1; fi87++ {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                if ((42 > 4) || (false && (true && false))) {
                    var v88 float32 = -2.25
                    _ = v88
                }
                _ = fi87
            }
            var v89 int8 = (v28 >> (((v34 | v34)) & 63))
            _ = v89
        }
        _ = fi82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
