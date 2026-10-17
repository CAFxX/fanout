package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 232
    var v3 bool = (100 == -98)
    var v4 int64 = -1
    var v5 int8 = 1
    v6 := []int{127}
    var v7 string = "foo bar"
    v6 = append(v6, 2147483647)
    if (v3 || ((-(v4)) > 0)) {
        _ = len(v6)
        switch 42 {
        case 32767:
            switch 42 {
            case 2147483647:
                if false {
                    for fi8 := 0; fi8 < 7; fi8++ {
                        var v9 string = "foo bar"
                        _ = fi8
                        _ = v9
                    }
                }
            }
        case 503:
            if len(v6) > 0 {
                v6[0] = 128
            }
        }
        v4 = v4
    }
    fmt.Println(v1, v4, v7)
    v10 := GenS1{F0: -120, F1: (42 != 0)}
    if v3 {
        var v11 bool = v3
        v12 := GenS1{F0: 49, F1: (7 <= 32767)}
        var v13 uint = 381
        _ = v11
        _ = v12
        _ = v13
    } else {
        v2 = 128
    }
    _ = len(v6)
    v2 += (v2 - 255)
    var v14 uint16 = 2
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v15 := GenS1{F0: 1, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
