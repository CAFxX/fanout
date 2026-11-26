package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int; F2 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 bool; F2 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = 1
    var v3 int32 = -445
    var v4 int32 = int32(v1)
    var v5 int = 9223372036854775807
    var v6 uint64 = 1
    switch v5 {
    case -1:
        var v7 string = ("a" + "hello")
        _ = v7
    default:
        var v8 uint32 = 116
        _ = v8
    }
    var v9 uint64 = (^(v6))
    v10 := []int{394}
    v5 -= (v5 % (((v5 ^ -18)) | 1))
    var v11 int = (v5 & (v5 & (-(v5))))
    var v12 string = (("a" + "\n\t") + "a")
    var v13 uint16 = 666
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 float64 = 3.14159
        v16 := GenS1{F0: v11}
        if ("" < ("foo bar" + v12)) {
            if true {
                if false {
                    var v17 int = (v11 / (((v5 % ((fi14) | 1))) | 1))
                    _ = v17
                }
                var v18 int = int(v1)
                _ = v18
            } else {
                if len(v10) > 0 {
                    v10[0] = (fi14 << ((4294967295) & 63))
                }
            }
            if false {
                _ = len(v10)
                var v19 int8 = (^(v2))
                _ = v19
            } else {
                if true {
                    if ((v9 * v6) > v9) {
                        v13 = 2
                        var v20 int = v11
                        _ = v20
                    } else {
                        v21 := GenS1{F0: (v11 << ((4294967295) & 63))}
                        v22 := GenS1{F0: -1}
                        _ = v21
                        _ = v22
                    }
                    v12 = ("a" + (((v12 + "foo bar") + (v12 + v12)) + v12))
                    v23 := GenS2{F0: 99, F1: 9223372036854775807, F2: v3}
                    _ = v23
                }
                v24 := GenS2{F0: 255, F1: 1, F2: (v4 ^ 0)}
                v13 = v13
                _ = v24
            }
            var v25 uint = 1
            _ = v25
        } else {
            var v26 uint64 = v9
            _ = v26
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v27 int8 = (v2 | v2)
    v9 = 291
    v28 := GenS3{F0: 2147483648, F1: 127, F2: v13}
    switch (-(v5)) {
    case -9223372036854775808:
        v29 := GenS1{F0: 32767}
        _ = v29
    case 32767:
        v30 := GenS4{F0: v12, F1: 223}
        _ = v30
    default:
        var v31 float64 = 0.0
        _ = v31
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := GenS5{F0: uint64(v13), F1: (^(v4)), F2: v6}
        v13 = 32767
        fmt.Println(v11, v3, v6)
        _ = fi32
        _ = v33
    }
    if (!true) {
        switch v11 {
        case 128:
            v5 = (v11 << ((4294967295) & 63))
        default:
            _ = len(v10)
        }
        var v34 uint8 = uint8(v6)
        var v35 int32 = v4
        _ = v34
        _ = v35
    } else {
        switch v12 {
        case "":
            if (v1 > -27) {
                v10 = append(v10, int(v4))
                var v36 int16 = 1000
                v4 = -2147483648
                _ = v36
            } else {
                fmt.Println(v11, v1, v6)
                v37 := GenS5{F0: (v6 << ((2) & 63)), F1: v4, F2: v9}
                for fi38 := 0; fi38 < 5; fi38++ {
                    switch v12 {
                    case "\n\t":
                        if (true && ((127 != 255) || ((fi38 / ((-9223372036854775808) | 1)) > fi38))) {
                            fmt.Println(v6, v11, v4)
                            v39 := GenS6{F0: uint(v6)}
                            switch 127 {
                            case -213:
                                v6 = (v6 + v9)
                            case 42:
                                if ((^(v1)) >= -804) {
                                    for fi40 := 0; fi40 < 12; fi40++ {
                                        if false {
                                            v4 = 1
                                            v2 = (-(v27))
                                        }
                                        v41 := GenS6{F0: 2}
                                        _ = fi40
                                        _ = v41
                                    }
                                }
                            default:
                                fmt.Println(v6, v1)
                            }
                            _ = v39
                        } else {
                            _ = len(v10)
                            var v42 string = v12
                            _ = v42
                        }
                    case "foo bar":
                        var v43 bool = false
                        _ = v43
                    case "a":
                        for fi44 := 0; fi44 < 18; fi44++ {
                            v45 := GenS4{F0: v12, F1: 187}
                            v4 = (v3 / (((-(v3))) | 1))
                            _ = fi44
                            _ = v45
                        }
                    }
                    if len(v10) > 0 {
                        v10[0] = (v11 - v11)
                    }
                    _ = fi38
                }
                _ = v37
            }
        case "hello":
            v46 := GenS5{F0: 42, F1: 42, F2: 4294967295}
            _ = v46
        default:
            v47 := GenS7{F0: (-(v27))}
            _ = v47
        }
        switch "foo bar" {
        case "":
            if len(v10) > 0 {
                v10[0] = v11
            }
        case "\n\t":
            var v48 int = v5
            _ = v48
        }
        v49 := GenS8{F0: v9, F1: false, F2: false}
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
