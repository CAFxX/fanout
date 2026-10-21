package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 string = ("" + "\n\t")
    v4 := []string{v3, "\n\t", "", ""}
    v5 := GenS1{F0: 128, F1: -1}
    v6 := GenS2{F0: 0, F1: 9223372036854775807}
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v3, fi7)
        fmt.Println(fi7, fi7)
        v3 = "foo bar"
        _ = fi7
    }
    v8 := GenS2{F0: 0, F1: 1}
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v3 = v3
        v4 = append(v4, (v3 + ("\n\t" + "")))
    }
    var v9 int32 = 906
    var v10 bool = false
    var v11 uint8 = 127
    var v12 int16 = 1000
    v1[2] = 255
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int64 = 9223372036854775807
        _ = fi13
        _ = v14
    }
    if v10 {
        v4 = append(v4, v3)
        var v15 int8 = 30
        _ = v15
    }
    fmt.Println(v11)
    fmt.Println(v12, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
