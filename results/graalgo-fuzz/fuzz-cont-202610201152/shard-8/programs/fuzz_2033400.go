package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int16, p1 uint32) string {
    if (42 < 42) {
        p0 += 960
    } else {
        if (0 <= 102) {
            p0 = p0
            for fi5 := 0; fi5 < 9; fi5++ {
                var v6 uint = 310
                if (35 == 100) {
                    fmt.Println(p0, p0, p0)
                }
                _ = fi5
                _ = v6
            }
        } else {
            var v7 float64 = -2.25
            _ = v7
        }
    }
    if false {
        p0 = p0
        v8 := []int{-734}
        if false {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 bool = (!(42 <= 0))
                v10 = ("" <= (((("hello" + "hello") + "hello") + "a") + ((("" + "\n\t") + ("" + "\n\t")) + "\n\t")))
                _ = fi9
                _ = v10
            }
        } else {
            p0 = (p0 + p0)
            fmt.Println(p1)
        }
        _ = v8
    }
    v11 := []int{42, 0, 32768, -19}
    _ = p0
    _ = p1
    _ = v11
    return "a"
}

func main() {
    var v1 string = "foo bar"
    var v2 float32 = 0.1
    var v3 int8 = -1
    var v4 int8 = -1
    v12 := []int{402, 187, -2147483648}
    fmt.Println(v4, v2, v3)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    switch 838 {
    case -829:
        switch -9223372036854775808 {
        case 433:
            var v13 bool = false
            _ = v13
        }
    default:
        fmt.Println(v3, v4)
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS1{F0: -366}
        _ = fi14
        _ = v15
    }
    switch -2147483648 {
    case -724:
        v2 = (float32((1e10 * ((1e10 - 0.1) / (((1.5 / (float64(3.14159) / -2.25)) / 3.14159) + 1e10)))) * 1.5)
    case 352:
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint32 = 738
            for fi18 := 0; fi18 < 8; fi18++ {
                v3 -= (v4 << ((442) & 63))
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
    case 32767:
        var v19 uint8 = 156
        _ = v19
    }
    v1 = "\n\t"
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS2{F0: 4294967295}
        var v22 int32 = -2147483648
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int64 = -378
    if true {
        if (v3 > 103) {
            var v24 uint16 = 749
            _ = v24
        }
    }
    if (true && true) {
        var v25 uint16 = 2
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
