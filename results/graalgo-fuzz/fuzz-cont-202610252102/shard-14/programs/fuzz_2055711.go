package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 int64 }
func genFn0(p0 int32, p1 int, p2 uint64) uint32 {
    var v4 int32 = (p0 | p0)
    p2 += p2
    v5 := []int{(p1 | -9223372036854775808)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint64) uint32 {
    fmt.Println(p0)
    var v6 float32 = 1e10
    v7 := GenS1{F0: 471}
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 bool = ("hello" > "")
    var v2 uint = 18446744073709551615
    var v3 int8 = 127
    v9 := GenS2{F0: (-(v3)), F1: v1, F2: 9223372036854775807}
    v3 += (v3 * (v3 ^ (v3 & -128)))
    v10 := []int{32768}
    v10 = append(v10, 128)
    v3 = (v3 << (((v2 ^ (v2 + v2))) & 63))
    v2 = 18446744073709551615
    var v11 string = ("\n\t" + (("a" + "foo bar") + ("hello" + "")))
    for fi12 := 0; fi12 < 12; fi12++ {
        if (!true) {
            fmt.Println(v3, v3)
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int = -9223372036854775808
                _ = fi13
                _ = v14
            }
            _ = len(v10)
        } else {
            switch v11 {
            case "":
                v15 := GenS1{F0: 65535}
                _ = v15
            case "\n\t":
                v10 = append(v10, -9223372036854775808)
            default:
                _ = len(v10)
            }
            fmt.Println(v3)
            v10 = append(v10, 0)
        }
        v3 += (-(v3))
        v2 -= (v2 << (((v2 & (v2 * (v2 - (v2 - v2))))) & 63))
        _ = fi12
    }
    if (true || true) {
        v11 = "hello"
        v2 = v2
        var v16 uint32 = 4294967295
        _ = v16
    }
    fmt.Println(v1, v1, v2)
    fmt.Println(v3, v1)
    v3 -= -128
    v2 -= v2
    if ((!(63 > 1000)) || v1) {
        v3 = v3
        fmt.Println(v11, v1)
    } else {
        var v17 uint16 = 42
        for fi18 := 0; fi18 < 3; fi18++ {
            _ = len(v10)
            var v19 bool = ("\n\t" <= v11)
            v2 = v2
            _ = fi18
            _ = v19
        }
        var v20 uint32 = 4294967295
        _ = v17
        _ = v20
    }
    var v21 int = 230
    fmt.Println(v21, v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
