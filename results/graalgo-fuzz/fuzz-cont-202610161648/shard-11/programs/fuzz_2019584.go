package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 string; F2 int32 }

func main() {
    var v1 uint = 42
    var v2 uint64 = 4294967295
    v3 := make(map[int]int)
    v4 := GenS1{F0: v2, F1: 42, F2: -32768}
    var v5 int64 = 88
    fmt.Println(v2, v5)
    if true {
        v5 = 1
    }
    if false {
        v6 := []int{4294967295}
        _ = v6
    } else {
        v7 := []int{42, 1, 257, -228}
        v1 = 535
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        switch -9223372036854775808 {
        case 4294967295:
            v9 := GenS1{F0: v2, F1: 643, F2: -32768}
            _ = v9
        default:
            var v10 int16 = int16(fi8)
            _ = v10
        }
        switch "" {
        case "a":
            fmt.Println(v1)
        }
        _ = fi8
    }
    v11 := GenS2{F0: 1000, F1: -110, F2: true}
    v3[127] = -2147483648
    var v12 int32 = 0
    switch 0 {
    case -26:
        v2 -= 306
    case -403:
        switch ("a" + "hello") {
        case "a":
            var v13 bool = false
            _ = v13
        case "":
            var v14 int8 = 6
            _ = v14
        default:
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 float64 = (((0.1 / -2.25) / (1e10 - 0.1)) - (0.1 * -2.25))
                _ = fi15
                _ = v16
            }
        }
    }
    if (779 == 32767) {
        if (true && true) {
            v2 = v2
        }
        v2 += 0
        var v17 int16 = 1
        _ = v17
    } else {
        var v18 float32 = ((0.1 + float32(((-2.25 - 3.14159) - (((float64((0.0 * 1.5)) + 1.5) / ((1e10 - ((3.14159 - -2.25) * (1.5 - (((((float64((1.5 * 1e10)) * ((1.5 * 1.5) - (1.5 + float64(float32(float64(3.14159)))))) - 3.14159) / 3.14159) / 1.5) + 3.14159)))) - 1e10)) - 3.14159)))) * 0.1)
        _ = v18
    }
    v19 := GenS1{F0: (^(v2)), F1: 2147483647, F2: 1000}
    if (0 < 1) {
        var v20 int32 = v12
        v21 := GenS2{F0: -965, F1: -128, F2: (!(!false))}
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 13; fi22++ {
            for fi23 := 0; fi23 < 12; fi23++ {
                switch fi23 {
                case -821:
                    v2 += v2
                    if false {
                        fmt.Println(v2)
                        v24 := GenS3{F0: "\n\t"}
                        _ = v24
                    }
                case -1:
                    if true {
                        if false {
                            v25 := GenS4{F0: v12, F1: (("a" + "\n\t") + "hello"), F2: (v12 ^ v12)}
                            _ = v25
                        }
                        v12 = 2147483647
                        var v26 uint = (v1 | (v1 >> (((v1 << ((v1) & 63))) & 63)))
                        _ = v26
                    }
                default:
                    v27 := []int{fi22, fi23, fi22}
                    v2 -= (v2 * (v2 + v2))
                    var v28 uint32 = 42
                    _ = v27
                    _ = v28
                }
                var v29 uint64 = v2
                _ = fi23
                _ = v29
            }
            var v30 int32 = 1
            _ = fi22
            _ = v30
        }
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 int16 = -1
            _ = fi31
            _ = v32
        }
    }
    fmt.Println(v5)
    var v33 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
