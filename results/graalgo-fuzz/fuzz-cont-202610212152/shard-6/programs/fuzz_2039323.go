package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int; F2 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    var v3 int = 0
    var v4 uint16 = 36
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: -1}
    var v7 uint16 = v4
    if (-1 > 1) {
        var v8 bool = (false || (false && true))
        _ = v8
    } else {
        v4 = 978
    }
    fmt.Println(v4, v4)
    switch "" {
    case "foo bar":
        _ = len(v2)
    case "":
        v9 := GenS1{F0: 42}
        _ = v9
    default:
        for fi10 := 0; fi10 < 4; fi10++ {
            switch "hello" {
            case "hello":
                for fi11 := 0; fi11 < 5; fi11++ {
                    v3 -= -211
                    if len(v2) > 0 {
                        v2[0] = v3
                    }
                    if false {
                        var v12 int32 = -2147483648
                        _ = v12
                    } else {
                        fmt.Println(v7, fi10, v4)
                        v3 -= (fi11 >> ((341) & 63))
                    }
                    _ = fi11
                }
            case "":
                for fi13 := 0; fi13 < 11; fi13++ {
                    fmt.Println(v4, fi10, v7)
                    _ = fi13
                }
            }
            var v14 uint = 65535
            _ = fi10
            _ = v14
        }
    }
    var v15 int32 = -1
    _ = len(v1)
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS1{F0: 987}
            var v19 int64 = -495
            _ = len(v5)
            v7 -= v4
            if false {
                var v20 float64 = 0.1
                var v21 uint32 = 42
                _ = v20
                _ = v21
            }
            v3 = 9223372036854775807
            _ = fi17
            _ = v18
            _ = v19
        }
        if false {
            var v22 string = ""
            var v23 int8 = 1
            var v24 float64 = 1e10
            _ = v22
            _ = v23
            _ = v24
        } else {
            fmt.Println(v4)
        }
        _ = fi16
    }
    var v25 int = (v3 << ((2) & 63))
    v15 = v15
    var v26 float32 = 1e10
    v2 = append(v2, v25)
    if (true && false) {
        v27 := GenS2{F0: 9223372036854775807, F1: 1, F2: -9223372036854775808}
        v15 = 42
        v28 := GenS1{F0: -9223372036854775808}
        _ = v27
        _ = v28
    } else {
        var v29 float32 = ((v26 + -2.25) - 0.1)
        for fi30 := 0; fi30 < 13; fi30++ {
            v26 = (v29 - (((0.0 - (v29 * v26)) * 1.5) - 0.0))
            _ = fi30
        }
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
