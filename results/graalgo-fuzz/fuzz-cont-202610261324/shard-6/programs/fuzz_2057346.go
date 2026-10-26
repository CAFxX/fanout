package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]int)
    fmt.Println(v3)
    var v5 uint64 = 42
    var v6 string = ("" + "")
    v6 = ("foo bar" + "foo bar")
    v5 += v5
    if (!(-32768 == 42)) {
        v5 = 42
        if ((true && true) && true) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v5 = 18446744073709551615
                v8 := []int{}
                if false {
                    switch "a" {
                    case "\n\t":
                        if (1 != 32767) {
                            if (true && (!((((uint32(v5) > 2) && true) && (-2147483648 != (fi7 % (((v3 ^ (fi7 + (fi7 | (-(v3)))))) | 1)))) && (false && false)))) {
                                var v9 int8 = 1
                                v9 -= v9
                                v2["a"] = 1
                                _ = v9
                            } else {
                                var v10 uint16 = 406
                                var v11 uint32 = 522
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                _ = v10
                                _ = v11
                            }
                            fmt.Println(v6)
                        } else {
                            if len(v8) > 0 {
                                v8[0] = 9223372036854775807
                            }
                        }
                    default:
                        v5 = (v5 % ((1) | 1))
                    }
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v13 := GenS1{F0: 127, F1: 223, F2: v5}
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        var v14 float32 = float32(3.14159)
                        var v15 uint = 429
                        switch v3 {
                        case 1:
                            v14 = 1.5
                        case 128:
                            v5 = 584
                        default:
                            v5 = v5
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = len(v4)
                } else {
                    fmt.Println(v3)
                    var v16 int32 = -2147483648
                    _ = len(v2)
                    _ = v16
                }
                _ = fi7
                _ = v8
            }
            var v17 float64 = -2.25
            if (((float32((0.1 * v17)) == (0.1 * 1e10)) && false) || false) {
                v6 = (v6 + "hello")
            }
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 2; fi18++ {
            if false {
                var v19 int16 = 541
                v5 = 2
                _ = v19
            } else {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v4[2147483647] = (^(v3))
                    v5 += (v5 >> ((2) & 63))
                    _ = fi20
                }
                v21 := GenS2{F0: (fi18 * (fi18 & 895)), F1: ((104 < 179) || (true && false))}
                _ = v21
            }
            var v22 uint32 = 2147483648
            _ = fi18
            _ = v22
        }
        var v23 int = v3
        _ = v23
    }
    if (0 >= 0) {
        var v24 int16 = 2
        _ = v24
    } else {
        fmt.Println(v5, v3)
        switch v6 {
        case "":
            var v25 int16 = -32768
            _ = v25
        case "a":
            fmt.Println(v6, v6, v5)
        case "foo bar":
            var v26 int = (v3 * v3)
            _ = v26
        }
    }
    if false {
        v27 := []int{-1, v3}
        _ = v27
    } else {
        v6 = "\n\t"
        for fi28 := 0; fi28 < 2; fi28++ {
            if true {
                var v29 float32 = float32(0.0)
                v29 -= float32((0.0 / 3.14159))
                _ = v29
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                switch (fi28 / ((v3) | 1)) {
                case -56:
                    v3 = (fi28 | int(v5))
                case -9223372036854775808:
                    v31 := GenS3{F0: 0}
                    _ = v31
                default:
                    if true {
                        var v32 uint16 = 267
                        _ = v32
                    } else {
                        _ = len(v1)
                        v33 := v1[65535]
                        fmt.Println(v33, v33, v33)
                        _ = v33
                    }
                }
                _ = fi30
            }
            var v34 uint8 = 2
            _ = fi28
            _ = v34
        }
    }
    var v35 float64 = 1.5
    var v36 int64 = 42
    v3 -= v3
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 uint32 = 2147483647
        _ = fi37
        _ = v38
    }
    var v39 uint8 = 225
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
