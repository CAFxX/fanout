package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0() uint64 {
    var v7 int16 = 0
    v7 = v7
    v7 += v7
    switch "foo bar" {
    case "foo bar":
        v8 := GenS1{F0: 14}
        _ = v8
    case "\n\t":
        v7 -= -32768
    case "hello":
        switch "hello" {
        case "foo bar":
            v7 = v7
        }
    default:
        v7 += (v7 | v7)
    }
    _ = v7
    return 885
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1.5
    var v3 uint16 = 65535
    v4 := []int{}
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]int)
    _ = len(v1)
    if (1 != 2) {
        var v9 uint8 = 1
        v3 -= 255
        v2 -= float32(1.5)
        _ = v9
    } else {
        var v10 uint = 42
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 uint64 = 1
                switch (-(fi12)) {
                case 1:
                    v13 += (v13 % (((^(v13))) | 1))
                    if (42 != 1) {
                        v14 := v1[uint32(v10)]
                        var v15 uint8 = 14
                        _ = v14
                        _ = v15
                    }
                case 737:
                    v16 := GenS1{F0: -128}
                    _ = v16
                default:
                    v13 += 18446744073709551615
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v17 := GenS2{F0: true, F1: 42}
        _ = v10
        _ = v17
    }
    v3 = 255
    var v18 int32 = 1
    v4 = append(v4, 32767)
    var v19 uint8 = 188
    if ((!true) || ("\n\t" >= "\n\t")) {
        fmt.Println(v3, v18)
        var v20 bool = false
        _ = v20
    }
    v2 -= v2
    var v21 bool = (v2 != -2.25)
    v22 := GenS2{F0: v21, F1: 4294967295}
    _ = len(v1)
    var v23 float32 = v2
    v18 -= -1
    v24 := GenS1{F0: 0}
    switch -9223372036854775808 {
    case 32768:
        v25 := GenS2{F0: v21, F1: 1}
        _ = v25
    case -694:
        v26 := GenS2{F0: ("\n\t" <= ""), F1: 628}
        _ = v26
    }
    v4 = append(v4, 4294967295)
    var v27 uint8 = v19
    var v28 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
