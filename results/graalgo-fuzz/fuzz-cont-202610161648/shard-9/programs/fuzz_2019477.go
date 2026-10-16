package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 int16, p1 int16, p2 string) int32 {
    p2 = p2
    v7 := []int{-9223372036854775808, 128, 127, 32767}
    p2 = "\n\t"
    var v8 uint8 = 128
    var v9 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []string{"foo bar", (("\n\t" + "hello") + "\n\t"), ("" + ("foo bar" + ("" + "a")))}
    var v2 int16 = 1000
    v3 := make(map[string]int)
    var v4 float64 = 1.5
    var v5 uint16 = 42
    v6 := make(map[int]int)
    var v10 uint = 1
    fmt.Println(v5, v5, v5)
    var v11 uint32 = uint32(v5)
    var v12 uint64 = 546
    v11 += v11
    if false {
        switch -9223372036854775808 {
        case 32768:
            _ = len(v3)
        }
    } else {
        v12 -= v12
        switch "" {
        case "hello":
            if (1 < -1) {
                if (!false) {
                    var v13 uint8 = 255
                    fmt.Println(v5, v12, v12)
                    _ = v13
                } else {
                    if (!true) {
                        if ((((0.0 * 1e10) / (1e10 + 1e10)) < (((1.5 - (-2.25 - 0.1)) * 3.14159) * -2.25)) || ((v12 ^ v12) > v12)) {
                            var v14 uint32 = 1
                            v15 := GenS1{F0: (v14 << ((v10) & 63))}
                            var v16 int32 = genFn0(693, (v2 >> ((65535) & 63)), ("" + "hello"))
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        v10 -= 1
                    }
                    v17 := GenS1{F0: v11}
                    fmt.Println(v2)
                    _ = v17
                }
                var v18 int16 = v2
                v19 := GenS1{F0: v11}
                _ = v18
                _ = v19
            }
        default:
            v10 = 2
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 bool = false
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := GenS2{F0: 1, F1: 172, F2: 868}
                var v24 int = (fi20 | (^(fi22)))
                _ = fi22
                _ = v23
                _ = v24
            }
            fmt.Println(v11, v4)
            _ = fi20
            _ = v21
        }
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint = 42
        _ = fi25
        _ = v26
    }
    v27 := GenS3{F0: false, F1: v5}
    switch (((("a" + "foo bar") + "foo bar") + "\n\t") + "foo bar") {
    case "hello":
        v12 = uint64(v5)
    case "":
        var v28 uint = v10
        _ = v28
    default:
        var v29 uint64 = uint64(v5)
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
