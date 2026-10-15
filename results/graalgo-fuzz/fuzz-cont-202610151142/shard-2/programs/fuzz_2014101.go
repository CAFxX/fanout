package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }

func main() {
    var v1 string = "a"
    var v2 uint64 = 18446744073709551615
    var v3 int = 4294967295
    var v4 uint32 = 0
    var v5 uint = 152
    v6 := []int{992, 1, (v3 * -695), v3}
    if false {
        fmt.Println(v1)
        var v7 bool = (42 < 42)
        _ = v7
    }
    var v8 string = v1
    v2 = v2
    v3 = (v3 * 127)
    v9 := GenS1{F0: 100, F1: 25}
    fmt.Println(v2)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint8 = 118
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(v8)
            if (true || (!(2 >= 255))) {
                v13 := GenS1{F0: -128, F1: 32767}
                v2 = v2
                _ = v13
            } else {
                if len(v6) > 0 {
                    v6[0] = (v3 + (v3 & (fi12 ^ fi10)))
                }
                v11 = uint8(v4)
                v14 := GenS2{F0: v4, F1: 255, F2: v5}
                _ = v14
            }
            _ = fi12
        }
        var v15 int = 32768
        _ = fi10
        _ = v11
        _ = v15
    }
    if true {
        switch v3 {
        case 32767:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            for fi16 := 0; fi16 < 12; fi16++ {
                fmt.Println(v4, fi16)
                var v17 uint8 = 3
                v8 = ""
                _ = fi16
                _ = v17
            }
        }
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
