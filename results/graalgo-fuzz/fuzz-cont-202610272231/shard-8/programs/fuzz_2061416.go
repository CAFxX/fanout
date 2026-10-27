package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }

func main() {
    var v1 float64 = 1e10
    var v2 bool = (false && false)
    v3 := []string{"\n\t", "foo bar", ("" + "a"), ("\n\t" + (("a" + "foo bar") + "foo bar"))}
    var v4 uint = 0
    v1 -= ((v1 + (v1 * v1)) - float64(3.14159))
    v5 := GenS1{F0: (v4 + v4)}
    v4 += v4
    if (2147483647 <= 2147483648) {
        var v6 uint = 546
        _ = v6
    } else {
        v7 := GenS1{F0: (v4 % ((42) | 1))}
        _ = v7
    }
    var v8 uint = (^(v4))
    for fi9 := 0; fi9 < 1; fi9++ {
        if (!v2) {
            if v2 {
                var v10 uint32 = 65535
                switch "hello" {
                case "":
                    v11 := GenS1{F0: (v4 ^ uint(v10))}
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 18; fi12++ {
                        v1 -= v1
                        if (((((v1 - v1) >= v1) || v2) && v2) && v2) {
                            v1 = v1
                        }
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            var v13 float64 = ((v1 * 1.5) * (v1 * (v1 + ((v1 - -2.25) + (-2.25 * v1)))))
                            _ = v13
                        case "":
                            v14 := GenS1{F0: (v4 & 42)}
                            _ = v14
                        default:
                            fmt.Println(fi12, v10, fi9)
                        }
                        _ = fi12
                    }
                }
                _ = v10
            } else {
                var v15 int = (fi9 + (^(fi9)))
                _ = v15
            }
            if (!(((1000 < -1) && (118 > 4294967295)) || (1 != 0))) {
                if v2 {
                    var v16 int32 = 0
                    var v17 uint8 = 1
                    _ = v16
                    _ = v17
                }
            } else {
                fmt.Println(v4, v2, v2)
                var v18 int8 = 100
                _ = v18
            }
        }
        var v19 uint = (^(v8))
        _ = fi9
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        switch (^(fi20)) {
        case 9223372036854775807:
            v4 += 2
        default:
            for fi21 := 0; fi21 < 6; fi21++ {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                if false {
                    fmt.Println(v2, v1, v2)
                    var v22 uint16 = 75
                    v23 := GenS2{F0: -9223372036854775808, F1: -32768, F2: 300}
                    var v24 uint8 = 128
                    _ = v22
                    _ = v23
                    _ = v24
                }
                fmt.Println(v8, v8)
                _ = fi21
            }
        }
        _ = fi20
    }
    if false {
        fmt.Println(v4, v4, v2)
    }
    var v25 float32 = ((1.5 * (-2.25 - -2.25)) - -2.25)
    _ = len(v3)
    var v26 uint8 = 42
    var v27 uint32 = 2147483647
    for fi28 := 0; fi28 < 7; fi28++ {
        v2 = v2
        v4 = v8
        _ = fi28
    }
    v29 := GenS1{F0: (v8 - (v8 - v4))}
    var v30 uint32 = v27
    v26 -= v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
