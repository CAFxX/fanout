package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 int32, p1 int) int32 {
    var v6 float64 = 3.14159
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint = 18446744073709551615
        var v9 float64 = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 0}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return p0
}
func genFn1(p0 uint16, p1 bool, p2 bool) uint {
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            switch -525 {
            case 0:
                var v13 int16 = 0
                _ = v13
            case 32767:
                v14 := []int{(fi12 - 32768), (fi12 / (((fi11 / ((698) | 1))) | 1))}
                _ = v14
            case 9223372036854775807:
                var v15 uint8 = 42
                _ = v15
            }
            var v16 uint64 = 4294967295
            v17 := []int{-1}
            var v18 string = ("foo bar" + "\n\t")
            p2 = p1
            _ = fi12
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi11
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        p0 -= p0
        _ = fi19
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint32 = 42
    var v2 bool = true
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]int)
    var v5 uint8 = 80
    v20 := GenS2{F0: 0, F1: 1}
    v1 = (^(v1))
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint64 = 1
        _ = fi21
        _ = v22
    }
    if v2 {
        v2 = (uint8(v1) != 255)
    }
    fmt.Println(v2)
    var v23 uint = 2
    if v2 {
        if (1 != 0) {
            switch -161 {
            case 32768:
                switch 127 {
                case -83:
                    var v24 float32 = 1e10
                    _ = v24
                default:
                    fmt.Println(v3, v23, v3)
                }
            }
        }
        switch 32768 {
        case -2147483648:
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int64 = v3
                _ = fi25
                _ = v26
            }
        case 128:
            var v27 uint16 = 255
            _ = v27
        }
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
