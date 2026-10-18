package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint8) int64 {
    p0 -= 2
    v4 := []int{-87, 9223372036854775807, 2147483647}
    _ = p0
    _ = v4
    return -9223372036854775808
}
func genFn1() int64 {
    var v5 string = ("foo bar" + "\n\t")
    v5 = v5
    _ = v5
    return -260
}

func main() {
    var v1 string = (((("a" + ("a" + "foo bar")) + (((("\n\t" + "\n\t") + "hello") + ("foo bar" + "a")) + "a")) + (("hello" + "hello") + "foo bar")) + "foo bar")
    var v2 string = v1
    var v3 float32 = float32(0.0)
    var v6 int32 = -1
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 0, F1: 165}
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v6, v3)
    if true {
        switch v1 {
        case "foo bar":
            v9 := []int{128, 32767}
            _ = v9
        case "":
            v3 = (((0.1 + 1.5) / ((3.14159 / float32(-2.25)) - v3)) * v3)
        }
    } else {
        if (42 <= 364) {
            var v10 int16 = -1
            v6 -= (v6 * 921)
            switch -672 {
            case 9223372036854775807:
                switch v1 {
                case "foo bar":
                    fmt.Println(v10, v3)
                case "":
                    v11 := []int{-1}
                    _ = v11
                default:
                    var v12 uint16 = 2
                    _ = v12
                }
            case 1:
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 int8 = 1
                    _ = fi13
                    _ = v14
                }
            case -1:
                if true {
                    v1 = (v2 + ("a" + ""))
                    v3 += (v3 + (float32(1.5) - v3))
                    v15 := GenS1{F0: 170, F1: 247}
                    _ = v15
                }
            }
            _ = v10
        } else {
            for fi16 := 0; fi16 < 9; fi16++ {
                if false {
                    v3 -= (3.14159 - v3)
                }
                _ = fi16
            }
            var v17 float64 = float64((float32(3.14159) / v3))
            _ = v17
        }
        v18 := []int{-2147483648, -243}
        v2 = ((v2 + "\n\t") + ((("hello" + v1) + v2) + ((v1 + v2) + "hello")))
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 bool = (false && true)
        _ = fi19
        _ = v20
    }
    switch v1 {
    case "a":
        v21 := GenS1{F0: 127, F1: 127}
        _ = v21
    default:
        if (2147483647 == -1) {
            for fi22 := 0; fi22 < 7; fi22++ {
                v1 = v2
                _ = fi22
            }
        }
    }
    var v23 float64 = 1e10
    fmt.Println(v1)
    var v24 uint8 = 127
    v25 := GenS1{F0: 0, F1: 214}
    var v26 uint8 = v24
    v26 = 255
    for fi27 := 0; fi27 < 7; fi27++ {
        fmt.Println(v6)
        _ = fi27
    }
    v28 := []int{41, -2147483648, 963}
    v29 := GenS1{F0: v26, F1: (v26 / (((v26 ^ (v26 / ((127) | 1)))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v6
}
