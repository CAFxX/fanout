package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -1
    var v3 int64 = 9223372036854775807
    var v4 bool = (128 != 46)
    switch 42 {
    case 951:
        v2 -= v2
    }
    v4 = (!false)
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 14; fi5++ {
        v3 -= (v3 << ((529) & 63))
        var v6 uint32 = 0
        _ = fi5
        _ = v6
    }
    if (-47 > -1) {
        fmt.Println(v3, v3)
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: -1, F1: 42, F2: -1}
            _ = fi7
            _ = v8
        }
    } else {
        v1["foo bar"] = 42
        fmt.Println(v3, v3)
    }
    if v4 {
        fmt.Println(v2)
        v4 = v4
        var v9 uint8 = 128
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: -12, F1: 0, F2: v2}
        v12 := []int{fi10, 4294967295}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint16 = 210
    _ = len(v1)
    v3 = v3
    v4 = (!(!(v4 && (!((!v4) || v4)))))
    v14 := []int{32767}
    switch "\n\t" {
    case "\n\t":
        v15 := GenS1{F0: 100, F1: -115, F2: v2}
        _ = v15
    case "foo bar":
        v3 = 648
    case "a":
        var v16 uint32 = 42
        _ = v16
    }
    v17 := GenS2{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
