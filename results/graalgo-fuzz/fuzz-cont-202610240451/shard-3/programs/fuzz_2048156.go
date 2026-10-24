package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 float32 = (float32((1e10 * (float64(((0.1 + (float32(1.5) + 0.0)) * (1.5 / -2.25))) * 0.0))) - 0.1)
    var v2 int32 = -1
    var v3 uint32 = 2
    v4 := make(map[int]bool)
    v3 -= 65535
    v2 = v2
    switch 32768 {
    case -222:
        v5 := GenS1{F0: 2}
        _ = v5
    case -9223372036854775808:
        switch 416 {
        case 4294967295:
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 int = fi6
                for fi8 := 0; fi8 < 3; fi8++ {
                    fmt.Println(v7, v3)
                    _ = fi8
                }
                var v9 bool = true
                _ = fi6
                _ = v7
                _ = v9
            }
        }
    default:
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int32 = 2147483647
            v2 = (v11 ^ (v11 ^ v11))
            _ = fi10
            _ = v11
        }
    }
    var v12 float32 = float32(-2.25)
    if false {
        fmt.Println(v2)
        v1 = ((((v1 * v12) + v1) / (((v1 / v12) + 1.5) + v12)) * v1)
    } else {
        v13 := []int{0, -249, -9223372036854775808, -1}
        _ = v13
    }
    v14 := []int{-207, 555}
    if (285 >= 2) {
        var v15 string = ("\n\t" + "")
        v3 = 2147483647
        _ = v15
    }
    var v16 int = -201
    v16 -= (v16 + 0)
    if ((1 > -9223372036854775808) && (1e10 != v12)) {
        v12 += 0.0
        var v17 uint8 = 127
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 int16 = 1
            fmt.Println(v19, fi18)
            var v20 int32 = v2
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    } else {
        _ = len(v14)
        switch "\n\t" {
        case "hello":
            for fi21 := 0; fi21 < 15; fi21++ {
                if (true && true) {
                    v1 = (v12 - 1.5)
                    if (1e10 >= float32(1.5)) {
                        var v22 int32 = (v2 << ((1) & 63))
                        v12 = ((v1 / ((v12 + 0.1) / v12)) * v12)
                        var v23 int = (-(v16))
                        _ = v22
                        _ = v23
                    }
                    v24 := v4[(fi21 & fi21)]
                    _ = v24
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 bool = (true && (true || (0.1 < 3.14159)))
                    v2 += v2
                    v27 := GenS1{F0: v3}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 17; fi28++ {
                    fmt.Println(fi21)
                    _ = fi28
                }
                _ = fi21
            }
        }
    }
    v29 := GenS2{F0: 42}
    v1 = (((0.1 * (float32(1.5) - (1.5 / v1))) - (float32(((-2.25 / ((0.0 - 3.14159) + 0.1)) - 3.14159)) * v1)) / v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v14
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
