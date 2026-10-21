package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int; F2 int16 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 68
    v3 := make(map[uint32]int)
    var v4 uint8 = (^(v2))
    if (!false) {
        fmt.Println(v2, v1, v4)
    }
    if ((0.0 + 0.0) < 0.0) {
        if false {
            for fi5 := 0; fi5 < 9; fi5++ {
                delete(v3, 80)
                fmt.Println(v1, v4)
                _ = fi5
            }
            fmt.Println(v1)
            var v6 uint64 = 2
            _ = v6
        } else {
            v7 := GenS1{F0: 0, F1: 536}
            _ = v7
        }
        fmt.Println(v4, v4, v4)
    }
    fmt.Println(v2, v1)
    v2 += (v2 * (v2 % (((v4 << ((0) & 63))) | 1)))
    var v8 uint16 = 2
    v2 = (v2 * v2)
    v9 := GenS2{F0: 4294967295}
    v10 := GenS2{F0: 9223372036854775807}
    for fi11 := 0; fi11 < 17; fi11++ {
        switch v1 {
        case "hello":
            switch ("a" + "a") {
            case "a":
                var v12 uint64 = 612
                _ = v12
            case "hello":
                for fi13 := 0; fi13 < 6; fi13++ {
                    switch v1 {
                    case "foo bar":
                        if true {
                            v14 := v3[0]
                            v15 := GenS1{F0: 1, F1: 42}
                            _ = v14
                            _ = v15
                        } else {
                            v1 = v1
                            switch ("a" + v1) {
                            case "foo bar":
                                v1 = v1
                            case "":
                                v1 = (v1 + v1)
                            default:
                                v1 = "hello"
                            }
                        }
                    default:
                        var v16 string = v1
                        _ = v16
                    }
                    _ = fi13
                }
            }
        }
        v17 := GenS3{F0: 712, F1: (fi11 >> ((484) & 63)), F2: -32768}
        _ = fi11
        _ = v17
    }
    if false {
        var v18 float64 = -2.25
        var v19 uint64 = 4294967295
        if true {
            v20 := GenS4{F0: 9223372036854775807}
            v21 := []int{}
            _ = v20
            _ = v21
        }
        _ = v18
        _ = v19
    } else {
        var v22 int8 = 0
        var v23 uint16 = v8
        _ = v22
        _ = v23
    }
    if (1 < 796) {
        var v24 uint = 643
        var v25 uint64 = 0
        _ = v24
        _ = v25
    } else {
        var v26 int8 = 122
        for fi27 := 0; fi27 < 9; fi27++ {
            fmt.Println(v4, v26)
            _ = fi27
        }
        _ = v26
    }
    v4 += (v2 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
