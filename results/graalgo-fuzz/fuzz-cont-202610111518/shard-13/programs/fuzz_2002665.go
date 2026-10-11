package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 string; F1 int8; F2 string }
type GenS6 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0() int64 {
    v7 := GenS1{F0: 599}
    var v8 int8 = 100
    _ = v7
    _ = v8
    return 852
}

func main() {
    var v1 int8 = 0
    var v2 int32 = -103
    var v3 uint16 = 2
    var v4 int64 = -390
    v5 := make(map[int]uint32)
    v6 := []int{2147483647}
    var v9 uint16 = v3
    fmt.Println(v4, v4, v9)
    if (!(381 != 1)) {
        v2 = (v2 & 375)
        v6 = append(v6, 933)
        if true {
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(v2)
                if (618 <= 42) {
                    var v11 uint16 = (v3 >> ((255) & 63))
                    var v12 int16 = -1
                    var v13 int8 = (v1 / (((v1 * 127)) | 1))
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        _ = len(v6)
                        if true {
                            v15 := GenS2{F0: 42, F1: 127, F2: fi14}
                            v3 = v3
                            fmt.Println(fi14, v1, fi14)
                            switch ("hello" + "") {
                            case "hello":
                                v3 += (^(v9))
                            case "foo bar":
                                var v16 int16 = -217
                                _ = v16
                            case "\n\t":
                                var v17 uint64 = 931
                                _ = v17
                            default:
                                switch "a" {
                                case "foo bar":
                                    var v18 uint = 123
                                    _ = v18
                                }
                            }
                            v2 = (v2 % (((v2 % (((v2 / (((v2 & -727)) | 1))) | 1))) | 1))
                            _ = v15
                        } else {
                            v9 = 65535
                            if ((v3 % ((v9) | 1)) > (^(v3))) {
                                v3 = v3
                            } else {
                                v2 = (^(v2))
                                _ = len(v5)
                            }
                            if (fi10 == 127) {
                                v1 += -1
                                var v19 uint8 = 128
                                v20 := GenS3{F0: genFn0(), F1: 1000, F2: 32767}
                                _ = v19
                                _ = v20
                            }
                        }
                        _ = fi14
                    }
                    if (18446744073709551615 >= 4294967295) {
                        var v21 int64 = (v4 % (((^(v4))) | 1))
                        delete(v5, fi10)
                        _ = v21
                    } else {
                        fmt.Println(v9, fi10)
                        if true {
                            var v22 int = fi10
                            var v23 uint = 194
                            if len(v6) > 0 {
                                v6[0] = (fi10 >> ((v23) & 63))
                            }
                            _ = v22
                            _ = v23
                        } else {
                            var v24 uint16 = v9
                            v25 := GenS4{F0: 42, F1: v1}
                            fmt.Println(v24)
                            _ = v24
                            _ = v25
                        }
                    }
                }
                _ = fi10
            }
        } else {
            var v26 string = ("" + ("foo bar" + ("foo bar" + "\n\t")))
            fmt.Println(v1)
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 bool = ((((0 > 255) || ((v4 % (((v4 - 42)) | 1)) != (v4 | (v4 + v4)))) || (!false)) || (true && true))
        if v28 {
            delete(v5, fi27)
            fmt.Println(v9, v2)
        }
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 bool = v28
            v4 = v4
            fmt.Println(v9, v28, v1)
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v2, v3)
    v31 := GenS5{F0: ("foo bar" + ""), F1: 100, F2: "\n\t"}
    v3 += (v9 - (v9 - (v3 % ((42) | 1))))
    if (!true) {
        for fi32 := 0; fi32 < 2; fi32++ {
            if (1.5 >= ((((float32(0.1) / 1.0) - (-2.25 + (1e10 + float32(0.0)))) / 1e10) / 3.14159)) {
                for fi33 := 0; fi33 < 7; fi33++ {
                    v3 += v9
                    var v34 float32 = -2.25
                    _ = len(v6)
                    _ = fi33
                    _ = v34
                }
                var v35 uint8 = 206
                _ = v35
            } else {
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
                var v36 uint = 476
                var v37 int = (fi32 / ((fi32) | 1))
                _ = v36
                _ = v37
            }
            _ = fi32
        }
        v38 := GenS6{F0: 2, F1: v1, F2: 4294967295}
        if false {
            for fi39 := 0; fi39 < 9; fi39++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if (false || false) {
                    v9 -= (v3 - (v3 << ((124) & 63)))
                }
                _ = fi39
            }
            delete(v5, 127)
        }
        _ = v38
    }
    delete(v5, 0)
    var v40 int64 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
