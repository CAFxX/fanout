package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
func genFn0() int8 {
    var v7 uint = 65535
    v8 := []int{}
    v7 -= (^(v7))
    v9 := GenS1{F0: false, F1: -9223372036854775808}
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() bool {
    v10 := []int{112}
    if (1 > 18446744073709551615) {
        v10 = append(v10, 2147483647)
    }
    var v11 int = 128
    var v12 uint32 = 0
    for fi13 := 0; fi13 < 18; fi13++ {
        v12 = 0
        switch v11 {
        case -9223372036854775808:
            var v14 int = fi13
            _ = v14
        case -91:
            for fi15 := 0; fi15 < 13; fi15++ {
                _ = len(v10)
                v11 -= v11
                var v16 int64 = -834
                _ = fi15
                _ = v16
            }
        default:
            v17 := GenS1{F0: (!(false || ((true || false) && (((!((((-2.25 / (float64(0.1) - float64((3.14159 - float32(0.1))))) * 0.0) * (3.14159 * -2.25)) >= 0.1)) && false) && (30 < 860))))), F1: (v11 ^ (fi13 / (((fi13 << ((2) & 63))) | 1)))}
            _ = v17
        }
        fmt.Println(v11)
        _ = fi13
    }
    _ = v10
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]uint32)
    var v3 float64 = 1e10
    v4 := make(map[string]string)
    var v5 uint = 807
    var v6 string = "a"
    v18 := v4[v6]
    var v19 uint16 = 0
    v6 = v18
    v5 = (v5 + 18446744073709551615)
    v1 = v1
    switch v6 {
    case "foo bar":
        var v20 int32 = -643
        _ = v20
    case "a":
        if (-128 >= -32768) {
            if ((v19 | (v19 >> (((v5 ^ v5)) & 63))) > (^(v19))) {
                v5 = 4294967295
                fmt.Println(v5, v19, v6)
                v19 = (v19 + v19)
            }
        }
    case "":
        fmt.Println(v6)
    default:
        for fi21 := 0; fi21 < 9; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                _ = len(v2)
                _ = fi22
            }
            var v23 bool = (774 > 49)
            _ = fi21
            _ = v23
        }
    }
    v6 = ""
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS1{F0: (!false), F1: (fi24 ^ (-(fi24)))}
        var v26 bool = (-1 == 1)
        v6 = (v18 + ((v18 + (("foo bar" + ((v18 + v18) + v18)) + "")) + "\n\t"))
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := v4[v18]
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
