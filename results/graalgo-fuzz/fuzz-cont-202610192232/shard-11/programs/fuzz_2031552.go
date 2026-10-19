package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (float64(3.14159) / 1.0)
    v3 := make(map[int]string)
    var v4 int = -501
    v5 := make(map[int]bool)
    var v6 string = (("hello" + "hello") + ((("" + "\n\t") + ("\n\t" + "a")) + ""))
    v4 = v4
    fmt.Println(v4, v2, v2)
    v4 += (v4 - v4)
    fmt.Println(v2)
    switch (v4 << ((255) & 63)) {
    case -1:
        v7 := GenS1{F0: int16(v4), F1: -751, F2: 0}
        _ = v7
    case 32768:
        fmt.Println(v4, v4, v6)
    default:
        if (0 <= 0) {
            fmt.Println(v6, v6)
            v8 := GenS1{F0: -1, F1: 42, F2: 2147483647}
            _ = v8
        } else {
            switch 4294967295 {
            case 0:
                var v9 float64 = ((v2 - float64(0.0)) + 0.1)
                _ = v9
            }
        }
    }
    v2 += v2
    v2 -= v2
    var v10 int64 = 733
    if ((v10 / (((-(v10))) | 1)) != v10) {
        v2 += float64(0.1)
        v11 := []int{4294967295}
        _ = v11
    }
    if false {
        v12 := []int{(-(v4))}
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                v10 = v10
                if false {
                    if true {
                        switch v6 {
                        case "":
                            fmt.Println(v2, fi13)
                        case "foo bar":
                            v2 = ((float64(3.14159) + (((1.5 / (v2 / (0.0 - v2))) * -2.25) - float64((-2.25 - 1e10)))) + v2)
                            if len(v12) > 0 {
                                v12[0] = -456
                            }
                            v12 = append(v12, 724)
                            var v15 float64 = -2.25
                            _ = v15
                        }
                        var v16 float64 = v2
                        if ("a" != v6) {
                            fmt.Println(v16)
                            switch v4 {
                            case -1:
                                v16 += (0.0 / v2)
                            }
                        }
                        _ = v16
                    } else {
                        v6 = v6
                        var v17 uint16 = 742
                        _ = v17
                    }
                    var v18 int8 = -128
                    _ = v18
                } else {
                    v2 -= ((float64(((-2.25 - 0.0) - -2.25)) - float64(((-2.25 + 1e10) / float32(v2)))) / v2)
                    _ = len(v12)
                    v6 = (v6 + ((v6 + ("a" + (v6 + v6))) + (v6 + v6)))
                }
                v2 += ((-2.25 * (float64(-2.25) - v2)) - v2)
                _ = fi14
            }
            var v19 bool = true
            fmt.Println(fi13, v10, v19)
            _ = fi13
            _ = v19
        }
        _ = v12
    }
    v20 := GenS1{F0: -145, F1: (v10 - v10), F2: 2147483647}
    if (0 == -32768) {
        v21 := GenS1{F0: -32768, F1: 42, F2: 42}
        var v22 int32 = -2147483648
        v4 = v4
        _ = v21
        _ = v22
    } else {
        _ = len(v1)
        v23 := GenS2{F0: 2147483648}
        _ = v23
    }
    _ = len(v1)
    v24 := v3[v4]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
