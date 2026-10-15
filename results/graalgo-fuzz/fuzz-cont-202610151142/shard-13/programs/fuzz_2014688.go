package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool }
func genFn0(p0 int, p1 uint8) uint8 {
    p1 = (p1 / (((p1 << ((42) & 63))) | 1))
    v6 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1(p0 bool) uint16 {
    fmt.Println(p0)
    p0 = (!p0)
    _ = p0
    return 32767
}

func main() {
    v1 := []int{-406, 42}
    var v2 int64 = -1
    var v3 uint = 42
    v4 := make(map[int]uint32)
    var v5 int16 = 773
    var v7 uint = 186
    switch -975 {
    case 127:
        v1 = append(v1, 9223372036854775807)
    case 609:
        if false {
            var v8 uint16 = 32767
            var v9 int8 = int8(v8)
            _ = v8
            _ = v9
        }
    default:
        if (0 > 280) {
            for fi10 := 0; fi10 < 2; fi10++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi10
            }
            var v11 string = ""
            _ = v11
        } else {
            v2 = v2
            var v12 uint32 = 1
            if (("a" + "foo bar") > "a") {
                var v13 uint64 = uint64(v12)
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS1{F0: v2, F1: 1000, F2: (v13 << (((v7 - v7)) & 63))}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            _ = v12
        }
    }
    v3 = 65535
    v16 := GenS2{F0: false}
    v17 := GenS1{F0: v2, F1: v5, F2: 1}
    for fi18 := 0; fi18 < 16; fi18++ {
        v5 += v5
        v4[(fi18 * 9223372036854775807)] = 656
        _ = fi18
    }
    v7 += v7
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 float32 = 3.14159
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
