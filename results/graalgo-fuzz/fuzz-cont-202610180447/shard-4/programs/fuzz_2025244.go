package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 string }
func genFn0() uint16 {
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v6)
        switch -698 {
        case 128:
            var v8 bool = (!(true && false))
            _ = v8
        case 0:
            v9 := GenS1{F0: 2}
            _ = v9
        case 1:
            var v10 int16 = 32767
            _ = v10
        }
        var v11 uint16 = 0
        _ = fi7
        _ = v11
    }
    switch -1 {
    case 112:
        if true {
            v6 = v6
            fmt.Println(v6)
            v12 := GenS1{F0: 18446744073709551615}
            _ = v12
        } else {
            var v13 int8 = 110
            switch "hello" {
            case "\n\t":
                v14 := GenS2{F0: (v6 / ((689) | 1)), F1: 4294967295}
                _ = v14
            default:
                fmt.Println(v13)
            }
            _ = v13
        }
    }
    _ = v6
    return 2
}

func main() {
    var v1 uint8 = 127
    v2 := []string{"hello", ""}
    var v3 uint8 = v1
    v4 := make(map[int]int)
    v5 := []int{}
    if (-35 <= -128) {
        v1 -= v3
        for fi15 := 0; fi15 < 13; fi15++ {
            if (100 > -86) {
                v1 = 1
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                var v16 int32 = 0
                _ = v16
            }
            _ = fi15
        }
        if false {
            if true {
                if true {
                    v17 := GenS1{F0: uint(v1)}
                    delete(v4, -580)
                    _ = v17
                }
                var v18 int8 = -18
                _ = v18
            }
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 bool = true
        _ = fi19
        _ = v20
    }
    if true {
        v21 := GenS3{F0: 299, F1: -9223372036854775808}
        _ = len(v4)
        var v22 int = 127
        _ = v21
        _ = v22
    } else {
        var v23 uint8 = 128
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v3 = v1
        var v25 int64 = 1
        v26 := GenS4{F0: "foo bar"}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 bool = (-1 >= 9223372036854775807)
    v5 = append(v5, 680)
    if (v1 > (v1 / ((2) | 1))) {
        v3 -= (^(v3))
    }
    v28 := GenS3{F0: 65535, F1: 42}
    var v29 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
