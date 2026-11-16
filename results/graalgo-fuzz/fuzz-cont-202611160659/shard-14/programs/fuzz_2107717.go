package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
func genFn0() int32 {
    var v5 bool = ((-2.25 * 0.1) <= 1.5)
    v6 := []int{127, 42, 77, 2147483647}
    var v7 int64 = 529
    switch "\n\t" {
    case "hello":
        v5 = v5
    case "":
        var v8 bool = v5
        _ = v8
    default:
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint16 = 0
            _ = len(v6)
            _ = fi9
            _ = v10
        }
    }
    v6 = append(v6, 32767)
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1() string {
    v11 := []int{429}
    var v12 int32 = -2147483648
    _ = v11
    _ = v12
    return ("\n\t" + ("hello" + ("" + "a")))
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 1
    var v3 uint8 = 2
    var v4 int8 = 42
    var v13 int8 = (v4 << ((42) & 63))
    var v14 uint32 = 114
    v3 += v3
    fmt.Println(v4, v4)
    var v15 uint8 = (v3 | 42)
    switch "foo bar" {
    case "":
        switch 0 {
        case 177:
            v14 = (v14 % (((v14 ^ v14)) | 1))
        default:
            fmt.Println(v14, v14)
        }
    case "\n\t":
        v15 = (v15 & (v15 - 42))
    case "a":
        var v16 int64 = -458
        _ = v16
    }
    if false {
        v2 = v2
        if false {
            v2 += -1
            var v17 uint32 = v14
            if true {
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS1{F0: int(v1), F1: v2}
                    _ = fi18
                    _ = v19
                }
                v15 += v3
            } else {
                var v20 string = "\n\t"
                fmt.Println(v15)
                _ = v20
            }
            _ = v17
        } else {
            v21 := GenS1{F0: -2147483648, F1: 1000}
            _ = v21
        }
        v22 := GenS1{F0: 42, F1: (v2 << ((42) & 63))}
        _ = v22
    }
    var v23 uint = 132
    switch ("hello" + "") {
    case "a":
        v24 := GenS1{F0: -9223372036854775808, F1: -1}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
