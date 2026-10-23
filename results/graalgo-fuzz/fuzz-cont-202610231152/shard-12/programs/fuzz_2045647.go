package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 string }

func main() {
    v1 := []int{4294967295, 32768, 2147483647, -9223372036854775808}
    var v2 int8 = 84
    var v3 float32 = (((float32(0.1) + 3.14159) + 0.1) / 1e10)
    var v4 int = -826
    var v5 int8 = (v2 % (((v2 / (((v2 | (^(v2)))) | 1))) | 1))
    var v6 int32 = -2147483648
    v3 -= v3
    v6 = (v6 - (v6 % ((v6) | 1)))
    var v7 uint32 = 1
    var v8 uint16 = 0
    v9 := GenS1{F0: ((v2 * v2) >= (v2 % ((v2) | 1)))}
    for fi10 := 0; fi10 < 3; fi10++ {
        if (!(v7 <= 42)) {
            v11 := GenS1{F0: (!((true && true) && (-1 > -9223372036854775808)))}
            v12 := GenS1{F0: (false && true)}
            var v13 string = ((("a" + ("\n\t" + "\n\t")) + ("" + ("a" + ("foo bar" + "\n\t")))) + ("" + ("foo bar" + (("\n\t" + "a") + (((("hello" + ("\n\t" + "a")) + ((("hello" + "foo bar") + "") + (("\n\t" + ("foo bar" + "")) + ("" + ("hello" + ""))))) + ("foo bar" + "hello")) + "a")))))
            _ = v11
            _ = v12
            _ = v13
        } else {
            if (!true) {
                var v14 bool = (false || (v7 == (v7 + v7)))
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            v5 += int8(v6)
            switch (fi15 + (v4 ^ v4)) {
            case 127:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v16 int16 = 831
            _ = fi15
            _ = v16
        }
        _ = fi10
    }
    if true {
        fmt.Println(v2, v8, v7)
        v17 := GenS2{F0: 403, F1: "a"}
        _ = v17
    } else {
        var v18 uint64 = 1
        var v19 uint32 = (v7 * (v7 % (((v7 * 1)) | 1)))
        var v20 int = v4
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: ((891 != v8) || ((((0.0 * ((1.5 - ((-2.25 / (((-2.25 / (-2.25 - 0.0)) + 0.1) - 1e10)) + (1e10 - 3.14159))) + (0.1 / ((1.5 / 1.5) - float64((float32((float64((0.0 * 1.5)) / (0.0 + 3.14159))) / v3)))))) <= -2.25) && ((2 < 128) || false)) && true))}
    var v22 float32 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
