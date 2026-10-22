package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 string; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 string; F1 int32; F2 uint8 }
func genFn0(p0 uint8, p1 bool) int {
    p1 = p1
    var v5 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 int8 = -103
    v2 := []int{32767, 32767, 4294967295, -771}
    var v3 uint = 687
    v4 := []int{-924, 128, 32767}
    for fi6 := 0; fi6 < 19; fi6++ {
        if (!(!false)) {
            if (42 >= -1) {
                v7 := GenS1{F0: 32768}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 int64 = 42
                    var v10 int16 = 1000
                    v9 = -855
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                if ((3.14159 - (3.14159 + 0.0)) >= -2.25) {
                    v11 := GenS1{F0: 2147483647}
                    switch (fi6 >> ((v3) & 63)) {
                    case 1:
                        fmt.Println(v3, fi6)
                    case 121:
                        var v12 int32 = 42
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v14 := GenS2{F0: 0, F1: ("\n\t" + "hello")}
                        if len(v2) > 0 {
                            v2[0] = fi6
                        }
                        var v15 int16 = -953
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = v11
                } else {
                    v4 = append(v4, (fi6 - 0))
                    if ((3.14159 - 0.0) <= -2.25) {
                        for fi16 := 0; fi16 < 19; fi16++ {
                            v17 := GenS1{F0: fi16}
                            v1 -= (v1 - 0)
                            var v18 uint8 = 167
                            v19 := GenS2{F0: 0, F1: ((("foo bar" + "\n\t") + "\n\t") + "a")}
                            var v20 int = (fi16 << (((v3 ^ v3)) & 63))
                            _ = fi16
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        v3 = v3
                        fmt.Println(v3)
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                var v21 int = fi6
                _ = v21
            }
            if true {
                var v22 uint8 = 1
                v1 = v1
                _ = v22
            }
        } else {
            _ = len(v4)
            if ((v3 - 109) >= v3) {
                _ = len(v4)
                v23 := GenS3{F0: 42}
                _ = v23
            } else {
                fmt.Println(v1)
            }
            fmt.Println(v3)
        }
        var v24 uint32 = 2147483648
        _ = fi6
        _ = v24
    }
    v25 := GenS4{F0: (v3 % (((v3 - 4294967295)) | 1)), F1: "\n\t", F2: 126}
    var v26 uint32 = 2
    switch "a" {
    case "a":
        v1 = v1
    case "hello":
        v1 += v1
    case "":
        switch "a" {
        case "foo bar":
            switch "\n\t" {
            case "hello":
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
            case "\n\t":
                var v27 bool = true
                _ = v27
            case "a":
                for fi28 := 0; fi28 < 16; fi28++ {
                    v29 := GenS5{F0: v26, F1: 65535, F2: -1}
                    var v30 uint16 = uint16(v26)
                    var v31 uint = (^(v3))
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
            default:
                v32 := GenS5{F0: 1, F1: 363, F2: 42}
                _ = v32
            }
        case "hello":
            for fi33 := 0; fi33 < 13; fi33++ {
                v1 = (-(v1))
                _ = fi33
            }
        case "":
            v26 += v26
        default:
            fmt.Println(v3)
        }
    default:
        var v34 int = int(v3)
        _ = v34
    }
    _ = len(v2)
    v3 = (v3 | uint(v26))
    for fi35 := 0; fi35 < 1; fi35++ {
        if true {
            v26 -= 2147483647
            fmt.Println(v26, fi35)
        } else {
            var v36 uint8 = 255
            v1 -= (v1 ^ (v1 | -50))
            _ = v36
        }
        fmt.Println(v3, v3)
        v37 := GenS6{F0: 1, F1: uint16(v3)}
        _ = fi35
        _ = v37
    }
    v1 = (v1 & (v1 / (((v1 << (((v3 | (v3 >> ((1) & 63)))) & 63))) | 1)))
    fmt.Println(v26, v3, v1)
    if true {
        var v38 uint8 = 57
        v39 := GenS1{F0: -9223372036854775808}
        var v40 bool = false
        _ = v38
        _ = v39
        _ = v40
    }
    var v41 int64 = -768
    switch (((("a" + "\n\t") + "\n\t") + "hello") + "hello") {
    case "foo bar":
        var v42 uint = v3
        _ = v42
    default:
        v3 -= (v3 / (((v3 + 18446744073709551615)) | 1))
    }
    v43 := GenS6{F0: -709, F1: 2}
    v44 := GenS7{F0: 1, F1: 1}
    switch -722 {
    case 32768:
        if true {
            if (v41 != v41) {
                switch 1 {
                case 42:
                    var v45 int16 = -32768
                    _ = v45
                }
            } else {
                var v46 float32 = 1e10
                var v47 string = "hello"
                switch 32767 {
                case 4294967295:
                    v47 = v47
                case -9223372036854775808:
                    if (-2.25 != (float32(1.5) * (v46 * -2.25))) {
                        if len(v2) > 0 {
                            v2[0] = genFn0(18, (false && false))
                        }
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v4 = append(v4, 1)
                        } else {
                            if true {
                                for fi48 := 0; fi48 < 1; fi48++ {
                                    if true {
                                        v46 = 0.1
                                        v41 -= v41
                                        if false {
                                            var v49 float32 = (3.14159 / v46)
                                            fmt.Println(fi48, v26)
                                            _ = v49
                                        }
                                    } else {
                                        var v50 int = 4294967295
                                        v2 = append(v2, fi48)
                                        _ = v50
                                    }
                                    _ = fi48
                                }
                                var v51 uint64 = 1
                                _ = v51
                            } else {
                                v41 -= (-(v41))
                            }
                        }
                        var v52 int8 = 89
                        _ = v52
                    } else {
                        var v53 int64 = v41
                        v3 = v3
                        _ = v53
                    }
                }
                _ = v46
                _ = v47
            }
        }
    case 764:
        v41 = 9223372036854775807
    default:
        var v54 uint64 = 462
        _ = v54
    }
    for fi55 := 0; fi55 < 13; fi55++ {
        v56 := GenS8{F0: "hello", F1: 42, F2: uint8(v26)}
        if (-100 >= v1) {
            var v57 float64 = 1.5
            for fi58 := 0; fi58 < 13; fi58++ {
                var v59 int16 = 0
                var v60 float64 = 1e10
                v1 = v1
                _ = fi58
                _ = v59
                _ = v60
            }
            _ = v57
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi55
        _ = v56
    }
    var v61 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v44
    _ = v61
}
