package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }
func genFn0() int {
    for fi6 := 0; fi6 < 17; fi6++ {
        if (float64((((1.5 - 3.14159) + 1.5) * ((((0.1 - 0.1) + 1.5) / 3.14159) + float32(0.1)))) == 3.14159) {
            var v7 bool = (42 != 32767)
            v7 = (0.0 <= float32((1e10 / 1.0)))
            _ = v7
        } else {
            var v8 int16 = -32768
            if true {
                v8 += v8
                v9 := GenS1{F0: 128, F1: 0}
                _ = v9
            }
            v10 := []int{fi6}
            _ = v10
            _ = v8
        }
        fmt.Println(fi6)
        _ = fi6
    }
    var v11 int8 = 42
    _ = v11
    return 128
}
func genFn1(p0 string, p1 uint32, p2 bool) uint16 {
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int32 = 42
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        p2 = (p2 && p2)
        _ = fi14
    }
    if p2 {
        var v15 uint64 = 574
        v16 := GenS2{F0: 32767, F1: 42, F2: 0}
        var v17 uint = 174
        _ = v15
        _ = v16
        _ = v17
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 4294967295
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v5 := []string{"a", (("" + "\n\t") + ("a" + ("\n\t" + ""))), "foo bar", "foo bar"}
    var v18 uint32 = 42
    v18 = (v18 ^ v18)
    var v19 uint8 = 127
    v20 := GenS1{F0: v19, F1: 1}
    var v21 int = 2147483647
    v19 -= (v19 ^ 61)
    var v22 int8 = -128
    delete(v4, 42)
    switch (v21 >> ((1) & 63)) {
    case 767:
        var v23 uint16 = 65535
        _ = v23
    }
    v24 := GenS1{F0: (v19 << ((4294967295) & 63)), F1: (^(v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
