package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
func genFn0(p0 int32, p1 uint16, p2 string) uint32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: -128, F1: -32768, F2: 42}
        fmt.Println(p2)
        _ = fi7
        _ = v8
    }
    var v9 int32 = p0
    p0 += (v9 >> ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 4294967295
}

func main() {
    v1 := []int{475}
    v2 := []string{}
    var v3 int16 = -731
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    var v6 int32 = -579
    for fi10 := 0; fi10 < 4; fi10++ {
        switch (("\n\t" + ("\n\t" + (("" + ("" + "hello")) + "\n\t"))) + "hello") {
        case "a":
            var v11 float32 = 0.1
            _ = v11
        default:
            v6 = -2147483648
        }
        fmt.Println(v3, v6, fi10)
        v6 += v6
        _ = fi10
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v6 = -2147483648
        _ = fi12
    }
    v3 -= (v3 + -32768)
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v3)
        if (((0.1 * -2.25) + 3.14159) < 0.1) {
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(fi13, fi13, fi14)
                _ = fi14
            }
        }
        var v15 string = "hello"
        _ = fi13
        _ = v15
    }
    var v16 int64 = 42
    _ = len(v2)
    var v17 float64 = 0.1
    var v18 uint8 = 152
    if (false || ((1e10 + (((0.1 - float32(1.5)) + 3.14159) + (0.0 - -2.25))) <= 1e10)) {
        var v19 uint64 = 370
        fmt.Println(v18, v18, v18)
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
