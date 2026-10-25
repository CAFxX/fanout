package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int64; F2 int }
type GenS3 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := []uint32{2147483647}
    var v2 int16 = 1000
    v3 := make(map[string]uint32)
    v4 := []uint32{204, 808, 42}
    v5 := []uint32{2147483647, 455, 0}
    v2 -= (v2 - (v2 % ((1000) | 1)))
    _ = len(v5)
    if (1 >= 42) {
        switch "hello" {
        case "foo bar":
            var v6 int8 = 127
            _ = v6
        case "\n\t":
            for fi7 := 0; fi7 < 5; fi7++ {
                v2 -= (v2 & (v2 >> ((95) & 63)))
                _ = len(v5)
                switch (fi7 >> ((0) & 63)) {
                case 32767:
                    v8 := GenS1{F0: 127}
                    _ = v8
                case -543:
                    for fi9 := 0; fi9 < 16; fi9++ {
                        if false {
                            fmt.Println(v2)
                            if (0 >= 65535) {
                                fmt.Println(fi7, fi7, fi9)
                                v2 = -458
                                v2 = v2
                            } else {
                                v10 := v3["foo bar"]
                                var v11 uint64 = 2
                                _ = v10
                                _ = v11
                            }
                            fmt.Println(fi9, v2)
                            var v12 uint32 = 772
                            _ = v12
                        }
                        var v13 uint32 = 334
                        v14 := GenS2{F0: "foo bar", F1: 866, F2: fi7}
                        _ = fi9
                        _ = v13
                        _ = v14
                    }
                case -2147483648:
                    if true {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            fmt.Println(fi7, fi7)
                            v2 = -677
                            _ = fi15
                        }
                        v2 += (v2 ^ v2)
                    } else {
                        if false {
                            v16 := GenS2{F0: ((("\n\t" + "hello") + "hello") + (("\n\t" + "\n\t") + "a")), F1: 9223372036854775807, F2: (-(fi7))}
                            if len(v4) > 0 {
                                v4[0] = 2147483648
                            }
                            var v17 uint = 18446744073709551615
                            _ = v16
                            _ = v17
                        }
                        switch "\n\t" {
                        case "":
                            for fi18 := 0; fi18 < 9; fi18++ {
                                var v19 uint = 4294967295
                                _ = fi18
                                _ = v19
                            }
                        case "hello":
                            var v20 string = "foo bar"
                            _ = v20
                        default:
                            fmt.Println(fi7, fi7)
                        }
                    }
                }
                _ = fi7
            }
        case "":
            v21 := GenS1{F0: 141}
            _ = v21
        }
    }
    var v22 bool = (!false)
    var v23 float64 = 1e10
    fmt.Println(v22, v23, v22)
    for fi24 := 0; fi24 < 7; fi24++ {
        v23 = float64((float32(v23) / -2.25))
        v25 := GenS3{F0: 42, F1: 0, F2: 255}
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: 53}
    v27 := GenS5{F0: 0}
    var v28 int8 = -1
    v29 := GenS6{F0: -566, F1: v28}
    if (18446744073709551615 > 1) {
        if ((v28 << ((2) & 63)) >= v28) {
            var v30 int64 = -632
            if true {
                v31 := GenS7{F0: 9223372036854775807, F1: 1}
                v32 := GenS1{F0: 128}
                switch 42 {
                case 127:
                    if ((v30 * v30) >= 232) {
                        v30 = v30
                        v33 := GenS1{F0: 165}
                        _ = v33
                    } else {
                        var v34 uint16 = 2
                        _ = v34
                    }
                default:
                    v22 = v22
                }
                _ = v31
                _ = v32
            } else {
                var v35 uint = 2
                v36 := GenS8{F0: 792, F1: 315}
                fmt.Println(v35, v22)
                _ = v35
                _ = v36
            }
            _ = v30
        }
    }
    v2 = (v2 * 6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
