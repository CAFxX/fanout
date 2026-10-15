package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 9223372036854775807
    var v3 string = "\n\t"
    var v4 uint = 1
    v5 := make(map[string]bool)
    if (1.5 <= 1.5) {
        switch v2 {
        case 32767:
            var v6 int16 = 42
            _ = v6
        case -536:
            v4 = (^(v4))
        default:
            v5[("hello" + "foo bar")] = (1 > 2)
        }
    } else {
        if (v4 != v4) {
            var v7 float64 = 1e10
            _ = v7
        } else {
            v8 := GenS1{F0: 32767, F1: 425, F2: 0}
            if (42 != 0) {
                var v9 float32 = float32((float64(0.0) / 1e10))
                if false {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 uint = (v4 << ((65535) & 63))
                        _ = fi10
                        _ = v11
                    }
                    var v12 int16 = 1000
                    _ = v12
                } else {
                    var v13 uint32 = 387
                    _ = v13
                }
                v14 := []int{v2, v2}
                _ = v14
                _ = v9
            }
            var v15 int64 = 0
            _ = v15
            _ = v8
        }
        if false {
            v16 := []int{v2}
            _ = v16
        }
        fmt.Println(v2)
    }
    v2 -= 128
    fmt.Println(v4, v3, v4)
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 string = ((v3 + "\n\t") + "foo bar")
            v4 = v4
            v19 = (("foo bar" + "") + v3)
            _ = fi18
            _ = v19
        }
        fmt.Println(v2, v3)
        fmt.Println(v2, v4, v4)
        _ = fi17
    }
    if true {
        v20 := []int{}
        if false {
            for fi21 := 0; fi21 < 12; fi21++ {
                if (-2147483648 != 989) {
                    v4 += (v4 * v4)
                    var v22 float64 = ((((3.14159 * 1e10) / ((((float64(((0.0 - (1e10 / 1.0)) - -2.25)) / 3.14159) * 1e10) / 0.1) + 0.0)) - (0.0 * (3.14159 - (((float64(0.1) - 1e10) + 0.0) / 0.1)))) / (1.5 + (0.1 / ((0.1 - 1.5) - 3.14159))))
                    v23 := GenS1{F0: (fi21 % ((v2) | 1)), F1: 162, F2: 0}
                    _ = v22
                    _ = v23
                } else {
                    v2 = fi21
                }
                _ = fi21
            }
        } else {
            var v24 uint32 = 899
            fmt.Println(v4, v4)
            _ = v24
        }
        var v25 float64 = (3.14159 - 3.14159)
        _ = v20
        _ = v25
    }
    var v26 float64 = 0.1
    for fi27 := 0; fi27 < 8; fi27++ {
        v26 += ((1e10 + 0.1) * (v26 + -2.25))
        _ = fi27
    }
    switch -2147483648 {
    case -375:
        if (486 < 0) {
            var v28 uint = v4
            v29 := []int{4294967295}
            var v30 int64 = 961
            _ = v28
            _ = v29
            _ = v30
        } else {
            var v31 float32 = 1.5
            _ = len(v1)
            _ = v31
        }
    default:
        var v32 uint64 = 18446744073709551615
        _ = v32
    }
    v33 := GenS1{F0: v2, F1: 18446744073709551615, F2: 1}
    var v34 float64 = v26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
