package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string; F1 int8; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint8 = 127
    var v2 float32 = ((((3.14159 - -2.25) + 3.14159) - (3.14159 - (float32((0.1 - 1e10)) * 0.1))) * 1e10)
    var v3 int16 = 32767
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    var v6 float32 = v2
    for fi7 := 0; fi7 < 16; fi7++ {
        if (560 == 42) {
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS1{F0: "hello", F1: 2}
                v10 := GenS2{F0: "hello", F1: 100, F2: 0}
                if true {
                    v2 += v6
                } else {
                    v11 := []int{(fi7 * 9223372036854775807), (fi7 ^ fi8), -584}
                    v11 = append(v11, fi7)
                    v12 := GenS3{F0: -2147483648}
                    switch (fi8 + -2147483648) {
                    case 1:
                        var v13 string = ("\n\t" + "foo bar")
                        _ = v13
                    }
                    if len(v11) > 0 {
                        v11[0] = fi7
                    }
                    _ = v11
                    _ = v12
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            var v14 int16 = (v3 / ((v3) | 1))
            v6 -= v2
            _ = v14
        } else {
            v3 -= 1000
            for fi15 := 0; fi15 < 18; fi15++ {
                switch fi7 {
                case 127:
                    var v16 bool = true
                    _ = v16
                case 9223372036854775807:
                    var v17 uint64 = 389
                    _ = v17
                default:
                    v18 := []int{(fi7 / (((fi7 << ((42) & 63))) | 1)), fi15}
                    _ = v18
                }
                _ = fi15
            }
        }
        _ = fi7
    }
    v19 := GenS4{F0: v3, F1: v3}
    v20 := GenS3{F0: 2147483647}
    switch "foo bar" {
    case "a":
        v21 := []int{81, 2147483647}
        _ = v21
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        v1 = (v1 & v1)
    default:
        var v22 int16 = v3
        _ = v22
    }
    var v23 int16 = v3
    v2 += ((v2 * v2) + ((3.14159 + 3.14159) - 1.5))
    var v24 int16 = (v3 >> ((0) & 63))
    v25 := []int{32768, 856}
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint32 = 2
        v3 = v3
        _ = fi26
        _ = v27
    }
    var v28 uint16 = 32767
    for fi29 := 0; fi29 < 5; fi29++ {
        var v30 uint8 = v1
        _ = fi29
        _ = v30
    }
    fmt.Println(v3)
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := v5[1]
        v33 := GenS1{F0: "", F1: 0}
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 float64 = 1.5
    v23 = (v23 + 223)
    v3 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
