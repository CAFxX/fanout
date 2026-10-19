package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 int16; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    v3 := []string{("\n\t" + "foo bar")}
    v4 := make(map[uint32]string)
    var v5 uint = 65535
    v6 := GenS1{F0: (v5 % (((v5 - (v5 * (^(v5))))) | 1)), F1: 2147483647}
    v2 = v2
    if true {
        v7 := v1["a"]
        var v8 bool = (1e10 < 0.0)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        if (1 == 188) {
            v5 = 684
            v5 -= v5
        }
        v10 := GenS2{F0: 42, F1: -9223372036854775808, F2: 42}
        if ((-145 < 42) && true) {
            switch (fi9 & fi9) {
            case -9223372036854775808:
                for fi11 := 0; fi11 < 5; fi11++ {
                    fmt.Println(fi9, fi11)
                    _ = fi11
                }
            case -900:
                v2 = v2
            default:
                var v12 bool = v2
                _ = v12
            }
            fmt.Println(fi9, v2)
            var v13 int16 = 1
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    if true {
        for fi14 := 0; fi14 < 14; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS3{F0: 42}
                switch "foo bar" {
                case "foo bar":
                    var v17 float32 = 1.5
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    } else {
        v2 = v2
    }
    var v18 int64 = -9223372036854775808
    v2 = (((((1e10 + 1.5) - (0.1 - 1.5)) - 3.14159) < 0.1) || (!((-9223372036854775808 >= v18) && v2)))
    v19 := v1[("" + "foo bar")]
    if v2 {
        var v20 float32 = (-2.25 * (1e10 + ((0.1 / 1.0) / (-2.25 - (1.5 + 0.0)))))
        fmt.Println(v20, v2)
        _ = v20
    } else {
        v1[v19] = v19
        fmt.Println(v5)
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 string = (v19 + "foo bar")
                var v24 float32 = float32(1.5)
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
    }
    if v2 {
        v19 = v19
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 bool = v2
            if false {
                delete(v1, (v19 + ("\n\t" + v19)))
                v2 = (128 < (fi25 ^ fi25))
                v19 = "hello"
            }
            v27 := GenS1{F0: v5, F1: 42}
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 int64 = (v18 % ((987) | 1))
        _ = v28
    } else {
        v29 := GenS4{F0: 632, F1: 1}
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 float64 = -2.25
        for fi32 := 0; fi32 < 20; fi32++ {
            switch (("foo bar" + v19) + ((v19 + (v19 + v19)) + "a")) {
            case "hello":
                switch fi30 {
                case 32768:
                    v33 := GenS5{F0: 232, F1: v5, F2: 128}
                    _ = v33
                case -906:
                    _ = len(v3)
                }
            case "a":
                v19 = "a"
            case "\n\t":
                if (65535 >= 533) {
                    if false {
                        if v2 {
                            switch (v19 + v19) {
                            case "\n\t":
                                if ((v31 + v31) == v31) {
                                    v2 = ((!v2) && v2)
                                    if false {
                                        var v34 uint32 = 65535
                                        _ = v34
                                    } else {
                                        if (!v2) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v5, v2)
                                            case "a":
                                                if (true || v2) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            default:
                                                var v35 float64 = v31
                                                _ = v35
                                            }
                                        } else {
                                            var v36 string = "foo bar"
                                            _ = v36
                                        }
                                        _ = len(v3)
                                        v37 := GenS6{F0: 0}
                                        _ = v37
                                    }
                                    fmt.Println(v19, v18, v2)
                                    v18 = v18
                                    var v38 int = 500
                                    fmt.Println(fi30, fi32, v18)
                                    v39 := GenS4{F0: 0, F1: 1}
                                    _ = v38
                                    _ = v39
                                }
                            case "foo bar":
                                if ((false && true) && v2) {
                                    var v40 string = v19
                                    v18 += v18
                                    v41 := GenS6{F0: v18}
                                    _ = v40
                                    _ = v41
                                }
                            }
                            v42 := GenS5{F0: uint64(v5), F1: v5, F2: 255}
                            _ = v42
                        }
                        v43 := GenS2{F0: (^(v5)), F1: v18, F2: 42}
                        fmt.Println(v18, v19, v31)
                        switch (fi30 >> ((v5) & 63)) {
                        case -569:
                            v2 = false
                        case -436:
                            v19 = (((v19 + "") + "\n\t") + v19)
                        case 0:
                            fmt.Println(v19, fi32)
                        }
                        _ = v43
                    } else {
                        fmt.Println(fi32, fi32, v31)
                    }
                    var v44 float64 = (0.1 * -2.25)
                    _ = v44
                } else {
                    if v2 {
                        v45 := GenS5{F0: 4294967295, F1: v5, F2: 14}
                        v31 += ((float64(((((0.0 * 0.1) + (1e10 * 0.0)) / (float32(3.14159) - ((0.0 / -2.25) + float32(v31)))) / float32(v31))) * v31) / (1e10 * v31))
                        var v46 int8 = 100
                        _ = v45
                        _ = v46
                    }
                }
                var v47 int16 = 778
                switch ((v19 + "\n\t") + v19) {
                case "hello":
                    v48 := GenS7{F0: fi30}
                    _ = v48
                case "a":
                    if (100 < 103) {
                        v5 = 701
                        if ((0.0 / float32(v31)) != 0.1) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v49 int8 = 42
                        _ = v49
                    } else {
                        if (0 <= 2) {
                            if (128 == 128) {
                                var v50 int32 = -1
                                var v51 float64 = 0.0
                                _ = v50
                                _ = v51
                            }
                            v52 := GenS3{F0: 13}
                            if ("a" < "hello") {
                                var v53 uint16 = 2
                                _ = v53
                            }
                            _ = v52
                        } else {
                            if (4294967295 <= 776) {
                                v54 := GenS6{F0: v18}
                                if false {
                                    var v55 uint8 = 76
                                    _ = len(v3)
                                    var v56 uint16 = 1
                                    _ = v55
                                    _ = v56
                                }
                                _ = v54
                            } else {
                                fmt.Println(v31, v5)
                                if v2 {
                                    var v57 uint16 = 1
                                    _ = v57
                                }
                            }
                            var v58 int32 = -237
                            _ = v58
                        }
                        var v59 uint64 = 2
                        _ = v59
                    }
                case "foo bar":
                    var v60 uint = 255
                    _ = v60
                }
                _ = v47
            default:
                if v2 {
                    if (1 <= 2) {
                        v61 := GenS8{F0: 0, F1: 474, F2: v2}
                        _ = v61
                    } else {
                        v62 := v1[""]
                        if ((v18 << (((v5 + (v5 >> (((v5 & (v5 >> ((v5) & 63)))) & 63)))) & 63)) != 42) {
                            v63 := v1[("" + "foo bar")]
                            var v64 int8 = 42
                            delete(v1, v19)
                            _ = v63
                            _ = v64
                        }
                        _ = v62
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = v19
                    }
                    v5 = (v5 / ((v5) | 1))
                }
            }
            _ = fi32
        }
        var v65 int8 = -50
        _ = fi30
        _ = v31
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
