package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 1
    var v2 int8 = -40
    v3 := []int{-703, 32767, 32768}
    var v4 int8 = (v2 >> ((1) & 63))
    v5 := []uint32{}
    fmt.Println(v1, v1, v4)
    var v6 int16 = 42
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 -= v2
    var v7 uint64 = 18446744073709551615
    if true {
        v1 += v1
    } else {
        if false {
            v1 -= (v1 << ((4294967295) & 63))
        }
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int8 = v2
            _ = fi8
            _ = v9
        }
    }
    v4 = 0
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            switch "\n\t" {
            case "foo bar":
                v4 += v2
            default:
                fmt.Println(v1, v6, v6)
            }
            var v12 int64 = 9223372036854775807
            if len(v3) > 0 {
                v3[0] = fi11
            }
            _ = fi11
            _ = v12
        }
        var v13 uint64 = v7
        _ = len(v3)
        _ = fi10
        _ = v13
    }
    var v14 bool = true
    switch "hello" {
    case "":
        _ = len(v3)
    case "\n\t":
        fmt.Println(v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
