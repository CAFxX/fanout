package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 uint32 = 2147483648
    var v2 int64 = -9223372036854775808
    var v3 float32 = float32(((1.5 - 1e10) - (((-2.25 / (1.5 * 3.14159)) - 3.14159) - 1.5)))
    v4 := []string{}
    var v5 int8 = -1
    var v6 uint32 = v1
    switch -609 {
    case -2147483648:
        v4 = append(v4, "")
    case 32768:
        v7 := GenS1{F0: 127, F1: "\n\t"}
        _ = v7
    case 312:
        if (0 > 0) {
            var v8 uint32 = v6
            if len(v4) > 0 {
                v4[0] = ("hello" + "foo bar")
            }
            v8 = (^(v6))
            _ = v8
        } else {
            var v9 int64 = (v2 % (((v2 - (^(v2)))) | 1))
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        _ = fi10
    }
    if false {
        v3 += v3
    } else {
        fmt.Println(v6)
        if (618 <= 412) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 = (v5 & v5)
        }
        fmt.Println(v5, v3)
    }
    var v11 uint32 = (v6 | (v6 * (v6 >> ((42) & 63))))
    var v12 bool = (75 == 42)
    var v13 string = "foo bar"
    v1 = (v6 % ((4294967295) | 1))
    v12 = v12
    var v14 float64 = -2.25
    v3 = v3
    if v12 {
        fmt.Println(v2, v6)
        fmt.Println(v12, v3, v13)
        var v15 float64 = v14
        _ = v15
    } else {
        fmt.Println(v13)
        fmt.Println(v6, v2, v14)
    }
    v16 := GenS1{F0: (v5 >> ((65535) & 63)), F1: v13}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
