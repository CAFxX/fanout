package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 int64 = 1
    v3 := []string{}
    v3 = append(v3, "")
    v1 = ("foo bar" + "")
    v4 := GenS1{F0: 1, F1: 65535}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float32 = float32((-2.25 + 0.0))
        _ = fi5
        _ = v6
    }
    fmt.Println(v2)
    v7 := GenS2{F0: 42, F1: 247, F2: -813}
    v2 += v2
    var v8 uint = 1
    if true {
        fmt.Println(v8, v8, v1)
    } else {
        fmt.Println(v2, v8, v2)
    }
    v1 = "a"
    if (true || (!false)) {
        var v9 int8 = -128
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if true {
                fmt.Println(v1, v2, v2)
                if false {
                    var v12 uint64 = 654
                    _ = v12
                } else {
                    var v13 uint16 = 42
                    var v14 bool = ((true || false) && ((false && false) && true))
                    fmt.Println(v8, fi11, fi10)
                    _ = v13
                    _ = v14
                }
                v2 += (v2 % (((v2 << ((65535) & 63))) | 1))
            } else {
                var v15 int16 = 863
                v8 = 666
                v3 = append(v3, v1)
                v16 := GenS1{F0: (fi10 + (fi11 | (fi10 / (((fi11 >> ((65535) & 63))) | 1)))), F1: 475}
                _ = v15
                _ = v16
            }
            var v17 int32 = -2147483648
            _ = fi11
            _ = v17
        }
        v1 = v1
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
