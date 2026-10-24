package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64, p1 int32, p2 uint32) int16 {
    var v6 bool = false
    var v7 int8 = 5
    v7 += (v7 << ((1) & 63))
    fmt.Println(v6, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 float64 = 0.0
    var v4 float32 = float32(0.1)
    v5 := []int{0, 32767}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint8 = 140
        var v10 uint8 = v9
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float64 = v3
    fmt.Println(v3, v4)
    if (134 != 127) {
        for fi12 := 0; fi12 < 10; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 uint32 = 2147483647
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 20; fi15++ {
                switch (("a" + "hello") + "foo bar") {
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        fmt.Println(fi12, fi15)
                    case "hello":
                        var v16 int8 = 127
                        _ = v16
                    case "a":
                        if (42 != 255) {
                            v17 := v2[fi15]
                            v11 = (-2.25 + v11)
                            _ = v17
                        } else {
                            switch "a" {
                            case "":
                                v18 := GenS1{F0: 100, F1: 1, F2: ((true && false) || false)}
                                _ = v18
                            case "a":
                                if false {
                                    fmt.Println(fi12, v11)
                                    var v19 int = fi15
                                    if ((fi12 * v19) > fi12) {
                                        v1[594] = fi15
                                        var v20 float32 = float32((float64(-2.25) - v11))
                                        _ = v20
                                    } else {
                                        v21 := v1[542]
                                        switch (fi12 + (v21 << ((0) & 63))) {
                                        case -1:
                                            v19 = -9223372036854775808
                                        case 42:
                                            delete(v2, v21)
                                        case 2147483647:
                                            v22 := GenS1{F0: 1, F1: -9223372036854775808, F2: (((^(fi15)) > v19) || (!false))}
                                            _ = v22
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = v19
                                        }
                                        _ = v21
                                    }
                                    _ = v19
                                } else {
                                    v11 += (1.5 + v11)
                                    _ = len(v5)
                                }
                            default:
                                if (2 == 1) {
                                    v23 := GenS2{F0: 9223372036854775807}
                                    v24 := GenS1{F0: -1, F1: -586, F2: true}
                                    _ = v23
                                    _ = v24
                                } else {
                                    v25 := GenS2{F0: 9223372036854775807}
                                    v11 = v11
                                    _ = len(v5)
                                    _ = v25
                                }
                            }
                            _ = len(v5)
                            var v26 uint64 = 18446744073709551615
                            _ = v26
                        }
                    }
                case "a":
                    if false {
                        v11 -= v3
                        switch "a" {
                        case "hello":
                            v27 := GenS2{F0: 1}
                            _ = v27
                        case "\n\t":
                            var v28 uint16 = 2
                            _ = v28
                        case "":
                            v11 = v3
                        default:
                            if (!false) {
                                fmt.Println(v11, fi12)
                                v5 = append(v5, (fi15 ^ fi15))
                            }
                        }
                    } else {
                        var v29 uint16 = 957
                        v3 = -2.25
                        _ = v29
                    }
                    var v30 uint32 = 65535
                    _ = v30
                default:
                    var v31 uint8 = 234
                    _ = v31
                }
                var v32 int32 = -127
                _ = fi15
                _ = v32
            }
            switch fi12 {
            case -1:
                for fi33 := 0; fi33 < 1; fi33++ {
                    var v34 int16 = -32768
                    if ((v3 > v3) && true) {
                        if len(v5) > 0 {
                            v5[0] = 128
                        }
                        v35 := v2[(fi33 + fi12)]
                        _ = len(v5)
                        _ = v35
                    } else {
                        _ = len(v5)
                        v4 = (1.5 - (v4 - float32((3.14159 / v11))))
                    }
                    _ = fi33
                    _ = v34
                }
            }
            _ = fi12
        }
    } else {
        var v36 uint8 = 0
        v2[32768] = 128
        _ = v36
    }
    v3 = float64(0.0)
    v11 = v11
    var v37 bool = false
    if (-2147483648 >= 0) {
        v4 = (-2.25 - (v4 / v4))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
