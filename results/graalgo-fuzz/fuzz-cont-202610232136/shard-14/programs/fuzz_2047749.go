package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32) uint {
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 -= p0
            var v8 float64 = (0.1 * (-2.25 / -2.25))
            _ = fi7
            _ = v8
        }
        if (true || true) {
            v9 := GenS1{F0: 127}
            _ = v9
        }
        p0 = p0
        _ = fi6
    }
    v10 := []int{-1, 128, int(p0), 4294967295}
    p0 = 1
    _ = p0
    _ = v10
    return 925
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = 0.0
    var v3 string = ((v1 + (v1 + v1)) + (v1 + v1))
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    fmt.Println(v2, v3)
    v3 = v3
    if true {
        v1 = (v3 + ((v3 + "\n\t") + "hello"))
        v11 := GenS1{F0: 1}
        var v12 bool = false
        _ = v11
        _ = v12
    } else {
        var v13 uint32 = 65535
        var v14 uint = 255
        _ = v13
        _ = v14
    }
    v3 = "a"
    var v15 uint8 = 1
    var v16 int16 = -793
    v17 := GenS2{F0: 204}
    v18 := []int{37, 42, 742}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
