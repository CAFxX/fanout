package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16 }
func genFn0() int {
    for fi4 := 0; fi4 < 16; fi4++ {
        if false {
            v5 := GenS1{F0: 96, F1: -9223372036854775808}
            v6 := GenS1{F0: 42, F1: (fi4 * (fi4 & -2147483648))}
            _ = v5
            _ = v6
        }
        switch ("a" + (("hello" + ("a" + ("" + ""))) + ("foo bar" + "a"))) {
        case "":
            var v7 uint64 = 2
            _ = v7
        case "hello":
            v8 := GenS1{F0: 42, F1: fi4}
            _ = v8
        default:
            fmt.Println(fi4)
        }
        _ = fi4
    }
    v9 := []int{-363}
    var v10 uint8 = 128
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 uint = 559
    var v2 int16 = 1000
    var v3 uint16 = 1
    fmt.Println(v2)
    var v11 int32 = -1
    v11 = (v11 | (^(v11)))
    v11 = (v11 ^ (v11 - (v11 ^ -1)))
    switch "hello" {
    case "\n\t":
        v12 := []int{127}
        _ = v12
    case "foo bar":
        switch "" {
        case "hello":
            var v13 float64 = -2.25
            _ = v13
        }
    }
    switch 32768 {
    case 0:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int16 = v2
            v15 = (^(v15))
            v16 := GenS1{F0: (^(v1)), F1: fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
    case 8:
        v11 = (v11 ^ (v11 >> (((^(v1))) & 63)))
    }
    v17 := []int{-573, -328}
    v18 := GenS2{F0: (^(v3))}
    var v19 uint32 = 2
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v2, v3)
        _ = fi20
    }
    var v21 uint16 = (v3 * (v3 << ((v1) & 63)))
    var v22 uint16 = 1
    var v23 int16 = int16(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
