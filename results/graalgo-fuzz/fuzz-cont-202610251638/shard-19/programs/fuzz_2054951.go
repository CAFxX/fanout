package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint64 }
func genFn0() uint64 {
    var v5 float32 = 0.1
    v6 := GenS1{F0: 32767}
    v5 = v5
    for fi7 := 0; fi7 < 1; fi7++ {
        v5 = float32(1e10)
        v5 += (0.0 * v5)
        _ = fi7
    }
    _ = v5
    _ = v6
    return 94
}
func genFn1(p0 int8) uint32 {
    p0 -= p0
    switch ("hello" + "\n\t") {
    case "foo bar":
        p0 -= 0
    case "":
        v8 := []int{0, -515, 0}
        _ = v8
    }
    var v9 bool = false
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 int = 42
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 string = "foo bar"
    v1 = (v1 - (v1 + (-(v1))))
    v1 -= (v1 - 1)
    var v10 int8 = -128
    v1 = v1
    switch (v4 + "") {
    case "hello":
        var v11 bool = false
        _ = v11
    default:
        v12 := GenS2{F0: 18446744073709551615, F1: 42, F2: 1}
        _ = v12
    }
    fmt.Println(v4)
    _ = len(v3)
    if false {
        var v13 int16 = -339
        if (65535 > 704) {
            if true {
                v3 = append(v3, v1)
                if true {
                    var v14 uint32 = 187
                    _ = v14
                }
                v10 = (v10 ^ (-(v10)))
            } else {
                v10 += (v10 + (v10 - v10))
            }
            v15 := GenS3{F0: 1000}
            switch (v1 - 9) {
            case -9223372036854775808:
                v16 := GenS4{F0: 944, F1: 999}
                _ = v16
            case 32767:
                v4 = v4
            case 0:
                v17 := GenS1{F0: 255}
                _ = v17
            }
            _ = v15
        }
        _ = v13
    }
    var v18 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
