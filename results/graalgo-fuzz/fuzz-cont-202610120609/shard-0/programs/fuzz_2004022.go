package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 string }

func main() {
    var v1 int8 = 1
    var v2 uint32 = 1
    var v3 int8 = v1
    var v4 string = ("hello" + "\n\t")
    v5 := make(map[uint32]string)
    var v6 uint32 = (v2 % ((v2) | 1))
    v1 += v3
    switch 2147483647 {
    case 128:
        v7 := GenS1{F0: 42, F1: 128}
        _ = v7
    case -463:
        v8 := []int{961, -2147483648, 9223372036854775807, 589}
        _ = v8
    case -482:
        for fi9 := 0; fi9 < 9; fi9++ {
            if true {
                v10 := GenS1{F0: 9223372036854775807, F1: 4}
                _ = v10
            }
            _ = fi9
        }
    default:
        v2 -= (v6 ^ 0)
    }
    if true {
        v6 += (v6 ^ (v2 ^ 1))
    } else {
        v5[v6] = v4
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint64 = 431
            v13 := GenS1{F0: -9223372036854775808, F1: 1}
            v14 := []int{(fi11 << ((2) & 63))}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v5[v2] = ""
    }
    v15 := []int{2147483647, 4294967295, -9223372036854775808}
    if false {
        var v16 float64 = (((0.0 / 3.14159) + -2.25) - 0.1)
        _ = v16
    }
    _ = len(v15)
    for fi17 := 0; fi17 < 11; fi17++ {
        if (!false) {
            if true {
                var v18 int = (fi17 | fi17)
                var v19 float32 = ((3.14159 / (3.14159 / ((3.14159 * float32((3.14159 / ((1e10 * float64(1e10)) - -2.25)))) / float32(1e10)))) + -2.25)
                _ = v18
                _ = v19
            } else {
                v2 += (v6 * (v6 * v6))
            }
            v15 = append(v15, (fi17 & fi17))
        }
        var v20 int64 = 42
        _ = fi17
        _ = v20
    }
    var v21 uint = 2
    if (true && false) {
        if ((-2.25 / (3.14159 - (1.5 * (1e10 / 1e10)))) > 3.14159) {
            if (!(false || (false && false))) {
                if (-2147483648 <= 42) {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 uint16 = 42
                        _ = fi22
                        _ = v23
                    }
                    v4 = "foo bar"
                } else {
                    fmt.Println(v3, v21, v1)
                }
                v24 := GenS1{F0: -1, F1: 123}
                _ = v24
            } else {
                var v25 int = 1
                _ = v25
            }
        }
        v26 := GenS1{F0: 9223372036854775807, F1: 2}
        v21 = v21
        _ = v26
    } else {
        var v27 float32 = 0.0
        _ = v27
    }
    var v28 int32 = 2147483647
    if true {
        var v29 float32 = 1.5
        v6 = v6
        switch 128 {
        case 621:
            var v30 uint = (^(v21))
            _ = v30
        default:
            var v31 uint64 = 1
            _ = v31
        }
        _ = v29
    } else {
        _ = len(v15)
    }
    var v32 uint8 = 31
    for fi33 := 0; fi33 < 11; fi33++ {
        fmt.Println(v1, v1, v21)
        v34 := v5[v2]
        if true {
            v35 := GenS2{F0: (v4 + v34)}
            _ = v35
        } else {
            if (!(0 != v6)) {
                var v36 uint32 = 2147483647
                _ = v36
            }
            v37 := GenS2{F0: v4}
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    var v38 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
