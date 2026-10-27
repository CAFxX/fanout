package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 string, p1 int64) int {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int64 = (-(p1))
        v7 := []int{-1}
        v6 = p1
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 2147483647, F1: (31 == 75), F2: 2}
    p1 -= -9223372036854775808
    _ = p0
    _ = p1
    _ = v8
    return 32767
}
func genFn1(p0 uint32, p1 bool) string {
    switch "" {
    case "hello":
        var v9 bool = ((((-2.25 / 1e10) + 1.5) != (float64(float32(3.14159)) + 0.1)) || p1)
        _ = v9
    case "\n\t":
        p0 -= p0
    case "a":
        p1 = (2 > 136)
    default:
        p1 = p1
    }
    v10 := []int{32767}
    _ = p0
    _ = p1
    _ = v10
    return ""
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 1.5
    v3 := make(map[string]string)
    var v4 bool = false
    v2 -= v2
    if v4 {
        v2 -= 1.5
        v11 := []int{genFn0("a", 1), 32768, -9223372036854775808, 265}
        for fi12 := 0; fi12 < 4; fi12++ {
            if ((569 <= 2) || v4) {
                if v4 {
                    var v13 uint32 = 734
                    if (281 != 65535) {
                        for fi14 := 0; fi14 < 10; fi14++ {
                            if (0 >= 38) {
                                if len(v11) > 0 {
                                    v11[0] = genFn0(("\n\t" + "foo bar"), 0)
                                }
                                v15 := GenS2{F0: "\n\t", F1: 2}
                                var v16 string = (("a" + "hello") + "\n\t")
                                _ = v15
                                _ = v16
                            }
                            var v17 int32 = 1
                            _ = len(v3)
                            _ = fi14
                            _ = v17
                        }
                        v4 = v4
                        v3[(("\n\t" + ("" + "hello")) + ("\n\t" + "\n\t"))] = ""
                    }
                    for fi18 := 0; fi18 < 11; fi18++ {
                        switch "a" {
                        case "\n\t":
                            var v19 int32 = 179
                            _ = v19
                        default:
                            v3["a"] = "a"
                        }
                        v20 := GenS2{F0: "\n\t", F1: 42}
                        _ = fi18
                        _ = v20
                    }
                    _ = v13
                }
            }
            var v21 int32 = -39
            v22 := GenS2{F0: ("hello" + ("\n\t" + "foo bar")), F1: 155}
            _ = fi12
            _ = v21
            _ = v22
        }
        _ = v11
    } else {
        v1[128] = 32767
    }
    fmt.Println(v4, v4)
    var v23 uint8 = 0
    var v24 int = 2147483647
    if (1 <= 96) {
        v25 := []int{v24, (-(v24))}
        _ = v25
    }
    var v26 uint16 = 2
    _ = len(v1)
    var v27 uint16 = v26
    v28 := []int{(^(v24))}
    for fi29 := 0; fi29 < 4; fi29++ {
        for fi30 := 0; fi30 < 3; fi30++ {
            _ = len(v28)
            v26 = v26
            _ = fi30
        }
        _ = fi29
    }
    var v31 uint16 = 255
    if (!true) {
        v32 := v1[(v24 + genFn0("hello", -235))]
        delete(v3, ((("\n\t" + "\n\t") + "foo bar") + (("hello" + "a") + "\n\t")))
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
