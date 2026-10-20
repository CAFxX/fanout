package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = 695
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v5 int16 = -369
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint16 = 582
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(v7)
            _ = fi8
        }
        var v9 uint64 = 81
        _ = fi6
        _ = v7
        _ = v9
    }
    var v10 uint = 42
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int8 = 42
            var v14 int8 = v13
            var v15 uint64 = 42
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int16 = (v5 % (((v5 & (v5 % ((-32768) | 1)))) | 1))
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v2 = v2
    var v19 int32 = (v2 >> ((v10) & 63))
    var v20 int8 = 1
    if false {
        delete(v4, 528)
    }
    if (!false) {
        var v21 int32 = (v19 * (^(v2)))
        if false {
            var v22 uint32 = (^(v1))
            _ = v22
        } else {
            v23 := GenS1{F0: 710}
            v10 = v10
            _ = v23
        }
        if false {
            delete(v3, 864)
        } else {
            var v24 uint = v10
            _ = v24
        }
        _ = v21
    }
    v25 := []int{128, 284, 272, -84}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
