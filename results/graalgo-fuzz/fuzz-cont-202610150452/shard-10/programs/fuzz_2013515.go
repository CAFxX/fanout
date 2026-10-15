package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 int64 }

func main() {
    var v1 uint32 = 0
    v2 := []int{-641, 32768}
    v3 := []int{-913, 127}
    v4 := make(map[uint32]bool)
    var v5 uint8 = 0
    v6 := []int{-348, -401, 32768}
    var v7 int = 32767
    switch (v7 % ((v7) | 1)) {
    case -2147483648:
        v1 += (^(v1))
    case 2147483647:
        v8 := v4[4294967295]
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 0, F2: 2147483647}
    v10 := GenS1{F0: 1, F1: 729, F2: 42}
    v4[(v1 | v1)] = true
    switch v7 {
    case 0:
        _ = len(v4)
    case 2147483647:
        if (uint(v1) == 1) {
            v5 -= (^(v5))
            if (0 < -1) {
                if false {
                    v5 += (^(v5))
                }
                var v11 int16 = -32768
                switch 32767 {
                case -615:
                    v7 = v7
                case 0:
                    var v12 float64 = -2.25
                    _ = v12
                case -9223372036854775808:
                    switch v7 {
                    case 2147483647:
                        var v13 float64 = 3.14159
                        _ = v13
                    default:
                        v11 = -876
                    }
                default:
                    var v14 string = "foo bar"
                    _ = v14
                }
                _ = v11
            } else {
                switch 757 {
                case 32767:
                    v1 -= (v1 + 42)
                }
            }
        }
    case 32767:
        for fi15 := 0; fi15 < 4; fi15++ {
            switch v7 {
            case 42:
                var v16 uint16 = 255
                _ = v16
            case 128:
                var v17 uint16 = 274
                _ = v17
            }
            if true {
                if (!(!true)) {
                    v2 = append(v2, fi15)
                    var v18 uint64 = 18446744073709551615
                    var v19 float64 = -2.25
                    _ = v18
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        if (1.5 < (0.1 - -2.25)) {
                            if false {
                                var v21 int16 = 1000
                                var v22 int = (fi15 | fi20)
                                _ = v21
                                _ = v22
                            }
                        }
                        fmt.Println(fi20, v1, v5)
                        _ = fi20
                    }
                    if ("a" <= "\n\t") {
                        if (169 > 2) {
                            v6 = append(v6, int(v1))
                            v7 += (v7 / ((fi15) | 1))
                            switch (("\n\t" + ("" + "foo bar")) + "\n\t") {
                            case "hello":
                                switch "\n\t" {
                                case "hello":
                                    v3 = append(v3, fi15)
                                case "a":
                                    for fi23 := 0; fi23 < 9; fi23++ {
                                        v24 := GenS2{F0: 69}
                                        _ = fi23
                                        _ = v24
                                    }
                                case "":
                                    fmt.Println(v1, v7)
                                }
                            case "foo bar":
                                var v25 int8 = 0
                                _ = v25
                            }
                        }
                        v26 := GenS3{F0: (v5 | (v5 ^ (v5 ^ v5))), F1: 882, F2: 682}
                        for fi27 := 0; fi27 < 2; fi27++ {
                            fmt.Println(fi15)
                            _ = fi27
                        }
                        _ = v26
                    }
                }
                v28 := GenS1{F0: -1, F1: 42, F2: 42}
                var v29 float32 = 1e10
                _ = v28
                _ = v29
            }
            v30 := GenS2{F0: 0}
            _ = fi15
            _ = v30
        }
    }
    if len(v2) > 0 {
        v2[0] = (v7 << ((713) & 63))
    }
    v1 += (v1 * (^(v1)))
    fmt.Println(v5, v5, v7)
    fmt.Println(v5)
    v31 := GenS4{F0: 127, F1: true, F2: 785}
    var v32 uint16 = 40
    for fi33 := 0; fi33 < 3; fi33++ {
        switch (v7 & (fi33 & 137)) {
        case 4294967295:
            fmt.Println(v5, v1, v1)
        case 128:
            v1 -= 42
        case -1:
            fmt.Println(v5)
        default:
            v34 := GenS5{F0: (v7 & (-(fi33))), F1: v32, F2: 9223372036854775807}
            _ = v34
        }
        v35 := v4[(v1 * (v1 ^ (^(v1))))]
        _ = fi33
        _ = v35
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
