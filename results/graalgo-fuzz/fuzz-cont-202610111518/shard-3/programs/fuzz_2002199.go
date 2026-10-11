package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }

func main() {
    v1 := []uint32{2147483648, 877, 65535}
    v2 := []uint32{4294967295, 4294967295, 368, 2147483648}
    var v3 uint16 = 1
    var v4 uint16 = 151
    var v5 float64 = float64(0.0)
    for fi6 := 0; fi6 < 14; fi6++ {
        switch (("foo bar" + "") + ("foo bar" + (("hello" + ("foo bar" + ("foo bar" + "\n\t"))) + "\n\t"))) {
        case "foo bar":
            var v7 string = ("" + "\n\t")
            _ = v7
        default:
            switch fi6 {
            case -9223372036854775808:
                v5 = v5
            case 128:
                v8 := GenS1{F0: 255, F1: 1, F2: 0}
                _ = v8
            default:
                v1 = append(v1, 2147483647)
            }
        }
        _ = fi6
    }
    var v9 int16 = 307
    switch -9223372036854775808 {
    case 42:
        v10 := GenS1{F0: 0, F1: 127, F2: -9223372036854775808}
        _ = v10
    case 127:
        fmt.Println(v4, v3, v3)
    case 711:
        var v11 uint = 65535
        _ = v11
    }
    fmt.Println(v3)
    fmt.Println(v9)
    if (255 < 255) {
        switch "\n\t" {
        case "foo bar":
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS1{F0: 2, F1: 36, F2: 286}
                _ = fi12
                _ = v13
            }
        case "hello":
            if true {
                var v14 float32 = 3.14159
                _ = v14
            }
        case "\n\t":
            v15 := GenS1{F0: 999, F1: -1, F2: 804}
            _ = v15
        default:
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 int64 = -522
                _ = fi16
                _ = v17
            }
        }
        var v18 uint16 = v3
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v20 := GenS1{F0: 0, F1: 0, F2: 42}
        fmt.Println(v3, v4, v5)
        var v21 int = int(v9)
        _ = fi19
        _ = v20
        _ = v21
    }
    switch "a" {
    case "\n\t":
        v5 = v5
    default:
        var v22 int16 = 202
        _ = v22
    }
    v5 -= 1e10
    v23 := GenS1{F0: 399, F1: 42, F2: -215}
    var v24 int16 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
