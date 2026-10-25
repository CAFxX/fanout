package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -648
    var v3 int8 = 1
    var v4 int64 = 1
    v4 = (v4 >> ((305) & 63))
    var v5 bool = (true && true)
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 1}
        v8 := []int{fi6, 529, (fi6 - (^(fi6)))}
        var v9 uint8 = 42
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: 65535}
        _ = fi10
        _ = v11
    }
    v12 := v1["hello"]
    var v13 uint64 = 2
    v14 := GenS1{F0: 781}
    v15 := []int{128, -865, -1, 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
