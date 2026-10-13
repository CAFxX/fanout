package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int64 }
func genFn0(p0 int8, p1 int) int {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    p0 -= p0
    v10 := GenS1{F0: p1, F1: 658}
    if (false || (2147483647 <= 2147483647)) {
        v11 := GenS1{F0: (p1 ^ p1), F1: 826}
        _ = v11
    } else {
        v12 := GenS1{F0: 9223372036854775807, F1: 0}
        switch ("foo bar" + ((("" + "\n\t") + ("" + ((((("\n\t" + "a") + ((("foo bar" + "\n\t") + "a") + "")) + "\n\t") + "") + ("foo bar" + ("\n\t" + "a"))))) + ((("hello" + "\n\t") + ("" + "")) + "a"))) {
        case "hello":
            if ((!(!true)) || true) {
                var v13 float32 = -2.25
                v13 -= 1.5
                if false {
                    p1 += -878
                    var v14 string = (("" + "hello") + "a")
                    _ = v14
                } else {
                    v15 := GenS1{F0: (p1 & -387), F1: 42}
                    var v16 int16 = -1
                    if true {
                        var v17 uint16 = 32767
                        _ = v17
                    }
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
        default:
            for fi18 := 0; fi18 < 2; fi18++ {
                p0 += p0
                switch -2147483648 {
                case 569:
                    p1 += (fi18 | fi18)
                case 32767:
                    p0 = p0
                case -726:
                    p0 -= p0
                default:
                    v19 := []int{(p1 / ((fi18) | 1)), 128, (p1 - p1)}
                    _ = v19
                }
                var v20 int32 = 42
                _ = fi18
                _ = v20
            }
        }
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return p1
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    v3 := []string{("" + ("a" + ("hello" + "a"))), ("hello" + "")}
    var v4 int8 = 127
    var v5 int = 2147483647
    var v6 float32 = 1e10
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 int = -220
        v5 = 1
        v23 := GenS1{F0: (v22 / (((fi21 | (fi21 << ((65535) & 63)))) | 1)), F1: 606}
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v5 -= v5
        v25 := GenS1{F0: v5, F1: 266}
        var v26 uint = 108
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 bool = (v6 > float32(((0.1 + v1) * ((v1 / (3.14159 + v1)) * v1))))
        _ = fi27
        _ = v28
    }
    switch "foo bar" {
    case "\n\t":
        v1 -= v1
    case "":
        fmt.Println(v5, v4, v1)
    }
    switch -9223372036854775808 {
    case 128:
        for fi29 := 0; fi29 < 8; fi29++ {
            v30 := GenS2{F0: 255, F1: fi29, F2: 0}
            v31 := v2["a"]
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    v2["a"] = 42
    delete(v2, "foo bar")
    v32 := GenS3{F0: 42, F1: 127, F2: 182}
    fmt.Println(v1, v1)
    _ = len(v2)
    var v33 float64 = (3.14159 - 0.0)
    switch "\n\t" {
    case "hello":
        v34 := GenS3{F0: 453, F1: v4, F2: -101}
        _ = v34
    case "":
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS3{F0: 65535, F1: (v4 / ((100) | 1)), F2: 9223372036854775807}
            _ = fi35
            _ = v36
        }
    }
    v37 := GenS4{F0: 1, F1: 69, F2: 1}
    var v38 float64 = (0.0 - v33)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
