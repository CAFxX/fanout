package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 bool }
type GenS2 struct { F0 bool; F1 uint8; F2 string }
type GenS3 struct { F0 int8 }
func genFn0() int {
    v4 := GenS1{F0: 0, F1: 0, F2: (false || true)}
    v5 := GenS2{F0: ((-1 <= 1) && (653 > 65535)), F1: 1, F2: (("hello" + "\n\t") + "\n\t")}
    var v6 uint16 = 154
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float32 = (0.0 / (1.5 / 1.5))
        switch "\n\t" {
        case "":
            v9 := GenS1{F0: -128, F1: 993, F2: (-2.25 > float32(((0.0 + float64(0.1)) / (1e10 - (0.0 / ((1.5 + float64(0.0)) - 0.1))))))}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if false {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: 100, F1: 42, F2: ((!false) && true)}
            _ = fi10
            _ = v11
        }
        fmt.Println(v6)
        var v12 uint64 = 42
        _ = v12
    }
    _ = v4
    _ = v5
    _ = v6
    return 6
}
func genFn1(p0 uint8) int16 {
    var v13 uint = 2
    var v14 bool = false
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v13, v14, p0)
        _ = fi15
    }
    var v16 uint64 = 18446744073709551615
    fmt.Println(v14, p0)
    _ = p0
    _ = v13
    _ = v14
    _ = v16
    return 0
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint16 = 363
    v3 := make(map[string]string)
    v2 -= (v2 / ((v2) | 1))
    fmt.Println(v2)
    v17 := v3[("a" + "\n\t")]
    v17 = ("hello" + "foo bar")
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 uint16 = (v2 ^ v2)
        switch "foo bar" {
        case "a":
            v20 := GenS1{F0: 42, F1: -1, F2: (true || true)}
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v17)
    delete(v3, "")
    v2 -= v2
    switch ((("foo bar" + v17) + v17) + "a") {
    case "":
        v3[""] = ("hello" + ("" + "foo bar"))
    default:
        if (v17 >= v17) {
            v21 := v3[(v17 + (v17 + ("\n\t" + v17)))]
            if len(v1) > 0 {
                v1[0] = v21
            }
            v2 -= v2
            _ = v21
        }
    }
    switch genFn0() {
    case 32768:
        v22 := GenS3{F0: 0}
        _ = v22
    case 77:
        fmt.Println(v17, v17)
    case 128:
        v17 = (v17 + v17)
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                v2 -= 634
                _ = fi24
            }
            _ = fi23
        }
    }
    switch genFn0() {
    case 102:
        v17 = v17
    case 320:
        var v25 bool = (true && false)
        _ = v25
    }
    var v26 uint8 = 1
    var v27 uint16 = v2
    v28 := GenS2{F0: true, F1: (v26 % ((v26) | 1)), F2: "\n\t"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
