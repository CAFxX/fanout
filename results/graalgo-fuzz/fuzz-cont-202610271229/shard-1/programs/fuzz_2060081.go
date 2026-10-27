package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{}
        var v9 uint8 = 128
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float32 = (0.0 * 0.1)
    var v11 int16 = -32768
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 string = ("foo bar" + ((("\n\t" + ("a" + "foo bar")) + ("" + "foo bar")) + ""))
    v2 := make(map[uint32]bool)
    var v3 int8 = 49
    var v4 int8 = -128
    var v5 bool = true
    var v6 int8 = -128
    _ = len(v2)
    var v12 int = 526
    var v13 int32 = -2147483648
    if true {
        fmt.Println(v3, v4, v5)
        switch (v12 - (v12 * int(v13))) {
        case 2147483647:
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 int8 = -128
                _ = fi14
                _ = v15
            }
        }
    }
    v6 = (^(v3))
    _ = len(v2)
    v16 := GenS1{F0: v13}
    v4 = v6
    if v5 {
        v17 := []int{128}
        if (!(v5 && (18446744073709551615 > 18446744073709551615))) {
            var v18 int32 = v13
            for fi19 := 0; fi19 < 8; fi19++ {
                if (v5 || (v5 && (113 > (-(v4))))) {
                    v13 += v13
                }
                v6 -= v4
                _ = fi19
            }
            _ = v18
        }
        v4 = (-(v6))
        _ = v17
    }
    if v5 {
        var v20 uint32 = 56
        _ = v20
    }
    var v21 float32 = float32(0.1)
    v22 := []int{}
    fmt.Println(v21)
    fmt.Println(v21, v21, v13)
    v23 := GenS2{F0: 4294967295, F1: v4}
    v3 = -80
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
