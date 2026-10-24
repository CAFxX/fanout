package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int16 = 32767
    v2 := []uint32{1, 2147483648, 2, 2}
    v3 := []string{"\n\t", "foo bar"}
    var v4 float32 = 1.5
    var v5 uint = 65535
    var v6 float32 = v4
    fmt.Println(v5, v4, v6)
    var v7 int16 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS1{F0: 211, F1: 1}
    _ = len(v3)
    var v9 bool = ("\n\t" == "foo bar")
    v5 += v5
    var v10 uint32 = 199
    fmt.Println(v7, v10)
    fmt.Println(v10, v4, v10)
    v5 = (v5 - 18446744073709551615)
    v10 -= 2147483648
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 float64 = float64(v4)
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
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
