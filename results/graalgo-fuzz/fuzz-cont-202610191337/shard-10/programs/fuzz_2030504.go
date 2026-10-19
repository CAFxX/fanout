package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int8) uint16 {
    p0 = p0
    p0 += p0
    v5 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v5
    return 261
}
func genFn1(p0 int32, p1 uint32, p2 uint8) uint {
    p2 -= (^(p2))
    if true {
        p0 = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    return 912
}

func main() {
    var v1 int8 = 53
    v2 := []uint32{859, 42}
    var v3 string = (("foo bar" + "\n\t") + ((("hello" + "foo bar") + ("" + ("foo bar" + "\n\t"))) + ("" + "a")))
    var v4 int64 = -9223372036854775808
    if (18446744073709551615 < 0) {
        v1 -= (^(v1))
    } else {
        v4 = (v4 / (((v4 >> ((255) & 63))) | 1))
        for fi6 := 0; fi6 < 18; fi6++ {
            fmt.Println(v4)
            fmt.Println(fi6, v4, fi6)
            _ = fi6
        }
    }
    var v7 string = (v3 + v3)
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v3, v3, v3)
        _ = fi8
    }
    fmt.Println(v7, v3, v3)
    if true {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 float64 = -2.25
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 9223372036854775807}
        v12 := GenS1{F0: (^(v4))}
        _ = v11
        _ = v12
    } else {
        fmt.Println(v4)
        v1 -= (-(v1))
        switch -1 {
        case -2147483648:
            var v13 string = v3
            _ = v13
        case 695:
            for fi14 := 0; fi14 < 20; fi14++ {
                switch v7 {
                case "a":
                    v4 -= v4
                default:
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v1 += 80
                        fmt.Println(v4)
                        _ = fi15
                    }
                }
                _ = fi14
            }
        }
    }
    var v16 int64 = (-(v4))
    fmt.Println(v7, v4, v4)
    var v17 int = int(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
