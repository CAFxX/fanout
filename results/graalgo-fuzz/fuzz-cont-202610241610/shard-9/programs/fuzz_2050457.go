package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool, p1 int64, p2 int8) int64 {
    v5 := []int{128}
    v6 := GenS1{F0: 277, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}

func main() {
    var v1 uint64 = 1
    var v2 int16 = 0
    var v3 int = 32767
    var v4 float64 = 1.5
    var v7 bool = (false || true)
    if (2147483647 >= 2147483647) {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = int64(v2)
            _ = fi8
            _ = v9
        }
        v10 := []int{627, v3, (v3 | v3)}
        _ = v10
    }
    if v7 {
        v11 := GenS1{F0: v1, F1: 671}
        _ = v11
    } else {
        v12 := []int{}
        _ = v12
    }
    v3 += 1
    for fi13 := 0; fi13 < 2; fi13++ {
        switch (v3 | (^(fi13))) {
        case -1:
            var v14 uint8 = 63
            _ = v14
        case 32767:
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int64 = 449
                v17 := []int{-514}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        var v18 uint = 50
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 float64 = (1e10 / v4)
            _ = fi19
            _ = v20
        }
        _ = fi13
        _ = v18
    }
    if (772 <= v1) {
        v21 := []int{v3, (v3 + v3), 1, (v3 >> ((192) & 63))}
        _ = v21
    }
    if (false && (v1 > v1)) {
        v22 := []int{v3, v3}
        v23 := GenS1{F0: 4294967295, F1: 42}
        _ = v22
        _ = v23
    } else {
        if (1 == 255) {
            fmt.Println(v1, v1)
            var v24 uint32 = 2147483648
            _ = v24
        } else {
            v4 -= 0.0
        }
        v2 -= v2
        v1 += (v1 * (v1 / ((42) | 1)))
    }
    v25 := GenS2{F0: -126, F1: (v1 >> ((289) & 63))}
    v4 -= v4
    v26 := GenS2{F0: 1, F1: v1}
    for fi27 := 0; fi27 < 18; fi27++ {
        switch "\n\t" {
        case "":
            for fi28 := 0; fi28 < 20; fi28++ {
                switch "\n\t" {
                case "a":
                    v29 := GenS3{F0: v7, F1: (^(v3))}
                    _ = v29
                case "hello":
                    v4 -= v4
                }
                _ = fi28
            }
        }
        var v30 uint64 = 42
        _ = fi27
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS4{F0: (-(v2))}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        if (float32(v4) <= -2.25) {
            fmt.Println(v7)
        } else {
            v4 = v4
            var v34 float32 = (3.14159 - (-2.25 - (float32(1e10) + 3.14159)))
            if (fi33 == 379) {
                var v35 uint = 42
                _ = v35
            }
            _ = v34
        }
        _ = fi33
    }
    switch -2147483648 {
    case -201:
        var v36 float32 = 3.14159
        _ = v36
    case 127:
        v37 := GenS3{F0: v7, F1: (v3 | (-(v3)))}
        _ = v37
    default:
        var v38 string = ""
        _ = v38
    }
    var v39 uint8 = 198
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v7
}
