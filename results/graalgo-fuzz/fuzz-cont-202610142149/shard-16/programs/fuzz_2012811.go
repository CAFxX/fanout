package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v4 := []string{(((("\n\t" + "") + ("hello" + "a")) + "foo bar") + "\n\t"), "\n\t", "a"}
    var v5 float64 = -2.25
    v5 = 1e10
    fmt.Println(v1, v5)
    v1 = v1
    fmt.Println(v1)
    if false {
        if true {
            var v6 int64 = -9223372036854775808
            fmt.Println(v6, v5, v1)
            _ = v6
        } else {
            v3[("foo bar" + ("foo bar" + "foo bar"))] = 4294967295
        }
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        if (0 <= 42) {
            v3[(("hello" + ("\n\t" + "hello")) + "\n\t")] = 2
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v8 float64 = ((v5 / float64((1e10 / 1e10))) * (v5 * 0.1))
            switch fi7 {
            case 42:
                v5 = v8
            }
            _ = v8
        }
        v9 := GenS1{F0: 0}
        _ = fi7
        _ = v9
    }
    v1 = 51
    if false {
        switch "a" {
        case "a":
            v1 = v1
        case "\n\t":
            if true {
                v10 := v2[32767]
                var v11 int32 = 2147483647
                v12 := GenS1{F0: 1}
                _ = v10
                _ = v11
                _ = v12
            }
        default:
            v1 -= (v1 / ((v1) | 1))
        }
    } else {
        var v13 bool = (-1 > -388)
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 float64 = -2.25
        v16 := GenS2{F0: -1, F1: 0}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 1, F1: -1}
    if false {
        var v18 int32 = 0
        _ = v18
    }
    v5 = 0.1
    _ = len(v3)
    if true {
        v5 = v5
    } else {
        switch 42 {
        case 42:
            var v19 string = "foo bar"
            _ = v19
        }
        v20 := v2[0]
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(v5)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
