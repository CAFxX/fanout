package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint32; F1 int }
func genFn0(p0 uint8) int8 {
    v7 := []int{32767}
    if (true && (p0 > p0)) {
        p0 = 174
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v8 float64 = 3.14159
        _ = v8
    } else {
        p0 += p0
    }
    var v9 float64 = -2.25
    var v10 uint = 42
    _ = p0
    _ = v10
    _ = v7
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    var v4 uint8 = 0
    v5 := make(map[string]int)
    var v6 int8 = 0
    switch 0 {
    case -283:
        v11 := GenS1{F0: ((v6 == v6) || true), F1: 9223372036854775807, F2: 1}
        _ = v11
    case 9223372036854775807:
        delete(v1, 65535)
    }
    v6 = genFn0(v4)
    v6 += (v6 & (v6 / ((v6) | 1)))
    v6 -= v6
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 float64 = 1.5
        _ = fi12
        _ = v13
    }
    var v14 float64 = -2.25
    v6 += 1
    var v15 int = -1
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(v14, v15)
        var v17 int16 = -1
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 4294967295, F1: (v4 < (^(v4)))}
    switch (v15 - v15) {
    case -911:
        if false {
            v19 := GenS3{F0: 42, F1: v15}
            v2 = append(v2, 4294967295)
            _ = v19
        } else {
            switch "" {
            case "hello":
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
            }
            var v20 uint16 = 32767
            _ = len(v3)
            _ = v20
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
