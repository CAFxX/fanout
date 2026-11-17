package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0() uint64 {
    if (-9223372036854775808 >= 9223372036854775807) {
        var v4 uint64 = 695
        _ = v4
    }
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(fi5)
        fmt.Println(fi5)
        var v6 int8 = -44
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2
    var v3 float64 = ((-2.25 * (3.14159 * (1e10 / 3.14159))) * 1e10)
    v8 := GenS1{F0: 42, F1: "foo bar", F2: 0}
    fmt.Println(v3, v2)
    v2 = v2
    var v9 uint32 = 2
    if ("" >= ("a" + "")) {
        v10 := GenS1{F0: 25, F1: "", F2: int16(v2)}
        _ = v10
    } else {
        v11 := GenS1{F0: 42, F1: "a", F2: 248}
        _ = v11
    }
    var v12 uint32 = v2
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            switch "\n\t" {
            case "hello":
                v3 -= v3
            case "a":
                v9 -= v2
            default:
                v15 := GenS1{F0: 1, F1: "\n\t", F2: 1000}
                _ = v15
            }
            delete(v1, fi13)
            if ((347 != -32768) && false) {
                if (-2.25 != 3.14159) {
                    if (127 != 0) {
                        v3 -= -2.25
                    } else {
                        var v16 int64 = 9223372036854775807
                        v3 = 3.14159
                        if true {
                            v17 := v1[fi13]
                            var v18 int = (fi14 + fi13)
                            var v19 int32 = 715
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        _ = v16
                    }
                } else {
                    var v20 uint64 = genFn0()
                    if true {
                        fmt.Println(v20, v12, v2)
                    } else {
                        var v21 float64 = ((((v3 / (v3 / v3)) * v3) * 0.0) * (3.14159 - 1e10))
                        var v22 string = (("foo bar" + (("\n\t" + ("foo bar" + ((("" + "") + ((((("\n\t" + "") + "foo bar") + "a") + "hello") + (("foo bar" + "\n\t") + ""))) + ("foo bar" + "hello")))) + "a")) + "")
                        _ = v21
                        _ = v22
                    }
                    var v23 uint8 = 1
                    v24 := []int{fi13}
                    if ((v20 ^ (v20 >> ((18446744073709551615) & 63))) < genFn0()) {
                        v25 := GenS2{F0: ((false || (fi13 < fi14)) && true), F1: true}
                        _ = v25
                    }
                    _ = v20
                    _ = v23
                    _ = v24
                }
                if (!false) {
                    v26 := []int{}
                    v3 -= 0.1
                    _ = v26
                }
                delete(v1, (^(fi13)))
            } else {
                switch "" {
                case "a":
                    v9 = (v2 / ((v9) | 1))
                default:
                    var v27 uint32 = (v9 | v9)
                    _ = v27
                }
            }
            _ = fi14
        }
        switch (fi13 ^ (fi13 ^ (fi13 ^ fi13))) {
        case -1:
            v28 := []int{fi13}
            _ = v28
        case 105:
            v3 = (v3 / ((v3 + v3) + 1e10))
        }
        _ = fi13
    }
    var v29 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v8
    _ = v9
}
