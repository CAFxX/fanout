package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -128
    v3 := make(map[string]bool)
    var v4 int16 = 42
    var v5 bool = false
    v6 := v3[((("" + "foo bar") + ("\n\t" + "foo bar")) + "foo bar")]
    var v7 float32 = 1e10
    v6 = v5
    v2 -= -128
    var v8 float64 = ((3.14159 * (((1e10 / 3.14159) * (0.0 + 0.0)) / 1.0)) * float64(v7))
    v4 -= (-(v4))
    if (!v6) {
        var v9 uint8 = 0
        fmt.Println(v7)
        var v10 int8 = v2
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            v8 -= ((3.14159 / 0.1) - (v8 + v8))
            _ = fi11
        }
        v8 += v8
    }
    v12 := GenS1{F0: v2}
    if false {
        v8 = v8
        switch -788 {
        case -1:
            var v13 int64 = 810
            _ = v13
        case 222:
            var v14 int = 128
            _ = v14
        case 127:
            if ((v2 ^ -128) > (v2 * 25)) {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v5, v2)
                case "a":
                    v1[42] = ("a" + "a")
                case "":
                    var v15 float64 = 3.14159
                    _ = v15
                }
            }
        default:
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 string = ("foo bar" + "foo bar")
                var v18 uint16 = 42
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v19 := []int{4294967295, -486, 619, 4294967295}
        _ = v19
    }
    v20 := []int{127, -582}
    var v21 string = "hello"
    switch v21 {
    case "":
        var v22 int32 = 405
        _ = v22
    case "hello":
        v23 := GenS1{F0: 100}
        _ = v23
    }
    v6 = (true && (2 == 4294967295))
    v24 := v3[((v21 + "a") + v21)]
    v6 = (0.0 <= float32((v8 * (((0.0 - (v8 * v8)) / v8) - 0.0))))
    if (!false) {
        v25 := GenS1{F0: 76}
        v2 -= 42
        _ = v25
    } else {
        var v26 uint16 = 165
        v21 = ""
        v27 := GenS1{F0: v2}
        _ = v26
        _ = v27
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
