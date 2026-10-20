package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 uint, p1 int64) uint64 {
    p0 = (p0 / (((p0 << ((65535) & 63))) | 1))
    fmt.Println(p1, p1)
    v5 := []int{1, 1, int(p0), 1}
    p0 += (p0 % ((p0) | 1))
    p1 = (p1 << (((^(p0))) & 63))
    _ = p0
    _ = p1
    _ = v5
    return 536
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 uint32 = 2147483647
    if false {
        if true {
            switch "\n\t" {
            case "a":
                v4 += 2
            default:
                if ((584 < 2147483647) && (true || true)) {
                    if true {
                        switch "hello" {
                        case "foo bar":
                            v1 -= v1
                        case "a":
                            var v6 float32 = -2.25
                            _ = v6
                        case "\n\t":
                            v7 := GenS1{F0: v1, F1: false, F2: -128}
                            _ = v7
                        }
                        var v8 float32 = 1.5
                        _ = v8
                    } else {
                        switch 32767 {
                        case 127:
                            if (genFn0(462, v1) != 0) {
                                var v9 uint = 2
                                _ = len(v3)
                                if false {
                                    fmt.Println(v9)
                                } else {
                                    switch 32767 {
                                    case -278:
                                        v10 := []int{303}
                                        _ = v10
                                    case -180:
                                        fmt.Println(v1, v9, v1)
                                    default:
                                        var v11 uint64 = genFn0(v9, v1)
                                        _ = v11
                                    }
                                }
                                _ = v9
                            }
                        case 0:
                            v1 += 5
                        case 32767:
                            var v12 uint64 = 18446744073709551615
                            _ = v12
                        }
                        fmt.Println(v4, v1)
                    }
                    v4 += 4294967295
                }
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                switch "a" {
                case "hello":
                    v14 := []int{}
                    _ = v14
                case "":
                    if true {
                        fmt.Println(v1)
                    } else {
                        fmt.Println(v1)
                        v15 := v3[fi13]
                        if true {
                            var v16 uint = 263
                            v17 := GenS2{F0: 65535, F1: -588}
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    }
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 bool = true
                    v4 = (v4 * v4)
                    _ = fi18
                    _ = v19
                }
                _ = fi13
            }
        } else {
            v20 := GenS3{F0: 481}
            _ = v20
        }
        v21 := GenS2{F0: 4294967295, F1: 42}
        v22 := []int{32768}
        _ = v21
        _ = v22
    }
    _ = len(v2)
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := GenS1{F0: -9223372036854775808, F1: false, F2: -1}
        delete(v2, 356)
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 4294967295, F1: 2147483647}
    v26 := v2[332]
    v27 := GenS3{F0: -123}
    var v28 int8 = -128
    delete(v3, v26)
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 int = 171
        _ = fi29
        _ = v30
    }
    v26 = (^(v26))
    v1 = (v1 * (v1 | 9223372036854775807))
    var v31 uint8 = 1
    v32 := []int{v26, 2147483647, 0}
    v4 = 2147483647
    v33 := GenS1{F0: (v1 - (v1 - v1)), F1: false, F2: int8(v4)}
    v26 -= (v26 >> ((2) & 63))
    v34 := GenS4{F0: 1, F1: false, F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
