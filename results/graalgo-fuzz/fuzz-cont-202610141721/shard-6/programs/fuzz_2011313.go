package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0(p0 int32, p1 uint16, p2 int32) int8 {
    var v5 uint16 = p1
    var v6 float64 = -2.25
    var v7 bool = true
    v8 := []int{}
    p1 += 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    var v3 int16 = 42
    v4 := make(map[uint32]bool)
    var v9 string = "\n\t"
    v4[903] = ((2 < 179) || true)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 float64 = float64(((1.5 - -2.25) * float32(1e10)))
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint = 909
            v1 = v1
            v9 = v9
            switch int(v13) {
            case 32768:
                var v14 int = fi12
                _ = v14
            case -982:
                v15 := GenS1{F0: 222}
                _ = v15
            }
            v16 := GenS2{F0: 0, F1: (v3 * 0), F2: 2}
            _ = fi12
            _ = v13
            _ = v16
        }
        v9 = v9
        _ = fi10
        _ = v11
    }
    v2[int(v1)] = (!(!false))
    v1 -= v1
    v3 = 42
    v9 = "foo bar"
    for fi17 := 0; fi17 < 12; fi17++ {
        v2[fi17] = (!true)
        fmt.Println(v1, v1, v1)
        _ = fi17
    }
    v18 := GenS2{F0: 9223372036854775807, F1: (-(v3)), F2: 2}
    switch (v9 + v9) {
    case "foo bar":
        v3 = v3
    }
    v3 += (v3 | (v3 & (v3 % ((v3) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
