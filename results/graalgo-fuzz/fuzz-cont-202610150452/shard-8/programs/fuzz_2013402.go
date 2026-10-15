package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 1
    v2 := []int{-855, 276, 4294967295, 128}
    var v3 int = 661
    for fi4 := 0; fi4 < 20; fi4++ {
        if (100 > -128) {
            fmt.Println(fi4, v1, fi4)
            v1 = 127
            v1 = (v1 << ((18446744073709551615) & 63))
        }
        _ = fi4
    }
    var v5 int8 = -1
    var v6 uint8 = v1
    fmt.Println(v3, v6)
    v2 = append(v2, v3)
    v6 += (^(v6))
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v2 = append(v2, (v3 % ((v3) | 1)))
            _ = fi8
        }
        _ = fi7
    }
    v9 := GenS1{F0: 657}
    if true {
        if true {
            _ = len(v2)
            var v10 int64 = -9223372036854775808
            var v11 float32 = 1.5
            _ = v10
            _ = v11
        }
        fmt.Println(v1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
