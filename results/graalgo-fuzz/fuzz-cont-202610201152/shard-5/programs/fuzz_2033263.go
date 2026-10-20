package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0() uint8 {
    v5 := GenS1{F0: 0, F1: 533}
    var v6 uint = 391
    if true {
        v6 = 255
        switch -99 {
        case 128:
            v6 = v6
        case 9223372036854775807:
            v6 += (v6 ^ v6)
        case -549:
            v7 := []int{-773, -282, 576}
            _ = v7
        default:
            switch 2147483647 {
            case -2147483648:
                if true {
                    v8 := GenS1{F0: -849, F1: (v6 | (^(v6)))}
                    var v9 uint32 = 2147483648
                    v6 = (v6 >> ((v6) & 63))
                    _ = v8
                    _ = v9
                } else {
                    fmt.Println(v6)
                    v6 = (^(v6))
                }
            case -999:
                v6 += v6
            }
        }
    }
    v10 := GenS1{F0: 127, F1: v6}
    for fi11 := 0; fi11 < 3; fi11++ {
        switch (^(fi11)) {
        case 42:
            fmt.Println(v6, fi11)
        case 128:
            var v12 int16 = 1
            _ = v12
        case -845:
            v13 := []int{}
            _ = v13
        }
        _ = fi11
    }
    _ = v10
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int64 = 1
    v2 := []string{(("a" + "\n\t") + ""), "", ("" + "hello"), "foo bar"}
    v3 := []int{-665}
    var v4 int8 = 100
    var v14 float32 = (1.5 * 0.1)
    v1 = (v1 / ((v1) | 1))
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v14)
    }
    v14 += 1.5
    v14 = v14
    v4 = 1
    var v15 uint16 = 0
    switch 561 {
    case 127:
        var v16 int = 32767
        _ = v16
    default:
        var v17 int32 = -745
        _ = v17
    }
    fmt.Println(v1, v4, v15)
    v4 = (v4 & -36)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
