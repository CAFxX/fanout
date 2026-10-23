package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 bool }

func main() {
    var v1 int = 42
    v2 := []int{v1, 272, v1, v1}
    var v3 uint8 = 127
    if (!true) {
        if false {
            v4 := GenS1{F0: 1, F1: 919, F2: 0}
            _ = v4
        }
        fmt.Println(v1, v3)
        var v5 uint = 4294967295
        _ = v5
    } else {
        v6 := GenS2{F0: 784}
        var v7 bool = (2 <= 4294967295)
        _ = v6
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            if true {
                for fi9 := 0; fi9 < 14; fi9++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v1 += 4294967295
                    _ = fi9
                }
                var v10 bool = (460 > 662)
                _ = v10
            } else {
                v3 -= 127
                var v11 int16 = 651
                _ = v11
            }
            var v12 uint16 = 32767
            v1 += fi8
            _ = fi8
            _ = v12
        }
        var v13 float32 = (float32(0.1) / 3.14159)
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint64 = 1
            var v16 int = (^(fi14))
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        v3 = v3
    }
    switch ("hello" + ("" + ((("foo bar" + "") + ((("\n\t" + "a") + "foo bar") + "foo bar")) + "a"))) {
    case "":
        v17 := GenS2{F0: 18446744073709551615}
        _ = v17
    case "foo bar":
        var v18 float32 = -2.25
        _ = v18
    default:
        var v19 int32 = -811
        _ = v19
    }
    var v20 float64 = 1e10
    if (0 != 1) {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 int = v1
            if false {
                if true {
                    if (478 > 42) {
                        for fi23 := 0; fi23 < 12; fi23++ {
                            switch (fi23 - v1) {
                            case -779:
                                if true {
                                    v20 += 1e10
                                    v24 := GenS3{F0: -1}
                                    switch ("" + "a") {
                                    case "foo bar":
                                        var v25 uint8 = (v3 / (((v3 & (v3 >> ((1) & 63)))) | 1))
                                        _ = v25
                                    default:
                                        if false {
                                            v1 = 127
                                        } else {
                                            fmt.Println(fi23, v22, v3)
                                            v22 = (v22 / ((0) | 1))
                                        }
                                    }
                                    _ = v24
                                }
                                var v26 float64 = v20
                                switch -2147483648 {
                                case -1:
                                    fmt.Println(fi21, v26, v1)
                                    if len(v2) > 0 {
                                        v2[0] = (fi21 & (v22 << ((18446744073709551615) & 63)))
                                    }
                                default:
                                    v27 := GenS3{F0: -735}
                                    _ = v27
                                }
                                _ = v26
                            default:
                                var v28 int8 = -128
                                _ = v28
                            }
                            v2 = append(v2, (-(fi23)))
                            switch "a" {
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = fi23
                        }
                        var v29 uint8 = (v3 * 2)
                        _ = v29
                    } else {
                        for fi30 := 0; fi30 < 3; fi30++ {
                            fmt.Println(fi30, fi21, v22)
                            switch ("a" + "hello") {
                            case "foo bar":
                                v31 := GenS4{F0: 64, F1: (("" + ("" + "")) + "a")}
                                _ = v31
                            }
                            var v32 int = v1
                            _ = fi30
                            _ = v32
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                } else {
                    switch "hello" {
                    case "a":
                        v3 += (v3 & v3)
                    case "foo bar":
                        v33 := GenS5{F0: ("a" + "foo bar")}
                        _ = v33
                    }
                }
                fmt.Println(v3, v22)
            }
            if len(v2) > 0 {
                v2[0] = v22
            }
            _ = fi21
            _ = v22
        }
        if false {
            switch (v1 >> ((42) & 63)) {
            case 1:
                if false {
                    var v34 uint = 21
                    v2 = append(v2, (v1 * 714))
                    _ = v34
                } else {
                    var v35 float64 = float64(0.0)
                    v3 -= (v3 >> ((4294967295) & 63))
                    switch v1 {
                    case 2147483647:
                        fmt.Println(v3, v1, v3)
                    case 144:
                        if false {
                            var v36 uint16 = 2
                            v3 = 2
                            v35 = 0.0
                            _ = v36
                        } else {
                            _ = len(v2)
                        }
                    default:
                        fmt.Println(v20, v35)
                    }
                    _ = v35
                }
            }
            fmt.Println(v20, v3)
        } else {
            fmt.Println(v20)
        }
        if (32767 > 42) {
            fmt.Println(v20)
        }
    }
    if (!(65535 == 121)) {
        v37 := GenS2{F0: 0}
        var v38 float64 = (-2.25 * v20)
        if true {
            fmt.Println(v3, v38, v38)
        } else {
            v39 := GenS6{F0: 349, F1: 2}
            _ = v39
        }
        _ = v37
        _ = v38
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v40 := GenS7{F0: 128, F1: 1}
    v41 := GenS8{F0: true}
    v42 := GenS7{F0: 0, F1: 122}
    if false {
        v2 = append(v2, -9223372036854775808)
        if false {
            if (4294967295 != 2147483648) {
                var v43 uint32 = 2147483648
                _ = v43
            }
            switch ("\n\t" + "hello") {
            case "hello":
                v1 = 9223372036854775807
            default:
                v44 := GenS5{F0: "\n\t"}
                _ = v44
            }
        }
    }
    fmt.Println(v3, v20)
    for fi45 := 0; fi45 < 19; fi45++ {
        var v46 uint32 = 65535
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v40
    _ = v41
    _ = v42
}
