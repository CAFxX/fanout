package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int64; F1 int; F2 uint32 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 int }
type GenS9 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint) uint8 {
    p0 = (^(p0))
    if false {
        var v5 string = ""
        var v6 bool = (140 != 255)
        v5 = "a"
        _ = v5
        _ = v6
    } else {
        fmt.Println(p0)
        v7 := GenS1{F0: 1, F1: 2, F2: 575}
        _ = v7
    }
    if (-1 <= -128) {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint = p0
            v10 := GenS1{F0: -843, F1: 32767, F2: 512}
            if false {
                v9 = 836
            } else {
                switch (fi8 >> (((^(p0))) & 63)) {
                case 127:
                    v11 := GenS2{F0: 0, F1: 919}
                    _ = v11
                default:
                    v9 = v9
                }
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                p0 += (^(p0))
                var v14 float64 = 0.0
                p0 = (p0 | 1)
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := []int{fi15, (fi15 >> (((p0 << ((p0) & 63))) & 63))}
                v17 := GenS3{F0: 128}
                if (true || true) {
                    var v18 uint8 = 68
                    var v19 uint64 = 4294967295
                    var v20 bool = (true || true)
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi12
        }
        v21 := GenS4{F0: 255, F1: 42}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 5; fi22++ {
            switch (fi22 << (((^(p0))) & 63)) {
            case 9223372036854775807:
                var v23 bool = true
                _ = v23
            case -187:
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 uint = p0
                    if (0 != 127) {
                        v25 = (v25 << ((1) & 63))
                        var v26 int16 = -1
                        if ((84 > 128) && false) {
                            switch 9223372036854775807 {
                            case 978:
                                v25 = v25
                            case 9223372036854775807:
                                fmt.Println(fi22, fi24)
                                if true {
                                    v27 := []int{}
                                    fmt.Println(fi22, v26)
                                    v26 += 0
                                    _ = v27
                                }
                                v28 := []int{32767, (fi24 * 42)}
                                if false {
                                    var v29 float32 = 0.1
                                    if false {
                                        v29 -= ((v29 + 3.14159) / 1e10)
                                    }
                                    _ = v29
                                }
                                _ = v28
                            case 127:
                                v30 := GenS1{F0: -1, F1: 675, F2: 0}
                                _ = v30
                            }
                            var v31 int16 = v26
                            v26 = v26
                            _ = v31
                        }
                        _ = v26
                    }
                    if false {
                        v32 := GenS5{F0: 1, F1: fi22, F2: 21}
                        _ = v32
                    } else {
                        switch "" {
                        case "foo bar":
                            var v33 int32 = 647
                            _ = v33
                        case "hello":
                            var v34 float32 = ((1.5 / 1.0) * 1.5)
                            _ = v34
                        case "":
                            p0 -= p0
                        }
                    }
                    _ = fi24
                    _ = v25
                }
            }
            fmt.Println(p0)
            v35 := GenS1{F0: 1, F1: 1, F2: 338}
            _ = fi22
            _ = v35
        }
        if true {
            p0 -= (p0 % (((p0 / (((p0 & p0)) | 1))) | 1))
            var v36 bool = ((((860 < 42) && (2147483647 == 2147483647)) && false) && true)
            fmt.Println(p0, v36)
            _ = v36
        }
        var v37 int64 = 676
        _ = v37
    }
    v38 := []int{42, 0, 583}
    _ = p0
    _ = v38
    return 88
}

