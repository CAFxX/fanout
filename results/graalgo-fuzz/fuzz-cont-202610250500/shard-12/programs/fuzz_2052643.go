package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int32 = 42
    v2 := []string{"hello"}
    var v3 uint16 = 222
    v1 = v1
    if (62 > 0) {
        v3 = 853
    } else {
        var v4 float32 = (3.14159 + (((0.1 * (-2.25 + 1.5)) + 1e10) * (0.0 - 1e10)))
        v4 -= (v4 - ((((v4 * ((((-2.25 / float32(0.1)) + v4) * ((float32((((1e10 - 1.5) + 0.0) / -2.25)) / v4) * (3.14159 / float32(0.1)))) / 3.14159)) + 1.5) / ((float32((float64(v4) + 1.5)) / 0.1) + (v4 + v4))) + 0.1))
        if true {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            v3 += v3
        } else {
            var v5 int8 = 42
            _ = v5
        }
        _ = v4
    }
    v1 -= (v1 & v1)
    fmt.Println(v1, v3)
    if false {
        if len(v2) > 0 {
            v2[0] = ""
        }
    } else {
        for fi6 := 0; fi6 < 1; fi6++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch int(v1) {
            case 127:
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            }
            if len(v2) > 0 {
                v2[0] = ""
            }
            _ = fi6
        }
        v7 := GenS1{F0: 2, F1: v3}
        v1 -= -2147483648
        _ = v7
    }
    var v8 int16 = 0
    var v9 string = ""
    if (true || (1.5 <= 0.1)) {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v10 := GenS1{F0: 61, F1: 1}
            _ = v10
        }
        v3 = v3
    }
    fmt.Println(v3, v8, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
