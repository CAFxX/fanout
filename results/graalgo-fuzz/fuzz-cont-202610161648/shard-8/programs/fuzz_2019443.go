package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 bool = (!(false && (!false)))
    var v2 uint16 = 32767
    var v3 bool = v1
    var v4 int32 = 459
    switch 170 {
    case -9223372036854775808:
        var v5 int8 = 33
        _ = v5
    case 42:
        switch 9223372036854775807 {
        case -730:
            v1 = (v1 || false)
        default:
            for fi6 := 0; fi6 < 10; fi6++ {
                if (((((((("\n\t" + "") + "") > (("a" + "") + "foo bar")) || v1) || ((795 > 42) && false)) || v1) && (!(1.5 < 3.14159))) || ((847 != 2) && false)) {
                    v3 = (-445 > -995)
                    v7 := GenS1{F0: (v2 * v2), F1: (!v1), F2: (!(true || (v3 || false)))}
                    v2 = v2
                    _ = v7
                }
                _ = fi6
            }
        }
    case -1:
        switch "a" {
        case "hello":
            v8 := []int{4294967295}
            _ = v8
        default:
            var v9 uint64 = 622
            _ = v9
        }
    }
    var v10 int16 = -1
    v11 := GenS1{F0: 1, F1: (!v1), F2: (128 != 113)}
    v1 = v1
    var v12 int = 128
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 float64 = -2.25
        _ = fi13
        _ = v14
    }
    fmt.Println(v4, v4)
    fmt.Println(v1, v3, v1)
    v3 = ("a" < "foo bar")
    if false {
        var v15 int16 = 32767
        _ = v15
    } else {
        var v16 int32 = 1
        var v17 int16 = (v10 >> ((255) & 63))
        switch "foo bar" {
        case "foo bar":
            switch 42 {
            case 9223372036854775807:
                fmt.Println(v16, v12, v2)
            default:
                var v18 uint16 = (v2 ^ (v2 ^ v2))
                _ = v18
            }
        case "":
            v17 -= v10
        default:
            fmt.Println(v3)
        }
        _ = v16
        _ = v17
    }
    var v19 int64 = 42
    switch (v12 / ((v12) | 1)) {
    case -892:
        switch "a" {
        case "":
            var v20 int32 = v4
            _ = v20
        case "\n\t":
            var v21 int8 = 14
            _ = v21
        case "foo bar":
            var v22 uint = 46
            _ = v22
        default:
            switch v12 {
            case 500:
                var v23 uint16 = 1
                _ = v23
            }
        }
    case -1:
        var v24 int32 = (v4 & (v4 << ((940) & 63)))
        _ = v24
    case 2147483647:
        var v25 uint16 = (v2 + (v2 * 32767))
        _ = v25
    }
    var v26 uint16 = 42
    fmt.Println(v4, v2, v12)
    v27 := GenS2{F0: int64(v10)}
    v3 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
