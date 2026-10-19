package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int32) bool {
    var v5 uint16 = 560
    p0 = 2147483647
    var v6 uint = 609
    v7 := []int{}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -515
    v3 := make(map[string]int)
    v4 := []string{"a", "", ("" + (("" + "\n\t") + (("a" + (("hello" + "a") + "")) + ("" + "hello"))))}
    if ((42 < -1) && true) {
        v2 -= v2
    }
    fmt.Println(v2)
    v2 = (v2 + (v2 * (v2 * v2)))
    var v8 bool = false
    if v8 {
        if (1 > 0) {
            for fi9 := 0; fi9 < 20; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    fmt.Println(fi10, fi10)
                    v11 := GenS1{F0: 0, F1: v2}
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        }
        var v12 bool = ((1 > 2) && ((42 < 302) && (!true)))
        _ = v12
    } else {
        v1[743] = -184
    }
    if v8 {
        v13 := GenS2{F0: (-2.25 < ((-2.25 - 0.1) - 1.5))}
        v2 -= -585
        _ = v13
    }
    var v14 uint = 65535
    v8 = v8
    var v15 uint64 = 1
    var v16 uint64 = 0
    var v17 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
