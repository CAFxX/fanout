package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint8) string {
    v6 := GenS1{F0: 1}
    var v7 float32 = -2.25
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("\n\t" + ("\n\t" + (("" + "hello") + "\n\t")))
}
func genFn1() uint16 {
    switch -37 {
    case 1:
        var v8 int16 = 1
        _ = v8
    case 9223372036854775807:
        v9 := GenS1{F0: -779}
        _ = v9
    default:
        var v10 float32 = (0.0 - 1e10)
        _ = v10
    }
    if ((3.14159 / (float32((((1e10 - 0.1) + float64(0.0)) + 0.0)) - 0.1)) >= -2.25) {
        var v11 float32 = (-2.25 - (3.14159 / 3.14159))
        fmt.Println(v11)
        v12 := []int{-955}
        _ = v11
        _ = v12
    }
    return 2
}

func main() {
    v1 := []string{"foo bar", ("a" + ("" + "\n\t")), (("a" + "foo bar") + ("" + (("" + "\n\t") + "a")))}
    v2 := []int{-9223372036854775808, 42, 32767, 2147483647}
    var v3 uint8 = 58
    v4 := []uint32{2, 809}
    v5 := []string{"hello", "hello", "a"}
    v3 -= v3
    fmt.Println(v3)
    for fi13 := 0; fi13 < 7; fi13++ {
        v3 = (v3 + (v3 + (v3 | v3)))
        _ = fi13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v14 := GenS1{F0: 42}
    var v15 uint64 = 146
    v16 := GenS1{F0: -9223372036854775808}
    fmt.Println(v3, v3)
    for fi17 := 0; fi17 < 8; fi17++ {
        if (!((true || (((1 != 974) && (((0.1 * 0.0) / (3.14159 - 0.0)) == ((1.5 / 1.5) + 1e10))) && false)) || true)) {
            var v18 float32 = -2.25
            _ = v18
        }
        var v19 uint = 903
        _ = fi17
        _ = v19
    }
    v15 = (v15 % ((v15) | 1))
    var v20 bool = false
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint8 = v3
        v1 = append(v1, "hello")
        switch ("foo bar" + "a") {
        case "\n\t":
            fmt.Println(v3, v3, v15)
        case "":
            v2 = append(v2, 32768)
        case "a":
            v3 = 82
        }
        _ = fi21
        _ = v22
    }
    v20 = v20
    fmt.Println(v3, v20, v20)
    var v23 int8 = 100
    var v24 uint8 = (v3 % ((127) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
