package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0(p0 uint, p1 uint16, p2 int16) uint64 {
    fmt.Println(p2, p2, p2)
    var v5 int = 1
    v6 := GenS1{F0: (p1 - p1), F1: (0 != 42)}
    v7 := GenS1{F0: (p1 % ((p1) | 1)), F1: (true && (-2.25 < -2.25))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 774
}
func genFn1() uint8 {
    v8 := []int{234, 32768}
    v8 = append(v8, -9223372036854775808)
    _ = v8
    return 2
}

func main() {
    var v1 uint8 = 183
    var v2 uint32 = 2147483648
    var v3 bool = ((((("\n\t" + "\n\t") + ("foo bar" + "a")) + "a") + "foo bar") >= "hello")
    var v4 int8 = -58
    v9 := []int{42, 2147483647}
    var v10 uint8 = 255
    for fi11 := 0; fi11 < 13; fi11++ {
        v9 = append(v9, fi11)
        _ = fi11
    }
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            v4 = (v4 << ((0) & 63))
            var v14 int16 = -16
            _ = fi13
            _ = v14
        }
        var v15 int = 32768
        _ = fi12
        _ = v15
    }
    v3 = ("a" >= "\n\t")
    fmt.Println(v3, v3, v1)
    var v16 int = 4294967295
    if (!(0 <= 255)) {
        v1 = 42
        v4 = (v4 % ((42) | 1))
        v1 -= (^(v1))
    }
    v17 := GenS1{F0: 0, F1: v3}
    var v18 string = ("foo bar" + ("\n\t" + ""))
    v9 = append(v9, -53)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
