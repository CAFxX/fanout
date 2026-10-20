package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint16 }
func genFn0() int16 {
    v6 := GenS1{F0: 32767, F1: "\n\t"}
    var v7 float32 = 3.14159
    _ = v6
    _ = v7
    return 282
}
func genFn1() int16 {
    var v8 uint = 416
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int64 = 1
        _ = fi9
        _ = v10
    }
    fmt.Println(v8)
    v8 -= v8
    var v11 float32 = ((float32(float64(float32(0.0))) - 1e10) / ((3.14159 - ((1e10 / 3.14159) * (0.1 + 1e10))) + float32(-2.25)))
    _ = v11
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    var v3 uint = 2
    var v4 int32 = -1
    var v5 int = 4294967295
    switch ("" + "hello") {
    case "hello":
        fmt.Println(v5, v5, v5)
    case "\n\t":
        if (42 >= 242) {
            v12 := []int{(-(v5)), v5, v5}
            fmt.Println(v4, v4)
            _ = v12
        }
    }
    var v13 uint8 = uint8(v3)
    switch (("\n\t" + ("a" + "a")) + "\n\t") {
    case "\n\t":
        for fi14 := 0; fi14 < 15; fi14++ {
            v13 += v13
            v15 := []int{(fi14 - -1), 2147483647}
            _ = fi14
            _ = v15
        }
    case "hello":
        var v16 uint16 = 32767
        _ = v16
    case "":
        v17 := []int{}
        _ = v17
    default:
        v18 := GenS1{F0: 65535, F1: ""}
        _ = v18
    }
    v19 := GenS1{F0: 77, F1: ""}
    v20 := GenS1{F0: 2, F1: "\n\t"}
    if (true || true) {
        for fi21 := 0; fi21 < 10; fi21++ {
            v4 += -2147483648
            _ = fi21
        }
        v22 := GenS2{F0: 971}
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        switch "\n\t" {
        case "a":
            switch 0 {
            case 9223372036854775807:
                switch "hello" {
                case "a":
                    fmt.Println(v4, v4, fi23)
                case "foo bar":
                    delete(v2, 42)
                case "\n\t":
                    fmt.Println(fi23, v5, v13)
                default:
                    v4 -= v4
                }
            case -2147483648:
                for fi24 := 0; fi24 < 4; fi24++ {
                    v5 -= (fi23 - 42)
                    v4 += v4
                    v25 := GenS1{F0: 194, F1: ("a" + "")}
                    _ = fi24
                    _ = v25
                }
            case 32768:
                v2[1] = 2147483647
            }
        default:
            var v26 float64 = 1.5
            _ = v26
        }
        _ = fi23
    }
    fmt.Println(v3, v4)
    if true {
        fmt.Println(v5, v4, v13)
        if true {
            v3 = v3
            v3 += (^(v3))
        } else {
            v3 -= v3
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
