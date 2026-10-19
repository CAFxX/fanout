package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8) int {
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 += p0
        var v5 float32 = 3.14159
        _ = fi4
        _ = v5
    }
    p0 = (p0 ^ 246)
    _ = p0
    return 4294967295
}
func genFn1(p0 int, p1 uint8) int8 {
    var v6 uint16 = 255
    fmt.Println(p0, p0)
    fmt.Println(p0, p0, v6)
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 65535
    v3 := []int{32768, 439}
    switch 128 {
    case -2147483648:
        if ((1.5 != (0.0 * 1.5)) || true) {
            v2 = (v2 & (v2 / ((0) | 1)))
            v2 -= (v2 & v2)
            v7 := GenS1{F0: 42}
            _ = v7
        } else {
            fmt.Println(v2)
        }
    case 0:
        fmt.Println(v2)
    default:
        v2 = v2
    }
    var v8 uint32 = (v2 << ((2) & 63))
    v2 = v2
    v2 = (v2 << ((4294967295) & 63))
    var v9 int16 = 32767
    v9 = (v9 + -950)
    if (v9 == (v9 / ((v9) | 1))) {
        fmt.Println(v9, v9)
        v8 = (v8 | (v8 >> ((267) & 63)))
        var v10 bool = false
        _ = v10
    } else {
        var v11 uint64 = 42
        var v12 uint64 = v11
        _ = v11
        _ = v12
    }
    fmt.Println(v2)
    switch 276 {
    case 32768:
        v2 = v2
    case 4294967295:
        switch ("foo bar" + (((("\n\t" + "\n\t") + "foo bar") + "hello") + "\n\t")) {
        case "a":
            if false {
                var v13 uint = 120
                v2 -= 4294967295
                _ = v13
            }
        }
    case -1:
        v14 := GenS1{F0: 32767}
        _ = v14
    default:
        var v15 string = ""
        _ = v15
    }
    fmt.Println(v9, v8)
    var v16 uint32 = v2
    v2 += (v2 << ((169) & 63))
    v17 := GenS2{F0: -32768}
    fmt.Println(v9, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
