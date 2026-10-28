package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int64; F1 uint; F2 uint }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 uint64 }

func main() {
    var v1 int64 = -873
    var v2 uint64 = 0
    var v3 uint64 = v2
    var v4 uint32 = 2
    var v5 uint = 4294967295
    v6 := []int{128, 0, 42}
    var v7 uint = v5
    var v8 float64 = (1e10 + 1e10)
    _ = len(v6)
    for fi9 := 0; fi9 < 3; fi9++ {
        v6 = append(v6, -471)
        v6 = append(v6, (fi9 << (((v7 - (v7 << ((2) & 63)))) & 63)))
        for fi10 := 0; fi10 < 6; fi10++ {
            if (true && true) {
                fmt.Println(fi10)
                _ = len(v6)
                if len(v6) > 0 {
                    v6[0] = -572
                }
            } else {
                var v11 int32 = 1
                _ = v11
            }
            _ = fi10
        }
        _ = fi9
    }
    v2 = 1
    var v12 int32 = 929
    for fi13 := 0; fi13 < 5; fi13++ {
        v5 -= (v5 / ((uint(v4)) | 1))
        if (true && ((v1 | -955) > (-(v1)))) {
            var v14 int64 = v1
            _ = v14
        } else {
            v15 := GenS1{F0: 128, F1: (v1 % ((v1) | 1)), F2: (v12 * v12)}
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS2{F0: (v2 | 325)}
            v2 += (v2 | v2)
            fmt.Println(v7, v3)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v4 = 954
    if (3.14159 != ((1e10 * (1.5 * 1e10)) * ((-2.25 / ((3.14159 - (0.1 + float32(1e10))) * float32(v8))) / 1.5))) {
        var v19 uint64 = v3
        _ = v19
    } else {
        v6 = append(v6, -1)
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
        if (false || (!false)) {
            if len(v6) > 0 {
                v6[0] = 224
            }
            if (v1 < 1) {
                fmt.Println(v4)
                v20 := GenS2{F0: 4294967295}
                switch -329 {
                case 260:
                    v5 += 2
                case 42:
                    _ = len(v6)
                case -798:
                    fmt.Println(v5, v1)
                default:
                    var v21 int64 = (^(v1))
                    _ = v21
                }
                _ = v20
            }
            switch "\n\t" {
            case "a":
                switch 662 {
                case 4294967295:
                    v22 := GenS3{F0: -991}
                    _ = v22
                case 9223372036854775807:
                    var v23 string = ""
                    _ = v23
                default:
                    switch "" {
                    case "a":
                        switch 4294967295 {
                        case 42:
                            fmt.Println(v5, v3, v1)
                        case 571:
                            switch "hello" {
                            case "a":
                                v8 = (-2.25 * v8)
                            case "\n\t":
                                v7 = uint(v4)
                            case "foo bar":
                                v2 = v3
                            default:
                                var v24 bool = ((v5 | 2) >= (v5 / (((v5 + (^(v7)))) | 1)))
                                _ = v24
                            }
                        case 0:
                            switch int(v1) {
                            case 980:
                                if false {
                                    v7 = 203
                                } else {
                                    var v25 int = 107
                                    var v26 uint = v7
                                    if (((^(v1)) >= -1) || (!true)) {
                                        fmt.Println(v2, v5, v26)
                                        switch (v25 & 32768) {
                                        case 2147483647:
                                            fmt.Println(v1, v26, v3)
                                        case -2147483648:
                                            var v27 uint = (v7 % ((v7) | 1))
                                            _ = v27
                                        }
                                    } else {
                                        for fi28 := 0; fi28 < 6; fi28++ {
                                            if true {
                                                if (true && (0.1 > 1.5)) {
                                                    v29 := GenS4{F0: 46, F1: (^(v3)), F2: (243 > 982)}
                                                    var v30 uint8 = 128
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                            v31 := GenS5{F0: (fi28 >> (((v5 & 255)) & 63))}
                                            _ = fi28
                                            _ = v31
                                        }
                                        if false {
                                            if true {
                                                if (false && true) {
                                                    v5 += v7
                                                    v5 = v26
                                                    if true {
                                                        fmt.Println(v7)
                                                    } else {
                                                        v6 = append(v6, (v25 / ((v25) | 1)))
                                                        for fi32 := 0; fi32 < 16; fi32++ {
                                                            switch (("" + "hello") + ("a" + "\n\t")) {
                                                            case "foo bar":
                                                                v33 := GenS6{F0: (v26 - (v26 | (^(v7)))), F1: 0}
                                                                _ = v33
                                                            case "a":
                                                                var v34 int8 = 18
                                                                _ = v34
                                                            case "hello":
                                                                for fi35 := 0; fi35 < 3; fi35++ {
                                                                    v36 := GenS7{F0: fi32, F1: 0}
                                                                    v2 -= uint64(v4)
                                                                    _ = fi35
                                                                    _ = v36
                                                                }
                                                            }
                                                            for fi37 := 0; fi37 < 5; fi37++ {
                                                                v3 -= v2
                                                                _ = fi37
                                                            }
                                                            fmt.Println(v3, v4, v12)
                                                            _ = fi32
                                                        }
                                                        var v38 uint64 = 422
                                                        _ = v38
                                                    }
                                                }
                                                fmt.Println(v2, v2)
                                            } else {
                                                switch (v25 ^ (v25 >> ((v7) & 63))) {
                                                case 0:
                                                    v12 = -1
                                                case 9223372036854775807:
                                                    switch 32767 {
                                                    case 104:
                                                        var v39 int8 = 100
                                                        _ = v39
                                                    case 2147483647:
                                                        fmt.Println(v8, v4, v1)
                                                    case -1:
                                                        if len(v6) > 0 {
                                                            v6[0] = (v25 & v25)
                                                        }
                                                    }
                                                }
                                            }
                                            v6 = append(v6, v25)
                                        }
                                    }
                                    _ = v25
                                    _ = v26
                                }
                            case 810:
                                fmt.Println(v12, v12)
                            default:
                                var v40 float64 = v8
                                _ = v40
                            }
                        }
                    case "foo bar":
                        v41 := GenS8{F0: (^(v1)), F1: 644, F2: (^(v7))}
                        _ = v41
                    case "\n\t":
                        v42 := GenS9{F0: v2, F1: "a"}
                        _ = v42
                    default:
                        fmt.Println(v8, v8)
                    }
                }
            default:
                v43 := GenS3{F0: 975}
                _ = v43
            }
        } else {
            for fi44 := 0; fi44 < 2; fi44++ {
                switch (fi44 * fi44) {
                case 4294967295:
                    var v45 int64 = (^(v1))
                    _ = v45
                default:
                    for fi46 := 0; fi46 < 3; fi46++ {
                        v47 := GenS10{F0: 0}
                        var v48 int64 = 365
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                }
                v6 = append(v6, -9223372036854775808)
                _ = fi44
            }
        }
    }
    var v49 uint8 = 49
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
