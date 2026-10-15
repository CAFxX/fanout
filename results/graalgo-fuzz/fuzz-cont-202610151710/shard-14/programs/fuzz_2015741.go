package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint32) uint16 {
    var v5 uint = 48
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    if false {
        switch -2147483648 {
        case 128:
            fmt.Println(p0, v5)
        }
    } else {
        switch "hello" {
        case "foo bar":
            v5 -= v5
        case "":
            v5 += (v5 & 4294967295)
        case "a":
            v8 := []int{127, 4294967295}
            _ = v8
        default:
            var v9 uint16 = 42
            _ = v9
        }
        var v10 float32 = (-2.25 * -2.25)
        v11 := []int{2147483647}
        _ = v10
        _ = v11
    }
    p0 -= p0
    _ = p0
    _ = v5
    return 255
}

func main() {
    var v1 uint8 = 128
    var v2 uint = 4294967295
    var v3 bool = (2 == 382)
    v4 := []string{"\n\t"}
    v2 = (v2 & 18446744073709551615)
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v1, v2, v3)
        _ = fi12
    }
    var v13 int64 = 1
    v2 = (v2 ^ (v2 & (v2 | v2)))
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            v2 += 667
            _ = len(v4)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v16 uint64 = 362
            switch ("foo bar" + "foo bar") {
            case "hello":
                fmt.Println(v16, v13)
            }
            if v3 {
                var v17 int8 = 42
                _ = v17
            }
            v18 := GenS1{F0: fi15}
            _ = fi15
            _ = v16
            _ = v18
        }
        _ = fi14
    }
    v2 = v2
    switch ("\n\t" + (("a" + (("hello" + "\n\t") + "")) + ("\n\t" + "\n\t"))) {
    case "\n\t":
        v2 += 295
    case "foo bar":
        v19 := GenS2{F0: 2147483647, F1: 32767}
        _ = v19
    case "hello":
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 float32 = (0.1 - 3.14159)
            var v22 uint8 = (v1 * (v1 * v1))
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 uint8 = v1
    if v3 {
        if (!((0 != 65535) && true)) {
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 uint32 = 0
                _ = fi24
                _ = v25
            }
        } else {
            var v26 int16 = -1
            v27 := GenS2{F0: 807, F1: v26}
            v4 = append(v4, "\n\t")
            _ = v26
            _ = v27
        }
        var v28 int = 4294967295
        _ = v28
    } else {
        v29 := GenS2{F0: 42, F1: -32768}
        _ = v29
    }
    v30 := GenS2{F0: uint32(v1), F1: 42}
    v4 = append(v4, ("foo bar" + "foo bar"))
    v2 += (v2 * 195)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
}
