package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0(p0 int8, p1 int32, p2 string) int16 {
    p1 += (p1 ^ (p1 ^ 995))
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(fi7, p2, fi7)
        p1 = p1
        switch (fi7 / ((fi7) | 1)) {
        case 1:
            var v8 float64 = 1.5
            _ = v8
        case 4294967295:
            p0 = (p0 ^ 100)
        case 128:
            if (1e10 == 1e10) {
                var v9 uint = 651
                v10 := GenS1{F0: 1, F1: 65535, F2: 100}
                _ = v10
                _ = v9
            }
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    var v1 float64 = 0.0
    var v2 int16 = 305
    var v3 string = "a"
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]uint32)
    var v11 int64 = 0
    var v12 float64 = float64((0.1 - (3.14159 * ((1.5 / (-2.25 * 1e10)) + 0.0))))
    v13 := v5[2]
    fmt.Println(v3, v13)
    v14 := GenS1{F0: (v11 & v11), F1: 1, F2: -128}
    var v15 int = -998
    fmt.Println(v15)
    var v16 int8 = 0
    for fi17 := 0; fi17 < 2; fi17++ {
        switch -1 {
        case -2147483648:
            var v18 float32 = -2.25
            _ = v18
        }
        _ = fi17
    }
    _ = len(v6)
    v12 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