func main() {
    var v1 float64 = float64(0.1)
    var v2 float32 = 0.1
    var v3 int = -954
    v4 := make(map[int]bool)
    var v39 int64 = -359
    switch (v3 * (v3 >> ((360) & 63))) {
    case 127:
        v2 = (float32((1.5 - (((1e10 * float64(0.1)) / (1e10 + (((v1 - ((((v1 * float64(v2)) + float64(v2)) * (v1 - float64(v2))) - (-2.25 * 3.14159))) / 3.14159) * v1))) / 3.14159))) * (-2.25 + v2))
    case 0:
        v40 := GenS4{F0: 2, F1: -2147483648}
        _ = v40
    case -2147483648:
        delete(v4, (-(v3)))
    }
    v3 = (v3 ^ 643)
    v2 = (0.1 + float32(v1))
    var v41 uint32 = 2
    v3 += 42
    var v42 int32 = 2147483647
    for fi43 := 0; fi43 < 14; fi43++ {
        v44 := []int{0, (v3 * v3), (v3 + fi43)}
        _ = fi43
        _ = v44
    }
    switch v3 {
    case -161:
        var v45 float64 = (v1 * v1)
        _ = v45
    default:
        v46 := GenS6{F0: "hello", F1: 202}
        _ = v46
    }
    v41 = v41
    if false {
        if (2 == 210) {
            v2 = v2
            v47 := GenS7{F0: 0}
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        var v49 uint8 = 0
        _ = fi48
        _ = v49
    }
    v50 := GenS8{F0: 32, F1: v3}
    for fi51 := 0; fi51 < 3; fi51++ {
        var v52 int64 = -1
        fmt.Println(v3)
        if false {
            var v53 int16 = -1
            if (1 >= 86) {
                v54 := GenS5{F0: 1, F1: v3, F2: 916}
                v42 = v42
                v55 := []int{(fi51 | (v3 / ((v3) | 1))), 4294967295}
                _ = v54
                _ = v55
            } else {
                _ = len(v4)
                var v56 int8 = 1
                var v57 int8 = (v56 * (v56 / ((v56) | 1)))
                _ = v56
                _ = v57
            }
            if true {
                v58 := GenS2{F0: 42, F1: v42}
                if true {
                    for fi59 := 0; fi59 < 9; fi59++ {
                        v42 = (v42 | 1)
                        v3 = (fi51 >> ((4294967295) & 63))
                        switch ((("" + ((("foo bar" + "foo bar") + "a") + ((((("foo bar" + "foo bar") + ("\n\t" + "a")) + "a") + "a") + ((("\n\t" + "\n\t") + "\n\t") + "")))) + "\n\t") + ("a" + "foo bar")) {
                        case "foo bar":
                            switch ("a" + "hello") {
                            case "":
                                var v60 uint32 = v41
                                _ = v60
                            case "a":
                                delete(v4, (v3 + v3))
                            default:
                                v61 := GenS8{F0: 42, F1: fi51}
                                _ = v61
                            }
                        case "\n\t":
                            var v62 uint8 = 42
                            _ = v62
                        case "hello":
                            var v63 uint32 = (^(v41))
                            _ = v63
                        default:
                            v64 := GenS5{F0: (^(v39)), F1: 0, F2: v41}
                            _ = v64
                        }
                        _ = fi59
                    }
                    if true {
                        var v65 uint = 1
                        _ = v65
                    } else {
                        v66 := GenS1{F0: (v42 - v42), F1: 2, F2: 1}
                        v3 -= 455
                        for fi67 := 0; fi67 < 20; fi67++ {
                            if true {
                                var v68 string = ("hello" + ("hello" + (("a" + ("" + "\n\t")) + "foo bar")))
                                _ = v68
                            } else {
                                v69 := GenS9{F0: 125, F1: 182}
                                v39 = (^(v39))
                                _ = v69
                            }
                            _ = fi67
                        }
                        _ = v66
                    }
                    fmt.Println(v41, v52)
                }
                _ = v58
            } else {
                var v70 float32 = v2
                switch fi51 {
                case -9223372036854775808:
                    if true {
                        v71 := GenS5{F0: v52, F1: v3, F2: v41}
                        _ = v71
                    } else {
                        v2 += -2.25
                        var v72 int32 = v42
                        _ = v72
                    }
                case 642:
                    v73 := []int{v3}
                    _ = v73
                }
                _ = v70
            }
            _ = v53
        } else {
            v74 := GenS8{F0: 100, F1: (fi51 | (v3 * v3))}
            v75 := []int{(v3 + 9223372036854775807), fi51, (fi51 % ((v3) | 1))}
            switch (v3 * -776) {
            case -1:
                fmt.Println(v39)
            }
            _ = v74
            _ = v75
        }
        _ = fi51
        _ = v52
    }
    v1 = (-2.25 + v1)
    v39 = (v39 >> ((uint(v41)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v50
}
