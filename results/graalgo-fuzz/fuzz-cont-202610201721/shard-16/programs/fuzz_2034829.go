package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8, p1 int16, p2 int8) string {
    p2 -= -128
    switch ("a" + "") {
    case "\n\t":
        for fi4 := 0; fi4 < 2; fi4++ {
            p2 = 42
            _ = fi4
        }
    case "hello":
        switch 32768 {
        case 1:
            switch ((("\n\t" + ("hello" + ("" + ("\n\t" + ("a" + ("\n\t" + "hello")))))) + ("" + ("a" + "foo bar"))) + "\n\t") {
            case "hello":
                p1 = (p1 ^ (p1 & (p1 / (((^(p1))) | 1))))
            default:
                if ((p2 | p2) < (p0 % ((p2) | 1))) {
                    p0 = (-(p0))
                }
            }
        default:
            if false {
                v5 := GenS1{F0: -9223372036854775808, F1: 32767}
                var v6 int32 = 364
                _ = v5
                _ = v6
            }
        }
    }
    v7 := []int{32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "a"
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    var v3 uint32 = 2147483647
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint8 = 0
        var v10 int8 = -111
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 0}
    delete(v2, "a")
    fmt.Println(v1, v1)
    var v12 int64 = 42
    fmt.Println(v12, v3, v12)
    var v13 int64 = v12
    v14 := []int{}
    if v1 {
        v13 += (v13 + v12)
    } else {
        if true {
            v15 := GenS2{F0: 106}
            v14 = append(v14, 4294967295)
            _ = v15
        } else {
            var v16 uint8 = 0
            v14 = append(v14, 32768)
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS1{F0: v13, F1: (fi17 * fi17)}
                fmt.Println(v13)
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    }
    v14 = append(v14, -9223372036854775808)
    var v19 float32 = float32(1.5)
    v19 -= v19
    v13 = 9223372036854775807
    v20 := GenS1{F0: (v12 >> ((636) & 63)), F1: 127}
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
