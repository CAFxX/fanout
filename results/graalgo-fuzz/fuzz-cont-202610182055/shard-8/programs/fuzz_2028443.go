package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0(p0 int8, p1 int32, p2 uint64) int64 {
    p2 += (p2 ^ (p2 << ((1) & 63)))
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1() uint8 {
    var v8 float64 = ((3.14159 / 0.1) - 0.1)
    for fi9 := 0; fi9 < 2; fi9++ {
        v8 = ((0.0 + v8) - ((float64(-2.25) / v8) / -2.25))
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{232, 32768, (fi10 << ((18446744073709551615) & 63)), -1}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 int32 = 481
    v13 := []int{-350, 42}
    _ = v12
    _ = v13
    _ = v8
    return 0
}

func main() {
    var v1 uint16 = 626
    var v2 int = -9223372036854775808
    var v3 int8 = 1
    var v4 uint16 = v1
    var v5 int8 = v3
    var v6 int32 = 42
    var v14 float32 = (-2.25 - (((((1.5 * 0.0) + 3.14159) + (1.5 + 1e10)) / (-2.25 + 1e10)) / (0.1 + -2.25)))
    v15 := GenS1{F0: (v2 + (-(v2))), F1: 342}
    fmt.Println(v2, v14)
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            if true {
                v4 -= (^(v1))
            }
            v14 -= 0.0
            fmt.Println(v3, v1)
            _ = fi17
        }
        _ = fi16
    }
    var v18 uint64 = 298
    var v19 bool = false
    v3 -= (v3 & (v3 % (((v5 / ((v5) | 1))) | 1)))
    v20 := GenS1{F0: 1, F1: 1}
    var v21 bool = v19
    if (1e10 <= v14) {
        v22 := []int{805, v2, (v2 + -94)}
        switch (("a" + "\n\t") + "") {
        case "foo bar":
            v21 = (v19 || true)
        case "hello":
            _ = len(v22)
        }
        switch v2 {
        case 602:
            v22 = append(v22, 0)
        case 2147483647:
            v2 = v2
        case 127:
            var v23 int = -2147483648
            _ = v23
        default:
            var v24 uint16 = (v1 / (((v4 % (((v4 * (^(v4)))) | 1))) | 1))
            _ = v24
        }
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
