package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }

func main() {
    v1 := []string{"", ("\n\t" + "a"), "foo bar"}
    v2 := make(map[uint32]int)
    var v3 string = "hello"
    var v4 int = -696
    v5 := make(map[uint32]uint32)
    var v6 int8 = 1
    fmt.Println(v4)
    if (false && false) {
        var v7 string = "\n\t"
        v6 = v6
        _ = v7
    } else {
        switch v3 {
        case "":
            fmt.Println(v6, v4, v6)
        default:
            fmt.Println(v4, v4, v6)
        }
        switch v4 {
        case -95:
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int = (v4 / (((v4 - (v4 * 42))) | 1))
                _ = fi8
                _ = v9
            }
        case -262:
            if len(v1) > 0 {
                v1[0] = ("\n\t" + ("\n\t" + "hello"))
            }
        }
        v3 = v3
    }
    switch v3 {
    case "a":
        switch 1 {
        case -2147483648:
            var v10 string = "foo bar"
            _ = v10
        case 42:
            var v11 int = -1
            _ = v11
        case 583:
            fmt.Println(v6, v6)
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                v6 -= (-(v6))
                var v13 string = v3
                v5[198] = 556
                _ = fi12
                _ = v13
            }
        }
    }
    var v14 int32 = 2147483647
    v14 -= v14
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v6, v4)
        for fi16 := 0; fi16 < 10; fi16++ {
            if false {
                var v17 float32 = (1.5 + -2.25)
                var v18 uint = 1
                v4 = v4
                _ = v17
                _ = v18
            } else {
                v6 = (v6 * (v6 * v6))
            }
            fmt.Println(v6)
            _ = fi16
        }
        _ = fi15
    }
    v3 = v3
    if true {
        v6 = (v6 / ((v6) | 1))
    } else {
        if true {
            var v19 uint8 = 9
            v19 = v19
            var v20 uint16 = 967
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: true, F1: -1, F2: 2}
            _ = v21
        }
        var v22 int8 = (v6 - v6)
        var v23 int32 = v14
        _ = v22
        _ = v23
    }
    var v24 uint = 2
    var v25 int16 = 1000
    fmt.Println(v4)
    v3 = v3
    var v26 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
