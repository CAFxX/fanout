package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16) uint {
    p0 = (p0 + p0)
    p0 = (p0 * 0)
    var v7 string = (("foo bar" + "a") + "\n\t")
    if (42 >= 42) {
        v7 = ((v7 + ("\n\t" + "foo bar")) + (("a" + v7) + v7))
        for fi8 := 0; fi8 < 18; fi8++ {
            p0 += (p0 + 767)
            _ = fi8
        }
    }
    _ = p0
    _ = v7
    return 129
}

func main() {
    v1 := []uint32{2, 4294967295, 2, 1}
    v2 := []int{-203, 1, -9223372036854775808, 2147483647}
    var v3 float64 = 0.1
    v4 := make(map[string]bool)
    var v5 int = 4294967295
    var v6 uint8 = 0
    v3 = (float64(float32(v3)) - (v3 / 0.1))
    v3 -= (1.5 - 0.1)
    fmt.Println(v6, v5, v6)
    v6 -= 42
    v9 := GenS1{F0: ("foo bar" + "foo bar")}
    v3 += v3
    var v10 float64 = (3.14159 * (1e10 * v3))
    fmt.Println(v3, v5, v10)
    v11 := GenS1{F0: ("hello" + "a")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
