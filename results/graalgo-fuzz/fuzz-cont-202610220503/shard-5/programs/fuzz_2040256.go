package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64; F1 int; F2 int16 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 int64; F2 uint8 }
type GenS10 struct { F0 uint32; F1 uint64; F2 int32 }
func genFn0(p0 int16) uint64 {
    v5 := GenS1{F0: 2, F1: 148}
    v6 := []int{2147483647, 845, 128, 293}
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint64, p1 string) uint16 {
    var v7 int = 2147483647
    if true {
        var v8 uint64 = p0
        v9 := GenS2{F0: -9223372036854775808, F1: 255, F2: 2147483647}
        v10 := []int{(v7 ^ (v7 * (v7 + v7))), (v7 & 32768), 42}
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int = 950
    v12 := GenS3{F0: genFn0(-1), F1: -1, F2: 32767}
    var v13 uint16 = 805
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v7
    return v13
}

func main() {
    var v1 uint = 769
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    v4 := []string{""}
    var v14 uint16 = 400
    for fi15 := 0; fi15 < 6; fi15++ {
        fmt.Println(v14)
        _ = fi15
    }
    switch 1 {
    case -182:
        v1 -= v1
    default:
        var v16 int32 = 42
        _ = v16
    }
    switch (("a" + (("a" + (("" + ((("" + "hello") + "hello") + (("hello" + "a") + "a"))) + "hello")) + ("a" + "a"))) + "a") {
    case "a":
        fmt.Println(v1, v14)
    case "foo bar":
        v1 = v1
    case "\n\t":
        var v17 float32 = 1.5
        _ = v17
    default:
        if (818 <= 330) {
            var v18 float64 = 3.14159
            v19 := GenS2{F0: -9223372036854775808, F1: (v14 << (((v1 & (v1 | v1))) & 63)), F2: 1}
            _ = v18
            _ = v19
        } else {
            v20 := GenS4{F0: 523, F1: 0}
            _ = v20
        }
    }
    var v21 int16 = -15
    if ((-2.25 > 1.5) && ((((3.14159 / ((0.1 / (-2.25 / 3.14159)) * 3.14159)) / 0.1) / (1.5 * float32((0.1 + -2.25)))) < (float32(0.1) - (1e10 + (((((3.14159 + 1.5) + 3.14159) / float32(0.1)) * 0.1) - 0.0))))) {
        var v22 int16 = 42
        var v23 string = ("" + "")
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 int8 = -128
            _ = len(v4)
            _ = fi24
            _ = v25
        }
        if (!false) {
            var v26 float32 = 3.14159
            for fi27 := 0; fi27 < 11; fi27++ {
                for fi28 := 0; fi28 < 5; fi28++ {
                    fmt.Println(fi27, v1)
                    _ = fi28
                }
                _ = fi27
            }
            _ = v26
        }
        if false {
            var v29 uint16 = (v14 << ((385) & 63))
            if (false || (!true)) {
                switch ("hello" + ("hello" + "a")) {
                case "\n\t":
                    fmt.Println(v14)
                case "foo bar":
                    if (42 >= 2147483647) {
                        var v30 uint = (v1 & v1)
                        v29 += v14
                        fmt.Println(v14, v1)
                        _ = v30
                    }
                }
                if (v29 == (v14 + v29)) {
                    if true {
                        if (!true) {
                            v31 := GenS2{F0: -9223372036854775808, F1: 1, F2: 2147483647}
                            for fi32 := 0; fi32 < 3; fi32++ {
                                if false {
                                    if true {
                                        v33 := GenS5{F0: -2147483648, F1: -83, F2: fi32}
                                        if true {
                                            var v34 bool = true
                                            v21 = (v21 & (v21 + v21))
                                            for fi35 := 0; fi35 < 19; fi35++ {
                                                v36 := v3[(fi32 | (fi32 >> ((v1) & 63)))]
                                                _ = fi35
                                                _ = v36
                                            }
                                            _ = v34
                                        } else {
                                            v1 -= v1
                                            fmt.Println(v14, v14, v21)
                                            delete(v2, -9223372036854775808)
                                        }
                                        v37 := GenS2{F0: -1, F1: 888, F2: -1}
                                        _ = v33
                                        _ = v37
                                    }
                                    v38 := v2[1]
                                    v39 := GenS6{F0: 42}
                                    _ = v38
                                    _ = v39
                                } else {
                                    delete(v3, fi32)
                                }
                                var v40 float32 = (-2.25 - 0.0)
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    if false {
                                        v42 := GenS7{F0: 0, F1: 0, F2: 0}
                                        v14 = v14
                                        _ = v42
                                    } else {
                                        if true {
                                            v40 += (-2.25 * 1.5)
                                        }
                                        var v43 int8 = -128
                                        _ = v43
                                    }
                                    switch -9223372036854775808 {
                                    case 525:
                                        v14 -= (v29 + (v29 * v29))
                                    case -111:
                                        if ((v29 & genFn1(uint64(v1), "")) < (v29 + (^(v29)))) {
                                            var v44 uint = 1
                                            _ = v44
                                        }
                                    case -2147483648:
                                        var v45 bool = (!true)
                                        _ = v45
                                    }
                                    _ = fi41
                                }
                                _ = fi32
                                _ = v40
                            }
                            _ = v31
                        } else {
                            switch ("\n\t" + "foo bar") {
                            case "":
                                v29 += (v29 * (v29 - (v29 << ((v1) & 63))))
                            }
                            v21 += (v21 | (^(v21)))
                        }
                        if (!true) {
                            for fi46 := 0; fi46 < 12; fi46++ {
                                _ = len(v3)
                                _ = fi46
                            }
                            v47 := GenS7{F0: 128, F1: int64(v1), F2: 4294967295}
                            if false {
                                v21 -= (v21 % ((v21) | 1))
                                v29 = v14
                                fmt.Println(v29, v14)
                            }
                            _ = v47
                        }
                    }
                }
                for fi48 := 0; fi48 < 11; fi48++ {
                    v21 += v21
                    if (730 == 793) {
                        _ = len(v2)
                        v49 := GenS8{F0: v21}
                        var v50 int32 = -589
                        _ = v49
                        _ = v50
                    }
                    _ = fi48
                }
            }
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                var v51 uint = 18446744073709551615
                _ = v51
            }
            _ = v29
        }
    }
    v52 := GenS6{F0: 42}
    v53 := GenS9{F0: -435, F1: -400, F2: 255}
    if true {
        v54 := GenS10{F0: 417, F1: 18446744073709551615, F2: 0}
        var v55 bool = (4294967295 > 1)
        _ = v54
        _ = v55
    } else {
        var v56 float64 = 1e10
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v56
    }
    switch "" {
    case "\n\t":
        v1 += v1
    }
    var v57 uint = v1
    v1 += v1
    for fi58 := 0; fi58 < 9; fi58++ {
        v1 += 42
        var v59 float32 = 3.14159
        _ = fi58
        _ = v59
    }
    v60 := GenS10{F0: uint32(v57), F1: 42, F2: 0}
    switch ("hello" + ((("a" + ("hello" + "a")) + "hello") + "hello")) {
    case "hello":
        v1 = 4294967295
    case "a":
        fmt.Println(v21, v14)
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v57
    _ = v60
}
