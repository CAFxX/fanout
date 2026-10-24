package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 bool, p1 uint32, p2 uint8) bool {
    p1 = (^(p1))
    v4 := []int{4294967295}
    p0 = p0
    if len(v4) > 0 {
        v4[0] = 791
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1(p0 string) string {
    p0 = "\n\t"
    switch 674 {
    case 127:
        var v5 int8 = 1
        _ = v5
    }
    v6 := []int{820, -557, -2147483648, 4294967295}
    _ = p0
    _ = v6
    return p0
}

func main() {
    var v1 float64 = float64(((float32((0.0 + ((0.0 * (-2.25 / 1.5)) - -2.25))) * (-2.25 / 3.14159)) * float32((0.0 - 1e10))))
    var v2 float64 = 3.14159
    var v3 int64 = -1
    v7 := GenS1{F0: 255, F1: -2147483648}
    v8 := GenS1{F0: 65535, F1: -1}
    var v9 string = "hello"
    var v10 string = v9
    v9 = ""
    v10 = v10
    v11 := []int{919, 9223372036854775807, 0, -2147483648}
    if false {
        if (-2147483648 > -130) {
            v9 = "foo bar"
            v3 = (-(v3))
        } else {
            v11 = append(v11, -371)
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    } else {
        var v12 float32 = 1.5
        _ = v12
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println(v3, v3, v9)
    if false {
        if (4294967295 != 0) {
            v2 -= v2
            v13 := GenS2{F0: ((-111 >= -128) && true), F1: 38}
            _ = v13
        } else {
            v10 = "foo bar"
        }
        var v14 float64 = v1
        for fi15 := 0; fi15 < 5; fi15++ {
            v11 = append(v11, fi15)
            _ = fi15
        }
        _ = v14
    } else {
        if ((((-431 <= 2147483647) || (false && (!true))) && false) || (1 > 89)) {
            for fi16 := 0; fi16 < 9; fi16++ {
                v2 = v2
                _ = len(v11)
                _ = fi16
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            v1 += (v1 + 3.14159)
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(fi17, v3, v1)
                v10 = v10
                v18 := GenS1{F0: 14, F1: 42}
                _ = fi17
                _ = v18
            }
            _ = len(v11)
        }
    }
    var v19 int = -855
    v19 = v19
    v20 := GenS1{F0: 214, F1: 761}
    fmt.Println(v1, v9, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
