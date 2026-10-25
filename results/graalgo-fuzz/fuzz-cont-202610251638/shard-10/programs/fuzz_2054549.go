package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]bool)
    var v3 uint64 = 2
    v4 := []int{32767}
    var v5 int8 = (v1 - 127)
    var v6 uint = 18446744073709551615
    v6 = (v6 / ((v6) | 1))
    fmt.Println(v1, v6)
    var v7 int16 = -132
    v5 = (-(v1))
    var v8 uint32 = 210
    switch -584 {
    case -2147483648:
        v9 := GenS1{F0: (v8 | 0)}
        _ = v9
    default:
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(v7)
            v11 := GenS1{F0: v8}
            fmt.Println(v1)
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if true {
            if (!true) {
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 string = ("\n\t" + "\n\t")
                    switch (v14 + (v14 + "hello")) {
                    case "hello":
                        var v15 uint32 = (v8 | 723)
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                v16 := GenS1{F0: 2147483648}
                _ = v16
            }
        } else {
            var v17 uint8 = 77
            var v18 uint8 = v17
            fmt.Println(v7, v17)
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            _ = len(v2)
            if len(v4) > 0 {
                v4[0] = (fi19 | int(v6))
            }
            _ = fi19
        }
        fmt.Println(v8, v1)
        _ = fi12
    }
    switch 0 {
    case 32767:
        switch 32768 {
        case -555:
            v4 = append(v4, 42)
        default:
            var v20 uint8 = 2
            _ = v20
        }
    case 0:
        if len(v4) > 0 {
            v4[0] = 412
        }
    case -319:
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
