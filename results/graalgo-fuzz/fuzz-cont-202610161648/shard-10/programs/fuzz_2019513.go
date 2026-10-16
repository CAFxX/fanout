package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint16, p1 int8, p2 int32) uint16 {
    p1 = (p1 / (((p1 / ((p1) | 1))) | 1))
    v7 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (^(p0))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := []uint32{2}
    var v5 uint64 = 0
    v6 := make(map[int]string)
    var v8 uint64 = (v5 & (v5 % ((4294967295) | 1)))
    delete(v3, 42)
    for fi9 := 0; fi9 < 6; fi9++ {
        v5 -= 42
        for fi10 := 0; fi10 < 1; fi10++ {
            if (!false) {
                delete(v1, 1)
                var v11 int16 = 1000
                _ = v11
            }
            _ = fi10
        }
        if false {
            var v12 uint8 = 1
            if (!false) {
                v8 = (v5 + (v8 + 1))
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 int64 = 0
                    switch "\n\t" {
                    case "\n\t":
                        switch "a" {
                        case "a":
                            if ("hello" > ((("foo bar" + "") + "") + "a")) {
                                var v15 uint64 = v8
                                v8 = (v8 & (v15 ^ v8))
                                v12 = (^(v12))
                                _ = v15
                            } else {
                                var v16 float32 = 0.0
                                v5 = v8
                                _ = v16
                            }
                        case "\n\t":
                            v17 := v1[15]
                            _ = v17
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case "a":
                        var v18 uint8 = (v12 >> ((255) & 63))
                        _ = v18
                    case "":
                        v8 -= (v5 * (v8 * v8))
                    default:
                        fmt.Println(v12)
                    }
                    _ = fi13
                    _ = v14
                }
                var v19 uint8 = 2
                _ = v19
            } else {
                var v20 uint64 = v5
                var v21 float32 = (((float32(0.0) * -2.25) + 0.1) / 1.0)
                v22 := GenS2{F0: 149, F1: v5}
                _ = v20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS3{F0: (fi9 % ((fi9) | 1)), F1: -1, F2: 42}
                if (v12 < (v12 * v12)) {
                    if false {
                        v8 = v5
                    } else {
                        v8 += v5
                        fmt.Println(fi23, fi23)
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        delete(v1, 609)
                    default:
                        switch ("foo bar" + "hello") {
                        case "hello":
                            var v25 uint16 = 264
                            _ = v25
                        case "foo bar":
                            var v26 float32 = -2.25
                            _ = v26
                        case "":
                            var v27 int32 = 1
                            _ = v27
                        default:
                            switch "a" {
                            case "":
                                var v28 string = "hello"
                                _ = v28
                            }
                        }
                    }
                }
                _ = fi23
                _ = v24
            }
            _ = v12
        }
        _ = fi9
    }
    v5 -= v5
    var v29 uint32 = 920
    v30 := GenS4{F0: 0, F1: "foo bar"}
    if (!false) {
        switch 32767 {
        case -924:
            v5 = (v5 * v8)
        case -2147483648:
            var v31 int = 0
            _ = v31
        case 1:
            var v32 bool = (32767 >= 4294967295)
            _ = v32
        }
    } else {
        if (!("hello" != "foo bar")) {
            for fi33 := 0; fi33 < 7; fi33++ {
                if ((v5 >> ((69) & 63)) == (v8 - v5)) {
                    v34 := GenS5{F0: -1}
                    _ = v34
                }
                switch (fi33 / ((-272) | 1)) {
                case 42:
                    var v35 int64 = -9223372036854775808
                    _ = v35
                case 501:
                    if true {
                        if (50 != -32768) {
                            switch (fi33 >> ((815) & 63)) {
                            case 32768:
                                _ = len(v4)
                            case 9223372036854775807:
                                if true {
                                    fmt.Println(v8)
                                }
                            default:
                                if true {
                                    switch fi33 {
                                    case 440:
                                        v8 = 18446744073709551615
                                    case -706:
                                        var v36 int8 = int8(v29)
                                        _ = v36
                                    default:
                                        var v37 uint = 18446744073709551615
                                        _ = v37
                                    }
                                    for fi38 := 0; fi38 < 13; fi38++ {
                                        var v39 bool = false
                                        _ = fi38
                                        _ = v39
                                    }
                                } else {
                                    v29 = v29
                                    v40 := GenS2{F0: 2, F1: v5}
                                    _ = v40
                                }
                            }
                        }
                    }
                case -9223372036854775808:
                    v41 := GenS2{F0: 2, F1: (v5 >> ((426) & 63))}
                    _ = v41
                default:
                    for fi42 := 0; fi42 < 17; fi42++ {
                        _ = len(v4)
                        _ = fi42
                    }
                }
                v43 := GenS6{F0: v8, F1: -1}
                _ = fi33
                _ = v43
            }
            v5 = (v8 & v5)
        }
    }
    fmt.Println(v8)
    v44 := GenS4{F0: 9223372036854775807, F1: ("\n\t" + "\n\t")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v8
}
