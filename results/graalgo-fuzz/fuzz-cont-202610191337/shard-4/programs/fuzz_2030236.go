package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0() uint64 {
    if (465 < 32767) {
        var v6 int32 = 1
        fmt.Println(v6)
        _ = v6
    }
    if false {
        var v7 int64 = 988
        if false {
            v7 = -240
        } else {
            if true {
                fmt.Println(v7)
                v8 := GenS1{F0: 114}
                v7 -= 1
                _ = v8
            }
            var v9 float64 = (((((((1.5 + 3.14159) * 1.5) * 1e10) - 1e10) / (1e10 * 0.1)) * float64((3.14159 - 0.1))) / 1.0)
            _ = v9
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(v7, fi10)
            var v11 int32 = -1
            _ = fi10
            _ = v11
        }
        _ = v7
    }
    return 0
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = 630
    v3 := make(map[int]int)
    var v4 int16 = 0
    var v5 uint16 = 42
    switch "a" {
    case "\n\t":
        v2 += 2
    case "a":
        if (false || (true || (true && false))) {
            if true {
                var v12 int64 = -661
                _ = v12
            }
            var v13 uint32 = 0
            fmt.Println(v5, v4, v2)
            _ = v13
        }
    }
    var v14 uint64 = 18446744073709551615
    if (-2147483648 == 4294967295) {
        var v15 int32 = 863
        if false {
            if true {
                fmt.Println(v4, v14)
            }
            var v16 uint32 = 621
            _ = v16
        }
        _ = v15
    } else {
        var v17 string = ("" + "hello")
        v18 := GenS2{F0: 550, F1: 4294967295}
        _ = len(v3)
        _ = v17
        _ = v18
    }
    var v19 float32 = ((1.5 / (-2.25 * -2.25)) / -2.25)
    fmt.Println(v4, v2)
    fmt.Println(v19)
    var v20 bool = false
    v19 -= (float32(0.0) / v19)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
