package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32, p1 int32, p2 uint) uint64 {
    p0 = p0
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 11; fi7++ {
        p1 -= p1
        _ = fi7
    }
    fmt.Println(v6, p1, p1)
    var v8 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return 748
}
func genFn1() string {
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 bool = false
            _ = fi9
            _ = v10
        }
    } else {
        v11 := []int{}
        _ = v11
    }
    switch 42 {
    case 128:
        v12 := []int{128, 4294967295, 579, 4294967295}
        _ = v12
    }
    return ("a" + "hello")
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = (v1 / 1.5)
    var v3 float64 = ((0.1 - 0.1) - 1e10)
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    switch "\n\t" {
    case "foo bar":
        v2 = -2.25
    case "a":
        fmt.Println(v1, v3)
    default:
        fmt.Println(v1, v1, v1)
    }
    var v13 int8 = 127
    v14 := []int{-2147483648}
    v13 = (v13 * 0)
    v15 := GenS1{F0: 610, F1: 42}
    v13 = v13
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            switch ("hello" + "hello") {
            case "hello":
                if (true || (!true)) {
                    var v18 int16 = -32768
                    _ = v18
                }
            case "a":
                var v19 string = (("a" + ((("\n\t" + ("\n\t" + ((("\n\t" + ("hello" + "a")) + "foo bar") + "hello"))) + "a") + ("foo bar" + ((("hello" + "foo bar") + "foo bar") + "foo bar")))) + "foo bar")
                _ = v19
            case "":
                v3 = float64(((0.0 / 1e10) * (float32(((v3 * (-2.25 + v3)) * (float64(((1e10 + (v2 + (3.14159 + v1))) + float32(0.0))) / (v3 * v3)))) * (v2 / (v2 - v1)))))
                v20 := GenS1{F0: -32768, F1: 42}
                var v21 uint16 = 260
                _ = v20
                _ = v21
            default:
                switch fi16 {
                case 9223372036854775807:
                    switch fi16 {
                    case 1:
                        v2 += v1
                        v13 -= (-(v13))
                        fmt.Println(v2, v13)
                    case 0:
                        var v22 int16 = 1
                        _ = v22
                    }
                    v13 -= (v13 - 122)
                default:
                    var v23 int64 = 9223372036854775807
                    _ = v23
                }
            }
            v24 := GenS1{F0: 1000, F1: 255}
            _ = fi17
            _ = v24
        }
        var v25 uint16 = 0
        v26 := GenS2{F0: 9223372036854775807}
        _ = fi16
        _ = v25
        _ = v26
    }
    if true {
        fmt.Println(v13, v3, v13)
        v13 = v13
    } else {
        v27 := v5[""]
        var v28 bool = (true && (v27 && v27))
        _ = v27
        _ = v28
    }
    switch "foo bar" {
    case "a":
        v14 = append(v14, -1)
    case "":
        v29 := GenS1{F0: 1, F1: 128}
        _ = v29
    case "\n\t":
        v30 := GenS2{F0: -1}
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 16; fi31++ {
            v32 := GenS2{F0: 0}
            _ = fi31
            _ = v32
        }
        _ = len(v14)
        if len(v14) > 0 {
            v14[0] = -276
        }
    } else {
        fmt.Println(v3, v13)
        fmt.Println(v3, v13, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
