package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{812}
    var v3 string = "foo bar"
    var v4 int32 = 843
    var v5 uint = 18446744073709551615
    var v6 uint16 = 842
    v7 := GenS1{F0: -9223372036854775808}
    v5 += 0
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v2 = append(v2, (fi8 << (((v5 & 907)) & 63)))
        fmt.Println(v3, v4, fi8)
        var v9 int = (^(fi8))
        _ = fi8
        _ = v9
    }
    v4 += int32(v6)
    v3 = ((((v3 + v3) + "foo bar") + "\n\t") + v3)
    v5 -= 18446744073709551615
    v1[127] = true
    v3 = (v3 + ("" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
