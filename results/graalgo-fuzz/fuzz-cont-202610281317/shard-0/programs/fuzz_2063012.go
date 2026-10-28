package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{173, 2147483648}
    v3 := make(map[uint32]string)
    v4 := []string{}
    for fi5 := 0; fi5 < 1; fi5++ {
        if true {
            var v6 int16 = -378
            v7 := GenS1{F0: -499}
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS2{F0: 71, F1: 2}
    var v9 uint8 = 2
    fmt.Println(v9)
    v10 := GenS1{F0: 0}
    fmt.Println(v9)
    for fi11 := 0; fi11 < 3; fi11++ {
        v9 = v9
        _ = fi11
    }
    v1[""] = ""
    v1["hello"] = "foo bar"
    v9 = (v9 >> ((uint(v9)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
