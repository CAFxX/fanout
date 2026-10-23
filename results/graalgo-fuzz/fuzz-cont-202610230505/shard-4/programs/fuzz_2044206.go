package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS4 struct { F0 int64; F1 string; F2 int8 }

func main() {
    var v1 uint = 65535
    var v2 uint8 = 255
    var v3 int64 = -571
    var v4 uint32 = 13
    var v5 uint8 = (v2 % ((v2) | 1))
    v2 += 77
    v6 := []int{32767, int(v4), 4294967295}
    if false {
        v2 += (v2 * 127)
    } else {
        v4 -= (v4 + 406)
        var v7 float64 = 1.5
        _ = v7
    }
    fmt.Println(v5, v1, v3)
    _ = len(v6)
    _ = len(v6)
    var v8 uint32 = (v4 & 299)
    v8 += 761
    v4 -= v4
    if true {
        var v9 float64 = -2.25
        switch "foo bar" {
        case "foo bar":
            if (v4 > v4) {
                v9 = v9
                var v10 float32 = 1e10
                _ = v10
            } else {
                var v11 uint32 = v8
                var v12 float64 = ((((v9 - -2.25) / ((float64((0.0 / float32(v9))) + v9) / 1.5)) + (float64(1e10) + v9)) - 0.0)
                v13 := GenS1{F0: v1}
                _ = v11
                _ = v12
                _ = v13
            }
        case "hello":
            v14 := GenS2{F0: 1000}
            _ = v14
        default:
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS3{F0: v2, F1: 32767, F2: -355}
                if false {
                    if false {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if (!false) {
                            switch "a" {
                            case "a":
                                _ = len(v6)
                            case "foo bar":
                                v17 := GenS4{F0: v3, F1: "a", F2: -75}
                                _ = v17
                            case "":
                                switch "a" {
                                case "\n\t":
                                    v18 := GenS2{F0: -1}
                                    _ = v18
                                case "a":
                                    if len(v6) > 0 {
                                        v6[0] = (fi15 << ((18446744073709551615) & 63))
                                    }
                                }
                            }
                            v4 += v8
                            var v19 int32 = -2147483648
                            _ = v19
                        } else {
                            v4 = 401
                            for fi20 := 0; fi20 < 2; fi20++ {
                                var v21 string = "hello"
                                _ = fi20
                                _ = v21
                            }
                            fmt.Println(v3, v1, v5)
                        }
                    }
                    for fi22 := 0; fi22 < 6; fi22++ {
                        if (32768 == fi15) {
                            switch "foo bar" {
                            case "foo bar":
                                var v23 uint64 = 330
                                _ = v23
                            }
                            if (((v9 * ((v9 * v9) * (v9 + v9))) / v9) > v9) {
                                var v24 string = "foo bar"
                                _ = v24
                            }
                        } else {
                            v25 := GenS2{F0: 572}
                            if true {
                                v26 := GenS3{F0: (^(v2)), F1: 788, F2: 629}
                                v8 += v4
                                var v27 uint16 = uint16(v2)
                                _ = v26
                                _ = v27
                            } else {
                                var v28 uint32 = v8
                                _ = v28
                            }
                            _ = v25
                        }
                        _ = fi22
                    }
                    v2 += (v2 + 42)
                }
                _ = fi15
                _ = v16
            }
        }
        v29 := GenS3{F0: (^(v5)), F1: 0, F2: -50}
        _ = v29
        _ = v9
    } else {
        var v30 string = "\n\t"
        _ = v30
    }
    var v31 string = (("hello" + ("\n\t" + "\n\t")) + ("hello" + "hello"))
    var v32 int16 = -1
    for fi33 := 0; fi33 < 2; fi33++ {
        v5 = 12
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 int8 = -26
            switch v31 {
            case "":
                _ = len(v6)
            default:
                var v36 int64 = v3
                _ = v36
            }
            var v37 uint16 = 32767
            _ = fi34
            _ = v35
            _ = v37
        }
        for fi38 := 0; fi38 < 8; fi38++ {
            _ = len(v6)
            if (!(!false)) {
                v39 := GenS3{F0: v5, F1: 427, F2: v32}
                var v40 bool = false
                _ = v39
                _ = v40
            } else {
                v2 = v2
                fmt.Println(v3)
            }
            _ = fi38
        }
        _ = fi33
    }
    if len(v6) > 0 {
        v6[0] = 127
    }
    v6 = append(v6, -898)
    v32 += (v32 ^ v32)
    if true {
        var v41 float32 = float32(1e10)
        _ = v41
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        fmt.Println(v4, v8)
        var v43 uint32 = uint32(v1)
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
