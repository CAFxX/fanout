package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = 2147483647
    var v3 int = 1
    var v4 uint64 = v1
    v4 += (^(v4))
    for fi5 := 0; fi5 < 11; fi5++ {
        switch (fi5 % ((fi5) | 1)) {
        case -593:
            v6 := []int{}
            _ = v6
        case 324:
            v7 := GenS1{F0: -1}
            _ = v7
        case 0:
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int8 = 110
                _ = fi8
                _ = v9
            }
        default:
            if true {
                var v10 int16 = 1000
                if false {
                    var v11 string = ""
                    _ = v11
                } else {
                    var v12 string = ("foo bar" + "foo bar")
                    var v13 int64 = 42
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        }
        v4 = (v4 - v1)
        _ = fi5
    }
    v14 := GenS1{F0: 100}
    v15 := []int{(v3 + (v3 | v3)), v3}
    v16 := GenS2{F0: 42}
    var v17 bool = true
    for fi18 := 0; fi18 < 19; fi18++ {
        v15 = append(v15, 2147483647)
        var v19 uint64 = v1
        v1 = v1
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
