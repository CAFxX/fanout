package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    var v4 uint32 = 566
    var v5 uint16 = 2
    var v6 float32 = 0.1
    var v7 float64 = 1e10
    v8 := []int{4294967295, 32768, -1, 2147483647}
    v6 += v6
    switch "hello" {
    case "\n\t":
        if false {
            switch "a" {
            case "":
                if (22 != 128) {
                    var v9 uint8 = 255
                    v5 = (v5 + (v5 / ((v5) | 1)))
                    _ = v9
                }
            case "hello":
                fmt.Println(v6, v6, v5)
            case "foo bar":
                v5 = (^(v5))
            }
            v4 = v2
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
    default:
        v10 := GenS1{F0: 4294967295, F1: -22}
        _ = v10
    }
    v7 = (0.1 - float64((v6 * float32(float64((v6 - ((0.1 - v6) / v6)))))))
    if (!(0 < 42)) {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: 1, F1: 887}
            _ = fi11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint64 = 42
            var v15 uint = 398
            _ = fi13
            _ = v14
            _ = v15
        }
        if (v5 >= v5) {
            if true {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 uint32 = v2
                    _ = fi16
                    _ = v17
                }
            }
        }
        var v18 bool = false
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v7 += v7
        fmt.Println(v5)
        _ = fi19
    }
    if false {
        var v20 int64 = 1
        _ = v20
    } else {
        v5 = (v5 & 427)
        if len(v8) > 0 {
            v8[0] = 127
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
