package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []uint32{0, 2147483648, 2}
    v2 := make(map[uint32]string)
    v3 := []string{"\n\t"}
    v4 := make(map[string]string)
    v5 := []int{805, 32768, 42, 480}
    var v6 uint32 = 392
    var v7 float64 = 0.1
    fmt.Println(v7, v6)
    fmt.Println(v7, v6)
    v6 += (v6 | 639)
    var v8 uint32 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS1{F0: -385}
    if len(v3) > 0 {
        v3[0] = ("hello" + ("" + "\n\t"))
    }
    v10 := v4["hello"]
    v6 = (v6 & 295)
    v10 = v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
