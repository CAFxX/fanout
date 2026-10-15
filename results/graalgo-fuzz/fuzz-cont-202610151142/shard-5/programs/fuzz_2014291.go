package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = -1
    var v2 float32 = -2.25
    var v3 float32 = v2
    var v4 uint = 735
    var v5 uint8 = 41
    var v6 float32 = v2
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint32 = 4294967295
        v9 := GenS1{F0: 4}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = 1e10
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint = v4
        _ = fi11
        _ = v12
    }
    var v13 int = 654
    if true {
        var v14 uint = (v4 & (v4 * (v4 ^ (v4 - v4))))
        var v15 string = "\n\t"
        v16 := []int{(v13 - 32768)}
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := []int{v13}
    v3 += v6
    switch (v13 / ((v13) | 1)) {
    case 2147483647:
        v18 := GenS1{F0: int8(v5)}
        _ = v18
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := GenS1{F0: 100}
            _ = fi19
            _ = v20
        }
    }
    var v21 int = v13
    var v22 int16 = -544
    var v23 int16 = (v22 & (v22 & -148))
    v4 += 0
    v24 := GenS1{F0: 1}
    _ = len(v17)
    var v25 uint64 = 4294967295
    if false {
        if ((v22 + v22) == 0) {
            var v26 int16 = v22
            _ = v26
        } else {
            for fi27 := 0; fi27 < 20; fi27++ {
                switch (-(v13)) {
                case 868:
                    v25 = v25
                default:
                    var v28 int = -1
                    _ = v28
                }
                _ = fi27
            }
        }
        fmt.Println(v1)
    } else {
        if len(v17) > 0 {
            v17[0] = (v21 << ((v4) & 63))
        }
        fmt.Println(v6, v10, v1)
        var v29 uint32 = uint32(v25)
        _ = v29
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 float64 = v10
        _ = fi30
        _ = v31
    }
    if true {
        var v32 float64 = float64(v6)
        _ = v32
    } else {
        v23 = v23
        v17 = append(v17, v21)
        var v33 int64 = -9223372036854775808
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
