package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int8 }
type GenS4 struct { F0 int32; F1 int64; F2 int32 }
func genFn0() int {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{}
        var v7 uint = 65535
        if (!true) {
            switch fi5 {
            case -2147483648:
                switch "hello" {
                case "":
                    for fi8 := 0; fi8 < 14; fi8++ {
                        var v9 int64 = -1
                        _ = fi8
                        _ = v9
                    }
                case "foo bar":
                    var v10 uint8 = uint8(v7)
                    _ = v10
                }
            case -9223372036854775808:
                switch ("" + "hello") {
                case "\n\t":
                    v11 := GenS1{F0: int32(fi5)}
                    _ = v11
                case "hello":
                    v6 = append(v6, (fi5 % ((2147483647) | 1)))
                case "foo bar":
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if (!(((!true) || true) || false)) {
                            fmt.Println(fi12, v7, fi12)
                            switch "hello" {
                            case "hello":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            case "":
                                switch fi5 {
                                case -1:
                                    var v13 int8 = 1
                                    _ = v13
                                case 32767:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                case -9223372036854775808:
                                    var v14 string = ((("hello" + "") + "") + "")
                                    _ = v14
                                }
                            default:
                                var v15 int32 = 42
                                v7 -= (v7 - v7)
                                _ = v15
                            }
                            v16 := GenS2{F0: (^(fi12)), F1: 0, F2: (^(v7))}
                            var v17 int64 = 615
                            v18 := GenS3{F0: (v7 << ((v7) & 63)), F1: 749, F2: -1}
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        _ = fi12
                    }
                default:
                    var v19 uint32 = 1
                    _ = v19
                }
            case 32767:
                v7 -= 18446744073709551615
            default:
                fmt.Println(v7, fi5)
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                fmt.Println(v7, fi5, fi20)
                if false {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    var v21 float64 = 0.1
                    _ = v21
                }
                v7 -= (v7 / ((708) | 1))
                _ = fi20
            }
            v22 := GenS4{F0: 210, F1: -874, F2: 971}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 float32 = (0.1 / float32((((3.14159 + 0.0) * -2.25) + (0.1 - (3.14159 - (1e10 - ((float64(((1e10 + 1e10) * 1e10)) - 1e10) / 1e10)))))))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v6 = append(v6, (fi23 | -695))
                _ = fi23
                _ = v24
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v25 := GenS2{F0: 32768, F1: 21, F2: 433}
    var v26 uint8 = 1
    v27 := []int{1, -9223372036854775808}
    _ = v25
    _ = v26
    _ = v27
    return 128
}

func main() {
    var v1 uint = 0
    v2 := []int{870, 62}
    var v3 int64 = -236
    v4 := make(map[int]uint32)
    var v28 bool = false
    v28 = ((v1 << (((^(v1))) & 63)) <= 2)
    v1 = v1
    for fi29 := 0; fi29 < 5; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 int64 = (v3 << ((314) & 63))
            delete(v4, genFn0())
            _ = fi30
            _ = v31
        }
        var v32 float32 = 0.1
        _ = fi29
        _ = v32
    }
    switch 128 {
    case 42:
        v33 := GenS1{F0: 0}
        _ = v33
    }
    v34 := v4[397]
    var v35 uint16 = 0
    if false {
        _ = len(v2)
        var v36 int8 = -107
        var v37 int = -1
        _ = v36
        _ = v37
    } else {
        _ = len(v2)
        for fi38 := 0; fi38 < 9; fi38++ {
            v3 -= 42
            var v39 int16 = 32767
            var v40 uint8 = 68
            _ = fi38
            _ = v39
            _ = v40
        }
        v28 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
