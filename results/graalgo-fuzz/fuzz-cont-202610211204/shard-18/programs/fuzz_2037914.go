package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0() int8 {
    var v5 int16 = 1000
    v6 := []int{128, -1}
    v5 = 42
    if true {
        var v7 bool = false
        var v8 uint16 = 0
        var v9 int = -1
        _ = v7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return int8(v5)
}

func main() {
    var v1 uint64 = 265
    v2 := make(map[string]bool)
    v3 := []string{"\n\t"}
    var v4 int = 9223372036854775807
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v1)
    var v10 int8 = 101
    var v11 int32 = 0
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    if false {
        var v12 bool = (!true)
        if (!false) {
            v12 = (!v12)
            delete(v2, "")
        }
        var v13 uint8 = 239
        _ = v12
        _ = v13
    } else {
        switch "" {
        case "\n\t":
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 19, F1: v10}
                var v16 uint8 = 2
                _ = fi14
                _ = v15
                _ = v16
            }
        case "":
            fmt.Println(v10, v1)
        default:
            var v17 uint8 = 127
            _ = v17
        }
        switch ("" + "hello") {
        case "foo bar":
            fmt.Println(v10)
        }
    }
    if false {
        var v18 uint8 = 1
        for fi19 := 0; fi19 < 11; fi19++ {
            if (!(18446744073709551615 > 612)) {
                v20 := GenS1{F0: (v18 * (v18 / ((0) | 1))), F1: v10}
                var v21 uint16 = 236
                v22 := GenS2{F0: v1}
                _ = v20
                _ = v21
                _ = v22
            } else {
                var v23 uint = 891
                _ = v23
            }
            v11 = (^(v11))
            v10 -= (^(v10))
            _ = fi19
        }
        _ = v18
    } else {
        var v24 int64 = -9223372036854775808
        _ = v24
    }
    fmt.Println(v1, v11, v10)
    var v25 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
