package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 int16 }
func genFn0(p0 uint32, p1 uint8, p2 int8) int8 {
    p1 = (p1 % ((p1) | 1))
    p1 -= 127
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            if false {
                if true {
                    v7 := []int{}
                    _ = v7
                }
                var v8 int8 = p2
                v9 := []int{fi6, fi6, 9223372036854775807}
                _ = v8
                _ = v9
            } else {
                p0 = (p0 << ((uint(p0)) & 63))
            }
            switch "hello" {
            case "a":
                var v10 uint = 149
                v11 := GenS1{F0: 18446744073709551615, F1: 515}
                p2 -= p2
                var v12 uint8 = 128
                _ = v10
                _ = v11
                _ = v12
            case "":
                var v13 int8 = (p2 ^ -128)
                _ = v13
            }
            _ = fi6
        }
        _ = fi5
    }
    v14 := []int{32768, -1, 128, -571}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return p2
}

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]bool)
    var v3 int64 = 217
    var v4 int = 32767
    switch ("a" + "a") {
    case "a":
        v15 := []int{}
        _ = v15
    case "":
        switch "a" {
        case "":
            var v16 uint16 = 42
            _ = v16
        }
    case "foo bar":
        v17 := GenS1{F0: 2, F1: 1000}
        _ = v17
    default:
        v18 := []int{}
        _ = v18
    }
    switch "a" {
    case "foo bar":
        v19 := v2[0]
        _ = v19
    case "a":
        var v20 float32 = float32((1e10 - 1e10))
        _ = v20
    case "hello":
        var v21 int64 = 9223372036854775807
        _ = v21
    }
    v4 = (v4 << ((0) & 63))
    switch v4 {
    case -2147483648:
        delete(v2, 544)
    default:
        v22 := []int{-800, (v4 / ((v4) | 1)), v4}
        _ = v22
    }
    var v23 uint64 = 1
    v2[2147483647] = (((((((0.1 + ((1.5 + (float32(0.1) - -2.25)) * 3.14159)) - (((0.0 - 3.14159) / 1.5) / (3.14159 + 3.14159))) * 0.1) * (1.5 - 3.14159)) - -2.25) > 1.5) && ((true || true) || true))
    var v24 int8 = v1
    if true {
        v25 := GenS2{F0: -684, F1: 1}
        v3 = (-(v3))
        _ = v25
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 int64 = v3
            _ = fi26
            _ = v27
        }
        var v28 float64 = 0.1
        _ = v28
    }
    v29 := GenS2{F0: 1, F1: 2}
    fmt.Println(v3, v3, v24)
    v24 += int8(v4)
    v30 := GenS1{F0: 677, F1: 362}
    for fi31 := 0; fi31 < 10; fi31++ {
        if (false || (!true)) {
            v4 = v4
            if false {
                v32 := v2[42]
                _ = v32
            } else {
                var v33 uint8 = 127
                _ = v33
            }
        } else {
            for fi34 := 0; fi34 < 6; fi34++ {
                v35 := GenS3{F0: (fi31 - -300), F1: (v24 >> ((0) & 63)), F2: -1}
                v24 = (v24 + (-(v24)))
                _ = fi34
                _ = v35
            }
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
