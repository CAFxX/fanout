package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 string }
type GenS4 struct { F0 bool; F1 bool; F2 bool }
func genFn0(p0 uint, p1 uint8) int {
    var v7 int32 = 0
    var v8 uint64 = 433
    for fi9 := 0; fi9 < 9; fi9++ {
        v8 = (^(v8))
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 int64, p1 uint8, p2 int16) uint {
    p1 = p1
    v10 := GenS1{F0: 4294967295, F1: -1, F2: 18446744073709551615}
    p2 -= 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2
}

func main() {
    v1 := []uint32{2147483647}
    v2 := make(map[int]uint32)
    var v3 float32 = 3.14159
    var v4 int16 = 42
    var v5 uint16 = 65535
    var v6 int8 = 42
    v11 := GenS1{F0: 28, F1: (v6 * 42), F2: 298}
    fmt.Println(v3)
    if (942 > 42) {
        if false {
            if true {
                var v12 uint32 = 240
                v4 = 1000
                _ = v12
            }
            var v13 float32 = float32(-2.25)
            delete(v2, 0)
            _ = v13
        } else {
            v14 := GenS1{F0: 917, F1: (v6 % ((100) | 1)), F2: 1}
            v6 -= (v6 - (v6 - 98))
            fmt.Println(v6, v5)
            _ = v14
        }
    }
    v6 += (v6 / (((v6 * v6)) | 1))
    if true {
        for fi15 := 0; fi15 < 13; fi15++ {
            if false {
                switch (fi15 ^ fi15) {
                case 42:
                    v6 += v6
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    v3 += 0.1
                    v6 = v6
                    if false {
                        v6 = (v6 / ((v6) | 1))
                        v3 = v3
                        fmt.Println(v3, v5)
                    } else {
                        v3 = v3
                        v4 -= -6
                    }
                    _ = fi16
                }
                v3 += (v3 - v3)
            } else {
                v6 -= v6
                switch (fi15 - (fi15 << ((255) & 63))) {
                case 712:
                    _ = len(v1)
                default:
                    switch "a" {
                    case "":
                        if true {
                            var v17 int32 = 0
                            fmt.Println(v4, fi15)
                            v17 -= v17
                            _ = v17
                        }
                    case "hello":
                        v3 -= v3
                    case "\n\t":
                        v18 := GenS2{F0: "", F1: (v4 | (v4 + (^(v4))))}
                        _ = v18
                    }
                }
                v2[fi15] = 1
            }
            for fi19 := 0; fi19 < 19; fi19++ {
                switch ("a" + ("a" + ("" + "foo bar"))) {
                case "\n\t":
                    _ = len(v2)
                default:
                    if ((((("foo bar" + "hello") + ("hello" + "a")) + "foo bar") < "hello") || true) {
                        if false {
                            _ = len(v1)
                            v20 := GenS1{F0: 18446744073709551615, F1: (v6 | v6), F2: 1}
                            _ = v20
                        } else {
                            v21 := GenS2{F0: ("\n\t" + "\n\t"), F1: v4}
                            _ = v21
                        }
                        fmt.Println(v5, fi19)
                        v22 := GenS3{F0: 190, F1: ((("foo bar" + "\n\t") + ("" + "foo bar")) + ("a" + "a")), F2: "a"}
                        _ = v22
                    }
                }
                v23 := GenS4{F0: false, F1: (65535 != 0), F2: true}
                _ = fi19
                _ = v23
            }
            _ = fi15
        }
    }
    var v24 float32 = (0.0 + v3)
    var v25 int16 = (v4 + (v4 ^ 42))
    var v26 uint = 18446744073709551615
    var v27 string = "hello"
    var v28 uint16 = v5
    v27 = v27
    for fi29 := 0; fi29 < 10; fi29++ {
        v24 += float32(float64(float32(-2.25)))
        _ = fi29
    }
    v28 = (v5 / (((v5 ^ (v28 - (^(v28))))) | 1))
    v30 := GenS2{F0: ("hello" + "hello"), F1: 1}
    fmt.Println(v28, v25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
