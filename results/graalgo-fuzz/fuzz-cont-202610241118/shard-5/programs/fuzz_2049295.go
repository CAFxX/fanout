package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -1
    var v3 uint16 = 53
    var v4 int = 127
    v5 := []uint32{487, 4294967295}
    v6 := []uint32{}
    v3 = v3
    var v7 int8 = 79
    var v8 uint = 255
    v6 = append(v6, 2147483647)
    fmt.Println(v4, v7, v8)
    var v9 uint64 = 42
    if (255 >= 2) {
        fmt.Println(v7, v4)
    }
    v6 = append(v6, 284)
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint32 = 1
            v7 = v7
            switch "foo bar" {
            case "foo bar":
                v13 := GenS1{F0: 1, F1: (("a" + ("a" + "foo bar")) + ("a" + ("a" + "\n\t"))), F2: 945}
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        var v14 int64 = -1
        _ = fi10
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint32 = 65535
        var v17 int32 = v2
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := v1[42]
        v7 = (-(v7))
        _ = fi18
        _ = v19
    }
    var v20 uint16 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
