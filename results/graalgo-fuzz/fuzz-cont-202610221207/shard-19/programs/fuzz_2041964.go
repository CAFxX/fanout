package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int8, p1 int16, p2 uint8) int16 {
    fmt.Println(p1)
    var v7 string = ((("a" + ("hello" + "\n\t")) + "foo bar") + ("" + ("a" + "a")))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1(p0 int8) uint {
    v8 := GenS1{F0: 42}
    fmt.Println(p0)
    if (4294967295 >= 79) {
        var v9 int = -556
        v9 = -613
        _ = v9
    }
    p0 = p0
    var v10 string = ""
    _ = p0
    _ = v10
    _ = v8
    return 255
}

func main() {
    v1 := []uint32{1, 0}
    var v2 int8 = -52
    var v3 uint = 65535
    var v4 uint = 255
    v5 := make(map[int]int)
    var v6 int = 0
    v11 := v5[v6]
    if false {
        for fi12 := 0; fi12 < 7; fi12++ {
            v4 += v3
            _ = len(v5)
            var v13 float32 = (1e10 - 1.5)
            _ = fi12
            _ = v13
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v1 = append(v1, 4294967295)
        var v14 float64 = ((-2.25 - (((1e10 / 1.5) - 1e10) * 1.5)) - 0.0)
        _ = v14
    }
    v6 = (v11 % ((v11) | 1))
    var v15 uint = (v3 >> ((0) & 63))
    var v16 float32 = 1.5
    v3 += v3
    if true {
        v3 = 958
        v2 -= v2
    }
    if (true && (127 > 127)) {
        v17 := GenS1{F0: -1}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 16; fi18++ {
            v6 = (v6 & (-(v6)))
            _ = fi18
        }
        var v19 bool = (((-9223372036854775808 > -9223372036854775808) && false) || false)
        _ = v19
    }
    if false {
        v6 = 9223372036854775807
    }
    v20 := GenS2{F0: 319}
    var v21 uint64 = 117
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
