package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32; F1 int8 }
func genFn0() int8 {
    var v4 bool = (!(3.14159 < 3.14159))
    v4 = (66 == 65535)
    v5 := GenS1{F0: -659, F1: -9223372036854775808}
    _ = v4
    _ = v5
    return -128
}

func main() {
    var v1 uint8 = 255
    var v2 int32 = -2147483648
    var v3 int16 = 0
    var v6 float64 = 0.0
    if (1 > 1) {
        if (false || true) {
            switch "hello" {
            case "\n\t":
                for fi7 := 0; fi7 < 8; fi7++ {
                    v3 -= (v3 / ((869) | 1))
                    v8 := GenS1{F0: 191, F1: 0}
                    v9 := GenS1{F0: (v3 | (^(v3))), F1: 1}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case "":
                v1 += v1
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: v3, F1: -9223372036854775808}
                _ = fi10
                _ = v11
            }
            var v12 uint16 = 32767
            _ = v12
        } else {
            switch "hello" {
            case "hello":
                var v13 string = ("foo bar" + "a")
                _ = v13
            default:
                var v14 int16 = 42
                _ = v14
            }
            v15 := GenS2{F0: v2, F1: 1}
            fmt.Println(v6, v1)
            _ = v15
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(v2)
            _ = fi16
        }
    }
    var v17 int16 = 1
    v18 := []int{128, 42, 127, 128}
    for fi19 := 0; fi19 < 1; fi19++ {
        v17 = v17
        v2 -= -1
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = fi19
    }
    v3 = (v3 - 1000)
    var v20 uint = 888
    var v21 int = 2147483647
    v2 = (v2 * v2)
    _ = len(v18)
    var v22 float32 = 1.5
    var v23 float64 = float64(v22)
    v6 = (v23 - v23)
    var v24 float64 = float64(float32((float64(((v22 * ((float32(-2.25) + -2.25) + 0.0)) + (float32(0.0) + 0.1))) / (0.1 * v23))))
    v21 = (v21 + (v21 - 32767))
    if (false && ((^(v1)) != (v1 + v1))) {
        v24 = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v6
}
