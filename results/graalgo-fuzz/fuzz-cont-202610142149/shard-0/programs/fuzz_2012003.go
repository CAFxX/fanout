package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0() int8 {
    var v7 uint = 21
    if false {
        if (!true) {
            v8 := GenS1{F0: 100, F1: 32767}
            var v9 uint64 = 0
            _ = v8
            _ = v9
        } else {
            var v10 bool = (true || false)
            var v11 float32 = (0.1 + 0.1)
            var v12 int32 = 2147483647
            _ = v10
            _ = v11
            _ = v12
        }
    }
    var v13 int16 = 1
    v7 = (v7 % (((v7 / ((v7) | 1))) | 1))
    _ = v13
    _ = v7
    return 127
}
func genFn1(p0 int32) int {
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := []int{(^(fi14)), (fi14 << ((18446744073709551615) & 63)), fi14, 0}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 float32 = 0.0
        _ = fi16
        _ = v17
    }
    if true {
        p0 += p0
    }
    _ = p0
    return -789
}

func main() {
    var v1 int8 = -31
    var v2 int8 = 100
    var v3 int16 = 0
    var v4 int16 = -32768
    v5 := []int{484}
    var v6 int8 = 1
    v3 = v3
    var v18 uint16 = 684
    var v19 uint8 = 127
    v3 = v3
    var v20 float64 = ((1e10 / 1.5) / ((0.1 - 1.5) * (-2.25 / 1.5)))
    fmt.Println(v6, v1)
    var v21 float64 = (v20 / (0.0 - v20))
    var v22 uint = 65535
    var v23 string = ("foo bar" + ("foo bar" + ("hello" + ("\n\t" + "foo bar"))))
    if false {
        v19 -= v19
        fmt.Println(v2, v2)
        var v24 float64 = v21
        _ = v24
    }
    v4 = 42
    v5 = append(v5, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
