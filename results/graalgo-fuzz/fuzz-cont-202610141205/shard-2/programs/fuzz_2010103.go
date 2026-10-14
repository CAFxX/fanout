package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 uint64) int32 {
    p0 = -9223372036854775808
    var v7 uint16 = 65535
    if true {
        switch 127 {
        case 127:
            var v8 uint64 = 0
            _ = v8
        case 0:
            v9 := []int{(-(p0)), (p0 / ((p0) | 1)), 32767}
            _ = v9
        default:
            var v10 float32 = -2.25
            _ = v10
        }
    } else {
        v11 := []int{p0, 32767, p0}
        p1 = (p1 & p1)
        var v12 uint64 = p1
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 788
}
func genFn1(p0 int64, p1 uint16) uint8 {
    p1 = p1
    v13 := GenS1{F0: ("a" + (("" + "\n\t") + "a")), F1: 1}
    _ = p0
    _ = p1
    _ = v13
    return uint8(p1)
}

func main() {
    var v1 int16 = -219
    v2 := []int{32768, 527, 32767}
    v3 := []uint32{437, 2147483648}
    v4 := []uint32{2147483648, 2147483647, 930, 0}
    v5 := make(map[uint32]string)
    var v6 int64 = 547
    var v14 uint = 4294967295
    fmt.Println(v1)
    v6 = (v6 + v6)
    var v15 uint32 = 1
    v16 := v5[(v15 - 2)]
    var v17 int32 = genFn0(576, 2)
    var v18 bool = (true && false)
    var v19 bool = (v17 > (v17 * 42))
    v6 -= (v6 - -978)
    fmt.Println(v15)
    for fi20 := 0; fi20 < 16; fi20++ {
        if true {
            switch (fi20 & int(v14)) {
            case 127:
                fmt.Println(v16)
            case 2147483647:
                v21 := GenS1{F0: (((((v16 + v16) + v16) + "foo bar") + v16) + "hello"), F1: (fi20 - (^(fi20)))}
                _ = v21
            case -2147483648:
                var v22 int = fi20
                _ = v22
            }
        } else {
            v23 := GenS2{F0: 65535, F1: 42}
            _ = v23
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
