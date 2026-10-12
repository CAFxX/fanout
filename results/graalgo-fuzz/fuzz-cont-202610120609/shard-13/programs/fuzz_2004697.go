package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 int }
func genFn0() uint {
    v6 := []int{818}
    var v7 int = 42
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 int64 = 1
    v2 := []string{("\n\t" + "\n\t"), "hello"}
    var v3 uint64 = 18446744073709551615
    v4 := []int{}
    v5 := make(map[int]string)
    v1 += (-(v1))
    v8 := GenS1{F0: (170 != 0), F1: 435}
    v9 := GenS2{F0: -141, F1: -9223372036854775808}
    for fi10 := 0; fi10 < 18; fi10++ {
        v3 = v3
        switch 42 {
        case 32768:
            v11 := GenS2{F0: 2147483647, F1: (fi10 / ((9223372036854775807) | 1))}
            _ = v11
        case -119:
            var v12 uint16 = 255
            _ = v12
        }
        v3 += v3
        _ = fi10
    }
    v1 -= v1
    v3 = (v3 >> ((4294967295) & 63))
    var v13 bool = false
    v14 := v5[-9223372036854775808]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
