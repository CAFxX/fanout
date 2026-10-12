package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    var v1 bool = false
    var v2 uint32 = 0
    var v3 int8 = -1
    var v4 uint64 = 731
    var v5 string = "\n\t"
    var v6 uint = 1
    var v7 uint8 = 71
    fmt.Println(v1, v7, v4)
    if v1 {
        switch v5 {
        case "hello":
            v2 = v2
        default:
            switch "" {
            case "a":
                fmt.Println(v4)
            case "\n\t":
                v2 += v2
            case "foo bar":
                var v8 int16 = 609
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint64 = 4294967295
            if false {
                var v11 int8 = -128
                v12 := []int{(^(fi9)), 32768, fi9}
                _ = v11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        if v1 {
            v7 -= (^(v7))
            v7 -= 128
        }
    }
    v13 := GenS1{F0: int(v6), F1: uint16(v7)}
    v4 += uint64(v2)
    v14 := GenS1{F0: 988, F1: 294}
    var v15 float64 = 3.14159
    var v16 uint = v6
    v6 -= (v6 * 18446744073709551615)
    for fi17 := 0; fi17 < 7; fi17++ {
        v16 = v6
        _ = fi17
    }
    v3 -= v3
    v5 = "\n\t"
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := []int{}
        v20 := GenS1{F0: (^(fi18)), F1: 189}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint16 = 255
    v22 := GenS1{F0: 214, F1: 387}
    switch 736 {
    case 42:
        var v23 int32 = -555
        _ = v23
    default:
        v24 := GenS1{F0: 128, F1: v21}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
