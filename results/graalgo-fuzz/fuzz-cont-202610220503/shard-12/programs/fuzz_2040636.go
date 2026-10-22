package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := []int{(fi4 << ((490) & 63))}
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 32767
    v7 := []int{127, 1, int(v6), 2147483647}
    _ = v6
    _ = v7
    return (v6 >> ((2) & 63))
}

func main() {
    var v1 int = 42
    var v2 uint64 = 425
    v3 := []uint32{2, 762, 1}
    if ((!(true && (!(!(323 <= -2147483648))))) || false) {
        v1 = v1
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 uint16 = 0
    v9 := GenS1{F0: 858}
    var v10 string = "\n\t"
    var v11 int64 = -1
    var v12 uint8 = 1
    var v13 uint32 = 2147483648
    v14 := GenS1{F0: 0}
    v11 += 42
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v13, v2)
        _ = fi15
    }
    switch v10 {
    case "\n\t":
        var v16 uint = 18446744073709551615
        _ = v16
    case "foo bar":
        if (false && true) {
            var v17 uint32 = (^(v13))
            v3 = append(v3, v13)
            fmt.Println(v11, v11, v11)
            _ = v17
        }
    default:
        var v18 int64 = 296
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
