package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 string = "\n\t"
    var v2 bool = (32767 > 9223372036854775807)
    var v3 int = 32768
    var v4 int64 = 374
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = -876
        switch 127 {
        case 1:
            v2 = (!(!v2))
        case -49:
            for fi8 := 0; fi8 < 5; fi8++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v9 uint64 = 835
                var v10 float64 = (((float64(-2.25) + -2.25) - 0.1) / 1.0)
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    v5 = append(v5, v3)
    if len(v5) > 0 {
        v5[0] = -122
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            if v2 {
                var v13 int8 = 100
                var v14 uint64 = 42
                _ = v13
                _ = v14
            } else {
                fmt.Println(fi11, v4, v1)
                v15 := GenS1{F0: v3}
                if v2 {
                    var v16 int16 = 32767
                    var v17 bool = v2
                    _ = v16
                    _ = v17
                }
                _ = v15
            }
            _ = fi12
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            _ = len(v5)
            _ = fi18
        }
        _ = fi11
    }
    switch "foo bar" {
    case "hello":
        v4 = v4
    default:
        v3 = v3
    }
    var v19 uint64 = 2
    var v20 uint64 = v19
    switch v3 {
    case 32767:
        fmt.Println(v1, v4)
    case -1:
        var v21 uint32 = 1
        _ = v21
    case 817:
        fmt.Println(v1, v1)
    default:
        if len(v5) > 0 {
            v5[0] = v3
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
