package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16, p1 bool, p2 int16) int16 {
    p2 = p2
    var v4 string = "hello"
    for fi5 := 0; fi5 < 6; fi5++ {
        v4 = "\n\t"
        v6 := []int{32768, (fi5 * (fi5 | (fi5 + (fi5 - 2147483647))))}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 255}
    v4 = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 float32 = float32((-2.25 * 3.14159))
    if false {
        var v8 int16 = -1
        v3 = (v3 + 3.14159)
        _ = v8
    } else {
        var v9 float32 = v3
        _ = v9
    }
    var v10 int16 = genFn0(42, true, -997)
    var v11 uint64 = 324
    var v12 uint16 = 42
    var v13 uint = 1
    v10 += -1
    var v14 int = -974
    v12 = (^(v12))
    switch "a" {
    case "":
        v15 := GenS1{F0: 1}
        _ = v15
    case "\n\t":
        for fi16 := 0; fi16 < 13; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                v14 -= v14
                _ = fi17
            }
            var v18 int32 = -546
            _ = fi16
            _ = v18
        }
    case "a":
        var v19 uint = v13
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
