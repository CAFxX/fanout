package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 bool; F2 int64 }
type GenS3 struct { F0 int32; F1 int; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 970
    v3 := []int{128}
    switch 9223372036854775807 {
    case -1:
        fmt.Println(v2)
    case 42:
        if ((false || (((float32(-2.25) / 1.0) == 3.14159) && (!(-32768 <= 128)))) && true) {
            for fi4 := 0; fi4 < 1; fi4++ {
                v2 = (v2 / (((v2 * v2)) | 1))
                _ = fi4
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            if true {
                v2 = 18446744073709551615
                for fi5 := 0; fi5 < 8; fi5++ {
                    v1[fi5] = fi5
                    v6 := GenS1{F0: -2147483648}
                    fmt.Println(v2, fi5)
                    _ = fi5
                    _ = v6
                }
                v7 := GenS1{F0: 2147483647}
                _ = v7
            }
        }
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        if false {
            v9 := GenS2{F0: fi8, F1: (-80 <= -128), F2: 1}
            _ = v9
        }
        _ = fi8
    }
    switch -365 {
    case -9223372036854775808:
        v2 = v2
    case 127:
        var v10 int = 96
        _ = v10
    default:
        v3 = append(v3, -376)
    }
    if false {
        var v11 uint8 = 13
        _ = v11
    }
    v2 -= v2
    v2 += 18446744073709551615
    v2 += 42
    var v12 float32 = ((3.14159 * 0.0) / 0.1)
    switch -207 {
    case 2147483647:
        fmt.Println(v2)
    default:
        delete(v1, 32767)
    }
    v13 := GenS3{F0: -295, F1: 42, F2: 42}
    var v14 int = 9223372036854775807
    fmt.Println(v2, v2, v12)
    fmt.Println(v12, v14)
    if len(v3) > 0 {
        v3[0] = (v14 | 42)
    }
    v12 += (0.1 * ((-2.25 / v12) * -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
