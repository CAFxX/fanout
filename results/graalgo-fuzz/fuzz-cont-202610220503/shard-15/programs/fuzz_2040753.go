package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{892}
    v3 := []string{""}
    var v4 int8 = -61
    v1 -= (float64(float32(v1)) - (v1 + float64(float32((v1 - float64((float32((v1 / ((v1 / ((float64(float32(((v1 / v1) + float64((((((1.5 * 1e10) + (1e10 + 0.1)) - 3.14159) + 0.1) + ((-2.25 - (0.1 + ((0.1 - (0.1 * 1.5)) - -2.25))) + 1.5)))))) + 1.5) - (v1 - float64(-2.25)))) / ((-2.25 + v1) * v1)))) + 0.0)))))))
    var v5 float32 = 3.14159
    var v6 bool = (!true)
    var v7 float32 = v5
    if len(v2) > 0 {
        v2[0] = 42
    }
    if len(v3) > 0 {
        v3[0] = ("a" + "")
    }
    v8 := GenS1{F0: false, F1: ("hello" + ""), F2: 659}
    var v9 uint32 = 4294967295
    switch 390 {
    case 118:
        var v10 string = "\n\t"
        _ = v10
    }
    switch "a" {
    case "foo bar":
        var v11 uint16 = 366
        _ = v11
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        v9 += 65535
    default:
        v4 -= v4
    }
    if true {
        fmt.Println(v7, v6, v6)
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        if (v6 && false) {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            fmt.Println(v9, v9)
        } else {
            v13 := GenS1{F0: (v6 || (v6 || v6)), F1: "a", F2: 18446744073709551615}
            if true {
                fmt.Println(v1, v6)
            } else {
                for fi14 := 0; fi14 < 13; fi14++ {
                    switch "hello" {
                    case "foo bar":
                        v6 = ((v5 * 3.14159) != v5)
                    default:
                        switch fi14 {
                        case -735:
                            v7 = (-2.25 + (0.1 - float32((float64(float32(v1)) + ((0.1 - 0.0) + (v1 / v1))))))
                        }
                    }
                    _ = fi14
                }
                var v15 uint = 0
                v2 = append(v2, 227)
                _ = v15
            }
            if v6 {
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(v4, v7)
                    v17 := GenS1{F0: (1 <= 1), F1: "hello", F2: 42}
                    _ = fi16
                    _ = v17
                }
                v18 := GenS1{F0: (((!v6) || v6) && ((v6 && (42 < 1)) && (1 < 255))), F1: "", F2: 384}
                _ = v18
            }
            _ = v13
        }
        var v19 string = "hello"
        fmt.Println(v1)
        _ = fi12
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            v1 = 1.5
            if v6 {
                var v21 uint64 = 4294967295
                _ = v21
            }
            v22 := GenS2{F0: v6, F1: 1}
            _ = v22
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
