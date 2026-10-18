package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int = 42
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := []int{-1, v8, v8}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if (255 <= 18446744073709551615) {
        var v11 uint64 = 2
        _ = v11
    } else {
        var v12 int16 = 1
        _ = v12
    }
    v13 := GenS1{F0: 1}
    switch (("a" + "\n\t") + ("hello" + ("" + "foo bar"))) {
    case "hello":
        var v14 uint8 = 100
        _ = v14
    case "a":
        var v15 string = "\n\t"
        _ = v15
    case "\n\t":
        v16 := GenS1{F0: 18446744073709551615}
        _ = v16
    }
    _ = v13
    return 255
}

func main() {
    v1 := []int{355, 1, 2147483647, 9223372036854775807}
    var v2 uint8 = 161
    var v3 uint32 = 393
    v4 := make(map[string]uint32)
    var v5 uint64 = 2
    var v6 uint8 = (v2 & (v2 * 2))
    var v17 int64 = -576
    fmt.Println(v5, v3, v5)
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS1{F0: uint(v3)}
            _ = len(v1)
            _ = fi19
            _ = v20
        }
        var v21 int32 = 1
        for fi22 := 0; fi22 < 18; fi22++ {
            delete(v4, "\n\t")
            _ = fi22
        }
        _ = fi18
        _ = v21
    }
    v2 = v6
    var v23 int = 32768
    v23 += (v23 * 2147483647)
    for fi24 := 0; fi24 < 19; fi24++ {
        switch "\n\t" {
        case "a":
            switch "hello" {
            case "\n\t":
                fmt.Println(v23, v3)
            }
        case "foo bar":
            var v25 string = "\n\t"
            _ = v25
        case "hello":
            switch ("" + "foo bar") {
            case "a":
                var v26 float64 = 0.0
                _ = v26
            }
        default:
            v5 = (v5 + 1)
        }
        var v27 uint64 = (v5 - v5)
        v3 = 2147483648
        _ = fi24
        _ = v27
    }
    v2 = (v6 | (v2 >> ((65535) & 63)))
    for fi28 := 0; fi28 < 12; fi28++ {
        fmt.Println(v6, v23)
        var v29 uint32 = v3
        var v30 uint32 = uint32(v5)
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
