package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
func genFn0(p0 string, p1 uint16) uint64 {
    if (-1 != -2147483648) {
        v6 := GenS1{F0: 267, F1: 4294967295}
        if true {
            var v7 int64 = -516
            _ = v7
        } else {
            v8 := GenS2{F0: -667}
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(p1, fi9, p1)
            p0 = ("" + "a")
            _ = fi9
        }
        _ = v6
    }
    var v10 uint16 = 0
    fmt.Println(p1, v10, p1)
    _ = p0
    _ = p1
    _ = v10
    return 570
}
func genFn1(p0 uint, p1 uint, p2 int16) int8 {
    fmt.Println(p1)
    p0 += p1
    p1 += (p1 << ((777) & 63))
    for fi11 := 0; fi11 < 12; fi11++ {
        p0 += (p1 - (p1 + (^(p1))))
        _ = fi11
    }
    if (true && true) {
        p1 += (p0 * p0)
    } else {
        var v12 int = 706
        var v13 int8 = -74
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint8 = 0
    var v2 int8 = -53
    var v3 int32 = 1
    var v4 float64 = 3.14159
    v5 := []uint32{0}
    var v14 int = 9223372036854775807
    switch v14 {
    case 351:
        v4 -= ((float64(((3.14159 + float32(v4)) / (-2.25 / 3.14159))) * (v4 - (((float64(1e10) + float64(1e10)) / ((((v4 / v4) / (((((v4 * (v4 / float64((1.5 + 1e10)))) - v4) - v4) + v4) * 1.5)) - 3.14159) + v4)) - ((((v4 * v4) - 3.14159) - 0.1) + -2.25)))) * -2.25)
    }
    v1 += (v1 / (((v1 + 0)) | 1))
    v3 += (v3 + (^(v3)))
    if len(v5) > 0 {
        v5[0] = 0
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint8 = v1
        var v17 uint32 = 4294967295
        if (42 == -9223372036854775808) {
            for fi18 := 0; fi18 < 8; fi18++ {
                switch fi18 {
                case -769:
                    if ((!(v17 < 65535)) || (v17 < 378)) {
                        fmt.Println(v16, v16)
                        var v19 int8 = v2
                        var v20 bool = (v4 != 0.1)
                        _ = v19
                        _ = v20
                    }
                case 157:
                    fmt.Println(v3, fi15)
                case -1:
                    switch (("" + "\n\t") + "foo bar") {
                    case "a":
                        var v21 string = (("\n\t" + ("foo bar" + ("a" + "hello"))) + "a")
                        _ = v21
                    case "\n\t":
                        if false {
                            fmt.Println(v17, v16)
                            fmt.Println(v3, fi15, v2)
                            fmt.Println(v16)
                            v22 := GenS3{F0: 1000}
                            _ = len(v5)
                            v16 = (^(v16))
                            v23 := GenS2{F0: (fi18 | fi18)}
                            _ = v22
                            _ = v23
                        } else {
                            if true {
                                var v24 int = 32767
                                if (42 == 18446744073709551615) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                _ = v24
                            }
                        }
                    default:
                        var v25 int32 = (v3 & 0)
                        _ = v25
                    }
                }
                var v26 uint16 = 255
                var v27 int32 = (v3 ^ (v3 / (((v3 | 42)) | 1)))
                _ = fi18
                _ = v26
                _ = v27
            }
            _ = len(v5)
            fmt.Println(v4, v2, v14)
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v5 = append(v5, 220)
    for fi28 := 0; fi28 < 1; fi28++ {
        v2 += (v2 + v2)
        v29 := GenS1{F0: v3, F1: uint64(v1)}
        _ = fi28
        _ = v29
    }
    v3 += v3
    v30 := GenS3{F0: -1}
    switch "hello" {
    case "a":
        var v31 bool = (42 <= 42)
        _ = v31
    case "foo bar":
        switch ("\n\t" + "") {
        case "a":
            var v32 bool = true
            _ = v32
        default:
            var v33 int16 = -32768
            _ = v33
        }
    }
    v3 = v3
    v34 := GenS2{F0: 1}
    v1 -= (v1 / (((v1 >> ((2) & 63))) | 1))
    v35 := GenS1{F0: 600, F1: 188}
    fmt.Println(v4, v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
