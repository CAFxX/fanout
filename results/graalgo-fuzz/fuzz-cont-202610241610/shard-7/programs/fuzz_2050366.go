package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 uint) int16 {
    p0 = p0
    var v6 uint64 = 42
    v7 := GenS1{F0: -99}
    _ = p0
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 int = -711
    var v2 uint32 = 42
    var v3 uint8 = 2
    var v4 int64 = 1
    var v5 bool = false
    for fi8 := 0; fi8 < 15; fi8++ {
        v5 = (-9223372036854775808 < (v4 ^ (v4 - v4)))
        v9 := []int{4294967295}
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: 9223372036854775807}
            v2 -= (v2 + (^(v2)))
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := []int{}
    for fi13 := 0; fi13 < 7; fi13++ {
        if v5 {
            v14 := GenS1{F0: v4}
            v1 = (v1 * (v1 % ((int(v4)) | 1)))
            _ = v14
        }
        _ = fi13
    }
    if len(v12) > 0 {
        v12[0] = v1
    }
    if (0 < 940) {
        v12 = append(v12, (-(v1)))
        v12 = append(v12, v1)
    }
    var v15 uint16 = 1
    switch "" {
    case "":
        v3 += v3
    case "hello":
        v4 -= (v4 * v4)
    case "\n\t":
        switch (v1 >> ((1) & 63)) {
        case 9223372036854775807:
            for fi16 := 0; fi16 < 10; fi16++ {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 uint8 = v3
                    v19 := GenS1{F0: v4}
                    if ((v18 * v18) != v18) {
                        v15 -= (v15 - (v15 / ((v15) | 1)))
                        fmt.Println(fi17, v2, fi17)
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        case -974:
            var v20 uint16 = (v15 ^ (v15 ^ (v15 % ((255) | 1))))
            _ = v20
        case -2147483648:
            _ = len(v12)
        default:
            fmt.Println(v3)
        }
    }
    v4 += (^(v4))
    fmt.Println(v3, v2, v3)
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint16 = 2
        _ = fi21
        _ = v22
    }
    if (float32((-2.25 + 0.1)) != ((-2.25 / (((0.1 - 0.0) / -2.25) * (float32(0.1) - -2.25))) + -2.25)) {
        var v23 uint = 0
        v15 = (^(v15))
        v12 = append(v12, v1)
        _ = v23
    }
    v15 -= (v15 ^ v15)
    v3 = (v3 & v3)
    switch (v1 >> ((2) & 63)) {
    case -38:
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS2{F0: "hello", F1: 1, F2: v2}
            if true {
                switch v1 {
                case 128:
                    _ = len(v12)
                case -283:
                    var v26 float64 = 0.1
                    _ = v26
                }
                var v27 int = -9223372036854775808
                _ = v27
            } else {
                _ = len(v12)
                v2 -= 231
                v28 := GenS3{F0: "\n\t"}
                _ = v28
            }
            var v29 uint32 = 4294967295
            _ = fi24
            _ = v25
            _ = v29
        }
    default:
        v30 := GenS1{F0: (v4 >> ((343) & 63))}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
