package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint16 = 654
        _ = fi4
        _ = v5
    }
    var v6 uint64 = 2
    _ = v6
    return 1000
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float32 = -2.25
    v3 := []int{-1, 32767, 150, 42}
    var v7 string = (("\n\t" + "\n\t") + "foo bar")
    v2 = v2
    if true {
        v8 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = v8
    } else {
        v3 = append(v3, 42)
    }
    if (42 > 321) {
        if false {
            fmt.Println(v2)
            v9 := GenS1{F0: 2, F1: 42}
            _ = v9
        }
        switch v7 {
        case "a":
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint64 = 2
                _ = fi10
                _ = v11
            }
        case "foo bar":
            v12 := GenS1{F0: 4294967295, F1: 66}
            _ = v12
        }
    } else {
        v1 += (v1 | v1)
        if true {
            var v13 float64 = 3.14159
            fmt.Println(v2, v2, v2)
            var v14 uint = 255
            _ = v13
            _ = v14
        } else {
            var v15 uint8 = 127
            v1 += v1
            var v16 int = 128
            _ = v15
            _ = v16
        }
        var v17 int8 = 127
        _ = v17
    }
    _ = len(v3)
    fmt.Println(v2)
    switch 32768 {
    case -632:
        v18 := GenS1{F0: 0, F1: 0}
        _ = v18
    case 127:
        v1 = (v1 | 746)
    default:
        v19 := GenS1{F0: 18446744073709551615, F1: -92}
        _ = v19
    }
    fmt.Println(v7, v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
