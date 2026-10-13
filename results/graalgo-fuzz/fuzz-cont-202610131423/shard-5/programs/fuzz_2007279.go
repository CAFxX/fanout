package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0() int16 {
    v6 := []int{128, 127}
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 9223372036854775807, F1: (fi7 != (fi7 ^ 32767)), F2: 1}
        _ = fi7
        _ = v8
    }
    var v9 uint = 705
    _ = v6
    _ = v9
    return 1000
}
func genFn1(p0 int8, p1 int) bool {
    var v10 int = (p1 - p1)
    switch "foo bar" {
    case "\n\t":
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := []int{-166, (p1 << ((42) & 63)), v10, v10}
            p1 = 380
            var v13 uint16 = 498
            _ = fi11
            _ = v12
            _ = v13
        }
    case "":
        switch ("" + "foo bar") {
        case "":
            var v14 int32 = 1
            _ = v14
        case "foo bar":
            fmt.Println(v10, p0)
        case "a":
            var v15 uint16 = 1
            _ = v15
        }
    case "a":
        if (0 > -32768) {
            for fi16 := 0; fi16 < 9; fi16++ {
                switch -2147483648 {
                case -200:
                    fmt.Println(p1)
                case 1:
                    v17 := GenS1{F0: -9223372036854775808, F1: true, F2: 1}
                    _ = v17
                case -2147483648:
                    var v18 int32 = 138
                    _ = v18
                default:
                    v19 := GenS1{F0: 9223372036854775807, F1: (-32768 == 1000), F2: 128}
                    _ = v19
                }
                _ = fi16
            }
        }
    default:
        var v20 uint8 = 127
        _ = v20
    }
    v21 := []int{v10, 32768, -421, -98}
    v22 := GenS1{F0: -1, F1: true, F2: 8}
    _ = p0
    _ = p1
    _ = v10
    _ = v21
    _ = v22
    return true
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]int)
    var v3 int = -9223372036854775808
    v4 := make(map[string]int)
    var v5 int = (v3 << ((20) & 63))
    v5 -= v3
    var v23 int8 = 42
    fmt.Println(v5)
    v23 += (-(v23))
    v3 = (v3 % ((v5) | 1))
    v1 -= (v1 ^ v1)
    fmt.Println(v23, v5)
    v23 = v23
    var v24 float32 = 1e10
    var v25 bool = false
    v24 = float32(1e10)
    v26 := v2[0]
    v5 = (v26 >> ((673) & 63))
    var v27 uint32 = 2
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            switch v26 {
            case 128:
                v30 := []int{v3, (fi29 - fi28), fi29}
                _ = v30
            default:
                v25 = v25
            }
            switch ("\n\t" + "hello") {
            case "foo bar":
                v5 = (v3 - -1)
            case "":
                delete(v4, "")
            }
            _ = fi29
        }
        _ = fi28
    }
    switch (v26 << ((1) & 63)) {
    case 63:
        v31 := GenS2{F0: 65535, F1: (v23 | (^(v23)))}
        _ = v31
    case -281:
        v1 -= v1
    case 128:
        v32 := v4["a"]
        _ = v32
    }
    var v33 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
