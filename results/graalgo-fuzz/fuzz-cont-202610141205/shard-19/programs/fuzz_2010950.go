package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 182
    var v3 int8 = 90
    switch 32768 {
    case 2147483647:
        var v4 uint16 = 0
        _ = v4
    case 9223372036854775807:
        var v5 int16 = 601
        _ = v5
    default:
        var v6 int32 = -994
        _ = v6
    }
    var v7 bool = false
    v3 = -63
    v2 += v2
    if (42 < uint16(v2)) {
        v2 = v2
        v8 := []int{}
        switch -2147483648 {
        case 1:
            switch 0 {
            case 229:
                fmt.Println(v7, v3, v3)
            case 9223372036854775807:
                switch v1 {
                case "\n\t":
                    if (!v7) {
                        var v9 int64 = -177
                        v1 = ""
                        _ = v9
                    }
                case "a":
                    v10 := GenS1{F0: uint32(v2)}
                    _ = v10
                case "":
                    switch v1 {
                    case "":
                        v3 = v3
                    case "\n\t":
                        for fi11 := 0; fi11 < 10; fi11++ {
                            var v12 uint32 = 2147483648
                            v13 := GenS2{F0: 127}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                default:
                    if len(v8) > 0 {
                        v8[0] = 42
                    }
                }
            }
        case -645:
            fmt.Println(v2)
        case -9223372036854775808:
            if (42 > 766) {
                v2 = (v2 - (v2 + 2))
                v8 = append(v8, -110)
            }
        }
        _ = v8
    } else {
        v1 = "hello"
        var v14 uint = 4294967295
        v15 := []int{4294967295, 127, 32768, -1}
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 2}
    fmt.Println(v7, v2, v2)
    v17 := GenS2{F0: 128}
    var v18 int16 = 0
    fmt.Println(v7)
    fmt.Println(v7, v2)
    v1 = v1
    switch 576 {
    case 207:
        var v19 float64 = 1.5
        _ = v19
    case 83:
        var v20 int16 = v18
        _ = v20
    case 0:
        v21 := []int{-9223372036854775808, 457}
        _ = v21
    default:
        var v22 string = ""
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
