package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = 0
    v3 := make(map[uint32]int)
    v4 := []int{-2147483648, 32767}
    _ = len(v3)
    v1 = (v1 / ((v1) | 1))
    var v5 uint32 = 295
    v6 := GenS1{F0: 0, F1: (-(v2)), F2: 42}
    switch -788 {
    case 773:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 32768:
        var v7 uint64 = 42
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if false {
            fmt.Println(fi8, fi8, v1)
        } else {
            switch ("" + ("" + ("" + "\n\t"))) {
            case "hello":
                var v9 uint32 = (v5 + (v5 & v5))
                _ = v9
            default:
                v4 = append(v4, (fi8 / ((42) | 1)))
            }
        }
        var v10 uint = 4294967295
        _ = fi8
        _ = v10
    }
    v2 += -1
    if true {
        v5 = (v5 % ((65535) | 1))
    } else {
        var v11 float64 = (-2.25 / 1.0)
        var v12 string = ((("foo bar" + "foo bar") + "hello") + ("a" + "hello"))
        fmt.Println(v5, v2)
        _ = v11
        _ = v12
    }
    v5 = v5
    if false {
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int = fi14
                switch (v15 >> ((65535) & 63)) {
                case 0:
                    v2 -= v2
                }
                var v16 float64 = 0.0
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        if true {
            v5 -= 2147483648
            var v17 uint8 = 127
            _ = v17
        } else {
            v18 := GenS1{F0: 42, F1: (v2 + (v2 >> ((0) & 63))), F2: 255}
            _ = v18
        }
        v19 := GenS2{F0: 2147483647, F1: v5}
        _ = v19
    }
    var v20 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
