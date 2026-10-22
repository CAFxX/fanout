package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
func genFn0() uint {
    if (-1 == -32768) {
        var v4 bool = (!(0 > -2147483648))
        _ = v4
    }
    v5 := []int{-1, 32768}
    var v6 bool = false
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int16 = 93
    var v2 float32 = 0.1
    var v3 uint = 2
    v1 = 0
    switch -1 {
    case 0:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float64 = (1.5 * (((-2.25 * float64(0.0)) / 0.1) * (3.14159 + 0.0)))
            v9 := []int{fi7, 127}
            v3 = genFn0()
            _ = fi7
            _ = v8
            _ = v9
        }
    case 1:
        v10 := GenS1{F0: -66, F1: -718, F2: (("foo bar" + "foo bar") + "")}
        _ = v10
    }
    var v11 float64 = 1.5
    v12 := GenS1{F0: 127, F1: -502, F2: "a"}
    fmt.Println(v1, v3)
    v1 -= (v1 / (((v1 * -1)) | 1))
    var v13 uint16 = 2
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint8 = 0
        _ = fi14
        _ = v15
    }
    fmt.Println(v13, v3, v13)
    for fi16 := 0; fi16 < 12; fi16++ {
        v13 += v13
        for fi17 := 0; fi17 < 14; fi17++ {
            if true {
                if (42 > 254) {
                    fmt.Println(fi17, v1)
                    v13 = (v13 & 32767)
                }
                var v18 bool = false
                _ = v18
            }
            var v19 bool = false
            v2 = 0.1
            _ = fi17
            _ = v19
        }
        _ = fi16
    }
    fmt.Println(v11, v2)
    if (448 >= 58) {
        v20 := GenS1{F0: 51, F1: -233, F2: "a"}
        var v21 int64 = -720
        v3 += v3
        _ = v20
        _ = v21
    }
    v22 := []int{int(v13), int(v3), -219}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
}
