package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int8, p1 uint16) int8 {
    var v7 float32 = -2.25
    if ("foo bar" == "") {
        if false {
            if (!false) {
                var v8 int = int(p1)
                v9 := []int{v8, 32768, v8}
                _ = v8
                _ = v9
            }
        } else {
            var v10 uint32 = 42
            p0 -= p0
            _ = v10
        }
    }
    var v11 bool = false
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return p0
}

func main() {
    v1 := []uint32{639, 0}
    var v2 int8 = -1
    var v3 uint64 = 90
    var v4 float32 = -2.25
    var v5 int = 9223372036854775807
    var v6 int16 = -1
    var v12 float64 = 1.5
    v1 = append(v1, 0)
    v3 -= 1
    var v13 uint64 = v3
    v4 += (v4 - (1.5 * (v4 + v4)))
    fmt.Println(v3, v3)
    var v14 uint = 1
    v15 := GenS1{F0: v5}
    v4 -= v4
    for fi16 := 0; fi16 < 20; fi16++ {
        v12 = float64(v4)
        _ = fi16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int64 = -258
        _ = fi17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
