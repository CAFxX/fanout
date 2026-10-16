package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) int32 {
    p0 = p0
    var v6 bool = (-2.25 != ((3.14159 * (3.14159 * (1e10 * 1e10))) * (0.0 / 3.14159)))
    _ = p0
    _ = v6
    return 903
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    v4 := []string{"a", (("\n\t" + ("foo bar" + (("a" + "foo bar") + "\n\t"))) + "a")}
    v5 := make(map[int]int)
    switch -574 {
    case 4294967295:
        v7 := GenS1{F0: 864, F1: 255}
        _ = v7
    case 0:
        v8 := GenS2{F0: v3}
        _ = v8
    }
    fmt.Println(v1, v3)
    v9 := GenS2{F0: (-2147483648 > 2147483647)}
    v10 := GenS2{F0: (66 != -103)}
    v11 := GenS1{F0: 683, F1: 42}
    var v12 uint8 = 42
    _ = len(v2)
    v3 = (42 == 18446744073709551615)
    fmt.Println(v12, v12, v12)
    var v13 float64 = (1.5 + float64(1e10))
    if false {
        switch "hello" {
        case "foo bar":
            v3 = true
        case "\n\t":
            v14 := GenS2{F0: (v3 && false)}
            _ = v14
        case "a":
            v13 = ((v1 - v1) - (v1 * (float64(((1e10 / (float32((float64(0.1) * (v1 + (0.1 + 1e10)))) * (1.5 / 1.0))) + float32((v1 * -2.25)))) / (v1 + v13))))
        default:
            _ = len(v4)
        }
    }
    var v15 uint16 = 65535
    if ((3.14159 > 1e10) && (42 != 1)) {
        switch 1 {
        case 773:
            _ = len(v4)
        default:
            v4 = append(v4, ((("foo bar" + "\n\t") + "hello") + ""))
        }
    }
    var v16 float64 = 0.0
    switch 315 {
    case -790:
        v17 := GenS2{F0: true}
        _ = v17
    }
    var v18 uint8 = v12
    for fi19 := 0; fi19 < 2; fi19++ {
        if v3 {
            v12 = v12
            v15 = v15
            for fi20 := 0; fi20 < 20; fi20++ {
                if ((!v3) && (v3 || v3)) {
                    fmt.Println(v1, v15, fi19)
                    var v21 uint = 2
                    _ = v21
                } else {
                    fmt.Println(v1, v12, fi20)
                    v18 += 212
                }
                switch 194 {
                case 32767:
                    _ = len(v4)
                case 26:
                    switch "a" {
                    case "\n\t":
                        _ = len(v2)
                    case "a":
                        var v22 int16 = -678
                        _ = v22
                    default:
                        v13 = float64(0.1)
                    }
                default:
                    fmt.Println(v16, v18)
                }
                switch "\n\t" {
                case "foo bar":
                    var v23 int32 = 0
                    _ = v23
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                default:
                    v13 = v13
                }
                _ = fi20
            }
        }
        var v24 uint = 255
        _ = fi19
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
