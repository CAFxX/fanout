package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = 95
    var v2 bool = ((!true) && true)
    v3 := make(map[string]int)
    var v4 uint64 = 1
    if true {
        var v5 int64 = 0
        if v2 {
            v6 := []int{}
            _ = v6
        }
        fmt.Println(v1)
        _ = v5
    }
    if true {
        v3["foo bar"] = 0
        if true {
            fmt.Println(v1)
        }
        var v7 uint16 = 32767
        _ = v7
    } else {
        var v8 int64 = -518
        for fi9 := 0; fi9 < 16; fi9++ {
            if true {
                v1 -= (v1 * v1)
                var v10 int64 = 9223372036854775807
                switch 564 {
                case 462:
                    v1 = v1
                }
                _ = v10
            }
            _ = fi9
        }
        _ = v8
    }
    v11 := []int{279}
    var v12 int16 = (v1 & (v1 << ((88) & 63)))
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if ((4294967295 >= 65535) && (1 == 100)) {
        var v13 int = 410
        var v14 uint = 255
        v15 := GenS1{F0: 255}
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 int16 = v1
        if v2 {
            var v17 uint16 = 2
            _ = v17
        }
        v18 := GenS2{F0: 127}
        _ = v16
        _ = v18
    }
    switch 127 {
    case 0:
        v12 = 32767
    case -9223372036854775808:
        v4 = 303
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
