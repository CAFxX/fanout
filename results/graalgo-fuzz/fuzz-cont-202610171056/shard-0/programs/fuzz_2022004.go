package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 bool; F1 int; F2 int32 }
type GenS6 struct { F0 int; F1 uint16; F2 int32 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 uint16; F2 int32 }
func genFn0(p0 int64) bool {
    p0 -= p0
    var v7 float32 = ((1.5 / 1e10) - 0.0)
    _ = p0
    _ = v7
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    v3 := []uint32{}
    var v4 uint16 = 202
    v5 := make(map[string]bool)
    v6 := []uint32{2147483647, 42, 65535}
    var v8 string = ""
    var v9 uint32 = 1
    if (128 < 609) {
        v9 -= 2
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 uint32 = v9
            switch (fi10 / (((fi10 << ((343) & 63))) | 1)) {
            case 4294967295:
                if ((((1.5 - 0.1) - (0.0 * (-2.25 * (-2.25 * -2.25)))) + -2.25) >= float32((1e10 + 0.0))) {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        switch ("foo bar" + ("a" + v8)) {
                        case "foo bar":
                            var v13 int64 = 1
                            _ = v13
                        case "hello":
                            v8 = (v8 + v8)
                        case "\n\t":
                            if v2 {
                                if false {
                                    if true {
                                        fmt.Println(fi12)
                                        fmt.Println(v2, v4, v11)
                                    } else {
                                        if true {
                                            var v14 float64 = (0.1 * float64((float32(float64(((0.0 * 3.14159) - (0.1 + ((0.0 + 0.1) * 0.1))))) - (-2.25 * (3.14159 - 3.14159)))))
                                            _ = v14
                                        } else {
                                            if v2 {
                                                var v15 uint32 = 97
                                                var v16 int8 = 42
                                                _ = v15
                                                _ = v16
                                            } else {
                                                v17 := GenS1{F0: v8, F1: v11}
                                                switch v8 {
                                                case "\n\t":
                                                    v11 += v11
                                                case "foo bar":
                                                    v4 = v4
                                                }
                                                _ = v17
                                            }
                                            v18 := GenS1{F0: v8, F1: 661}
                                            _ = v18
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                    var v19 uint8 = 1
                                    v2 = (!(575 > -947))
                                    _ = v19
                                }
                                fmt.Println(v8, fi10)
                                v6 = append(v6, (v9 ^ (v11 | (^(v11)))))
                            } else {
                                v20 := GenS2{F0: "a", F1: -1}
                                fmt.Println(fi12)
                                _ = v20
                            }
                        default:
                            v9 = (v9 * 4294967295)
                        }
                        _ = fi12
                    }
                }
            }
            _ = len(v1)
            _ = fi10
            _ = v11
        }
        v21 := GenS2{F0: v8, F1: -32768}
        switch (("hello" + ((v8 + "foo bar") + ("foo bar" + ""))) + v8) {
        case "foo bar":
            fmt.Println(v4)
        case "\n\t":
            if v2 {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v4 -= 32767
                    if ((v8 > "\n\t") || true) {
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 uint8 = 42
                            _ = fi23
                            _ = v24
                        }
                        var v25 uint16 = 1
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS2{F0: "", F1: -32768}
                            _ = fi26
                            _ = v27
                        }
                    }
                    var v28 uint8 = 48
                    _ = fi22
                    _ = v28
                }
            }
        default:
            for fi29 := 0; fi29 < 4; fi29++ {
                v30 := GenS2{F0: v8, F1: 0}
                v31 := GenS3{F0: -1, F1: 992, F2: (v8 + "a")}
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        _ = v21
    }
    var v32 uint64 = 0
    v33 := v5["a"]
    switch (v8 + v8) {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 float32 = 0.0
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := GenS2{F0: (v8 + "foo bar"), F1: 42}
        v2 = (v2 && v2)
        _ = fi36
        _ = v37
    }
    switch -871 {
    case 0:
        var v38 int = 127
        _ = v38
    case -9223372036854775808:
        v39 := GenS4{F0: 2147483647, F1: 0}
        _ = v39
    default:
        for fi40 := 0; fi40 < 15; fi40++ {
            v9 = v9
            for fi41 := 0; fi41 < 4; fi41++ {
                v42 := GenS2{F0: v8, F1: -1}
                var v43 int64 = 506
                v8 = ("" + (v8 + (v8 + (v8 + ("a" + v8)))))
                _ = fi41
                _ = v42
                _ = v43
            }
            if false {
                v4 = (v4 & v4)
                var v44 int32 = 1
                var v45 uint32 = 0
                _ = v44
                _ = v45
            } else {
                switch v8 {
                case "hello":
                    fmt.Println(v9, v8, v33)
                default:
                    if false {
                        v46 := GenS5{F0: (18446744073709551615 == 990), F1: (fi40 & -2147483648), F2: -907}
                        var v47 uint32 = (v9 / ((4294967295) | 1))
                        v48 := GenS5{F0: (uint8(v4) != 255), F1: (^(fi40)), F2: 2147483647}
                        _ = v46
                        _ = v47
                        _ = v48
                    } else {
                        switch (("hello" + v8) + (((v8 + v8) + (v8 + "\n\t")) + ((v8 + v8) + ((v8 + "") + ("" + (v8 + v8)))))) {
                        case "":
                            if (65535 < 2) {
                                v1[fi40] = fi40
                            }
                        case "hello":
                            v49 := GenS6{F0: fi40, F1: 0, F2: 0}
                            _ = v49
                        default:
                            v8 = v8
                        }
                        switch v8 {
                        case "hello":
                            v5[v8] = v33
                        }
                        v50 := GenS5{F0: (17 != 1), F1: fi40, F2: 0}
                        _ = v50
                    }
                }
            }
            _ = fi40
        }
    }
    v2 = (v33 || v33)
    var v51 float64 = ((0.1 + -2.25) - 1e10)
    if true {
        for fi52 := 0; fi52 < 8; fi52++ {
            for fi53 := 0; fi53 < 10; fi53++ {
                var v54 float32 = 0.1
                switch 42 {
                case 9223372036854775807:
                    switch int(v4) {
                    case -182:
                        v54 -= v54
                        var v55 string = "\n\t"
                        var v56 uint8 = 127
                        switch v8 {
                        case "foo bar":
                            v57 := GenS7{F0: 1000, F1: 100}
                            _ = v57
                        }
                        var v58 int = fi53
                        v32 -= v32
                        v8 = "a"
                        _ = v55
                        _ = v56
                        _ = v58
                    case 9223372036854775807:
                        v59 := GenS8{F0: v9}
                        _ = v59
                    }
                case 354:
                    if false {
                        v60 := v5[(v8 + "hello")]
                        v33 = (v2 && v33)
                        _ = v60
                    }
                case 503:
                    v33 = (!(int32(fi52) <= -154))
                }
                var v61 int16 = 73
                _ = fi53
                _ = v54
                _ = v61
            }
            if (v9 >= 2147483647) {
                var v62 int32 = -32
                v33 = v2
                var v63 int8 = 124
                _ = v62
                _ = v63
            } else {
                switch ("\n\t" + v8) {
                case "\n\t":
                    switch v8 {
                    case "foo bar":
                        v51 -= (3.14159 * -2.25)
                    case "":
                        switch fi52 {
                        case 32768:
                            v64 := GenS6{F0: int(v32), F1: 1, F2: 1}
                            _ = v64
                        case 928:
                            v51 -= (-2.25 + (v51 - v51))
                        case 2147483647:
                            v4 = uint16(v32)
                        default:
                            var v65 int8 = 0
                            _ = v65
                        }
                    default:
                        v33 = (!v2)
                    }
                default:
                    if v2 {
                        delete(v1, (^(fi52)))
                        fmt.Println(v33, v4, fi52)
                    }
                }
                v32 += v32
                v33 = false
            }
            v32 = 0
            _ = fi52
        }
        v66 := GenS4{F0: 345, F1: v32}
        v67 := GenS8{F0: 65535}
        _ = v66
        _ = v67
    }
    fmt.Println(v32, v2, v4)
    v33 = true
    if v2 {
        v68 := GenS9{F0: (float32((v51 - float64(((0.0 - 1e10) * -2.25)))) >= (0.0 + -2.25)), F1: 1, F2: 289}
        _ = v68
    } else {
        var v69 int = 231
        v2 = v33
        _ = v69
    }
    v70 := GenS2{F0: v8, F1: 32767}
    var v71 int = 0
    for fi72 := 0; fi72 < 10; fi72++ {
        fmt.Println(v9)
        _ = fi72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v71:", v71)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v70
    _ = v71
    _ = v8
    _ = v9
}
