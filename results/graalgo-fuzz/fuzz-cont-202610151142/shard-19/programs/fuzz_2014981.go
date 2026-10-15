package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 int16 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 20; fi4++ {
        if true {
            v5 := GenS1{F0: 2, F1: -2147483648, F2: -388}
            _ = v5
        }
        _ = fi4
    }
    v6 := GenS2{F0: 2147483647, F1: "hello", F2: 1}
    if false {
        switch "foo bar" {
        case "a":
            var v7 float64 = (1.5 - 0.1)
            _ = v7
        case "":
            var v8 int16 = 0
            _ = v8
        }
    } else {
        var v9 uint16 = 543
        _ = v9
    }
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (2147483648 < 848)
    v3 := make(map[string]int)
    v2 = v2
    var v10 int = 1
    if v2 {
        switch v10 {
        case 42:
            if (-2.25 == (1.5 / (3.14159 / (0.1 - (float64(3.14159) / ((((1e10 + 0.1) + 1e10) - 1.5) + 1e10)))))) {
                if v2 {
                    v11 := GenS2{F0: 188, F1: "hello", F2: 42}
                    var v12 uint64 = 42
                    _ = v11
                    _ = v12
                }
                var v13 string = "hello"
                _ = v13
            }
        }
        fmt.Println(v10, v2)
    }
    v14 := GenS1{F0: 45, F1: 906, F2: genFn0()}
    v10 = (v10 | 768)
    var v15 bool = (!(true || (v2 && v2)))
    v2 = ((v10 ^ 1) <= 397)
    var v16 uint = 255
    if true {
        var v17 bool = v15
        if (65535 >= 117) {
            v18 := v1["foo bar"]
            var v19 int64 = genFn0()
            _ = v18
            _ = v19
        } else {
            v16 = v16
            v3[((("a" + "a") + "hello") + ("" + ""))] = -501
            fmt.Println(v16)
        }
        _ = v17
    } else {
        fmt.Println(v10)
    }
    if ((!false) && v15) {
        v15 = (!(!(v15 && (126 <= 255))))
        if v15 {
            v20 := GenS1{F0: 42, F1: 409, F2: -9223372036854775808}
            var v21 uint32 = 886
            _ = v20
            _ = v21
        }
        v22 := []int{}
        _ = v22
    }
    switch "\n\t" {
    case "":
        var v23 int8 = 64
        _ = v23
    }
    var v24 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
}
