package main

import (
    "fmt"
)

func genFn0(p0 int8, p1 int32, p2 uint) int16 {
    var v4 uint32 = 2147483647
    var v5 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -32768
}

func main() {
    var v1 float64 = float64(((1.5 + 0.0) / (1e10 - 0.0)))
    var v2 uint = 42
    var v3 float32 = -2.25
    fmt.Println(v2, v1, v3)
    v6 := []int{-1}
    var v7 bool = true
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int16 = genFn0(2, -487, v2)
        fmt.Println(v7, v7, v9)
        var v10 int64 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = v2
    if false {
        if len(v6) > 0 {
            v6[0] = 127
        }
    } else {
        var v11 int32 = -953
        fmt.Println(v11, v11, v3)
        _ = v11
    }
    if (640 != -2147483648) {
        v7 = (485 != 0)
        if (!(2 >= 138)) {
            switch ("foo bar" + "foo bar") {
            case "hello":
                var v12 int = -9223372036854775808
                _ = v12
            case "foo bar":
                v3 = 0.0
            case "a":
                var v13 float32 = v3
                _ = v13
            default:
                var v14 int16 = 381
                _ = v14
            }
            var v15 string = ""
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 bool = v7
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int8 = -109
            _ = fi18
            _ = v19
        }
    }
    var v20 int8 = 127
    _ = len(v6)
    if len(v6) > 0 {
        v6[0] = 0
    }
    var v21 string = "hello"
    v20 = (-(v20))
    fmt.Println(v2)
    var v22 uint = (v2 / (((v2 % ((v2) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v6
    _ = v7
}
