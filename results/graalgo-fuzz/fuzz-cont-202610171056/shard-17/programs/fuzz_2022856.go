package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
func genFn0() uint {
    var v5 int8 = -1
    v5 = (v5 - 1)
    _ = v5
    return 65535
}

func main() {
    var v1 int8 = 0
    var v2 int32 = -1
    var v3 int32 = v2
    var v4 int8 = -122
    if false {
        v6 := []int{9223372036854775807, 410, -233}
        _ = v6
    }
    if (!(128 >= -920)) {
        v2 = 589
    }
    v1 += v4
    if false {
        var v7 uint16 = 2
        if (true && (!true)) {
            v8 := []int{128, 128}
            var v9 float64 = 1.5
            _ = v8
            _ = v9
        } else {
            var v10 float32 = 1.5
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := []int{(fi11 | fi11)}
            var v13 int64 = -652
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v7
    } else {
        switch 882 {
        case 15:
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint = 46
                var v16 string = "foo bar"
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    }
    v3 = (v2 + v2)
    v3 -= (v3 % ((377) | 1))
    v17 := GenS1{F0: 996, F1: 0}
    v1 += v1
    var v18 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
