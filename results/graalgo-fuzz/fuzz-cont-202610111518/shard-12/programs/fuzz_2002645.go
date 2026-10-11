package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint64 = 42
    var v2 int64 = -155
    var v3 int = -9223372036854775808
    var v4 int16 = -1
    var v5 string = ((((("foo bar" + "\n\t") + ("hello" + "\n\t")) + (("hello" + "\n\t") + "\n\t")) + "foo bar") + ((("a" + "hello") + "") + ("hello" + "")))
    v4 = v4
    v6 := []int{}
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(fi7, fi7)
        switch (^(fi7)) {
        case 524:
            if ((2147483647 <= -9223372036854775808) && (!false)) {
                var v8 int = (v3 % (((v3 >> ((65535) & 63))) | 1))
                _ = v8
            }
        case 32768:
            fmt.Println(fi7, v4)
        default:
            if ((fi7 & v3) != (fi7 | (v3 * v3))) {
                var v9 int32 = 48
                v2 = -9223372036854775808
                _ = v9
            } else {
                if (9223372036854775807 <= v3) {
                    var v10 int32 = 2147483647
                    _ = v10
                }
            }
        }
        _ = fi7
    }
    var v11 int64 = (v2 / ((v2) | 1))
    if false {
        var v12 bool = (42 > (v3 << ((4294967295) & 63)))
        _ = v12
    }
    var v13 int16 = -32768
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            if false {
                fmt.Println(v2)
            } else {
                switch (((v5 + v5) + v5) + v5) {
                case "foo bar":
                    v11 = v11
                }
                v16 := GenS1{F0: 255, F1: 189}
                v5 = "\n\t"
                _ = v16
            }
            fmt.Println(v5)
            var v17 float64 = -2.25
            _ = fi15
            _ = v17
        }
        fmt.Println(v4, v3, v5)
        v18 := GenS1{F0: 2, F1: 2}
        _ = fi14
        _ = v18
    }
    v5 = (v5 + v5)
    fmt.Println(v13, v5, v4)
    if true {
        v5 = (("hello" + "a") + ("a" + ((v5 + ("" + (v5 + v5))) + v5)))
        var v19 int16 = (v4 - (v4 >> ((2) & 63)))
        _ = v19
    }
    v20 := GenS1{F0: 2, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
