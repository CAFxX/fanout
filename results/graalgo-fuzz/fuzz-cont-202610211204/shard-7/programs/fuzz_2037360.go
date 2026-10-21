package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint16 = 27
    var v2 string = "a"
    var v3 int64 = 0
    var v4 int = 2147483647
    v5 := make(map[int]int)
    v6 := GenS1{F0: -2147483648}
    v1 += v1
    var v7 int8 = 1
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := v5[(v4 & 9223372036854775807)]
        _ = fi8
        _ = v9
    }
    v4 += -2147483648
    v10 := []int{}
    switch v4 {
    case 42:
        v1 += (v1 % (((^(v1))) | 1))
    }
    var v11 uint8 = 145
    if (580 <= 519) {
        var v12 bool = false
        if len(v10) > 0 {
            v10[0] = v4
        }
        _ = len(v10)
        _ = v12
    } else {
        var v13 int = v4
        for fi14 := 0; fi14 < 8; fi14++ {
            if (!(("a" < v2) && (((v1 & (v1 / (((v1 & 2)) | 1))) == 32767) && false))) {
                v5[321] = (fi14 ^ fi14)
                v11 = 42
            }
            var v15 float64 = (0.0 - ((-2.25 / -2.25) - 1.5))
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    if ((2 < 912) || true) {
        v16 := GenS2{F0: (v11 + v11), F1: v11}
        var v17 int8 = 11
        if ("a" >= v2) {
            fmt.Println(v2)
            var v18 uint16 = 65535
            _ = v18
        }
        _ = v16
        _ = v17
    } else {
        if (((!true) && (-1 >= 1)) && ((!true) && (v7 <= 1))) {
            if true {
                if true {
                    var v19 uint = 18446744073709551615
                    _ = v19
                } else {
                    var v20 uint64 = 267
                    for fi21 := 0; fi21 < 6; fi21++ {
                        if true {
                            for fi22 := 0; fi22 < 18; fi22++ {
                                _ = len(v10)
                                v11 += (v11 >> ((1) & 63))
                                _ = fi22
                            }
                            var v23 int32 = 950
                            v1 -= (^(v1))
                            _ = v23
                        } else {
                            v24 := GenS3{F0: v11, F1: (-(v4)), F2: (v7 >> ((42) & 63))}
                            fmt.Println(fi21, fi21, v1)
                            for fi25 := 0; fi25 < 3; fi25++ {
                                if false {
                                    _ = len(v5)
                                    v10 = append(v10, 1)
                                    var v26 string = "foo bar"
                                    _ = v26
                                } else {
                                    fmt.Println(fi25)
                                }
                                v1 = (v1 | v1)
                                v27 := GenS4{F0: 1000, F1: -704}
                                _ = fi25
                                _ = v27
                            }
                            _ = v24
                        }
                        var v28 uint = 1
                        _ = fi21
                        _ = v28
                    }
                    v29 := GenS5{F0: 0, F1: v3}
                    _ = v20
                    _ = v29
                }
            }
            if (!true) {
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 float32 = 1e10
                    v11 = (v11 % (((v11 & 0)) | 1))
                    _ = fi30
                    _ = v31
                }
                v3 -= v3
                v1 = (^(v1))
            }
            v32 := GenS3{F0: (^(v11)), F1: (v4 + 4294967295), F2: -1}
            _ = v32
        }
        for fi33 := 0; fi33 < 12; fi33++ {
            v34 := GenS1{F0: 1}
            v5[797] = (fi33 + 42)
            v2 = v2
            _ = fi33
            _ = v34
        }
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    if false {
        v35 := GenS3{F0: (v11 * v11), F1: 4294967295, F2: (v7 - (v7 + (v7 | -26)))}
        _ = v35
    } else {
        for fi36 := 0; fi36 < 16; fi36++ {
            v37 := GenS1{F0: -548}
            _ = fi36
            _ = v37
        }
    }
    v38 := GenS2{F0: v11, F1: (v11 % ((v11) | 1))}
    var v39 string = v2
    fmt.Println(v1, v11)
    v1 -= (v1 | v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
