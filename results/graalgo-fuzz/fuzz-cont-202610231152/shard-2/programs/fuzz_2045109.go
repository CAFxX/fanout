package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }

func main() {
    var v1 float64 = (-2.25 + 3.14159)
    v2 := []string{"hello", ("hello" + ("\n\t" + ""))}
    var v3 float64 = ((v1 - 1e10) / ((v1 / 1e10) * v1))
    v4 := make(map[string]bool)
    var v5 int64 = -9223372036854775808
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 787, F1: 0, F2: (v5 << ((487) & 63))}
    _ = len(v2)
    v5 += v5
    if ((-128 > -124) && (1 <= -1)) {
        delete(v6, 0)
    } else {
        _ = len(v4)
    }
    var v8 uint16 = 2
    var v9 int8 = 42
    v5 += -9223372036854775808
    var v10 string = "foo bar"
    var v11 bool = false
    var v12 int = -677
    v13 := GenS1{F0: 42, F1: 2, F2: (v5 >> ((82) & 63))}
    fmt.Println(v10, v12)
    _ = len(v6)
    var v14 string = (v10 + ("a" + v10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
