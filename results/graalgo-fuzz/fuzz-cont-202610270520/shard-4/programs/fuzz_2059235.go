package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = 135
    var v2 uint8 = 14
    var v3 int16 = 1
    var v4 float64 = 0.1
    for fi5 := 0; fi5 < 15; fi5++ {
        v1 = int64(v3)
        v6 := []int{(-(fi5)), fi5}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: "a", F2: 2}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8, -658, fi8, 370}
        _ = fi8
        _ = v9
    }
    v2 = v2
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 float64 = (v4 * (((v4 / ((v4 - 1.5) * v4)) + -2.25) + float64(0.0)))
        if ((fi10 - fi10) == 0) {
            v12 := GenS1{F0: 42, F1: ("foo bar" + "\n\t"), F2: 398}
            v13 := []int{-480}
            v4 = 0.0
            _ = v12
            _ = v13
        }
        if ((^(v2)) < v2) {
            v11 -= v4
        } else {
            v14 := []int{9223372036854775807, fi10, (fi10 >> ((0) & 63)), (fi10 - (fi10 | fi10))}
            v15 := GenS2{F0: 42}
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 uint8 = (v2 >> ((522) & 63))
    v17 := GenS2{F0: 4294967295}
    v18 := GenS1{F0: 217, F1: ("" + ("\n\t" + "foo bar")), F2: 324}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
