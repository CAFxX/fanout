package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 string, p1 uint16, p2 uint) uint64 {
    var v4 uint8 = 1
    v5 := GenS1{F0: p0}
    var v6 uint = (^(p2))
    p0 = p0
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{-1, (fi8 + (fi8 << ((65535) & 63))), -9223372036854775808}
            _ = fi8
            _ = v9
        }
        var v10 uint = p2
        v11 := GenS1{F0: ("\n\t" + p0)}
        _ = fi7
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = ((-1 <= 1) && (1 > 1000))
    v3 := make(map[uint32]string)
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v2, fi12)
        _ = fi12
    }
    v2 = (1 != -1)
    v13 := GenS1{F0: "\n\t"}
    var v14 int64 = -9223372036854775808
    v2 = v2
    for fi15 := 0; fi15 < 12; fi15++ {
        v14 += -9223372036854775808
        fmt.Println(v14)
        v14 = v14
        _ = fi15
    }
    v16 := []int{32767, 128}
    v17 := GenS1{F0: "a"}
    var v18 float64 = 0.1
    var v19 uint8 = 42
    var v20 uint8 = (v19 * 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
