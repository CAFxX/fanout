package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint16 = 0
    var v2 int32 = -2147483648
    var v3 float32 = 0.1
    var v4 uint64 = 1
    var v5 float32 = float32((-2.25 / 1e10))
    var v6 int64 = 42
    var v7 float64 = (1e10 + (3.14159 + (3.14159 - float64(v5))))
    v5 = v3
    for fi8 := 0; fi8 < 9; fi8++ {
        switch "hello" {
        case "hello":
            var v9 bool = (!((v2 + 1) < 2147483647))
            _ = v9
        case "":
            v4 = 2
        default:
            switch fi8 {
            case 127:
                var v10 int32 = v2
                _ = v10
            case 232:
                switch "foo bar" {
                case "\n\t":
                    var v11 int16 = 203
                    _ = v11
                case "":
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v13 := GenS1{F0: (v4 & 560)}
                        var v14 int8 = 100
                        var v15 string = (((("a" + ("\n\t" + ("\n\t" + "\n\t"))) + ("hello" + ("foo bar" + ""))) + ("" + "\n\t")) + ("" + "foo bar"))
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                case "a":
                    fmt.Println(v7, fi8, v1)
                }
            }
        }
        fmt.Println(v3, fi8)
        v16 := GenS1{F0: v4}
        _ = fi8
        _ = v16
    }
    var v17 uint64 = (v4 ^ v4)
    fmt.Println(v1)
    var v18 bool = false
    v19 := []int{-923, 0, 9223372036854775807}
    if false {
        v20 := GenS1{F0: v4}
        _ = v20
    } else {
        if v18 {
            _ = len(v19)
        }
        var v21 float64 = v7
        _ = v21
    }
    var v22 int64 = -58
    fmt.Println(v18)
    var v23 uint64 = 2
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 string = ("a" + "a")
        v26 := GenS1{F0: (v23 << ((2) & 63))}
        _ = fi24
        _ = v25
        _ = v26
    }
    v23 -= (v23 + 18446744073709551615)
    var v27 int = 32767
    v4 -= v4
    var v28 uint64 = (v4 | v17)
    v2 -= -2147483648
    fmt.Println(v27, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
