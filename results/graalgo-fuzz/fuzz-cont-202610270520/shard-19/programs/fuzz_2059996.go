package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 int) int32 {
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 += p0
        _ = fi6
    }
    var v7 int = (p0 / (((p0 / ((p0) | 1))) | 1))
    if (true || false) {
        p0 -= (v7 << ((2) & 63))
        v8 := []int{9223372036854775807, v7, 0}
        _ = v8
    }
    _ = p0
    _ = v7
    return -682
}
func genFn1(p0 uint16, p1 int, p2 uint8) uint32 {
    p2 -= p2
    if (false || (100 != -62)) {
        p0 = (p0 % ((p0) | 1))
    } else {
        p2 = 128
    }
    var v9 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    var v4 uint8 = 140
    v5 := []int{}
    var v10 uint64 = 0
    fmt.Println(v10, v10, v10)
    var v11 bool = (3.14159 == -2.25)
    _ = len(v1)
    v12 := GenS1{F0: (((("a" + ("" + "")) + "") == "\n\t") && (2147483647 >= -647)), F1: 170}
    v2 = v10
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int8 = 100
            _ = fi14
            _ = v15
        }
        var v16 uint32 = 1
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int16 = -364
            switch (fi17 * (^(fi13))) {
            case 32768:
                var v19 uint64 = 2
                _ = v19
            default:
                fmt.Println(v11, v4)
            }
            _ = fi17
            _ = v18
        }
        _ = fi13
        _ = v16
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 float32 = 1.5
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v10 = (v2 - (^(v2)))
    v11 = (753 != 303)
    var v23 string = (("a" + "a") + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
