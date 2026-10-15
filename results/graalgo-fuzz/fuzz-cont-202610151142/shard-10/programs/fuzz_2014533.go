package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 int = -2147483648
    v3 := make(map[string]string)
    v4 := make(map[uint32]bool)
    if (("a" + "hello") == "\n\t") {
        switch "a" {
        case "hello":
            v5 := GenS1{F0: -128, F1: 42, F2: 18446744073709551615}
            _ = v5
        default:
            _ = len(v4)
        }
        v2 = v2
        if false {
            for fi6 := 0; fi6 < 10; fi6++ {
                v2 += -428
                v2 += (-(v2))
                for fi7 := 0; fi7 < 4; fi7++ {
                    delete(v4, 4294967295)
                    var v8 float64 = 1e10
                    v2 -= (v2 * (fi6 % ((881) | 1)))
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
        } else {
            var v9 uint32 = 2147483647
            v9 += v9
            v10 := GenS1{F0: 1, F1: 9223372036854775807, F2: 831}
            _ = v10
            _ = v9
        }
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            v2 += (^(fi11))
            v12 := GenS1{F0: 127, F1: -1, F2: 0}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := []int{fi13}
            _ = len(v3)
            _ = fi13
            _ = v14
        }
    }
    var v15 uint = 0
    if (9223372036854775807 <= -1) {
        v2 += -9223372036854775808
        var v16 int = (v2 ^ (v2 & v2))
        _ = v16
    }
    v17 := GenS1{F0: -128, F1: 9223372036854775807, F2: 129}
    v18 := v4[4294967295]
    v19 := []int{(v2 - 42), -9223372036854775808}
    fmt.Println(v2, v2, v15)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v15 = (v15 + (v15 + (^(v15))))
    if (v15 > 2) {
        v2 -= v2
        v15 += 0
    }
    if true {
        _ = len(v1)
        v20 := GenS1{F0: -32, F1: 42, F2: v15}
        switch ("hello" + (("\n\t" + "") + ("hello" + "a"))) {
        case "hello":
            if (287 <= -334) {
                switch v2 {
                case 42:
                    fmt.Println(v18, v15)
                }
                if v18 {
                    var v21 float32 = -2.25
                    _ = v21
                } else {
                    v18 = true
                    fmt.Println(v2)
                }
            }
        case "":
            v3["a"] = (("foo bar" + ("hello" + "foo bar")) + "")
        default:
            var v22 int = (v2 + v2)
            _ = v22
        }
        _ = v20
    } else {
        switch ("" + ("hello" + ("" + "\n\t"))) {
        case "\n\t":
            v23 := GenS1{F0: 42, F1: 9223372036854775807, F2: (^(v15))}
            _ = v23
        }
    }
    switch "foo bar" {
    case "":
        var v24 uint16 = 251
        _ = v24
    case "\n\t":
        v15 += 65535
    case "a":
        var v25 int64 = 42
        _ = v25
    default:
        _ = len(v3)
    }
    var v26 int16 = -1
    v27 := GenS1{F0: 30, F1: 42, F2: 18446744073709551615}
    v18 = (3.14159 <= 0.1)
    v28 := GenS1{F0: 127, F1: 0, F2: (v15 * v15)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
