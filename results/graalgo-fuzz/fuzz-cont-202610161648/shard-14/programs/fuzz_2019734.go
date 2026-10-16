package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 string; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint; F1 int }
func genFn0(p0 uint, p1 bool, p2 int16) int32 {
    if ((p0 >> (((p0 * (p0 * 4294967295))) & 63)) != p0) {
        p1 = (p0 > p0)
        v5 := GenS1{F0: -1, F1: p2, F2: (p0 / (((^(p0))) | 1))}
        _ = v5
    } else {
        fmt.Println(p1)
        for fi6 := 0; fi6 < 8; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                if (((-2.25 < (-2.25 * 1e10)) || false) && p1) {
                    var v8 uint = (p0 | (p0 ^ p0))
                    _ = v8
                }
                v9 := GenS2{F0: 0, F1: -1}
                _ = fi7
                _ = v9
            }
            _ = fi6
        }
        v10 := []int{127, 32767}
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        p0 = (^(p0))
        _ = fi11
    }
    var v12 int16 = 42
    v13 := []int{526, 1, 495, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 42
}
func genFn1() int8 {
    if ((!true) || (!(1000 > 0))) {
        v14 := GenS3{F0: "\n\t", F1: 1, F2: 2}
        _ = v14
    } else {
        switch "foo bar" {
        case "hello":
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS4{F0: 1, F1: 207, F2: 906}
                if true {
                    if true {
                        var v17 int = (fi15 | (-(fi15)))
                        _ = v17
                    } else {
                        var v18 uint64 = 2
                        _ = v18
                    }
                    if true {
                        v19 := []int{(fi15 & 0), 4294967295, fi15, fi15}
                        var v20 int = (fi15 >> ((2) & 63))
                        _ = v19
                        _ = v20
                    } else {
                        switch fi15 {
                        case -2147483648:
                            var v21 uint = 4294967295
                            _ = v21
                        default:
                            v22 := GenS1{F0: -1, F1: -32768, F2: 255}
                            _ = v22
                        }
                        var v23 int8 = 1
                        switch (fi15 - (^(fi15))) {
                        case 32768:
                            if true {
                                fmt.Println(fi15, fi15)
                                v23 = v23
                                v23 += v23
                            } else {
                                var v24 int16 = -1
                                v25 := GenS1{F0: -731, F1: 1, F2: 255}
                                _ = v24
                                _ = v25
                            }
                        case -673:
                            var v26 int16 = 1000
                            _ = v26
                        default:
                            var v27 uint64 = 2
                            _ = v27
                        }
                        _ = v23
                    }
                } else {
                    v28 := []int{(fi15 - fi15), fi15, fi15}
                    _ = v28
                }
                _ = fi15
                _ = v16
            }
        case "":
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 int16 = 42
                _ = fi29
                _ = v30
            }
        default:
            var v31 uint8 = 1
            _ = v31
        }
    }
    var v32 uint32 = 42
    for fi33 := 0; fi33 < 14; fi33++ {
        fmt.Println(fi33)
        _ = fi33
    }
    _ = v32
    return 0
}

func main() {
    v1 := []string{"", ("" + "\n\t")}
    v2 := make(map[int]string)
    var v3 float32 = float32(((1.5 / 0.1) * (((((-2.25 - 0.0) / 1.0) - -2.25) * (3.14159 / (1e10 - 0.1))) / 1.5)))
    var v4 int32 = -870
    switch 127 {
    case 32768:
        var v34 int8 = 1
        _ = v34
    case 926:
        if (127 == 87) {
            for fi35 := 0; fi35 < 11; fi35++ {
                var v36 int16 = 42
                _ = fi35
                _ = v36
            }
        }
    default:
        fmt.Println(v3)
    }
    var v37 float32 = v3
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 bool = true
        for fi40 := 0; fi40 < 7; fi40++ {
            delete(v2, (fi40 & (fi38 * (fi40 << ((407) & 63)))))
            fmt.Println(v37)
            var v41 uint8 = 101
            if v39 {
                var v42 int64 = 423
                _ = v42
            } else {
                var v43 uint8 = v41
                _ = v43
            }
            v41 += v41
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    var v44 int32 = v4
    v3 -= (float32(3.14159) * v37)
    v3 += 0.1
    var v45 int8 = 42
    v46 := GenS5{F0: 2, F1: 898}
    v37 += v37
    var v47 float64 = ((0.0 / (3.14159 * (1.5 / (1.5 * (1.5 + ((1.5 * 1e10) + float64(v3))))))) - 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
