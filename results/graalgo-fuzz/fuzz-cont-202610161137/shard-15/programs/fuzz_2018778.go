package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 1
    var v3 bool = ((127 < 42) || (false || (0.0 <= 0.0)))
    v4 := []uint32{734, 410, 4294967295, 4294967295}
    var v5 string = "foo bar"
    fmt.Println(v3, v3, v5)
    var v6 int16 = -32768
    if (!(65535 >= 2147483648)) {
        var v7 uint32 = 42
        _ = v7
    }
    switch 35 {
    case 126:
        v6 = 1000
    case 42:
        v2 += (v2 - (v2 & (v2 + (v2 * 1))))
    case 2147483647:
        switch ((v5 + "hello") + "foo bar") {
        case "a":
            var v8 float64 = 0.1
            _ = v8
        case "":
            v9 := GenS1{F0: 127, F1: 9223372036854775807, F2: 1}
            _ = v9
        case "\n\t":
            for fi10 := 0; fi10 < 3; fi10++ {
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 string = v5
                    _ = fi11
                    _ = v12
                }
                var v13 int32 = v2
                _ = fi10
                _ = v13
            }
        }
    default:
        fmt.Println(v2, v5, v3)
    }
    var v14 uint = 4294967295
    switch 4294967295 {
    case -9223372036854775808:
        var v15 uint16 = 618
        _ = v15
    default:
        _ = len(v1)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int32 = 42
        v14 = (v14 ^ (v14 << ((v14) & 63)))
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 100, F1: 1, F2: -128}
    if (!v3) {
        fmt.Println(v3)
        var v19 string = v5
        v20 := GenS1{F0: 100, F1: -1, F2: 6}
        _ = v19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
