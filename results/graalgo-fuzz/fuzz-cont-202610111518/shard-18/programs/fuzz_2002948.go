package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    var v4 uint32 = 2147483647
    v5 := []uint32{}
    v6 := make(map[int]int)
    var v7 int8 = 100
    if true {
        if false {
            var v8 int64 = -1
            for fi9 := 0; fi9 < 18; fi9++ {
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 float64 = (0.1 - 1.5)
                    v12 := GenS1{F0: 0, F1: (v1 ^ (v1 | (v1 - (v1 & v1)))), F2: (uint16(v4) > 2)}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v4 = (v4 * v4)
                if true {
                    v5 = append(v5, (^(v4)))
                }
                _ = fi9
            }
            _ = v8
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                switch (("a" + "\n\t") + ("hello" + "foo bar")) {
                case "\n\t":
                    if (-2.25 <= 1e10) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                for fi14 := 0; fi14 < 19; fi14++ {
                    v7 = v7
                    v1 -= 18446744073709551615
                    _ = fi14
                }
                _ = fi13
            }
        }
        var v15 int = 1
        if ((-9223372036854775808 == 9223372036854775807) || true) {
            v16 := GenS1{F0: 356, F1: (v1 | (v1 >> ((18446744073709551615) & 63))), F2: false}
            _ = v16
        } else {
            var v17 float32 = 0.0
            _ = v17
        }
        _ = v15
    }
    if (255 == 853) {
        if (113 >= -379) {
            v7 = 23
        } else {
            fmt.Println(v7, v7, v7)
            v7 = (-(v7))
            var v18 bool = true
            _ = v18
        }
        if ((0.0 / 3.14159) > (3.14159 - 1e10)) {
            fmt.Println(v3)
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 string = "a"
                var v21 float64 = 3.14159
                v22 := GenS1{F0: 392, F1: v1, F2: false}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        }
        v4 = v4
    }
    v3 = v3
    switch 2147483647 {
    case 42:
        v23 := v6[-165]
        _ = v23
    case 9223372036854775807:
        var v24 uint32 = (v4 | (v4 + 0))
        _ = v24
    default:
        var v25 uint8 = 127
        _ = v25
    }
    fmt.Println(v7, v3, v3)
    if (194 <= 54) {
        var v26 int32 = (^(v3))
        var v27 int = 128
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v26
        _ = v27
    }
    if (!false) {
        if false {
            var v28 int16 = 658
            var v29 int = 128
            _ = v28
            _ = v29
        }
    } else {
        v4 = v4
        v30 := GenS1{F0: 255, F1: v1, F2: false}
        _ = v30
    }
    v1 = 4294967295
    var v31 int16 = 48
    fmt.Println(v1, v1)
    var v32 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
