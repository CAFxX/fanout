package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0() int64 {
    var v4 uint32 = 72
    fmt.Println(v4)
    switch "foo bar" {
    case "hello":
        v5 := GenS1{F0: 1, F1: 0, F2: (^(v4))}
        _ = v5
    case "a":
        v6 := []int{2147483647, 2147483647}
        _ = v6
    case "\n\t":
        fmt.Println(v4)
    default:
        var v7 int8 = 0
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    _ = v4
    _ = v8
    return 0
}

func main() {
    var v1 uint = 42
    v2 := []string{"foo bar"}
    var v3 float32 = -2.25
    switch 50 {
    case 32767:
        for fi9 := 0; fi9 < 18; fi9++ {
            if (!false) {
                var v10 bool = true
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 int = (fi9 / ((-483) | 1))
                    _ = fi11
                    _ = v12
                }
                v3 -= ((float32((-2.25 / 1.5)) + ((0.0 * ((((v3 - v3) * ((-2.25 * v3) / v3)) - v3) + 0.0)) * (v3 / v3))) / v3)
                _ = v10
            } else {
                v3 += v3
            }
            fmt.Println(v1, fi9, v1)
            _ = fi9
        }
    case 32768:
        var v13 uint16 = 902
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 2; fi14++ {
            v3 -= 0.1
            fmt.Println(fi14, v1, v3)
            v2 = append(v2, ("a" + "a"))
            _ = fi14
        }
        fmt.Println(v3, v1)
    } else {
        for fi15 := 0; fi15 < 2; fi15++ {
            if (991 == -1) {
                v16 := GenS1{F0: 42, F1: 765, F2: 1}
                _ = len(v2)
                v17 := GenS1{F0: 1, F1: 0, F2: 642}
                _ = v16
                _ = v17
            } else {
                fmt.Println(v3)
                var v18 string = "hello"
                v19 := GenS2{F0: -128, F1: 354}
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        var v20 int8 = 106
        _ = v20
    }
    switch -516 {
    case 32768:
        fmt.Println(v1, v3)
    case -1:
        var v21 string = ""
        _ = v21
    case 42:
        v2 = append(v2, "hello")
    }
    v2 = append(v2, ("\n\t" + "a"))
    switch "\n\t" {
    case "a":
        var v22 uint64 = 336
        _ = v22
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v1)
    v23 := GenS1{F0: -2147483648, F1: genFn0(), F2: 2147483648}
    var v24 string = "hello"
    var v25 uint = 42
    if ((555 < 1) || false) {
        var v26 int16 = 0
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS1{F0: -2147483648, F1: 407, F2: 0}
            var v29 int16 = v26
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
