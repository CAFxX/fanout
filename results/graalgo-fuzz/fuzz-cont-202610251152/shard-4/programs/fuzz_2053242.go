package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint8) uint32 {
    p1 = (p1 & p1)
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{(fi7 - (fi7 & (fi7 * (fi7 / (((fi7 ^ (-(fi7)))) | 1))))), (fi7 % ((fi7) | 1)), (fi7 >> ((89) & 63)), (^(fi7))}
        _ = fi7
        _ = v8
    }
    var v9 float32 = (-2.25 - 0.0)
    v10 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2147483647
}
func genFn1(p0 int32) int {
    v11 := GenS2{F0: true, F1: (29 >= 32767)}
    var v12 float64 = 0.1
    fmt.Println(p0)
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := []int{(fi14 << ((437) & 63)), (fi13 * fi14)}
            _ = fi14
            _ = v15
        }
        fmt.Println(v12, v12)
        _ = fi13
    }
    _ = p0
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 string = (("hello" + "\n\t") + "\n\t")
    var v2 int16 = 0
    var v3 uint16 = 272
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = (v3 << ((255) & 63))
    v6 := make(map[int]uint32)
    fmt.Println(v5, v1, v2)
    switch v1 {
    case "foo bar":
        v6[genFn1(-2147483648)] = 2147483647
    case "\n\t":
        var v16 int32 = -1
        _ = v16
    case "hello":
        fmt.Println(v4)
    default:
        v17 := GenS2{F0: ("a" == ((v1 + ((((((v1 + v1) + (("\n\t" + ((v1 + (v1 + v1)) + v1)) + ("\n\t" + ("hello" + (v1 + v1))))) + v1) + ((v1 + ((v1 + (v1 + (v1 + (("a" + v1) + v1)))) + ((((((v1 + "hello") + (v1 + (("" + "hello") + ((v1 + ("\n\t" + ((v1 + (v1 + ("hello" + "hello"))) + "foo bar"))) + (v1 + v1))))) + (v1 + "hello")) + "hello") + v1) + "a"))) + v1)) + "a") + "foo bar")) + ((v1 + "") + ("\n\t" + v1)))), F1: (128 <= 1)}
        _ = v17
    }
    v5 = v3
    v1 = "foo bar"
    v18 := v6[128]
    switch 739 {
    case 2147483647:
        var v19 uint8 = 218
        _ = v19
    case -2147483648:
        v1 = v1
    }
    v20 := GenS3{F0: 42}
    fmt.Println(v5, v1, v4)
    var v21 int32 = 308
    var v22 int = -2147483648
    if (32767 < (v22 / ((32767) | 1))) {
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS3{F0: 127}
            switch v1 {
            case "foo bar":
                var v25 uint8 = 42
                _ = v25
            case "\n\t":
                v4 = v4
            case "a":
                var v26 uint = 18446744073709551615
                _ = v26
            default:
                var v27 float64 = float64(float32(1.5))
                _ = v27
            }
            v2 += (v2 << ((1) & 63))
            _ = fi23
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
