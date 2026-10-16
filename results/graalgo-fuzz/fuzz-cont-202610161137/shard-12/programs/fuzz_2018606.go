package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int }
type GenS2 struct { F0 uint; F1 int32; F2 string }
func genFn0(p0 int) string {
    p0 -= (p0 >> ((387) & 63))
    v4 := []int{(p0 / ((254) | 1)), 553}
    fmt.Println(p0)
    v5 := GenS1{F0: 0, F1: -1, F2: p0}
    _ = p0
    _ = v4
    _ = v5
    return ("hello" + ("\n\t" + ("hello" + "\n\t")))
}
func genFn1(p0 int, p1 uint, p2 uint8) int {
    switch 1 {
    case 4294967295:
        for fi6 := 0; fi6 < 17; fi6++ {
            p2 = (p2 * uint8(p1))
            _ = fi6
        }
    case -2147483648:
        for fi7 := 0; fi7 < 13; fi7++ {
            p1 = p1
            _ = fi7
        }
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            p1 -= p1
            v9 := GenS2{F0: (p1 * (^(p1))), F1: 0, F2: "a"}
            var v10 int8 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    fmt.Println(p2)
    var v11 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p0 << (((p1 - 821)) & 63))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483648
    var v3 int64 = 358
    v2 = v2
    v3 += 176
    fmt.Println(v3, v2)
    v3 = (v3 ^ v3)
    var v12 bool = ((genFn1(genFn1(4294967295, 42, 0), 255, 2) == int(v3)) && (93 == 220))
    v13 := []int{}
    v14 := GenS1{F0: 950, F1: 1, F2: genFn1(9223372036854775807, 42, 127)}
    v15 := GenS1{F0: 65535, F1: 123, F2: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
