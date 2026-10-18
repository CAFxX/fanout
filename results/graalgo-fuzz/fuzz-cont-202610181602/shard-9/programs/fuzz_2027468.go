package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
func genFn0(p0 uint64, p1 int16) uint16 {
    var v4 int64 = 231
    v5 := []int{128, 32768, int(p1), -1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int8, p1 int64) uint32 {
    v6 := []int{-15, -1}
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(p0, p1, fi7)
        _ = fi7
    }
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 65535
}

func main() {
    var v1 uint8 = 128
    var v2 int64 = 414
    var v3 uint16 = 32767
    if ((false || true) || false) {
        switch ("hello" + "\n\t") {
        case "":
            var v9 int16 = 0
            _ = v9
        }
        var v10 uint = 378
        _ = v10
    }
    switch "" {
    case "":
        fmt.Println(v1, v1)
    case "\n\t":
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v1)
            _ = fi11
        }
    default:
        v1 = 72
    }
    fmt.Println(v2, v2, v3)
    v2 += v2
    var v12 int16 = 702
    var v13 int64 = v2
    var v14 string = "foo bar"
    var v15 string = v14
    if false {
        var v16 int32 = 42
        _ = v16
    } else {
        var v17 int64 = v13
        fmt.Println(v3)
        _ = v17
    }
    v18 := GenS1{F0: -1, F1: false, F2: 1}
    var v19 int = -2147483648
    fmt.Println(v3, v13)
    var v20 int16 = (v12 >> ((653) & 63))
    if true {
        v21 := []int{-9223372036854775808, (v19 & (v19 ^ 0))}
        _ = len(v21)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
