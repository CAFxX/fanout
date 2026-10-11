package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint32, p1 uint64, p2 uint16) uint {
    var v6 uint8 = 255
    p0 = 2
    p1 += p1
    fmt.Println(p0)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int64 = -1
        if (3.14159 < (0.1 - 1e10)) {
            if ("" > "hello") {
                if true {
                    if true {
                        for fi9 := 0; fi9 < 9; fi9++ {
                            var v10 string = "a"
                            fmt.Println(fi9, v10, fi7)
                            _ = fi9
                            _ = v10
                        }
                        var v11 bool = false
                        v6 -= (v6 % ((v6) | 1))
                        _ = v11
                    }
                }
            }
        }
        var v12 float32 = -2.25
        _ = fi7
        _ = v12
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    var v3 int = 2147483647
    v4 := []int{}
    var v5 int16 = -1
    v13 := GenS1{F0: 0, F1: 1, F2: 255}
    var v14 int16 = (v5 + v5)
    delete(v2, "hello")
    var v15 int = v3
    v16 := GenS2{F0: 42, F1: 32767}
    switch ("foo bar" + "a") {
    case "\n\t":
        v14 = (v5 >> ((4294967295) & 63))
    default:
        var v17 bool = (!true)
        _ = v17
    }
    v4 = append(v4, (v15 * 4294967295))
    var v18 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
