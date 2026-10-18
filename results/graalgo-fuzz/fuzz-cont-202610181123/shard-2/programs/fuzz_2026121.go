package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int32; F2 uint8 }

func main() {
    v1 := []int{32768}
    v2 := []int{42, -9223372036854775808, -550, 32768}
    v3 := make(map[uint32]bool)
    var v4 uint = 4294967295
    v5 := make(map[string]int)
    if (-98 >= 0) {
        var v6 int = 2147483647
        if true {
            delete(v5, "a")
            var v7 int16 = -32768
            _ = v7
        }
        var v8 int = (v6 + (v6 >> ((916) & 63)))
        _ = v6
        _ = v8
    } else {
        switch 42 {
        case 42:
            v4 -= (^(v4))
        case -415:
            v9 := GenS1{F0: 2}
            _ = v9
        default:
            v10 := GenS1{F0: 589}
            _ = v10
        }
        switch "\n\t" {
        case "a":
            var v11 int16 = 42
            _ = v11
        }
        v4 = (v4 + 255)
    }
    if true {
        v12 := GenS1{F0: 2}
        var v13 string = ("hello" + ("foo bar" + "\n\t"))
        var v14 int = 128
        _ = v12
        _ = v13
        _ = v14
    } else {
        if false {
            for fi15 := 0; fi15 < 5; fi15++ {
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 uint16 = 2
                    _ = fi16
                    _ = v17
                }
                fmt.Println(fi15, v4)
                v4 = 2
                _ = fi15
            }
        }
    }
    var v18 int8 = -128
    v18 = (v18 & v18)
    v18 += (v18 / ((v18) | 1))
    var v19 float64 = 1.5
    v20 := GenS2{F0: 9223372036854775807, F1: 472, F2: 1}
    fmt.Println(v18, v19, v4)
    var v21 int8 = (v18 | v18)
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
