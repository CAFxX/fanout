package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 uint32, p1 uint16, p2 uint8) uint64 {
    p1 = (p1 + 1)
    v5 := []int{-749, -251, 404}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 int = 127
    v2 := make(map[uint32]string)
    var v3 float64 = float64((0.1 * (0.1 / (1e10 / 1.5))))
    var v4 int64 = -9223372036854775808
    fmt.Println(v3, v4, v4)
    var v6 float64 = v3
    fmt.Println(v4, v1)
    _ = len(v2)
    var v7 float32 = float32(((float64((-2.25 + 0.1)) / -2.25) * 0.0))
    for fi8 := 0; fi8 < 1; fi8++ {
        if true {
            v6 = 3.14159
            if (!("a" > ("a" + ("\n\t" + ("hello" + ""))))) {
                var v9 uint64 = 806
                _ = v9
            }
            v4 = 125
        } else {
            v10 := GenS1{F0: 0, F1: 351}
            if (1 != 42) {
                var v11 uint64 = 42
                v7 = v7
                _ = v11
            } else {
                var v12 float64 = 1.5
                _ = v12
            }
            _ = v10
        }
        v2[4294967295] = "hello"
        _ = fi8
    }
    delete(v2, 2147483647)
    v13 := GenS2{F0: 377, F1: (v4 + v4)}
    v6 += v6
    v3 += (-2.25 / v3)
    _ = len(v2)
    var v14 float32 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
