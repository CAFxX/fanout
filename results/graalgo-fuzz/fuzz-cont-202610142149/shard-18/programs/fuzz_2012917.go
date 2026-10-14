package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int16 }
func genFn0(p0 int16) string {
    v4 := GenS1{F0: (65535 < 0), F1: (p0 % (((p0 * p0)) | 1)), F2: p0}
    p0 += -32768
    if true {
        v5 := GenS1{F0: false, F1: p0, F2: -1}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: (42 == 255), F1: (p0 % ((p0) | 1)), F2: (p0 << ((1) & 63))}
            for fi8 := 0; fi8 < 12; fi8++ {
                switch -209 {
                case 2147483647:
                    fmt.Println(fi8, fi6)
                default:
                    var v9 int = fi8
                    var v10 uint64 = 149
                    v11 := []int{}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                var v12 uint8 = 93
                v13 := GenS1{F0: true, F1: p0, F2: (^(p0))}
                _ = fi8
                _ = v12
                _ = v13
            }
            _ = fi6
            _ = v7
        }
        if (!(2147483647 != 0)) {
            var v14 uint8 = 128
            _ = v14
        }
    }
    var v15 uint = 1
    _ = p0
    _ = v15
    _ = v4
    return (("foo bar" + ("" + "hello")) + (("foo bar" + ("" + ((("a" + ("\n\t" + "a")) + "foo bar") + ((("foo bar" + "") + "foo bar") + "\n\t")))) + "a"))
}
func genFn1() int8 {
    v16 := []int{}
    if false {
        switch ("hello" + "a") {
        case "foo bar":
            var v17 uint16 = 32767
            _ = v17
        }
        var v18 int8 = 100
        _ = v18
    } else {
        if ("foo bar" > "\n\t") {
            if true {
                var v19 bool = true
                _ = v19
            }
        } else {
            var v20 uint16 = 0
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        if true {
            var v22 int16 = 42
            _ = v22
        }
        _ = fi21
    }
    var v23 int = -899
    if len(v16) > 0 {
        v16[0] = 0
    }
    _ = v16
    _ = v23
    return 3
}

func main() {
    var v1 bool = false
    var v2 int32 = -263
    var v3 float32 = 3.14159
    v3 -= 0.1
    var v24 int32 = (v2 * (^(v2)))
    var v25 int8 = -128
    v26 := []int{}
    fmt.Println(v3, v25, v2)
    if len(v26) > 0 {
        v26[0] = -601
    }
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    v2 += v24
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS1{F0: (((float32((((1.5 / 1e10) + (1.5 - ((1e10 * (-2.25 * 0.0)) / float64(((float32((3.14159 / 1.0)) + (v3 / -2.25)) * (v3 + v3)))))) + 0.1)) * (v3 * float32((float64(float32(-2.25)) - 1.5)))) + (v3 + ((1e10 * v3) - 0.0))) != v3), F1: 32767, F2: -1}
        switch "" {
        case "hello":
            var v29 string = "a"
            _ = v29
        case "":
            v26 = append(v26, (fi27 % ((fi27) | 1)))
        case "\n\t":
            if true {
                var v30 string = ("" + "a")
                fmt.Println(v24, v24, v30)
                _ = v30
            }
        }
        v3 = ((1e10 - v3) * (1.5 * v3))
        _ = fi27
        _ = v28
    }
    v3 -= 1e10
    v25 = 14
    v31 := GenS1{F0: (3.14159 <= 3.14159), F1: -1, F2: -391}
    fmt.Println(v3)
    var v32 float64 = float64(float32(((3.14159 * (0.1 * 1e10)) - (float64(float32((1.5 * 0.0))) - 0.0))))
    if len(v26) > 0 {
        v26[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
}
