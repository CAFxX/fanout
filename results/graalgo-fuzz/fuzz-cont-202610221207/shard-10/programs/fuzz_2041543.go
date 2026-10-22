package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 uint64) uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 789}
        p0 -= (^(p0))
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    p0 = p0
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int16 = 42
    var v2 float32 = 1e10
    var v3 uint8 = 1
    var v4 int16 = v1
    v3 += (v3 + (v3 ^ v3))
    var v9 uint = 255
    switch -391 {
    case -499:
        v10 := []int{-1, 32768, 1}
        _ = v10
    case -9223372036854775808:
        v3 -= 174
    case 127:
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS2{F0: int(v4), F1: (fi12 & (fi12 | fi12)), F2: 2147483648}
        var v14 int32 = -1
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float64 = (-2.25 / -2.25)
    v16 := []int{}
    v17 := GenS1{F0: 907}
    var v18 uint16 = 255
    v3 += (v3 | (v3 | (v3 + v3)))
    for fi19 := 0; fi19 < 3; fi19++ {
        v16 = append(v16, (-(fi19)))
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = fi19
    }
    var v20 int64 = 1
    v1 -= (v4 << ((v9) & 63))
    if false {
        var v21 int64 = (v20 * v20)
        _ = v21
    }
    v20 = v20
    v22 := GenS2{F0: -1, F1: 9223372036854775807, F2: 337}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
