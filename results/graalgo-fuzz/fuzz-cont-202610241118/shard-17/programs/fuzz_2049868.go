package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int; F1 string; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint; F2 string }
func genFn0(p0 int8) uint8 {
    fmt.Println(p0)
    v7 := []int{9223372036854775807, 42}
    switch "a" {
    case "":
        var v8 uint = 531
        _ = v8
    case "\n\t":
        v7 = append(v7, 2147483647)
    default:
        fmt.Println(p0)
    }
    v9 := GenS1{F0: ""}
    if len(v7) > 0 {
        v7[0] = 84
    }
    _ = p0
    _ = v7
    _ = v9
    return 127
}
func genFn1() uint8 {
    var v10 float64 = 0.1
    v11 := []int{9223372036854775807}
    var v12 uint = 255
    _ = v10
    _ = v11
    _ = v12
    return 255
}

func main() {
    v1 := []uint32{392}
    var v2 bool = ((true || true) || true)
    var v3 string = "\n\t"
    v4 := make(map[int]string)
    var v5 uint64 = 1
    var v6 int16 = 1
    var v13 int32 = -188
    var v14 string = ((v3 + "") + "\n\t")
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            if (1 >= 306) {
                if (-1 >= (fi16 >> ((711) & 63))) {
                    var v17 float64 = (1e10 / (1.5 * 1e10))
                    v13 = (v13 | v13)
                    if (!(2 >= 127)) {
                        var v18 float64 = v17
                        switch (fi15 % ((fi15) | 1)) {
                        case 1:
                            fmt.Println(fi16, v14, v2)
                        case 4294967295:
                            var v19 string = "\n\t"
                            _ = v19
                        default:
                            if v2 {
                                var v20 int16 = (v6 * (v6 - int16(v5)))
                                v21 := GenS2{F0: 9223372036854775807, F1: 65535}
                                var v22 float64 = float64(1.5)
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                        _ = v18
                    }
                    _ = v17
                }
            }
            v6 -= (v6 + v6)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi16
        }
        _ = fi15
    }
    v3 = v14
    if v2 {
        for fi23 := 0; fi23 < 7; fi23++ {
            for fi24 := 0; fi24 < 1; fi24++ {
                if v2 {
                    var v25 uint32 = 943
                    var v26 int16 = -32768
                    switch "" {
                    case "a":
                        var v27 int64 = 0
                        _ = v27
                    case "foo bar":
                        v28 := GenS3{F0: 42, F1: 1}
                        _ = v28
                    case "\n\t":
                        switch int(v13) {
                        case -459:
                            v6 += (v26 >> ((246) & 63))
                        default:
                            var v29 bool = v2
                            switch 249 {
                            case -2147483648:
                                v30 := GenS1{F0: "\n\t"}
                                _ = v30
                            case 42:
                                _ = len(v4)
                            }
                            var v31 bool = ((v5 << ((0) & 63)) >= 0)
                            v32 := GenS4{F0: (fi24 * fi23), F1: "", F2: 0}
                            _ = v29
                            _ = v31
                            _ = v32
                        }
                    default:
                        var v33 float64 = (0.1 + 0.1)
                        _ = v33
                    }
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
            var v34 bool = ((float32(0.1) + (3.14159 * (float32(1e10) / float32(1e10)))) == -2.25)
            _ = fi23
            _ = v34
        }
        var v35 int64 = int64(v6)
        for fi36 := 0; fi36 < 20; fi36++ {
            if (619 != 508) {
                fmt.Println(fi36, v5)
            }
            _ = fi36
        }
        _ = v35
    }
    if v2 {
        if (!true) {
            v37 := GenS3{F0: 32767, F1: 65535}
            v5 = (v5 >> ((18446744073709551615) & 63))
            v1 = append(v1, 2)
            _ = v37
        } else {
            _ = len(v1)
            var v38 uint16 = 32767
            _ = v38
        }
    }
    v39 := GenS3{F0: 127, F1: 255}
    var v40 int = 32768
    var v41 float64 = 1e10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v42 string = v3
    v43 := GenS5{F0: 2, F1: 255, F2: (v42 + v14)}
    switch v40 {
    case 1:
        var v44 uint32 = 2147483648
        _ = v44
    default:
        _ = len(v4)
    }
    v45 := GenS3{F0: v40, F1: 255}
    var v46 string = v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
