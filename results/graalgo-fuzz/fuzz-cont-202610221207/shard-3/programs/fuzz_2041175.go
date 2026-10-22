package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8; F2 bool }
func genFn0(p0 uint, p1 int, p2 int8) int16 {
    v5 := []int{p1, (p1 / (((p1 * (p1 ^ 42))) | 1)), p1}
    switch ("a" + "foo bar") {
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "a":
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                if len(v5) > 0 {
                    v5[0] = (p1 ^ (fi7 | fi6))
                }
                if (65535 <= 1) {
                    p1 = 295
                    v8 := GenS1{F0: 1}
                    var v9 uint32 = 863
                    _ = v8
                    _ = v9
                }
                var v10 int64 = 0
                _ = fi7
                _ = v10
            }
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1.5
    v3 := make(map[int]int)
    var v4 uint = 460
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS1{F0: 2}
        _ = fi11
        _ = v12
    }
    v13 := []int{1, 4294967295, 42}
    if true {
        if (!((4294967295 < 318) || true)) {
            if ("foo bar" == ("a" + ("\n\t" + "foo bar"))) {
                var v14 int16 = 1000
                fmt.Println(v2, v14)
                _ = v14
            }
            if true {
                var v15 bool = true
                _ = v15
            } else {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v4 = 1
                    v1[1] = (32767 != 0)
                    _ = fi16
                }
            }
        }
    } else {
        v17 := GenS2{F0: 562, F1: 218, F2: true}
        _ = v17
    }
    if ("\n\t" >= "a") {
        var v18 uint16 = 1
        _ = v18
    }
    var v19 uint8 = 96
    var v20 uint32 = 4294967295
    switch "" {
    case "hello":
        var v21 uint = v4
        _ = v21
    }
    v19 += 127
    delete(v1, v20)
    fmt.Println(v2, v4)
    v22 := GenS2{F0: 42, F1: v19, F2: (-1 <= 124)}
    v19 -= v19
    var v23 int64 = 9223372036854775807
    for fi24 := 0; fi24 < 6; fi24++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        var v25 int32 = 0
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 int8 = 42
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
