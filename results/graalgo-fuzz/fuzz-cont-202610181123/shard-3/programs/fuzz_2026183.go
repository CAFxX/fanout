package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int64; F1 int32; F2 int }
type GenS7 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 int64 = -1
    var v5 uint64 = 0
    var v6 float64 = 0.0
    v7 := []int{-9223372036854775808}
    var v8 string = "foo bar"
    v5 = (v5 ^ (v5 & 4294967295))
    var v9 string = v8
    v9 = (("" + "hello") + ("foo bar" + "a"))
    if true {
        fmt.Println(v9, v4)
    } else {
        var v10 float32 = 0.0
        v8 = "hello"
        var v11 float64 = 3.14159
        _ = v10
        _ = v11
    }
    var v12 float64 = (v6 - float64((1e10 / (float32(0.1) * (1e10 / ((-2.25 / (1e10 + ((float32(v6) - 1.5) + 0.0))) * 1.5))))))
    v9 = v9
    for fi13 := 0; fi13 < 18; fi13++ {
        v7 = append(v7, (fi13 / ((fi13) | 1)))
        _ = fi13
    }
    v14 := GenS1{F0: 1}
    if false {
        _ = len(v1)
        v15 := GenS2{F0: (-(v4)), F1: -1, F2: "\n\t"}
        _ = v15
    } else {
        switch 1 {
        case 127:
            var v16 uint = 255
            _ = v16
        case 366:
            var v17 int64 = v4
            _ = v17
        default:
            if (42 < 2) {
                var v18 float32 = 3.14159
                if (0 <= -160) {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        for fi20 := 0; fi20 < 6; fi20++ {
                            var v21 string = (v8 + v8)
                            v22 := GenS3{F0: v8}
                            v21 = ((v21 + (v8 + v21)) + ((("" + v21) + "") + v9))
                            v9 = "foo bar"
                            var v23 float32 = (v18 + (v18 + ((0.1 * (((v18 * v18) * 0.0) + (0.0 - 1e10))) - 1e10)))
                            var v24 int8 = 42
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        v8 = v9
                        var v25 int16 = -311
                        _ = fi19
                        _ = v25
                    }
                    if (42 > 11) {
                        v26 := GenS4{F0: v4, F1: (v5 >> ((65535) & 63)), F2: 32767}
                        _ = v26
                    } else {
                        fmt.Println(v6)
                        switch -158 {
                        case -712:
                            v27 := GenS5{F0: 2147483647, F1: 971}
                            _ = v27
                        case -799:
                            var v28 int8 = -55
                            _ = v28
                        case 1:
                            v7 = append(v7, 1)
                        }
                        switch -56 {
                        case 4294967295:
                            v29 := GenS6{F0: (-(v4)), F1: -17, F2: -9223372036854775808}
                            _ = v29
                        case 887:
                            var v30 uint64 = (v5 + v5)
                            _ = v30
                        default:
                            if len(v7) > 0 {
                                v7[0] = 832
                            }
                        }
                    }
                    for fi31 := 0; fi31 < 15; fi31++ {
                        var v32 uint8 = 42
                        v12 -= (v12 + 0.0)
                        _ = fi31
                        _ = v32
                    }
                }
                v33 := GenS6{F0: v4, F1: 42, F2: 32768}
                _ = v18
                _ = v33
            } else {
                v34 := GenS2{F0: -9223372036854775808, F1: 1, F2: v9}
                for fi35 := 0; fi35 < 7; fi35++ {
                    for fi36 := 0; fi36 < 8; fi36++ {
                        var v37 uint64 = 4294967295
                        _ = fi36
                        _ = v37
                    }
                    _ = fi35
                }
                for fi38 := 0; fi38 < 8; fi38++ {
                    fmt.Println(v5)
                    if false {
                        switch ("hello" + ("hello" + ((("" + ((v9 + v8) + v8)) + "") + ("" + v9)))) {
                        case "a":
                            for fi39 := 0; fi39 < 10; fi39++ {
                                var v40 int = 288
                                v8 = v9
                                _ = fi39
                                _ = v40
                            }
                        case "\n\t":
                            if ((true && (2147483647 > 42)) && (v5 <= (v5 ^ v5))) {
                                v3[("a" + (v9 + v9))] = ((!false) && (910 >= -2147483648))
                                _ = len(v7)
                                v9 = v9
                            }
                        }
                    }
                    _ = fi38
                }
                _ = v34
            }
        }
    }
    if true {
        fmt.Println(v6, v9, v9)
        v41 := GenS3{F0: v9}
        for fi42 := 0; fi42 < 14; fi42++ {
            if (!(!false)) {
                fmt.Println(v4)
                for fi43 := 0; fi43 < 9; fi43++ {
                    fmt.Println(fi43, v5, v12)
                    var v44 uint8 = 165
                    _ = fi43
                    _ = v44
                }
                for fi45 := 0; fi45 < 10; fi45++ {
                    _ = len(v7)
                    if (!false) {
                        v46 := GenS3{F0: ("\n\t" + "hello")}
                        v8 = "\n\t"
                        v47 := GenS7{F0: -2147483648}
                        _ = v46
                        _ = v47
                    } else {
                        var v48 int8 = -87
                        _ = v48
                    }
                    fmt.Println(fi42, v8)
                    _ = fi45
                }
            }
            for fi49 := 0; fi49 < 1; fi49++ {
                switch (fi42 & (fi42 * fi49)) {
                case 1:
                    var v50 int8 = 42
                    _ = v50
                case 2147483647:
                    var v51 uint32 = 2147483648
                    _ = v51
                case -2147483648:
                    v9 = v9
                default:
                    if ((v9 + v9) < "\n\t") {
                        v8 = "a"
                        var v52 int = fi49
                        v9 = "\n\t"
                        _ = v52
                    }
                }
                _ = fi49
            }
            _ = fi42
        }
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
