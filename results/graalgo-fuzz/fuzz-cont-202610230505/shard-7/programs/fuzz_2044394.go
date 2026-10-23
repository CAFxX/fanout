package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
func genFn0(p0 string, p1 string) int16 {
    v5 := []int{32767, 128}
    var v6 string = p0
    if false {
        if len(v5) > 0 {
            v5[0] = 128
        }
        fmt.Println(p0, p1)
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
    } else {
        _ = len(v5)
        var v7 uint = 18446744073709551615
        _ = v7
    }
    switch "" {
    case "hello":
        var v8 bool = (42 > 126)
        _ = v8
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v9 float32 = (-2.25 + 0.1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return -297
}

func main() {
    var v1 uint = 2
    v2 := []string{("" + "foo bar"), "\n\t", "hello", ("foo bar" + "\n\t")}
    v3 := []uint32{65535, 0, 0, 295}
    var v4 uint = (v1 ^ 42)
    fmt.Println(v4, v1)
    var v10 bool = false
    var v11 uint8 = 0
    switch 2147483647 {
    case -629:
        if len(v2) > 0 {
            v2[0] = ("a" + "hello")
        }
    case 2147483647:
        v10 = false
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 float64 = (((0.0 - 3.14159) + (float64(-2.25) + 0.1)) * 1e10)
            _ = fi12
            _ = v13
        }
    }
    v10 = (!(-32768 > int16(v4)))
    fmt.Println(v10)
    v4 -= 675
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 int64 = 9223372036854775807
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v11 -= v11
    v17 := GenS1{F0: genFn0("a", "hello"), F1: 2147483647}
    v11 -= (v11 | (v11 / (((v11 ^ (v11 * (v11 >> (((v4 | (v1 - (v4 / ((18446744073709551615) | 1))))) & 63))))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
