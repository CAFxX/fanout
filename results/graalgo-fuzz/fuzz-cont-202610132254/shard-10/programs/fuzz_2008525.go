package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 float32 = 1.5
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    if (42 >= 142) {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 float32 = 1.5
            v8 := []int{(fi6 >> ((96) & 63)), fi6, fi6}
            if true {
                var v9 uint32 = 674
                var v10 int32 = -2147483648
                for fi11 := 0; fi11 < 4; fi11++ {
                    switch fi11 {
                    case 2147483647:
                        v7 = v3
                    case -102:
                        v12 := GenS1{F0: -1}
                        _ = v12
                    case 128:
                        v9 = 2147483647
                    default:
                        if (false && true) {
                            v13 := GenS1{F0: 1}
                            _ = v13
                        }
                    }
                    fmt.Println(v9)
                    _ = fi11
                }
                _ = v10
                _ = v9
            } else {
                var v14 uint16 = 1
                _ = v14
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        switch 127 {
        case 43:
            fmt.Println(v5)
        default:
            switch "\n\t" {
            case "foo bar":
                if (32767 == -1) {
                    v15 := []int{127, 9223372036854775807, -576, -1}
                    var v16 string = "\n\t"
                    _ = v15
                    _ = v16
                }
            default:
                v17 := GenS1{F0: 0}
                _ = v17
            }
        }
        _ = len(v1)
    }
    v3 = (0.0 * (1.5 + (1.5 / 3.14159)))
    v5 = v5
    var v18 bool = (42 != -97)
    if ("hello" <= ("\n\t" + "")) {
        v5 += (^(v5))
    } else {
        var v19 int32 = -979
        v19 += 42
        _ = v19
    }
    if false {
        var v20 uint32 = v5
        _ = v20
    } else {
        v3 -= -2.25
        v5 = 965
        fmt.Println(v18)
    }
    v21 := GenS1{F0: 1}
    v22 := GenS1{F0: -281}
    var v23 int32 = 1
    if false {
        var v24 int = 677
        var v25 int8 = -28
        _ = v24
        _ = v25
    } else {
        v26 := []int{-2147483648}
        v3 = (v3 * ((float32(-2.25) + ((-2.25 - v3) - (float32(-2.25) * ((v3 + 1e10) + v3)))) / v3))
        _ = v26
    }
    v27 := []int{-2147483648, 0, 1, -9223372036854775808}
    if v18 {
        var v28 int64 = -1
        _ = v28
    }
    var v29 float64 = (1.5 + ((1.5 / 1.5) / float64((v3 / v3))))
    var v30 string = ("foo bar" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
