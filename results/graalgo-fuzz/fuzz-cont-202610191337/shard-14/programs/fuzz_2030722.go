package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int64; F2 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 int, p2 uint64) string {
    if false {
        p0 += (p0 - p1)
        for fi6 := 0; fi6 < 4; fi6++ {
            if true {
                p2 = p2
            } else {
                var v7 int64 = 0
                var v8 float32 = float32(0.1)
                v9 := []int{p0, fi6, p0}
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 2, F1: 42, F2: 70}
            _ = fi6
            _ = v10
        }
        if ("" > "a") {
            var v11 int = p0
            _ = v11
        }
    }
    if false {
        if true {
            p2 = (p2 >> ((992) & 63))
            switch p0 {
            case -621:
                var v12 uint64 = 18446744073709551615
                _ = v12
            default:
                p2 = p2
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 uint64 = (p2 & 4294967295)
                for fi15 := 0; fi15 < 7; fi15++ {
                    v14 = p2
                    var v16 int = -743
                    _ = fi15
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
        }
    } else {
        var v17 bool = (0.0 < float64(3.14159))
        _ = v17
    }
    v18 := []int{p1}
    p1 += (p1 / ((p0) | 1))
    v19 := GenS2{F0: 100}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return "a"
}

func main() {
    v1 := []uint32{}
    var v2 bool = false
    var v3 int8 = 0
    v4 := []int{-418}
    var v5 uint32 = 2147483647
    var v20 uint64 = 4294967295
    fmt.Println(v2)
    switch 127 {
    case 127:
        var v21 int64 = -810
        _ = v21
    case 564:
        for fi22 := 0; fi22 < 8; fi22++ {
            switch (fi22 << ((328) & 63)) {
            case 163:
                v5 -= (v5 << ((255) & 63))
            case -170:
                var v23 string = "\n\t"
                _ = v23
            default:
                var v24 float32 = 0.1
                _ = v24
            }
            switch "foo bar" {
            case "\n\t":
                var v25 uint8 = 128
                _ = v25
            default:
                v26 := GenS2{F0: v3}
                _ = v26
            }
            _ = fi22
        }
    case 4294967295:
        v1 = append(v1, v5)
    default:
        if (127 == 0) {
            var v27 float32 = (((-2.25 - float32((0.0 * (0.1 - 1e10)))) / 1e10) + (float32((0.1 - -2.25)) / (3.14159 / float32(-2.25))))
            var v28 int32 = -1
            fmt.Println(v3, v5, v5)
            _ = v27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        for fi30 := 0; fi30 < 3; fi30++ {
            v4 = append(v4, 127)
            if (127 < 26) {
                v31 := GenS3{F0: "a", F1: int64(v20), F2: 80}
                if (-483 < 9223372036854775807) {
                    switch 0 {
                    case -565:
                        v32 := GenS3{F0: "foo bar", F1: -977, F2: -1}
                        _ = v32
                    default:
                        if v2 {
                            switch "hello" {
                            case "\n\t":
                                var v33 int = fi29
                                _ = v33
                            case "hello":
                                v34 := GenS4{F0: ("hello" + "\n\t")}
                                _ = v34
                            case "foo bar":
                                var v35 int32 = 0
                                _ = v35
                            }
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v31
            }
            var v36 uint = 794
            fmt.Println(v5, v5, v2)
            var v37 float32 = 1e10
            _ = len(v1)
            _ = fi30
            _ = v36
            _ = v37
        }
        _ = fi29
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int16 = -327
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        v20 -= (v20 ^ 0)
        _ = fi40
    }
    var v41 uint64 = v20
    v42 := GenS2{F0: v3}
    var v43 bool = true
    if len(v1) > 0 {
        v1[0] = (v5 | (v5 & (v5 % ((v5) | 1))))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
