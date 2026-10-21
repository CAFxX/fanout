package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 string; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("hello" + ""), ("\n\t" + ("" + (("" + (("a" + "hello") + ("\n\t" + (("foo bar" + "a") + "\n\t")))) + ("foo bar" + (("a" + (("" + "") + "foo bar")) + ("a" + ("a" + ("a" + "foo bar"))))))))}
    v3 := make(map[uint32]uint32)
    v4 := []string{"a"}
    v5 := make(map[int]bool)
    if true {
        var v6 float32 = 1e10
        if (42 < 2) {
            var v7 bool = true
            v7 = v7
            _ = v7
        }
        _ = v6
    }
    var v8 int = 1
    var v9 float32 = 0.1
    var v10 string = "hello"
    var v11 int64 = 9223372036854775807
    if (!true) {
        if (v11 <= (v11 / (((v11 | -9223372036854775808)) | 1))) {
            if (true || (((float32(float64(v9)) * float32(0.0)) > v9) || (42 == 295))) {
                if true {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 uint64 = 351
                        _ = fi12
                        _ = v13
                    }
                    if false {
                        switch 42 {
                        case 397:
                            v9 += 3.14159
                        default:
                            var v14 float64 = -2.25
                            _ = v14
                        }
                    }
                }
                var v15 uint8 = 255
                _ = v15
            }
            v8 = 128
        } else {
            for fi16 := 0; fi16 < 8; fi16++ {
                _ = len(v3)
                _ = fi16
            }
            var v17 int8 = 127
            for fi18 := 0; fi18 < 4; fi18++ {
                if false {
                    v9 = (1e10 / float32(3.14159))
                } else {
                    var v19 int64 = 42
                    var v20 uint = 82
                    _ = v19
                    _ = v20
                }
                var v21 int8 = v17
                if (!(862 > 1)) {
                    v22 := GenS1{F0: -2147483648}
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if (false && true) {
                            if (!false) {
                                v21 = 127
                                if true {
                                    v24 := GenS2{F0: 1000, F1: 493}
                                    v21 += (v21 - (v21 * v21))
                                    _ = v24
                                } else {
                                    v17 = (v17 << ((0) & 63))
                                }
                            } else {
                                fmt.Println(v9, v8, v10)
                                v25 := GenS2{F0: -737, F1: 0}
                                _ = v25
                            }
                        }
                        _ = len(v2)
                        v26 := GenS1{F0: -2147483648}
                        _ = fi23
                        _ = v26
                    }
                    _ = v22
                }
                _ = fi18
                _ = v21
            }
            _ = v17
        }
    }
    fmt.Println(v11, v10)
    var v27 uint = 42
    switch v10 {
    case "":
        var v28 string = (((v10 + v10) + v10) + ((v10 + (v10 + ("hello" + v10))) + (v10 + v10)))
        _ = v28
    case "hello":
        var v29 uint64 = 462
        _ = v29
    case "foo bar":
        switch (v8 / (((^(v8))) | 1)) {
        case 711:
            v8 = 9223372036854775807
        case 606:
            fmt.Println(v11, v27)
        case 128:
            if true {
                v11 += v11
                v27 = 593
            } else {
                for fi30 := 0; fi30 < 19; fi30++ {
                    if (false && ("\n\t" == (v10 + v10))) {
                        v31 := GenS3{F0: 1}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v32 int16 = 1000
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS4{F0: (v11 * (-(v11)))}
                    var v34 uint64 = 2
                    _ = fi30
                    _ = v33
                    _ = v34
                }
                switch "a" {
                case "":
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v5[fi35] = false
                        _ = fi35
                    }
                case "\n\t":
                    if false {
                        fmt.Println(v10, v9, v8)
                        v1[4294967295] = (v8 >> ((1) & 63))
                        v36 := v5[v8]
                        _ = v36
                    } else {
                        v37 := GenS5{F0: (v27 & (v27 | (v27 / ((18446744073709551615) | 1)))), F1: "a", F2: 65535}
                        if false {
                            var v38 int8 = -12
                            v39 := GenS6{F0: 128, F1: v8, F2: 0}
                            v38 = -128
                            _ = v38
                            _ = v39
                        } else {
                            fmt.Println(v8, v8, v9)
                            v10 = v10
                            var v40 int64 = (v11 | (v11 + (v11 * v11)))
                            _ = v40
                        }
                        _ = v37
                    }
                default:
                    v8 -= (v8 % (((v8 % (((v8 / ((v8) | 1))) | 1))) | 1))
                }
            }
        }
    }
    v41 := GenS5{F0: v27, F1: "\n\t", F2: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v8
    _ = v9
}
