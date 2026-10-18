package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0(p0 int64, p1 int64, p2 uint32) int64 {
    var v6 int16 = 169
    v7 := GenS1{F0: 1, F1: ("a" != ("foo bar" + ("\n\t" + "a")))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 / (((-(p1))) | 1))
}

func main() {
    var v1 string = "a"
    v2 := []string{}
    var v3 int64 = 595
    v4 := []int{9223372036854775807}
    var v5 bool = false
    v5 = v5
    var v8 float64 = (1.5 + float64(0.1))
    if v5 {
        v8 += 0.0
    }
    var v9 bool = ("\n\t" <= v1)
    switch "a" {
    case "a":
        v4 = append(v4, 719)
    case "foo bar":
        fmt.Println(v5)
    default:
        v5 = v9
    }
    var v10 uint8 = 128
    switch -769 {
    case 32767:
        var v11 int64 = (v3 * (v3 | v3))
        _ = v11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "hello":
        v12 := GenS1{F0: 42, F1: false}
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 8; fi13++ {
            switch 0 {
            case 128:
                v10 = (v10 * v10)
            case 301:
                if len(v4) > 0 {
                    v4[0] = -2147483648
                }
            case 0:
                var v14 uint = 255
                _ = v14
            default:
                if v9 {
                    v15 := GenS1{F0: 733, F1: true}
                    var v16 uint16 = 65
                    if (!v9) {
                        if (2147483648 == 0) {
                            var v17 uint32 = 4294967295
                            if v5 {
                                v9 = (v10 != (v10 % ((v10) | 1)))
                            }
                            var v18 int16 = -1
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = v15
                    _ = v16
                } else {
                    v4 = append(v4, fi13)
                }
            }
            _ = fi13
        }
    default:
        fmt.Println(v3, v1)
    }
    switch 4294967295 {
    case 1:
        if len(v4) > 0 {
            v4[0] = 127
        }
    default:
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS1{F0: 255, F1: v9}
            _ = fi19
            _ = v20
        }
    }
    switch 1 {
    case 0:
        v8 = (v8 / v8)
    case 4294967295:
        v8 -= 3.14159
    }
    v21 := GenS1{F0: 1, F1: v9}
    v9 = false
    v9 = (-2.25 < 0.1)
    v1 = v1
    var v22 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
