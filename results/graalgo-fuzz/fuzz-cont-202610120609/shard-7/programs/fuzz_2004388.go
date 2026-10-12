package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16 }

func main() {
    var v1 int64 = 92
    v2 := make(map[int]string)
    var v3 int32 = 673
    var v4 int64 = 0
    v5 := GenS1{F0: 2, F1: 114, F2: int(v1)}
    for fi6 := 0; fi6 < 6; fi6++ {
        if (-1 == -427) {
            if false {
                v4 += (v1 & 949)
                switch 127 {
                case 1:
                    fmt.Println(v3, v3)
                default:
                    var v7 int32 = v3
                    _ = v7
                }
                fmt.Println(fi6)
            }
        } else {
            v8 := GenS1{F0: 42, F1: -1, F2: 1}
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 int8 = 1
                var v11 bool = true
                switch fi6 {
                case 0:
                    v12 := GenS1{F0: 973, F1: v10, F2: (fi6 * int(v4))}
                    _ = v12
                default:
                    v4 = v4
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
        v3 += (v3 << ((255) & 63))
        v3 = (v3 ^ (v3 + -1))
        _ = fi6
    }
    v13 := GenS1{F0: 2147483647, F1: 127, F2: int(v1)}
    if false {
        if (!(0 >= 42)) {
            v14 := GenS1{F0: 4294967295, F1: -1, F2: 127}
            _ = v14
        } else {
            var v15 string = "hello"
            var v16 bool = (true || true)
            _ = v15
            _ = v16
        }
        v17 := GenS2{F0: 9223372036854775807, F1: 7}
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        if true {
            v19 := GenS1{F0: 4294967295, F1: 0, F2: fi18}
            v20 := GenS3{F0: "foo bar", F1: 42, F2: 104}
            v3 = (v3 * (v3 >> ((65535) & 63)))
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            v3 -= 1
            v22 := GenS3{F0: "foo bar", F1: 931, F2: -1}
            _ = fi21
            _ = v22
        }
        switch "" {
        case "a":
            v23 := GenS4{F0: 268, F1: 42, F2: 2}
            _ = v23
        case "":
            for fi24 := 0; fi24 < 6; fi24++ {
                if false {
                    var v25 float32 = 1e10
                    var v26 int = fi18
                    v27 := GenS5{F0: "foo bar", F1: -168}
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    if (!false) {
                        v4 += (v4 & 42)
                    } else {
                        var v28 int = (fi18 % (((^(fi24))) | 1))
                        if true {
                            var v29 int8 = -100
                            _ = v29
                        } else {
                            var v30 uint32 = 0
                            v31 := []int{}
                            if false {
                                var v32 int16 = 1000
                                fmt.Println(v1)
                                if (!false) {
                                    if ((!(!(v3 <= v3))) || (!false)) {
                                        if false {
                                            var v33 uint8 = 0
                                            _ = v33
                                        }
                                    }
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        switch ((((("foo bar" + "hello") + "hello") + "a") + ("a" + "\n\t")) + "") {
                                        case "":
                                            switch (^(v28)) {
                                            case -183:
                                                var v34 uint32 = (v30 - 2)
                                                _ = v34
                                            case 32768:
                                                v28 -= (v28 << ((1) & 63))
                                            }
                                        case "hello":
                                            v32 = v32
                                        }
                                    }
                                }
                                v35 := v2[(fi24 + (fi24 >> ((0) & 63)))]
                                _ = len(v2)
                                if ((1e10 >= float32(0.0)) || true) {
                                    if (0 >= 32767) {
                                        var v36 string = (v35 + v35)
                                        _ = v36
                                    }
                                    if false {
                                        var v37 int16 = v32
                                        _ = v37
                                    }
                                } else {
                                    fmt.Println(v4, v32)
                                }
                                _ = v32
                                _ = v35
                            }
                            _ = len(v31)
                            _ = v30
                            _ = v31
                        }
                        var v38 uint64 = 1
                        _ = v28
                        _ = v38
                    }
                    v39 := GenS3{F0: "a", F1: 1, F2: -128}
                    v4 = (v1 - -9223372036854775808)
                    v40 := GenS3{F0: "\n\t", F1: 194, F2: 1}
                    if ((!(!true)) && (!true)) {
                        if false {
                            var v41 int8 = -128
                            switch "foo bar" {
                            case "a":
                                v3 += -1
                            case "hello":
                                v42 := GenS5{F0: "\n\t", F1: v1}
                                _ = v42
                            default:
                                fmt.Println(v41)
                            }
                            v43 := []int{(fi18 % ((944) | 1))}
                            _ = v41
                            _ = v43
                        } else {
                            v44 := GenS6{F0: 1, F1: v3}
                            if ((340 > 886) || false) {
                                var v45 float64 = -2.25
                                v46 := GenS5{F0: (("hello" + "hello") + "hello"), F1: -942}
                                _ = v45
                                _ = v46
                            }
                            var v47 uint16 = 369
                            delete(v2, fi18)
                            var v48 uint64 = 325
                            v47 = (v47 & v47)
                            _ = v44
                            _ = v47
                            _ = v48
                        }
                    } else {
                        if (false && ((-(fi18)) <= fi24)) {
                            if true {
                                v4 = v1
                                v49 := []int{9223372036854775807, 127, (fi24 << ((0) & 63)), fi24}
                                if (!(2147483647 < -2147483648)) {
                                    v4 -= -460
                                }
                                _ = v49
                            }
                            v50 := GenS4{F0: 1, F1: 413, F2: 452}
                            v51 := GenS4{F0: -1, F1: 18446744073709551615, F2: 904}
                            _ = v50
                            _ = v51
                        }
                        var v52 uint16 = 1
                        _ = v52
                    }
                    _ = v39
                    _ = v40
                }
                v1 -= v4
                switch "foo bar" {
                case "a":
                    if false {
                        v1 += v4
                        var v53 float32 = 0.0
                        v53 = float32((-2.25 / 3.14159))
                        var v54 int = (^(fi24))
                        _ = v53
                        _ = v54
                    } else {
                        v55 := GenS7{F0: 42, F1: 2, F2: 0}
                        _ = v55
                    }
                }
                _ = fi24
            }
        }
        _ = fi18
    }
    if (float32((0.0 * 0.1)) >= 0.1) {
        var v56 bool = true
        var v57 float64 = 3.14159
        var v58 int32 = v3
        _ = v56
        _ = v57
        _ = v58
    } else {
        v59 := GenS8{F0: 42}
        if true {
            v4 += 0
        } else {
            fmt.Println(v4, v1)
            for fi60 := 0; fi60 < 1; fi60++ {
                v61 := GenS2{F0: int(v3), F1: 0}
                switch "hello" {
                case "":
                    if (429 != 508) {
                        var v62 uint32 = 42
                        _ = v62
                    }
                case "hello":
                    fmt.Println(fi60, v1)
                default:
                    v3 = 1
                }
                if false {
                    v3 = 1
                    if true {
                        v4 = v1
                        v2[fi60] = "a"
                        var v63 float64 = float64(((1e10 + 3.14159) - 1.5))
                        _ = v63
                    } else {
                        v4 += v1
                        if (false || false) {
                            v64 := GenS4{F0: -1, F1: 42, F2: 32767}
                            if (1 > 468) {
                                v1 = -9223372036854775808
                                v1 = 42
                            } else {
                                delete(v2, (fi60 % ((127) | 1)))
                            }
                            _ = v64
                        } else {
                            for fi65 := 0; fi65 < 11; fi65++ {
                                fmt.Println(v3, fi65, v1)
                                var v66 string = ("hello" + "foo bar")
                                _ = fi65
                                _ = v66
                            }
                            v67 := []int{}
                            _ = v67
                        }
                    }
                } else {
                    v1 = (v1 % ((42) | 1))
                    v1 += v1
                    var v68 uint16 = 32767
                    _ = v68
                }
                _ = fi60
                _ = v61
            }
            v69 := GenS9{F0: 65535}
            _ = v69
        }
        if (!false) {
            fmt.Println(v1, v4)
            var v70 uint64 = 735
            fmt.Println(v3)
            _ = v70
        }
        _ = v59
    }
    v4 -= (v4 / (((v4 + -1)) | 1))
    v1 += (v1 % ((v4) | 1))
    var v71 int = -128
    switch "hello" {
    case "":
        for fi72 := 0; fi72 < 8; fi72++ {
            var v73 bool = ((("hello" + "") + "") < "")
            v4 = -455
            v74 := GenS9{F0: 1}
            _ = fi72
            _ = v73
            _ = v74
        }
    }
    var v75 int32 = -2147483648
    v76 := GenS2{F0: 42, F1: 255}
    v77 := GenS1{F0: 42, F1: 127, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v75
    _ = v76
    _ = v77
}
