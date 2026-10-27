package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 3.14159
    var v3 float64 = (-2.25 / (-2.25 * 3.14159))
    var v4 uint16 = 42
    var v5 int32 = 0
    v6 := make(map[string]int)
    switch "foo bar" {
    case "a":
        fmt.Println(v4)
    case "\n\t":
        v5 += v5
    }
    var v7 int = 252
    if (((((("hello" + "a") + "a") + "foo bar") >= "a") && true) && (!(false && (!false)))) {
        v2 = 0.1
        delete(v1, (v7 ^ (v7 >> ((49) & 63))))
        v8 := []int{v7, v7, (v7 / ((v7) | 1))}
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if (!(1e10 < -2.25)) {
            v10 := GenS1{F0: 683}
            v4 += (v4 % (((v4 << ((863) & 63))) | 1))
            _ = v10
        }
        v11 := GenS1{F0: (v4 / ((v4) | 1))}
        _ = fi9
        _ = v11
    }
    fmt.Println(v3)
    v4 -= (^(v4))
    v12 := []int{}
    var v13 float64 = v3
    var v14 uint64 = 42
    if (4294967295 >= 154) {
        var v15 int32 = (v5 / (((v5 * -2147483648)) | 1))
        _ = v15
    }
    var v16 uint16 = v4
    var v17 int8 = 95
    if false {
        var v18 int = (^(v7))
        _ = v18
    }
    if (v5 != int32(v4)) {
        delete(v6, "foo bar")
        if true {
            if (float32(v13) == 0.1) {
                var v19 int8 = v17
                var v20 uint = 65535
                v12 = append(v12, v7)
                _ = v19
                _ = v20
            }
            v21 := v1[v7]
            var v22 uint64 = v14
            _ = v21
            _ = v22
        }
    }
    var v23 uint16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
