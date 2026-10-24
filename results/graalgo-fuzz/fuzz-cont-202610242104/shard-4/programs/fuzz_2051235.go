package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 bool; F2 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    var v3 float32 = 0.0
    v4 := []int{2147483647}
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            switch "foo bar" {
            case "\n\t":
                if false {
                    switch fi6 {
                    case 42:
                        fmt.Println(v3)
                    default:
                        v7 := GenS1{F0: 9223372036854775807}
                        _ = v7
                    }
                    var v8 uint16 = 0
                    if false {
                        fmt.Println(v3, v3)
                        switch "\n\t" {
                        case "a":
                            for fi9 := 0; fi9 < 14; fi9++ {
                                var v10 float32 = 0.1
                                var v11 uint32 = 2147483647
                                var v12 float32 = v3
                                _ = fi9
                                _ = v10
                                _ = v11
                                _ = v12
                            }
                        case "\n\t":
                            v5[(fi6 * fi6)] = ""
                        case "hello":
                            if false {
                                var v13 bool = (((false || true) && false) && (!(42 <= 1)))
                                _ = v13
                            }
                        default:
                            v14 := GenS2{F0: -9223372036854775808}
                            _ = v14
                        }
                        if false {
                            fmt.Println(v8, v8)
                            for fi15 := 0; fi15 < 10; fi15++ {
                                v16 := GenS2{F0: (fi15 % ((int(v8)) | 1))}
                                if false {
                                    v17 := GenS3{F0: 42}
                                    switch (fi6 % ((int(v8)) | 1)) {
                                    case 42:
                                        var v18 int16 = -32768
                                        _ = v18
                                    case 2147483647:
                                        v19 := GenS3{F0: 660}
                                        _ = v19
                                    case 32767:
                                        v5[(fi15 << ((18446744073709551615) & 63))] = ("a" + "\n\t")
                                    default:
                                        delete(v5, fi6)
                                    }
                                    _ = v17
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = (("\n\t" + "a") + "hello")
                                    }
                                    if false {
                                        if (-39 >= -128) {
                                            switch "\n\t" {
                                            case "a":
                                                v20 := GenS1{F0: 0}
                                                _ = v20
                                            case "":
                                                fmt.Println(v8, v3)
                                                switch "foo bar" {
                                                case "hello":
                                                    _ = len(v1)
                                                case "":
                                                    _ = len(v4)
                                                }
                                                v8 = (v8 & (^(v8)))
                                            }
                                            v3 -= float32(0.0)
                                            if len(v4) > 0 {
                                                v4[0] = 42
                                            }
                                            _ = len(v4)
                                            var v21 int16 = 1000
                                            if (-807 != 806) {
                                                switch "hello" {
                                                case "foo bar":
                                                    var v22 int8 = 127
                                                    _ = v22
                                                case "a":
                                                    if (true && true) {
                                                        switch fi6 {
                                                        case 0:
                                                            var v23 uint8 = 14
                                                            _ = v23
                                                        }
                                                        v21 = v21
                                                        v21 += 42
                                                    } else {
                                                        v24 := GenS3{F0: 2}
                                                        if len(v2) > 0 {
                                                            v2[0] = ("a" + ("foo bar" + ("foo bar" + "hello")))
                                                        }
                                                        _ = v24
                                                    }
                                                }
                                                v25 := GenS2{F0: (-(fi6))}
                                                _ = v25
                                            }
                                            _ = v21
                                        } else {
                                            var v26 int = fi6
                                            _ = v26
                                        }
                                        v1[("" + "\n\t")] = 94
                                        v27 := GenS3{F0: 2}
                                        _ = v27
                                    }
                                }
                                if true {
                                    v28 := GenS2{F0: fi15}
                                    var v29 float64 = 3.14159
                                    _ = len(v2)
                                    _ = v28
                                    _ = v29
                                } else {
                                    v3 = float32(float64(v3))
                                    if true {
                                        if (0.1 == 3.14159) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            if (0 <= -650) {
                                                v30 := GenS2{F0: fi6}
                                                if (!((100 < 1) && (!(42 > 396)))) {
                                                    v3 = v3
                                                    var v31 uint8 = 83
                                                    fmt.Println(v3, fi6, fi15)
                                                    _ = v31
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                fmt.Println(fi6)
                                                _ = v30
                                            }
                                        }
                                    }
                                }
                                _ = fi15
                                _ = v16
                            }
                            fmt.Println(v8)
                        }
                    } else {
                        switch "foo bar" {
                        case "":
                            switch "hello" {
                            case "a":
                                v3 += float32((0.0 * (float64(float32(((3.14159 / 3.14159) + 1.5))) * 1.5)))
                            case "\n\t":
                                if (true || true) {
                                    v32 := GenS4{F0: 1, F1: 1}
                                    for fi33 := 0; fi33 < 19; fi33++ {
                                        v34 := GenS2{F0: -133}
                                        v8 -= v8
                                        v35 := GenS5{F0: (!false), F1: 127}
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v32
                                }
                            }
                        case "\n\t":
                            v8 = v8
                        }
                        v36 := GenS6{F0: 0}
                        for fi37 := 0; fi37 < 1; fi37++ {
                            v3 = 0.0
                            if (65535 != 65535) {
                                fmt.Println(fi37, v8)
                                var v38 uint = 255
                                var v39 uint64 = 817
                                var v40 uint64 = (v39 % (((^(v39))) | 1))
                                var v41 int = fi37
                                v42 := GenS7{F0: -2147483648, F1: (true && (("a" + "") != "foo bar")), F2: fi37}
                                v43 := GenS4{F0: 0, F1: v8}
                                v38 = (v38 >> ((uint(v40)) & 63))
                                var v44 float32 = v3
                                _ = len(v4)
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            fmt.Println(v8)
                            _ = fi37
                        }
                        _ = v36
                    }
                    _ = v8
                } else {
                    for fi45 := 0; fi45 < 12; fi45++ {
                        delete(v1, "")
                        _ = fi45
                    }
                    if (!(125 <= 255)) {
                        var v46 float32 = v3
                        _ = v46
                    }
                }
            case "foo bar":
                if false {
                    if false {
                        v3 = float32(0.0)
                    }
                    var v47 int64 = -282
                    v5[fi6] = ""
                    _ = v47
                }
            }
            v48 := GenS8{F0: true}
            for fi49 := 0; fi49 < 17; fi49++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v50 := GenS7{F0: 345, F1: false, F2: fi6}
                _ = len(v4)
                _ = fi49
                _ = v50
            }
            _ = v48
        }
        var v51 int32 = 42
        _ = fi6
        _ = v51
    }
    delete(v5, 810)
    fmt.Println(v3)
    v3 += (v3 - ((-2.25 + v3) + 1.5))
    var v52 int32 = -2147483648
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v3)
        var v53 uint64 = 1
        _ = v53
    } else {
        v54 := GenS8{F0: true}
        v55 := GenS9{F0: "hello"}
        _ = v54
        _ = v55
    }
    var v56 uint = 42
    v52 -= (v52 % ((v52) | 1))
    for fi57 := 0; fi57 < 2; fi57++ {
        var v58 int = (-(fi57))
        v58 = v58
        v56 = v56
        _ = fi57
        _ = v58
    }
    _ = len(v5)
    switch "" {
    case "a":
        v56 += v56
    default:
        if (-9223372036854775808 >= 42) {
            var v59 bool = false
            v60 := GenS4{F0: (v52 ^ v52), F1: 1}
            var v61 int32 = v52
            _ = v59
            _ = v60
            _ = v61
        }
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v52 += (-(v52))
    var v62 int64 = -9223372036854775808
    for fi63 := 0; fi63 < 9; fi63++ {
        v64 := GenS10{F0: -128, F1: (v52 % (((^(v52))) | 1))}
        fmt.Println(v3, v52)
        _ = fi63
        _ = v64
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v56
    _ = v62
}
