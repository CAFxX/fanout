package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int32; F1 uint }
func genFn0(p0 int32, p1 int16) int8 {
    v5 := GenS1{F0: 942}
    fmt.Println(p1)
    switch -523 {
    case -2147483648:
        v6 := []int{-2147483648, -902}
        _ = v6
    case 42:
        v7 := []int{-1}
        _ = v7
    default:
        if (65535 < 0) {
            var v8 int = 4294967295
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 100
}

func main() {
    var v1 bool = true
    var v2 int32 = 1
    var v3 uint = 729
    v4 := make(map[string]int)
    var v9 bool = v1
    v10 := []int{-339, 0, -2147483648}
    if false {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 uint8 = 2
            switch "\n\t" {
            case "":
                v1 = (-838 >= -1)
            case "hello":
                v2 = (v2 >> (((v3 % ((v3) | 1))) & 63))
            }
            _ = fi11
            _ = v12
        }
        var v13 uint = v3
        _ = v13
    } else {
        if ("a" <= "") {
            _ = len(v4)
        } else {
            v14 := GenS2{F0: 959, F1: -1}
            switch 9223372036854775807 {
            case 19:
                if ((!(v9 || (float32(0.0) >= 1e10))) && v1) {
                    fmt.Println(v2)
                } else {
                    var v15 int = 127
                    _ = v15
                }
            default:
                v16 := GenS2{F0: 2147483648, F1: 128}
                _ = v16
            }
            _ = v14
        }
    }
    v9 = true
    if v9 {
        var v17 bool = (v2 >= v2)
        _ = v17
    } else {
        v2 = -2147483648
        var v18 int64 = -272
        _ = v18
    }
    v19 := GenS2{F0: 1, F1: 32768}
    if v1 {
        for fi20 := 0; fi20 < 19; fi20++ {
            v3 += v3
            v2 += (v2 * 2147483647)
            var v21 uint8 = 128
            _ = fi20
            _ = v21
        }
        v9 = (0 == 42)
    } else {
        switch "foo bar" {
        case "hello":
            if v1 {
                if (!(!true)) {
                    if (-117 != 100) {
                        _ = len(v10)
                        v22 := v4[("a" + "\n\t")]
                        v23 := GenS2{F0: 42, F1: int(v3)}
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(v1, v3)
                }
            }
        case "foo bar":
            if v1 {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
        default:
            if v1 {
                v24 := v4["foo bar"]
                _ = v24
            } else {
                if (!(-2147483648 <= (v2 + -2147483648))) {
                    if (!true) {
                        var v25 int64 = -9223372036854775808
                        _ = v25
                    }
                    fmt.Println(v2, v1, v3)
                } else {
                    v26 := GenS3{F0: v2, F1: v3}
                    _ = v26
                }
                v3 -= 18446744073709551615
                var v27 uint = 4294967295
                _ = v27
            }
        }
        var v28 int = -1
        var v29 uint8 = 2
        _ = v28
        _ = v29
    }
    delete(v4, "hello")
    v30 := GenS1{F0: (v3 / ((v3) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
