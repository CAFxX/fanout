package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int32, p1 int, p2 bool) uint16 {
    var v4 int32 = (p0 % ((0) | 1))
    var v5 uint = 831
    fmt.Println(v5, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 65535
}
func genFn1(p0 int32) uint8 {
    p0 = -255
    p0 -= -228
    if (-1 == 1) {
        var v6 int32 = -929
        v7 := []int{397}
        _ = v6
        _ = v7
    }
    switch 128 {
    case 0:
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 -= p0
            _ = fi8
        }
    }
    _ = p0
    return 127
}

func main() {
    var v1 uint8 = 146
    var v2 int16 = 42
    var v3 uint8 = (v1 + v1)
    v9 := GenS1{F0: -368}
    var v10 uint64 = 0
    var v11 string = ("" + "\n\t")
    var v12 uint32 = 2147483647
    v1 = 40
    v3 += uint8(v12)
    switch (("a" + ("" + "hello")) + v11) {
    case "hello":
        fmt.Println(v10)
    default:
        v13 := GenS1{F0: -9223372036854775808}
        _ = v13
    }
    var v14 uint32 = (v12 + 2147483648)
    v14 = uint32(v3)
    var v15 bool = (!true)
    v16 := []int{4294967295, -2147483648, 2147483647}
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 uint = 484
        for fi19 := 0; fi19 < 15; fi19++ {
            if len(v16) > 0 {
                v16[0] = 32768
            }
            _ = fi19
        }
        if v15 {
            if ((fi17 >> (((v18 & 398)) & 63)) > -517) {
                var v20 bool = v15
                _ = v20
            } else {
                if (!(v15 || ((v15 && (float32(0.1) <= 1e10)) || ((3.14159 + float64((1e10 - ((1e10 / (0.1 + -2.25)) / 0.1)))) <= 1.5)))) {
                    v15 = v15
                    v11 = (v11 + v11)
                }
                v18 = 18446744073709551615
            }
        }
        _ = fi17
        _ = v18
    }
    v1 += (v3 / ((1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
