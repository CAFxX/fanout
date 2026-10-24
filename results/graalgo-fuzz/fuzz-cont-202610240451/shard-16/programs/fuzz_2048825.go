package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint = 681
    v2 := make(map[uint32]bool)
    var v3 int64 = 1
    v4 := make(map[uint32]int)
    v5 := []uint32{2, 42}
    var v6 string = (("\n\t" + "") + "hello")
    var v7 bool = (255 < 297)
    switch 4294967295 {
    case 42:
        v1 -= 1
    case 2147483647:
        v8 := GenS1{F0: ("a" + v6)}
        _ = v8
    case 32767:
        var v9 int16 = 32767
        _ = v9
    }
    var v10 uint8 = 1
    if v7 {
        v7 = v7
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 uint64 = 2
            delete(v4, 2)
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: ((v6 + v6) + v6)}
        _ = v13
    } else {
        var v14 int8 = 42
        var v15 int32 = 1
        switch -879 {
        case 127:
            _ = len(v2)
        case 726:
            if (v10 >= 42) {
                v15 += (-(v15))
            } else {
                for fi16 := 0; fi16 < 16; fi16++ {
                    fmt.Println(v7)
                    _ = fi16
                }
            }
        }
        _ = v14
        _ = v15
    }
    var v17 bool = v7
    v18 := GenS1{F0: (v6 + (v6 + v6))}
    fmt.Println(v6, v6, v1)
    v5 = append(v5, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
