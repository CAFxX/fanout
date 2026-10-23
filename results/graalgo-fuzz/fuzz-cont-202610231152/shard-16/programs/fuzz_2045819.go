package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int32 = 0
    var v2 uint16 = 65535
    var v3 bool = false
    var v4 uint64 = 231
    var v5 bool = true
    v6 := make(map[uint32]int)
    var v7 bool = v3
    for fi8 := 0; fi8 < 7; fi8++ {
        v2 = (v2 ^ (v2 ^ (v2 - (v2 << ((326) & 63)))))
        _ = fi8
    }
    var v9 int8 = 16
    if true {
        v10 := GenS1{F0: -170}
        v11 := GenS1{F0: 938}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 int8 = (^(v9))
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 uint32 = 42
    v3 = ((false || (!("a" <= "\n\t"))) && v3)
    var v16 int8 = v9
    v17 := []int{42, 555}
    for fi18 := 0; fi18 < 18; fi18++ {
        if false {
            fmt.Println(v9, v16, v9)
        }
        fmt.Println(fi18, v15, v15)
        _ = fi18
    }
    if (921 < (v2 + (v2 / (((v2 / ((v2) | 1))) | 1)))) {
        fmt.Println(v9, v5)
    } else {
        switch "hello" {
        case "":
            v2 = 255
        case "a":
            v19 := GenS1{F0: 42}
            _ = v19
        }
    }
    if ((1.5 > (1.5 * 0.1)) || false) {
        var v20 uint32 = (v15 + (v15 >> ((42) & 63)))
        v1 = -2147483648
        _ = v20
    }
    v9 = (-(v16))
    _ = len(v6)
    var v21 uint8 = 146
    var v22 bool = (v2 < 311)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
