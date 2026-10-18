package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint, p1 int8) int16 {
    var v5 uint16 = 32767
    v6 := GenS1{F0: 0, F1: (p0 - (^(p0))), F2: 42}
    if (((v5 | (^(v5))) == v5) || true) {
        v7 := []int{820, -2147483648, 32768, -2147483648}
        for fi8 := 0; fi8 < 17; fi8++ {
            v5 += v5
            var v9 bool = ((p1 & int8(v5)) != -37)
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    switch ((("\n\t" + "\n\t") + "hello") + "\n\t") {
    case "foo bar":
        v5 = (v5 / ((42) | 1))
    case "\n\t":
        v10 := []int{1, -296, -73, -2147483648}
        _ = v10
    }
    v5 -= (v5 ^ v5)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -191
}
func genFn1(p0 string) uint8 {
    var v11 int32 = 576
    v12 := []int{32768}
    if false {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint32 = 2147483648
            _ = fi13
            _ = v14
        }
        fmt.Println(p0)
        if len(v12) > 0 {
            v12[0] = 127
        }
    } else {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    p0 = "foo bar"
    _ = p0
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]string)
    var v3 uint32 = 742
    v4 := make(map[int]uint32)
    var v15 int = 2147483647
    fmt.Println(v15, v1, v3)
    var v16 uint = (v1 - (v1 & (^(v1))))
    v1 += (v16 % ((69) | 1))
    v17 := []int{v15, (-(v15))}
    v3 -= v3
    v18 := GenS2{F0: 1000}
    v1 += (v16 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
