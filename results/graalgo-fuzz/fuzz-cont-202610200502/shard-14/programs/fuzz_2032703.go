package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int8 }
type GenS3 struct { F0 string; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint16; F1 int16; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int16; F1 int16; F2 uint64 }

func main() {
    v1 := []int{}
    var v2 int8 = 0
    var v3 string = "\n\t"
    var v4 uint16 = 123
    v5 := make(map[uint32]int)
    v3 = v3
    if true {
        fmt.Println(v2)
        v6 := GenS1{F0: "hello", F1: (true || false), F2: v4}
        v1 = append(v1, -126)
        _ = v6
    }
    if true {
        v7 := GenS2{F0: (v2 + v2), F1: 2, F2: -1}
        _ = v7
    }
    if ((!false) && false) {
        v4 = 562
        v8 := GenS3{F0: v3, F1: 0, F2: v2}
        v9 := GenS1{F0: (v3 + "\n\t"), F1: false, F2: v4}
        _ = v8
        _ = v9
    }
    v1 = append(v1, -1)
    if (!false) {
        v5[449] = -9223372036854775808
        switch (v3 + (v3 + "foo bar")) {
        case "foo bar":
            fmt.Println(v2, v4, v4)
        case "hello":
            if (0 > 0) {
                var v10 uint = 42
                var v11 bool = false
                v2 = (v2 >> ((uint(v4)) & 63))
                _ = v10
                _ = v11
            } else {
                if false {
                    fmt.Println(v4)
                    v12 := GenS4{F0: v4, F1: 32767, F2: true}
                    var v13 uint64 = 767
                    _ = v12
                    _ = v13
                } else {
                    if false {
                        var v14 uint = 767
                        var v15 int8 = (v2 >> (((^(v14))) & 63))
                        v4 = (^(v4))
                        _ = v14
                        _ = v15
                    } else {
                        if true {
                            switch v3 {
                            case "hello":
                                fmt.Println(v4, v3, v4)
                            case "a":
                                switch 2147483647 {
                                case 9223372036854775807:
                                    fmt.Println(v2, v2, v4)
                                }
                            case "foo bar":
                                v16 := v5[2147483647]
                                _ = v16
                            }
                            delete(v5, 65535)
                        } else {
                            v17 := v5[739]
                            _ = v17
                        }
                        v2 -= 85
                        fmt.Println(v4)
                    }
                    if ((9223372036854775807 == -800) && false) {
                        v4 -= v4
                        fmt.Println(v4, v3, v4)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = len(v1)
                }
                v18 := GenS3{F0: "\n\t", F1: 568, F2: (v2 * v2)}
                _ = v18
            }
        default:
            switch -1 {
            case 2147483647:
                v19 := GenS3{F0: (("foo bar" + "\n\t") + v3), F1: 303, F2: 0}
                _ = v19
            case 0:
                fmt.Println(v2, v2, v3)
            case -211:
                _ = len(v5)
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    switch (v3 + "hello") {
    case "":
        v2 = v2
    }
    v20 := GenS4{F0: (^(v4)), F1: -1, F2: true}
    if true {
        v3 = v3
    } else {
        v21 := GenS1{F0: (v3 + v3), F1: (((v4 >> ((0) & 63)) > v4) && (0 != 486)), F2: 32767}
        _ = v21
    }
    v22 := GenS4{F0: v4, F1: 170, F2: true}
    v4 = v4
    var v23 uint32 = 0
    var v24 bool = false
    for fi25 := 0; fi25 < 11; fi25++ {
        switch (^(fi25)) {
        case 272:
            var v26 uint32 = v23
            _ = v26
        default:
            if (311 == 1000) {
                v27 := GenS2{F0: v2, F1: 255, F2: 42}
                switch (v3 + "foo bar") {
                case "\n\t":
                    v24 = (1000 == -48)
                case "":
                    delete(v5, (v23 >> ((234) & 63)))
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v27
            } else {
                v2 += 65
                v28 := GenS4{F0: v4, F1: 0, F2: ((108 == 128) && v24)}
                _ = v28
            }
        }
        _ = fi25
    }
    var v29 int = 4294967295
    for fi30 := 0; fi30 < 6; fi30++ {
        if (true && v24) {
            v31 := GenS4{F0: v4, F1: 1, F2: false}
            v29 += (-(v29))
            switch v3 {
            case "":
                switch (v29 * v29) {
                case 42:
                    _ = len(v5)
                case 0:
                    if v24 {
                        if false {
                            fmt.Println(v29)
                        } else {
                            v32 := v5[(v23 / ((v23) | 1))]
                            var v33 float32 = 0.0
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS4{F0: v4, F1: 42, F2: v24}
                        _ = v34
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = (^(v29))
                    }
                }
            case "a":
                if (9223372036854775807 != -9223372036854775808) {
                    if true {
                        var v35 uint64 = 42
                        _ = v35
                    }
                }
            }
            _ = v31
        } else {
            var v36 uint16 = v4
            _ = v36
        }
        if v24 {
            v3 = (v3 + v3)
            v37 := GenS2{F0: -128, F1: 161, F2: 1}
            v3 = v3
            _ = v37
        }
        _ = fi30
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        for fi39 := 0; fi39 < 2; fi39++ {
            v40 := v5[2147483648]
            if v24 {
                switch "a" {
                case "hello":
                    v23 += 0
                default:
                    if true {
                        fmt.Println(v40, fi39)
                        var v41 uint32 = v23
                        _ = v41
                    } else {
                        if (!((1 > 2) && (float64(float32(-2.25)) <= (1e10 * 0.0)))) {
                            if (v3 != "a") {
                                v42 := GenS2{F0: (^(v2)), F1: 45, F2: (v2 - v2)}
                                v4 -= (v4 % ((v4) | 1))
                                if false {
                                    v2 = 127
                                    v3 = (("hello" + ((v3 + (v3 + v3)) + "a")) + v3)
                                } else {
                                    if v24 {
                                        delete(v5, v23)
                                        if ((v24 || true) || false) {
                                            var v43 uint8 = 151
                                            fmt.Println(v43, v29)
                                            var v44 int = v29
                                            var v45 int16 = -1
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        } else {
                                            var v46 uint64 = 940
                                            _ = v46
                                        }
                                        var v47 uint64 = 2
                                        _ = v47
                                    }
                                    if (!(v24 && v24)) {
                                        v3 = ""
                                        if (false || ((fi38 ^ v29) < (fi39 >> ((0) & 63)))) {
                                            var v48 uint64 = 4294967295
                                            v49 := GenS5{F0: 0}
                                            v50 := GenS6{F0: -572, F1: uint8(v4)}
                                            var v51 uint32 = v23
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                    } else {
                                        v4 -= v4
                                        if (1 != 42) {
                                            var v52 float64 = (-2.25 / -2.25)
                                            switch "hello" {
                                            case "\n\t":
                                                switch v40 {
                                                case 9223372036854775807:
                                                    v53 := GenS7{F0: int16(v2), F1: 699, F2: 4294967295}
                                                    _ = v53
                                                case 420:
                                                    fmt.Println(fi38, v23, v4)
                                                case -1:
                                                    var v54 bool = v24
                                                    _ = v54
                                                }
                                            case "hello":
                                                fmt.Println(v29)
                                            }
                                            _ = v52
                                        }
                                    }
                                }
                                _ = v42
                            } else {
                                var v55 string = (v3 + (v3 + (("" + (v3 + v3)) + "\n\t")))
                                _ = v55
                            }
                            v56 := v5[(^(v23))]
                            _ = v56
                        }
                    }
                }
                var v57 uint32 = (^(v23))
                v58 := GenS4{F0: 312, F1: 654, F2: v24}
                var v59 float64 = (-2.25 - 0.0)
                fmt.Println(v59, v59, v2)
                _ = v57
                _ = v58
                _ = v59
            } else {
                _ = len(v5)
                v3 = v3
                v40 = fi38
            }
            _ = fi39
            _ = v40
        }
        switch v3 {
        case "hello":
            if (1 < -1) {
                if (v24 || ((v29 % (((fi38 + fi38)) | 1)) < fi38)) {
                    fmt.Println(v29)
                    _ = len(v5)
                }
                v1 = append(v1, -1)
                fmt.Println(fi38, v4)
            }
        case "a":
            var v60 int = v29
            _ = v60
        case "":
            for fi61 := 0; fi61 < 6; fi61++ {
                switch (-(fi61)) {
                case 1:
                    switch int(v4) {
                    case -9223372036854775808:
                        var v62 uint32 = v23
                        _ = v62
                    default:
                        fmt.Println(fi38, fi38)
                    }
                }
                _ = len(v1)
                v23 -= v23
                _ = fi61
            }
        }
        v2 = v2
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
