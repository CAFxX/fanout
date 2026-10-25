package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int }
func genFn0(p0 int16) int32 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 4294967295, F1: 32767}
        v8 := GenS1{F0: 157, F1: 2}
        p0 += 48
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: 42, F1: 2}
        _ = fi9
        _ = v10
    }
    p0 += p0
    v11 := GenS2{F0: 2147483647}
    _ = p0
    _ = v11
    return 805
}

func main() {
    var v1 float32 = (float32(0.0) / 1.0)
    var v2 uint8 = 0
    v3 := []int{-1, 0, -1}
    v4 := []uint32{760, 42, 2147483647, uint32(v2)}
    var v5 float32 = v1
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(v1, fi12, v1)
        _ = fi12
    }
    fmt.Println(v1)
    v1 += 3.14159
    v2 = v2
    switch 127 {
    case -2147483648:
        if true {
            switch "hello" {
            case "foo bar":
                v13 := GenS2{F0: 9223372036854775807}
                _ = v13
            case "":
                v2 -= 16
            }
        }
    case 127:
        v1 += 3.14159
    case 863:
        v2 = (^(v2))
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v5 = -2.25
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 uint32 = 2147483647
            var v17 int64 = 42
            if len(v4) > 0 {
                v4[0] = (v16 * (^(v16)))
            }
            var v18 uint16 = 85
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 4294967295, F1: 18}
        _ = fi14
        _ = v19
    }
    switch (("\n\t" + ("a" + "hello")) + ("foo bar" + "foo bar")) {
    case "\n\t":
        if false {
            v5 -= v5
        }
    case "a":
        v20 := GenS2{F0: 32768}
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int16 = -1
            switch "hello" {
            case "a":
                if false {
                    fmt.Println(v5, v23)
                    v2 -= v2
                    switch ("" + "\n\t") {
                    case "hello":
                        var v24 uint8 = 42
                        _ = v24
                    case "a":
                        v1 -= (((v1 / 1.5) / 1e10) * 3.14159)
                    case "\n\t":
                        var v25 int8 = 38
                        _ = v25
                    default:
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v23, v2, v2)
                        case "":
                            v1 += -2.25
                        default:
                            var v26 int16 = (v23 | v23)
                            _ = v26
                        }
                    }
                } else {
                    fmt.Println(v1, v2, v23)
                    var v27 uint32 = 65535
                    _ = len(v4)
                    _ = v27
                }
            case "\n\t":
                v28 := GenS2{F0: fi22}
                _ = v28
            default:
                v29 := GenS1{F0: 42, F1: 42}
                _ = v29
            }
            switch "\n\t" {
            case "\n\t":
                v1 = v1
                fmt.Println(fi21, fi22, fi22)
                v2 = v2
            default:
                v4 = append(v4, 42)
            }
            _ = fi22
            _ = v23
        }
        var v30 uint16 = 42
        _ = fi21
        _ = v30
    }
    switch -1 {
    case -1:
        var v31 uint = 364
        _ = v31
    case 42:
        v1 = (3.14159 + 0.1)
    case 1:
        for fi32 := 0; fi32 < 9; fi32++ {
            fmt.Println(v5, fi32, fi32)
            v5 = v5
            _ = fi32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
