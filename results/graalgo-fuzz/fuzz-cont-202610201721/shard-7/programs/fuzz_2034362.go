package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint64, p1 uint64) int8 {
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(p1, fi6)
        _ = fi6
    }
    fmt.Println(p1, p0)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: p0, F1: 42}
        v9 := []int{-686, fi7, 297}
        if (127 == 212) {
            var v10 string = "a"
            var v11 float32 = (((1e10 - 3.14159) / 0.1) / 1e10)
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 -= p1
    if (false || (-419 >= 2147483647)) {
        p0 -= p0
        v12 := GenS1{F0: 42, F1: -128}
        switch -9223372036854775808 {
        case -2147483648:
            v13 := []int{-750, 32768}
            _ = v13
        default:
            var v14 int8 = 127
            _ = v14
        }
        _ = v12
    }
    _ = p0
    _ = p1
    return 79
}
func genFn1(p0 int8, p1 int16) uint {
    if true {
        switch "" {
        case "foo bar":
            fmt.Println(p0)
        }
        v15 := GenS1{F0: 42, F1: 0}
        _ = v15
    }
    if true {
        fmt.Println(p1, p0)
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := []int{-9223372036854775808, 1, (fi16 % ((42) | 1))}
            _ = fi16
            _ = v17
        }
    }
    v18 := []int{2147483647, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v18
    return 600
}

func main() {
    var v1 uint = 782
    var v2 uint64 = 1
    var v3 uint16 = 42
    v4 := make(map[string]string)
    var v5 string = ("foo bar" + "a")
    v5 = (v5 + ("" + (("" + v5) + (v5 + (v5 + v5)))))
    var v19 uint64 = v2
    v3 -= (^(v3))
    v3 = uint16(v1)
    v4[("\n\t" + v5)] = "hello"
    v20 := GenS1{F0: v2, F1: 100}
    var v21 uint64 = v2
    switch 9223372036854775807 {
    case 4294967295:
        v21 += (^(v2))
    default:
        if (344 == -575) {
            var v22 float64 = (-2.25 + 0.1)
            _ = v22
        }
    }
    _ = len(v4)
    var v23 uint32 = 2147483648
    if (((^(v1)) == (^(v1))) || ("hello" <= v5)) {
        var v24 int32 = 0
        _ = v24
    } else {
        v2 = 784
        v19 += v2
        if true {
            v21 += v21
            v4[v5] = "foo bar"
        }
    }
    if (!false) {
        v25 := GenS1{F0: (v21 / ((0) | 1)), F1: 0}
        _ = v25
    }
    fmt.Println(v21)
    var v26 int16 = 557
    v1 = (^(v1))
    if ((42 > v19) && (v1 < (v1 + v1))) {
        if (!(0.1 != float32(((0.0 / -2.25) / 1.5)))) {
            v4[v5] = "foo bar"
        }
        v21 += 410
        var v27 uint32 = (v23 * (v23 + (v23 + v23)))
        _ = v27
    } else {
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 15; fi29++ {
                var v30 int32 = 42
                var v31 uint32 = (^(v23))
                v21 = 0
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 uint = (v1 % (((^(v1))) | 1))
            switch 106 {
            case 730:
                var v33 int8 = genFn0((v2 / (((v21 + 0)) | 1)), 190)
                _ = v33
            case -699:
                var v34 uint = (v32 ^ genFn1(42, v26))
                _ = v34
            }
            _ = fi28
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
