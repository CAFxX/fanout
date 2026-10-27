package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }

func main() {
    var v1 bool = ((true || (!false)) || false)
    var v2 float64 = ((((((float64(float32((1e10 + 3.14159))) + ((0.1 + 0.0) - 1e10)) + 1.5) + (0.1 + -2.25)) - 1e10) - 0.1) - 1e10)
    var v3 int64 = -9223372036854775808
    var v4 int16 = -1
    switch "\n\t" {
    case "foo bar":
        v5 := []int{0, -2147483648, 32768}
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{fi6, 2147483647}
            switch "a" {
            case "a":
                var v8 uint64 = 4294967295
                _ = v8
            case "\n\t":
                if v1 {
                    v7 = append(v7, (fi6 << ((525) & 63)))
                }
            case "hello":
                if len(v7) > 0 {
                    v7[0] = fi6
                }
            }
            _ = fi6
            _ = v7
        }
    case "":
        var v9 uint16 = 32767
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v3 = v3
        _ = fi10
    }
    var v11 float32 = 3.14159
    v12 := GenS1{F0: 1000, F1: 100, F2: 42}
    fmt.Println(v3, v3)
    v13 := GenS1{F0: v4, F1: 127, F2: 42}
    v14 := []int{1, 995}
    for fi15 := 0; fi15 < 10; fi15++ {
        switch (("" + "\n\t") + "\n\t") {
        case "foo bar":
            var v16 uint64 = 4294967295
            _ = v16
        case "":
            var v17 int = (fi15 - (fi15 * fi15))
            _ = v17
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
