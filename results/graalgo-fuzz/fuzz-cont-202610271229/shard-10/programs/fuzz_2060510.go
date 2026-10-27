package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint8 = 181
    v2 := []string{"foo bar"}
    var v3 float64 = 0.0
    var v4 int8 = 127
    var v5 string = "\n\t"
    var v6 int8 = 1
    fmt.Println(v6, v6, v5)
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 805, F1: 2147483647}
        v9 := GenS1{F0: 2, F1: 1}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch v5 {
    case "foo bar":
        var v10 uint16 = 32767
        _ = v10
    default:
        fmt.Println(v6, v4)
    }
    switch (("foo bar" + (v5 + v5)) + "\n\t") {
    case "a":
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int16 = 1000
            _ = fi11
            _ = v12
        }
    }
    _ = len(v2)
    fmt.Println(v5, v4, v4)
    _ = len(v2)
    if (!(0 > 352)) {
        v13 := GenS2{F0: 4294967295, F1: 2}
        _ = v13
    }
    v14 := GenS3{F0: false, F1: -1}
    var v15 int = 32767
    v16 := GenS1{F0: 775, F1: 2}
    switch v15 {
    case -9223372036854775808:
        for fi17 := 0; fi17 < 2; fi17++ {
            if false {
                fmt.Println(v3)
            } else {
                fmt.Println(v4, v1)
                switch (("\n\t" + v5) + ("hello" + v5)) {
                case "":
                    v18 := GenS2{F0: 4294967295, F1: 1}
                    _ = v18
                default:
                    v19 := GenS1{F0: 0, F1: 65535}
                    _ = v19
                }
                if (v1 == v1) {
                    if len(v2) > 0 {
                        v2[0] = v5
                    }
                    fmt.Println(v15, v4)
                } else {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        switch (fi17 / ((934) | 1)) {
                        case 128:
                            v21 := GenS2{F0: 2, F1: uint(v1)}
                            _ = v21
                        }
                        _ = fi20
                    }
                }
            }
            var v22 uint32 = 42
            _ = fi17
            _ = v22
        }
    case 355:
        var v23 float32 = float32((0.0 * v3))
        _ = v23
    case -160:
        if (0 > 247) {
            if false {
                v3 -= (1.5 - 3.14159)
                var v24 float32 = 3.14159
                v24 += 3.14159
                _ = v24
            }
            v2 = append(v2, v5)
        }
    default:
        fmt.Println(v15, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
