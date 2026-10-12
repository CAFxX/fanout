package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
func genFn0() string {
    switch "\n\t" {
    case "a":
        if (!((1 != -128) && false)) {
            v6 := []int{42, 9223372036854775807}
            _ = v6
        }
    case "foo bar":
        var v7 uint16 = 302
        _ = v7
    case "":
        if false {
            var v8 int = 2147483647
            v8 = (^(v8))
            v8 += v8
            _ = v8
        }
    }
    var v9 int = -813
    _ = v9
    return ""
}

func main() {
    var v1 int = -9223372036854775808
    var v2 bool = ((119 == 26) || ("foo bar" != ""))
    var v3 int16 = 76
    var v4 uint8 = 1
    var v5 int32 = 1
    v10 := []int{(v1 - v1), 0}
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 = append(v10, v1)
        _ = fi11
    }
    if len(v10) > 0 {
        v10[0] = (-(v1))
    }
    var v12 uint64 = 501
    v4 -= (v4 >> ((255) & 63))
    var v13 bool = v2
    fmt.Println(v13, v3, v12)
    _ = len(v10)
    v1 = v1
    switch (v1 % (((^(v1))) | 1)) {
    case -2147483648:
        v1 = v1
    case 962:
        switch (("a" + "") + ("hello" + "\n\t")) {
        case "hello":
            switch "a" {
            case "a":
                v4 -= (^(v4))
            default:
                v4 = v4
            }
        case "foo bar":
            var v14 string = "foo bar"
            _ = v14
        }
    case 1:
        if false {
            var v15 int64 = 9223372036854775807
            _ = v15
        }
    }
    v10 = append(v10, (v1 ^ v1))
    if v13 {
        var v16 bool = ((3.14159 + float64((3.14159 + 1e10))) == 1.5)
        v17 := GenS1{F0: (v1 / (((-(v1))) | 1))}
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 float64 = 3.14159
            _ = fi18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    v20 := GenS2{F0: v3}
    v4 -= (v4 * (v4 & v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
