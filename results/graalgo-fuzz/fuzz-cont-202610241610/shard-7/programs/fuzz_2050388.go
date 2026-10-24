package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0() int {
    for fi4 := 0; fi4 < 18; fi4++ {
        switch 9223372036854775807 {
        case 1:
            fmt.Println(fi4)
        case 0:
            fmt.Println(fi4)
        }
        _ = fi4
    }
    if (102 < 0) {
        var v5 float32 = 3.14159
        _ = v5
    } else {
        var v6 string = "hello"
        fmt.Println(v6)
        _ = v6
    }
    var v7 int16 = -1
    var v8 uint8 = 127
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 328
    var v3 uint8 = 42
    v9 := []int{128, 128, -2147483648, -1}
    switch ("a" + "foo bar") {
    case "a":
        _ = len(v9)
    case "\n\t":
        v3 = (v3 - 207)
    default:
        switch 32767 {
        case 128:
            if false {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v10 uint64 = 354
                _ = v10
            }
        case 4294967295:
            v11 := GenS1{F0: 100, F1: 4294967295}
            _ = v11
        }
    }
    var v12 float32 = (1e10 / 1.0)
    fmt.Println(v3)
    v12 += float32((1e10 * ((((0.0 - (0.0 - (((float64((3.14159 * (0.1 / (v12 * (0.0 / 1e10))))) - float64(v12)) + 0.1) * -2.25))) * (0.1 - ((((-2.25 / float64(v12)) - -2.25) / 1e10) * 1.5))) * float64(0.1)) + ((((1e10 - float64(3.14159)) * 1e10) + (0.1 * -2.25)) + 0.0))))
    if len(v9) > 0 {
        v9[0] = 32768
    }
    var v13 float64 = 3.14159
    var v14 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
