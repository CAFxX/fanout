package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint8 = 168
    var v2 uint64 = 0
    var v3 int32 = -1
    var v4 int8 = 42
    var v5 int32 = -312
    v6 := make(map[int]uint32)
    delete(v6, 32768)
    if (true || false) {
        v7 := GenS1{F0: v3}
        if (("" + ("hello" + "hello")) > (("foo bar" + ("hello" + "")) + ("" + ("a" + "")))) {
            fmt.Println(v5)
        }
        v8 := GenS2{F0: (v3 ^ (v3 << ((1) & 63))), F1: false, F2: -1}
        _ = v7
        _ = v8
    }
    var v9 uint = 4294967295
    var v10 uint32 = 4294967295
    v11 := GenS3{F0: (v9 - v9), F1: (-(v4))}
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS4{F0: 42}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        if (v9 == uint(v2)) {
            v3 += 611
            var v15 int16 = 454
            _ = v15
        }
        if true {
            var v16 bool = false
            v17 := []int{fi14, fi14, fi14, fi14}
            var v18 float64 = float64(((1e10 * 1e10) * ((((1.5 - ((float32((3.14159 - (0.1 * 0.1))) * ((float32(0.1) / 1e10) + -2.25)) - 3.14159)) + float32(0.0)) * -2.25) - (float32((0.0 + 3.14159)) - (0.0 / 1.0)))))
            _ = v16
            _ = v17
            _ = v18
        } else {
            delete(v6, (fi14 ^ (fi14 - 127)))
            fmt.Println(v4, v1)
            _ = len(v6)
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int32 = v5
        var v21 uint8 = v1
        var v22 string = "foo bar"
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: (v5 & 0), F1: (false || true), F2: -128}
    fmt.Println(v4, v9)
    v2 += v2
    var v24 uint8 = (v1 / ((v1) | 1))
    v25 := []int{42, -259}
    for fi26 := 0; fi26 < 11; fi26++ {
        if (v4 < (v4 / (((v4 << (((v9 % ((v9) | 1))) & 63))) | 1))) {
            v10 += v10
            _ = len(v6)
            v27 := GenS5{F0: 9223372036854775807}
            _ = v27
        }
        _ = fi26
    }
    switch 42 {
    case 296:
        switch "foo bar" {
        case "hello":
            if len(v25) > 0 {
                v25[0] = 32768
            }
        default:
            fmt.Println(v9, v10)
        }
    }
    fmt.Println(v5, v24, v2)
    v9 -= (v9 >> (((^(v9))) & 63))
    v1 = (v24 + 2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
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
