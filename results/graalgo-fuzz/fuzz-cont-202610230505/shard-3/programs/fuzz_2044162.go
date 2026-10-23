package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0(p0 int8) uint32 {
    if true {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 int64 = -735
            _ = fi5
            _ = v6
        }
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                fmt.Println(p0, fi7)
                switch (fi7 + (fi7 / ((0) | 1))) {
                case 2147483647:
                    var v8 uint64 = 843
                    _ = v8
                case -429:
                    var v9 bool = true
                    _ = v9
                }
                if false {
                    p0 = (^(p0))
                    var v10 uint16 = 65535
                    _ = v10
                }
                _ = fi7
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 float32 = (0.0 + (0.1 * 3.14159))
    var v2 string = ("hello" + "\n\t")
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    fmt.Println(v2, v2)
    switch -878 {
    case -1:
        fmt.Println(v2)
    default:
        v11 := v4[""]
        _ = v11
    }
    v1 += 1e10
    delete(v3, 326)
    var v12 bool = false
    var v13 int64 = -9223372036854775808
    var v14 int32 = -1
    fmt.Println(v13)
    switch "\n\t" {
    case "\n\t":
        for fi15 := 0; fi15 < 5; fi15++ {
            v12 = true
            _ = fi15
        }
    default:
        if ((v14 | (v14 << ((0) & 63))) < 95) {
            v16 := GenS1{F0: (v2 + v2), F1: 1}
            var v17 string = v2
            _ = v16
            _ = v17
        }
    }
    switch -2147483648 {
    case 722:
        var v18 uint = 42
        _ = v18
    case 4294967295:
        var v19 uint32 = 65535
        _ = v19
    }
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
