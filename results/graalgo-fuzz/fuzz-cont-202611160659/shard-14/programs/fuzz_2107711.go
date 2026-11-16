package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0() int {
    v6 := GenS1{F0: 0, F1: 0, F2: (-1 == -2147483648)}
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{(fi7 | (fi7 * (fi7 + fi7))), fi7, (fi7 % (((fi7 ^ (-(fi7)))) | 1)), fi7}
            if len(v8) > 0 {
                v8[0] = 42
            }
            switch "foo bar" {
            case "a":
                var v9 float64 = 1.5
                _ = v9
            case "":
                if true {
                    fmt.Println(fi7)
                    fmt.Println(fi7)
                } else {
                    var v10 float32 = 0.1
                    v10 = (1e10 * v10)
                    _ = v10
                }
            }
            _ = fi7
            _ = v8
        }
    }
    v11 := []int{128, 32767, -2147483648, 128}
    var v12 int32 = -2147483648
    _ = v11
    _ = v12
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 uint = 89
    var v2 uint = 18446744073709551615
    var v3 bool = (("a" < "foo bar") && true)
    var v4 float32 = 1e10
    v5 := make(map[string]bool)
    var v13 uint16 = 65535
    v14 := GenS2{F0: -1}
    v2 = 647
    if ((!(v3 && v3)) && v3) {
        v1 = v2
    } else {
        var v15 uint8 = 17
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        if (1 == 1) {
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 uint32 = 2147483647
                switch "" {
                case "a":
                    var v19 int32 = -625
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            var v20 int8 = 1
            v21 := v5[("foo bar" + "")]
            _ = v20
            _ = v21
        } else {
            fmt.Println(v3, v3, v1)
        }
        _ = fi16
    }
    v1 = (^(v1))
    v22 := v5[""]
    v1 = (v2 / (((v1 * v2)) | 1))
    if v22 {
        v3 = true
        fmt.Println(v3, v22, v3)
        v13 -= v13
    }
    var v23 float64 = (0.1 + float64(v4))
    if ((549 >= v13) && v3) {
        var v24 int16 = 0
        _ = v24
    }
    fmt.Println(v13, v4)
    var v25 uint = (v1 + 2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
