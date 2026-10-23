package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }

func main() {
    var v1 uint16 = 0
    v2 := []string{"hello", "foo bar"}
    var v3 int8 = 123
    var v4 int32 = -211
    var v5 float64 = 1.5
    v1 = 1
    v2 = append(v2, (("\n\t" + "") + ("hello" + "")))
    var v6 string = (("foo bar" + "hello") + "\n\t")
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 string = ("foo bar" + "a")
        var v9 string = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 -= 1.5
    v10 := GenS1{F0: 2147483648, F1: 89, F2: v1}
    var v11 uint8 = 127
    v1 += v1
    if false {
        switch 0 {
        case -9223372036854775808:
            _ = len(v2)
        default:
            var v12 string = v6
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int = (-(fi13))
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint32 = 0
                _ = fi15
                _ = v16
            }
            v4 = v4
            _ = fi13
            _ = v14
        }
        v17 := GenS1{F0: 42, F1: 4294967295, F2: 11}
        _ = v17
    }
    v11 += (v11 | (v11 & v11))
    v6 = v6
    if (1e10 != 1.5) {
        var v18 uint8 = v11
        var v19 int = 4294967295
        _ = v18
        _ = v19
    }
    fmt.Println(v5, v4)
    v11 += (v11 & uint8(v1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v3 = (v3 ^ (^(v3)))
        v6 = (((((v6 + v6) + "") + v6) + (((("" + ("foo bar" + v6)) + ("foo bar" + "a")) + "") + v6)) + "foo bar")
        v21 := GenS1{F0: 2147483647, F1: 1, F2: v1}
        _ = fi20
        _ = v21
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
