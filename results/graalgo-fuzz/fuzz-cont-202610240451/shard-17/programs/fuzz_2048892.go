package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 int32) bool {
    var v6 uint64 = 4294967295
    v7 := []int{619, 733, 295}
    var v8 float32 = (-2.25 / 1.0)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v6 < (v6 % (((v6 + (^(v6)))) | 1)))
}
func genFn1(p0 int32, p1 int8, p2 uint) int8 {
    v9 := GenS1{F0: 4294967295}
    v10 := GenS1{F0: 322}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -111
}

func main() {
    v1 := []int{32767, 42, 2147483647}
    var v2 int64 = -458
    v3 := make(map[uint32]string)
    var v4 int8 = 42
    var v5 int32 = 2147483647
    switch "a" {
    case "\n\t":
        var v11 int32 = (^(v5))
        _ = v11
    case "hello":
        fmt.Println(v5)
    case "foo bar":
        for fi12 := 0; fi12 < 17; fi12++ {
            v4 += -36
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := v3[2147483648]
                v2 -= (v2 & (v2 << ((4294967295) & 63)))
                switch (fi13 << ((76) & 63)) {
                case 745:
                    if (1.5 >= float64((0.1 * 0.1))) {
                        v15 := GenS1{F0: 70}
                        fmt.Println(v5, v5, v2)
                        var v16 uint64 = 1
                        _ = v15
                        _ = v16
                    }
                case 1:
                    v4 -= v4
                case 2147483647:
                    switch fi13 {
                    case -698:
                        v17 := v3[2]
                        _ = v17
                    default:
                        switch (fi13 >> ((255) & 63)) {
                        case -1:
                            v2 = v2
                        }
                        v2 += v2
                        if true {
                            switch (fi12 | fi13) {
                            case 2147483647:
                                var v18 uint8 = 127
                                _ = v18
                            default:
                                v4 -= (v4 << ((18446744073709551615) & 63))
                            }
                            v2 = (v2 << ((18446744073709551615) & 63))
                            v5 = v5
                            v19 := GenS1{F0: 18446744073709551615}
                            switch fi12 {
                            case 0:
                                var v20 uint8 = 178
                                _ = v20
                            case -52:
                                v14 = v14
                            default:
                                var v21 int64 = -1
                                _ = v21
                            }
                            _ = v19
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
            v5 += v5
            _ = fi12
        }
    }
    _ = len(v3)
    v2 += 0
    v22 := GenS2{F0: -1, F1: false}
    v23 := GenS3{F0: (^(v4))}
    switch ("" + "foo bar") {
    case "":
        if false {
            var v24 float32 = float32((1.5 - ((1.5 + 3.14159) / 1.0)))
            v25 := GenS1{F0: 639}
            _ = v24
            _ = v25
        } else {
            var v26 int32 = v5
            if (!true) {
                v27 := GenS1{F0: 2}
                _ = v27
            } else {
                for fi28 := 0; fi28 < 5; fi28++ {
                    if false {
                        for fi29 := 0; fi29 < 11; fi29++ {
                            v2 = (v2 << ((255) & 63))
                            _ = fi29
                        }
                    }
                    v2 -= (v2 % ((0) | 1))
                    switch "foo bar" {
                    case "":
                        v2 = -62
                    }
                    _ = fi28
                }
                v26 = (v26 / (((v26 << ((255) & 63))) | 1))
                v2 -= 526
            }
            v30 := GenS3{F0: (v4 >> ((18446744073709551615) & 63))}
            _ = v26
            _ = v30
        }
    }
    var v31 int16 = 1
    fmt.Println(v31, v31)
    var v32 int = 127
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 int32 = v5
        v35 := GenS3{F0: (-(v4))}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
