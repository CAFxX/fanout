package main

import (
    "fmt"
)

func genFn0(p0 uint16, p1 uint) uint {
    v5 := []int{-9223372036854775808}
    if false {
        for fi6 := 0; fi6 < 1; fi6++ {
            switch fi6 {
            case 32768:
                p0 = p0
            default:
                p0 -= p0
            }
            p0 += (^(p0))
            switch "hello" {
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = (^(fi6))
                }
            case "":
                if (fi6 != 229) {
                    p0 = (^(p0))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            case "foo bar":
                var v7 uint16 = p0
                _ = v7
            default:
                var v8 bool = (!(-120 != int8(p1)))
                _ = v8
            }
            _ = fi6
        }
        var v9 int64 = -9223372036854775808
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}

func main() {
    v1 := []int{-2147483648}
    var v2 float32 = 0.1
    var v3 uint = 1
    var v4 int8 = 100
    var v10 uint8 = 223
    var v11 bool = ((float32(0.0) + v2) > ((float32(float64(v2)) / v2) - v2))
    var v12 float32 = float32(1.5)
    fmt.Println(v2)
    var v13 float32 = (v2 - v2)
    v11 = v11
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v4, v3)
        if (2147483647 < -2147483648) {
            switch "hello" {
            case "hello":
                var v15 uint16 = 0
                _ = v15
            default:
                if v11 {
                    fmt.Println(v3, fi14, v12)
                    v11 = true
                } else {
                    var v16 uint16 = 32767
                    _ = v16
                }
            }
            _ = len(v1)
        }
        _ = fi14
    }
    v3 = genFn0(46, (v3 / ((v3) | 1)))
    fmt.Println(v4, v3, v10)
    for fi17 := 0; fi17 < 12; fi17++ {
        v10 = 2
        _ = fi17
    }
    var v18 uint16 = 65535
    fmt.Println(v3, v18, v2)
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint32 = 2147483647
        var v21 int8 = v4
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
