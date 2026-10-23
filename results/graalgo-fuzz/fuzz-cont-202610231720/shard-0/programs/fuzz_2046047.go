package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int64, p1 int16, p2 uint32) int32 {
    v5 := []int{1, 2147483647, 127}
    var v6 uint16 = 42
    var v7 int = 4294967295
    switch (v7 << ((1) & 63)) {
    case 0:
        v7 -= v7
    default:
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int = (v7 * (fi8 ^ (v7 << ((255) & 63))))
            v10 := GenS1{F0: ("" + ""), F1: 255, F2: (v9 * v9)}
            fmt.Println(v7, fi8)
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 bool = false
    var v2 float64 = float64(1e10)
    var v3 uint8 = 128
    var v4 string = "\n\t"
    v11 := []int{1, 128, 2147483647, 2147483647}
    switch -660 {
    case 525:
        v12 := GenS2{F0: -111, F1: 1, F2: 1000}
        _ = v12
    }
    v13 := GenS3{F0: v3}
    _ = len(v11)
    v14 := GenS4{F0: 18446744073709551615, F1: (v3 >> ((42) & 63))}
    for fi15 := 0; fi15 < 17; fi15++ {
        if (1 <= 36) {
            var v16 int16 = -1
            _ = v16
        } else {
            if (42 <= 4294967295) {
                fmt.Println(v3, fi15, v1)
                v4 = (v4 + "\n\t")
            }
        }
        _ = len(v11)
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi15
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v1 = true
    v4 = "a"
    for fi17 := 0; fi17 < 8; fi17++ {
        v1 = ((1.5 * 0.0) != (1e10 * (0.1 - 3.14159)))
        _ = fi17
    }
    switch v4 {
    case "foo bar":
        v18 := GenS5{F0: -922}
        _ = v18
    case "hello":
        var v19 uint16 = 292
        _ = v19
    case "":
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                if len(v11) > 0 {
                    v11[0] = -9223372036854775808
                }
                switch (-(fi21)) {
                case 42:
                    v2 += (1.5 + ((v2 * ((v2 + (v2 + 1e10)) / v2)) * 1e10))
                }
                var v22 int32 = 440
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    default:
        switch 32767 {
        case -910:
            if v1 {
                fmt.Println(v2, v1, v1)
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(v3)
                    v11 = append(v11, 724)
                    var v24 uint16 = uint16(v3)
                    _ = fi23
                    _ = v24
                }
            }
        default:
            switch "\n\t" {
            case "\n\t":
                v25 := GenS6{F0: 255}
                _ = v25
            case "a":
                var v26 float64 = float64(-2.25)
                _ = v26
            case "":
                var v27 int8 = 1
                _ = v27
            }
        }
    }
    fmt.Println(v2, v2, v3)
    switch 32767 {
    case -269:
        var v28 int16 = 1
        _ = v28
    case 32768:
        var v29 int32 = 0
        _ = v29
    default:
        var v30 uint16 = 42
        _ = v30
    }
    v4 = ""
    if v1 {
        var v31 string = (v4 + v4)
        v2 -= (0.1 * v2)
        for fi32 := 0; fi32 < 11; fi32++ {
            if v1 {
                fmt.Println(v2, v2)
            } else {
                var v33 int32 = 1
                switch fi32 {
                case -976:
                    fmt.Println(v2, v4, v31)
                case -9223372036854775808:
                    if v1 {
                        for fi34 := 0; fi34 < 20; fi34++ {
                            var v35 string = (v31 + v4)
                            _ = fi34
                            _ = v35
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    } else {
                        var v36 int32 = (^(v33))
                        for fi37 := 0; fi37 < 19; fi37++ {
                            var v38 uint64 = 42
                            if len(v11) > 0 {
                                v11[0] = 0
                            }
                            _ = fi37
                            _ = v38
                        }
                        _ = v36
                    }
                }
                _ = v33
            }
            _ = fi32
        }
        _ = v31
    }
    v4 = v4
    for fi39 := 0; fi39 < 2; fi39++ {
        var v40 uint16 = 65535
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
