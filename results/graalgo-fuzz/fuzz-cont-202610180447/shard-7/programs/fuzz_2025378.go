package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint8 = 117
    v2 := make(map[uint32]bool)
    var v3 bool = (!(((!false) || true) && false))
    var v4 uint32 = 581
    if true {
        var v5 int = 522
        if v3 {
            fmt.Println(v3, v1, v1)
            for fi6 := 0; fi6 < 18; fi6++ {
                v7 := GenS1{F0: v4}
                var v8 bool = v3
                v9 := v2[v4]
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            v2[(^(v4))] = false
        }
        v3 = (467 <= 65535)
        _ = v5
    }
    v4 = 734
    var v10 int32 = 101
    var v11 float64 = -2.25
    v3 = (v3 && true)
    fmt.Println(v3, v10)
    fmt.Println(v11, v3)
    switch -776 {
    case 4294967295:
        var v12 uint64 = 0
        _ = v12
    default:
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 float64 = 1.5
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS2{F0: (v4 - (^(v4)))}
        for fi17 := 0; fi17 < 3; fi17++ {
            switch "\n\t" {
            case "":
                v18 := GenS1{F0: v4}
                _ = v18
            default:
                var v19 int32 = 42
                _ = v19
            }
            var v20 uint16 = 872
            var v21 uint8 = v1
            _ = fi17
            _ = v20
            _ = v21
        }
        if false {
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := GenS1{F0: (v4 + (v4 << ((758) & 63)))}
                _ = fi22
                _ = v23
            }
            _ = len(v2)
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
