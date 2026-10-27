package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
func genFn0(p0 bool, p1 uint16, p2 int) uint64 {
    p2 = p2
    var v7 uint16 = 42
    v8 := GenS1{F0: 42, F1: (p1 / ((p1) | 1)), F2: 65535}
    p1 += v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 110
}

func main() {
    var v1 bool = (true || true)
    var v2 uint16 = 1
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 float64 = 1.5
    v6 := make(map[int]uint32)
    var v9 float32 = (1e10 - (3.14159 * 1.5))
    v10 := GenS2{F0: 43, F1: v2, F2: 128}
    v11 := GenS1{F0: 1, F1: (^(v2)), F2: 42}
    if false {
        switch 986 {
        case 32768:
            v12 := GenS1{F0: 2, F1: (v2 << ((2) & 63)), F2: 18446744073709551615}
            _ = v12
        case -850:
            v2 = v2
        case 4294967295:
            fmt.Println(v2, v5)
        default:
            v1 = v1
        }
        delete(v3, "hello")
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int64 = 1
            v15 := []int{843, fi13, 32767}
            var v16 float64 = (v5 - v5)
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    } else {
        var v17 string = "\n\t"
        v18 := []int{}
        _ = v17
        _ = v18
    }
    fmt.Println(v9, v5)
    var v19 uint = 1
    fmt.Println(v19, v9)
    var v20 int16 = -579
    if v1 {
        var v21 uint64 = 42
        switch ("hello" + "foo bar") {
        case "a":
            v22 := v6[9223372036854775807]
            _ = v22
        default:
            v23 := GenS1{F0: 255, F1: (v2 | (v2 / (((v2 ^ (v2 % ((210) | 1)))) | 1))), F2: 364}
            _ = v23
        }
        _ = v21
    }
    v19 = (v19 | 132)
    v24 := GenS1{F0: 121, F1: v2, F2: (v19 >> (((v19 % (((^(v19))) | 1))) & 63))}
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            if true {
                v9 = v9
                var v27 uint8 = 128
                _ = v27
            } else {
                v28 := []int{-2147483648, (fi26 % ((fi25) | 1)), 32768}
                _ = v28
            }
            _ = fi26
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            switch "hello" {
            case "":
                v20 -= 42
            case "\n\t":
                if v1 {
                    fmt.Println(v20, v9, fi25)
                    var v30 int64 = -352
                    _ = v30
                }
            }
            var v31 uint8 = 0
            _ = fi29
            _ = v31
        }
        v19 += v19
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
