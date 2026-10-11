package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint }
func genFn0(p0 int64, p1 uint, p2 uint64) uint8 {
    var v4 bool = (-9223372036854775808 < 1)
    p2 = p2
    fmt.Println(p1, p0)
    p1 -= p1
    var v5 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 128
}

func main() {
    var v1 bool = false
    v2 := []int{-810, 1, 32767}
    v3 := []string{"hello", "foo bar"}
    _ = len(v2)
    switch "a" {
    case "hello":
        switch ("" + ((((("a" + "hello") + ("" + "hello")) + "a") + (("foo bar" + ("a" + "")) + "\n\t")) + "foo bar")) {
        case "":
            var v6 int32 = 42
            _ = v6
        default:
            v7 := GenS1{F0: "a", F1: 733}
            _ = v7
        }
    case "\n\t":
        var v8 uint32 = 101
        _ = v8
    default:
        var v9 uint16 = 1
        _ = v9
    }
    v1 = (!v1)
    if (2147483647 <= 666) {
        var v10 int = 127
        var v11 int64 = 1
        v11 = v11
        _ = v10
        _ = v11
    } else {
        var v12 int64 = -673
        _ = v12
    }
    v13 := GenS2{F0: 0, F1: 1000}
    v14 := GenS1{F0: ("" + ((((("\n\t" + ((((("\n\t" + "foo bar") + "hello") + "hello") + "\n\t") + "a")) + "hello") + "hello") + "\n\t") + ("hello" + "foo bar"))), F1: -2147483648}
    if v1 {
        switch "\n\t" {
        case "hello":
            for fi15 := 0; fi15 < 15; fi15++ {
                if (1 >= -61) {
                    v16 := GenS3{F0: 42, F1: 0, F2: 2}
                    v17 := GenS3{F0: 1, F1: 42, F2: 2147483647}
                    _ = v16
                    _ = v17
                }
                v1 = v1
                _ = fi15
            }
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            if (v1 && (194 != 2)) {
                if ((((((1e10 * 0.0) - float32((-2.25 * -2.25))) / 1.0) + 1.5) + 3.14159) > -2.25) {
                    var v19 uint64 = 2
                    _ = v19
                }
                if false {
                    var v20 uint64 = 757
                    _ = v20
                }
                v1 = (!(42 >= 1))
            }
            if (!((v1 || false) && (true && (!(2147483647 != 0))))) {
                v1 = v1
                v3 = append(v3, (("\n\t" + ("\n\t" + "hello")) + "a"))
                var v21 uint8 = 255
                _ = v21
            }
            var v22 float64 = (0.0 / -2.25)
            _ = fi18
            _ = v22
        }
    }
    var v23 string = "\n\t"
    for fi24 := 0; fi24 < 5; fi24++ {
        v1 = ((v1 && (!(((996 >= 480) || v1) || (!(3.14159 != (3.14159 + (0.0 * (3.14159 / 3.14159)))))))) && (-2.25 > 0.1))
        if ((fi24 & fi24) <= (fi24 / (((fi24 & (fi24 / (((fi24 * (fi24 >> ((42) & 63)))) | 1)))) | 1))) {
            if (!v1) {
                var v25 string = (v23 + ("a" + v23))
                var v26 float64 = 3.14159
                _ = v25
                _ = v26
            }
        }
        _ = fi24
    }
    v27 := GenS4{F0: 1}
    v1 = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v27
    _ = v3
}
