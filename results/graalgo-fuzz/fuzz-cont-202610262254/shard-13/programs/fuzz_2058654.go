package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }

func main() {
    var v1 bool = false
    v2 := []int{127, -9223372036854775808, 265}
    var v3 uint64 = 0
    for fi4 := 0; fi4 < 18; fi4++ {
        v3 = 42
        _ = fi4
    }
    v5 := GenS1{F0: 255}
    v1 = (944 < -2147483648)
    switch 128 {
    case -9223372036854775808:
        for fi6 := 0; fi6 < 17; fi6++ {
            if v1 {
                v1 = v1
            } else {
                v7 := GenS2{F0: v3, F1: uint32(v3), F2: v3}
                _ = v7
            }
            if v1 {
                var v8 int64 = 1
                _ = v8
            }
            _ = fi6
        }
    case 128:
        v3 = v3
    default:
        v9 := GenS3{F0: v1}
        _ = v9
    }
    v1 = ((1e10 - ((3.14159 + float32(1e10)) + 0.1)) != (1.5 / float32(1.5)))
    var v10 float64 = float64(0.1)
    v10 = ((v10 * 1.5) + (v10 + ((v10 / (v10 * ((3.14159 * v10) - float64((0.1 + ((1.5 - 0.1) + ((3.14159 + 0.0) - -2.25))))))) - (v10 - (1e10 - v10)))))
    var v11 float64 = 1.5
    if ((v1 && (false || ((v3 ^ (v3 / ((v3) | 1))) < v3))) || (!(float32((v11 * v10)) >= ((0.1 + 1.5) + (-2.25 + ((0.1 / float32((v11 * v11))) - 1e10)))))) {
        v12 := GenS2{F0: (v3 * 568), F1: 1, F2: (v3 - 2)}
        v13 := GenS3{F0: (!v1)}
        _ = v12
        _ = v13
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (false && false) {
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
    }
    v14 := GenS4{F0: ((("foo bar" + "foo bar") + ("a" + ("\n\t" + ("\n\t" + "\n\t")))) + "")}
    switch "" {
    case "hello":
        v3 = 341
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v15 := GenS3{F0: v1}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v5
}
