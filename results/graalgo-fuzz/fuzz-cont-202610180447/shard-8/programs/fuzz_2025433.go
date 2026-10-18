package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int8 = -128
    var v2 int32 = 2147483647
    var v3 bool = (v1 == (v1 / ((v1) | 1)))
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: -879, F1: -1, F2: 4294967295}
        var v9 int8 = -120
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 bool = (v3 || true)
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    delete(v5, "foo bar")
    if true {
        var v12 int16 = 677
        _ = v12
    } else {
        fmt.Println(v1, v3, v3)
        fmt.Println(v2, v3)
        switch ("hello" + ("\n\t" + ((("foo bar" + "foo bar") + ("" + "foo bar")) + "foo bar"))) {
        case "a":
            v3 = true
        }
    }
    v13 := GenS1{F0: -1, F1: 1, F2: 337}
    if v3 {
        v14 := GenS2{F0: (v2 * -49), F1: 1}
        v15 := GenS2{F0: v2, F1: 253}
        _ = v14
        _ = v15
    }
    switch 42 {
    case -902:
        var v16 uint32 = 115
        _ = v16
    }
    v2 += v2
    var v17 uint8 = 2
    if v3 {
        if v3 {
            fmt.Println(v2)
        }
        if (!v3) {
            if (65535 <= 255) {
                v17 = v17
            } else {
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 bool = v3
                    fmt.Println(fi18)
                    fmt.Println(v3, v2)
                    _ = fi18
                    _ = v19
                }
            }
        }
        fmt.Println(v3)
    }
    v20 := []int{9223372036854775807, -9223372036854775808}
    _ = len(v20)
    var v21 uint = 42
    var v22 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
