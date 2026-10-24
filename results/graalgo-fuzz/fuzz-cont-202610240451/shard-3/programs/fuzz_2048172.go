package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{127, 127, 2147483647, 2147483647}
    var v2 int32 = -2147483648
    var v3 int8 = 42
    v4 := make(map[string]bool)
    var v5 float32 = 0.1
    v6 := make(map[string]int)
    v7 := GenS1{F0: 42}
    if true {
        var v8 uint8 = 104
        var v9 int16 = 1000
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 127}
        if true {
            var v11 int16 = 1
            v5 += (-2.25 / v5)
            _ = v11
        } else {
            v2 = (v2 & (v2 & v2))
            v2 = (v2 | (v2 & (v2 / (((v2 % ((-1) | 1))) | 1))))
            switch -9223372036854775808 {
            case 0:
                v5 = v5
            default:
                switch "foo bar" {
                case "\n\t":
                    var v12 uint8 = 1
                    _ = v12
                default:
                    v2 -= -454
                }
            }
        }
        _ = v10
    }
    v2 += (v2 - 42)
    var v13 float64 = 3.14159
    v4["hello"] = false
    fmt.Println(v3)
    if false {
        v14 := GenS2{F0: v2}
        var v15 string = ("\n\t" + "a")
        var v16 uint64 = 250
        _ = v14
        _ = v15
        _ = v16
    } else {
        var v17 uint64 = 1
        v2 += 724
        if false {
            v3 = (-(v3))
            fmt.Println(v17)
            fmt.Println(v17, v5)
        }
        _ = v17
    }
    if true {
        var v18 uint64 = 2
        _ = v18
    }
    switch "" {
    case "":
        fmt.Println(v5)
    case "hello":
        fmt.Println(v3)
    default:
        if (false || (((0.0 / -2.25) - 0.0) != v5)) {
            var v19 int = 990
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint8 = 255
                switch "\n\t" {
                case "":
                    v19 += (fi20 % ((v19) | 1))
                case "hello":
                    var v22 uint = 42
                    _ = v22
                case "\n\t":
                    var v23 uint = 451
                    _ = v23
                default:
                    _ = len(v1)
                }
                _ = fi20
                _ = v21
            }
            switch "" {
            case "":
                var v24 string = ""
                _ = v24
            case "foo bar":
                v25 := GenS2{F0: (^(v2))}
                _ = v25
            case "\n\t":
                fmt.Println(v19)
            }
            _ = v19
        }
    }
    v3 += v3
    fmt.Println(v13, v13, v2)
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 uint8 = 42
        if (-9223372036854775808 == 0) {
            switch "\n\t" {
            case "\n\t":
                if (0 < 775) {
                    v28 := GenS1{F0: (^(v27))}
                    _ = v28
                }
            default:
                fmt.Println(v3, v3)
            }
            _ = len(v6)
        } else {
            v29 := GenS2{F0: (^(v2))}
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    v13 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
