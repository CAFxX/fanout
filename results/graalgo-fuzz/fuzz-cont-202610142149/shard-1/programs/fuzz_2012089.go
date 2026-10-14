package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 uint = 255
    v4 := make(map[int]bool)
    v3 -= (v3 % (((v3 >> (((v3 & v3)) & 63))) | 1))
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 float64 = (float64((float32(-2.25) - -2.25)) + 0.0)
        var v7 int = (fi5 | fi5)
        var v8 bool = true
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{32767, 9223372036854775807, 800, 0}
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint64 = 4294967295
            v3 = 495
            if (false || true) {
                _ = len(v4)
                v3 = v3
                fmt.Println(fi11, v12, v3)
            } else {
                var v13 bool = true
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int = (fi15 | (fi15 * fi15))
        switch (fi15 | -1) {
        case 578:
            _ = len(v9)
        case 1:
            if false {
                switch ("" + "foo bar") {
                case "\n\t":
                    var v17 uint64 = 2
                    _ = v17
                case "":
                    _ = len(v9)
                default:
                    if true {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            _ = len(v9)
                            var v19 int8 = -1
                            var v20 uint8 = 0
                            fmt.Println(v19, v3, fi18)
                            if false {
                                var v21 int16 = 42
                                v21 = (v21 >> ((v3) & 63))
                                var v22 uint64 = 18446744073709551615
                                v23 := GenS1{F0: v14, F1: 2, F2: -1}
                                v24 := GenS2{F0: "hello", F1: -2147483648}
                                _ = v21
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                var v25 int32 = 42
                                switch "\n\t" {
                                case "":
                                    v3 -= 18446744073709551615
                                case "hello":
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        fmt.Println(v14)
                                    default:
                                        switch 32767 {
                                        case -2147483648:
                                            fmt.Println(v3, v25, v20)
                                        default:
                                            var v26 uint16 = 959
                                            var v27 uint64 = 18446744073709551615
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                }
                                _ = v25
                            }
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                }
                var v28 int64 = 707
                _ = v28
            }
        case 32767:
            v1[2] = 42
        default:
            for fi29 := 0; fi29 < 14; fi29++ {
                v16 = v16
                _ = fi29
            }
        }
        fmt.Println(v16, v16)
        _ = fi15
        _ = v16
    }
    v30 := v1[0]
    v31 := GenS3{F0: 2, F1: 1, F2: 326}
    v32 := GenS4{F0: v30}
    v14 = (v14 * uint16(v30))
    fmt.Println(v30, v14, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
