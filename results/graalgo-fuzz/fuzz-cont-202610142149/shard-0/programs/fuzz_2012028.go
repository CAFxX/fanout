package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]int)
    var v3 float32 = 0.0
    var v4 int32 = 0
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(v4)
        var v6 int16 = 1000
        if ((v1 | v1) == v1) {
            var v7 uint16 = 260
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    var v8 float64 = 3.14159
    v1 = v1
    v9 := GenS1{F0: "hello", F1: (false || (false && false)), F2: v1}
    v10 := GenS2{F0: 32767, F1: 18446744073709551615}
    for fi11 := 0; fi11 < 14; fi11++ {
        if (false && (true || (109 > 127))) {
            v8 = 3.14159
            if false {
                if false {
                    if true {
                        for fi12 := 0; fi12 < 12; fi12++ {
                            var v13 int32 = (v4 % (((^(v4))) | 1))
                            if (!(fi11 == (fi12 << ((0) & 63)))) {
                                v14 := GenS1{F0: "\n\t", F1: (!(((-186 != -957) && (!true)) || (((0.1 / (1e10 / -2.25)) * v3) <= 0.0))), F2: v1}
                                if (true || false) {
                                    var v15 uint64 = v1
                                    var v16 int64 = 600
                                    _ = v15
                                    _ = v16
                                } else {
                                    v17 := v2["a"]
                                    _ = v17
                                }
                                _ = v14
                            } else {
                                _ = len(v2)
                                v18 := GenS3{F0: int16(v1)}
                                _ = v18
                            }
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        v1 = 42
                        var v19 uint64 = 132
                        var v20 int = fi11
                        _ = v19
                        _ = v20
                    }
                } else {
                    v1 = 0
                }
                _ = len(v2)
            }
            if (543 > 584) {
                v8 -= float64(v3)
                fmt.Println(v4)
            } else {
                v8 = (v8 - float64(v3))
                v21 := GenS4{F0: 652}
                if (4294967295 >= 4294967295) {
                    var v22 string = "\n\t"
                    _ = v22
                } else {
                    var v23 uint = 255
                    v4 -= v4
                    _ = v23
                }
                _ = v21
            }
        }
        _ = fi11
    }
    v1 = (v1 & (v1 % (((v1 & 1)) | 1)))
    for fi24 := 0; fi24 < 16; fi24++ {
        switch fi24 {
        case 104:
            v25 := GenS5{F0: ((v8 - v8) > float64((0.1 / (0.1 * v3))))}
            _ = v25
        default:
            if (255 > 42) {
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 uint8 = 1
                    _ = fi26
                    _ = v27
                }
                var v28 int16 = -388
                _ = v28
            }
        }
        v4 = -1
        switch fi24 {
        case 1:
            var v29 float64 = -2.25
            _ = v29
        case -602:
            v4 = v4
        default:
            for fi30 := 0; fi30 < 2; fi30++ {
                v3 = v3
                var v31 int = -307
                switch "hello" {
                case "":
                    var v32 uint16 = 88
                    _ = v32
                }
                _ = fi30
                _ = v31
            }
        }
        _ = fi24
    }
    v33 := []int{32767, 1}
    var v34 bool = (false || (((!(!false)) || false) || (75 <= 199)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v8
    _ = v9
}
