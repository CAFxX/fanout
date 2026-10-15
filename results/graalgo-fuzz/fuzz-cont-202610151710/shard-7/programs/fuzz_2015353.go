package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int8; F1 bool }

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 825
    var v3 uint8 = 1
    v4 := make(map[uint32]bool)
    v5 := []int{}
    for fi6 := 0; fi6 < 20; fi6++ {
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        fmt.Println(fi6, v2)
        switch (fi6 >> ((475) & 63)) {
        case 4294967295:
            v7 := GenS1{F0: fi6, F1: 274}
            _ = v7
        case -2147483648:
            fmt.Println(fi6, v1, v1)
        }
        _ = fi6
    }
    var v8 uint8 = v3
    var v9 bool = ((v3 / ((uint8(v2)) | 1)) < 2)
    var v10 string = ("hello" + "")
    if true {
        v11 := GenS2{F0: 32768, F1: v2, F2: 1}
        _ = v11
    } else {
        v12 := GenS2{F0: 9223372036854775807, F1: (v2 | (v2 | v2)), F2: 0}
        _ = v12
    }
    var v13 int64 = 0
    if v9 {
        v1 = uint32(v2)
        var v14 float64 = ((0.1 + (1e10 - 3.14159)) * float64(1.5))
        v1 -= v1
        _ = v14
    } else {
        v15 := GenS3{F0: 758}
        var v16 uint8 = 127
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        switch (-(fi17)) {
        case 600:
            var v18 int16 = 42
            _ = v18
        case 127:
            if v9 {
                v5 = append(v5, fi17)
                fmt.Println(v13)
                if false {
                    v19 := GenS1{F0: fi17, F1: 9223372036854775807}
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 int32 = -2147483648
                        fmt.Println(v2, v8)
                        if v9 {
                            v22 := v4[(v1 << ((2) & 63))]
                            delete(v4, v1)
                            _ = v22
                        } else {
                            if v9 {
                                v4[uint32(v8)] = (v9 || v9)
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = 6
                                }
                                var v23 uint = 1
                                v9 = false
                                _ = len(v4)
                                _ = v23
                            }
                            if len(v5) > 0 {
                                v5[0] = fi20
                            }
                            if v9 {
                                v8 -= 233
                                v9 = (v9 || (!false))
                            }
                        }
                        _ = fi20
                        _ = v21
                    }
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v2 = v2
                        if true {
                            switch fi17 {
                            case 0:
                                var v25 uint = 42
                                _ = v25
                            case -9223372036854775808:
                                var v26 float64 = 1.5
                                _ = v26
                            default:
                                v3 += v3
                            }
                            switch v10 {
                            case "a":
                                if (-939 <= 0) {
                                    v27 := GenS1{F0: (fi17 ^ fi24), F1: (v13 * (v13 + (v13 & (v13 & v13))))}
                                    _ = v27
                                } else {
                                    fmt.Println(v10, v2, fi24)
                                }
                            case "":
                                fmt.Println(v8, v1)
                            case "\n\t":
                                var v28 int = (fi24 + (fi24 * (-(fi24))))
                                _ = v28
                            default:
                                var v29 float32 = 0.0
                                _ = v29
                            }
                            switch (fi24 - fi24) {
                            case 128:
                                var v30 uint32 = v1
                                _ = v30
                            case 0:
                                var v31 uint = uint(v1)
                                _ = v31
                            }
                        } else {
                            if true {
                                v32 := GenS1{F0: (-(fi24)), F1: -1}
                                if (4294967295 < 4294967295) {
                                    switch 9223372036854775807 {
                                    case 32768:
                                        v1 -= 820
                                    default:
                                        v33 := GenS1{F0: (fi24 + (fi24 ^ fi24)), F1: 9223372036854775807}
                                        _ = v33
                                    }
                                }
                                _ = v32
                            } else {
                                v5 = append(v5, (fi17 * fi17))
                                var v34 bool = true
                                var v35 uint64 = 293
                                _ = v34
                                _ = v35
                            }
                        }
                        v4[v1] = ((0.1 / 3.14159) != 0.0)
                        _ = fi24
                    }
                    _ = v19
                }
            }
        case 280:
            var v36 int8 = 1
            _ = v36
        }
        v9 = true
        switch fi17 {
        case -2147483648:
            for fi37 := 0; fi37 < 14; fi37++ {
                v38 := GenS2{F0: (^(fi17)), F1: 941, F2: v13}
                if true {
                    _ = len(v5)
                    var v39 float32 = 0.1
                    switch fi17 {
                    case 665:
                        _ = len(v5)
                    case -726:
                        v1 -= (^(v1))
                        v40 := GenS4{F0: v10, F1: 18446744073709551615}
                        var v41 uint8 = v8
                        _ = v40
                        _ = v41
                    }
                    _ = v39
                } else {
                    var v42 uint32 = v1
                    _ = v42
                }
                v43 := GenS4{F0: "hello", F1: 4294967295}
                _ = fi37
                _ = v38
                _ = v43
            }
        case 42:
            v44 := GenS5{F0: -1, F1: (v1 > (v1 & v1))}
            _ = v44
        case 2147483647:
            v13 += v13
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
