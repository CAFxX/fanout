package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int32; F2 int64 }

func main() {
    var v1 bool = (!false)
    var v2 uint64 = 247
    v3 := []int{961, 128}
    var v4 uint8 = 1
    var v5 uint32 = 83
    v1 = true
    v5 -= v5
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
    }
    var v8 int64 = 1
    v8 = (v8 & -1)
    _ = len(v3)
    var v9 float64 = (-2.25 * 3.14159)
    fmt.Println(v1, v1)
    switch "a" {
    case "foo bar":
        v5 = v5
    case "hello":
        var v10 uint8 = (^(v4))
        _ = v10
    default:
        v11 := GenS1{F0: 1, F1: "a", F2: uint(v4)}
        _ = v11
    }
    var v12 string = "\n\t"
    switch v12 {
    case "hello":
        v8 = (-(v8))
    }
    v3 = append(v3, 2147483647)
    switch 0 {
    case -2147483648:
        switch (v12 + v12) {
        case "a":
            if v1 {
                v13 := GenS2{F0: 426}
                v14 := GenS3{F0: 0}
                _ = v13
                _ = v14
            }
        case "foo bar":
            if v1 {
                if (35 <= 0) {
                    if (uint(v2) != 4294967295) {
                        v5 += 4294967295
                        v3 = append(v3, 32768)
                        v15 := GenS4{F0: 65535, F1: 42, F2: (v8 * v8)}
                        _ = v15
                    } else {
                        v16 := GenS4{F0: 439, F1: 106, F2: int64(v2)}
                        _ = v16
                    }
                    v17 := GenS2{F0: -1}
                    _ = v17
                } else {
                    fmt.Println(v5, v12, v4)
                }
                _ = len(v3)
                v4 = v4
            } else {
                v2 = 18446744073709551615
                v12 = (v12 + v12)
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case -85:
        var v18 int8 = 103
        _ = v18
    case 42:
        if (!true) {
            v3 = append(v3, -791)
            var v19 bool = (-32768 == 159)
            _ = v19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
