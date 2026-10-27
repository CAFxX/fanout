package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    var v4 float64 = 3.14159
    v5 := []uint32{744}
    var v6 int64 = -183
    v6 = v6
    var v7 int8 = -128
    v3[((("" + "\n\t") + ("" + "")) + ("\n\t" + "\n\t"))] = false
    _ = len(v5)
    var v8 uint64 = 18446744073709551615
    for fi9 := 0; fi9 < 14; fi9++ {
        v2 += v2
        v2 = (v2 - (v2 % ((v2) | 1)))
        v4 = (v4 / v4)
        _ = fi9
    }
    v4 -= v4
    if (true && false) {
        switch 42 {
        case 938:
            var v10 int32 = -626
            _ = v10
        case 1:
            var v11 float32 = ((0.1 + 1.5) + ((3.14159 - 1e10) / ((0.0 - 0.0) - (-2.25 * 0.1))))
            _ = v11
        case 32767:
            v2 += (v2 & (^(v2)))
        default:
            v12 := GenS1{F0: 127, F1: 32767}
            _ = v12
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int8 = (^(v7))
            _ = fi13
            _ = v14
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 int8 = v7
            _ = fi15
            _ = v16
        }
        var v17 int16 = -32768
        _ = v17
    }
    if (!("a" < "foo bar")) {
        if len(v5) > 0 {
            v5[0] = 561
        }
        if (!false) {
            v18 := GenS2{F0: v8, F1: (((false || (!true)) || ((0 >= 1000) || (!((v2 * (v2 / ((994) | 1))) > 2)))) && (498 >= 65535)), F2: v6}
            switch 32768 {
            case 4294967295:
                v19 := GenS3{F0: 2}
                _ = v19
            case 32768:
                if (0 >= v7) {
                    var v20 uint = (v2 ^ v2)
                    v20 = 0
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v2 += v2
                        for fi22 := 0; fi22 < 12; fi22++ {
                            v8 = (v8 * uint64(v2))
                            v8 = v8
                            v23 := GenS1{F0: (v7 + v7), F1: 42}
                            _ = fi22
                            _ = v23
                        }
                        _ = fi21
                    }
                    _ = v20
                } else {
                    var v24 bool = false
                    _ = v24
                }
            }
            v5 = append(v5, 933)
            _ = v18
        }
        fmt.Println(v7, v7, v7)
    } else {
        v4 = ((0.0 + v4) / ((float64(-2.25) + 0.1) / float64(3.14159)))
        fmt.Println(v8, v8, v6)
    }
    v25 := GenS4{F0: 0, F1: 1}
    for fi26 := 0; fi26 < 14; fi26++ {
        if true {
            var v27 float32 = (float32(3.14159) - 1.5)
            for fi28 := 0; fi28 < 7; fi28++ {
                v8 = v8
                _ = fi28
            }
            _ = v27
        } else {
            v2 = v2
            for fi29 := 0; fi29 < 9; fi29++ {
                v7 = v7
                if len(v5) > 0 {
                    v5[0] = 310
                }
                _ = fi29
            }
            fmt.Println(fi26, v2)
        }
        _ = fi26
    }
    var v30 string = ("a" + "foo bar")
    v6 = (v6 ^ (v6 % (((v6 + -540)) | 1)))
    for fi31 := 0; fi31 < 18; fi31++ {
        v3[(v30 + ((v30 + "foo bar") + (v30 + "foo bar")))] = ((!(v7 >= -1)) || ((v6 / ((v6) | 1)) >= (^(v6))))
        v6 -= (v6 % (((v6 + 42)) | 1))
        _ = fi31
    }
    v3[v30] = true
    if false {
        _ = len(v1)
        v32 := GenS4{F0: 0, F1: (v2 / ((v2) | 1))}
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
