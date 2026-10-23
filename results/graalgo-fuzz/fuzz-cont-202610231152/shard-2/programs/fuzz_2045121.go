package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0(p0 string, p1 uint16) uint64 {
    v4 := GenS1{F0: 65535, F1: 642, F2: -2147483648}
    fmt.Println(p0, p1)
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(fi5, p1)
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 815
}

func main() {
    v1 := []uint32{65535, 1, 121, 6}
    var v2 uint = 255
    v3 := make(map[uint32]uint32)
    v2 = v2
    v2 -= v2
    v6 := GenS1{F0: 941, F1: 754, F2: 811}
    if true {
        if (100 < -3) {
            delete(v3, 842)
        } else {
            v2 = 1
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 string = ("" + ("\n\t" + "\n\t"))
                _ = fi7
                _ = v8
            }
        }
        var v9 int64 = 593
        _ = v9
    }
    fmt.Println(v2)
    var v10 uint = 582
    for fi11 := 0; fi11 < 9; fi11++ {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = fi11
    }
    var v12 uint32 = 0
    var v13 bool = (-299 >= 0)
    fmt.Println(v12)
    v14 := v3[v12]
    v14 = v12
    v2 += 981
    v3[v14] = 171
    v15 := GenS1{F0: 905, F1: 0, F2: 42}
    _ = len(v1)
    for fi16 := 0; fi16 < 16; fi16++ {
        v14 -= (^(v12))
        _ = fi16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
}
