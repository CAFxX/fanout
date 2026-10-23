package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 int64, p2 uint8) int8 {
    var v4 int16 = 376
    var v5 uint = 857
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return int8(v4)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 int32 = -2147483648
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint64 = 4294967295
        switch -9223372036854775808 {
        case 1:
            fmt.Println(v3, fi6)
        case 2147483647:
            switch "hello" {
            case "":
                v8 := []int{(fi6 % ((32768) | 1)), fi6}
                _ = v8
            case "\n\t":
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: 1, F1: 32767}
                    _ = fi9
                    _ = v10
                }
            case "hello":
                v2[(fi6 % (((fi6 >> ((65535) & 63))) | 1))] = 2
            }
        }
        v11 := []int{fi6, fi6, (fi6 >> ((4294967295) & 63)), fi6}
        _ = fi6
        _ = v11
        _ = v7
    }
    var v12 uint = 2
    if (!false) {
        v13 := GenS2{F0: ""}
        fmt.Println(v12)
        _ = v13
    } else {
        v3 = (-(v3))
        v3 -= (v3 % ((v3) | 1))
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 = (v3 << ((v12) & 63))
        v12 -= (v12 & v12)
        v12 = 18446744073709551615
        _ = fi14
    }
    switch "hello" {
    case "":
        fmt.Println(v3, v12)
    case "a":
        v15 := GenS3{F0: false}
        _ = v15
    case "foo bar":
        switch "" {
        case "a":
            var v16 float32 = 3.14159
            _ = v16
        default:
            fmt.Println(v3, v3)
        }
    }
    if (127 >= -375) {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int8 = 1
            v19 := v1[fi17]
            v19 = -2147483648
            _ = fi17
            _ = v18
            _ = v19
        }
        v3 += v3
    }
    if ((((-2.25 * 3.14159) + -2.25) - float64((-2.25 * (1e10 * (-2.25 + 1e10))))) <= -2.25) {
        _ = len(v2)
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS1{F0: 115, F1: 32767}
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    }
    v12 = 2
    fmt.Println(v3)
    var v23 int16 = 726
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
}
