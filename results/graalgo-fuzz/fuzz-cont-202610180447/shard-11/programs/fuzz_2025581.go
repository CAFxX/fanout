package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }

func main() {
    var v1 string = ("foo bar" + ("\n\t" + "\n\t"))
    var v2 uint8 = 2
    var v3 float64 = 0.0
    var v4 uint = 42
    var v5 uint32 = 2
    var v6 uint = 699
    if false {
        if (!(false || true)) {
            v4 = 4294967295
            v7 := []int{-2147483648, 32768, -2147483648}
            _ = v7
        } else {
            var v8 int8 = 42
            _ = v8
        }
        v9 := GenS1{F0: 42, F1: (v5 << (((v6 * (v6 & v6))) & 63))}
        _ = v9
    } else {
        v10 := []int{-2147483648, 4294967295, 860, 9223372036854775807}
        v11 := GenS1{F0: 42, F1: v5}
        var v12 int64 = int64(v2)
        _ = v10
        _ = v11
        _ = v12
    }
    switch 263 {
    case 128:
        var v13 int64 = 9223372036854775807
        _ = v13
    case 1:
        v14 := GenS1{F0: 936, F1: 454}
        _ = v14
    case 862:
        v6 = v6
    }
    v1 = v1
    var v15 uint64 = 2
    fmt.Println(v4)
    var v16 uint64 = 0
    var v17 string = v1
    v18 := []int{1, 2147483647, 0}
    if len(v18) > 0 {
        v18[0] = 42
    }
    switch "" {
    case "hello":
        var v19 int8 = -23
        _ = v19
    case "a":
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    default:
        var v20 int64 = -1
        _ = v20
    }
    if true {
        v21 := GenS1{F0: 0, F1: (v5 % ((281) | 1))}
        _ = len(v18)
        _ = v21
    } else {
        if false {
            fmt.Println(v5, v16)
            fmt.Println(v17)
        } else {
            var v22 string = "hello"
            var v23 int32 = -2147483648
            _ = v22
            _ = v23
        }
        if len(v18) > 0 {
            v18[0] = 523
        }
    }
    var v24 uint64 = uint64(v6)
    v25 := GenS1{F0: 774, F1: (v5 << ((v6) & 63))}
    _ = len(v18)
    for fi26 := 0; fi26 < 14; fi26++ {
        if false {
            var v27 uint8 = v2
            _ = v27
        }
        switch (((v17 + "foo bar") + (v17 + v17)) + "\n\t") {
        case "":
            if false {
                v4 += (v6 << ((v4) & 63))
                v28 := GenS1{F0: 570, F1: v5}
                _ = v28
            } else {
                if true {
                    v29 := GenS2{F0: (1e10 == 0.1), F1: uint16(v5)}
                    var v30 int16 = -932
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS1{F0: -9223372036854775808, F1: (v5 - (v5 - 4294967295))}
                    _ = v31
                }
            }
        case "hello":
            fmt.Println(v24)
        }
        var v32 string = (v17 + "foo bar")
        _ = fi26
        _ = v32
    }
    var v33 int16 = 42
    v4 = (v6 * v6)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
