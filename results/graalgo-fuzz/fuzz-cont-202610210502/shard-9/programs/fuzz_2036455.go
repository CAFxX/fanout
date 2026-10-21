package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 42
    v3 := []string{"a", "a"}
    v4 := []int{1, -2147483648, -774}
    var v5 uint16 = 686
    var v6 int16 = -32768
    delete(v1, 42)
    fmt.Println(v2, v5, v6)
    switch 32767 {
    case 978:
        if ((v6 ^ v6) > v6) {
            fmt.Println(v2)
            v5 -= (^(v5))
            var v7 int8 = -1
            _ = v7
        } else {
            var v8 uint32 = 4294967295
            fmt.Println(v2, v2, v5)
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        if true {
            v10 := GenS1{F0: (v5 << ((533) & 63))}
            _ = v10
        } else {
            var v11 uint16 = 32767
            _ = v11
        }
        var v12 int8 = 71
        _ = fi9
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v6, v2, v5)
        for fi14 := 0; fi14 < 5; fi14++ {
            _ = len(v3)
            _ = fi14
        }
        _ = fi13
    }
    switch 15 {
    case -1:
        v6 = (v6 & (v6 << ((4294967295) & 63)))
    case 4294967295:
        if (false && true) {
            fmt.Println(v6, v5)
            var v15 int64 = 0
            _ = v15
        } else {
            v5 += v5
        }
    case 127:
        v16 := GenS1{F0: v5}
        _ = v16
    }
    fmt.Println(v6)
    fmt.Println(v6, v6)
    v5 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
