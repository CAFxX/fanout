package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 int16 }

func main() {
    v1 := []int{377, 4294967295, 42}
    var v2 string = ""
    var v3 int = 2147483647
    var v4 string = (v2 + (v2 + v2))
    v5 := make(map[int]int)
    fmt.Println(v3, v4)
    fmt.Println(v2)
    var v6 bool = (false && (0 < -509))
    v7 := GenS1{F0: -716, F1: 197}
    v3 = v3
    switch v3 {
    case -497:
        v2 = v4
    case 128:
        if (462 < 0) {
            v8 := GenS1{F0: 42, F1: 131}
            var v9 bool = (((!false) && (364 <= 4294967295)) && v6)
            if false {
                _ = len(v5)
                if (!(((v3 != v3) || v9) || v9)) {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        v3 = 32767
                        var v11 bool = (v6 && false)
                        _ = fi10
                        _ = v11
                    }
                }
                var v12 float64 = 0.1
                _ = v12
            } else {
                v9 = v9
                for fi13 := 0; fi13 < 9; fi13++ {
                    fmt.Println(v2)
                    _ = fi13
                }
            }
            _ = v8
            _ = v9
        } else {
            var v14 int32 = -1
            _ = v14
        }
    case 127:
        var v15 float32 = 0.0
        _ = v15
    default:
        switch v4 {
        case "\n\t":
            fmt.Println(v3, v2)
        case "":
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS2{F0: 1000, F1: v6, F2: -1}
                if ((1.5 * float32((-2.25 + -2.25))) < ((1e10 / (0.1 / 1e10)) * (float32(-2.25) + (3.14159 + float32(1e10))))) {
                    v2 = (v2 + "\n\t")
                    var v18 string = (v4 + (v4 + "a"))
                    _ = v18
                } else {
                    if (!(v6 && true)) {
                        var v19 int32 = 365
                        v20 := GenS2{F0: 42, F1: true, F2: 0}
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(v4)
                    }
                    if len(v1) > 0 {
                        v1[0] = -2147483648
                    }
                }
                fmt.Println(v2, v3)
                _ = fi16
                _ = v17
            }
        default:
            for fi21 := 0; fi21 < 20; fi21++ {
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := GenS1{F0: -513, F1: 2}
                    v2 = v2
                    _ = fi22
                    _ = v23
                }
                v3 -= (v3 + (v3 >> ((65535) & 63)))
                for fi24 := 0; fi24 < 19; fi24++ {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v6)
                    default:
                        v4 = v2
                    }
                    fmt.Println(v2, v4)
                    _ = len(v1)
                    _ = fi24
                }
                _ = fi21
            }
        }
    }
    v25 := GenS1{F0: 42, F1: 6}
    var v26 int8 = -128
    fmt.Println(v2, v26, v26)
    fmt.Println(v4, v2, v6)
    v26 -= 0
    var v27 uint8 = 146
    v27 -= v27
    if v6 {
        _ = len(v5)
        var v28 string = (("\n\t" + "foo bar") + v4)
        _ = v28
    }
    var v29 uint32 = 2147483648
    var v30 uint64 = 2
    v6 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
