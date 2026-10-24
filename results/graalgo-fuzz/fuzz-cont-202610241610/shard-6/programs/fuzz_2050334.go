package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 uint; F1 string; F2 int32 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int }

func main() {
    var v1 int16 = -110
    v2 := []string{}
    var v3 float64 = -2.25
    var v4 uint32 = 2147483648
    v5 := GenS1{F0: 4294967295, F1: 35}
    if false {
        v4 += (v4 % (((v4 & v4)) | 1))
        if true {
            switch "a" {
            case "\n\t":
                var v6 int32 = 1
                _ = v6
            case "a":
                var v7 uint = 4294967295
                _ = v7
            }
            v3 -= (v3 - (v3 / (v3 / ((((0.0 * v3) / 1.5) * v3) - v3))))
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                v1 = 136
                _ = fi8
            }
            var v9 int = 1
            var v10 int32 = 998
            _ = v10
            _ = v9
        }
        var v11 bool = false
        _ = v11
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        v4 += v4
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -2147483648:
        fmt.Println(v3, v3)
    default:
        v12 := GenS1{F0: 1, F1: 2}
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int16 = 42
        v1 = (v1 - 1000)
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS2{F0: 65535, F1: (v4 ^ (v4 & (v4 >> ((1) & 63)))), F2: 42}
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v1)
    var v17 float32 = 1.5
    fmt.Println(v17)
    v18 := GenS2{F0: 32767, F1: v4, F2: 255}
    fmt.Println(v1)
    fmt.Println(v17, v3, v4)
    for fi19 := 0; fi19 < 9; fi19++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) > 0 {
            v2[0] = ""
        }
        if true {
            v17 = (3.14159 + 1.5)
            var v20 string = "\n\t"
            v1 -= (v1 | 0)
            _ = v20
        } else {
            var v21 int64 = 1
            for fi22 := 0; fi22 < 14; fi22++ {
                v1 = (^(v1))
                var v23 uint32 = 42
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        _ = fi19
    }
    if false {
        for fi24 := 0; fi24 < 20; fi24++ {
            switch "hello" {
            case "":
                v25 := GenS3{F0: "\n\t", F1: 0, F2: 0}
                _ = v25
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            default:
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                v3 = (0.1 / (float64(0.1) + 3.14159))
                v27 := GenS4{F0: 394, F1: (("a" + ("foo bar" + "")) + "a"), F2: -1}
                var v28 int8 = 81
                switch "foo bar" {
                case "a":
                    v17 = v17
                case "hello":
                    if ((((v17 * ((0.0 + 0.0) * (v17 * ((0.0 + v17) * 3.14159)))) / float32(v3)) + 1e10) == v17) {
                        v28 += (v28 << ((304) & 63))
                        var v29 int64 = -9223372036854775808
                        var v30 uint64 = 2
                        switch "" {
                        case "hello":
                            v28 = 1
                        }
                        if false {
                            v31 := GenS2{F0: 255, F1: 2147483648, F2: 255}
                            switch (fi26 << ((65535) & 63)) {
                            case -689:
                                var v32 uint = 18446744073709551615
                                _ = v32
                            default:
                                v28 = -128
                            }
                            v1 -= int16(fi26)
                            _ = v31
                        } else {
                            v4 = v4
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                        }
                        _ = v29
                        _ = v30
                    } else {
                        if true {
                            v4 += v4
                            var v33 uint16 = 32767
                            _ = v33
                        } else {
                            v3 += v3
                            v4 -= v4
                            v34 := GenS5{F0: (-(v1)), F1: v1}
                            _ = v34
                        }
                    }
                case "\n\t":
                    v35 := GenS6{F0: (v28 & (v28 & 1))}
                    v3 = 0.0
                    v3 += v3
                    _ = v35
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi24
        }
    }
    var v36 float64 = (v3 / v3)
    for fi37 := 0; fi37 < 15; fi37++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi37
    }
    v4 -= (v4 | v4)
    if true {
        var v38 uint16 = 42
        _ = v38
    } else {
        switch 806 {
        case 4294967295:
            if (true && (float32(((v36 - (v3 / v36)) / 1e10)) < 0.0)) {
                for fi39 := 0; fi39 < 16; fi39++ {
                    if true {
                        fmt.Println(fi39)
                    }
                    if false {
                        for fi40 := 0; fi40 < 2; fi40++ {
                            var v41 uint16 = 1
                            v42 := GenS7{F0: (v4 % (((^(v4))) | 1)), F1: 524}
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    } else {
                        _ = len(v2)
                    }
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + ("\n\t" + ("" + "")))
                    }
                    _ = fi39
                }
                v17 += (v17 * (v17 + (float32(1e10) * 0.0)))
                var v43 uint16 = 42
                _ = v43
            }
        }
    }
    var v44 uint = 0
    v3 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
    _ = v5
}
