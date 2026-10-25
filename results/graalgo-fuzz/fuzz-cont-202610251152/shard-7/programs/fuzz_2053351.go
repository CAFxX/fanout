package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 string }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    var v1 int16 = -248
    v2 := make(map[string]int)
    var v3 int16 = (v1 >> ((18446744073709551615) & 63))
    v1 = v3
    var v4 float32 = 0.0
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 += -577
        _ = fi6
    }
    var v7 int32 = 2147483647
    if (42 < 1) {
        fmt.Println(v4, v5, v5)
        v8 := []int{4294967295, 128}
        fmt.Println(v7, v3, v4)
        _ = v8
    }
    if (true || true) {
        var v9 int32 = (v7 & v5)
        if (18446744073709551615 < 42) {
            for fi10 := 0; fi10 < 18; fi10++ {
                v1 = 1000
                if (!(4294967295 < 0)) {
                    fmt.Println(v3, v5, v7)
                } else {
                    v4 = 0.1
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v4 = v4
                        fmt.Println(v7)
                        var v12 uint64 = 220
                        _ = fi11
                        _ = v12
                    }
                }
                fmt.Println(v4)
                _ = fi10
            }
        }
        var v13 uint = 0
        _ = v13
        _ = v9
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if true {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint = 0
                var v17 string = "hello"
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 int8 = 11
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 uint8 = 42
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            switch 32768 {
            case -2147483648:
                for fi21 := 0; fi21 < 3; fi21++ {
                    switch "foo bar" {
                    case "foo bar":
                        v22 := GenS1{F0: 919, F1: 775, F2: "foo bar"}
                        _ = v22
                    case "\n\t":
                        v7 = v5
                    default:
                        _ = len(v2)
                    }
                    v1 -= 42
                    switch fi21 {
                    case 806:
                        v1 += -32768
                    case 4294967295:
                        v23 := []int{-1, 9223372036854775807}
                        _ = v23
                    }
                    _ = fi21
                }
            default:
                v24 := v2[("hello" + "foo bar")]
                _ = v24
            }
            v7 = (v7 << ((519) & 63))
        }
        _ = fi14
    }
    if (((true || false) && (!(32767 == 4294967295))) || (!true)) {
        switch "foo bar" {
        case "a":
            switch 32767 {
            case -486:
                fmt.Println(v5, v5, v1)
            case 127:
                var v25 int32 = 122
                _ = v25
            case 0:
                var v26 bool = ("foo bar" <= "hello")
                _ = v26
            }
        case "\n\t":
            v27 := []int{128, -9223372036854775808, -1}
            _ = v27
        default:
            var v28 float32 = 0.0
            _ = v28
        }
        if false {
            var v29 float32 = v4
            var v30 float32 = v29
            _ = v29
            _ = v30
        }
    } else {
        var v31 uint16 = 2
        fmt.Println(v5, v7)
        v4 += v4
        _ = v31
    }
    switch "" {
    case "\n\t":
        fmt.Println(v4, v1, v5)
    case "a":
        var v32 uint8 = 1
        _ = v32
    }
    v33 := []int{127, 974}
    v4 -= (v4 - ((3.14159 * v4) * -2.25))
    _ = len(v2)
    delete(v2, ("" + ""))
    fmt.Println(v5, v1)
    v34 := GenS2{F0: 42, F1: 42}
    v2["a"] = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v7
}
