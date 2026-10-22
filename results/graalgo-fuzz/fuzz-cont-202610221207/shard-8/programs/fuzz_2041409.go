package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int, p1 uint) uint8 {
    var v4 int8 = -11
    var v5 float32 = (1e10 + float32(0.1))
    var v6 int8 = (v4 * v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 153
}

func main() {
    var v1 uint = 950
    v2 := []string{(("\n\t" + ("" + ("\n\t" + "a"))) + ""), ("foo bar" + "\n\t"), "a"}
    var v3 uint32 = 474
    fmt.Println(v3, v1)
    fmt.Println(v1)
    if (1e10 != ((1e10 * (-2.25 + 1.5)) * -2.25)) {
        v7 := GenS1{F0: 42, F1: 846, F2: 255}
        v3 = (v3 / (((v3 << ((v1) & 63))) | 1))
        var v8 float64 = 0.0
        _ = v7
        _ = v8
    }
    fmt.Println(v1)
    fmt.Println(v3, v3)
    v3 = (v3 << (((v1 % (((^(v1))) | 1))) & 63))
    v3 = 2147483648
    if (true || false) {
        fmt.Println(v3, v1)
        v3 += (v3 | v3)
    }
    switch 2147483647 {
    case -2147483648:
        v9 := GenS1{F0: -353, F1: v1, F2: 0}
        _ = v9
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 0:
        fmt.Println(v3, v1)
    default:
        v10 := GenS2{F0: 2147483647}
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
