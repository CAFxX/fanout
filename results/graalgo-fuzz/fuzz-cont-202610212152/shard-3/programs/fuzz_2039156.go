package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int16; F1 uint64; F2 uint64 }
func genFn0() int32 {
    if true {
        v5 := GenS1{F0: 0, F1: 18446744073709551615, F2: 1}
        _ = v5
    }
    var v6 float32 = float32(float64((1e10 * (0.0 - float32(-2.25)))))
    v6 = (0.0 * v6)
    _ = v6
    return 2147483647
}
func genFn1(p0 int64) int64 {
    v7 := []int{540}
    if (915 < 0) {
        p0 += (p0 * (p0 << ((2) & 63)))
        p0 -= (p0 << ((18446744073709551615) & 63))
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = p0
    _ = v7
    return p0
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    if false {
        _ = len(v2)
        var v8 uint64 = 0
        for fi9 := 0; fi9 < 18; fi9++ {
            if false {
                switch fi9 {
                case 463:
                    switch "" {
                    case "foo bar":
                        v8 = (^(v8))
                    }
                default:
                    v8 = v8
                }
                if true {
                    v8 = v8
                    delete(v4, 2147483648)
                }
            } else {
                _ = len(v3)
                var v10 uint8 = 42
                _ = v10
            }
            v3[65535] = 63
            _ = fi9
        }
        _ = v8
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    var v11 uint32 = 0
    var v12 int = 1
    v12 += (v12 << ((255) & 63))
    v12 += v12
    var v13 bool = true
    for fi14 := 0; fi14 < 19; fi14++ {
        switch ("foo bar" + "\n\t") {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "hello":
            var v15 uint32 = 303
            _ = v15
        case "\n\t":
            var v16 int = (^(fi14))
            _ = v16
        default:
            for fi17 := 0; fi17 < 6; fi17++ {
                switch (v12 >> ((2) & 63)) {
                case -229:
                    v13 = false
                case 42:
                    if true {
                        v18 := v3[2]
                        switch "foo bar" {
                        case "hello":
                            if v13 {
                                var v19 uint64 = 42
                                _ = v19
                            } else {
                                if v13 {
                                    v11 = 431
                                    var v20 uint32 = (v11 & 2147483647)
                                    var v21 float64 = 1.5
                                    v22 := GenS2{F0: -32768, F1: 228}
                                    if v13 {
                                        v23 := GenS3{F0: ("foo bar" + ""), F1: 871, F2: 1}
                                        if (1e10 <= 1e10) {
                                            fmt.Println(v20, v21)
                                        }
                                        if ((v20 << ((uint(v18)) & 63)) != (^(v20))) {
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = fi17
                                                }
                                                fmt.Println(v21)
                                            } else {
                                                switch (-(fi14)) {
                                                case -1:
                                                    if (!(65535 >= 32767)) {
                                                        var v24 string = "a"
                                                        switch ("a" + v24) {
                                                        case "hello":
                                                            var v25 int32 = -2147483648
                                                            _ = v25
                                                        case "a":
                                                            fmt.Println(v11, fi14)
                                                        case "\n\t":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v26 int = fi17
                                                            if false {
                                                                v27 := GenS4{F0: 0, F1: 0}
                                                                var v28 uint8 = 128
                                                                switch ("foo bar" + "a") {
                                                                case "a":
                                                                    v11 += (v11 * (^(v18)))
                                                                case "\n\t":
                                                                    if (((^(fi17)) >= fi14) && v13) {
                                                                        v29 := GenS1{F0: -1, F1: 4294967295, F2: (v28 & 127)}
                                                                        _ = v29
                                                                    } else {
                                                                        _ = len(v4)
                                                                        var v30 uint = 4294967295
                                                                        v28 = 42
                                                                        _ = v30
                                                                    }
                                                                }
                                                                _ = v27
                                                                _ = v28
                                                            }
                                                            _ = v26
                                                        }
                                                        _ = v24
                                                    }
                                                case 9223372036854775807:
                                                    delete(v4, (^(v18)))
                                                default:
                                                    fmt.Println(v12, v12)
                                                }
                                                if true {
                                                    delete(v2, v12)
                                                    v18 -= 65535
                                                }
                                            }
                                            v31 := GenS3{F0: "a", F1: 0, F2: 2}
                                            _ = v31
                                        } else {
                                            fmt.Println(v11, v21)
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v13 = (2 > 42)
                                            _ = len(v3)
                                        }
                                        _ = v23
                                    } else {
                                        v32 := GenS1{F0: genFn1(-1), F1: uint64(v18), F2: 2}
                                        if v13 {
                                            if len(v1) > 0 {
                                                v1[0] = 127
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                if v13 {
                                                    var v33 float64 = (float64(1e10) + v21)
                                                    var v34 string = "foo bar"
                                                    fmt.Println(v18)
                                                    var v35 uint8 = 42
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                }
                                            case "":
                                                if true {
                                                    switch "a" {
                                                    case "":
                                                        var v36 int32 = genFn0()
                                                        _ = v36
                                                    }
                                                    if (!(483 == 0)) {
                                                        var v37 uint = 255
                                                        _ = v37
                                                    } else {
                                                        var v38 int = fi14
                                                        delete(v2, -9223372036854775808)
                                                        _ = len(v3)
                                                        v2[v12] = "a"
                                                        _ = v38
                                                    }
                                                    v39 := GenS2{F0: 1000, F1: 128}
                                                    _ = v39
                                                }
                                            case "foo bar":
                                                v13 = v13
                                            default:
                                                var v40 uint8 = 255
                                                _ = v40
                                            }
                                            if v13 {
                                                switch -655 {
                                                case 42:
                                                    fmt.Println(v12, v20)
                                                default:
                                                    v20 = 65535
                                                }
                                            } else {
                                                delete(v4, v20)
                                            }
                                        }
                                        _ = v32
                                    }
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v41 uint16 = 1
                                    v3[(^(v11))] = 664
                                    var v42 uint64 = 943
                                    _ = v41
                                    _ = v42
                                }
                                v11 = v18
                                if v13 {
                                    var v43 uint = 938
                                    v43 += v43
                                    _ = len(v4)
                                    _ = v43
                                }
                            }
                        case "a":
                            if v13 {
                                var v44 uint64 = 4294967295
                                _ = v44
                            } else {
                                switch "" {
                                case "hello":
                                    switch "a" {
                                    case "":
                                        v11 -= 488
                                    default:
                                        var v45 uint8 = 42
                                        _ = v45
                                    }
                                }
                            }
                        }
                        v46 := GenS3{F0: "hello", F1: 1, F2: v18}
                        _ = v18
                        _ = v46
                    } else {
                        fmt.Println(v12, v11)
                        var v47 int16 = -1
                        if false {
                            v48 := GenS5{F0: -286, F1: 88}
                            var v49 int8 = 0
                            v50 := GenS6{F0: (v47 ^ v47), F1: 2, F2: 42}
                            var v51 uint8 = 128
                            var v52 int16 = -32768
                            var v53 uint = 2
                            _ = v48
                            _ = v49
                            _ = v50
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        _ = v47
                    }
                    var v54 int64 = -1
                    v54 = (v54 % ((v54) | 1))
                    _ = v54
                default:
                    v13 = v13
                }
                v55 := GenS5{F0: -1, F1: 0}
                _ = fi17
                _ = v55
            }
        }
        var v56 int64 = genFn1(1)
        var v57 int8 = 100
        _ = fi14
        _ = v56
        _ = v57
    }
    v13 = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
