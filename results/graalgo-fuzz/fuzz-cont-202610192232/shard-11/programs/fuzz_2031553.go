package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    var v1 uint = 0
    var v2 uint32 = 470
    v3 := []int{-2147483648, -1}
    var v4 float64 = 1.5
    v5 := make(map[string]string)
    v2 = v2
    if (v2 <= (^(v2))) {
        var v6 uint8 = 128
        _ = v6
    }
    switch 810 {
    case 0:
        var v7 uint = (v1 ^ v1)
        _ = v7
    case 9223372036854775807:
        var v8 string = "a"
        _ = v8
    default:
        v2 = v2
    }
    v9 := GenS1{F0: 42, F1: "hello"}
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    var v10 uint8 = uint8(v1)
    _ = len(v5)
    if (true && (!(630 <= 4294967295))) {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: -1, F1: "foo bar"}
            _ = fi11
            _ = v12
        }
        var v13 int16 = 1000
        delete(v5, ("hello" + ("foo bar" + "hello")))
        _ = v13
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS1{F0: -9223372036854775808, F1: (("hello" + ("a" + "")) + "foo bar")}
                _ = fi15
                _ = v16
            }
            var v17 uint32 = (v2 % (((^(v2))) | 1))
            _ = fi14
            _ = v17
        }
    }
    fmt.Println(v10)
    v18 := v5[("hello" + "a")]
    v10 -= (v10 * (v10 % ((v10) | 1)))
    delete(v5, v18)
    fmt.Println(v18)
    if len(v3) > 0 {
        v3[0] = 32768
    }
    v19 := v5[v18]
    v5[(("foo bar" + v18) + ((v19 + v18) + "hello"))] = "foo bar"
    v20 := v5[((v19 + "hello") + (v19 + (v19 + (v19 + "\n\t"))))]
    switch v18 {
    case "\n\t":
        var v21 int32 = 1
        _ = v21
    case "a":
        for fi22 := 0; fi22 < 9; fi22++ {
            if false {
                if len(v3) > 0 {
                    v3[0] = (fi22 ^ fi22)
                }
                v18 = "foo bar"
            }
            _ = fi22
        }
    case "foo bar":
        switch 9223372036854775807 {
        case -9223372036854775808:
            v2 = (v2 % ((v2) | 1))
        case 4294967295:
            delete(v5, (v20 + v18))
        case -481:
            v23 := GenS1{F0: -505, F1: v20}
            _ = v23
        default:
            v24 := GenS1{F0: -9223372036854775808, F1: ""}
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
