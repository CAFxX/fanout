package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int16; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint32; F1 int64; F2 int64 }
type GenS8 struct { F0 int8; F1 uint32; F2 int }

func main() {
    v1 := []uint32{2}
    var v2 uint = 429
    var v3 uint = v2
    var v4 uint = v2
    var v5 uint = v3
    var v6 int16 = 779
    fmt.Println(v3, v3)
    fmt.Println(v4, v5, v3)
    var v7 uint64 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "foo bar" {
    case "\n\t":
        switch ("foo bar" + ("foo bar" + "foo bar")) {
        case "\n\t":
            v1 = append(v1, 135)
        case "hello":
            var v8 uint8 = 42
            _ = v8
        case "a":
            var v9 uint16 = 1
            _ = v9
        default:
            if (!(1 == 1)) {
                fmt.Println(v3, v5)
                fmt.Println(v7)
                if (!(!true)) {
                    var v10 int = -2147483648
                    fmt.Println(v5, v3, v10)
                    var v11 uint32 = 1
                    _ = v10
                    _ = v11
                } else {
                    if ((false || (65535 != 1)) || true) {
                        if true {
                            fmt.Println(v3, v3)
                            fmt.Println(v5, v2)
                            var v12 bool = (-100 >= 1)
                            _ = v12
                        }
                    } else {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            fmt.Println(v5, v3)
                            switch (fi13 / (((fi13 % (((fi13 & (fi13 % (((fi13 * fi13)) | 1)))) | 1))) | 1)) {
                            case 1:
                                v14 := GenS1{F0: 1, F1: ("a" + "a")}
                                _ = v14
                            case -706:
                                fmt.Println(v3, fi13)
                            case -303:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v15 := GenS2{F0: 4294967295}
                                _ = v15
                            }
                            _ = fi13
                        }
                    }
                    switch 32768 {
                    case -473:
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v2)
                        default:
                            fmt.Println(v6, v5, v7)
                        }
                    case -9223372036854775808:
                        v16 := GenS1{F0: 18446744073709551615, F1: ""}
                        _ = v16
                    default:
                        var v17 uint64 = (^(v7))
                        _ = v17
                    }
                    if (117 > 1) {
                        if (1 == 267) {
                            if (true && true) {
                                v1 = append(v1, 2147483648)
                            }
                            var v18 int16 = v6
                            _ = v18
                        } else {
                            fmt.Println(v4, v6, v2)
                            if ((218 <= 2) && ((v4 & v4) < (v4 % (((v5 >> (((v4 * v2)) & 63))) | 1)))) {
                                v5 += (v3 ^ 227)
                            } else {
                                if (0.1 > (float64((3.14159 * (0.1 - 1.5))) - (1.5 + (0.1 + 1.5)))) {
                                    fmt.Println(v6)
                                }
                                for fi19 := 0; fi19 < 17; fi19++ {
                                    if ((v7 == v7) && false) {
                                        var v20 uint16 = 255
                                        var v21 uint = (v5 ^ v5)
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 bool = false
                                    _ = fi19
                                    _ = v22
                                }
                                for fi23 := 0; fi23 < 6; fi23++ {
                                    v24 := GenS2{F0: 756}
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        fmt.Println(v7, v6)
                        fmt.Println(v3, v5)
                        v25 := GenS3{F0: 344}
                        _ = v25
                    }
                }
            } else {
                v26 := GenS4{F0: 128, F1: v6, F2: 127}
                v7 = (v7 | 816)
                var v27 int16 = -32768
                _ = v26
                _ = v27
            }
        }
    case "a":
        v5 -= 973
    case "foo bar":
        v28 := GenS2{F0: 2147483648}
        _ = v28
    default:
        for fi29 := 0; fi29 < 4; fi29++ {
            v3 += 42
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi30 := 0; fi30 < 7; fi30++ {
                if true {
                    if true {
                        if (!true) {
                            var v31 int8 = int8(fi29)
                            var v32 int64 = int64(v3)
                            _ = v31
                            _ = v32
                        } else {
                            v2 += (v2 << ((v5) & 63))
                            var v33 uint32 = 65535
                            _ = v33
                        }
                        v34 := GenS5{F0: 0}
                        _ = v34
                    } else {
                        v3 = 947
                    }
                    v6 = (v6 | v6)
                    var v35 uint8 = 1
                    _ = v35
                }
                _ = fi30
            }
            _ = fi29
        }
    }
    v36 := GenS4{F0: 288, F1: v6, F2: 42}
    if true {
        var v37 int32 = 2147483647
        if (v7 != v7) {
            var v38 float64 = 1.5
            _ = v38
        }
        _ = v37
    } else {
        if false {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            v7 += (v7 ^ (^(v7)))
            for fi39 := 0; fi39 < 18; fi39++ {
                v5 += uint(v7)
                _ = fi39
            }
        }
        var v40 int64 = 42
        v3 = (^(v3))
        _ = v40
    }
    var v41 string = ("" + "\n\t")
    v42 := GenS4{F0: 32768, F1: (v6 % (((v6 * (v6 + (v6 | -427)))) | 1)), F2: -111}
    if len(v1) > 0 {
        v1[0] = 1
    }
    v41 = (v41 + "")
    switch 0 {
    case -9223372036854775808:
        v43 := GenS6{F0: 68, F1: 1}
        _ = v43
    case 951:
        fmt.Println(v4)
    default:
        fmt.Println(v4)
    }
    if (((!(!true)) && (float32(((0.1 / 1.0) - 1.5)) <= float32((-2.25 / 0.1)))) && (!false)) {
        if len(v1) > 0 {
            v1[0] = 42
        }
        for fi44 := 0; fi44 < 11; fi44++ {
            fmt.Println(v3, v41)
            fmt.Println(v41, v6)
            _ = fi44
        }
    }
    v3 += v2
    var v45 float64 = 0.1
    var v46 uint16 = 845
    var v47 bool = (float32((((((((((1e10 / 1.0) + v45) * (float64(((0.0 - -2.25) - ((-2.25 / 1.5) - -2.25))) - v45)) - 0.1) + 3.14159) / 1.0) + 0.1) + v45) * v45)) <= 3.14159)
    if (42 > int32(v7)) {
        fmt.Println(v3)
        if (-2.25 <= (3.14159 - 3.14159)) {
            if v47 {
                fmt.Println(v4)
            }
            for fi48 := 0; fi48 < 1; fi48++ {
                v41 = "hello"
                _ = fi48
            }
            var v49 int8 = 0
            _ = v49
        } else {
            v50 := GenS7{F0: 174, F1: 9223372036854775807, F2: -9223372036854775808}
            _ = v50
        }
        _ = len(v1)
    } else {
        v51 := GenS8{F0: -105, F1: 2147483647, F2: 127}
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
