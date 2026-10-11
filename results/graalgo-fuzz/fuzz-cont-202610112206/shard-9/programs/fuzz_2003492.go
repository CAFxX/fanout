package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint32, p1 int32, p2 int8) int64 {
    var v4 int8 = p2
    var v5 int16 = 1000
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 float64 = 0.1
        switch 0 {
        case 32767:
            fmt.Println(v4, p1, v7)
        default:
            fmt.Println(v5)
        }
        var v8 uint64 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 float64 = 1.5
    var v2 int8 = 42
    var v3 uint = 18446744073709551615
    var v9 uint = v3
    v10 := []int{-9223372036854775808, 128, 979, 4294967295}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v1 = v1
    var v11 uint32 = 2147483647
    v3 = 855
    fmt.Println(v2)
    if false {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        var v12 uint = (v9 - (^(v3)))
        _ = v12
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 string = "a"
            switch (fi13 - fi13) {
            case 2147483647:
                var v15 float32 = 3.14159
                _ = v15
            case 0:
                v11 = v11
            default:
                if (false || false) {
                    var v16 int32 = 0
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if (!(v16 > 0)) {
                            if false {
                                v11 = v11
                            } else {
                                var v18 uint8 = 255
                                v18 = 127
                                v19 := GenS1{F0: -998}
                                var v20 uint = v9
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS2{F0: 42, F1: (v16 >> ((18446744073709551615) & 63)), F2: 32767}
                            switch (fi17 ^ (fi17 >> ((v3) & 63))) {
                            case 2147483647:
                                v10 = append(v10, fi17)
                            case 184:
                                var v22 int64 = 374
                                _ = v22
                            case -454:
                                v9 += v3
                                var v23 uint8 = 42
                                v23 = 42
                                v9 -= (v9 | (^(v3)))
                                _ = v23
                            }
                            _ = v21
                        } else {
                            v9 += (^(v3))
                            v1 = (-2.25 - ((float64(1.5) * v1) * (-2.25 + ((v1 - v1) - (1e10 * (v1 * ((v1 * v1) + float64(0.0))))))))
                        }
                        _ = fi17
                    }
                    fmt.Println(v9)
                    _ = v16
                } else {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 float32 = 1.5
                        v26 := GenS2{F0: 1000, F1: -2147483648, F2: 432}
                        v27 := GenS3{F0: 42, F1: -2147483648}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
            }
            _ = fi13
            _ = v14
        }
        var v28 int32 = 2147483647
        v2 = 100
        _ = v28
    }
    switch "" {
    case "a":
        var v29 uint8 = 1
        _ = v29
    case "hello":
        var v30 uint16 = 32767
        _ = v30
    default:
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 int = fi31
            var v33 bool = (352 <= 596)
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    var v34 float32 = 1.5
    var v35 uint64 = 1
    if (1 <= 1) {
        v3 = 4294967295
        v9 = v3
    } else {
        v34 -= (float32(v1) * ((0.0 + 3.14159) * v34))
        v36 := GenS1{F0: 454}
        _ = v36
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        if true {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            if len(v10) > 0 {
                v10[0] = fi37
            }
            var v38 uint8 = 138
            _ = v38
        }
        for fi39 := 0; fi39 < 2; fi39++ {
            var v40 bool = (!false)
            v41 := GenS3{F0: -1, F1: -2147483648}
            var v42 int = -151
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
        }
        fmt.Println(v9)
        _ = fi37
    }
    var v43 int16 = -977
    var v44 float32 = float32(v1)
    switch "hello" {
    case "foo bar":
        switch 131 {
        case 127:
            var v45 string = ""
            _ = v45
        case -596:
            fmt.Println(v2)
        case 42:
            v10 = append(v10, 2147483647)
        default:
            for fi46 := 0; fi46 < 18; fi46++ {
                v43 -= v43
                _ = fi46
            }
        }
    }
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v43
    _ = v44
    _ = v9
}
