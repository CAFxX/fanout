package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint64; F2 string }
func genFn0() uint {
    v6 := GenS1{F0: 0, F1: 42}
    switch -9223372036854775808 {
    case 845:
        var v7 int64 = 0
        _ = v7
    case 42:
        v8 := []int{127, 0, -137, 9223372036854775807}
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint64 = 635
            switch fi9 {
            case 155:
                var v11 int64 = -978
                _ = v11
            }
            _ = fi9
            _ = v10
        }
    } else {
        var v12 int64 = 1
        _ = v12
    }
    switch "a" {
    case "":
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS1{F0: 42, F1: 1}
            _ = fi13
            _ = v14
        }
    case "\n\t":
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := []int{(fi15 << ((18446744073709551615) & 63)), (fi15 | 791), 9223372036854775807, (-(fi15))}
            _ = fi15
            _ = v16
        }
    case "foo bar":
        v17 := GenS1{F0: -2147483648, F1: 1}
        _ = v17
    }
    v18 := GenS1{F0: 42, F1: 1}
    _ = v18
    _ = v6
    return 798
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]uint32)
    var v3 int64 = 587
    var v4 uint8 = 131
    v5 := make(map[string]uint32)
    if (!(!((32768 != 1) && (0.0 == 1e10)))) {
        var v19 int64 = v3
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS2{F0: 123, F1: 4294967295, F2: "a"}
            delete(v2, fi20)
            fmt.Println(v19)
            _ = fi20
            _ = v21
        }
        switch -1 {
        case 128:
            if true {
                fmt.Println(v19)
                var v22 int8 = -1
                _ = v22
            }
        case 9223372036854775807:
            delete(v5, "a")
        }
        _ = v19
    }
    fmt.Println(v1)
    _ = len(v2)
    delete(v5, "hello")
    v1 = v1
    var v23 uint8 = v4
    var v24 int16 = -32768
    v25 := v2[9223372036854775807]
    var v26 bool = true
    fmt.Println(v25, v24, v24)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
