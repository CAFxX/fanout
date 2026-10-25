package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint16; F2 int }
func genFn0(p0 int32, p1 uint, p2 bool) uint {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint = 1
        if p2 {
            for fi7 := 0; fi7 < 12; fi7++ {
                fmt.Println(p1, fi5)
                _ = fi7
            }
        } else {
            v6 -= p1
            if (-280 > -9223372036854775808) {
                if (65535 > 42) {
                    fmt.Println(fi5)
                }
                v8 := []int{4294967295, 359}
                var v9 uint8 = 0
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: -889, F1: 65535}
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    var v11 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p1
}
func genFn1() uint16 {
    v12 := GenS1{F0: 0, F1: 4294967295}
    v13 := GenS1{F0: 0, F1: 660}
    var v14 int32 = 690
    _ = v12
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = -2.25
    var v3 bool = (v2 <= (3.14159 - ((float32(((1e10 - (1.5 / 1.5)) / float64((-2.25 * (1.5 / ((0.0 + v2) - v2)))))) - float32(float64(v2))) + 1e10)))
    var v4 uint = 2
    v15 := []int{595, -9223372036854775808, -2147483648}
    if (!(v4 == 42)) {
        var v16 uint32 = 42
        switch int(v1) {
        case 32768:
            v17 := GenS1{F0: v1, F1: v16}
            _ = v17
        case 0:
            v16 += v16
        case 2147483647:
            v15 = append(v15, 32768)
        }
        _ = v16
    } else {
        var v18 bool = (!(647 == -2147483648))
        v2 = v2
        v2 -= v2
        _ = v18
    }
    v15 = append(v15, 32767)
    if v3 {
        v4 = v4
        if true {
            var v19 float32 = v2
            v20 := GenS2{F0: (^(v4)), F1: 1, F2: -714}
            var v21 string = (("\n\t" + "") + "\n\t")
            _ = v19
            _ = v20
            _ = v21
        }
    }
    var v22 bool = v3
    v1 = (v1 & v1)
    v22 = v22
    switch "foo bar" {
    case "hello":
        var v23 string = ("" + "a")
        _ = v23
    case "\n\t":
        var v24 uint32 = 65535
        _ = v24
    }
    var v25 float32 = (1e10 * (v2 / v2))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
