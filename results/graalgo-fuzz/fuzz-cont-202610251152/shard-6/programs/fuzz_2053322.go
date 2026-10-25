package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int16; F2 bool }

func main() {
    var v1 int32 = -530
    v2 := []uint32{}
    v3 := []string{"foo bar", ("\n\t" + "foo bar")}
    var v4 int32 = -853
    _ = len(v3)
    switch ("foo bar" + ("foo bar" + (("hello" + (("" + (("" + ((("" + "a") + "a") + ("foo bar" + "foo bar"))) + ("\n\t" + ("foo bar" + "a")))) + ("\n\t" + "hello"))) + "hello"))) {
    case "\n\t":
        switch -397 {
        case 127:
            if false {
                if false {
                    fmt.Println(v4, v4)
                } else {
                    if (4294967295 > 1) {
                        var v5 int16 = 1
                        var v6 int16 = v5
                        _ = v5
                        _ = v6
                    }
                }
                v1 = v1
                v1 = (^(v4))
            } else {
                var v7 bool = false
                _ = v7
            }
        default:
            if (2147483647 == (^(v4))) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v8 uint8 = 1
                _ = v8
            }
        }
    }
    var v9 bool = (false && true)
    v1 += v1
    v1 = (v1 & v1)
    if (4294967295 != 754) {
        if v9 {
            fmt.Println(v4, v1, v1)
            v4 += v4
        }
        if (-2.25 != (3.14159 * (((0.1 + float64(-2.25)) / (1e10 + 0.0)) / ((3.14159 / -2.25) + (-2.25 * 1.5))))) {
            if ((v1 > (v4 | 42)) || (-51 == -89)) {
                fmt.Println(v4, v4)
                var v10 bool = v9
                _ = v10
            } else {
                if (!((4294967295 != 2) && v9)) {
                    if (-928 <= 1000) {
                        var v11 int32 = (v4 - v1)
                        var v12 uint = 4294967295
                        if false {
                            var v13 int16 = 259
                            var v14 bool = false
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                    }
                } else {
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 int64 = -884
                        if (!(42 >= 4294967295)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            var v17 int32 = (v4 | 42)
                            _ = v17
                        }
                        _ = fi15
                        _ = v16
                    }
                }
                if v9 {
                    if v9 {
                        var v18 int8 = -128
                        _ = v18
                    } else {
                        v3 = append(v3, "\n\t")
                    }
                    for fi19 := 0; fi19 < 3; fi19++ {
                        var v20 int64 = 889
                        v4 = (v1 << ((0) & 63))
                        var v21 int64 = (v20 & v20)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                fmt.Println(v1)
            }
            if (170 == 193) {
                v9 = v9
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 uint8 = 255
                    v24 := GenS1{F0: 2147483648, F1: v9, F2: uint32(v23)}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int32 = v1
                    if v9 {
                        var v27 float32 = 1.5
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
            } else {
                if v9 {
                    v28 := GenS2{F0: "\n\t", F1: (^(v4)), F2: v4}
                    _ = v28
                }
                fmt.Println(v1)
                for fi29 := 0; fi29 < 13; fi29++ {
                    fmt.Println(v4, v1, v9)
                    v3 = append(v3, "hello")
                    v9 = false
                    _ = fi29
                }
            }
        }
        fmt.Println(v9)
    } else {
        for fi30 := 0; fi30 < 3; fi30++ {
            fmt.Println(v9, v4, v4)
            if true {
                v9 = (127 >= 255)
                v1 -= (v4 & -2147483648)
            }
            switch ("\n\t" + "a") {
            case "\n\t":
                var v31 uint16 = 65535
                _ = v31
            default:
                v1 = v1
            }
            _ = fi30
        }
    }
    v32 := GenS3{F0: ("hello" + ("a" + ""))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4, v9, v9)
    if ((true && (v9 || (v9 || v9))) || v9) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            v3 = append(v3, "foo bar")
            v1 = -148
            switch "a" {
            case "\n\t":
                for fi34 := 0; fi34 < 17; fi34++ {
                    fmt.Println(fi33, v1)
                    fmt.Println(v9)
                    v1 = v1
                    _ = fi34
                }
            default:
                for fi35 := 0; fi35 < 6; fi35++ {
                    v36 := GenS4{F0: (!v9), F1: 184, F2: v9}
                    var v37 string = "a"
                    var v38 int32 = (v1 << ((18446744073709551615) & 63))
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
            }
            _ = fi33
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v9, v9, v1)
    var v39 uint = 2
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v9
}
