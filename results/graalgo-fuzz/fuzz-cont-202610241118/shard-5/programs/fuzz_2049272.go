package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0() uint {
    if false {
        for fi5 := 0; fi5 < 20; fi5++ {
            if true {
                var v6 string = "\n\t"
                _ = v6
            } else {
                fmt.Println(fi5)
            }
            _ = fi5
        }
        switch -759 {
        case 32767:
            var v7 float64 = (((1e10 / (1e10 / 1.0)) * 1.5) * ((float64(1.5) / 1e10) - 0.1))
            _ = v7
        }
    }
    var v8 uint = 4294967295
    fmt.Println(v8)
    _ = v8
    return (v8 - v8)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 uint64 = 1
    var v4 uint8 = 128
    switch "" {
    case "foo bar":
        v4 = (v4 + v4)
    case "\n\t":
        v9 := v1[-9223372036854775808]
        _ = v9
    case "":
        if true {
            v4 += (v4 / ((v4) | 1))
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(v4, v3)
        _ = fi10
    }
    v11 := GenS1{F0: -78, F1: 1}
    var v12 int64 = -9223372036854775808
    v13 := []int{-1, -879, 4294967295}
    v4 -= 128
    v4 = v4
    var v14 int16 = -32768
    var v15 string = "hello"
    delete(v1, int(v3))
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint8 = 0
        _ = fi16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
