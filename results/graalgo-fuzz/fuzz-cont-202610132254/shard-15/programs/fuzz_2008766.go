package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = 42
    var v3 float32 = -2.25
    fmt.Println(v1, v3)
    var v4 uint16 = 2
    switch 42 {
    case -114:
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := []int{(fi5 << ((4294967295) & 63)), 41, fi5}
            v2 = v2
            var v7 string = ""
            _ = fi5
            _ = v6
            _ = v7
        }
    case 42:
        v8 := []int{2147483647, 127, -9223372036854775808, 59}
        _ = v8
    }
    fmt.Println(v2)
    if true {
        var v9 uint16 = v4
        _ = v9
    }
    var v10 int = 785
    fmt.Println(v10)
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int64 = int64(v10)
        v13 := GenS1{F0: 0, F1: 0}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{v10}
    v1 = v1
    if true {
        if (32767 >= v4) {
            switch ("hello" + ("a" + ("" + (((("a" + (("" + "\n\t") + ((("hello" + "") + "\n\t") + ("foo bar" + ("foo bar" + "\n\t"))))) + "foo bar") + "") + ("\n\t" + ""))))) {
            case "a":
                if false {
                    var v15 uint64 = 2
                    v3 -= (((-2.25 / (0.1 * ((v3 * 0.0) + v3))) / 1.5) / (v3 + -2.25))
                    switch 2147483647 {
                    case 129:
                        var v16 bool = (false && false)
                        _ = v16
                    default:
                        fmt.Println(v2)
                    }
                    _ = v15
                }
            case "foo bar":
                if ((0 == uint8(v1)) || ((-126 <= -1) || false)) {
                    v17 := GenS1{F0: 0, F1: 51}
                    _ = len(v14)
                    _ = v17
                }
            case "\n\t":
                var v18 bool = true
                _ = v18
            }
            var v19 uint32 = v1
            v2 += (v2 & (v2 ^ 9223372036854775807))
            _ = v19
        } else {
            var v20 int8 = 20
            v20 = v20
            _ = v20
        }
    }
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
