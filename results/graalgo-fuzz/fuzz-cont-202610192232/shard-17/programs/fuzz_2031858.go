package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    var v3 int32 = 0
    var v4 uint16 = 284
    v5 := []string{("foo bar" + ("\n\t" + "foo bar")), "a", "foo bar"}
    v6 := make(map[string]string)
    var v7 bool = ((v3 > v3) || (!false))
    var v8 uint16 = 585
    switch -937 {
    case 140:
        v4 = (v8 - 2)
    default:
        var v9 uint32 = 4294967295
        _ = v9
    }
    fmt.Println(v7, v8)
    if true {
        fmt.Println(v1, v7, v7)
        v10 := GenS1{F0: 645, F1: uint8(v8), F2: "\n\t"}
        _ = v10
    } else {
        if v7 {
            _ = len(v2)
            fmt.Println(v4, v3, v8)
        } else {
            v4 = (v8 + v8)
        }
        v4 -= 255
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        if len(v5) > 0 {
            v5[0] = ""
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int32 = v3
            var v14 uint64 = 615
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 uint8 = 255
    v16 := GenS2{F0: 4294967295, F1: 18446744073709551615}
    fmt.Println(v1)
    if ((v3 - (v3 << ((18446744073709551615) & 63))) != 42) {
        v5 = append(v5, ("hello" + "foo bar"))
    }
    if v7 {
        v15 += v15
        var v17 uint16 = v4
        _ = v17
    } else {
        var v18 int = 9223372036854775807
        var v19 uint8 = 2
        v8 = v8
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 706, F1: 0}
    v3 = (v3 | -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
