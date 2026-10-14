package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 int8 = 127
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    var v6 float32 = (3.14159 - 3.14159)
    v3 -= 42
    var v7 uint64 = 18446744073709551615
    if false {
        v8 := GenS1{F0: (^(v7)), F1: 42, F2: 4294967295}
        switch (((("a" + ("hello" + ("" + ("" + "")))) + "a") + ("" + "foo bar")) + (("a" + "hello") + (("foo bar" + "\n\t") + "\n\t"))) {
        case "foo bar":
            if true {
                var v9 string = (((("\n\t" + ("" + ("foo bar" + "\n\t"))) + ("hello" + "\n\t")) + ("a" + "foo bar")) + "\n\t")
                if (42 != 1) {
                    switch "hello" {
                    case "foo bar":
                        var v10 uint16 = 572
                        _ = v10
                    case "hello":
                        v11 := GenS1{F0: v7, F1: -1, F2: 1}
                        _ = v11
                    case "a":
                        v3 = -107
                    }
                } else {
                    var v12 int32 = 1
                    _ = v12
                }
                v13 := []int{128, 32768, 85}
                _ = v13
                _ = v9
            }
        case "\n\t":
            v5[uint32(v7)] = 65535
        default:
            _ = len(v2)
        }
        _ = v8
    } else {
        v14 := []int{}
        v3 -= (v3 ^ (v3 / (((v3 - (v3 << ((373) & 63)))) | 1)))
        var v15 int64 = -816
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 float64 = (float64(0.1) - 0.1)
        var v20 int8 = 127
        _ = fi18
        _ = v19
        _ = v20
    }
    if (((0.1 == float64(v6)) && true) && false) {
        var v21 int8 = (v3 % (((v3 & (v3 >> ((733) & 63)))) | 1))
        _ = v21
    } else {
        v3 += v3
        var v22 int8 = v3
        v3 = -18
        _ = v22
    }
    fmt.Println(v7, v3)
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 bool = true
            _ = fi24
            _ = v25
        }
        var v26 int8 = v3
        _ = fi23
        _ = v26
    }
    if (v7 <= v7) {
        var v27 bool = true
        v28 := GenS1{F0: (v7 % ((v7) | 1)), F1: -269, F2: -9223372036854775808}
        var v29 uint16 = 1
        _ = v27
        _ = v28
        _ = v29
    } else {
        v3 = 1
        v30 := GenS1{F0: (^(v7)), F1: 0, F2: 32768}
        var v31 uint16 = 702
        _ = v30
        _ = v31
    }
    v32 := []int{128, 32767, 930}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
