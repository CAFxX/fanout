package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint8 = 194
    v2 := []string{}
    var v3 uint = 18446744073709551615
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = v4
    var v6 int = -769
    v7 := GenS1{F0: 65535, F1: 460}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 657, F1: 42}
        v10 := GenS1{F0: 42, F1: -9223372036854775808}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint = 972
    v12 := GenS1{F0: 2, F1: 9223372036854775807}
    v3 = (v11 - (v11 - 15))
    v11 = v3
    for fi13 := 0; fi13 < 2; fi13++ {
        v6 -= v6
        var v14 uint32 = 2147483647
        v5 += uint64(v1)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v3 -= v11
        _ = fi15
    }
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "\n\t")
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
