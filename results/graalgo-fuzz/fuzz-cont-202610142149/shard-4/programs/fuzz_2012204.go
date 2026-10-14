package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("" + "foo bar"), "foo bar"}
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    var v5 float64 = (3.14159 - 1.5)
    var v6 int8 = 100
    var v7 int = -383
    switch ("\n\t" + "") {
    case "a":
        v6 = (v6 / ((v6) | 1))
    case "hello":
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint32 = 42
            v10 := GenS1{F0: 443}
            _ = fi8
            _ = v10
            _ = v9
        }
    case "":
        for fi11 := 0; fi11 < 15; fi11++ {
            v7 += (^(fi11))
            if false {
                if false {
                    v12 := GenS2{F0: -2147483648, F1: v6, F2: 1}
                    if (!true) {
                        v2 = append(v2, "")
                        var v13 int64 = 42
                        _ = v13
                    }
                    _ = len(v3)
                    _ = v12
                } else {
                    v14 := GenS1{F0: 771}
                    v5 = v5
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v5 = v5
                        _ = fi15
                    }
                    _ = v14
                }
                v16 := GenS2{F0: 0, F1: -3, F2: 65535}
                _ = v16
            } else {
                _ = len(v2)
            }
            _ = fi11
        }
    default:
        delete(v1, v7)
    }
    v17 := GenS2{F0: -506, F1: (v6 * v6), F2: 0}
    for fi18 := 0; fi18 < 19; fi18++ {
        switch v7 {
        case -2147483648:
            if (float32(v5) <= 0.0) {
                if false {
                    switch (v7 + fi18) {
                    case 906:
                        if (false || ((-2147483648 != 42) || (3.14159 >= ((v5 - (float64(3.14159) * (0.1 * 3.14159))) + 0.1)))) {
                            var v19 int16 = 1000
                            _ = v19
                        }
                    default:
                        var v20 int16 = 42
                        _ = v20
                    }
                }
                v5 = 1.5
                if (false && (152 <= 9223372036854775807)) {
                    var v21 float64 = float64(0.0)
                    switch fi18 {
                    case -2147483648:
                        var v22 int64 = 177
                        _ = v22
                    }
                    if false {
                        if false {
                            v5 += (v21 + 0.1)
                            if true {
                                fmt.Println(v6, v21, v5)
                                v6 -= (v6 % (((v6 << ((4294967295) & 63))) | 1))
                            } else {
                                if false {
                                    v6 += v6
                                } else {
                                    fmt.Println(fi18, v21)
                                }
                                if false {
                                    var v23 uint8 = 239
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v21 = float64(((0.1 / float32(float64(((3.14159 * 0.1) + (-2.25 + (float32(v5) / 1.5)))))) * 1.5))
                                    _ = v23
                                }
                            }
                        }
                        v24 := GenS3{F0: 42, F1: 42}
                        if (true && true) {
                            for fi25 := 0; fi25 < 1; fi25++ {
                                var v26 int64 = 9223372036854775807
                                if (("foo bar" + ("\n\t" + "")) > "hello") {
                                    if false {
                                        v21 = 3.14159
                                        v21 += (-2.25 - v21)
                                    }
                                }
                                if false {
                                    v27 := GenS2{F0: 1, F1: (v6 + v6), F2: 532}
                                    _ = v27
                                } else {
                                    v26 = -823
                                }
                                _ = fi25
                                _ = v26
                            }
                            fmt.Println(v21)
                            fmt.Println(v6)
                        } else {
                            _ = len(v2)
                            var v28 float64 = (v21 * v5)
                            for fi29 := 0; fi29 < 10; fi29++ {
                                v30 := GenS4{F0: false, F1: 32767}
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        }
                        _ = v24
                    }
                    _ = v21
                } else {
                    v4[65535] = true
                    for fi31 := 0; fi31 < 16; fi31++ {
                        switch "\n\t" {
                        case "\n\t":
                            v32 := GenS1{F0: 38}
                            _ = v32
                        case "hello":
                            fmt.Println(v5, v7, fi31)
                        case "":
                            fmt.Println(v5, fi18)
                        default:
                            fmt.Println(v7)
                        }
                        if (!true) {
                            var v33 int = -301
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v33
                        } else {
                            v34 := GenS5{F0: "", F1: fi18}
                            v5 = 3.14159
                            _ = v34
                        }
                        _ = fi31
                    }
                    v35 := GenS1{F0: 32767}
                    _ = v35
                }
            }
        }
        _ = fi18
    }
    var v36 bool = (42 == -59)
    var v37 uint64 = 18446744073709551615
    v5 = (((v5 / v5) / 1.0) * (v5 * v5))
    v37 = v37
    v38 := v4[65535]
    var v39 float32 = -2.25
    switch "foo bar" {
    case "foo bar":
        if (-9223372036854775808 == 9223372036854775807) {
            var v40 uint64 = v37
            var v41 uint8 = 128
            _ = v40
            _ = v41
        } else {
            v42 := GenS2{F0: 778, F1: -43, F2: 212}
            for fi43 := 0; fi43 < 20; fi43++ {
                delete(v1, -231)
                for fi44 := 0; fi44 < 4; fi44++ {
                    var v45 int32 = -2147483648
                    var v46 int = (v7 & (fi43 % ((-2147483648) | 1)))
                    if false {
                        if v38 {
                            v47 := GenS1{F0: 0}
                            if false {
                                switch "a" {
                                case "\n\t":
                                    if v38 {
                                        if (!((!((42 != 1) || v36)) || true)) {
                                            v1[v7] = ""
                                            v2 = append(v2, "foo bar")
                                            delete(v1, fi44)
                                            var v48 float32 = v39
                                            fmt.Println(fi44, v37)
                                            _ = v48
                                        }
                                    }
                                default:
                                    var v49 float64 = v5
                                    _ = v49
                                }
                            } else {
                                var v50 int64 = 9223372036854775807
                                _ = v50
                            }
                            _ = v47
                        }
                        v51 := GenS6{F0: 255, F1: 128}
                        _ = v51
                    } else {
                        v52 := GenS3{F0: 2, F1: 42}
                        var v53 int = fi43
                        _ = v52
                        _ = v53
                    }
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                var v54 float64 = v5
                _ = fi43
                _ = v54
            }
            _ = v42
        }
    case "\n\t":
        if true {
            for fi55 := 0; fi55 < 11; fi55++ {
                v56 := GenS7{F0: 2, F1: 127}
                var v57 uint64 = 706
                _ = fi55
                _ = v56
                _ = v57
            }
        }
    default:
        switch v7 {
        case 9223372036854775807:
            v58 := GenS7{F0: 284, F1: 128}
            _ = v58
        case 42:
            for fi59 := 0; fi59 < 6; fi59++ {
                for fi60 := 0; fi60 < 14; fi60++ {
                    fmt.Println(fi59)
                    v36 = v38
                    _ = fi60
                }
                v38 = v38
                _ = fi59
            }
        }
    }
    v61 := v1[(^(v7))]
    var v62 int8 = v6
    switch (v61 + (v61 + v61)) {
    case "\n\t":
        v63 := GenS8{F0: 65535, F1: (v7 >> ((65535) & 63))}
        _ = v63
    case "hello":
        switch v61 {
        case "foo bar":
            var v64 int64 = -1
            _ = v64
        }
    case "a":
        v36 = false
    default:
        fmt.Println(v37)
    }
    switch v61 {
    case "hello":
        var v65 int = 9223372036854775807
        _ = v65
    default:
        for fi66 := 0; fi66 < 2; fi66++ {
            v38 = false
            _ = fi66
        }
    }
    var v67 uint32 = 638
    var v68 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v67
    _ = v68
    _ = v7
}
