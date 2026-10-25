package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint32 }

func main() {
    var v1 string = ""
    var v2 bool = (!(116 <= 18446744073709551615))
    var v3 float64 = (1e10 * (-2.25 - 0.1))
    var v4 float32 = 1e10
    if false {
        v5 := []int{1, 93}
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int32 = 1
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    switch v1 {
    case "":
        v8 := []int{9223372036854775807}
        _ = v8
    case "foo bar":
        fmt.Println(v1, v4, v4)
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v2 = v2
        fmt.Println(v4, v3, v4)
        _ = fi9
    }
    switch -2147483648 {
    case 1:
        var v10 float64 = v3
        _ = v10
    }
    v3 = (v3 / (float64(v4) / 3.14159))
    fmt.Println(v2, v2)
    if v2 {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint16 = 255
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS1{F0: 127}
    v14 := GenS2{F0: -600, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
