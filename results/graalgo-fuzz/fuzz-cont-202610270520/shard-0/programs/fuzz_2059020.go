package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
func genFn0(p0 uint32, p1 uint16) int16 {
    v5 := GenS1{F0: (int(p1) > 196), F1: ("foo bar" + "foo bar"), F2: 28}
    var v6 uint32 = (^(p0))
    var v7 uint32 = (v6 ^ (v6 | 2147483647))
    p0 = (v6 + (v7 << ((4294967295) & 63)))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -876
    v3 := []string{"foo bar"}
    var v4 int8 = 127
    var v8 uint16 = 0
    var v9 int8 = v4
    v9 = v9
    var v10 uint32 = 2147483647
    var v11 float32 = 3.14159
    switch "" {
    case "a":
        v10 = v10
    case "\n\t":
        if false {
            if (-2147483648 < -2147483648) {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    if true {
                        var v12 float64 = 0.1
                        _ = v12
                    } else {
                        v13 := GenS1{F0: (581 != 2), F1: "\n\t", F2: 42}
                        var v14 uint64 = 0
                        _ = v13
                        _ = v14
                    }
                case "":
                    if (true || (18446744073709551615 < 542)) {
                        var v15 uint8 = 237
                        _ = v15
                    } else {
                        var v16 int8 = -128
                        var v17 float64 = -2.25
                        _ = v16
                        _ = v17
                    }
                default:
                    v11 += (v11 - v11)
                }
                v2 -= v2
            } else {
                switch 418 {
                case 2147483647:
                    _ = len(v3)
                case 32767:
                    _ = len(v3)
                default:
                    v8 = 255
                }
            }
        }
    }
    fmt.Println(v8, v2, v2)
    var v18 int64 = 230
    v19 := GenS1{F0: false, F1: ("hello" + (("foo bar" + ("hello" + "")) + "")), F2: 54}
    var v20 int16 = 1000
    v11 -= v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
