package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    var v4 float64 = 0.1
    v1 += -1
    for fi5 := 0; fi5 < 9; fi5++ {
        v4 = float64(3.14159)
        _ = fi5
    }
    v6 := v2[1]
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            fmt.Println(v6)
        }
        v8 := GenS1{F0: 1}
        var v9 int = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v2)
    switch v6 {
    case 20:
        var v10 bool = (v1 < (v1 & int32(v6)))
        _ = v10
    default:
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v1)
            v2[42] = 9223372036854775807
            _ = fi11
        }
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            v1 = v1
            _ = fi13
        }
        _ = fi12
    }
    if true {
        v4 = float64(float32((1.5 + v4)))
    } else {
        v14 := []int{v6, -29, -1, (v6 | 277)}
        for fi15 := 0; fi15 < 3; fi15++ {
            if true {
                v16 := GenS1{F0: 21}
                v3 = (v3 << ((65535) & 63))
                _ = v16
            } else {
                var v17 bool = true
                var v18 int = fi15
                _ = v17
                _ = v18
            }
            var v19 uint = 2
            v1 = (v1 % ((v1) | 1))
            _ = fi15
            _ = v19
        }
        v20 := GenS2{F0: 0}
        _ = v14
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
