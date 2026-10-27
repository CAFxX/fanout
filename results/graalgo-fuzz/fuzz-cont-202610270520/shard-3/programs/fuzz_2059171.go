package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int64; F1 string; F2 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int64; F1 int16; F2 int32 }
type GenS6 struct { F0 int; F1 string }

func main() {
    var v1 uint = 0
    var v2 uint16 = uint16(v1)
    v3 := []string{}
    var v4 uint8 = 0
    v5 := make(map[string]int)
    v2 = v2
    if true {
        v4 = 0
    } else {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int64 = 42
            switch "\n\t" {
            case "a":
                switch ("a" + ("a" + "\n\t")) {
                case "":
                    v7 = (v7 / (((v7 + v7)) | 1))
                case "foo bar":
                    v3 = append(v3, "hello")
                default:
                    if false {
                        var v8 string = "\n\t"
                        _ = v8
                    } else {
                        var v9 float64 = -2.25
                        v10 := GenS1{F0: -1, F1: -946}
                        _ = v10
                        _ = v9
                    }
                }
            }
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int32 = 0
                v4 = (^(v4))
                _ = fi11
                _ = v12
            }
            _ = fi6
            _ = v7
        }
    }
    var v13 int64 = 1
    var v14 uint64 = 616
    for fi15 := 0; fi15 < 20; fi15++ {
        if (v2 > (v2 << ((v1) & 63))) {
            if true {
                v2 += (^(v2))
            } else {
                switch "foo bar" {
                case "":
                    v5["hello"] = fi15
                case "hello":
                    v2 = uint16(v4)
                case "\n\t":
                    v16 := GenS1{F0: 1, F1: 1}
                    _ = v16
                default:
                    v2 -= 332
                }
                for fi17 := 0; fi17 < 11; fi17++ {
                    v4 += v4
                    var v18 float64 = 3.14159
                    if (false && ("foo bar" <= ("" + ""))) {
                        var v19 float32 = 0.0
                        v13 += (v13 & v13)
                        var v20 float32 = (v19 / (v19 * 3.14159))
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int16 = 1000
                        var v22 int8 = -52
                        var v23 int64 = (v13 * v13)
                        if false {
                            var v24 float64 = (((v18 - (v18 - (float64(0.0) - (float64(0.1) + -2.25)))) + 0.0) * 1.5)
                            v25 := GenS2{F0: v23, F1: "", F2: ((("hello" + ("hello" + "hello")) + ("a" + "")) + "foo bar")}
                            v26 := GenS3{F0: uint64(v4), F1: -817}
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            var v27 uint32 = 1
                            _ = v27
                        }
                        fmt.Println(v4)
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi17
                    _ = v18
                }
            }
            var v28 int64 = v13
            for fi29 := 0; fi29 < 4; fi29++ {
                v5[("\n\t" + ("" + ""))] = (-(fi15))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v30 := GenS3{F0: v14, F1: -635}
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            fmt.Println(v13, v1, v4)
        }
        if false {
            v2 = 32767
            v31 := GenS1{F0: -1, F1: 42}
            var v32 int32 = -1
            _ = v31
            _ = v32
        } else {
            if (-1 < -2147483648) {
                v14 = (v14 | (v14 + (v14 / ((v14) | 1))))
            }
            v13 -= v13
            switch (fi15 & (fi15 / ((-9223372036854775808) | 1))) {
            case -296:
                for fi33 := 0; fi33 < 2; fi33++ {
                    switch "a" {
                    case "a":
                        var v34 uint16 = (v2 ^ v2)
                        _ = v34
                    default:
                        v35 := GenS1{F0: 22, F1: -183}
                        var v36 bool = ((-(fi15)) == 722)
                        if (("hello" + ("foo bar" + ("foo bar" + "a"))) > (("a" + "hello") + ("hello" + ("a" + "")))) {
                            v37 := GenS1{F0: 100, F1: 501}
                            v13 = (-(v13))
                            if v36 {
                                v38 := GenS4{F0: (v2 + v2), F1: fi33}
                                _ = v38
                            }
                            if true {
                                v14 = 1
                                v39 := GenS1{F0: 9, F1: 1}
                                _ = v39
                            }
                            _ = v37
                        }
                        _ = v35
                        _ = v36
                    }
                    _ = fi33
                }
            case 9223372036854775807:
                fmt.Println(fi15, v1)
            default:
                var v40 bool = (42 == 65535)
                _ = v40
            }
        }
        _ = fi15
    }
    var v41 int32 = 1
    switch "" {
    case "foo bar":
        v1 += (v1 & v1)
    }
    switch "a" {
    case "":
        var v42 uint32 = 65535
        _ = v42
    case "a":
        if (!true) {
            if false {
                var v43 float32 = float32(1.5)
                v4 += uint8(v1)
                v44 := GenS5{F0: (v13 - (v13 << (((v1 / (((v1 & 205)) | 1))) & 63))), F1: 902, F2: 0}
                _ = v43
                _ = v44
            }
        } else {
            fmt.Println(v13)
            v1 = v1
            var v45 uint = (v1 >> ((0) & 63))
            _ = v45
        }
    }
    v46 := GenS1{F0: -107, F1: v41}
    switch "\n\t" {
    case "a":
        var v47 bool = false
        _ = v47
    default:
        v1 = uint(v2)
    }
    if (v2 == 2) {
        v2 = (v2 - (v2 << (((v1 / ((2) | 1))) & 63)))
        for fi48 := 0; fi48 < 18; fi48++ {
            if (100 == 0) {
                var v49 uint = (v1 + (v1 / ((uint(v4)) | 1)))
                var v50 int = (fi48 | (fi48 + fi48))
                v3 = append(v3, ("a" + "hello"))
                _ = v49
                _ = v50
            } else {
                v14 = 2
                for fi51 := 0; fi51 < 2; fi51++ {
                    switch "a" {
                    case "foo bar":
                        var v52 int = fi51
                        _ = v52
                    case "\n\t":
                        var v53 int8 = -66
                        _ = v53
                    case "":
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    }
                    v54 := GenS6{F0: int(v4), F1: ("hello" + "")}
                    fmt.Println(v14, v1)
                    _ = fi51
                    _ = v54
                }
            }
            var v55 float64 = 0.0
            _ = fi48
            _ = v55
        }
        var v56 float32 = (3.14159 - 3.14159)
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v5
}
