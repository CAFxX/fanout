package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{65535}
    v3 := []int{32767, 27, -9223372036854775808, 128}
    var v4 uint16 = 0
    var v5 uint32 = 21
    v6 := []int{-339, 4294967295}
    v3 = append(v3, 2147483647)
    v4 -= (v4 + (v4 / (((^(v4))) | 1)))
    switch -272 {
    case 2147483647:
        v5 += 65535
    default:
        var v7 int8 = 0
        _ = v7
    }
    var v8 int32 = 42
    for fi9 := 0; fi9 < 4; fi9++ {
        v8 += v8
        _ = fi9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := v1[128]
        var v12 uint64 = 18446744073709551615
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    switch -368 {
    case 4294967295:
        switch 886 {
        case -2147483648:
            if (592 != 65535) {
                for fi13 := 0; fi13 < 14; fi13++ {
                    if (18446744073709551615 >= 255) {
                        fmt.Println(v5, v4, v8)
                        v14 := GenS1{F0: (v5 / (((v5 >> ((2) & 63))) | 1)), F1: (v8 << ((203) & 63))}
                        _ = v14
                    } else {
                        switch fi13 {
                        case 32767:
                            switch "" {
                            case "hello":
                                v15 := GenS1{F0: (v5 >> ((748) & 63)), F1: (v8 << ((4294967295) & 63))}
                                _ = v15
                            case "":
                                fmt.Println(v8)
                            default:
                                v8 += v8
                            }
                        case 341:
                            v16 := GenS2{F0: ""}
                            _ = v16
                        case -9223372036854775808:
                            if (15 <= (v4 ^ 2)) {
                                v4 = 1
                            }
                        }
                        switch "hello" {
                        case "\n\t":
                            v8 += v8
                        case "a":
                            fmt.Println(v8, v5, v8)
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        _ = len(v3)
                    case "hello":
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v1[fi13] = (v5 << ((996) & 63))
                            _ = fi17
                        }
                    case "":
                        if false {
                            if true {
                                v4 += 32767
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = (fi13 * int(v4))
                                }
                            }
                            v5 = v5
                            v18 := GenS3{F0: -9223372036854775808, F1: -9223372036854775808}
                            _ = v18
                        } else {
                            v5 -= (v5 / ((2147483648) | 1))
                        }
                    default:
                        var v19 uint16 = v4
                        _ = v19
                    }
                    v20 := GenS4{F0: (v5 & 1)}
                    _ = fi13
                    _ = v20
                }
            }
        }
    case -643:
        switch -1 {
        case 766:
            if (false && true) {
                v5 = v5
                var v21 float32 = ((0.0 * 0.1) / 1.0)
                _ = v21
            } else {
                _ = len(v1)
            }
        case 0:
            if false {
                v4 = v4
            } else {
                var v22 float64 = 0.1
                var v23 uint32 = v5
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v22
                _ = v23
            }
        case 32768:
            v5 = (v5 & 0)
        }
    case 32767:
        v8 -= v8
    default:
        switch 0 {
        case 256:
            switch (("foo bar" + ("" + "a")) + "hello") {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        case 292:
            v4 += v4
        case -2147483648:
            v24 := GenS2{F0: (((("a" + "foo bar") + "") + "a") + "hello")}
            _ = v24
        default:
            switch (("" + ("\n\t" + "")) + "a") {
            case "a":
                var v25 int16 = 1000
                _ = v25
            case "\n\t":
                var v26 uint32 = v5
                _ = v26
            }
        }
    }
    v27 := GenS5{F0: 643, F1: int64(v4)}
    for fi28 := 0; fi28 < 13; fi28++ {
        if (!(!true)) {
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 bool = false
                if ((v4 >> ((65535) & 63)) >= v4) {
                    _ = len(v1)
                }
                _ = fi29
                _ = v30
            }
        }
        _ = fi28
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        fmt.Println(v5)
        switch "a" {
        case "hello":
            switch fi31 {
            case 2147483647:
                v4 = v4
            default:
                var v32 uint8 = 118
                _ = v32
            }
        case "a":
            v8 = -878
        }
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
