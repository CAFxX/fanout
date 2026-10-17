package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int64 = 42
    var v2 uint8 = 89
    var v3 uint = 488
    var v4 uint8 = (v2 * (v2 & (v2 % (((v2 % (((v2 + (v2 % ((v2) | 1)))) | 1))) | 1))))
    fmt.Println(v3)
    v4 = 0
    var v5 int = 4294967295
    var v6 int = (v5 % ((v5) | 1))
    fmt.Println(v3, v2)
    v6 -= v6
    v7 := []int{}
    v2 = v4
    if ((827 != 0) || (1 == 1)) {
        var v8 bool = (42 < 18446744073709551615)
        fmt.Println(v2)
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 string = "\n\t"
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    if len(v7) > 0 {
        v7[0] = -1
    }
    fmt.Println(v4)
    fmt.Println(v1, v4, v1)
    switch (v5 - (v6 % (((v5 << ((v3) & 63))) | 1))) {
    case -2147483648:
        v11 := GenS1{F0: 449}
        _ = v11
    case 42:
        var v12 string = ""
        _ = v12
    }
    switch v5 {
    case -142:
        if len(v7) > 0 {
            v7[0] = (v5 << (((v3 + 817)) & 63))
        }
    }
    if true {
        var v13 int = int(v1)
        v7 = append(v7, -2147483648)
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
