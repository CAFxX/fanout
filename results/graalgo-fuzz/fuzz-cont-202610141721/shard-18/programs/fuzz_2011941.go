package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 int32, p1 bool, p2 string) string {
    fmt.Println(p2, p2, p0)
    v5 := []int{32768, 32768, 42, 9223372036854775807}
    if p1 {
        for fi6 := 0; fi6 < 14; fi6++ {
            v5 = append(v5, (fi6 & fi6))
            if len(v5) > 0 {
                v5[0] = (fi6 + fi6)
            }
            _ = fi6
        }
        var v7 int = 32767
        _ = v7
    } else {
        if (127 != 105) {
            var v8 uint = 255
            if (944 <= 0) {
                v9 := GenS1{F0: 106, F1: 1, F2: 1}
                _ = v9
            }
            if len(v5) > 0 {
                v5[0] = -14
            }
            _ = v8
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 int32, p1 int64, p2 string) uint16 {
    for fi10 := 0; fi10 < 16; fi10++ {
        p2 = p2
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int = (fi11 * (fi10 ^ (fi10 - (-(fi11)))))
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v13 float32 = 3.14159
    fmt.Println(v13)
    var v14 float64 = (0.1 - float64((v13 / (v13 + ((v13 / v13) + 3.14159)))))
    fmt.Println(v1)
    var v15 int8 = -128
    v14 = float64(v13)
    v16 := GenS2{F0: (v1 / ((v1) | 1)), F1: "hello", F2: "\n\t"}
    v1 += -32768
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS1{F0: (v15 + 12), F1: 0, F2: 42}
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: v1, F1: "hello", F2: (("foo bar" + "a") + ("foo bar" + ""))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
