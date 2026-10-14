package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int }
func genFn0(p0 string) string {
    switch p0 {
    case "foo bar":
        var v4 uint16 = 42
        _ = v4
    default:
        if (561 >= 559) {
            for fi5 := 0; fi5 < 11; fi5++ {
                if (730 < 1) {
                    fmt.Println(fi5, p0)
                    switch fi5 {
                    case 259:
                        var v6 int16 = -178
                        _ = v6
                    case 2147483647:
                        var v7 int16 = -32768
                        _ = v7
                    }
                } else {
                    var v8 int64 = 793
                    v9 := GenS1{F0: 1, F1: 0, F2: 4294967295}
                    v8 = (^(v8))
                    _ = v8
                    _ = v9
                }
                _ = fi5
            }
        } else {
            var v10 uint64 = 2
            v10 += v10
            switch ("a" + "hello") {
            case "hello":
                var v11 int = 2147483647
                _ = v11
            }
            _ = v10
        }
    }
    v12 := []int{-1}
    _ = p0
    _ = v12
    return (("hello" + p0) + (p0 + p0))
}

func main() {
    var v1 uint8 = 65
    var v2 bool = true
    var v3 string = "a"
    v1 += v1
    v2 = (32767 == 1)
    var v13 int16 = 827
    var v14 float64 = float64((1e10 + 1e10))
    var v15 uint32 = 2147483648
    fmt.Println(v3, v14, v1)
    fmt.Println(v14, v15)
    v16 := GenS1{F0: 1, F1: 42, F2: -2147483648}
    v3 = "\n\t"
    v2 = v2
    v17 := GenS1{F0: 2, F1: 1, F2: 9223372036854775807}
    v18 := GenS1{F0: uint16(v15), F1: -1, F2: 127}
    if v2 {
        fmt.Println(v3, v14, v15)
    }
    v19 := []int{42, -974, 395, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
