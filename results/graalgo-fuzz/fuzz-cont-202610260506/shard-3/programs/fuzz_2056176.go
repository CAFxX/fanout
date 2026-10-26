package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint32 }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{"foo bar"}
    var v2 int64 = 1
    v3 := make(map[int]uint32)
    var v4 string = "\n\t"
    var v5 float32 = 0.0
    v6 := make(map[string]uint32)
    v6[v4] = 65535
    if true {
        if false {
            fmt.Println(v4)
            var v7 float32 = (v5 / 1.5)
            if true {
                fmt.Println(v4)
            } else {
                fmt.Println(v5)
                var v8 uint = 0
                var v9 bool = (390 >= 0)
                _ = v8
                _ = v9
            }
            _ = v7
        }
    }
    switch (v4 + ("foo bar" + v4)) {
    case "hello":
        var v10 int16 = -32768
        _ = v10
    case "foo bar":
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 bool = true
            _ = fi11
            _ = v12
        }
    case "":
        v13 := GenS1{F0: 1, F1: (v2 / ((v2) | 1)), F2: 4294967295}
        _ = v13
    default:
        v14 := GenS1{F0: 42, F1: 0, F2: 2147483648}
        _ = v14
    }
    delete(v6, v4)
    for fi15 := 0; fi15 < 7; fi15++ {
        if false {
            v16 := GenS2{F0: (fi15 / ((0) | 1))}
            _ = len(v1)
            _ = v16
        }
        var v17 int16 = 42
        _ = fi15
        _ = v17
    }
    v1 = append(v1, v4)
    var v18 uint64 = 18446744073709551615
    _ = len(v6)
    switch v4 {
    case "a":
        v19 := GenS2{F0: 42}
        _ = v19
    case "":
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 float64 = 1e10
            _ = fi20
            _ = v21
        }
    case "hello":
        v18 = v18
    }
    fmt.Println(v4, v4)
    var v22 int8 = 100
    if ((-(v2)) <= (v2 % (((v2 % (((v2 % ((v2) | 1))) | 1))) | 1))) {
        var v23 int32 = 752
        if false {
            var v24 uint64 = v18
            _ = v24
        }
        _ = v23
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 int8 = v22
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 uint64 = (^(v18))
            _ = fi28
            _ = v29
        }
        var v30 int8 = (v22 & -63)
        _ = fi27
        _ = v30
    }
    var v31 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
