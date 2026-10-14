package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 string, p1 string) string {
    if (42 < -9223372036854775808) {
        var v4 int = -653
        p0 = p1
        p1 = p1
        _ = v4
    } else {
        p0 = "foo bar"
    }
    p1 = "hello"
    _ = p0
    _ = p1
    return ((p1 + "hello") + p1)
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]bool)
    var v3 float64 = (1e10 * 3.14159)
    v1 -= v1
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(v3)
        var v6 string = "foo bar"
        _ = fi5
        _ = v6
    }
    if (255 != 62) {
        v3 -= (v3 - v3)
        var v7 uint = 4294967295
        fmt.Println(v3, v3)
        _ = v7
    }
    var v8 uint64 = (v1 + (v1 / ((v1) | 1)))
    fmt.Println(v8, v3, v1)
    if true {
        if true {
            v3 = v3
            switch ("" + "foo bar") {
            case "":
                if (!false) {
                    v3 += ((v3 - v3) * 0.0)
                }
            }
        } else {
            if (-128 != 1) {
                v9 := v2[9223372036854775807]
                if v9 {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v8 = 0
                        v11 := GenS1{F0: 0, F1: -1}
                        v1 = 338
                        _ = fi10
                        _ = v11
                    }
                    v8 = 711
                    v1 = 483
                }
                _ = v9
            }
        }
        var v12 int32 = 771
        _ = v12
    } else {
        _ = len(v2)
        fmt.Println(v8)
    }
    if (840 > 42) {
        v13 := GenS1{F0: 2, F1: -32768}
        v14 := GenS1{F0: 82, F1: 32767}
        _ = v13
        _ = v14
    }
    v15 := []int{32767}
    v16 := GenS1{F0: 215, F1: 1}
    v8 -= (v8 & 2)
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 string = "foo bar"
            _ = fi18
            _ = v19
        }
        v2[0] = false
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
