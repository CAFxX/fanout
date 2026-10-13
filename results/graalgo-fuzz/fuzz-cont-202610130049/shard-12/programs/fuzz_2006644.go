package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    var v1 float32 = float32((-2.25 - (-2.25 - 1.5)))
    v2 := make(map[int]bool)
    var v3 uint32 = 0
    v1 -= v1
    v4 := GenS1{F0: 681}
    if false {
        v2[696] = (0 <= 1)
        var v5 int64 = 0
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 float64 = -2.25
            var v8 int = fi6
            switch (fi6 / ((42) | 1)) {
            case 32768:
                var v9 int16 = 42
                _ = v9
            case 950:
                var v10 int = (v8 * 127)
                _ = v10
            case -2147483648:
                if true {
                    v11 := GenS1{F0: 2}
                    v12 := GenS1{F0: 42}
                    _ = v11
                    _ = v12
                } else {
                    v7 -= v7
                    var v13 int8 = 42
                    switch "hello" {
                    case "a":
                        if ((v8 + -970) == fi6) {
                            v14 := GenS2{F0: v3, F1: 25}
                            _ = v14
                        }
                    }
                    _ = v13
                }
            default:
                _ = len(v2)
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    }
    var v15 float32 = (((0.0 - 3.14159) * 0.1) * float32((-2.25 + 0.0)))
    var v16 int64 = 145
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int = int(v3)
        var v19 uint64 = 4294967295
        v3 = (v3 % ((440) | 1))
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            v3 += 2
        }
        fmt.Println(v1, v15)
    } else {
        v20 := GenS1{F0: 2}
        for fi21 := 0; fi21 < 2; fi21++ {
            if (((-2.25 / 1.0) > 0.1) || (false && false)) {
                v3 += v3
            } else {
                fmt.Println(v3)
                switch (("hello" + "foo bar") + "hello") {
                case "hello":
                    v16 += (^(v16))
                default:
                    v22 := []int{-1, fi21, fi21, (fi21 / ((fi21) | 1))}
                    _ = v22
                }
            }
            _ = fi21
        }
        if true {
            fmt.Println(v3, v3)
            delete(v2, 545)
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS1{F0: 18446744073709551615}
                v16 = v16
                if (("hello" > "foo bar") && (!("foo bar" < "foo bar"))) {
                    var v25 uint64 = 42
                    var v26 int8 = 42
                    _ = v25
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
        }
        _ = v20
    }
    if (806 == 1) {
        if true {
            v27 := v2[32768]
            for fi28 := 0; fi28 < 16; fi28++ {
                fmt.Println(v15, fi28, v16)
                _ = fi28
            }
            _ = v27
        }
    }
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
