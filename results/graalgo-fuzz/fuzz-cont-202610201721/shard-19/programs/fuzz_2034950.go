package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 string; F2 uint32 }

func main() {
    var v1 bool = true
    var v2 float32 = 1e10
    var v3 string = "\n\t"
    var v4 int16 = -1
    var v5 uint64 = 1
    v6 := []uint32{945}
    var v7 float64 = 0.1
    switch v3 {
    case "hello":
        switch -9223372036854775808 {
        case 9223372036854775807:
            if (0 < -159) {
                if v1 {
                    v6 = append(v6, 275)
                    fmt.Println(v4, v7)
                }
            } else {
                for fi8 := 0; fi8 < 12; fi8++ {
                    fmt.Println(v7, v7)
                    v9 := GenS1{F0: v3, F1: 52}
                    _ = fi8
                    _ = v9
                }
            }
        case 128:
            v3 = ""
        default:
            switch v3 {
            case "foo bar":
                v6 = append(v6, 730)
            case "hello":
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 float32 = float32(v7)
                    if v1 {
                        v12 := GenS2{F0: 107, F1: -331, F2: 797}
                        v3 = v3
                        _ = v12
                    } else {
                        fmt.Println(fi10, v4, v7)
                    }
                    _ = fi10
                    _ = v11
                }
            case "a":
                v4 = v4
            }
        }
    case "":
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v1 = v1
        v14 := GenS3{F0: (^(v5))}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: -101, F1: 42, F2: 4294967295}
    if (!(52 != -103)) {
        if (42 != -128) {
            fmt.Println(v1)
            fmt.Println(v5)
        }
        var v16 uint64 = v5
        var v17 int = 32768
        _ = v16
        _ = v17
    } else {
        v3 = ("foo bar" + (v3 + ""))
        if v1 {
            var v18 int16 = (v4 - v4)
            _ = v18
        } else {
            if (2147483647 < 2147483648) {
                v19 := GenS4{F0: -119}
                switch 42 {
                case -9223372036854775808:
                    if v1 {
                        for fi20 := 0; fi20 < 10; fi20++ {
                            for fi21 := 0; fi21 < 14; fi21++ {
                                v1 = (float64(v2) >= ((((v7 * (v7 - ((float64(v2) / (float64(((float32(1.5) * v2) * (v2 - v2))) + v7)) * float64(v2)))) - v7) - ((v7 + v7) + v7)) + 1e10))
                                _ = fi21
                            }
                            var v22 uint = 2
                            _ = len(v6)
                            _ = fi20
                            _ = v22
                        }
                    }
                default:
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v3 = v3
                        fmt.Println(v5, v5)
                        var v24 string = (v3 + ((v3 + "foo bar") + "hello"))
                        _ = fi23
                        _ = v24
                    }
                }
                _ = v19
            } else {
                var v25 uint64 = v5
                v26 := GenS5{F0: 321, F1: (v3 + ("foo bar" + "\n\t")), F2: 718}
                _ = v25
                _ = v26
            }
        }
        v3 = v3
    }
    switch ((v3 + (v3 + "a")) + "a") {
    case "foo bar":
        v1 = (!(v5 <= 945))
    case "a":
        if v1 {
            var v27 uint8 = 1
            v5 -= v5
            _ = v27
        } else {
            v28 := GenS5{F0: 1, F1: "a", F2: 2147483647}
            v7 = v7
            fmt.Println(v7, v5, v2)
            _ = v28
        }
    }
    var v29 int64 = 949
    v1 = true
    v30 := GenS2{F0: 127, F1: -1, F2: 2}
    var v31 int16 = 0
    var v32 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
