package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int = 128
    v2 := make(map[int]bool)
    var v3 int32 = -2147483648
    var v4 int16 = 820
    v5 := []uint32{}
    v3 = (^(v3))
    for fi6 := 0; fi6 < 13; fi6++ {
        if (!true) {
            for fi7 := 0; fi7 < 10; fi7++ {
                _ = len(v2)
                _ = fi7
            }
            v8 := GenS1{F0: ((1.5 > (float32(1e10) * 1e10)) || true), F1: 0}
            if (false || false) {
                var v9 float32 = 0.1
                for fi10 := 0; fi10 < 8; fi10++ {
                    v1 = (fi6 >> ((0) & 63))
                    fmt.Println(v3, v4)
                    v9 = (0.1 * 0.0)
                    _ = fi10
                }
                v11 := GenS2{F0: v4}
                _ = v11
                _ = v9
            }
            _ = v8
        } else {
            if true {
                v12 := GenS3{F0: -92, F1: 2}
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 bool = (false || (0.1 > 1e10))
                    v4 = 0
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                if (v4 != v4) {
                    v3 -= (v3 + (v3 & 1))
                    for fi15 := 0; fi15 < 14; fi15++ {
                        if (!(539 == 4294967295)) {
                            switch "a" {
                            case "a":
                                v16 := GenS1{F0: false, F1: 41}
                                _ = v16
                            default:
                                if false {
                                    switch ("hello" + "hello") {
                                    case "":
                                        v17 := v2[(fi6 | (fi15 << ((2) & 63)))]
                                        _ = v17
                                    }
                                    v3 = (v3 & (-(v3)))
                                    var v18 bool = (fi15 > (-(fi6)))
                                    _ = v18
                                } else {
                                    if false {
                                        var v19 int64 = 0
                                        v20 := GenS4{F0: 1, F1: v3}
                                        switch ("foo bar" + "hello") {
                                        case "hello":
                                            v21 := GenS3{F0: int8(fi15), F1: 0}
                                            _ = v21
                                        case "a":
                                            if (((42 <= 255) || (2 < 0)) || false) {
                                                var v22 int32 = v3
                                                fmt.Println(v1, v3, v4)
                                                switch fi6 {
                                                case 32768:
                                                    switch fi15 {
                                                    case 2147483647:
                                                        var v23 string = (((("a" + ("\n\t" + "\n\t")) + "a") + ("" + "foo bar")) + "\n\t")
                                                        _ = v23
                                                    }
                                                }
                                                _ = v22
                                            }
                                        default:
                                            switch v1 {
                                            case 32768:
                                                var v24 uint = 1
                                                _ = v24
                                            default:
                                                fmt.Println(fi6)
                                            }
                                        }
                                        _ = v19
                                        _ = v20
                                    }
                                    fmt.Println(v4, v1)
                                    if false {
                                        v25 := GenS3{F0: 127, F1: 42}
                                        var v26 uint64 = 332
                                        _ = v25
                                        _ = v26
                                    }
                                    v5 = append(v5, 0)
                                }
                            }
                            switch ("a" + "a") {
                            case "hello":
                                var v27 uint16 = 484
                                _ = v27
                            case "":
                                var v28 string = "hello"
                                _ = v28
                            }
                            if false {
                                fmt.Println(v1, v3)
                            }
                        }
                        _ = fi15
                    }
                    v1 += (v1 >> ((1) & 63))
                }
                var v29 int64 = -854
                _ = v29
            }
        }
        v4 += -1
        _ = fi6
    }
    v30 := GenS4{F0: 42, F1: (v3 >> ((2) & 63))}
    var v31 int32 = (^(v3))
    var v32 float64 = (((float64(1.5) + -2.25) * 3.14159) - ((3.14159 + 1.5) + 1.5))
    var v33 uint = 1
    var v34 int32 = v31
    fmt.Println(v34)
    v3 -= 1
    v35 := v2[v1]
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v36 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
