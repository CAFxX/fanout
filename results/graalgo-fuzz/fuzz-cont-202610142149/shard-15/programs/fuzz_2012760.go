package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint16; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint64 }
func genFn0(p0 bool, p1 int64) uint {
    v6 := []int{-2147483648, 0, -9223372036854775808, 42}
    p0 = (false && p0)
    var v7 int8 = 104
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 uint8 = 186
    var v4 bool = (false && ((!(32767 == 0)) || (-2.25 == 1.5)))
    v5 := []string{}
    if (false || v4) {
        v8 := GenS1{F0: 32767}
        if true {
            for fi9 := 0; fi9 < 16; fi9++ {
                v5 = append(v5, "")
                _ = fi9
            }
            fmt.Println(v4, v4)
        } else {
            v10 := GenS1{F0: -563}
            _ = v10
        }
        _ = v8
    } else {
        var v11 int8 = 42
        var v12 int8 = (v11 << ((uint(v3)) & 63))
        _ = v11
        _ = v12
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        if (1 == 18446744073709551615) {
            v4 = ((1.5 * 1.5) < ((0.1 / 1e10) / 0.1))
        }
        switch (("foo bar" + "hello") + "foo bar") {
        case "":
            switch "\n\t" {
            case "\n\t":
                var v14 uint16 = 674
                _ = v14
            case "a":
                v4 = v4
            }
        case "a":
            var v15 int32 = -1
            _ = v15
        case "\n\t":
            v16 := GenS2{F0: -1, F1: 32767, F2: -412}
            _ = v16
        default:
            v17 := GenS1{F0: fi13}
            _ = v17
        }
        _ = fi13
    }
    if (true && ((1e10 / 1.0) == 0.1)) {
        fmt.Println(v4)
        var v18 uint16 = 1
        _ = v18
    }
    switch "a" {
    case "\n\t":
        var v19 int64 = 33
        _ = v19
    case "a":
        if true {
            var v20 uint64 = 42
            v21 := v1[32768]
            _ = v20
            _ = v21
        }
    case "hello":
        v22 := v1[32767]
        _ = v22
    }
    switch "hello" {
    case "":
        v3 -= 127
    case "hello":
        v4 = v4
    case "\n\t":
        delete(v2, 0)
    default:
        v23 := GenS1{F0: -44}
        _ = v23
    }
    _ = len(v1)
    if v4 {
        fmt.Println(v3, v3)
        v4 = (128 < -674)
        v2[257] = -1
    } else {
        switch 0 {
        case 0:
            var v24 int16 = 1
            _ = v24
        case 32768:
            var v25 bool = false
            _ = v25
        case 2147483647:
            if v4 {
                fmt.Println(v4, v4)
                if true {
                    if false {
                        fmt.Println(v4)
                    } else {
                        v3 += v3
                    }
                    fmt.Println(v3, v4)
                } else {
                    if ((18446744073709551615 <= 852) && v4) {
                        if (42 >= -1) {
                            var v26 uint16 = 551
                            v26 = v26
                            _ = v26
                        }
                        v5 = append(v5, "")
                    }
                    if len(v5) > 0 {
                        v5[0] = ("\n\t" + ("foo bar" + "\n\t"))
                    }
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                }
            } else {
                v27 := GenS1{F0: 154}
                _ = v27
            }
        default:
            for fi28 := 0; fi28 < 7; fi28++ {
                switch "\n\t" {
                case "hello":
                    _ = len(v1)
                case "foo bar":
                    v3 += v3
                default:
                    v29 := GenS3{F0: -2147483648}
                    _ = v29
                }
                if (-162 > -1) {
                    switch "a" {
                    case "":
                        for fi30 := 0; fi30 < 15; fi30++ {
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(fi28, fi30)
                            case "":
                                v5 = append(v5, "\n\t")
                            case "foo bar":
                                v3 = (v3 / (((v3 + 1)) | 1))
                            default:
                                var v31 int64 = -9223372036854775808
                                _ = v31
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi30
                        }
                    case "foo bar":
                        fmt.Println(v3, v3, v3)
                    case "a":
                        var v32 uint64 = 42
                        _ = v32
                    default:
                        v33 := GenS4{F0: 1, F1: -2147483648, F2: 0}
                        _ = v33
                    }
                } else {
                    v34 := GenS3{F0: 2147483647}
                    _ = v34
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    if (((v3 % ((21) | 1)) > (v3 << ((2) & 63))) || (((v3 ^ v3) >= (v3 & (v3 | v3))) || (!true))) {
                        v36 := GenS2{F0: 1, F1: 1, F2: -9223372036854775808}
                        _ = v36
                    }
                    _ = fi35
                }
                _ = fi28
            }
        }
        var v37 int8 = 127
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
