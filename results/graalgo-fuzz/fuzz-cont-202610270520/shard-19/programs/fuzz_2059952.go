package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint64 = 352
    v2 := []uint32{}
    var v3 int8 = 1
    var v4 bool = true
    var v5 string = "\n\t"
    if len(v2) > 0 {
        v2[0] = 0
    }
    switch v5 {
    case "a":
        v5 = "foo bar"
    case "hello":
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 bool = false
            var v8 int64 = 0
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        v4 = v4
    }
    var v9 int32 = -845
    var v10 uint64 = (v1 >> ((0) & 63))
    v3 -= v3
    var v11 uint = 239
    var v12 int = -9223372036854775808
    var v13 bool = (552 == (v9 << (((v11 % ((18446744073709551615) | 1))) & 63)))
    for fi14 := 0; fi14 < 17; fi14++ {
        switch 128 {
        case 4294967295:
            v5 = (v5 + (v5 + "\n\t"))
        case 0:
            v4 = v13
        case 42:
            v2 = append(v2, 2147483647)
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            switch (fi15 | (v12 << (((v11 + (v11 ^ v11))) & 63))) {
            case 9223372036854775807:
                var v16 uint = v11
                _ = v16
            case 459:
                v17 := GenS1{F0: -1}
                _ = v17
            }
            var v18 int64 = 9223372036854775807
            _ = fi15
            _ = v18
        }
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v19 := GenS1{F0: (v3 << ((4294967295) & 63))}
            _ = v19
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
