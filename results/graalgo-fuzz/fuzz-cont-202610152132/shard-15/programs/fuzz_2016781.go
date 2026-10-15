package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int }
func genFn0() uint32 {
    var v5 int16 = 42
    v6 := []int{}
    v7 := GenS1{F0: -1, F1: 100}
    _ = v5
    _ = v6
    _ = v7
    return 530
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 32767
    var v3 int8 = 42
    v4 := make(map[string]bool)
    fmt.Println(v2, v2)
    v3 += (-(v3))
    var v8 int8 = (v3 & (v3 - v3))
    var v9 uint16 = 32767
    var v10 int16 = 817
    var v11 int64 = -993
    v12 := []int{-2147483648, (v2 + v2)}
    v10 = 1
    switch (^(v2)) {
    case -189:
        for fi13 := 0; fi13 < 7; fi13++ {
            if (("" + "") < "a") {
                switch ((((("\n\t" + "a") + "a") + ("a" + "")) + "\n\t") + "") {
                case "foo bar":
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v15 := GenS1{F0: -2147483648, F1: -68}
                        _ = fi14
                        _ = v15
                    }
                case "":
                    if false {
                        if true {
                            v3 = (v3 % ((v8) | 1))
                            v16 := GenS1{F0: v2, F1: v3}
                            _ = v16
                        } else {
                            _ = len(v1)
                        }
                    } else {
                        var v17 float64 = 3.14159
                        v2 = (fi13 % (((fi13 % (((-(v2))) | 1))) | 1))
                        if (v2 < fi13) {
                            if len(v12) > 0 {
                                v12[0] = (-(fi13))
                            }
                        }
                        _ = v17
                    }
                }
            } else {
                v18 := GenS1{F0: (fi13 >> ((65535) & 63)), F1: v8}
                _ = v18
            }
            switch fi13 {
            case -1:
                var v19 int64 = (v11 + v11)
                _ = v19
            default:
                var v20 float32 = 3.14159
                _ = v20
            }
            _ = fi13
        }
    case 0:
        if (((v10 >> ((0) & 63)) <= (v10 + -32768)) || true) {
            fmt.Println(v2, v8)
        } else {
            var v21 float64 = 0.0
            _ = v21
        }
    }
    var v22 uint32 = genFn0()
    switch 128 {
    case 408:
        if (false || (v22 != (v22 / ((v22) | 1)))) {
            for fi23 := 0; fi23 < 1; fi23++ {
                fmt.Println(v11, v11)
                _ = fi23
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                v8 = (-(v8))
                _ = fi24
            }
        } else {
            v25 := GenS2{F0: v2}
            var v26 string = "hello"
            var v27 int = (v2 | v2)
            _ = v25
            _ = v26
            _ = v27
        }
    case -9223372036854775808:
        v28 := GenS2{F0: v2}
        _ = v28
    default:
        fmt.Println(v2)
    }
    var v29 int8 = -1
    switch "" {
    case "foo bar":
        var v30 int = (v2 & (v2 & (v2 * v2)))
        _ = v30
    case "":
        v31 := GenS2{F0: v2}
        _ = v31
    case "a":
        v8 = v29
    }
    switch "" {
    case "foo bar":
        v29 -= v8
    case "a":
        var v32 uint = 18446744073709551615
        _ = v32
    case "":
        fmt.Println(v10)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
