package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    v1 := []uint32{2147483648}
    var v2 uint8 = 1
    var v3 uint32 = 76
    fmt.Println(v3)
    var v4 float32 = (1.5 * (1.5 - 3.14159))
    v2 = (v2 ^ v2)
    fmt.Println(v3, v2)
    v5 := GenS1{F0: 0}
    if false {
        var v6 float32 = v4
        if true {
            if ((!(false && false)) || false) {
                for fi7 := 0; fi7 < 14; fi7++ {
                    v6 = (v4 + float32(0.0))
                    if len(v1) > 0 {
                        v1[0] = (^(v3))
                    }
                    var v8 string = ("\n\t" + (("hello" + ("" + "")) + ""))
                    _ = fi7
                    _ = v8
                }
                v9 := GenS2{F0: -2147483648, F1: 42}
                var v10 int8 = -23
                _ = v10
                _ = v9
            } else {
                v2 = (^(v2))
                fmt.Println(v6, v3)
            }
            var v11 int8 = -116
            var v12 float32 = v6
            _ = v11
            _ = v12
        } else {
            _ = len(v1)
            if (!((!false) && (-73 != 70))) {
                v1 = append(v1, v3)
                switch 42 {
                case 672:
                    if true {
                        var v13 bool = true
                        _ = v13
                    }
                case 128:
                    v4 = 3.14159
                case 1:
                    switch 9223372036854775807 {
                    case -345:
                        for fi14 := 0; fi14 < 3; fi14++ {
                            var v15 uint16 = uint16(v2)
                            _ = fi14
                            _ = v15
                        }
                    default:
                        v3 -= v3
                    }
                }
            } else {
                var v16 float32 = v4
                v2 -= 255
                _ = v16
            }
        }
        switch 53 {
        case 9223372036854775807:
            v1 = append(v1, (v3 ^ (^(v3))))
        case 32767:
            v1 = append(v1, v3)
        }
        _ = v6
    }
    if (100 <= -20) {
        var v17 string = ("foo bar" + "hello")
        var v18 float32 = (v4 / v4)
        _ = v17
        _ = v18
    }
    var v19 int8 = -128
    var v20 int16 = -32768
    var v21 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
