package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 uint8, p1 int, p2 uint) int32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            p2 -= 255
            _ = fi8
        }
        _ = fi7
    }
    var v9 float64 = 3.14159
    v10 := GenS1{F0: -413, F1: ("" + "hello"), F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -316
}
func genFn1(p0 uint) int64 {
    if true {
        if (-1 <= -31) {
            var v11 uint16 = 42
            _ = v11
        }
        v12 := GenS1{F0: 42, F1: "", F2: uint64(p0)}
        _ = v12
    } else {
        var v13 uint64 = 42
        _ = v13
    }
    if (!(false || (!true))) {
        var v14 float32 = (float32(1e10) + 1e10)
        for fi15 := 0; fi15 < 2; fi15++ {
            fmt.Println(p0, fi15)
            var v16 float64 = 3.14159
            for fi17 := 0; fi17 < 15; fi17++ {
                switch (fi15 % ((fi17) | 1)) {
                case 4294967295:
                    v18 := GenS1{F0: -1, F1: "", F2: 355}
                    _ = v18
                case 0:
                    var v19 int8 = -70
                    _ = v19
                default:
                    var v20 uint16 = 297
                    _ = v20
                }
                var v21 uint64 = 124
                _ = fi17
                _ = v21
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    var v22 float32 = 3.14159
    fmt.Println(v22, p0)
    _ = p0
    _ = v22
    return 337
}

func main() {
    v1 := []string{""}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v4 int16 = 42
    var v5 bool = true
    v6 := make(map[string]bool)
    if ((false || v5) || (!v5)) {
        v3[1] = ""
        v23 := GenS2{F0: 42, F1: 65535}
        var v24 uint8 = 100
        _ = v23
        _ = v24
    }
    var v25 int = -2147483648
    v25 = 2147483647
    if v5 {
        _ = len(v1)
        switch (v25 % ((0) | 1)) {
        case -918:
            v25 = 0
        case 127:
            _ = len(v1)
        default:
            _ = len(v1)
        }
    } else {
        v5 = true
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := GenS2{F0: -141, F1: 1}
            var v28 float32 = 3.14159
            v4 = v4
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    var v29 uint32 = 2147483648
    fmt.Println(v4, v5)
    fmt.Println(v29, v29, v5)
    var v30 int = -9223372036854775808
    _ = len(v1)
    v4 = (v4 + v4)
    v31 := v6["foo bar"]
    v31 = v31
    v4 -= v4
    if v5 {
        v29 = 2147483648
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v29)
        }
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 float32 = ((3.14159 / 0.1) - 0.1)
            fmt.Println(v30, v31, v25)
            delete(v6, "foo bar")
            _ = fi33
            _ = v34
        }
        v35 := GenS2{F0: 9223372036854775807, F1: v29}
        _ = fi32
        _ = v35
    }
    if v5 {
        v25 -= v25
    } else {
        fmt.Println(v31)
        var v36 int64 = -1
        if (1 >= v29) {
            v5 = (true || v31)
            for fi37 := 0; fi37 < 7; fi37++ {
                v36 += 9223372036854775807
                for fi38 := 0; fi38 < 16; fi38++ {
                    if v31 {
                        v25 += v25
                    } else {
                        var v39 float64 = 0.0
                        if ("" > ("" + ("\n\t" + "\n\t"))) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v40 := GenS3{F0: (v36 ^ -1), F1: 42}
                            if true {
                                switch "a" {
                                case "\n\t":
                                    v41 := GenS4{F0: 255}
                                    _ = v41
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                case "hello":
                                    var v42 int = -2147483648
                                    _ = v42
                                default:
                                    v43 := GenS2{F0: 42, F1: (v29 | v29)}
                                    _ = v43
                                }
                                if true {
                                    if false {
                                        var v44 int16 = (v4 % ((1) | 1))
                                        var v45 string = "foo bar"
                                        _ = v44
                                        _ = v45
                                    } else {
                                        switch ("foo bar" + "hello") {
                                        case "foo bar":
                                            var v46 float32 = 0.0
                                            _ = v46
                                        case "hello":
                                            var v47 int32 = 498
                                            var v48 uint16 = 31
                                            fmt.Println(v47, v48, v31)
                                            _ = v47
                                            _ = v48
                                        default:
                                            v49 := GenS5{F0: (v29 % ((v29) | 1)), F1: (v36 - (v36 * (^(v36)))), F2: 946}
                                            _ = v49
                                        }
                                        if true {
                                            var v50 float32 = 1e10
                                            _ = v50
                                        }
                                    }
                                    switch (v30 * fi38) {
                                    case 785:
                                        if false {
                                            var v51 int64 = v36
                                            if v31 {
                                                var v52 float64 = float64(1e10)
                                                v30 += (v25 >> ((42) & 63))
                                                var v53 uint64 = 249
                                                _ = v52
                                                _ = v53
                                            } else {
                                                v54 := GenS3{F0: (v36 & (-(v51))), F1: v29}
                                                _ = v54
                                            }
                                            _ = v51
                                        } else {
                                            fmt.Println(v39)
                                            var v55 int8 = 30
                                            _ = v55
                                        }
                                    case 128:
                                        v29 = v29
                                    case 200:
                                        v56 := v3[v29]
                                        _ = v56
                                    }
                                } else {
                                    if (!false) {
                                        var v57 uint16 = 255
                                        v58 := GenS6{F0: 67}
                                        var v59 int64 = (^(v36))
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v60 int = (v25 ^ int(v36))
                                        if (636 <= 2) {
                                            v4 -= v4
                                            var v61 int32 = -1
                                            _ = v61
                                        } else {
                                            var v62 uint64 = 281
                                            v36 = 85
                                            _ = v62
                                        }
                                        _ = v60
                                    }
                                    if true {
                                        switch fi38 {
                                        case 638:
                                            v39 += v39
                                        case 235:
                                            fmt.Println(v29, v39)
                                        case 127:
                                            v36 -= v36
                                        }
                                        var v63 string = "\n\t"
                                        _ = v63
                                    } else {
                                        if v5 {
                                            v1 = append(v1, ("a" + ""))
                                        } else {
                                            v64 := GenS7{F0: 0}
                                            var v65 string = "foo bar"
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                    v30 -= v25
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v39, v5, v30)
                                    }
                                }
                                v4 = (v4 % ((42) | 1))
                            } else {
                                v66 := GenS8{F0: 128, F1: 64, F2: 2}
                                _ = v66
                            }
                            _ = v40
                        }
                        fmt.Println(v30, v39)
                        _ = v39
                    }
                    _ = fi38
                }
                v29 += (^(v29))
                _ = fi37
            }
        }
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
