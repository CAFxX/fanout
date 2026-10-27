package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 uint16 = 42
    var v2 bool = (!(!(!true)))
    var v3 bool = (!((false || (42 >= -1)) && true))
    var v4 int64 = 42
    var v5 uint = 2
    v2 = ("a" > "a")
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint8 = 111
        _ = fi6
        _ = v7
    }
    if v3 {
        switch 4294967295 {
        case 127:
            var v8 uint16 = 946
            _ = v8
        case -724:
            switch "a" {
            case "\n\t":
                var v9 uint32 = 65535
                _ = v9
            default:
                var v10 int8 = int8(v1)
                _ = v10
            }
        case -2147483648:
            v5 -= (v5 << ((v5) & 63))
        default:
            var v11 int32 = 0
            _ = v11
        }
    }
    var v12 int16 = -1
    var v13 int32 = 0
    v14 := []int{-442, 128}
    if ((v1 - (v1 | 953)) > v1) {
        var v15 uint64 = 42
        if (v3 || false) {
            switch 906 {
            case 32768:
                v14 = append(v14, 0)
            case 32767:
                var v16 int32 = int32(v12)
                _ = v16
            case -588:
                v4 = v4
            }
        } else {
            v1 += v1
        }
        var v17 uint16 = v1
        _ = v15
        _ = v17
    }
    v3 = v2
    switch 32767 {
    case -551:
        var v18 int16 = (^(v12))
        _ = v18
    case 127:
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int = fi19
            v21 := GenS1{F0: fi19, F1: 67}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS1{F0: -1, F1: 127}
    var v23 uint = v5
    var v24 uint16 = (v1 + (^(v1)))
    v2 = true
    v3 = (!(!v3))
    if v3 {
        switch "hello" {
        case "a":
            v12 += (-(v12))
        case "\n\t":
            if len(v14) > 0 {
                v14[0] = 2147483647
            }
        default:
            v4 -= (v4 << ((v23) & 63))
        }
        _ = len(v14)
        var v25 int8 = int8(v5)
        _ = v25
    } else {
        _ = len(v14)
        switch "\n\t" {
        case "\n\t":
            switch -9223372036854775808 {
            case 2147483647:
                _ = len(v14)
            case -136:
                v26 := GenS1{F0: -2147483648, F1: 1}
                _ = v26
            case -1:
                fmt.Println(v13, v2)
            default:
                switch 128 {
                case 127:
                    v12 -= v12
                case 2147483647:
                    v2 = v3
                case 408:
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                default:
                    v27 := GenS1{F0: 127, F1: 127}
                    _ = v27
                }
            }
        default:
            fmt.Println(v23)
        }
    }
    var v28 int32 = 42
    v2 = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
