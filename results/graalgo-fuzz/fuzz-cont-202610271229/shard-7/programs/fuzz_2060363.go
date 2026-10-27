package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    v3 := make(map[string]bool)
    var v4 uint32 = 2147483647
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 bool = ((float64(1.5) + ((3.14159 / 1e10) - 0.0)) <= float64(1.5))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint32 = v4
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        var v9 string = "foo bar"
        _ = v9
    }
    fmt.Println(v4, v2)
    if true {
        var v10 int32 = 336
        var v11 int32 = -2147483648
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: -128, F1: 1, F2: 1}
        switch -243 {
        case 9223372036854775807:
            var v13 uint8 = 185
            _ = v13
        case 4294967295:
            v3[("" + "")] = ((v4 << ((0) & 63)) >= v4)
        case -9223372036854775808:
            if (true && true) {
                v2 = 9223372036854775807
            }
        }
        _ = v12
    }
    delete(v3, (("foo bar" + ("foo bar" + (("foo bar" + "hello") + "\n\t"))) + ""))
    v2 += v2
    fmt.Println(v2, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
