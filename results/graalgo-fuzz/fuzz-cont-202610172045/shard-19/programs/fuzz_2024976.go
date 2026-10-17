package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint8, p2 uint32) uint8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        p1 -= 42
        _ = fi5
    }
    fmt.Println(p2, p1, p0)
    if (false || false) {
        v6 := GenS1{F0: 736, F1: 19, F2: p1}
        _ = v6
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    return (p0 << ((192) & 63))
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 1
    var v3 int16 = 0
    v4 := []int{127, 4294967295, -2147483648, 42}
    v7 := GenS1{F0: 2, F1: 0, F2: v2}
    _ = len(v4)
    v1 -= (-(v1))
    v3 -= (v3 / ((v3) | 1))
    v2 = genFn0(v2, v2, 967)
    var v8 int8 = 127
    v8 = v8
    fmt.Println(v2, v2)
    var v9 uint32 = 894
    switch int(v2) {
    case -1:
        for fi10 := 0; fi10 < 12; fi10++ {
            switch "a" {
            case "\n\t":
                v2 = v2
            case "a":
                var v11 float64 = (0.0 - (1e10 - float64(1e10)))
                _ = v11
            case "":
                var v12 int8 = (^(v8))
                _ = v12
            default:
                v3 = (v3 ^ (v3 ^ v3))
            }
            for fi13 := 0; fi13 < 7; fi13++ {
                if false {
                    var v14 uint = 18446744073709551615
                    var v15 float32 = 0.0
                    fmt.Println(v2, v14, v14)
                    var v16 uint64 = 42
                    var v17 int8 = (v8 + int8(v9))
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            _ = len(v4)
                        case "hello":
                            v2 += v2
                        case "a":
                            switch (fi10 + fi10) {
                            case 778:
                                v18 := GenS1{F0: (v16 / ((uint64(v9)) | 1)), F1: (v16 & v16), F2: (^(v2))}
                                _ = v18
                            case -255:
                                var v19 bool = false
                                _ = v19
                            default:
                                var v20 int8 = 42
                                _ = v20
                            }
                        default:
                            v4 = append(v4, -184)
                        }
                    }
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    v8 -= v8
                    switch 42 {
                    case 32768:
                        v3 -= (v3 ^ v3)
                    }
                }
                switch fi13 {
                case 32768:
                    var v21 uint16 = 42
                    if len(v4) > 0 {
                        v4[0] = fi10
                    }
                    v3 += v3
                    _ = v21
                case 128:
                    v3 += v3
                case -861:
                    var v22 int16 = v3
                    _ = v22
                default:
                    if true {
                        var v23 uint32 = 2147483648
                        v24 := GenS1{F0: 42, F1: uint64(v23), F2: v2}
                        if true {
                            v23 = (v23 | (v9 | (^(v9))))
                            v9 -= (v23 + v9)
                            _ = len(v4)
                        }
                        _ = v23
                        _ = v24
                    } else {
                        switch "" {
                        case "":
                            v25 := GenS2{F0: 115, F1: (^(fi10)), F2: (2 != 1)}
                            _ = v25
                        case "a":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch "a" {
                            case "\n\t":
                                if true {
                                    v26 := GenS1{F0: 4294967295, F1: 2, F2: genFn0((v2 - (v2 << ((0) & 63))), (v2 << ((18446744073709551615) & 63)), v9)}
                                    _ = v26
                                } else {
                                    var v27 uint16 = 517
                                    _ = v27
                                }
                            case "hello":
                                v28 := GenS1{F0: 996, F1: 42, F2: v2}
                                _ = v28
                            default:
                                v8 += (v8 / ((42) | 1))
                            }
                        case "\n\t":
                            var v29 float32 = (-2.25 - 1.5)
                            _ = v29
                        default:
                            v2 += (^(v2))
                        }
                    }
                    var v30 string = "\n\t"
                    _ = v30
                }
                _ = fi13
            }
            v31 := GenS3{F0: -1}
            _ = fi10
            _ = v31
        }
    case 916:
        fmt.Println(v8, v9)
    }
    v32 := GenS3{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
