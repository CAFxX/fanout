package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 bool; F2 uint16 }

func main() {
    v1 := []int{127}
    v2 := make(map[int]string)
    var v3 int16 = -1
    var v4 float32 = -2.25
    var v5 string = "a"
    var v6 int32 = 218
    v6 -= (v6 * v6)
    for fi7 := 0; fi7 < 13; fi7++ {
        v3 = v3
        var v8 uint16 = 790
        _ = fi7
        _ = v8
    }
    var v9 int = 860
    var v10 int16 = (v3 - -842)
    var v11 float64 = (1.5 * (1e10 - ((3.14159 - 0.0) * (float64((v4 * 0.1)) + float64(v4)))))
    switch v5 {
    case "hello":
        v12 := GenS1{F0: 42, F1: 42, F2: (v9 % (((v9 << ((242) & 63))) | 1))}
        _ = v12
    case "a":
        delete(v2, (^(v9)))
    case "\n\t":
        var v13 int = v9
        _ = v13
    default:
        v3 = 1
    }
    if false {
        var v14 int32 = v6
        _ = v14
    } else {
        v15 := GenS1{F0: 2, F1: 646, F2: v9}
        _ = v15
    }
    if false {
        if true {
            var v16 float32 = (v4 / v4)
            v4 = v16
            _ = v16
        } else {
            v10 = (v10 + (-(v3)))
            v4 += (v4 / (v4 + (v4 / (v4 - float32((((3.14159 * 0.0) - v11) - float64(v4)))))))
        }
        v17 := GenS1{F0: 1, F1: 463, F2: 0}
        v18 := GenS1{F0: 0, F1: 2, F2: v9}
        _ = v17
        _ = v18
    } else {
        fmt.Println(v6, v11)
        delete(v2, -1)
        if (2 >= 2) {
            for fi19 := 0; fi19 < 2; fi19++ {
                if true {
                    v20 := v2[(fi19 ^ int(v6))]
                    _ = v20
                } else {
                    var v21 uint16 = 2
                    var v22 bool = true
                    _ = v21
                    _ = v22
                }
                switch "hello" {
                case "a":
                    fmt.Println(v6, v10, v11)
                }
                _ = fi19
            }
        }
    }
    v23 := GenS2{F0: 42}
    v24 := v2[(v9 - 32768)]
    v25 := GenS3{F0: false, F1: true, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
