package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16; F1 int; F2 int }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
func genFn0(p0 uint64, p1 uint64) uint {
    if (!true) {
        v5 := GenS1{F0: p0, F1: 1}
        _ = v5
    }
    p1 = p1
    _ = p0
    _ = p1
    return 137
}

func main() {
    var v1 string = ""
    v2 := make(map[int]int)
    var v3 uint16 = 0
    var v4 uint16 = (v3 | (v3 ^ v3))
    if (-2147483648 != 2147483647) {
        var v6 int8 = 127
        if true {
            v7 := []int{}
            _ = v7
        }
        _ = v6
    }
    v3 = v3
    for fi8 := 0; fi8 < 5; fi8++ {
        v1 = ((v1 + "foo bar") + "foo bar")
        if false {
            switch "hello" {
            case "\n\t":
                var v9 float32 = 0.1
                _ = v9
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 uint = 991
                v11 += 712
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 7; fi12++ {
                switch (^(fi8)) {
                case -806:
                    if ((1000 >= int16(v4)) || true) {
                        var v13 uint32 = 1
                        v1 = "hello"
                        fmt.Println(v13, v3)
                        _ = v13
                    } else {
                        if (101 > 42) {
                            v14 := v2[127]
                            v15 := GenS1{F0: 0, F1: (fi8 - fi12)}
                            if true {
                                v1 = v1
                            }
                            _ = v14
                            _ = v15
                        }
                        v16 := []int{(^(fi12)), 32767, (fi12 + fi12), (fi8 ^ fi12)}
                        _ = v16
                    }
                case 711:
                    delete(v2, (fi8 / ((fi8) | 1)))
                case 128:
                    fmt.Println(v3, v3)
                default:
                    v17 := GenS2{F0: (0.0 != (0.0 / (-2.25 + float64(0.1)))), F1: (!false)}
                    _ = v17
                }
                if true {
                    if (42 != 0) {
                        v18 := []int{}
                        var v19 uint = genFn0(4294967295, 42)
                        _ = v18
                        _ = v19
                    }
                }
                switch 128 {
                case 2147483647:
                    fmt.Println(fi12, v1)
                case 806:
                    var v20 uint32 = 2147483648
                    _ = v20
                case -2147483648:
                    fmt.Println(v4, v3)
                default:
                    if false {
                        var v21 int = (-(fi8))
                        var v22 int64 = -82
                        _ = v21
                        _ = v22
                    } else {
                        v23 := []int{(fi8 * (fi8 | (fi12 >> ((2) & 63)))), (fi8 ^ (fi8 % (((^(fi8))) | 1))), fi12, (fi8 | fi8)}
                        switch "foo bar" {
                        case "a":
                            v3 += (v4 / (((v4 + v3)) | 1))
                        case "hello":
                            var v24 int = 9223372036854775807
                            _ = v24
                        case "foo bar":
                            var v25 string = v1
                            _ = v25
                        default:
                            v26 := v2[fi12]
                            _ = v26
                        }
                        _ = v23
                    }
                }
                switch v1 {
                case "a":
                    if (false && (1e10 > float32((float64(0.0) - (1.5 - 1e10))))) {
                        var v27 int32 = 2147483647
                        var v28 int8 = -128
                        if false {
                            if (!((v1 == v1) && false)) {
                                fmt.Println(v4)
                                v29 := GenS2{F0: (false && false), F1: (!true)}
                                _ = v29
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v30 float32 = 3.14159
                                _ = v30
                            case "hello":
                                var v31 int16 = int16(v27)
                                _ = v31
                            case "a":
                                v32 := GenS1{F0: 1, F1: (fi8 >> ((42) & 63))}
                                _ = v32
                            }
                            var v33 float32 = float32(3.14159)
                            _ = v33
                        } else {
                            v1 = "hello"
                            v3 += (v3 - v3)
                            var v34 float32 = 0.0
                            _ = len(v2)
                            _ = v34
                        }
                        _ = v27
                        _ = v28
                    }
                case "\n\t":
                    fmt.Println(v1, v4)
                default:
                    var v35 int16 = 1
                    _ = v35
                }
                v36 := []int{}
                _ = fi12
                _ = v36
            }
        } else {
            v37 := GenS3{F0: 0, F1: fi8, F2: fi8}
            for fi38 := 0; fi38 < 20; fi38++ {
                v1 = "foo bar"
                _ = fi38
            }
            _ = v37
        }
        _ = fi8
    }
    if false {
        for fi39 := 0; fi39 < 19; fi39++ {
            _ = len(v2)
            _ = fi39
        }
    }
    fmt.Println(v1, v3)
    for fi40 := 0; fi40 < 19; fi40++ {
        fmt.Println(v1, v3)
        v41 := GenS2{F0: false, F1: (!true)}
        _ = fi40
        _ = v41
    }
    var v42 uint = 567
    v3 = (v3 & v3)
    v3 = v4
    if true {
        v43 := GenS4{F0: -1, F1: 166, F2: 772}
        var v44 uint16 = (v4 % ((v4) | 1))
        var v45 uint = v42
        _ = v43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
