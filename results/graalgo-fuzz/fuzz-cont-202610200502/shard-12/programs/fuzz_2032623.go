package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 int16; F2 string }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int; F1 uint16; F2 uint64 }
func genFn0(p0 int, p1 uint, p2 uint32) bool {
    var v5 int64 = -423
    v5 -= v5
    var v6 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (int8(v6) < 127)
}

func main() {
    var v1 float32 = 0.0
    var v2 float64 = (0.1 - 0.0)
    v3 := []uint32{}
    var v4 int8 = 0
    var v7 float32 = v1
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v8 int32 = -2147483648
        _ = v8
    case "":
        v9 := GenS1{F0: 42, F1: "foo bar"}
        _ = v9
    case "a":
        if (!false) {
            fmt.Println(v1, v7, v2)
            var v10 int8 = v4
            var v11 uint32 = 2
            _ = v10
            _ = v11
        }
    }
    var v12 bool = true
    v13 := GenS2{F0: v12, F1: (-(v4)), F2: 85}
    switch 42 {
    case 127:
        v14 := GenS3{F0: 18446744073709551615}
        _ = v14
    case 0:
        for fi15 := 0; fi15 < 5; fi15++ {
            switch "\n\t" {
            case "\n\t":
                fmt.Println(fi15, v4)
            case "hello":
                v16 := GenS2{F0: (false || true), F1: (v4 * v4), F2: 42}
                _ = v16
            case "":
                var v17 uint64 = 1
                _ = v17
            default:
                v7 += v7
            }
            var v18 uint64 = 1
            var v19 float64 = float64(0.0)
            _ = fi15
            _ = v18
            _ = v19
        }
    case 139:
        v1 = float32(v2)
    default:
        switch -240 {
        case 680:
            if false {
                if (v1 == v7) {
                    fmt.Println(v2)
                } else {
                    v20 := GenS2{F0: v12, F1: (v4 + (v4 - (-(v4)))), F2: 1}
                    switch 1 {
                    case 42:
                        fmt.Println(v12, v12)
                    case 2147483647:
                        var v21 int = -911
                        _ = v21
                    case -1:
                        v7 += 1.5
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v20
                }
                if false {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        fmt.Println(v2)
                    case "a":
                        v4 += 124
                    case "hello":
                        fmt.Println(v4, v4)
                    }
                    if (!(-9223372036854775808 == -1)) {
                        v22 := GenS4{F0: 42, F1: 246, F2: ""}
                        for fi23 := 0; fi23 < 14; fi23++ {
                            if v12 {
                                var v24 string = ""
                                _ = v24
                            }
                            var v25 uint = 65535
                            _ = fi23
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 10; fi26++ {
                            switch (-(fi26)) {
                            case -1:
                                v3 = append(v3, 65535)
                            case 32767:
                                if (("hello" + "") != "\n\t") {
                                    v27 := GenS3{F0: 1}
                                    v2 = 1.5
                                    v2 = 0.0
                                    _ = v27
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if true {
                                        v28 := GenS1{F0: 933, F1: "\n\t"}
                                        v7 -= float32((float64(v1) / float64(v7)))
                                        v3 = append(v3, 2147483647)
                                        _ = v28
                                    } else {
                                        var v29 bool = false
                                        _ = v29
                                    }
                                    if v12 {
                                        _ = len(v3)
                                        var v30 uint32 = 2147483648
                                        v12 = (18446744073709551615 >= 18446744073709551615)
                                        _ = v30
                                    }
                                }
                            default:
                                var v31 uint16 = 65535
                                _ = v31
                            }
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            _ = fi26
                        }
                        _ = v22
                    } else {
                        v32 := GenS5{F0: 705, F1: 2, F2: 1}
                        _ = v32
                    }
                    if (2 <= 0) {
                        var v33 uint64 = 4294967295
                        _ = v33
                    }
                }
            }
        case -103:
            v1 -= 0.0
        case 42:
            var v34 string = ("foo bar" + "")
            _ = v34
        default:
            var v35 int32 = 2147483647
            _ = v35
        }
    }
    v36 := GenS1{F0: 65535, F1: "a"}
    v3 = append(v3, 2)
    v4 = (v4 >> ((699) & 63))
    for fi37 := 0; fi37 < 19; fi37++ {
        _ = len(v3)
        _ = fi37
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        v39 := GenS6{F0: 0, F1: 32767, F2: 4294967295}
        for fi40 := 0; fi40 < 12; fi40++ {
            v7 = (-2.25 * 1e10)
            var v41 int64 = -9223372036854775808
            if true {
                v1 += ((float32(0.1) - -2.25) * v7)
                v1 -= v1
            } else {
                switch fi38 {
                case 32767:
                    v4 = (v4 & int8(v41))
                case -9223372036854775808:
                    var v42 int16 = 1
                    _ = v42
                case -2147483648:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if ((0 >= 42) || (303 > 1)) {
                    var v43 int64 = (v41 << ((27) & 63))
                    var v44 string = "a"
                    _ = v43
                    _ = v44
                }
            }
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    v3 = append(v3, 65535)
    for fi45 := 0; fi45 < 10; fi45++ {
        var v46 int32 = 111
        var v47 int = fi45
        v47 += (^(fi45))
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v7
}
