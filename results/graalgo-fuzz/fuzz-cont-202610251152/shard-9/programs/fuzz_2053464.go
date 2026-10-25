package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64; F1 int }

func main() {
    var v1 uint32 = 668
    v2 := make(map[int]int)
    v3 := []uint32{(v1 & v1)}
    switch 114 {
    case 2147483647:
        for fi4 := 0; fi4 < 10; fi4++ {
            if false {
                fmt.Println(fi4, v1)
            } else {
                switch "foo bar" {
                case "\n\t":
                    v1 -= v1
                case "hello":
                    var v5 uint64 = 1
                    _ = v5
                case "a":
                    if false {
                        v6 := GenS1{F0: v1, F1: 42}
                        v7 := v2[313]
                        var v8 float64 = 3.14159
                        _ = v6
                        _ = v7
                        _ = v8
                    }
                default:
                    switch ("hello" + "hello") {
                    case "foo bar":
                        v1 += v1
                    case "":
                        var v9 bool = false
                        _ = v9
                    case "a":
                        var v10 uint = 559
                        _ = v10
                    default:
                        var v11 bool = ((uint(v1) < 474) || true)
                        _ = v11
                    }
                }
                for fi12 := 0; fi12 < 14; fi12++ {
                    fmt.Println(fi4)
                    _ = fi12
                }
            }
            if ((true || true) && (1e10 >= ((-2.25 + 1.5) + 0.1))) {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v1 += (v1 & 267)
                    v1 = v1
                    switch (fi4 + (^(fi13))) {
                    case 128:
                        if (true || ((-138 >= -405) || true)) {
                            fmt.Println(fi13, fi13)
                        } else {
                            if (759 < 823) {
                                v1 += 4294967295
                                delete(v2, 32768)
                                v1 = (v1 - v1)
                                _ = len(v3)
                            } else {
                                var v14 uint16 = 0
                                v1 = 0
                                if true {
                                    var v15 uint16 = (v14 >> ((0) & 63))
                                    var v16 int16 = -32768
                                    v15 = (v14 >> ((674) & 63))
                                    _ = v15
                                    _ = v16
                                }
                                switch ((("" + ("\n\t" + "\n\t")) + "hello") + "a") {
                                case "\n\t":
                                    if (0.0 == 1.5) {
                                        var v17 int = 595
                                        var v18 int64 = -1
                                        _ = v17
                                        _ = v18
                                    }
                                case "foo bar":
                                    var v19 uint = 59
                                    _ = v19
                                case "hello":
                                    fmt.Println(fi13, fi13)
                                default:
                                    v1 = 1
                                }
                                _ = v14
                            }
                            v1 -= (v1 - 65535)
                        }
                    default:
                        v1 = v1
                    }
                    _ = fi13
                }
                fmt.Println(fi4, v1)
                var v20 int16 = 138
                _ = v20
            }
            if (2147483647 >= 42) {
                _ = len(v3)
                fmt.Println(v1, fi4)
                v1 = (^(v1))
            } else {
                v21 := GenS1{F0: (v1 ^ v1), F1: -1}
                _ = v21
            }
            _ = fi4
        }
    case -9223372036854775808:
        var v22 uint64 = 42
        _ = v22
    }
    if true {
        v1 = v1
    }
    var v23 float32 = 1.5
    _ = len(v3)
    var v24 int64 = 480
    v25 := GenS2{F0: v24, F1: -598}
    switch -471 {
    case 2147483647:
        v23 = (v23 / 1.0)
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        if len(v3) > 0 {
            v3[0] = 877
        }
        v24 = (v24 >> ((1) & 63))
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
