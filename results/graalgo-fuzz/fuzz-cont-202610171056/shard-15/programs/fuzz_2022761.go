package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]uint32)
    v3 := []uint32{4294967295}
    v4 := []uint32{273, 4294967295, 4294967295, 2}
    var v5 int = 128
    v6 := make(map[uint32]string)
    var v7 uint8 = 128
    var v8 int32 = (v1 >> ((18446744073709551615) & 63))
    var v9 uint16 = 1
    if (!false) {
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                v7 = v7
                _ = fi11
            }
            _ = fi10
        }
    }
    v12 := GenS1{F0: (!false)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            v7 = v7
            var v15 uint16 = v9
            fmt.Println(v9, fi14, v5)
            _ = fi14
            _ = v15
        }
        if (-119 != -146) {
            fmt.Println(v8, v7)
            v5 = v5
        }
        var v16 float32 = (0.1 + 1.5)
        _ = fi13
        _ = v16
    }
    switch "foo bar" {
    case "":
        fmt.Println(v8)
    case "\n\t":
        fmt.Println(v8, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
