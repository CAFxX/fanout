package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8) int8 {
    var v4 int16 = -413
    v5 := GenS1{F0: ("foo bar" + "foo bar")}
    _ = p0
    _ = v4
    _ = v5
    return -6
}

func main() {
    v1 := []int{}
    var v2 uint32 = 65535
    var v3 float32 = (3.14159 / 0.1)
    var v6 uint32 = v2
    v6 += 361
    v7 := GenS1{F0: ("" + "hello")}
    v3 -= 0.0
    fmt.Println(v2, v6)
    v8 := GenS1{F0: ("foo bar" + "")}
    v9 := GenS1{F0: "hello"}
    v1 = append(v1, 42)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v3 = v3
        _ = len(v1)
        v10 := GenS2{F0: 30}
        _ = v10
    }
    if (773 <= 2) {
        v6 += 4294967295
        fmt.Println(v6)
        var v11 uint8 = 19
        _ = v11
    }
    fmt.Println(v6, v2, v2)
    var v12 float64 = 1e10
    switch ("a" + ("hello" + "")) {
    case "foo bar":
        switch 42 {
        case 32768:
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(v3)
                if (!(1 == -128)) {
                    switch -9223372036854775808 {
                    case 477:
                        var v14 int16 = 63
                        _ = v14
                    case 127:
                        var v15 uint = uint(v2)
                        _ = v15
                    case 9223372036854775807:
                        var v16 float64 = (0.1 / v12)
                        _ = v16
                    default:
                        var v17 float64 = (1.5 + v12)
                        _ = v17
                    }
                    if (!((("hello" + "foo bar") + "foo bar") > "a")) {
                        var v18 uint16 = 255
                        if len(v1) > 0 {
                            v1[0] = fi13
                        }
                        _ = v18
                    }
                }
                v12 -= v12
                _ = fi13
            }
        case 42:
            v2 += 65535
        }
    case "\n\t":
        var v19 float32 = ((v3 - -2.25) + 3.14159)
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint16 = 65535
        v22 := GenS2{F0: (v21 % ((v21) | 1))}
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v2, v3, v12)
    v12 += v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
