package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 string) int {
    var v5 uint32 = 0
    if (1 == 127) {
        switch 2147483647 {
        case 2147483647:
            p0 = p0
        }
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint8 = 1
            v8 := GenS1{F0: uint8(v5)}
            v7 = (v7 << ((782) & 63))
            _ = fi6
            _ = v7
            _ = v8
        }
        if (true || true) {
            var v9 uint = 28
            var v10 int16 = 32767
            var v11 int32 = 2147483647
            _ = v10
            _ = v11
            _ = v9
        }
    } else {
        v12 := GenS1{F0: 249}
        fmt.Println(v5, p0)
        _ = v12
    }
    fmt.Println(v5, v5)
    fmt.Println(p0, v5)
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 float64 = 1.5
    var v2 float64 = -2.25
    v3 := make(map[int]bool)
    var v4 int8 = 36
    if (false || (9223372036854775807 > -9223372036854775808)) {
        var v13 int64 = 42
        _ = v13
    }
    v2 -= (v2 - 0.0)
    for fi14 := 0; fi14 < 11; fi14++ {
        v2 += float64((0.1 + (1e10 + 3.14159)))
        if (((6 == 0) && false) && (4294967295 < 529)) {
            v1 -= float64((-2.25 + 0.1))
        }
        _ = fi14
    }
    switch 127 {
    case -409:
        v15 := []int{-249, genFn0(("" + "hello")), 94, 127}
        _ = v15
    default:
        switch ("\n\t" + "hello") {
        case "":
            v16 := []int{-2147483648, genFn0("a"), 0}
            _ = v16
        case "hello":
            v1 += (v1 / (v2 / 1e10))
        case "\n\t":
            v17 := GenS1{F0: 0}
            _ = v17
        }
    }
    var v18 uint64 = 18446744073709551615
    var v19 string = "hello"
    fmt.Println(v4)
    var v20 string = v19
    v2 -= v1
    v18 = v18
    switch "" {
    case "foo bar":
        var v21 int32 = 1
        _ = v21
    }
    fmt.Println(v1, v2, v19)
    if (float32(v2) < (0.1 * (1e10 + ((-2.25 - 1.5) / (1e10 / 3.14159))))) {
        v22 := v3[303]
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
