package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int64; F2 int8 }
func genFn0(p0 bool, p1 int8) uint {
    for fi6 := 0; fi6 < 9; fi6++ {
        switch "\n\t" {
        case "a":
            p0 = p0
        case "hello":
            switch "hello" {
            case "":
                switch ("\n\t" + ("foo bar" + ("a" + "foo bar"))) {
                case "":
                    switch 0 {
                    case 127:
                        fmt.Println(p0)
                    case -1:
                        for fi7 := 0; fi7 < 2; fi7++ {
                            v8 := []int{}
                            v9 := GenS1{F0: 42}
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                    default:
                        var v10 int16 = -324
                        _ = v10
                    }
                default:
                    v11 := GenS2{F0: 42}
                    _ = v11
                }
            case "\n\t":
                v12 := []int{}
                _ = v12
            default:
                var v13 uint8 = 240
                _ = v13
            }
        default:
            var v14 string = "foo bar"
            _ = v14
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            if (((3.14159 != 3.14159) && false) && false) {
                fmt.Println(fi15)
                if (42 > 0) {
                    v16 := []int{(^(fi15)), (fi15 << ((638) & 63)), fi6}
                    if p0 {
                        v16 = append(v16, 32767)
                        var v17 float32 = ((1.5 * float32((1.5 / (0.0 + ((float64((1e10 * 1e10)) / float64(1e10)) + 1e10))))) / (1e10 - 1.5))
                        var v18 uint64 = 2
                        _ = v17
                        _ = v18
                    }
                    if ((127 == 63) || p0) {
                        fmt.Println(p1, p0)
                        p1 += (p1 << ((1) & 63))
                    } else {
                        p0 = p0
                    }
                    v19 := GenS3{F0: "", F1: 0, F2: p1}
                    fmt.Println(fi15)
                    _ = v16
                    _ = v19
                }
                var v20 float64 = (1e10 + (float64(1.5) / 1.0))
                _ = v20
            } else {
                var v21 uint8 = 1
                _ = v21
            }
            _ = fi15
        }
        _ = fi6
    }
    var v22 uint16 = 65535
    var v23 uint64 = 0
    if (4294967295 < 42) {
        switch "" {
        case "":
            if p0 {
                v24 := []int{0}
                fmt.Println(v23)
                if (float64(-2.25) == 1e10) {
                    v25 := GenS2{F0: 1}
                    _ = v25
                }
                _ = v24
            } else {
                v23 += v23
            }
        }
    } else {
        var v26 int16 = 181
        var v27 uint = 405
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 473
}
func genFn1(p0 uint32, p1 uint64, p2 uint64) int32 {
    if ((true || false) || false) {
        v28 := []int{4294967295, 83, -117, 378}
        _ = v28
    } else {
        v29 := []int{194, 2147483647, 0}
        var v30 string = "\n\t"
        switch "" {
        case "a":
            var v31 float64 = (0.1 + (0.0 - float64(3.14159)))
            _ = v31
        case "hello":
            fmt.Println(p1, p0)
        case "foo bar":
            if (!(true && true)) {
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 float64 = (((0.1 - ((1.5 / (1e10 + 1e10)) * float64(1e10))) + float64(0.0)) / 1.5)
                    var v34 string = v30
                    for fi35 := 0; fi35 < 7; fi35++ {
                        v36 := GenS2{F0: 42}
                        _ = fi35
                        _ = v36
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                if ((true || (!(1.5 == float32(0.1)))) && (-1 > -52)) {
                    v37 := GenS2{F0: -949}
                    _ = v37
                }
            }
        }
        _ = v29
        _ = v30
    }
    v38 := []int{-2147483648, 128}
    fmt.Println(p2, p2, p2)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    return 340
}

func main() {
    var v1 float64 = (0.0 - 1.5)
    v2 := make(map[string]int)
    v3 := []int{-1, 886}
    v4 := make(map[int]int)
    var v5 int8 = 55
    _ = len(v3)
    var v39 float64 = (((1.5 + 0.0) + v1) + v1)
    var v40 int64 = -1
    v5 -= -128
    v39 += ((v39 - v39) + 0.0)
    var v41 int = 127
    var v42 uint8 = 255
    var v43 int32 = 0
    var v44 uint8 = (v42 + 129)
    fmt.Println(v44)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
