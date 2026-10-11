package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 0
    var v3 uint = 0
    v4 := []uint32{0, 1, 4294967295}
    v5 := []int{}
    v6 := []string{""}
    v1 = 0.0
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if ((v2 << ((0) & 63)) >= v2) {
        v3 = v3
    }
    v7 := GenS1{F0: 655}
    v8 := GenS1{F0: 764}
    var v9 uint16 = 32767
    var v10 uint32 = 4294967295
    switch 127 {
    case 32768:
        _ = len(v6)
    default:
        switch 127 {
        case 266:
            for fi11 := 0; fi11 < 4; fi11++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v10 = (v10 | 0)
                _ = fi11
            }
        case -1:
            if len(v4) > 0 {
                v4[0] = 619
            }
        case 4294967295:
            v6 = append(v6, "a")
        default:
            v9 = v9
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = v1
    switch 128 {
    case -1:
        v9 += v9
    case -326:
        v9 -= (^(v9))
    case 32767:
        var v12 uint64 = 605
        _ = v12
    default:
        for fi13 := 0; fi13 < 12; fi13++ {
            if len(v4) > 0 {
                v4[0] = v10
            }
            v14 := GenS1{F0: 2147483647}
            _ = fi13
            _ = v14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
