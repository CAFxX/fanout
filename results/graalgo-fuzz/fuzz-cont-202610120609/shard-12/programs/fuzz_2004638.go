package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{32767}
    v3 := make(map[string]int)
    var v4 uint32 = 1
    fmt.Println(v4)
    v5 := GenS1{F0: 0}
    if true {
        var v6 int = 4294967295
        var v7 uint32 = v4
        var v8 uint64 = 1
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 string = ("a" + ((("hello" + "a") + "\n\t") + "\n\t"))
        _ = v9
    }
    var v10 int8 = -128
    v4 += (v4 >> ((42) & 63))
    if false {
        fmt.Println(v10, v10)
        for fi11 := 0; fi11 < 7; fi11++ {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 float64 = 1.5
                var v14 int32 = -560
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int64 = -120
            _ = fi15
            _ = v16
        }
    }
    var v17 int64 = 1
    fmt.Println(v10, v10, v10)
    if false {
        v18 := GenS1{F0: 410}
        for fi19 := 0; fi19 < 7; fi19++ {
            if (!true) {
                var v20 string = "\n\t"
                if false {
                    if true {
                        v20 = "foo bar"
                        fmt.Println(v20, v20, v10)
                        if len(v2) > 0 {
                            v2[0] = fi19
                        }
                    } else {
                        fmt.Println(v4, v17)
                    }
                    v21 := v1[v20]
                    fmt.Println(v17)
                    _ = v21
                }
                var v22 int64 = (v17 % ((445) | 1))
                _ = v20
                _ = v22
            }
            _ = fi19
        }
        _ = v18
    } else {
        var v23 bool = false
        _ = v23
    }
    if true {
        var v24 uint64 = 4294967295
        _ = v24
    }
    v25 := GenS2{F0: uint(v4), F1: -2147483648, F2: 255}
    for fi26 := 0; fi26 < 14; fi26++ {
        switch "" {
        case "foo bar":
            var v27 uint16 = 2
            _ = v27
        case "\n\t":
            fmt.Println(v4)
        case "":
            for fi28 := 0; fi28 < 9; fi28++ {
                v10 = (v10 | (v10 - v10))
                var v29 uint8 = 255
                _ = fi28
                _ = v29
            }
        }
        var v30 string = ""
        _ = fi26
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
