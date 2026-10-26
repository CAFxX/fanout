package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{222, 2147483647, 65535, 1}
    var v2 uint64 = 0
    var v3 int8 = -59
    var v4 int32 = 1
    v5 := make(map[string]int)
    var v6 uint16 = 1
    v4 -= (v4 & -2147483648)
    fmt.Println(v2)
    fmt.Println(v3, v4)
    switch -2147483648 {
    case 128:
        fmt.Println(v6, v4, v6)
    default:
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 bool = false
            var v9 int32 = v4
            v10 := GenS1{F0: 0}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    }
    v11 := GenS1{F0: (v2 / (((^(v2))) | 1))}
    v2 -= v2
    if true {
        for fi12 := 0; fi12 < 16; fi12++ {
            v4 = v4
            v6 -= v6
            _ = fi12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 int = fi14
                fmt.Println(v3, v2)
                v15 = v15
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        v6 = v6
    }
    if true {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := v5[("hello" + ((("\n\t" + "\n\t") + "a") + "foo bar"))]
            var v18 uint8 = uint8(v2)
            var v19 uint = 4294967295
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    }
    fmt.Println(v2, v2)
    _ = len(v1)
    switch 42 {
    case 32767:
        v3 = (v3 ^ (-(v3)))
    }
    fmt.Println(v4)
    var v20 float64 = ((0.1 + (3.14159 - 0.1)) - (3.14159 * 1e10))
    v20 += (v20 * v20)
    var v21 uint = 65535
    switch ((("foo bar" + "hello") + ("foo bar" + "")) + "") {
    case "foo bar":
        var v22 int = 128
        _ = v22
    }
    v20 -= v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
