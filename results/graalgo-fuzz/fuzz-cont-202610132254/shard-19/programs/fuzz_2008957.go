package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int16) uint16 {
    var v7 int8 = -74
    v7 = v7
    var v8 bool = (49 > 0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 bool) int16 {
    fmt.Println(p0)
    if ((float64((0.1 - (1e10 * 0.1))) + (((1.5 / 1.0) + (-2.25 / -2.25)) + 1e10)) != 0.0) {
        p0 = true
        p0 = true
    } else {
        v9 := GenS1{F0: 237, F1: 1, F2: -40}
        var v10 int16 = 0
        _ = v10
        _ = v9
    }
    v11 := []int{}
    var v12 uint = 0
    v13 := GenS1{F0: 9223372036854775807, F1: -663, F2: -128}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -9223372036854775808
    var v3 float64 = 0.0
    var v4 uint16 = 0
    v5 := make(map[uint32]string)
    var v6 int64 = 0
    switch "\n\t" {
    case "":
        v3 = 0.1
    case "hello":
        v3 = v3
    }
    switch (("\n\t" + "hello") + "a") {
    case "\n\t":
        for fi14 := 0; fi14 < 20; fi14++ {
            if ((v4 >> ((451) & 63)) <= genFn0(532, 288)) {
                v6 -= (v2 << ((762) & 63))
                v3 = (-2.25 / v3)
                var v15 uint32 = 507
                _ = v15
            } else {
                fmt.Println(v6, v2)
                v16 := []int{}
                _ = v16
            }
            _ = fi14
        }
    default:
        fmt.Println(v6, v6, v4)
    }
    var v17 float64 = float64(-2.25)
    v4 = (v4 << ((800) & 63))
    switch "hello" {
    case "foo bar":
        v18 := GenS2{F0: 42}
        _ = v18
    case "\n\t":
        v19 := GenS1{F0: (v6 / (((^(v2))) | 1)), F1: 2147483647, F2: -45}
        _ = v19
    case "a":
        if (0.1 < 0.1) {
            v3 += (((((v3 * v3) - v3) - v17) / v3) - float64(float32(float64((((3.14159 - -2.25) + 1e10) + 1.5)))))
            v20 := []int{2147483647}
            v4 = (v4 << ((103) & 63))
            _ = v20
        } else {
            v1[-31] = 32768
            var v21 float32 = -2.25
            _ = v21
        }
    default:
        switch int(v2) {
        case -970:
            var v22 int = 422
            _ = v22
        default:
            for fi23 := 0; fi23 < 5; fi23++ {
                delete(v5, 0)
                for fi24 := 0; fi24 < 20; fi24++ {
                    v4 = 42
                    _ = fi24
                }
                _ = fi23
            }
        }
    }
    var v25 uint = 48
    delete(v1, 32768)
    v1[-157] = -187
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
