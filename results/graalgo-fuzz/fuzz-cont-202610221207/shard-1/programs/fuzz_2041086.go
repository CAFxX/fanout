package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 string = "\n\t"
    var v2 bool = true
    v3 := make(map[int]string)
    var v4 uint16 = 1
    var v5 uint8 = 128
    v6 := []string{((("foo bar" + "\n\t") + v1) + v1), (v1 + "hello")}
    v2 = (!v2)
    for fi7 := 0; fi7 < 20; fi7++ {
        v5 = (v5 % (((v5 % ((v5) | 1))) | 1))
        var v8 int64 = -1
        _ = fi7
        _ = v8
    }
    _ = len(v3)
    v4 -= v4
    v4 = (v4 | 42)
    var v9 int64 = 9223372036854775807
    var v10 float32 = (1e10 + -2.25)
    if ((1e10 > v10) || (((v1 == "hello") && v2) && true)) {
        v4 -= v4
        v9 -= int64(v5)
        var v11 int8 = int8(v5)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: 0}
            for fi14 := 0; fi14 < 1; fi14++ {
                switch (^(fi14)) {
                case 42:
                    if v2 {
                        v9 = (v9 % ((v9) | 1))
                        if false {
                            v4 += (v4 | 273)
                            v1 = v1
                            if len(v6) > 0 {
                                v6[0] = "foo bar"
                            }
                        }
                        var v15 uint32 = 1
                        var v16 int32 = 2147483647
                        _ = v15
                        _ = v16
                    } else {
                        switch (-(fi14)) {
                        case 2147483647:
                            v10 = -2.25
                        case -10:
                            var v17 int32 = -2147483648
                            _ = v17
                        }
                        var v18 uint16 = (v4 * v4)
                        v19 := GenS1{F0: 42}
                        _ = v18
                        _ = v19
                    }
                default:
                    fmt.Println(v1, v10)
                }
                var v20 int32 = -2147483648
                _ = fi14
                _ = v20
            }
            if ((fi12 << ((42) & 63)) >= fi12) {
                for fi21 := 0; fi21 < 10; fi21++ {
                    v5 = v5
                    if v2 {
                        v10 += v10
                        var v22 uint64 = 1
                        v10 = (-2.25 - 1.5)
                        _ = v22
                    } else {
                        if (!true) {
                            v3[fi21] = (v1 + v1)
                        } else {
                            switch v1 {
                            case "":
                                switch fi12 {
                                case 42:
                                    v10 = float32(0.1)
                                case 2147483647:
                                    v5 -= (^(v5))
                                }
                            case "foo bar":
                                var v23 int8 = 127
                                _ = v23
                            case "\n\t":
                                var v24 int8 = -1
                                _ = v24
                            default:
                                if v2 {
                                    if (32767 != (^(v4))) {
                                        var v25 string = v1
                                        v1 = v25
                                        _ = v25
                                    } else {
                                        if (508 != 65535) {
                                            if ((fi12 << ((158) & 63)) > fi12) {
                                                v9 = v9
                                            }
                                        } else {
                                            var v26 uint8 = v5
                                            v27 := GenS1{F0: 32767}
                                            _ = v26
                                            _ = v27
                                        }
                                        fmt.Println(v10, fi21, v2)
                                        if v2 {
                                            v4 = v4
                                        }
                                    }
                                    v28 := GenS2{F0: (^(v5)), F1: 42}
                                    _ = len(v6)
                                    var v29 bool = false
                                    if (true && v29) {
                                        v9 = 9223372036854775807
                                        if v2 {
                                            var v30 int64 = (v9 >> ((527) & 63))
                                            var v31 float32 = v10
                                            _ = len(v6)
                                            _ = v30
                                            _ = v31
                                        } else {
                                            _ = len(v6)
                                            switch 32768 {
                                            case 42:
                                                var v32 int = 891
                                                _ = v32
                                            case -505:
                                                if (v10 != ((v10 / (v10 / v10)) + v10)) {
                                                    if true {
                                                        var v33 uint = 303
                                                        switch fi12 {
                                                        case 128:
                                                            v34 := GenS1{F0: -724}
                                                            _ = v34
                                                        }
                                                        v35 := GenS1{F0: 32767}
                                                        if true {
                                                            if false {
                                                                var v36 int8 = 45
                                                                var v37 int = fi21
                                                                var v38 float32 = v10
                                                                _ = v36
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                            var v39 int16 = int16(v33)
                                                            if v2 {
                                                                v5 -= 255
                                                                if (v33 < v33) {
                                                                    v40 := GenS3{F0: (v33 << ((v33) & 63)), F1: 623}
                                                                    fmt.Println(v4)
                                                                    _ = v40
                                                                } else {
                                                                    var v41 uint32 = 2147483647
                                                                    var v42 bool = false
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                fmt.Println(v10)
                                                            } else {
                                                                delete(v3, (fi12 & 80))
                                                                if true {
                                                                    fmt.Println(v1, v2)
                                                                    if (!(fi12 == (fi21 - fi21))) {
                                                                        v9 -= v9
                                                                        if (1.5 > (v10 - v10)) {
                                                                            switch v1 {
                                                                            case "foo bar":
                                                                                var v43 float64 = (1e10 / 1.5)
                                                                                _ = v43
                                                                            }
                                                                            v4 -= v4
                                                                            if v29 {
                                                                                v4 += 32767
                                                                            } else {
                                                                                v44 := GenS2{F0: (v5 + (v5 % ((v5) | 1))), F1: (^(v33))}
                                                                                _ = len(v6)
                                                                                _ = v44
                                                                            }
                                                                        }
                                                                        v10 += v10
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                    }
                                                                    var v45 uint64 = 4294967295
                                                                    _ = v45
                                                                } else {
                                                                    v39 = v39
                                                                    v46 := GenS2{F0: v5, F1: v33}
                                                                    v47 := GenS4{F0: (v9 - (^(v9)))}
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                fmt.Println(v1, v39)
                                                            }
                                                            _ = v39
                                                        } else {
                                                            var v48 uint8 = (^(v5))
                                                            v29 = (!(32767 < 0))
                                                            _ = v48
                                                        }
                                                        _ = v33
                                                        _ = v35
                                                    } else {
                                                        v49 := GenS5{F0: "\n\t"}
                                                        _ = v49
                                                    }
                                                    if v29 {
                                                        if ((fi21 & (^(fi21))) > fi12) {
                                                            if v29 {
                                                                v50 := GenS4{F0: v9}
                                                                _ = v50
                                                            } else {
                                                                v3[fi21] = v1
                                                                var v51 uint = 2
                                                                _ = v51
                                                            }
                                                            var v52 float64 = (0.1 * 0.0)
                                                            _ = v52
                                                        } else {
                                                            fmt.Println(v10, v5, v10)
                                                        }
                                                        var v53 int32 = -747
                                                        _ = v53
                                                    } else {
                                                        var v54 int64 = v9
                                                        _ = v54
                                                    }
                                                }
                                            }
                                            switch -1 {
                                            case -1:
                                                v2 = (!(v29 || (!((!(!v29)) && v2))))
                                            case 42:
                                                var v55 uint32 = 42
                                                _ = v55
                                            case 9223372036854775807:
                                                var v56 int8 = -127
                                                _ = v56
                                            default:
                                                v57 := GenS6{F0: v1}
                                                _ = v57
                                            }
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v10 -= float32(-2.25)
                                    }
                                    var v58 uint = 325
                                    _ = v28
                                    _ = v29
                                    _ = v58
                                } else {
                                    v1 = v1
                                }
                            }
                            var v59 int32 = int32(fi21)
                            _ = v59
                        }
                        var v60 float32 = (((v10 - (v10 / float32(1.5))) - v10) * (v10 / 0.1))
                        v61 := GenS7{F0: v5}
                        _ = v60
                        _ = v61
                    }
                    _ = fi21
                }
                if ((v2 || (-61 >= 1)) || (-18 > -19)) {
                    var v62 uint64 = 489
                    _ = v62
                } else {
                    var v63 int8 = 0
                    var v64 int = (^(fi12))
                    _ = v63
                    _ = v64
                }
                _ = len(v3)
            }
            _ = fi12
            _ = v13
        }
        fmt.Println(v4, v4, v2)
    }
    if len(v6) > 0 {
        v6[0] = ((v1 + "\n\t") + ((v1 + v1) + v1))
    }
    v1 = (("" + "") + ((v1 + "foo bar") + (v1 + v1)))
    for fi65 := 0; fi65 < 3; fi65++ {
        v4 = 697
        fmt.Println(v4, v10)
        _ = fi65
    }
    switch 296 {
    case 9223372036854775807:
        fmt.Println(v1, v10, v10)
    }
    var v66 int16 = -1
    var v67 uint8 = (v5 | (v5 ^ (v5 ^ (^(v5)))))
    v3[2147483647] = v1
    v67 += v67
    var v68 uint64 = 0
    switch v1 {
    case "":
        var v69 int32 = 946
        _ = v69
    case "foo bar":
        var v70 uint8 = 168
        _ = v70
    default:
        var v71 float32 = (float32(((((float64(1e10) + -2.25) * ((0.0 - (-2.25 + float64((v10 + v10)))) + 0.1)) / 1e10) + 3.14159)) / (((((v10 * v10) + v10) - 1e10) / (v10 - (v10 + float32(3.14159)))) * -2.25))
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v9
}
