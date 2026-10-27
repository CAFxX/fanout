package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
func genFn0(p0 uint16, p1 string) uint {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint = 65535
        fmt.Println(p1, v5, p0)
        _ = fi4
        _ = v5
    }
    p1 = p1
    fmt.Println(p1, p0)
    var v6 uint32 = 0
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 int64 = 990
    v2 := make(map[string]uint32)
    var v3 int64 = v1
    var v7 int32 = 468
    v3 += v3
    v1 += v3
    var v8 int = 32767
    delete(v2, ("a" + ""))
    v8 = 127
    v9 := GenS1{F0: 1, F1: false, F2: 0}
    var v10 int8 = 4
    v11 := GenS1{F0: 2147483648, F1: false, F2: v10}
    var v12 bool = (false || ((true && (false && ((true || ((1.5 * (1.5 - -2.25)) < 0.0)) && ((3.14159 >= -2.25) || true)))) && (!true)))
    v10 += -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
