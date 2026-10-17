package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 int }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int8 }
func genFn0() int32 {
    v6 := GenS1{F0: 0, F1: 2147483648}
    var v7 uint = 0
    var v8 uint32 = 4294967295
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 bool = false
    var v5 int32 = 197
    v5 -= (v5 << ((766) & 63))
    switch "foo bar" {
    case "hello":
        v5 += v5
    case "\n\t":
        fmt.Println(v4, v4)
    case "a":
        switch -9223372036854775808 {
        case 32767:
            v5 = (-(v5))
        default:
            var v9 int = 32767
            _ = v9
        }
    default:
        if v4 {
            var v10 float64 = (float64(1.5) + 3.14159)
            switch 42 {
            case 2147483647:
                v11 := GenS2{F0: 359, F1: 962, F2: 32768}
                _ = v11
            case 9223372036854775807:
                v12 := GenS2{F0: 0, F1: 754, F2: 1}
                _ = v12
            case -1:
                var v13 bool = (v4 && (255 == 512))
                _ = v13
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                if v4 {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 uint8 = 1
                        v16 += v16
                        v4 = v4
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS3{F0: "", F1: "\n\t", F2: (fi14 / (((fi14 | -9223372036854775808)) | 1))}
                    v18 := GenS3{F0: ("hello" + "\n\t"), F1: "foo bar", F2: 127}
                    _ = v17
                    _ = v18
                }
                switch fi14 {
                case 4294967295:
                    v5 = (^(v5))
                case 128:
                    v4 = v4
                case -441:
                    var v19 uint64 = 18446744073709551615
                    _ = v19
                default:
                    fmt.Println(v4, v10, v5)
                }
                _ = fi14
            }
            _ = v10
        } else {
            v20 := GenS2{F0: 2147483648, F1: 398, F2: 2147483647}
            v21 := GenS2{F0: 2, F1: -469, F2: 0}
            for fi22 := 0; fi22 < 11; fi22++ {
                fmt.Println(fi22, v5)
                _ = fi22
            }
            _ = v20
            _ = v21
        }
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 int64 = 0
            _ = fi24
            _ = v25
        }
        v26 := GenS4{F0: 0, F1: 127}
        for fi27 := 0; fi27 < 20; fi27++ {
            if (v4 && (v4 && ((((float32(1e10) + 1e10) - 0.0) - (-2.25 / 0.1)) <= ((((3.14159 + -2.25) / (3.14159 * 1e10)) + (0.1 / ((1e10 + 3.14159) - 3.14159))) / 1e10)))) {
                v5 += v5
                var v28 uint16 = 311
                _ = v28
            } else {
                v29 := []int{(fi27 & (fi23 + (fi27 & fi23))), fi23, (fi23 | (fi23 | fi23)), fi23}
                _ = v29
            }
            _ = fi27
        }
        _ = fi23
        _ = v26
    }
    v1[""] = 1
    if v4 {
        if (!v4) {
            v5 -= (v5 / ((-1) | 1))
        }
        var v30 uint = 634
        _ = v30
    } else {
        if true {
            if (42 == 746) {
                for fi31 := 0; fi31 < 5; fi31++ {
                    v5 += genFn0()
                    _ = fi31
                }
                for fi32 := 0; fi32 < 3; fi32++ {
                    v5 += (v5 >> ((42) & 63))
                    if (247 == 18446744073709551615) {
                        for fi33 := 0; fi33 < 11; fi33++ {
                            switch fi32 {
                            case -1:
                                if (65535 <= 1) {
                                    if true {
                                        var v34 string = "foo bar"
                                        _ = v34
                                    } else {
                                        var v35 uint32 = 0
                                        if (255 == 42) {
                                            fmt.Println(v5, v4)
                                        } else {
                                            var v36 int64 = 9223372036854775807
                                            v5 += v5
                                            switch fi33 {
                                            case 4294967295:
                                                v37 := GenS5{F0: 2}
                                                _ = v37
                                            default:
                                                var v38 uint64 = 2
                                                _ = v38
                                            }
                                            _ = v36
                                        }
                                        if v4 {
                                            v39 := []int{(fi33 + fi33), -742, (fi32 ^ fi32)}
                                            if v4 {
                                                var v40 uint32 = (v35 % (((^(v35))) | 1))
                                                delete(v3, fi33)
                                                if v4 {
                                                    var v41 string = ""
                                                    _ = v41
                                                }
                                                var v42 uint16 = 494
                                                _ = v40
                                                _ = v42
                                            }
                                            _ = v39
                                        }
                                        _ = v35
                                    }
                                }
                            default:
                                var v43 uint16 = 0
                                _ = v43
                            }
                            _ = fi33
                        }
                        var v44 uint8 = 234
                        var v45 uint16 = 255
                        _ = v44
                        _ = v45
                    } else {
                        _ = len(v3)
                        v4 = v4
                    }
                    _ = fi32
                }
                if v4 {
                    var v46 uint = 2
                    _ = v46
                }
            } else {
                var v47 float64 = 3.14159
                v48 := GenS6{F0: 712, F1: 0}
                _ = v47
                _ = v48
            }
            var v49 int8 = 22
            v49 += (v49 / (((v49 * v49)) | 1))
            _ = v49
        }
    }
    switch "a" {
    case "\n\t":
        v50 := []int{127, 17}
        _ = v50
    case "hello":
        fmt.Println(v5)
    case "a":
        var v51 float64 = 1e10
        _ = v51
    default:
        for fi52 := 0; fi52 < 12; fi52++ {
            var v53 int = fi52
            _ = fi52
            _ = v53
        }
    }
    for fi54 := 0; fi54 < 5; fi54++ {
        for fi55 := 0; fi55 < 10; fi55++ {
            v56 := GenS2{F0: 1, F1: (-(fi55)), F2: (fi55 >> ((748) & 63))}
            _ = fi55
            _ = v56
        }
        var v57 int = fi54
        _ = fi54
        _ = v57
    }
    var v58 uint64 = 550
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
}
