package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string }
func genFn0() int16 {
    var v7 int64 = 1
    v7 = v7
    _ = v7
    return -1
}
func genFn1(p0 string, p1 int16) int {
    v8 := GenS1{F0: -9223372036854775808}
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS2{F0: -128, F1: "\n\t"}
            p0 = (("\n\t" + p0) + p0)
            v12 := GenS2{F0: 0, F1: p0}
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(fi9)
        _ = fi9
    }
    var v13 int16 = 821
    for fi14 := 0; fi14 < 19; fi14++ {
        if (3.14159 >= (1e10 - (3.14159 - (3.14159 + float32(1.5))))) {
            v15 := []int{fi14, 32767, -302}
            v16 := GenS1{F0: 9223372036854775807}
            p1 = (v13 - -630)
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    switch 32768 {
    case 913:
        switch "a" {
        case "":
            var v17 uint32 = 2147483647
            _ = v17
        default:
            switch 886 {
            case -976:
                v18 := GenS1{F0: -1}
                _ = v18
            case -580:
                v13 += 42
            case 128:
                var v19 int16 = -32768
                _ = v19
            default:
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 int16 = (-(p1))
                    v22 := GenS2{F0: 7, F1: ("\n\t" + p0)}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        }
    case 9223372036854775807:
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint16 = 42
            fmt.Println(p1, v24, p1)
            _ = fi23
            _ = v24
        }
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    return 855
}

func main() {
    v1 := []int{127, 2147483647, 962}
    var v2 int8 = 100
    var v3 float64 = (-2.25 * 0.1)
    var v4 float32 = (3.14159 * 0.0)
    var v5 int8 = (v2 | v2)
    var v6 float32 = (v4 + ((((-2.25 - v4) / 3.14159) * 1e10) + v4))
    v3 += 1e10
    fmt.Println(v2, v5, v6)
    var v25 uint16 = 32767
    var v26 float32 = 0.0
    v27 := GenS2{F0: v5, F1: ""}
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS2{F0: (v2 / ((v5) | 1)), F1: ""}
        v3 += (3.14159 * v3)
        v6 -= float32(v3)
        _ = fi28
        _ = v29
    }
    switch ("" + ((("hello" + (("" + ("" + (("" + "a") + (("hello" + "foo bar") + "a")))) + ("foo bar" + "foo bar"))) + "hello") + ("a" + ("foo bar" + (("foo bar" + "hello") + ((("a" + "foo bar") + "a") + ("foo bar" + ("" + "foo bar")))))))) {
    case "":
        v30 := GenS1{F0: 743}
        _ = v30
    case "hello":
        v6 += v6
    default:
        if (v3 == v3) {
            var v31 uint = 42
            if false {
                v32 := GenS2{F0: 1, F1: ("" + ("\n\t" + "foo bar"))}
                var v33 uint32 = 2147483647
                if len(v1) > 0 {
                    v1[0] = 1
                }
                _ = v32
                _ = v33
            }
            if len(v1) > 0 {
                v1[0] = 128
            }
            _ = v31
        }
    }
    _ = len(v1)
    var v34 int8 = (v5 | -1)
    v3 -= v3
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v34, v6, v6)
        v36 := GenS2{F0: (v2 ^ v5), F1: "foo bar"}
        _ = fi35
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
