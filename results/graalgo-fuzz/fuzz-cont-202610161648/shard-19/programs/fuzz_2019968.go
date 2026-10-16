package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8 }
func genFn0() uint32 {
    var v7 int8 = 0
    v7 += (v7 % ((v7) | 1))
    switch "" {
    case "":
        if (!(-1 >= 623)) {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := []int{-119, fi8, (fi8 * (fi8 % (((fi8 >> ((4294967295) & 63))) | 1)))}
                v7 = 42
                if len(v9) > 0 {
                    v9[0] = (fi8 >> ((18446744073709551615) & 63))
                }
                _ = fi8
                _ = v9
            }
        } else {
            var v10 bool = true
            _ = v10
        }
    case "\n\t":
        switch 127 {
        case 1:
            v11 := GenS1{F0: -2147483648, F1: false, F2: 82}
            _ = v11
        case 32767:
            switch -2147483648 {
            case 2147483647:
                v12 := GenS1{F0: -467, F1: true, F2: 170}
                _ = v12
            }
        default:
            v13 := []int{}
            _ = v13
        }
    default:
        v14 := []int{-9223372036854775808, 370, 32768, 42}
        _ = v14
    }
    _ = v7
    return 872
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483648, 2147483647, 2147483647}
    v3 := make(map[int]string)
    var v4 int64 = 42
    var v5 string = "foo bar"
    v6 := []uint32{2, 0, 42, 173}
    for fi15 := 0; fi15 < 19; fi15++ {
        if (255 == 1) {
            for fi16 := 0; fi16 < 8; fi16++ {
                v2 = append(v2, 161)
                v4 = (v4 - v4)
                switch v5 {
                case "":
                    if true {
                        v17 := GenS1{F0: (fi15 + fi15), F1: true, F2: 241}
                        if true {
                            var v18 uint64 = 0
                            var v19 int32 = 2147483647
                            if true {
                                fmt.Println(v19)
                                v19 += v19
                                if true {
                                    if false {
                                        v20 := GenS2{F0: -400, F1: 42}
                                        var v21 int8 = 1
                                        if ((v19 / (((v19 ^ v19)) | 1)) < 0) {
                                            fmt.Println(fi15, v21)
                                        }
                                        switch (("" + v5) + (v5 + v5)) {
                                        case "":
                                            var v22 float32 = -2.25
                                            _ = v22
                                        case "hello":
                                            v5 = ("hello" + v5)
                                        }
                                        _ = v20
                                        _ = v21
                                    }
                                    v19 -= (v19 / (((v19 * int32(v18))) | 1))
                                    fmt.Println(v18, v19, fi16)
                                } else {
                                    if (-9223372036854775808 < fi15) {
                                        fmt.Println(v4, v5)
                                        v4 += (v4 >> ((4294967295) & 63))
                                    } else {
                                        if (!true) {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v5 = (v5 + v5)
                                            switch fi16 {
                                            case 9223372036854775807:
                                                var v23 float32 = ((1.5 * (float32(1e10) + (3.14159 / 1.0))) + 1.5)
                                                _ = v23
                                            default:
                                                switch (fi16 - fi15) {
                                                case 4294967295:
                                                    if true {
                                                        fmt.Println(v19)
                                                    }
                                                case 611:
                                                    var v24 uint8 = 128
                                                    _ = v24
                                                case 462:
                                                    var v25 int8 = -34
                                                    _ = v25
                                                default:
                                                    var v26 bool = true
                                                    v18 = v18
                                                    v26 = false
                                                    _ = v26
                                                }
                                            }
                                        } else {
                                            _ = len(v1)
                                        }
                                        fmt.Println(v5)
                                        var v27 uint8 = 2
                                        _ = v27
                                    }
                                }
                            } else {
                                v4 = v4
                                var v28 string = (v5 + v5)
                                _ = v28
                            }
                            _ = v18
                            _ = v19
                        } else {
                            v29 := GenS3{F0: 2147483648}
                            fmt.Println(fi15, v5)
                            var v30 int = 127
                            switch v5 {
                            case "\n\t":
                                switch v5 {
                                case "\n\t":
                                    v31 := GenS1{F0: fi16, F1: (!(!(true && false))), F2: 128}
                                    _ = v31
                                default:
                                    v32 := GenS1{F0: -890, F1: (!true), F2: 0}
                                    _ = v32
                                }
                            case "foo bar":
                                v5 = ("hello" + "")
                            case "a":
                                switch v5 {
                                case "\n\t":
                                    v33 := GenS2{F0: 415, F1: 32767}
                                    _ = v33
                                case "hello":
                                    v30 = (fi15 + (v30 / ((fi15) | 1)))
                                }
                            default:
                                v30 = fi16
                            }
                            if false {
                                if false {
                                    var v34 uint16 = 36
                                    if ((true || (3.14159 <= float32(-2.25))) || false) {
                                        v5 = v5
                                        v5 = ""
                                        v4 = int64(v34)
                                        v35 := GenS4{F0: v34, F1: (v34 | v34)}
                                        fmt.Println(v5, v5)
                                        _ = v35
                                    }
                                    _ = v34
                                }
                                fmt.Println(fi16)
                                switch fi16 {
                                case 127:
                                    v5 = (v5 + (v5 + v5))
                                case -923:
                                    v5 = ("" + ("\n\t" + v5))
                                }
                                if true {
                                    var v36 int = fi16
                                    v30 = 103
                                    _ = v36
                                } else {
                                    switch ("a" + v5) {
                                    case "":
                                        v5 = ("a" + (v5 + (((v5 + v5) + ("a" + v5)) + (v5 + v5))))
                                    }
                                }
                            }
                            _ = v29
                            _ = v30
                        }
                        _ = v17
                    } else {
                        var v37 bool = (false && (179 < 18446744073709551615))
                        fmt.Println(v4, v37, fi15)
                        v38 := GenS1{F0: -944, F1: (((v5 == v5) || (false && (!v37))) && (0 < 1)), F2: 255}
                        v39 := GenS4{F0: 42, F1: 46}
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                case "\n\t":
                    delete(v1, fi16)
                case "a":
                    switch fi16 {
                    case 127:
                        var v40 float64 = 0.1
                        _ = v40
                    case 128:
                        var v41 int8 = 100
                        switch ("\n\t" + v5) {
                        case "":
                            v42 := GenS1{F0: 1, F1: ((18446744073709551615 >= 4294967295) || true), F2: 2}
                            _ = v42
                        case "\n\t":
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                            }
                        case "a":
                            var v43 string = v5
                            _ = v43
                        default:
                            _ = len(v3)
                        }
                        _ = v41
                    }
                }
                _ = fi16
            }
        } else {
            v44 := GenS5{F0: 42}
            _ = v44
        }
        if ((true || true) && (true && false)) {
            v4 = 9223372036854775807
            fmt.Println(fi15, v4)
            v4 = 9223372036854775807
        }
        if (false && true) {
            for fi45 := 0; fi45 < 4; fi45++ {
                v46 := v3[565]
                v5 = (v46 + v46)
                _ = fi45
                _ = v46
            }
        }
        _ = fi15
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        v5 = (v5 + "\n\t")
        _ = fi47
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        switch (^(fi48)) {
        case 864:
            if len(v2) > 0 {
                v2[0] = 216
            }
        case 2147483647:
            for fi49 := 0; fi49 < 18; fi49++ {
                v5 = v5
                v50 := GenS6{F0: false}
                fmt.Println(fi48, fi49, fi49)
                _ = fi49
                _ = v50
            }
        case -623:
            v51 := GenS7{F0: -74}
            _ = v51
        default:
            v4 += (v4 >> ((255) & 63))
        }
        v4 = (v4 - 1)
        _ = fi48
    }
    v5 = v5
    fmt.Println(v5, v4)
    _ = len(v6)
    v4 -= -9223372036854775808
    var v52 uint16 = 504
    var v53 int16 = int16(v4)
    var v54 uint = 4294967295
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
