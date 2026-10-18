package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int; F1 int32; F2 int32 }
type GenS5 struct { F0 int64; F1 uint8; F2 int64 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 string; F1 string; F2 int }

func main() {
    v1 := []int{1, 4294967295, 9223372036854775807}
    var v2 int64 = 510
    v3 := []int{}
    var v4 uint16 = 640
    var v5 int32 = 42
    switch int(v5) {
    case 683:
        v6 := GenS1{F0: (v2 % ((v2) | 1)), F1: 673}
        _ = v6
    }
    v2 -= v2
    v7 := GenS1{F0: (v2 + (v2 >> ((4294967295) & 63))), F1: 4294967295}
    var v8 float32 = float32(1e10)
    var v9 string = "foo bar"
    if true {
        switch "foo bar" {
        case "":
            var v10 int64 = v2
            _ = v10
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            if (2 >= 515) {
                if true {
                    var v11 int8 = 68
                    var v12 uint = 0
                    v5 = (v5 | v5)
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 15; fi13++ {
                    if (v8 != v8) {
                        v9 = (v9 + (v9 + v9))
                        if true {
                            switch "a" {
                            case "foo bar":
                                v9 = v9
                            }
                        } else {
                            if (!(!false)) {
                                v2 = 9223372036854775807
                                v3 = append(v3, (fi13 + (fi13 >> ((255) & 63))))
                                fmt.Println(fi13, fi13)
                            } else {
                                _ = len(v1)
                            }
                            switch (v9 + v9) {
                            case "hello":
                                v14 := GenS2{F0: -1, F1: 459, F2: 44}
                                _ = v14
                            }
                            var v15 string = "\n\t"
                            _ = v15
                        }
                        switch "\n\t" {
                        case "a":
                            var v16 uint8 = 2
                            _ = v16
                        case "\n\t":
                            fmt.Println(v9)
                        case "foo bar":
                            if true {
                                v4 += v4
                                v17 := GenS1{F0: v2, F1: fi13}
                                _ = v17
                            }
                        }
                    } else {
                        v2 -= v2
                        var v18 int = fi13
                        for fi19 := 0; fi19 < 3; fi19++ {
                            var v20 int = (fi13 >> ((0) & 63))
                            v2 = -9223372036854775808
                            if true {
                                _ = len(v3)
                                if false {
                                    v5 -= int32(v2)
                                    v1 = append(v1, v20)
                                } else {
                                    var v21 float64 = float64((float32(3.14159) - v8))
                                    if len(v3) > 0 {
                                        v3[0] = 143
                                    }
                                    v4 -= (^(v4))
                                    _ = v21
                                }
                                v22 := GenS2{F0: 42, F1: 2147483648, F2: 1}
                                _ = v22
                            }
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    }
                    _ = fi13
                }
            }
        default:
            if (-1 < 1) {
                if true {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if ((fi23 >> ((97) & 63)) == fi23) {
                            v24 := GenS3{F0: 32767, F1: v9}
                            var v25 uint64 = 2
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(v8, v9)
                        }
                        if (!false) {
                            if (2 != -96) {
                                var v26 int = -472
                                v26 = -9223372036854775808
                                _ = v26
                            }
                            v8 = (v8 / (v8 + 0.0))
                        }
                        _ = fi23
                    }
                    if true {
                        v4 = (v4 * v4)
                        if true {
                            v27 := GenS4{F0: 4294967295, F1: (v5 / (((-(v5))) | 1)), F2: (v5 * 42)}
                            if len(v3) > 0 {
                                v3[0] = -1
                            }
                            for fi28 := 0; fi28 < 16; fi28++ {
                                if (65535 <= (v4 | (v4 >> ((4294967295) & 63)))) {
                                    fmt.Println(fi28, v5)
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        v30 := GenS5{F0: (v2 * (v2 + 42)), F1: 127, F2: -975}
                                        v31 := GenS2{F0: 0, F1: 42, F2: 1}
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                } else {
                                    var v32 int64 = (v2 ^ (v2 >> ((2) & 63)))
                                    switch fi28 {
                                    case -142:
                                        if len(v3) > 0 {
                                            v3[0] = (fi28 % (((-(fi28))) | 1))
                                        }
                                    case 2147483647:
                                        if ((1.5 / 1e10) != 0.0) {
                                            var v33 string = (v9 + v9)
                                            var v34 uint32 = 2
                                            v8 = ((-2.25 - v8) - v8)
                                            _ = v33
                                            _ = v34
                                        }
                                    default:
                                        v35 := GenS6{F0: 42, F1: 42}
                                        _ = v35
                                    }
                                    fmt.Println(v5, v4)
                                    _ = v32
                                }
                                v2 = (^(v2))
                                _ = fi28
                            }
                            _ = v27
                        }
                    }
                    v36 := GenS7{F0: v9, F1: (v9 + "\n\t"), F2: 32768}
                    _ = v36
                }
                var v37 string = v9
                if len(v1) > 0 {
                    v1[0] = 42
                }
                _ = v37
            } else {
                fmt.Println(v4, v2, v2)
                for fi38 := 0; fi38 < 17; fi38++ {
                    if false {
                        _ = len(v3)
                    }
                    fmt.Println(fi38, v2)
                    if len(v3) > 0 {
                        v3[0] = 875
                    }
                    _ = fi38
                }
            }
        }
    } else {
        fmt.Println(v9, v5)
        for fi39 := 0; fi39 < 20; fi39++ {
            for fi40 := 0; fi40 < 7; fi40++ {
                var v41 string = (v9 + v9)
                _ = fi40
                _ = v41
            }
            v8 = 1e10
            _ = fi39
        }
        v42 := GenS2{F0: 1000, F1: 2147483647, F2: 255}
        _ = v42
    }
    v4 -= (v4 * v4)
    switch -986 {
    case 127:
        v43 := GenS3{F0: -32768, F1: (v9 + v9)}
        _ = v43
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        for fi45 := 0; fi45 < 8; fi45++ {
            var v46 int8 = 0
            var v47 uint32 = 2147483648
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v5 = int32(v2)
            _ = fi45
            _ = v46
            _ = v47
        }
        if len(v1) > 0 {
            v1[0] = 0
        }
        v48 := GenS2{F0: 616, F1: 65535, F2: 65}
        _ = fi44
        _ = v48
    }
    v49 := GenS5{F0: (v2 >> ((18446744073709551615) & 63)), F1: 255, F2: -9223372036854775808}
    var v50 int32 = (v5 & v5)
    var v51 int8 = 100
    v4 = (v4 + v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v7
    _ = v8
    _ = v9
}
