package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 string }

func main() {
    var v1 int16 = -1
    var v2 uint = 255
    var v3 bool = (1e10 == -2.25)
    var v4 uint8 = 35
    var v5 uint64 = 18446744073709551615
    if false {
        var v6 int8 = 115
        var v7 uint = v2
        v3 = v3
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 79, F1: 1, F2: 42}
        _ = v8
    }
    v9 := []int{127, 32768, 128, 127}
    v1 -= v1
    v2 = v2
    _ = len(v9)
    var v10 int8 = 42
    switch "\n\t" {
    case "foo bar":
        var v11 int16 = v1
        _ = v11
    case "":
        if false {
            switch ((((("foo bar" + (("a" + "a") + ("hello" + "a"))) + ("hello" + "hello")) + (("a" + "foo bar") + ("foo bar" + ("a" + "")))) + ("hello" + "hello")) + "hello") {
            case "hello":
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 uint8 = 1
                    _ = fi12
                    _ = v13
                }
            case "":
                v10 -= v10
            }
        } else {
            fmt.Println(v2)
            var v14 int32 = 1
            for fi15 := 0; fi15 < 13; fi15++ {
                switch fi15 {
                case -281:
                    var v16 uint16 = 83
                    _ = v16
                case -1:
                    fmt.Println(v3, v3, v2)
                }
                switch (fi15 % (((fi15 % (((fi15 % ((32768) | 1))) | 1))) | 1)) {
                case 299:
                    v2 += (v2 + 18446744073709551615)
                case -718:
                    var v17 int = fi15
                    _ = v17
                default:
                    switch (fi15 % (((fi15 / ((fi15) | 1))) | 1)) {
                    case 336:
                        v2 = (v2 >> ((uint(v5)) & 63))
                    default:
                        v18 := GenS1{F0: 4, F1: uint64(v4), F2: (v5 + (v5 | 972))}
                        _ = v18
                    }
                }
                _ = fi15
            }
            _ = v14
        }
    }
    var v19 uint64 = (^(v5))
    if (v3 || false) {
        var v20 float32 = 1.5
        if v3 {
            var v21 int64 = 9223372036854775807
            v9 = append(v9, 9223372036854775807)
            v4 = 42
            _ = v21
        }
        var v22 bool = false
        _ = v20
        _ = v22
    } else {
        v23 := GenS2{F0: "foo bar"}
        v4 = 92
        var v24 int = 32768
        _ = v23
        _ = v24
    }
    v10 -= (-(v10))
    fmt.Println(v3, v1, v4)
    var v25 int32 = -1
    v9 = append(v9, 128)
    var v26 int32 = v25
    switch 32767 {
    case 9223372036854775807:
        switch 1 {
        case -859:
            var v27 float64 = 3.14159
            _ = v27
        case -300:
            v26 -= 42
        case 32767:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        default:
            v10 += 19
        }
    case 127:
        var v28 uint8 = (v4 & v4)
        _ = v28
    }
    if len(v9) > 0 {
        v9[0] = 127
    }
    fmt.Println(v3, v3)
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := GenS1{F0: 35, F1: v19, F2: (^(v5))}
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
