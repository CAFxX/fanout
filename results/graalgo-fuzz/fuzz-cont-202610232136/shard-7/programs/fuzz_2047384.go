package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string }
func genFn0() int16 {
    v5 := []int{668, 1}
    var v6 float64 = (0.0 * 1e10)
    _ = v5
    _ = v6
    return 155
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]bool)
    var v3 int16 = -403
    var v4 int = -738
    var v7 string = "hello"
    if (0 > 258) {
        fmt.Println(v4)
        var v8 float64 = -2.25
        _ = v8
    }
    var v9 int = (v4 & (v4 / (((v4 ^ (v4 & 1))) | 1)))
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 bool = ((2147483647 > (fi10 / ((2147483647) | 1))) && false)
        _ = fi10
        _ = v11
    }
    switch (v7 + "") {
    case "foo bar":
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint = 0
            _ = fi12
            _ = v13
        }
    case "\n\t":
        var v14 int64 = 42
        _ = v14
    }
    fmt.Println(v7)
    v15 := GenS1{F0: 10}
    var v16 int32 = -650
    var v17 int16 = 42
    fmt.Println(v7, v16)
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint32 = v1
        _ = fi18
        _ = v19
    }
    v9 = (v9 % ((42) | 1))
    v20 := GenS2{F0: v3, F1: ((("a" + v7) + (("" + (v7 + v7)) + v7)) + v7)}
    v1 -= 2147483648
    fmt.Println(v3, v17, v1)
    if true {
        var v21 int = v4
        _ = v21
    }
    switch (v9 * -379) {
    case 0:
        v16 = (^(v16))
    case 42:
        var v22 float64 = -2.25
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 float32 = (1e10 - 0.0)
        v7 = "foo bar"
        fmt.Println(v16)
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
