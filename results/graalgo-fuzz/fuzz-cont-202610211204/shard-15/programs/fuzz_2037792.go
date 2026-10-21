package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 32767
    var v3 int = 4294967295
    var v4 int64 = (v1 * v1)
    var v5 float64 = (-2.25 * -2.25)
    v6 := make(map[string]string)
    v4 = 42
    v1 += (v4 << ((65535) & 63))
    var v7 uint16 = 65535
    var v8 bool = true
    if v8 {
        v9 := []int{127}
        _ = v9
    }
    fmt.Println(v8)
    delete(v6, "a")
    v10 := GenS1{F0: 42}
    v11 := GenS1{F0: -1}
    v12 := GenS2{F0: v3}
    v13 := v6["\n\t"]
    v1 = 42
    v14 := []int{(v3 * v3), (v3 | v3)}
    switch (("a" + v13) + "") {
    case "foo bar":
        v15 := v6[v13]
        _ = v15
    default:
        delete(v6, (v13 + v13))
    }
    v16 := GenS1{F0: (^(v2))}
    var v17 int16 = (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
