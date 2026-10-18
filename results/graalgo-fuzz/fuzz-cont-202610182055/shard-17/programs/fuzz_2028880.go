package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
func genFn0(p0 int, p1 uint8, p2 int8) int8 {
    p2 += p2
    v7 := GenS1{F0: "hello", F1: p2, F2: (p1 - 180)}
    v8 := GenS1{F0: (("hello" + ("hello" + "foo bar")) + "\n\t"), F1: (p2 >> ((749) & 63)), F2: p1}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -55
}

func main() {
    var v1 float64 = (-2.25 + 0.1)
    var v2 int = -930
    var v3 int16 = -370
    v4 := []string{("" + "hello"), "hello", "foo bar"}
    v5 := []string{}
    v6 := make(map[uint32]bool)
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9)
        for fi10 := 0; fi10 < 1; fi10++ {
            v1 += -2.25
            _ = fi10
        }
        _ = fi9
    }
    var v11 int8 = 1
    fmt.Println(v2)
    if (566 <= 0) {
        switch "\n\t" {
        case "hello":
            v5 = append(v5, "a")
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = (("hello" + "foo bar") + "a")
            }
        case "a":
            for fi12 := 0; fi12 < 19; fi12++ {
                if (0.0 != 3.14159) {
                    v5 = append(v5, ("" + "foo bar"))
                    var v13 int8 = 24
                    _ = v13
                } else {
                    _ = len(v5)
                    var v14 bool = true
                    v11 = genFn0(128, 222, v11)
                    _ = v14
                }
                var v15 int32 = -814
                if true {
                    v3 -= v3
                    delete(v6, 2147483648)
                    v1 -= (v1 / ((v1 - float64(1.5)) - (1.5 - v1)))
                } else {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        fmt.Println(v11)
                        _ = fi16
                    }
                }
                _ = fi12
                _ = v15
            }
        }
        fmt.Println(v1)
        fmt.Println(v11, v1)
    } else {
        var v17 int16 = -325
        v2 -= -2147483648
        _ = v17
    }
    fmt.Println(v3, v1)
    fmt.Println(v3, v3, v2)
    v11 = (v11 | (v11 % (((v11 ^ 0)) | 1)))
    for fi18 := 0; fi18 < 18; fi18++ {
        if (100 >= genFn0((fi18 % ((v2) | 1)), 1, v11)) {
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(fi18)
                v11 = (v11 ^ v11)
                _ = fi19
            }
            fmt.Println(v1, fi18, v3)
        }
        _ = len(v4)
        v20 := GenS1{F0: "", F1: v11, F2: 0}
        _ = fi18
        _ = v20
    }
    fmt.Println(v11)
    var v21 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
