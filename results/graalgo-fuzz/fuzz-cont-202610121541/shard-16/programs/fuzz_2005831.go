package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint16) uint8 {
    fmt.Println(p0)
    switch 1 {
    case -631:
        var v6 uint8 = (p0 << ((0) & 63))
        _ = v6
    case -2147483648:
        p0 -= 127
    case 157:
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{}
            var v9 int = -9223372036854775808
            for fi10 := 0; fi10 < 6; fi10++ {
                v8 = append(v8, fi7)
                fmt.Println(p1, fi7)
                _ = fi10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v11 float64 = 0.1
    v12 := GenS1{F0: -32768}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 int32 = 1
    v4 := []int{42, 4294967295, -759, 1}
    var v5 int32 = (v3 - 2147483647)
    v13 := GenS1{F0: -32768}
    if true {
        fmt.Println(v3)
    }
    var v14 bool = false
    var v15 bool = v14
    v16 := GenS1{F0: 0}
    for fi17 := 0; fi17 < 17; fi17++ {
        if (!(-128 > 72)) {
            var v18 uint32 = 2147483648
            _ = v18
        }
        _ = fi17
    }
    v3 = v3
    var v19 int8 = 100
    var v20 uint8 = 127
    v15 = (!(false || v14))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
