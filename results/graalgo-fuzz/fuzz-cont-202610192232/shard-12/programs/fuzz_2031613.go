package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint8 = 1
    var v2 int64 = 9223372036854775807
    var v3 int64 = (v2 | (-(v2)))
    for fi4 := 0; fi4 < 18; fi4++ {
        v1 -= v1
        fmt.Println(v2)
        _ = fi4
    }
    var v5 int64 = v3
    v6 := GenS1{F0: 230}
    var v7 uint8 = 0
    switch "\n\t" {
    case "hello":
        var v8 uint = 65535
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 12; fi9++ {
            v3 -= int64(fi9)
            var v10 int8 = 100
            if false {
                v10 += v10
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 int64 = 9223372036854775807
                    _ = fi11
                    _ = v12
                }
                var v13 uint32 = 2
                _ = v13
            }
            _ = fi9
            _ = v10
        }
    case "\n\t":
        v14 := []int{}
        _ = v14
    }
    var v15 string = ("hello" + "")
    v16 := GenS2{F0: 18446744073709551615}
    v1 += 1
    switch -9223372036854775808 {
    case -72:
        var v17 float32 = (0.0 / 1.5)
        _ = v17
    case 2147483647:
        fmt.Println(v7, v5)
    default:
        v5 -= 1
    }
    switch "" {
    case "hello":
        v18 := []int{0, -699}
        _ = v18
    case "a":
        v3 = -475
    default:
        v7 -= v7
    }
    switch ((v15 + v15) + (v15 + "foo bar")) {
    case "":
        var v19 int64 = v2
        _ = v19
    case "a":
        v1 = 99
    case "foo bar":
        var v20 string = v15
        _ = v20
    default:
        v21 := []int{}
        _ = v21
    }
    var v22 float64 = 0.1
    v2 += v2
    var v23 float64 = (v22 + ((float64((1.5 / ((1e10 * (1e10 + float32(-2.25))) - (0.1 * 1e10)))) + 0.0) - v22))
    fmt.Println(v5, v22, v1)
    switch 42 {
    case 42:
        v7 += v7
    default:
        if ((4294967295 > 728) && (true && false)) {
            var v24 string = "a"
            var v25 float64 = (0.1 / v22)
            _ = v24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
