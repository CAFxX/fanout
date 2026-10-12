package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 uint8 }
func genFn0(p0 uint64, p1 int16, p2 uint32) int8 {
    var v7 uint32 = (p2 - p2)
    v8 := GenS1{F0: 2}
    var v9 uint16 = 154
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -115
}

func main() {
    var v1 float64 = float64(((((0.0 * 0.1) + 1.5) * 0.0) * 1.5))
    var v2 int32 = -1
    v3 := make(map[int]string)
    v4 := []uint32{2}
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    switch "a" {
    case "\n\t":
        var v10 int8 = genFn0(4294967295, 42, 2147483648)
        _ = v10
    case "":
        for fi11 := 0; fi11 < 13; fi11++ {
            v1 = 0.0
            _ = fi11
        }
    case "a":
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 10; fi13++ {
                v1 += 1e10
                _ = fi13
            }
            for fi14 := 0; fi14 < 13; fi14++ {
                v1 = float64(0.1)
                _ = fi14
            }
            v2 = (v2 ^ -147)
            _ = fi12
        }
    }
    _ = len(v4)
    var v15 int16 = -1
    var v16 int64 = 42
    var v17 string = "foo bar"
    v17 = v17
    v18 := GenS1{F0: 42}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2)
    if true {
        var v19 float32 = 0.1
        var v20 int8 = 1
        _ = v19
        _ = v20
    } else {
        v2 += (^(v2))
    }
    var v21 int8 = genFn0(2, (v15 | (v15 + v15)), 433)
    v22 := GenS1{F0: 1}
    var v23 int64 = int64(v15)
    if false {
        var v24 int16 = 0
        if true {
            var v25 uint64 = 706
            _ = v25
        } else {
            v1 -= 1e10
            v24 += v15
        }
        if (!(1 == 42)) {
            var v26 uint32 = 705
            switch (v17 + "foo bar") {
            case "":
                switch v17 {
                case "foo bar":
                    v17 = v17
                case "a":
                    switch 1 {
                    case -9223372036854775808:
                        var v27 int64 = v16
                        _ = v27
                    case 280:
                        for fi28 := 0; fi28 < 9; fi28++ {
                            v21 += -128
                            v2 = (^(v2))
                            if (((0.1 - float32((v1 - v1))) + -2.25) > 1e10) {
                                switch 0 {
                                case 1:
                                    if true {
                                        for fi29 := 0; fi29 < 12; fi29++ {
                                            v15 = (v24 % ((0) | 1))
                                            v2 += (v2 << ((255) & 63))
                                            if len(v4) > 0 {
                                                v4[0] = (^(v26))
                                            }
                                            _ = fi29
                                        }
                                        v30 := GenS1{F0: 42}
                                        fmt.Println(v24)
                                        _ = v30
                                    }
                                case -756:
                                    v4 = append(v4, v26)
                                }
                            } else {
                                fmt.Println(v1)
                            }
                            _ = fi28
                        }
                    default:
                        var v31 bool = (false && true)
                        _ = v31
                    }
                case "\n\t":
                    v32 := GenS1{F0: uint16(v26)}
                    _ = v32
                }
            case "a":
                v17 = "a"
            default:
                if true {
                    if len(v4) > 0 {
                        v4[0] = (v26 % ((654) | 1))
                    }
                    v23 -= v16
                    v24 = (v15 | v24)
                } else {
                    fmt.Println(v24, v26)
                }
            }
            if true {
                v33 := GenS2{F0: (v21 ^ (v21 << ((290) & 63))), F1: true, F2: 1}
                _ = len(v4)
                if false {
                    if false {
                        v24 -= v15
                    } else {
                        v34 := GenS2{F0: 0, F1: (!(v23 >= int64(v2))), F2: 128}
                        var v35 uint8 = 124
                        v26 += v26
                        _ = v34
                        _ = v35
                    }
                    var v36 int16 = -334
                    _ = v36
                } else {
                    fmt.Println(v15, v24, v23)
                    _ = len(v3)
                    v17 = ""
                }
                _ = v33
            }
            _ = v26
        }
        _ = v24
    } else {
        fmt.Println(v1, v2, v2)
        fmt.Println(v23, v21, v16)
    }
    fmt.Println(v21, v17)
    fmt.Println(v2, v1)
    switch 42 {
    case 128:
        v1 += v1
    case -1:
        v37 := GenS1{F0: 255}
        _ = v37
    case 32768:
        if (1 <= 170) {
            var v38 uint16 = 32767
            _ = v38
        } else {
            var v39 int16 = v15
            _ = len(v6)
            _ = v39
        }
    }
    var v40 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
