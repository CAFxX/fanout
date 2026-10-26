package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{128, 2147483647}
    var v2 int16 = 0
    v3 := []uint32{0, 2147483648}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v2 += (v2 | v2)
    }
    v2 -= 0
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(v2)
        v2 = (v2 | v2)
        _ = fi4
    }
    v2 -= (v2 >> ((255) & 63))
    _ = len(v1)
    v1 = append(v1, 590)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        for fi5 := 0; fi5 < 13; fi5++ {
            if (42 < 42) {
                for fi6 := 0; fi6 < 15; fi6++ {
                    _ = len(v1)
                    _ = fi6
                }
            }
            v7 := GenS1{F0: 2}
            _ = fi5
            _ = v7
        }
    }
    if true {
        var v8 int64 = -1
        _ = v8
    }
    v9 := GenS2{F0: -2147483648}
    v2 = v2
    var v10 int32 = -684
    if true {
        fmt.Println(v2, v2)
    }
    fmt.Println(v10, v10)
    var v11 uint32 = 1
    switch "\n\t" {
    case "a":
        fmt.Println(v2)
    case "":
        v3 = append(v3, 0)
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int8 = -114
            fmt.Println(fi13, v10, fi12)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v2 = (v2 << ((42) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
