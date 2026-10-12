package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int8 }
func genFn0(p0 int32) uint {
    var v6 int = -2147483648
    var v7 string = ("a" + "\n\t")
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: int16(p0), F1: 1, F2: -1}
        _ = fi8
        _ = v9
    }
    if true {
        v10 := []int{(v6 >> ((65535) & 63)), (v6 ^ v6), (v6 ^ v6)}
        var v11 uint = 2
        _ = len(v10)
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{4294967295, 42, 127}
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint8 = 79
        _ = fi12
        _ = v13
    }
    var v14 float64 = 1.5
    v5[5] = (1000 <= 1)
    var v15 uint = 947
    var v16 uint8 = 127
    fmt.Println(v16)
    if true {
        v14 = 3.14159
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 int64 = -1
            _ = fi17
            _ = v18
        }
    }
    if false {
        switch 128 {
        case 128:
            var v19 float64 = v14
            _ = v19
        }
        var v20 uint16 = 782
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
