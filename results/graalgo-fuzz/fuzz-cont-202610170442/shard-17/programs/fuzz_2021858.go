package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS4 struct { F0 string; F1 uint; F2 bool }
type GenS5 struct { F0 int64; F1 int; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.1
    v3 := []int{32768, 9223372036854775807, -527}
    var v4 uint = 257
    for fi5 := 0; fi5 < 18; fi5++ {
        if (42 <= 344) {
            var v6 uint8 = 2
            _ = v6
        } else {
            var v7 int64 = 971
            _ = v7
        }
        v8 := GenS1{F0: 129, F1: -623, F2: v4}
        _ = fi5
        _ = v8
    }
    v9 := GenS2{F0: 4294967295}
    var v10 bool = (!false)
    for fi11 := 0; fi11 < 1; fi11++ {
        delete(v1, "foo bar")
        v12 := GenS3{F0: 0, F1: 4294967295, F2: 152}
        _ = fi11
        _ = v12
    }
    v3 = append(v3, 352)
    v13 := v1["a"]
    v13 = v13
    var v14 uint = 18446744073709551615
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int32 = 1
            switch 9223372036854775807 {
            case 127:
                v10 = false
            case 839:
                v17 = v17
                var v18 int32 = 42
                _ = v18
            default:
                var v19 int = (fi15 * fi16)
                _ = v19
            }
            var v20 uint64 = 159
            _ = fi16
            _ = v17
            _ = v20
        }
        var v21 uint64 = 378
        for fi22 := 0; fi22 < 1; fi22++ {
            if ((fi15 << ((v4) & 63)) != (fi22 | (fi15 | 2147483647))) {
                var v23 int16 = -792
                _ = v23
            }
            v10 = true
            var v24 uint64 = 42
            v24 = 676
            if (0 >= -9223372036854775808) {
                if (848 >= 861) {
                    v13 = v13
                }
                var v25 uint8 = 198
                var v26 int32 = 425
                v27 := GenS4{F0: "", F1: (v4 >> ((v4) & 63)), F2: v10}
                v3 = append(v3, 744)
                var v28 int64 = -1
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            } else {
                v21 -= (v24 & v21)
                v2 = 1.5
            }
            switch v13 {
            case "hello":
                if v10 {
                    switch -9223372036854775808 {
                    case 1:
                        switch v13 {
                        case "a":
                            switch 42 {
                            case 389:
                                v10 = false
                            case 32767:
                                if v10 {
                                    v4 = 1
                                } else {
                                    v29 := GenS5{F0: -577, F1: 42, F2: (255 == 61)}
                                    _ = len(v3)
                                    var v30 int8 = 0
                                    _ = v29
                                    _ = v30
                                }
                            default:
                                fmt.Println(v14, v2, fi22)
                            }
                        case "\n\t":
                            var v31 int8 = 1
                            _ = v31
                        }
                        var v32 int8 = 93
                        var v33 uint64 = v21
                        _ = v32
                        _ = v33
                    case -9223372036854775808:
                        v21 += v24
                    case 0:
                        if (-2147483648 >= -1) {
                            switch fi15 {
                            case 32768:
                                if (42 >= 401) {
                                    fmt.Println(v10, v4, v13)
                                    v34 := GenS4{F0: v13, F1: (^(v14)), F2: (!(!((-664 >= -740) || v10)))}
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(fi15)
                                    default:
                                        var v35 uint16 = 88
                                        _ = v35
                                    }
                                    _ = v34
                                } else {
                                    _ = len(v1)
                                    switch v13 {
                                    case "a":
                                        if ((v14 / ((v4) | 1)) <= 4294967295) {
                                            v21 = (v21 - 18446744073709551615)
                                        }
                                    }
                                    v13 = "foo bar"
                                }
                            case 128:
                                if false {
                                    var v36 bool = (0 < -268)
                                    if (42 == (fi22 & (fi15 * 32768))) {
                                        fmt.Println(v14, fi15, v21)
                                        switch 4294967295 {
                                        case -9223372036854775808:
                                            v2 = float64(float32((float64(float32((1e10 / v2))) - (3.14159 * (v2 * v2)))))
                                        }
                                        var v37 int64 = 1
                                        _ = v37
                                    } else {
                                        v21 = 381
                                    }
                                    fmt.Println(fi22, v2)
                                    if (2147483647 >= -309) {
                                        v38 := GenS5{F0: 42, F1: fi22, F2: false}
                                        if len(v3) > 0 {
                                            v3[0] = fi22
                                        }
                                        delete(v1, ((v13 + "\n\t") + (v13 + ("a" + v13))))
                                        _ = v38
                                    }
                                    _ = v36
                                }
                            case -691:
                                fmt.Println(v21)
                            default:
                                if (0 != 0) {
                                    var v39 int32 = 0
                                    fmt.Println(v24, v13, fi22)
                                    v40 := GenS2{F0: (v21 % ((v24) | 1))}
                                    var v41 float64 = float64(float32((((v2 / v2) * ((0.0 / (v2 - float64(1e10))) + v2)) / (v2 * (-2.25 + (v2 - v2))))))
                                    var v42 float32 = (3.14159 * 3.14159)
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                            }
                        }
                    }
                }
            case "a":
                v13 = v13
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if v10 {
                var v43 int64 = 42
                _ = v43
            }
            fmt.Println(v13)
            _ = fi22
            _ = v24
        }
        _ = fi15
        _ = v21
    }
    switch 4294967295 {
    case -459:
        v2 = (v2 / v2)
    case 9223372036854775807:
        var v44 float32 = -2.25
        _ = v44
    }
    var v45 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v9
}
