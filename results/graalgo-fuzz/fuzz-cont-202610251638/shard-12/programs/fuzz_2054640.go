package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint; F1 bool }

func main() {
    v1 := []uint32{}
    var v2 string = "foo bar"
    var v3 float32 = float32(float64(float32((0.0 + (-2.25 * 0.0)))))
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    v6 := []int{}
    var v7 int64 = -9223372036854775808
    _ = len(v6)
    if (true || false) {
        for fi8 := 0; fi8 < 12; fi8++ {
            switch (("hello" + (((("foo bar" + "") + ((("hello" + v2) + "foo bar") + (((v2 + v2) + v2) + v2))) + "foo bar") + "a")) + v2) {
            case "\n\t":
                var v9 float32 = (float32(1.5) * v3)
                _ = v9
            }
            _ = fi8
        }
        v7 = (v7 & (v7 + (^(v7))))
        var v10 int64 = -9223372036854775808
        _ = v10
    } else {
        v11 := v5[((v2 + v2) + v2)]
        var v12 int = (v11 & (v11 >> ((2) & 63)))
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 0, F1: v2}
    var v14 uint8 = 127
    v7 = v7
    var v15 float32 = v3
    v7 = (v7 + (v7 - v7))
    v16 := v5[v2]
    _ = len(v4)
    switch (^(v16)) {
    case 71:
        v17 := GenS1{F0: 1, F1: v2}
        _ = v17
    case -1:
        v18 := GenS2{F0: 217, F1: true}
        _ = v18
    case -136:
        v19 := GenS1{F0: 1000, F1: "foo bar"}
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v5[v2] = fi20
        _ = fi20
    }
    if (false || true) {
        v16 = (-(v16))
    } else {
        var v21 float32 = 3.14159
        _ = v21
    }
    fmt.Println(v16, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
