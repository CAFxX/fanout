package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0() int64 {
    var v4 bool = ("foo bar" == "")
    var v5 float64 = 1.5
    v5 = ((0.0 * (v5 + 1.5)) * v5)
    if (!v4) {
        switch "" {
        case "":
            var v6 int8 = 38
            _ = v6
        case "a":
            v5 += float64(0.1)
        case "foo bar":
            v7 := []int{4294967295, 32767, 190}
            _ = v7
        }
        var v8 uint = 255
        switch (("" + "foo bar") + "\n\t") {
        case "\n\t":
            var v9 bool = ("\n\t" < "a")
            _ = v9
        default:
            var v10 uint = 2
            _ = v10
        }
        _ = v8
    }
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 bool = ((-1 == 537) || (true || false))
    var v2 int = -9223372036854775808
    var v3 int32 = -153
    var v11 bool = v1
    fmt.Println(v2, v1)
    var v12 uint16 = 65535
    fmt.Println(v3, v1, v3)
    v11 = v1
    v12 = (^(v12))
    v13 := GenS1{F0: 79, F1: (-(v2))}
    var v14 string = "a"
    v15 := GenS1{F0: 128, F1: 9223372036854775807}
    var v16 uint = 2
    if v11 {
        var v17 string = ("" + v14)
        _ = v17
    } else {
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v14)
            var v19 int32 = int32(v12)
            if (v2 <= (fi18 - (v2 % ((v2) | 1)))) {
                v2 = v2
                for fi20 := 0; fi20 < 3; fi20++ {
                    v21 := GenS1{F0: 2, F1: (fi20 & fi20)}
                    _ = fi20
                    _ = v21
                }
                v22 := []int{(fi18 ^ 42), v2, -2147483648}
                _ = v22
            }
            _ = fi18
            _ = v19
        }
        v16 = 1
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int = (-(v2))
            _ = fi23
            _ = v24
        }
    }
    v25 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
