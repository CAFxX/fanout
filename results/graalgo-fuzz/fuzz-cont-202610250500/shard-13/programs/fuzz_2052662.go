package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 bool, p1 uint8) uint16 {
    v7 := GenS1{F0: -76, F1: ("" + "foo bar")}
    fmt.Println(p0)
    for fi8 := 0; fi8 < 7; fi8++ {
        switch "\n\t" {
        case "":
            p1 = (^(p1))
        default:
            if (true && false) {
                switch fi8 {
                case 42:
                    p0 = p0
                case -9223372036854775808:
                    if (-2147483648 == 2147483647) {
                        v9 := GenS1{F0: 127, F1: ("foo bar" + "")}
                        _ = v9
                    }
                }
                var v10 int = (fi8 & -446)
                if false {
                    v10 += (fi8 / ((-2147483648) | 1))
                    var v11 int = -2147483648
                    _ = v11
                }
                _ = v10
            }
        }
        switch "\n\t" {
        case "":
            if p0 {
                p0 = p0
                v12 := []int{fi8}
                _ = v12
            }
        }
        p0 = p0
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    v3 := []uint32{2147483648, 157, 2147483648}
    v4 := []uint32{0, 2}
    var v5 uint = 396
    var v6 string = ""
    v13 := GenS1{F0: -1, F1: "foo bar"}
    var v14 int8 = 58
    if (2 < 2) {
        switch v6 {
        case "a":
            var v15 string = v6
            _ = v15
        }
    } else {
        if false {
            if len(v4) > 0 {
                v4[0] = 2
            }
            v16 := GenS1{F0: v14, F1: "foo bar"}
            _ = v16
        }
        switch "hello" {
        case "a":
            v17 := GenS1{F0: v14, F1: v6}
            _ = v17
        case "foo bar":
            fmt.Println(v14, v14, v14)
        }
        fmt.Println(v14, v14, v14)
    }
    v6 = ""
    if (false || true) {
        if true {
            var v18 uint8 = 91
            _ = len(v4)
            _ = v18
        }
    } else {
        var v19 float32 = 1.5
        var v20 int = 2147483647
        delete(v2, (("hello" + v6) + v6))
        _ = v19
        _ = v20
    }
    var v21 int16 = -654
    switch "hello" {
    case "a":
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := GenS1{F0: 42, F1: (v6 + (("\n\t" + (v6 + "\n\t")) + (v6 + (((v6 + "a") + ("hello" + v6)) + v6))))}
            var v24 bool = true
            _ = fi22
            _ = v23
            _ = v24
        }
    case "":
        v5 = v5
    default:
        v25 := v2[v6]
        _ = v25
    }
    v14 -= v14
    var v26 int16 = 304
    v27 := GenS1{F0: v14, F1: (v6 + (v6 + "a"))}
    fmt.Println(v6, v5)
    v6 = (v6 + "\n\t")
    var v28 int8 = (v14 + (v14 ^ -128))
    fmt.Println(v6, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
