package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0(p0 string) uint32 {
    fmt.Println(p0)
    v4 := []int{32768, 0}
    p0 = ""
    _ = p0
    _ = v4
    return 2147483648
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"", ("" + "hello")}
    var v3 float32 = 0.1
    for fi5 := 0; fi5 < 17; fi5++ {
        if (42 == 1) {
            fmt.Println(fi5, fi5)
        }
        v3 += ((v3 - (1.5 + 3.14159)) / float32((1e10 - float64(3.14159))))
        v3 -= 0.0
        _ = fi5
    }
    fmt.Println(v3)
    v3 += v3
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    switch 1 {
    case -9223372036854775808:
        switch "" {
        case "":
            v7 := GenS1{F0: ((true || (false || true)) && false), F1: 4294967295, F2: 2}
            _ = v7
        case "\n\t":
            fmt.Println(v3)
        }
    case 127:
        if false {
            var v8 uint = 2
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = v8
        }
    case 62:
        if (false || false) {
            if false {
                if ((v3 * -2.25) == v3) {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    if true {
                        var v9 uint8 = 0
                        var v10 int8 = -128
                        _ = v10
                        _ = v9
                    } else {
                        v3 = (((float32(((1e10 - 0.0) + 0.1)) * (v3 / (1e10 * v3))) + -2.25) * 0.1)
                    }
                    v3 += v3
                }
            }
        }
    }
    v3 -= 0.0
    if (-9223372036854775808 > 1) {
        v11 := GenS1{F0: (((1.5 / 1.0) * float64(-2.25)) >= 3.14159), F1: 1, F2: 2147483647}
        fmt.Println(v3)
        switch "foo bar" {
        case "hello":
            var v12 int64 = -581
            _ = v12
        case "":
            v13 := GenS1{F0: (32768 < 49), F1: 1, F2: 4294967295}
            _ = v13
        }
        _ = v11
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint32 = 42
            var v16 bool = true
            v17 := GenS1{F0: v16, F1: 1, F2: (v15 % ((0) | 1))}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: false, F1: 343, F2: 2147483647}
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v3)
    var v20 uint32 = 2147483647
    v21 := v1[("hello" + ((("a" + "foo bar") + "a") + "foo bar"))]
    if (117 < 127) {
        for fi22 := 0; fi22 < 15; fi22++ {
            fmt.Println(fi22, fi22, v3)
            _ = fi22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
