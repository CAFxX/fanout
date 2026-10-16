package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint32 }
func genFn0(p0 bool, p1 int32, p2 uint) int32 {
    v7 := GenS1{F0: 1}
    var v8 float64 = (float64(-2.25) - 0.0)
    v9 := []int{32767, 42, 127, 127}
    _ = len(v9)
    fmt.Println(p0, p0, v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 1.5
    v3 := []string{"a", "hello"}
    var v4 int = 756
    v5 := make(map[int]bool)
    var v6 int = (v4 >> ((2) & 63))
    if (!false) {
        v4 += v4
    }
    v2 = (0.0 / 1.5)
    v10 := GenS1{F0: -891}
    delete(v1, v6)
    var v11 uint16 = 1
    v4 -= 4294967295
    fmt.Println(v2, v6)
    for fi12 := 0; fi12 < 1; fi12++ {
        v6 -= 4294967295
        var v13 uint16 = (v11 + v11)
        _ = fi12
        _ = v13
    }
    if (-428 >= 1) {
        for fi14 := 0; fi14 < 15; fi14++ {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v4 = v6
            _ = fi14
        }
        fmt.Println(v4)
        v4 = 127
    }
    switch ("\n\t" + ("a" + "\n\t")) {
    case "\n\t":
        v15 := GenS1{F0: 1000}
        _ = v15
    case "":
        if (-1 != 42) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v6, v4, v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v16 := GenS2{F0: -171, F1: 171, F2: 42}
    var v17 string = ("a" + ("" + ("" + "\n\t")))
    v1[(v6 ^ v6)] = 0
    fmt.Println(v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
