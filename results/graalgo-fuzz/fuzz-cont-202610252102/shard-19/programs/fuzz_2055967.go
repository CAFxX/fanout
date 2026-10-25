package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint, p1 int8) int8 {
    var v5 int = 961
    p1 += p1
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v5)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return (p1 % ((100) | 1))
}

func main() {
    var v1 int16 = -309
    var v2 int = 1
    var v3 int8 = 0
    var v4 uint8 = 127
    v7 := GenS1{F0: (v4 % (((v4 * v4)) | 1)), F1: "hello", F2: 0}
    if false {
        v8 := []int{v2, (v2 - 65), v2, v2}
        for fi9 := 0; fi9 < 17; fi9++ {
            _ = len(v8)
            _ = fi9
        }
        if false {
            v8 = append(v8, (v2 << ((42) & 63)))
            v2 = 32768
            fmt.Println(v1)
        } else {
            fmt.Println(v1, v1, v2)
            v2 = (-(v2))
        }
        _ = v8
    } else {
        var v10 string = "a"
        if (700 > 255) {
            fmt.Println(v4, v2, v10)
            v2 = v2
            if ((870 != 1) && (!false)) {
                var v11 int64 = 0
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := []int{-2147483648, (v2 | 32767), v2}
                    if (false || true) {
                        v2 = v2
                        var v14 int32 = -761
                        if (!false) {
                            var v15 bool = false
                            _ = v15
                        }
                        _ = v14
                    } else {
                        v16 := GenS1{F0: (v4 | v4), F1: v10, F2: v11}
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
        } else {
            v17 := GenS2{F0: 0}
            _ = v17
        }
        v3 -= (^(v3))
        _ = v10
    }
    v3 = v3
    v1 = (v1 / (((v1 >> ((4294967295) & 63))) | 1))
    if false {
        v3 = (v3 * v3)
        fmt.Println(v4, v3, v1)
    } else {
        switch "" {
        case "\n\t":
            v4 = v4
        case "":
            fmt.Println(v4, v3)
        case "foo bar":
            fmt.Println(v1, v1, v3)
        }
        var v18 uint16 = 255
        var v19 uint32 = 701
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := []int{}
        v4 += v4
        _ = fi20
        _ = v21
    }
    v1 += v1
    if true {
        v22 := []int{v2}
        var v23 string = "hello"
        _ = v22
        _ = v23
    }
    var v24 int = v2
    if (0.0 == 0.1) {
        fmt.Println(v4)
        var v25 int = (^(v24))
        _ = v25
    } else {
        var v26 float32 = 1e10
        v27 := []int{128, int(v3), int(v4), (v24 % (((v24 - (-(v24)))) | 1))}
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v7
}
