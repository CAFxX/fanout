package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
type GenS4 struct { F0 int; F1 uint32; F2 string }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    var v3 string = v1
    v4 := make(map[uint32]string)
    _ = len(v2)
    switch -826 {
    case 0:
        v5 := GenS1{F0: 255, F1: (-1 > 32767), F2: 34}
        _ = v5
    case 4294967295:
        fmt.Println(v3)
    case 904:
        switch v3 {
        case "hello":
            var v6 int32 = 0
            _ = v6
        case "foo bar":
            v7 := v2[953]
            _ = v7
        case "a":
            if true {
                v8 := []int{1, 229, 552, -2147483648}
                _ = v8
            }
        default:
            v4[562] = v1
        }
    default:
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS2{F0: 18446744073709551615, F1: false}
            for fi11 := 0; fi11 < 15; fi11++ {
                v3 = "\n\t"
                var v12 uint64 = 876
                v13 := GenS1{F0: 4294967295, F1: (-1 > 77), F2: 42}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
    }
    v14 := GenS3{F0: 127, F1: 1, F2: ((v3 + "hello") + "a")}
    v15 := []int{42, -9223372036854775808, -545}
    var v16 string = v1
    var v17 uint64 = 2
    v17 -= (v17 & (v17 / (((v17 - v17)) | 1)))
    v18 := GenS4{F0: 4294967295, F1: 364, F2: ((v1 + "foo bar") + (v16 + v1))}
    if (float64(0.0) != ((((float64((float32((0.0 - 3.14159)) - ((((1.5 * ((((((-2.25 + 3.14159) + float32(-2.25)) + -2.25) - 1.5) * float32(1e10)) + 1e10)) - -2.25) + 1.5) * 0.1))) - (float64(-2.25) + (-2.25 * ((1.5 * -2.25) * float64(3.14159))))) * (0.1 + 0.1)) + float64(1.5)) / 3.14159)) {
        v17 += v17
        fmt.Println(v17, v3, v16)
    }
    var v19 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
