package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0() uint16 {
    var v5 int16 = 0
    var v6 float64 = 0.1
    fmt.Println(v6)
    if true {
        v6 -= 0.1
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := []int{(fi7 >> ((65535) & 63))}
            v8 = append(v8, fi7)
            v9 := GenS1{F0: 124, F1: 324}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{1}
        _ = v10
    } else {
        fmt.Println(v6)
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := []int{fi11, fi11, (fi11 % (((fi11 + (-(fi11)))) | 1))}
            switch "" {
            case "a":
                fmt.Println(v6)
            case "hello":
                if (!(0 < 721)) {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 int16 = (v5 % ((v5) | 1))
                        _ = fi13
                        _ = v14
                    }
                }
            }
            _ = fi11
            _ = v12
        }
        v15 := []int{0, 32768, 1}
        _ = v15
    }
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 uint = 1
    var v2 float64 = 3.14159
    var v3 int8 = 0
    v4 := []string{"foo bar", "\n\t"}
    v2 += ((1e10 - (v2 + 1e10)) - 3.14159)
    v1 -= 0
    v16 := GenS1{F0: 100, F1: 65535}
    v2 = (((-2.25 / 0.1) / ((v2 + v2) / (((v2 - (v2 + 0.1)) - v2) + 3.14159))) * (3.14159 - 3.14159))
    v1 -= (v1 & (^(v1)))
    if (!(42 > 309)) {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int32 = -1
            var v19 uint64 = 18446744073709551615
            _ = fi17
            _ = v18
            _ = v19
        }
        v1 = v1
    }
    v1 -= 1
    if (((-593 != -2147483648) && true) || false) {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    }
    var v20 string = ""
    if (-134 != 438) {
        v20 = ""
    } else {
        v20 = (v20 + "\n\t")
        var v21 int8 = v3
        var v22 int8 = 42
        _ = v21
        _ = v22
    }
    var v23 float64 = 3.14159
    _ = len(v4)
    var v24 bool = ((65535 < 87) || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
