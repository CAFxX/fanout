package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }

func main() {
    var v1 uint64 = 507
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    var v4 int8 = 47
    for fi5 := 0; fi5 < 8; fi5++ {
        delete(v3, -2147483648)
        var v6 uint8 = 2
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(fi7 - fi7), (fi7 >> ((4294967295) & 63)), (fi7 - (fi7 << ((489) & 63))), (^(fi7))}
        _ = fi7
        _ = v8
    }
    var v9 int8 = 73
    if false {
        var v10 string = ""
        var v11 uint = 0
        var v12 uint = v11
        _ = v10
        _ = v11
        _ = v12
    }
    v9 = v4
    var v13 int32 = -704
    var v14 uint64 = (v1 & v1)
    v3[-537] = 9223372036854775807
    v15 := GenS1{F0: 255, F1: 1, F2: (v13 ^ v13)}
    v16 := v3[-928]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
