package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = -594
    v2 := []int{-9223372036854775808}
    var v3 uint8 = 42
    v4 := make(map[string]string)
    var v5 string = ("foo bar" + "\n\t")
    if false {
        if true {
            var v6 float64 = ((3.14159 + (0.1 - 1e10)) - (3.14159 / ((3.14159 / 1.0) - (0.0 * (1e10 - 0.1)))))
            _ = v6
        }
    } else {
        var v7 uint64 = uint64(v3)
        v5 = (v5 + "\n\t")
        var v8 uint64 = (v7 % ((v7) | 1))
        _ = v7
        _ = v8
    }
    v1 = (v1 / ((1) | 1))
    switch 42 {
    case -954:
        for fi9 := 0; fi9 < 3; fi9++ {
            v1 -= (v1 ^ (v1 - (v1 | 2147483647)))
            _ = fi9
        }
    case 0:
        v5 = v5
    }
    fmt.Println(v3, v5, v5)
    for fi10 := 0; fi10 < 4; fi10++ {
        if false {
            v11 := GenS1{F0: v5}
            _ = v11
        }
        _ = len(v2)
        for fi12 := 0; fi12 < 3; fi12++ {
            v3 += (v3 << ((18446744073709551615) & 63))
            _ = fi12
        }
        _ = fi10
    }
    switch "hello" {
    case "hello":
        if false {
            v1 += -955
            var v13 float64 = -2.25
            _ = v13
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 string = v5
                var v16 uint16 = 444
                _ = fi14
                _ = v15
                _ = v16
            }
            if false {
                v17 := GenS1{F0: ("a" + v5)}
                _ = v17
            } else {
                v4[v5] = (v5 + "hello")
                fmt.Println(v5, v1, v1)
            }
            var v18 int8 = 31
            _ = v18
        }
    case "foo bar":
        for fi19 := 0; fi19 < 7; fi19++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi19
        }
    case "a":
        fmt.Println(v3, v5)
    default:
        fmt.Println(v3, v3)
    }
    v1 -= (-(v1))
    if len(v2) > 0 {
        v2[0] = 42
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 string = ""
            _ = fi21
            _ = v22
        }
        var v23 int64 = -9223372036854775808
        _ = fi20
        _ = v23
    }
    fmt.Println(v5, v3, v3)
    delete(v4, v5)
    v1 = v1
    v1 += (v1 % ((v1) | 1))
    fmt.Println(v1, v3, v3)
    v24 := GenS1{F0: "hello"}
    switch -556 {
    case 728:
        var v25 int32 = (v1 ^ (v1 + v1))
        _ = v25
    case 2147483647:
        fmt.Println(v3, v1)
    case -209:
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 string = ""
            var v28 uint32 = 175
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        var v29 int16 = -32768
        _ = v29
    }
    switch -606 {
    case 498:
        v3 = 228
    case 2147483647:
        for fi30 := 0; fi30 < 8; fi30++ {
            v1 -= v1
            var v31 int16 = 359
            _ = fi30
            _ = v31
        }
    case -9223372036854775808:
        fmt.Println(v1, v1, v5)
    default:
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
