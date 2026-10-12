package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []int{42, 32767, 914}
    var v2 int16 = 42
    v3 := []int{int(v2), 0, 32768, 9223372036854775807}
    var v4 uint32 = 406
    var v5 string = "foo bar"
    v4 += 4294967295
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int32 = 746
        _ = fi6
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS1{F0: 2}
    v9 := GenS1{F0: 127}
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(v2, v2)
        _ = fi10
    }
    switch v5 {
    case "hello":
        v4 = 916
    }
    var v11 uint32 = 539
    var v12 int64 = -9223372036854775808
    v4 -= 0
    var v13 string = v5
    v14 := GenS2{F0: -32768, F1: 304}
    v15 := GenS2{F0: v2, F1: 704}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
