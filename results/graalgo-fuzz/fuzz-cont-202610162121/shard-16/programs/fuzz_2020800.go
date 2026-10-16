package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint = 4294967295
    var v2 int64 = 42
    var v3 uint8 = 232
    var v4 int8 = 127
    v5 := make(map[uint32]int)
    v6 := []string{"a"}
    if (0 < -260) {
        for fi7 := 0; fi7 < 4; fi7++ {
            if ((1.5 + 0.1) < 1e10) {
                var v8 uint = (v1 >> ((2) & 63))
                _ = v8
            } else {
                var v9 uint8 = v3
                var v10 uint16 = 25
                _ = v10
                _ = v9
            }
            switch fi7 {
            case -2147483648:
                fmt.Println(v3)
            case 4294967295:
                switch "foo bar" {
                case "\n\t":
                    if (true && true) {
                        var v11 uint8 = v3
                        var v12 int64 = v2
                        _ = v11
                        _ = v12
                    }
                }
            default:
                var v13 float32 = 1e10
                _ = v13
            }
            v14 := v5[4294967295]
            _ = fi7
            _ = v14
        }
    } else {
        fmt.Println(v2, v1, v4)
    }
    switch 42 {
    case -651:
        v6 = append(v6, ("hello" + "a"))
    case 128:
        delete(v5, 255)
    default:
        v2 = v2
    }
    fmt.Println(v3, v3)
    for fi15 := 0; fi15 < 1; fi15++ {
        if ((v1 & 65535) > (v1 << (((^(v1))) & 63))) {
            _ = len(v6)
        } else {
            var v16 bool = false
            var v17 float64 = 1e10
            _ = len(v5)
            _ = v16
            _ = v17
        }
        v1 -= v1
        _ = fi15
    }
    v4 += v4
    _ = len(v5)
    switch -956 {
    case -997:
        if true {
            v5[63] = 2147483647
        } else {
            v3 = v3
            fmt.Println(v2)
            v18 := GenS1{F0: 4294967295}
            _ = v18
        }
    case 721:
        v19 := GenS1{F0: 128}
        _ = v19
    case -183:
        if false {
            for fi20 := 0; fi20 < 4; fi20++ {
                v2 += (v2 / (((-(v2))) | 1))
                var v21 int16 = 32767
                v22 := v5[2147483647]
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    }
    _ = len(v6)
    v23 := GenS1{F0: 0}
    v24 := GenS1{F0: 0}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
