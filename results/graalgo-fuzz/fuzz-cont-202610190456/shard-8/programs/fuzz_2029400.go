package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []int{32768}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 float64 = 1.5
    var v5 bool = (703 != 765)
    var v6 float32 = 0.1
    v4 = float64(1.5)
    v7 := GenS1{F0: -1}
    v6 = v6
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = float32(0.0)
        _ = len(v2)
        var v10 uint64 = 2
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: -1}
    v4 -= ((0.0 * 1.5) + 0.0)
    switch -9223372036854775808 {
    case -1:
        fmt.Println(v5)
    case -984:
        var v12 int = 524
        _ = v12
    case 685:
        switch -429 {
        case -2147483648:
            v6 = v6
        case -892:
            switch (("a" + "hello") + "a") {
            case "":
                delete(v2, 1)
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    var v13 bool = ((!v5) || v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
