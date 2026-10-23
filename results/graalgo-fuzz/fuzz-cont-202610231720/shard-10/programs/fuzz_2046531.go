package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0(p0 uint32, p1 uint64, p2 uint16) string {
    var v5 int16 = -32768
    var v6 bool = (4294967295 != 42)
    v7 := GenS1{F0: -1, F1: ""}
    p0 = (p0 - 1)
    var v8 uint16 = uint16(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ("\n\t" + "\n\t")
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 65535
    var v3 float64 = float64((1.5 + ((1e10 / 1.0) * 0.0)))
    var v4 uint8 = 255
    _ = len(v1)
    var v9 uint = 417
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 9; fi11++ {
        v3 = 0.1
        _ = fi11
    }
    switch "" {
    case "foo bar":
        if (-756 <= -819) {
            fmt.Println(v9, v9, v9)
            v12 := []int{-880, 32768, 4294967295, -2147483648}
            _ = v12
        }
    }
    var v13 int64 = 58
    var v14 uint = v9
    fmt.Println(v3, v10)
    v4 = 0
    var v15 float64 = ((3.14159 * (float64(1.5) + v3)) * v3)
    if true {
        v16 := GenS1{F0: 0, F1: "a"}
        _ = v16
    } else {
        if (255 != (v9 + (v9 - 255))) {
            v2 -= v2
        } else {
            fmt.Println(v9, v14)
        }
    }
    delete(v1, -9223372036854775808)
    _ = len(v1)
    var v17 int = 127
    switch 2147483647 {
    case 42:
        v17 -= (v17 << ((2) & 63))
    default:
        v18 := GenS1{F0: 0, F1: ("a" + ("" + ("" + "\n\t")))}
        _ = v18
    }
    var v19 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
