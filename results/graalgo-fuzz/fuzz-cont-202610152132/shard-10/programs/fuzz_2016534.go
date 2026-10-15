package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 int; F2 uint32 }
type GenS9 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint64) uint32 {
    p0 -= 1
    p0 -= (p0 / (((p0 + 18446744073709551615)) | 1))
    var v7 int = -858
    v8 := GenS1{F0: -2147483648}
    v9 := GenS2{F0: 889}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 918
}
func genFn1(p0 bool, p1 uint16) int16 {
    if (18446744073709551615 == 208) {
        var v10 int64 = 9223372036854775807
        p1 = p1
        switch -363 {
        case 127:
            p0 = true
        case -9223372036854775808:
            var v11 int = 32768
            _ = v11
        default:
            v12 := []int{}
            _ = v12
        }
        _ = v10
    }
    var v13 int32 = -2147483648
    var v14 uint8 = 215
    fmt.Println(v14, p0)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 1000
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 uint = 2
    v5 := make(map[uint32]bool)
    var v6 int16 = 0
    v15 := GenS1{F0: -2147483648}
    v16 := GenS3{F0: 42}
    var v17 int64 = v1
    v17 += v1
    v18 := GenS4{F0: 78}
    if false {
        v4 = v4
        var v19 int16 = v6
        _ = v19
    }
    v1 = 42
    v20 := []int{4294967295}
    var v21 uint16 = uint16(v4)
    if false {
        switch 127 {
        case 2147483647:
            v22 := GenS5{F0: (v1 % ((v17) | 1))}
            _ = v22
        case 4294967295:
            if (false && false) {
                if (v4 == 4294967295) {
                    if (-1 == 0) {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            v2["a"] = int(v6)
                            switch fi23 {
                            case 0:
                                switch "" {
                                case "a":
                                    if (fi23 < fi23) {
                                        switch (^(fi23)) {
                                        case -9223372036854775808:
                                            var v24 uint = 1
                                            _ = v24
                                        case 37:
                                            var v25 uint = 0
                                            _ = v25
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            var v26 int16 = (v6 + v6)
                                            _ = v26
                                        case "foo bar":
                                            v17 = 42
                                        default:
                                            for fi27 := 0; fi27 < 6; fi27++ {
                                                v28 := GenS6{F0: (-(fi27))}
                                                v17 = v17
                                                _ = fi27
                                                _ = v28
                                            }
                                        }
                                        var v29 int = (fi23 / (((fi23 >> ((uint(v21)) & 63))) | 1))
                                        _ = v29
                                    }
                                case "foo bar":
                                    if (!((1 != 52) && true)) {
                                        for fi30 := 0; fi30 < 14; fi30++ {
                                            var v31 int64 = 9223372036854775807
                                            var v32 int32 = -1
                                            if false {
                                                var v33 uint64 = 474
                                                var v34 string = "foo bar"
                                                _ = v33
                                                _ = v34
                                            } else {
                                                switch (fi23 % (((fi23 | (fi23 / (((-(fi30))) | 1)))) | 1)) {
                                                case 128:
                                                    switch fi23 {
                                                    case 127:
                                                        var v35 int32 = v32
                                                        _ = v35
                                                    case 42:
                                                        var v36 int32 = (-(v32))
                                                        fmt.Println(v32, v32, v36)
                                                        _ = v36
                                                    default:
                                                        var v37 bool = false
                                                        _ = v37
                                                    }
                                                    if false {
                                                        switch "hello" {
                                                        case "a":
                                                            switch "a" {
                                                            case "a":
                                                                v38 := GenS1{F0: (-(v32))}
                                                                _ = v38
                                                            default:
                                                                var v39 float64 = 1e10
                                                                _ = v39
                                                            }
                                                        case "foo bar":
                                                            var v40 int8 = 0
                                                            _ = v40
                                                        case "":
                                                            v5[350] = (!false)
                                                        default:
                                                            v41 := GenS2{F0: (v6 << (((v4 >> ((v4) & 63))) & 63))}
                                                            v42 := GenS1{F0: (v32 - 2147483647)}
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    }
                                                }
                                                v31 += 881
                                                v43 := v5[2]
                                                _ = v43
                                            }
                                            v44 := GenS7{F0: (v6 + v6)}
                                            _ = fi30
                                            _ = v31
                                            _ = v32
                                            _ = v44
                                        }
                                        v1 += v17
                                        v45 := GenS8{F0: (^(v21)), F1: -1, F2: 2}
                                        _ = v45
                                    }
                                default:
                                    v46 := v5[49]
                                    _ = v46
                                }
                            default:
                                switch -286 {
                                case -6:
                                    v6 = (v6 * genFn1(false, (v21 & v21)))
                                case 32767:
                                    var v47 uint32 = genFn0(974)
                                    _ = v47
                                case -9223372036854775808:
                                    for fi48 := 0; fi48 < 15; fi48++ {
                                        v21 += v21
                                        _ = fi48
                                    }
                                default:
                                    for fi49 := 0; fi49 < 3; fi49++ {
                                        var v50 uint16 = v21
                                        v17 += 1
                                        v51 := GenS7{F0: (v6 >> ((v4) & 63))}
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                }
                            }
                            _ = fi23
                        }
                        _ = len(v2)
                    }
                } else {
                    _ = len(v2)
                    delete(v5, 778)
                }
                var v52 float32 = float32((float64(((((0.0 + 1e10) + 0.1) / -2.25) + 1e10)) - (1.5 - 0.1)))
                _ = v52
            }
        case -611:
            v53 := GenS9{F0: 4294967295, F1: (v6 ^ int16(v21))}
            _ = v53
        default:
            var v54 float32 = 1.5
            _ = v54
        }
    }
    fmt.Println(v6, v6, v6)
    v6 = int16(v4)
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
