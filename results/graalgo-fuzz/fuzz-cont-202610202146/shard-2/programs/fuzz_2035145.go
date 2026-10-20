package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 string; F1 int64; F2 int16 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]bool)
    var v3 int32 = -824
    var v4 uint16 = 0
    if false {
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := []int{fi5, 32767, fi5}
            _ = fi5
            _ = v6
        }
        if (246 > 42) {
            for fi7 := 0; fi7 < 8; fi7++ {
                if ((2 >= 2) && true) {
                    v8 := v2[1]
                    v1 = 1e10
                    _ = v8
                } else {
                    v1 = -2.25
                    for fi9 := 0; fi9 < 5; fi9++ {
                        switch "" {
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                fmt.Println(fi7, v1)
                            case "foo bar":
                                if true {
                                    _ = len(v2)
                                    var v10 float64 = 3.14159
                                    _ = v10
                                } else {
                                    v4 = v4
                                }
                            case "":
                                v11 := GenS1{F0: "\n\t", F1: 18446744073709551615}
                                _ = v11
                            }
                        case "a":
                            v12 := GenS2{F0: (v4 << ((124) & 63)), F1: 1, F2: 2}
                            _ = v12
                        default:
                            _ = len(v2)
                        }
                        var v13 int8 = 56
                        _ = fi9
                        _ = v13
                    }
                }
                var v14 uint32 = 658
                _ = fi7
                _ = v14
            }
        } else {
            v15 := []int{827}
            var v16 int64 = 800
            _ = v15
            _ = v16
        }
    }
    v17 := GenS2{F0: (v4 - (v4 + 32767)), F1: 105, F2: 2}
    var v18 int16 = 1000
    v19 := GenS3{F0: 42, F1: 42}
    if ((("" + "foo bar") + ("hello" + ("\n\t" + "a"))) < "\n\t") {
        v18 += v18
        v20 := v2[0]
        _ = v20
    }
    v21 := []int{32767, -488}
    v22 := GenS1{F0: ("hello" + ((("" + "\n\t") + "a") + "hello")), F1: 0}
    v18 = -1
    v1 -= (float32(float64(v1)) - v1)
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(v4, fi23)
            _ = fi24
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            v4 = (v4 - v4)
            switch 2147483647 {
            case 1:
                v26 := GenS3{F0: -9223372036854775808, F1: -1}
                _ = v26
            case 410:
                _ = len(v2)
            case 716:
                fmt.Println(fi25)
                switch ("hello" + ("hello" + "hello")) {
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        v4 += v4
                    case "foo bar":
                        var v27 uint8 = 1
                        _ = v27
                    default:
                        fmt.Println(fi25, v4, fi25)
                    }
                case "a":
                    v18 += 987
                case "":
                    _ = len(v2)
                default:
                    if false {
                        v3 = (v3 >> ((0) & 63))
                        if true {
                            var v28 int8 = 127
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            var v29 uint = 0
                            _ = len(v2)
                            v29 = v29
                            v30 := GenS1{F0: "foo bar", F1: 42}
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        v31 := GenS4{F0: ((("" + "a") + (("\n\t" + "a") + "a")) + ""), F1: 1, F2: (v18 / (((v18 >> ((255) & 63))) | 1))}
                        _ = v31
                    }
                }
                v1 += 0.0
            default:
                fmt.Println(fi23, v18, v3)
            }
            _ = fi25
        }
        var v32 string = "\n\t"
        _ = fi23
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
