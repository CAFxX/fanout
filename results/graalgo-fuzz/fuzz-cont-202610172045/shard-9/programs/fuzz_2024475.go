package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int16; F2 int }
type GenS5 struct { F0 int16; F1 int16 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]int)
    var v3 int32 = 42
    var v4 uint64 = 18446744073709551615
    var v5 int64 = 9223372036854775807
    fmt.Println(v3, v3)
    if false {
        v4 -= (v4 - v4)
        switch 552 {
        case 2147483647:
            v3 -= (v3 + (v3 & v3))
        default:
            v3 -= (v3 + 55)
        }
    } else {
        var v6 bool = true
        if v6 {
            v7 := GenS1{F0: (v4 - v4), F1: 255, F2: (v5 / ((42) | 1))}
            v8 := GenS1{F0: v4, F1: 1, F2: -9223372036854775808}
            v1 = ((((1.5 / 1e10) + 1e10) + v1) - (1e10 + (3.14159 / v1)))
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                if v6 {
                    v1 = -2.25
                } else {
                    var v10 uint = 255
                    var v11 int32 = (v3 * 1)
                    _ = v10
                    _ = v11
                }
                switch (fi9 | fi9) {
                case 59:
                    var v12 string = (("hello" + "foo bar") + ("hello" + ""))
                    _ = v12
                case -9223372036854775808:
                    v13 := GenS2{F0: 42}
                    _ = v13
                case 128:
                    var v14 int64 = v5
                    _ = v14
                default:
                    v4 += v4
                }
                switch -713 {
                case 9223372036854775807:
                    if false {
                        delete(v2, 1)
                        var v15 bool = (v3 != v3)
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v2[65535] = fi9
                            switch "\n\t" {
                            case "\n\t":
                                var v17 string = ""
                                _ = v17
                            case "":
                                v18 := []int{}
                                _ = v18
                            case "a":
                                if ((v5 & v5) >= 248) {
                                    if ((v4 >> ((2) & 63)) != 0) {
                                        v19 := GenS1{F0: (^(v4)), F1: 499, F2: -1}
                                        v3 = (v3 - (v3 << ((2) & 63)))
                                        v20 := []int{}
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                        _ = v19
                                        _ = v20
                                    }
                                }
                            default:
                                fmt.Println(v5)
                                if true {
                                    v6 = v6
                                }
                            }
                            v21 := GenS3{F0: 0}
                            v22 := GenS1{F0: v4, F1: 0, F2: 1}
                            v23 := []int{(fi16 | fi9), (-(fi9)), -1}
                            _ = fi16
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        var v24 int32 = v3
                        if (-32768 <= -1) {
                            fmt.Println(v4, v3, v6)
                            fmt.Println(v24, v5, v1)
                        }
                        _ = v24
                    }
                case -847:
                    fmt.Println(v5)
                default:
                    v1 += 3.14159
                }
                _ = fi9
            }
            fmt.Println(v6)
            fmt.Println(v1)
        }
        _ = v6
    }
    v2[4294967295] = -1
    v25 := GenS4{F0: (9223372036854775807 > -952), F1: 1000, F2: 32768}
    v26 := GenS1{F0: (v4 % (((v4 / (((v4 ^ v4)) | 1))) | 1)), F1: 470, F2: v5}
    var v27 int = 127
    v4 = v4
    v28 := GenS5{F0: -763, F1: 1}
    var v29 int = (v27 >> ((65535) & 63))
    fmt.Println(v27, v3, v29)
    var v30 int8 = -24
    var v31 float32 = 1e10
    v1 = v1
    v5 = v5
    switch (((("hello" + ("\n\t" + "a")) + "hello") + "\n\t") + "a") {
    case "":
        if (false || false) {
            var v32 int16 = -491
            v1 = v1
            v33 := v2[19]
            _ = v32
            _ = v33
        }
    case "hello":
        var v34 int8 = (v30 << ((9) & 63))
        _ = v34
    default:
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 int = v27
            v5 -= v5
            _ = fi35
            _ = v36
        }
    }
    var v37 float64 = float64(v31)
    v38 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
