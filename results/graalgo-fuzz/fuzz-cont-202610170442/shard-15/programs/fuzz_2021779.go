package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []string{}
    var v2 uint = 4294967295
    v3 := make(map[string]int)
    var v4 uint8 = 2
    var v5 uint16 = 1
    var v6 float32 = 1e10
    _ = len(v3)
    v2 += (^(v2))
    v5 -= v5
    if false {
        v7 := GenS1{F0: 261}
        v6 -= v6
        _ = v7
    } else {
        v2 = v2
        switch 127 {
        case -726:
            _ = len(v3)
        case -391:
            var v8 float32 = v6
            _ = v8
        case 1:
            var v9 uint = (^(v2))
            _ = v9
        default:
            for fi10 := 0; fi10 < 18; fi10++ {
                v4 -= (v4 >> ((v2) & 63))
                _ = fi10
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: 1}
        _ = fi11
        _ = v12
    }
    v1 = append(v1, ("\n\t" + "a"))
    v2 -= v2
    switch -843 {
    case 42:
        v13 := v3["\n\t"]
        _ = v13
    case -490:
        v5 += v5
    case -662:
        var v14 uint8 = 211
        _ = v14
    default:
        var v15 bool = ((2 >= 349) || (0 >= 784))
        _ = v15
    }
    var v16 uint64 = 1
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
