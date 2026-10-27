package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 bool; F2 int8 }
type GenS4 struct { F0 int32; F1 uint }
func genFn0(p0 int8, p1 uint64, p2 uint32) uint {
    fmt.Println(p2, p2, p0)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 581
}

func main() {
    var v1 int8 = -1
    var v2 int8 = v1
    var v3 uint8 = 205
    v4 := make(map[int]uint32)
    var v5 float64 = (float64((-2.25 * 0.0)) * float64((float32(-2.25) + float32(((-2.25 * (1.5 / -2.25)) - 0.0)))))
    var v6 int64 = -1
    var v7 uint16 = 2
    if false {
        var v8 int = -736
        var v9 string = ("foo bar" + "")
        _ = len(v4)
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7)
        var v10 uint = 255
        switch (("a" + ("" + (("\n\t" + "foo bar") + "hello"))) + ("foo bar" + "")) {
        case "hello":
            if ("foo bar" > "") {
                v11 := GenS1{F0: 0, F1: 2147483647, F2: -96}
                var v12 int64 = (v6 << ((2) & 63))
                _ = v11
                _ = v12
            } else {
                v7 = (^(v7))
                for fi13 := 0; fi13 < 14; fi13++ {
                    if false {
                        for fi14 := 0; fi14 < 6; fi14++ {
                            if (!true) {
                                v15 := GenS2{F0: 18446744073709551615}
                                _ = v15
                            }
                            var v16 bool = true
                            if false {
                                var v17 uint32 = 2147483647
                                switch "foo bar" {
                                case "foo bar":
                                    var v18 uint16 = (^(v7))
                                    _ = v18
                                case "":
                                    v19 := GenS1{F0: (fi14 & 32767), F1: fi13, F2: int8(v17)}
                                    _ = v19
                                }
                                _ = v17
                            }
                            _ = fi14
                            _ = v16
                        }
                        v20 := v4[(fi13 / ((32768) | 1))]
                        var v21 bool = true
                        _ = v20
                        _ = v21
                    } else {
                        fmt.Println(v2)
                        v22 := GenS3{F0: "hello", F1: false, F2: (v1 * v1)}
                        var v23 int16 = -32768
                        _ = v22
                        _ = v23
                    }
                    if false {
                        v1 += (v2 / (((v2 | (v2 % (((v2 * (v1 * (v2 | 1)))) | 1)))) | 1))
                        var v24 int8 = (v1 | v2)
                        _ = v24
                    }
                    v25 := GenS4{F0: -805, F1: 567}
                    _ = fi13
                    _ = v25
                }
            }
        default:
            v26 := []int{634}
            _ = v26
        }
        _ = v10
    }
    var v27 uint16 = v7
    v27 -= v27
    v28 := GenS4{F0: -1, F1: 198}
    for fi29 := 0; fi29 < 8; fi29++ {
        v7 = v27
        _ = fi29
    }
    v2 += int8(v7)
    var v30 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
