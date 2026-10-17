package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0(p0 string) uint64 {
    fmt.Println(p0)
    p0 = p0
    p0 = "a"
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            p0 = ((((p0 + "foo bar") + p0) + p0) + "\n\t")
            p0 = p0
            p0 = (p0 + (p0 + ((p0 + "foo bar") + p0)))
            _ = fi8
        }
        if (("foo bar" + "\n\t") == p0) {
            var v9 uint64 = 18446744073709551615
            var v10 int64 = 862
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    return 118
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = v1
    var v4 int8 = 0
    var v5 int = -1
    var v6 int64 = -841
    var v11 uint8 = 0
    v12 := GenS1{F0: v2, F1: (^(v11))}
    var v13 int64 = (v6 ^ (-(v6)))
    v14 := GenS1{F0: 2, F1: uint8(v2)}
    v15 := []int{579, (v5 + 32767), 128, v5}
    var v16 uint8 = (v11 & uint8(v2))
    fmt.Println(v6, v2, v16)
    for fi17 := 0; fi17 < 10; fi17++ {
        v3 = v3
        v18 := GenS1{F0: v2, F1: v16}
        _ = fi17
        _ = v18
    }
    var v19 int32 = 2147483647
    for fi20 := 0; fi20 < 12; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint8 = uint8(v3)
            _ = fi21
            _ = v22
        }
        var v23 uint8 = (^(v11))
        _ = fi20
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
