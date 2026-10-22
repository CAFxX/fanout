package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 string; F1 string }

func main() {
    var v1 int16 = 944
    v2 := make(map[int]string)
    var v3 float32 = (1.5 / -2.25)
    var v4 int = 322
    var v5 int = v4
    switch (v4 ^ (-(v5))) {
    case 659:
        if true {
            if (!(218 != 249)) {
                var v6 uint32 = 779
                _ = v6
            }
            if false {
                for fi7 := 0; fi7 < 3; fi7++ {
                    var v8 uint16 = 0
                    _ = fi7
                    _ = v8
                }
                fmt.Println(v1)
                var v9 int8 = 100
                _ = v9
            } else {
                fmt.Println(v1, v1, v1)
            }
            v10 := GenS1{F0: 10}
            _ = v10
        }
    case 0:
        if false {
            fmt.Println(v1)
            for fi11 := 0; fi11 < 2; fi11++ {
                v5 = fi11
                if (32767 == 32767) {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        fmt.Println(v1)
                        _ = fi12
                    }
                } else {
                    var v13 int64 = -285
                    _ = v13
                }
                _ = fi11
            }
            var v14 string = "foo bar"
            _ = v14
        } else {
            var v15 int = (-(v4))
            fmt.Println(v1, v5)
            for fi16 := 0; fi16 < 13; fi16++ {
                fmt.Println(v5, v4, v1)
                switch "\n\t" {
                case "hello":
                    var v17 int8 = -128
                    _ = v17
                case "a":
                    v18 := GenS2{F0: "hello", F1: 2147483647}
                    _ = v18
                case "\n\t":
                    if true {
                        var v19 int = (v4 | 32768)
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v19 += 32767
                            v21 := v2[fi16]
                            var v22 uint64 = 539
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        var v23 int64 = 42
                        _ = v19
                        _ = v23
                    } else {
                        var v24 float64 = 3.14159
                        if false {
                            var v25 float32 = 0.0
                            _ = v25
                        }
                        var v26 uint32 = 2
                        _ = v24
                        _ = v26
                    }
                default:
                    fmt.Println(fi16)
                }
                var v27 uint16 = 529
                _ = fi16
                _ = v27
            }
            _ = v15
        }
    case 32768:
        fmt.Println(v1)
    default:
        var v28 int8 = 127
        _ = v28
    }
    switch "" {
    case "":
        if (255 < 812) {
            for fi29 := 0; fi29 < 14; fi29++ {
                switch fi29 {
                case -513:
                    if true {
                        v30 := GenS3{F0: 2, F1: -374}
                        var v31 int16 = 575
                        var v32 int16 = 42
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                case -211:
                    v33 := []int{}
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v35 := GenS4{F0: 1}
                        fmt.Println(v3, v3, v4)
                        v36 := []int{(fi29 & v4), -2147483648, (-(fi34))}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
                v37 := GenS5{F0: (fi29 * v5), F1: 0, F2: 42}
                _ = fi29
                _ = v37
            }
        }
    }
    switch "a" {
    case "hello":
        fmt.Println(v1, v1)
    case "a":
        v38 := GenS1{F0: 42}
        _ = v38
    }
    var v39 int32 = 2147483647
    if true {
        v5 = v4
        var v40 uint16 = 32767
        _ = v40
    } else {
        v5 = (^(v4))
    }
    v39 = (-(v39))
    v41 := GenS2{F0: "", F1: (^(v39))}
    delete(v2, v4)
    fmt.Println(v4, v5, v39)
    v39 += int32(v4)
    if true {
        v42 := GenS1{F0: 0}
        if (9223372036854775807 >= -489) {
            fmt.Println(v3, v5, v3)
            var v43 int = v5
            _ = v43
        }
        _ = v42
    }
    var v44 uint16 = 976
    if false {
        v45 := GenS1{F0: 127}
        _ = v45
    }
    v44 = v44
    var v46 int16 = -1
    if ((!(true || (103 != 42))) && (v3 < float32(float64((float32(1e10) - float32(-2.25)))))) {
        for fi47 := 0; fi47 < 20; fi47++ {
            var v48 bool = false
            if (v48 && v48) {
                var v49 float64 = ((-2.25 / 1.5) + (0.0 + 3.14159))
                v44 = v44
                _ = v49
            }
            _ = fi47
            _ = v48
        }
    } else {
        v50 := v2[(v5 % ((-1) | 1))]
        if (-2.25 != 3.14159) {
            v51 := []int{(v4 >> ((311) & 63))}
            for fi52 := 0; fi52 < 19; fi52++ {
                for fi53 := 0; fi53 < 15; fi53++ {
                    v51 = append(v51, (v5 / (((fi53 - v4)) | 1)))
                    if (0.1 <= float32((1e10 * ((0.1 - (3.14159 - 0.0)) + 1e10)))) {
                        v54 := v2[9223372036854775807]
                        delete(v2, (fi52 * v5))
                        v39 = v39
                        _ = v54
                    }
                    _ = fi53
                }
                var v55 uint8 = 6
                _ = fi52
                _ = v55
            }
            switch -2147483648 {
            case 758:
                v3 = v3
            case 298:
                for fi56 := 0; fi56 < 8; fi56++ {
                    fmt.Println(v44, v4)
                    v57 := GenS5{F0: fi56, F1: 1, F2: (v39 + v39)}
                    for fi58 := 0; fi58 < 16; fi58++ {
                        var v59 uint = 0
                        if false {
                            var v60 int8 = 114
                            _ = v60
                        } else {
                            if true {
                                delete(v2, 32768)
                                switch (v4 * (v5 - 1)) {
                                case 4294967295:
                                    v51 = append(v51, fi58)
                                case 828:
                                    var v61 int64 = -1
                                    v1 = v1
                                    _ = v61
                                case 0:
                                    fmt.Println(v59, v39, v4)
                                default:
                                    if (0 != 147) {
                                        var v62 uint32 = 572
                                        v62 += 4294967295
                                        _ = v62
                                    } else {
                                        v63 := GenS6{F0: (v59 ^ 2), F1: "a"}
                                        v44 -= (v44 / ((v44) | 1))
                                        v64 := GenS7{F0: 42, F1: 127}
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                if true {
                                    v44 -= v44
                                }
                            } else {
                                if true {
                                    fmt.Println(v44, v3)
                                    if false {
                                        v44 = v44
                                    } else {
                                        fmt.Println(v59)
                                        v65 := GenS5{F0: fi56, F1: 255, F2: v39}
                                        v50 = "hello"
                                        v44 -= 295
                                        _ = v65
                                    }
                                    switch v50 {
                                    case "a":
                                        v66 := GenS8{F0: v50, F1: v50}
                                        _ = v66
                                    default:
                                        fmt.Println(v3, v39)
                                    }
                                    if ((120 != 127) && (v44 >= (v44 % ((v44) | 1)))) {
                                        v67 := GenS3{F0: 53, F1: 1}
                                        if len(v51) > 0 {
                                            v51[0] = (fi56 % ((-1) | 1))
                                        }
                                        _ = v67
                                    } else {
                                        switch "" {
                                        case "a":
                                            v68 := GenS8{F0: "hello", F1: "hello"}
                                            _ = v68
                                        case "":
                                            if (2 >= uint8(v59)) {
                                                v39 += 2147483647
                                                var v69 uint64 = 4294967295
                                                _ = v69
                                            }
                                        default:
                                            var v70 float32 = float32(3.14159)
                                            _ = v70
                                        }
                                    }
                                    v44 = (v44 - v44)
                                }
                            }
                            v71 := GenS8{F0: "hello", F1: ("\n\t" + ("foo bar" + v50))}
                            _ = v71
                        }
                        _ = fi58
                        _ = v59
                    }
                    _ = fi56
                    _ = v57
                }
            }
            _ = v51
        }
        _ = v50
    }
    var v72 int = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v44
    _ = v46
    _ = v5
    _ = v72
}
