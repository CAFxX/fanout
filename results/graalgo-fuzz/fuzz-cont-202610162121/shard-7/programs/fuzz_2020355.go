package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
func genFn0(p0 int8) int8 {
    var v4 int8 = (-(p0))
    v5 := []int{1, -170}
    if false {
        v4 = (v4 << ((4294967295) & 63))
        for fi6 := 0; fi6 < 2; fi6++ {
            switch "" {
            case "foo bar":
                if (true && false) {
                    var v7 int32 = -2147483648
                    var v8 int8 = v4
                    _ = v7
                    _ = v8
                }
            case "a":
                var v9 float64 = -2.25
                _ = v9
            default:
                var v10 uint32 = 65535
                _ = v10
            }
            v4 -= 100
            var v11 int16 = -500
            _ = fi6
            _ = v11
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int8, p1 uint) string {
    p0 -= (p0 * p0)
    for fi12 := 0; fi12 < 16; fi12++ {
        p1 = (p1 >> ((2) & 63))
        fmt.Println(p1)
        p1 -= p1
        _ = fi12
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return ("hello" + (("\n\t" + "foo bar") + ("foo bar" + "a")))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]bool)
    var v13 float32 = (1e10 / -2.25)
    v13 = ((float32(-2.25) - 0.1) * float32(0.1))
    v13 += ((((v13 / v13) * ((v13 / -2.25) + (v13 / v13))) - float32(1.5)) + v13)
    v2 -= v2
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int = (fi14 ^ (fi14 + fi14))
        _ = fi14
        _ = v15
    }
    fmt.Println(v2, v2)
    v16 := GenS1{F0: 1, F1: 0}
    var v17 uint = 4294967295
    v17 -= 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
