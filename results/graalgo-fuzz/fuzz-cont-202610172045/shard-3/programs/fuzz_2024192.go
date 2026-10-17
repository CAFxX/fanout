package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
func genFn0(p0 int32, p1 int8) uint64 {
    p1 = (p1 ^ 100)
    fmt.Println(p1, p1)
    if (2 <= 42) {
        v7 := []int{-1, -203}
        p0 += p0
        _ = v7
    }
    var v8 uint16 = 32767
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 uint) uint {
    v9 := []int{-2147483648, 1, -607, 32767}
    v9 = append(v9, 1)
    for fi10 := 0; fi10 < 7; fi10++ {
        p0 -= 18446744073709551615
        if len(v9) > 0 {
            v9[0] = fi10
        }
        _ = fi10
    }
    _ = p0
    _ = v9
    return (p0 ^ (p0 | (p0 | 9)))
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 uint8 = 2
    v6 := make(map[uint32]string)
    var v11 uint = 1
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (fi12 ^ fi12) {
        case 0:
            for fi13 := 0; fi13 < 1; fi13++ {
                v11 = (v11 >> ((genFn1((v11 * genFn1(genFn1((v11 - v11)))))) & 63))
                _ = fi13
            }
        default:
            for fi14 := 0; fi14 < 13; fi14++ {
                if false {
                    var v15 bool = (false || (true && (!false)))
                    var v16 uint32 = 395
                    v17 := GenS1{F0: true, F1: 0}
                    var v18 int64 = 0
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi14
            }
        }
        _ = fi12
    }
    var v19 uint = (v11 << ((v11) & 63))
    v1 = (v1 >> ((10) & 63))
    for fi20 := 0; fi20 < 2; fi20++ {
        if false {
            fmt.Println(v5, v19)
            var v21 float32 = float32(1e10)
            _ = v21
        } else {
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint64 = 2
                if true {
                    v24 := GenS1{F0: (1 == 65535), F1: 32767}
                    if ((v5 ^ 255) == (v5 | (v5 * 34))) {
                        v11 = v19
                    }
                    v23 = genFn0(2147483647, v1)
                    _ = v24
                }
                v23 = v23
                _ = fi22
                _ = v23
            }
            var v25 float32 = 3.14159
            v25 = 1e10
            _ = v25
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int8 = -128
            v27 = (v27 - v1)
            _ = fi26
            _ = v27
        }
        v28 := GenS1{F0: (float32(float64(3.14159)) == 0.0), F1: 69}
        _ = fi20
        _ = v28
    }
    var v29 uint = 0
    for fi30 := 0; fi30 < 3; fi30++ {
        if true {
            v19 += (v19 % ((v11) | 1))
        }
        var v31 uint16 = 967
        fmt.Println(v1, v19, v19)
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: 42, F1: 42, F2: 1}
    fmt.Println(v11, v5)
    var v33 int64 = 0
    v34 := v2[2147483647]
    var v35 uint = 299
    switch "a" {
    case "":
        v34 = v34
    case "\n\t":
        var v36 int64 = 174
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
