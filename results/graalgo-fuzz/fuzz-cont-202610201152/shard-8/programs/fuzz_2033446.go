package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }

func main() {
    var v1 uint64 = 0
    var v2 uint64 = (v1 % (((v1 | v1)) | 1))
    var v3 uint32 = 2147483648
    var v4 uint8 = 128
    var v5 string = ""
    v2 = (v1 * v1)
    var v6 uint32 = (v3 - v3)
    var v7 bool = false
    for fi8 := 0; fi8 < 18; fi8++ {
        v3 -= v6
        var v9 int8 = 47
        for fi10 := 0; fi10 < 19; fi10++ {
            if (v7 || (((0.0 * 1.5) + 1.5) > float32(((-2.25 * float64(0.0)) - 0.1)))) {
                switch fi8 {
                case -9223372036854775808:
                    switch v5 {
                    case "":
                        var v11 uint16 = 2
                        _ = v11
                    case "\n\t":
                        var v12 uint32 = 2147483647
                        _ = v12
                    default:
                        var v13 string = "hello"
                        _ = v13
                    }
                case 4294967295:
                    v14 := []int{}
                    _ = v14
                case 127:
                    if (!v7) {
                        fmt.Println(v9)
                        switch 2147483647 {
                        case -1:
                            var v15 float32 = 3.14159
                            _ = v15
                        default:
                            v16 := []int{128}
                            _ = v16
                        }
                        v1 = uint64(v4)
                    }
                }
            }
            var v17 uint16 = 826
            v2 = (v2 * uint64(v4))
            _ = fi10
            _ = v17
        }
        _ = fi8
        _ = v9
    }
    var v18 float32 = (float32(1e10) * 0.1)
    switch -9223372036854775808 {
    case 32767:
        var v19 uint = 65535
        _ = v19
    case 127:
        var v20 uint8 = (v4 ^ uint8(v3))
        _ = v20
    case 32768:
        v1 += v2
    default:
        var v21 string = "foo bar"
        _ = v21
    }
    switch ("" + "hello") {
    case "":
        switch -1 {
        case 811:
            v22 := GenS1{F0: 0, F1: v7, F2: v5}
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
