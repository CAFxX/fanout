package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int64; F1 uint8; F2 int32 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 string = "a"
        var v9 uint8 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = 95
    fmt.Println(v10)
    v11 := []int{251, 652}
    _ = v10
    _ = v11
    return v10
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 9223372036854775807
    v3 := []uint32{2, 198, 2147483647}
    var v4 int16 = 1000
    v5 := make(map[uint32]string)
    var v6 uint16 = 65535
    var v12 int32 = 2147483647
    var v13 bool = true
    v12 = (v12 * v12)
    v4 += (v4 ^ 0)
    v14 := GenS1{F0: 797, F1: v12}
    v4 -= -140
    v15 := GenS2{F0: (v4 * v4)}
    v16 := GenS2{F0: 348}
    for fi17 := 0; fi17 < 14; fi17++ {
        switch "" {
        case "a":
            for fi18 := 0; fi18 < 7; fi18++ {
                switch "hello" {
                case "hello":
                    v4 -= -32768
                }
                _ = fi18
            }
        case "foo bar":
            fmt.Println(v13, v4, v2)
        case "hello":
            if v13 {
                if false {
                    fmt.Println(v2, fi17)
                }
                var v19 string = "\n\t"
                _ = v19
            } else {
                var v20 uint64 = 4294967295
                _ = v20
            }
        default:
            if v13 {
                var v21 uint64 = 42
                var v22 uint = 658
                switch ("hello" + "") {
                case "foo bar":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        if (1e10 >= 0.1) {
                            v24 := GenS3{F0: 65535, F1: 100, F2: (^(v6))}
                            v12 -= (v12 * (v12 - -2147483648))
                            v2 = 344
                            _ = v24
                        } else {
                            v6 = v6
                        }
                        var v25 string = ""
                        _ = fi23
                        _ = v25
                    }
                case "a":
                    if false {
                        switch fi17 {
                        case 4294967295:
                            switch fi17 {
                            case 4294967295:
                                v26 := GenS4{F0: 218}
                                _ = v26
                            case -904:
                                v27 := GenS5{F0: (v21 * v21), F1: (v21 + (v21 / (((v21 & (v21 ^ 940))) | 1)))}
                                _ = v27
                            }
                        case -997:
                            for fi28 := 0; fi28 < 2; fi28++ {
                                if (2147483647 >= v12) {
                                    var v29 uint16 = (v6 * v6)
                                    _ = v29
                                } else {
                                    var v30 int32 = 820
                                    v5[1] = "\n\t"
                                    _ = v30
                                }
                                v6 = uint16(v22)
                                _ = fi28
                            }
                        default:
                            var v31 uint16 = v6
                            _ = v31
                        }
                    } else {
                        var v32 uint8 = 42
                        var v33 string = ("" + "")
                        for fi34 := 0; fi34 < 9; fi34++ {
                            switch int(v4) {
                            case 42:
                                if (!v13) {
                                    switch ("" + (("foo bar" + "hello") + (((v33 + "\n\t") + "foo bar") + v33))) {
                                    case "hello":
                                        v21 += v21
                                    default:
                                        v12 = (v12 >> (((v22 >> ((v22) & 63))) & 63))
                                    }
                                }
                            default:
                                delete(v1, "\n\t")
                            }
                            _ = fi34
                        }
                        _ = v32
                        _ = v33
                    }
                default:
                    for fi35 := 0; fi35 < 14; fi35++ {
                        var v36 uint16 = 388
                        if (v13 || (0 != 504)) {
                            fmt.Println(v12, v21)
                            var v37 int8 = 0
                            v2 -= (v2 ^ (v2 * v2))
                            _ = v37
                        }
                        v38 := v1["\n\t"]
                        _ = fi35
                        _ = v36
                        _ = v38
                    }
                }
                _ = v21
                _ = v22
            }
        }
        var v39 uint16 = v6
        v6 = 266
        _ = fi17
        _ = v39
    }
    var v40 int16 = v4
    switch ("" + "") {
    case "hello":
        for fi41 := 0; fi41 < 14; fi41++ {
            if (!(!(((3.14159 + 1e10) * (3.14159 * 0.0)) <= -2.25))) {
                v12 -= 0
            } else {
                var v42 int = 32767
                _ = v42
            }
            _ = fi41
        }
    case "\n\t":
        switch "foo bar" {
        case "foo bar":
            v43 := GenS6{F0: v2, F1: 1, F2: (v12 | v12)}
            _ = v43
        case "":
            fmt.Println(v13)
        default:
            v4 = -362
        }
    case "foo bar":
        var v44 float64 = 1e10
        _ = v44
    default:
        var v45 int16 = v40
        _ = v45
    }
    var v46 float64 = 1.5
    v1[(("a" + "hello") + ("foo bar" + ""))] = 2147483647
    v47 := GenS1{F0: 65535, F1: v12}
    v12 += v12
    for fi48 := 0; fi48 < 9; fi48++ {
        v40 -= v4
        if v13 {
            for fi49 := 0; fi49 < 20; fi49++ {
                v46 = 0.1
                switch "a" {
                case "\n\t":
                    v50 := GenS6{F0: 42, F1: uint8(v6), F2: (v12 % (((v12 - (v12 - (v12 - -1)))) | 1))}
                    v2 = (v2 * (v2 << ((42) & 63)))
                    _ = v50
                case "a":
                    fmt.Println(v46, v4, v13)
                case "":
                    v13 = false
                }
                v51 := GenS7{F0: (v2 | v2), F1: v2}
                _ = fi49
                _ = v51
            }
            v52 := GenS3{F0: 1, F1: -61, F2: 42}
            _ = v52
        }
        if false {
            var v53 uint64 = 18446744073709551615
            var v54 bool = (-128 == int8(fi48))
            var v55 int16 = v40
            _ = v53
            _ = v54
            _ = v55
        } else {
            if v13 {
                v56 := GenS8{F0: 31}
                _ = v56
            } else {
                for fi57 := 0; fi57 < 3; fi57++ {
                    var v58 float32 = 1.5
                    _ = fi57
                    _ = v58
                }
                delete(v1, "hello")
                var v59 int = 632
                _ = v59
            }
            for fi60 := 0; fi60 < 9; fi60++ {
                var v61 float32 = 1.5
                var v62 float32 = (0.0 + v61)
                _ = fi60
                _ = v61
                _ = v62
            }
            var v63 int = fi48
            _ = v63
        }
        _ = fi48
    }
    v64 := GenS3{F0: 440, F1: -59, F2: v6}
    v65 := GenS9{F0: 4294967295, F1: 42, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
