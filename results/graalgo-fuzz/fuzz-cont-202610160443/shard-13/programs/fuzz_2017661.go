package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
func genFn0(p0 uint64, p1 bool, p2 int8) uint16 {
    v4 := GenS1{F0: (!(p1 || (0 > 661)))}
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint16 = 2
            var v8 uint8 = 128
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint16 = 1
        _ = fi5
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 798
}
func genFn1(p0 int, p1 uint8) int8 {
    v10 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v10
    return 100
}

func main() {
    v1 := []string{(((("" + "") + ("a" + "foo bar")) + "a") + ("a" + "foo bar")), "", ("\n\t" + "a")}
    var v2 uint32 = 2147483647
    var v3 uint = 674
    var v11 int = -921
    switch ("foo bar" + (("" + ("" + "\n\t")) + ("foo bar" + ""))) {
    case "\n\t":
        v3 += v3
    case "a":
        var v12 uint32 = v2
        _ = v12
    case "hello":
        fmt.Println(v11, v2, v3)
    }
    v13 := GenS1{F0: false}
    v14 := GenS1{F0: (42 >= 624)}
    fmt.Println(v11)
    v11 -= (v11 & (v11 - (v11 * v11)))
    var v15 int32 = -61
    fmt.Println(v15, v3)
    v11 += v11
    v2 = v2
    if (((0.1 + float32(-2.25)) == 0.1) && ((true && false) && true)) {
        if (false || (false || ((!true) || false))) {
            var v16 int8 = 42
            v16 -= v16
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS2{F0: (-(fi17))}
                _ = fi17
                _ = v18
            }
            _ = v16
        } else {
            fmt.Println(v2)
        }
        var v19 int32 = v15
        v20 := GenS1{F0: true}
        _ = v19
        _ = v20
    } else {
        if ((float32(1.5) != 0.0) && (!((0.0 * 3.14159) <= (1e10 * float32(float64(0.1)))))) {
            var v21 bool = (true && ((float32(1.5) + float32(1.5)) != 0.0))
            v15 = 42
            var v22 uint32 = 2
            _ = v21
            _ = v22
        }
        v11 = (v11 / (((v11 / (((v11 >> ((771) & 63))) | 1))) | 1))
        var v23 int32 = (v15 * v15)
        _ = v23
    }
    if ((^(v15)) < (v15 + v15)) {
        var v24 uint = (v3 - v3)
        _ = v24
    } else {
        var v25 bool = (!false)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v26 uint8 = 150
        _ = v25
        _ = v26
    }
    var v27 uint64 = 632
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
}
