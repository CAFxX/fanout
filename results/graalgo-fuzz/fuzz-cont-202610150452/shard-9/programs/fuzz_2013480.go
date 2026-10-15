package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int; F2 bool }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 uint64; F1 int64; F2 bool }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 int8; F1 int16 }

func main() {
    v1 := []uint32{}
    v2 := []int{681}
    var v3 bool = (369 <= 634)
    var v4 uint16 = 2
    fmt.Println(v4, v3)
    var v5 uint32 = 1
    fmt.Println(v4, v3, v4)
    fmt.Println(v5, v4)
    fmt.Println(v3, v5)
    for fi6 := 0; fi6 < 6; fi6++ {
        if ("" >= "a") {
            if false {
                switch (fi6 / (((fi6 ^ fi6)) | 1)) {
                case 128:
                    switch fi6 {
                    case 32768:
                        _ = len(v2)
                    case 1:
                        v7 := GenS1{F0: 0}
                        _ = v7
                    case 965:
                        for fi8 := 0; fi8 < 20; fi8++ {
                            _ = len(v1)
                            v3 = false
                            _ = fi8
                        }
                    default:
                        fmt.Println(fi6)
                    }
                case 696:
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v5 = (v5 >> ((1) & 63))
                        _ = fi9
                    }
                case -9223372036854775808:
                    for fi10 := 0; fi10 < 17; fi10++ {
                        var v11 int64 = -1
                        if v3 {
                            v12 := GenS1{F0: 100}
                            v13 := GenS1{F0: 1}
                            if v3 {
                                if len(v2) > 0 {
                                    v2[0] = (-(fi6))
                                }
                            } else {
                                if ((float64(3.14159) == 0.0) || v3) {
                                    var v14 int8 = 127
                                    v2 = append(v2, (fi6 % ((fi6) | 1)))
                                    if v3 {
                                        v5 += v5
                                        v5 = (v5 >> ((42) & 63))
                                    } else {
                                        v15 := GenS2{F0: v14}
                                        _ = v15
                                    }
                                    _ = v14
                                } else {
                                    v2 = append(v2, (-(fi6)))
                                    v1 = append(v1, (v5 % ((v5) | 1)))
                                }
                                var v16 uint = 2
                                _ = v16
                            }
                            _ = v12
                            _ = v13
                        } else {
                            v17 := GenS3{F0: -2147483648}
                            var v18 uint32 = v5
                            if v3 {
                                var v19 uint64 = 18446744073709551615
                                _ = v19
                            } else {
                                v18 += (v18 * (^(v18)))
                                fmt.Println(v18, v11)
                            }
                            _ = v17
                            _ = v18
                        }
                        _ = len(v2)
                        var v20 uint8 = 2
                        v3 = (!v3)
                        if len(v1) > 0 {
                            v1[0] = (v5 << ((541) & 63))
                        }
                        _ = fi10
                        _ = v11
                        _ = v20
                    }
                }
                v21 := GenS4{F0: 0, F1: -1, F2: (!(-2147483648 != -447))}
                if ((float64(3.14159) != ((-2.25 + (float64(0.1) * 3.14159)) - float64(((1.5 + 3.14159) * 0.1)))) && ((!(!true)) || ((float32(3.14159) - float32((-2.25 / -2.25))) == ((0.1 * 0.1) / (1e10 + float32(3.14159)))))) {
                    if v3 {
                        for fi22 := 0; fi22 < 2; fi22++ {
                            fmt.Println(fi22, v5, v4)
                            v4 += v4
                            if (v3 || (!true)) {
                                v4 -= 32767
                            }
                            _ = fi22
                        }
                    } else {
                        fmt.Println(v5)
                    }
                    switch (fi6 + (fi6 % ((fi6) | 1))) {
                    case 1:
                        for fi23 := 0; fi23 < 9; fi23++ {
                            v5 -= v5
                            switch ("hello" + ((("" + "foo bar") + "\n\t") + "hello")) {
                            case "hello":
                                if v3 {
                                    v24 := GenS5{F0: "a", F1: uint64(v4)}
                                    _ = v24
                                } else {
                                    var v25 int64 = 199
                                    v5 -= v5
                                    var v26 uint16 = 51
                                    fmt.Println(v4)
                                    if len(v1) > 0 {
                                        v1[0] = 956
                                    }
                                    v27 := GenS6{F0: fi23, F1: v4}
                                    var v28 int64 = (v25 / ((v25) | 1))
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            case "\n\t":
                                if (-493 != -9223372036854775808) {
                                    if len(v1) > 0 {
                                        v1[0] = 319
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v29 uint32 = (v5 | (v5 & (v5 / ((1) | 1))))
                                    _ = v29
                                }
                            case "":
                                switch "\n\t" {
                                case "hello":
                                    v30 := GenS3{F0: 1}
                                    _ = v30
                                case "foo bar":
                                    switch ("a" + ("a" + ("foo bar" + "hello"))) {
                                    case "":
                                        if v3 {
                                            v31 := GenS7{F0: 4294967295, F1: -9223372036854775808, F2: v3}
                                            switch (-(fi23)) {
                                            case 127:
                                                if len(v2) > 0 {
                                                    v2[0] = (fi6 >> ((24) & 63))
                                                }
                                            }
                                            _ = v31
                                        }
                                    default:
                                        v32 := GenS2{F0: 127}
                                        _ = v32
                                    }
                                }
                            }
                            var v33 int8 = -128
                            _ = fi23
                            _ = v33
                        }
                    default:
                        v34 := GenS7{F0: 18446744073709551615, F1: 9223372036854775807, F2: v3}
                        _ = v34
                    }
                } else {
                    if v3 {
                        var v35 bool = ((v4 >= v4) || false)
                        switch "a" {
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = 207
                            }
                        default:
                            var v36 string = "foo bar"
                            _ = v36
                        }
                        _ = v35
                    }
                    fmt.Println(fi6, v3, fi6)
                }
                _ = v21
            }
            fmt.Println(fi6, v5, v3)
            v37 := GenS7{F0: 18446744073709551615, F1: -940, F2: true}
            _ = v37
        }
        fmt.Println(v4)
        _ = fi6
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        if len(v1) > 0 {
            v1[0] = 5
        }
        _ = fi38
    }
    fmt.Println(v3, v4, v5)
    v3 = ((560 > 886) && false)
    for fi39 := 0; fi39 < 17; fi39++ {
        v2 = append(v2, (fi39 % ((fi39) | 1)))
        v40 := GenS8{F0: v5, F1: 0}
        v41 := GenS3{F0: 0}
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v5, v3)
    if v3 {
        v5 += (v5 * 959)
        var v42 string = "a"
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v42
    } else {
        var v43 uint = 4294967295
        switch 4294967295 {
        case -593:
            if len(v2) > 0 {
                v2[0] = -53
            }
        case 127:
            fmt.Println(v4, v5)
        }
        for fi44 := 0; fi44 < 8; fi44++ {
            fmt.Println(v3, v5, fi44)
            _ = fi44
        }
        _ = v43
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        for fi46 := 0; fi46 < 14; fi46++ {
            v3 = v3
            fmt.Println(fi45, fi45)
            _ = fi46
        }
        for fi47 := 0; fi47 < 18; fi47++ {
            if len(v2) > 0 {
                v2[0] = (-(fi47))
            }
            _ = fi47
        }
        _ = fi45
    }
    var v48 int64 = -1
    for fi49 := 0; fi49 < 15; fi49++ {
        v4 += 42
        if (613 != 0) {
            fmt.Println(v4, v3)
            if len(v2) > 0 {
                v2[0] = (-(fi49))
            }
        } else {
            v50 := GenS8{F0: (v5 - (v5 >> ((0) & 63))), F1: 18446744073709551615}
            _ = v50
        }
        fmt.Println(v4)
        _ = fi49
    }
    if true {
        for fi51 := 0; fi51 < 17; fi51++ {
            v5 += v5
            v52 := GenS3{F0: -2147483648}
            v53 := GenS9{F0: -55, F1: 1}
            _ = fi51
            _ = v52
            _ = v53
        }
    } else {
        _ = len(v1)
        var v54 float32 = 0.1
        _ = v54
    }
    v48 -= v48
    fmt.Println(v4, v5, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
