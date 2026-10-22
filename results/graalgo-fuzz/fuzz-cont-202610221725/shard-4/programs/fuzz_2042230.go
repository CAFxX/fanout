package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0() string {
    var v7 float32 = (0.1 * 0.1)
    var v8 uint32 = 2
    v7 += (((v7 + v7) - 0.0) - v7)
    v7 = (v7 + ((v7 + (-2.25 * v7)) - ((1.5 + float32(float64(((v7 - 0.0) / (v7 * (0.1 + (((3.14159 + v7) / v7) + 1.5))))))) + v7)))
    fmt.Println(v7, v8)
    _ = v7
    _ = v8
    return (("a" + "a") + (("\n\t" + "a") + (("\n\t" + "") + "")))
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    v4 := []string{"\n\t", ""}
    var v5 string = "a"
    v6 := []int{(v1 / ((v1) | 1))}
    v9 := v2[("\n\t" + (v5 + "foo bar"))]
    v9 = v9
    var v10 uint16 = 528
    switch v5 {
    case "\n\t":
        if (-1 <= -378) {
            switch v1 {
            case -2147483648:
                v5 = v5
            case 5:
                fmt.Println(v5)
            case 2147483647:
                var v11 string = ("a" + v5)
                _ = v11
            default:
                v12 := GenS1{F0: -9223372036854775808}
                _ = v12
            }
        } else {
            v1 = 127
        }
    case "a":
        var v13 int16 = 0
        _ = v13
    }
    var v14 float32 = 1e10
    v15 := GenS1{F0: -9223372036854775808}
    switch v1 {
    case 2147483647:
        if len(v6) > 0 {
            v6[0] = v1
        }
    }
    var v16 bool = (v9 && (-419 == -1))
    if v9 {
        if (537 == 1) {
            fmt.Println(v16, v10)
        } else {
            switch v5 {
            case "":
                var v17 float64 = 0.1
                _ = v17
            case "a":
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 bool = (!v16)
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 bool = v9
                        var v22 uint16 = v10
                        v10 = (v22 & v10)
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                }
            case "\n\t":
                if v9 {
                    v10 = v10
                    if (42 <= 33) {
                        v10 = v10
                    }
                    fmt.Println(v10, v16)
                } else {
                    v23 := GenS1{F0: 1}
                    _ = v23
                }
            default:
                v6 = append(v6, 32768)
            }
        }
    }
    v24 := GenS2{F0: 6, F1: (^(v10)), F2: 18446744073709551615}
    switch ("foo bar" + (("" + "foo bar") + v5)) {
    case "":
        if (v10 != v10) {
            fmt.Println(v16, v14)
        }
    }
    var v25 int8 = -1
    var v26 uint8 = 1
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 string = ""
        v25 = v25
        if ((v25 / ((v25) | 1)) <= (v25 | (v25 & (v25 - (v25 / (((v25 - v25)) | 1)))))) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v29 float32 = (v14 + v14)
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
