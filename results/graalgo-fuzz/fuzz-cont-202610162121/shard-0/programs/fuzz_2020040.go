package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 int; F1 uint }
func genFn0() uint8 {
    var v6 int8 = 0
    var v7 float64 = 0.0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint = 642
    var v2 uint32 = 2147483647
    var v3 float64 = 1.5
    var v4 uint16 = 0
    v5 := make(map[int]string)
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int16 = -142
        _ = fi8
        _ = v9
    }
    var v10 int8 = 1
    v11 := GenS1{F0: v4, F1: 845, F2: 32767}
    var v12 uint64 = 18446744073709551615
    var v13 int8 = -33
    var v14 float32 = (3.14159 + 3.14159)
    v1 = v1
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(v1)
        if ((v12 - v12) < (v12 * v12)) {
            v14 += v14
            _ = len(v5)
            v14 -= (v14 + float32(float64((v14 - (v14 / 1.5)))))
        } else {
            v16 := GenS1{F0: v4, F1: -316, F2: -435}
            v2 = v2
            _ = v16
        }
        v10 = v10
        _ = fi15
    }
    v10 -= (^(v13))
    if (!(v14 <= 1.5)) {
        var v17 uint16 = (v4 % (((v4 | (v4 / ((749) | 1)))) | 1))
        v2 = (v2 & v2)
        _ = v17
    }
    v18 := GenS1{F0: (v4 ^ (v4 >> (((^(v1))) & 63))), F1: 352, F2: 1}
    v14 += v14
    if ("hello" < ("a" + "\n\t")) {
        v19 := v5[int(v12)]
        var v20 uint8 = 136
        var v21 int64 = 940
        _ = v19
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                switch "foo bar" {
                case "foo bar":
                    v14 = 1e10
                case "a":
                    v24 := GenS2{F0: fi22, F1: v1}
                    _ = v24
                }
                v25 := []int{(fi22 & 841), fi23}
                v26 := GenS1{F0: v4, F1: 1000, F2: 1000}
                v1 -= 255
                _ = fi23
                _ = v25
                _ = v26
            }
            _ = fi22
        }
    }
    if (!(-488 <= 299)) {
        v27 := GenS2{F0: -83, F1: uint(v4)}
        v12 -= (^(v12))
        _ = v27
    } else {
        switch "hello" {
        case "a":
            switch "a" {
            case "\n\t":
                var v28 bool = true
                _ = v28
            case "foo bar":
                v3 = float64(v14)
            default:
                if false {
                    v29 := GenS1{F0: 326, F1: 0, F2: -333}
                    v12 = v12
                    _ = v29
                }
            }
        case "":
            var v30 uint8 = 127
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
