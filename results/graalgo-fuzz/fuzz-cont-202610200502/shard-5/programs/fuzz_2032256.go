package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint32, p2 uint8) int16 {
    if (748 != 254) {
        var v7 int64 = -1
        _ = v7
    }
    fmt.Println(p0, p0)
    var v8 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return int16(p0)
}
func genFn1(p0 int32, p1 int32) int64 {
    v9 := []int{248}
    _ = len(v9)
    p1 -= (p0 / ((p1) | 1))
    v9 = append(v9, 1)
    var v10 int = 2147483647
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 518
}

func main() {
    var v1 int16 = -812
    v2 := make(map[int]bool)
    v3 := []string{("" + "foo bar")}
    var v4 uint32 = 42
    var v5 float64 = (1.5 * (-2.25 + (float64(float32(1.5)) / 3.14159)))
    v6 := []string{"", ("hello" + "a")}
    v5 += (v5 + 0.1)
    v11 := GenS1{F0: 0}
    switch "hello" {
    case "":
        v1 += 32767
    case "hello":
        var v12 bool = (-1 <= -1)
        _ = v12
    case "\n\t":
        if (v4 >= v4) {
            var v13 uint32 = v4
            var v14 int = -9223372036854775808
            if false {
                var v15 int = v14
                _ = v15
            } else {
                v14 = v14
            }
            _ = v13
            _ = v14
        }
    default:
        v16 := GenS2{F0: (v4 >> ((299) & 63))}
        _ = v16
    }
    var v17 uint16 = 255
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 bool = false
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 100}
    if true {
        _ = len(v3)
        if false {
            v1 = (-(v1))
        } else {
            switch (("\n\t" + "") + "foo bar") {
            case "":
                _ = len(v3)
            }
        }
        v5 -= 1e10
    } else {
        var v21 int16 = genFn0(1, v4, 128)
        _ = v21
    }
    var v22 float64 = float64(1.5)
    var v23 float64 = 0.0
    v17 = 42
    var v24 int32 = -591
    v1 = (-(v1))
    var v25 uint32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
