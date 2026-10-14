package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    v4 := []int{-2147483648}
    var v5 bool = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = 958
    }
    if v5 {
        if (255 <= 1) {
            var v6 int32 = 150
            for fi7 := 0; fi7 < 3; fi7++ {
                for fi8 := 0; fi8 < 18; fi8++ {
                    if v5 {
                        if v5 {
                            fmt.Println(v6)
                            v5 = (255 <= 255)
                            if v5 {
                                v9 := GenS1{F0: 159, F1: 391}
                                _ = v9
                            }
                        } else {
                            v5 = (1 < 1)
                            v6 = (v6 % (((v6 % ((v6) | 1))) | 1))
                        }
                    }
                    var v10 int16 = 1000
                    switch "\n\t" {
                    case "hello":
                        var v11 uint64 = 1
                        _ = v11
                    case "a":
                        switch fi7 {
                        case -418:
                            v12 := GenS1{F0: 255, F1: (v6 << ((255) & 63))}
                            _ = v12
                        default:
                            delete(v1, (fi8 * fi7))
                        }
                    default:
                        var v13 float64 = (-2.25 - ((0.1 + (0.1 - float64(3.14159))) + 0.0))
                        _ = v13
                    }
                    _ = fi8
                    _ = v10
                }
                if ((v5 || v5) && (-2.25 < (-2.25 - 3.14159))) {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v6 = v6
                        _ = fi14
                    }
                    v6 += (v6 + (^(v6)))
                }
                _ = fi7
            }
            _ = v6
        }
    }
    v15 := GenS1{F0: 128, F1: 0}
    v3[""] = 2147483647
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            if (42 > 42) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS2{F0: (^(fi17))}
                    _ = len(v1)
                    var v19 float32 = 0.0
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                switch "\n\t" {
                case "foo bar":
                    var v20 int16 = 1
                    _ = v20
                case "\n\t":
                    delete(v1, (-(fi16)))
                case "":
                    _ = len(v2)
                default:
                    switch fi16 {
                    case -1:
                        var v21 float32 = 0.1
                        _ = v21
                    }
                }
            } else {
                var v22 float32 = 1e10
                switch "" {
                case "\n\t":
                    v5 = ((2 < 121) && false)
                default:
                    v23 := v3["\n\t"]
                    _ = v23
                }
                _ = v22
            }
            v5 = false
            var v24 uint8 = 125
            _ = v24
        } else {
            if v5 {
                v25 := v2[(fi16 * fi16)]
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 uint = 42
                    v28 := GenS3{F0: 127, F1: int64(fi16), F2: 0}
                    v29 := GenS3{F0: 127, F1: 1, F2: 255}
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                if v5 {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        var v31 int32 = -1
                        v25 = (fi16 << ((65535) & 63))
                        _ = fi30
                        _ = v31
                    }
                    var v32 int32 = 42
                    _ = v32
                }
                _ = v25
            }
            var v33 int16 = 32767
            _ = v33
        }
        if v5 {
            v5 = ((!(612 >= 32767)) || v5)
        }
        _ = fi16
    }
    var v34 int = -497
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
