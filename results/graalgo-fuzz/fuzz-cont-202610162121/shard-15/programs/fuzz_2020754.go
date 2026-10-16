package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint32, p1 int8, p2 string) int64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 17; fi7++ {
            p1 += (p1 / ((p1) | 1))
            _ = fi7
        }
        _ = fi6
    }
    v8 := []int{127}
    var v9 uint = 65535
    p2 = p2
    if len(v8) > 0 {
        v8[0] = 42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -905
}

func main() {
    v1 := []string{"hello", ((("a" + "a") + ("" + "hello")) + ("foo bar" + "\n\t")), ("" + "hello"), "hello"}
    var v2 float32 = float32(1.5)
    var v3 uint = 1
    var v4 float64 = 3.14159
    var v5 int8 = 17
    fmt.Println(v3, v3)
    var v10 uint8 = 255
    var v11 float64 = 3.14159
    var v12 int = 920
    switch v12 {
    case -2:
        v4 = v11
    case 0:
        if len(v1) > 0 {
            v1[0] = (("a" + "foo bar") + "foo bar")
        }
    case 32768:
        v13 := GenS1{F0: -2147483648}
        _ = v13
    default:
        _ = len(v1)
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v12, v5)
        _ = len(v1)
        _ = fi14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!true) {
        v15 := GenS1{F0: -562}
        v3 -= (^(v3))
        v4 = float64(1.5)
        _ = v15
    } else {
        if ((true || (18446744073709551615 <= 742)) && ((v3 - v3) == v3)) {
            v3 = 65535
            v11 -= v11
        } else {
            for fi16 := 0; fi16 < 2; fi16++ {
                v10 += v10
                var v17 uint = (v3 & 37)
                _ = fi16
                _ = v17
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v18 := GenS1{F0: 42}
    var v19 int16 = -32768
    fmt.Println(v10, v12, v11)
    if ((v2 + 1.5) >= float32(v11)) {
        var v20 float32 = (-2.25 * float32(v11))
        _ = v20
    }
    var v21 uint64 = 329
    v2 += float32(-2.25)
    v2 -= (v2 - v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
