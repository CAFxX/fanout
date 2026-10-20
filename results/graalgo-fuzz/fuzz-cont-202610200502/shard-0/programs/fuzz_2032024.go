package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint32; F2 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 int32 }
type GenS4 struct { F0 int }

func main() {
    var v1 float64 = (3.14159 * 1.5)
    v2 := make(map[int]bool)
    v3 := []int{1, 42}
    v4 := []int{-349, -674, 32768, 32768}
    var v5 int64 = 9223372036854775807
    v6 := []uint32{}
    if len(v4) > 0 {
        v4[0] = 696
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        switch (fi7 ^ -592) {
        case 9223372036854775807:
            switch fi7 {
            case -729:
                v8 := GenS1{F0: false}
                _ = v8
            case 974:
                v9 := v2[(fi7 << ((1) & 63))]
                _ = v9
            case -914:
                delete(v2, (fi7 | (-(fi7))))
            default:
                if (100 == 1) {
                    v10 := GenS2{F0: fi7, F1: 1, F2: true}
                    _ = v10
                } else {
                    var v11 uint = 255
                    var v12 int64 = (v5 * -1)
                    _ = v11
                    _ = v12
                }
            }
        case -1:
            if (((v1 * 0.1) <= (v1 + (((v1 * v1) * -2.25) - ((1e10 + (0.0 / (float64(0.1) / float64(1e10)))) * v1)))) && (false || true)) {
                v1 = v1
                v5 = (v5 - (v5 << ((18446744073709551615) & 63)))
                v2[(^(fi7))] = true
            }
        case 32767:
            v13 := GenS1{F0: false}
            _ = v13
        default:
            v1 = v1
        }
        var v14 uint = 4294967295
        _ = fi7
        _ = v14
    }
    var v15 int32 = -2147483648
    v15 = v15
    v1 += v1
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int64 = (v5 - (v5 << ((4294967295) & 63)))
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(v15)
                v15 += v15
                switch fi16 {
                case 32767:
                    var v19 uint16 = 1
                    _ = v19
                case 9223372036854775807:
                    fmt.Println(fi16, v17, v5)
                case 127:
                    if true {
                        fmt.Println(fi16, fi18, fi18)
                        v17 -= -1
                    } else {
                        v6 = append(v6, 17)
                        if false {
                            var v20 int32 = (-(v15))
                            _ = v20
                        }
                        v21 := GenS1{F0: (true || true)}
                        _ = v21
                    }
                default:
                    switch "\n\t" {
                    case "":
                        v2[-1] = (true && (!false))
                        fmt.Println(v17, v5)
                        var v22 int32 = -1
                        _ = v22
                    case "\n\t":
                        v1 += (v1 * ((float64((((1e10 - 0.0) / (3.14159 / 3.14159)) + 3.14159)) / 1.5) + (1.5 + -2.25)))
                    case "foo bar":
                        if true {
                            _ = len(v3)
                            fmt.Println(v5)
                            var v23 uint8 = 255
                            _ = v23
                        } else {
                            v24 := GenS2{F0: (fi16 * (fi18 - fi16)), F1: 42, F2: false}
                            _ = v24
                        }
                    default:
                        if false {
                            v25 := GenS1{F0: (1.5 < 1e10)}
                            if (((fi18 ^ fi18) != (fi16 >> ((784) & 63))) && (817 >= 0)) {
                                var v26 int8 = 42
                                var v27 int16 = 1000
                                var v28 uint8 = 255
                                _ = v26
                                _ = v27
                                _ = v28
                            } else {
                                if (!false) {
                                    var v29 float64 = 0.0
                                    var v30 int32 = (v15 - (v15 + (v15 | -2147483648)))
                                    fmt.Println(v1, v17, v5)
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 int16 = -1
                                    v17 -= 0
                                    _ = v31
                                }
                            }
                            fmt.Println(v15, v1, v17)
                            var v32 int64 = v17
                            var v33 int16 = 42
                            var v34 int32 = -631
                            var v35 int = 32768
                            var v36 int8 = 100
                            _ = v25
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                            _ = v36
                        } else {
                            v1 -= ((3.14159 / v1) - ((0.1 * (3.14159 - (float64(((0.1 + 0.1) / 3.14159)) * v1))) - float64(0.0)))
                            var v37 int64 = v5
                            v38 := GenS1{F0: false}
                            _ = v37
                            _ = v38
                        }
                    }
                }
                _ = fi18
            }
            var v39 bool = (42 > 0)
            _ = fi16
            _ = v17
            _ = v39
        }
    }
    v2[127] = (!false)
    switch "foo bar" {
    case "foo bar":
        v15 = (v15 / ((v15) | 1))
    case "a":
        var v40 uint32 = 392
        _ = v40
    }
    var v41 int8 = 42
    for fi42 := 0; fi42 < 15; fi42++ {
        v43 := GenS3{F0: v15, F1: 932, F2: int32(v41)}
        var v44 float32 = 1e10
        var v45 bool = (42 > 193)
        _ = fi42
        _ = v43
        _ = v44
        _ = v45
    }
    if len(v4) > 0 {
        v4[0] = 32767
    }
    v46 := GenS4{F0: -588}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v5
    _ = v6
}
