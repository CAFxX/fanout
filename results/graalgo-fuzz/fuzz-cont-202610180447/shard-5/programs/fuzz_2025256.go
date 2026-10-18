package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []string{"a", "hello", "hello"}
    v2 := []uint32{1, 637, 2147483648, 2147483648}
    v3 := []string{"a", ("\n\t" + ("" + "")), "foo bar", ("a" + "a")}
    var v4 int8 = -1
    v5 := make(map[int]int)
    v4 += (v4 ^ v4)
    fmt.Println(v4)
    var v6 uint16 = 2
    v7 := GenS1{F0: v6, F1: -1}
    _ = len(v2)
    var v8 int16 = -906
    v6 += (v6 >> ((683) & 63))
    var v9 float32 = ((1.5 * float32(1e10)) - 1e10)
    fmt.Println(v9)
    var v10 float64 = float64((v9 * ((v9 - v9) + (1e10 - 1.5))))
    v10 += 3.14159
    for fi11 := 0; fi11 < 18; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi11
    }
    var v12 uint8 = 94
    fmt.Println(v9, v6, v4)
    var v13 int16 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
