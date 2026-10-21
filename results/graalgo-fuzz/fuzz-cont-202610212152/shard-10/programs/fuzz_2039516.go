package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }
func genFn0(p0 int16) uint16 {
    fmt.Println(p0)
    p0 += p0
    fmt.Println(p0)
    p0 -= 42
    var v4 uint = 18446744073709551615
    _ = p0
    _ = v4
    return 42
}

func main() {
    v1 := []string{"", "\n\t"}
    var v2 int = 32768
    var v3 int64 = -382
    if true {
        if len(v1) > 0 {
            v1[0] = ("a" + "hello")
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "")
        }
    }
    v5 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: -1}
    v2 = -9223372036854775808
    fmt.Println(v2, v2)
    var v6 float64 = (0.1 * ((0.0 + 0.0) / 0.1))
    var v7 uint32 = 2147483647
    var v8 int16 = 0
    fmt.Println(v3, v3, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
