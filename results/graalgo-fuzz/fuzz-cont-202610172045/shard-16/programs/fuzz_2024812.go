package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint8 }

func main() {
    v1 := []int{1, -860}
    v2 := make(map[int]int)
    v3 := []int{-2147483648, 127, 4294967295}
    var v4 uint64 = 1
    var v5 uint16 = 1
    var v6 string = "\n\t"
    v6 = v6
    var v7 uint = 472
    v8 := GenS1{F0: v7, F1: 65535, F2: 42}
    if (!false) {
        v6 = v6
        v7 = v7
        var v9 float64 = 3.14159
        _ = v9
    } else {
        fmt.Println(v4, v4, v7)
        fmt.Println(v7)
    }
    var v10 float64 = 3.14159
    v7 += v7
    for fi11 := 0; fi11 < 5; fi11++ {
        if ((v5 / ((v5) | 1)) == 2) {
            var v12 uint8 = uint8(v5)
            fmt.Println(fi11, fi11, v4)
            v2[(^(fi11))] = (fi11 * fi11)
            _ = v12
        } else {
            var v13 string = ("\n\t" + (v6 + ""))
            _ = v13
        }
        v7 -= v7
        v14 := GenS1{F0: (v7 & (v7 & 18446744073709551615)), F1: 0, F2: 1000}
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        if ("hello" == v6) {
            var v16 uint8 = 44
            if true {
                var v17 uint32 = 468
                _ = len(v1)
                _ = v17
            }
            var v18 int16 = 42
            _ = v16
            _ = v18
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi15
    }
    v19 := GenS1{F0: 2, F1: 2, F2: 32767}
    v20 := GenS2{F0: 4294967295, F1: 127, F2: 1}
    var v21 uint = (v7 % (((v7 + 594)) | 1))
    fmt.Println(v6, v21)
    switch 32767 {
    case -601:
        fmt.Println(v4, v7)
    case 295:
        for fi22 := 0; fi22 < 2; fi22++ {
            v5 -= (v5 | 84)
            _ = fi22
        }
    default:
        fmt.Println(v4, v4, v7)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
