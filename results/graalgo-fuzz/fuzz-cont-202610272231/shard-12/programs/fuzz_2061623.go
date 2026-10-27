package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 950
    var v3 float32 = 0.0
    v4 := make(map[string]int)
    var v5 int = int(v2)
    var v6 float64 = (-2.25 - -2.25)
    fmt.Println(v5)
    v7 := GenS1{F0: 876}
    v8 := GenS1{F0: 543}
    var v9 int64 = 42
    var v10 uint16 = 65535
    var v11 string = ("" + "\n\t")
    v4["hello"] = (v5 & -2147483648)
    v6 += v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
