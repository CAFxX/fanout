package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 int = 0
    var v2 uint16 = 32767
    var v3 float64 = (1e10 + ((0.1 * 1e10) + (float64(1e10) - 0.0)))
    var v4 uint32 = 2147483647
    v5 := []uint32{(^(v4)), (v4 % ((v4) | 1)), (v4 / ((v4) | 1))}
    var v6 uint8 = 34
    var v7 string = (("" + ("\n\t" + "a")) + ("hello" + "a"))
    v8 := GenS1{F0: v4, F1: 0, F2: 255}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int64 = 1
        if false {
            var v11 uint64 = 42
            for fi12 := 0; fi12 < 20; fi12++ {
                switch 127 {
                case 127:
                    switch 0 {
                    case 576:
                        var v13 float32 = 1e10
                        _ = v13
                    }
                case 32768:
                    if false {
                        var v14 int = (^(fi9))
                        v3 += v3
                        _ = len(v5)
                        v6 -= v6
                        v7 = "hello"
                        _ = v14
                    }
                default:
                    v1 += 572
                }
                if (true || false) {
                    v1 -= (fi12 << ((42) & 63))
                    var v15 int = int(v2)
                    _ = v15
                } else {
                    v4 += (v4 & (v4 | (v4 << ((255) & 63))))
                    v7 = ""
                }
                if false {
                    if false {
                        if false {
                            v10 += 932
                        }
                    }
                    v16 := GenS2{F0: 607, F1: (v10 < v10), F2: v7}
                    switch ("hello" + (("a" + v7) + "foo bar")) {
                    case "foo bar":
                        v5 = append(v5, v4)
                    case "a":
                        v4 = (v4 >> ((4294967295) & 63))
                    default:
                        v11 += (v11 * (v11 - (v11 + v11)))
                    }
                    _ = v16
                }
                v4 = 2
                v11 -= 2
                v17 := GenS2{F0: 0, F1: true, F2: v7}
                _ = fi12
                _ = v17
            }
            if false {
                v4 -= v4
            } else {
                var v18 int8 = -114
                _ = v18
            }
            _ = v11
        } else {
            v19 := GenS3{F0: v1}
            for fi20 := 0; fi20 < 7; fi20++ {
                v3 -= v3
                switch (v7 + ((("hello" + (("\n\t" + (v7 + ((v7 + v7) + (((v7 + ("a" + v7)) + (v7 + v7)) + v7)))) + ((v7 + v7) + v7))) + (v7 + v7)) + "\n\t")) {
                case "foo bar":
                    v3 = 1e10
                default:
                    v21 := GenS4{F0: -1}
                    _ = v21
                }
                var v22 float32 = 3.14159
                _ = fi20
                _ = v22
            }
            _ = v19
        }
        _ = fi9
        _ = v10
    }
    v7 = ("\n\t" + "\n\t")
    if (!(-2.25 >= (((v3 / (v3 / 1.0)) - (v3 / v3)) + 1.5))) {
        fmt.Println(v7, v3, v7)
        v23 := GenS5{F0: int32(v4), F1: 51, F2: -531}
        _ = v23
    }
    var v24 uint64 = 65
    var v25 uint16 = v2
    v26 := GenS4{F0: -2147483648}
    v6 = (v6 % ((127) | 1))
    if (218 > 128) {
        v27 := GenS4{F0: 42}
        fmt.Println(v24)
        v2 = v2
        _ = v27
    } else {
        v28 := GenS4{F0: 2147483647}
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 string = "foo bar"
        fmt.Println(v24, v3)
        _ = fi29
        _ = v30
    }
    v2 = 32767
    v5 = append(v5, v4)
    v25 = (^(v2))
    if (!false) {
        var v31 uint32 = (v4 ^ (v4 >> ((4294967295) & 63)))
        _ = v31
    } else {
        v32 := GenS6{F0: (v1 | v1)}
        if false {
            v2 += 321
            var v33 int32 = -713
            _ = v33
        } else {
            v34 := GenS7{F0: 944, F1: -1}
            _ = len(v5)
            for fi35 := 0; fi35 < 5; fi35++ {
                v36 := GenS8{F0: 32767}
                fmt.Println(v6)
                for fi37 := 0; fi37 < 7; fi37++ {
                    switch "hello" {
                    case "\n\t":
                        v4 += 1
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v7, v25, v25)
                    v25 -= 576
                    _ = fi37
                }
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
