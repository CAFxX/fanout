package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8) uint64 {
    fmt.Println(p0)
    p0 = p0
    p0 -= -128
    v5 := []int{4294967295, 0}
    p0 += (p0 << ((1) & 63))
    _ = p0
    _ = v5
    return 289
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = -2.25
    var v3 uint64 = 42
    v4 := []string{"", ""}
    v3 -= 541
    var v6 uint32 = 4294967295
    if false {
        var v7 int64 = -435
        fmt.Println(v3, v6, v3)
        v7 -= -485
        _ = v7
    } else {
        var v8 uint = uint(v3)
        _ = v8
    }
    fmt.Println(v6)
    if (!false) {
        var v9 uint = 1
        var v10 int8 = 100
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi11
        }
        v12 := GenS1{F0: 1, F1: 265, F2: ("" + ("\n\t" + "\n\t"))}
        v6 -= (v6 | 2147483647)
        _ = v12
    }
    v2 = (-2.25 / (0.1 - v2))
    _ = len(v1)
    if (9223372036854775807 >= -9223372036854775808) {
        if (86 <= 1) {
            v3 += (v3 * (v3 ^ (^(v3))))
            var v13 int32 = 2147483647
            _ = v13
        }
        if (!false) {
            v1["a"] = ("hello" + "\n\t")
            v6 = v6
        } else {
            v3 = 4294967295
            v2 = float32((3.14159 - (float64((v2 + v2)) * 1e10)))
        }
        fmt.Println(v2, v6)
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS2{F0: 1000}
            _ = fi14
            _ = v15
        }
        var v16 int32 = -2147483648
        _ = v16
    }
    v17 := v1["hello"]
    if (!(((v17 + (v17 + v17)) + (v17 + v17)) != v17)) {
        var v18 uint = 255
        var v19 float32 = ((((v2 * ((float32(((0.0 - float64(v2)) + 0.0)) + (v2 + v2)) / float32(1.5))) / (v2 + 1e10)) + (0.0 / v2)) - v2)
        delete(v1, "hello")
        _ = v18
        _ = v19
    }
    switch "" {
    case "":
        fmt.Println(v3)
    case "hello":
        v6 -= v6
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
