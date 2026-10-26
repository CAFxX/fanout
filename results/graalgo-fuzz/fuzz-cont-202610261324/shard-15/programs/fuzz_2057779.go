package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar"}
    var v4 uint16 = 32767
    v5 := []int{-2147483648, 2147483647, 764, 32768}
    v6 := GenS1{F0: 0, F1: 127, F2: "a"}
    v7 := GenS1{F0: 0, F1: 1, F2: "\n\t"}
    fmt.Println(v4)
    v4 += (v4 - v4)
    v4 = 554
    switch "foo bar" {
    case "\n\t":
        var v8 float32 = 1.5
        _ = v8
    }
    if true {
        var v9 uint32 = 65535
        fmt.Println(v9, v4)
        _ = v9
    } else {
        var v10 bool = false
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = v10
    }
    fmt.Println(v4)
    v3 = append(v3, "\n\t")
    v11 := GenS1{F0: 4294967295, F1: 128, F2: ""}
    fmt.Println(v4)
    v4 += (v4 << ((799) & 63))
    v4 -= v4
    v4 -= v4
    v12 := v2[""]
    v12 += 600
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
