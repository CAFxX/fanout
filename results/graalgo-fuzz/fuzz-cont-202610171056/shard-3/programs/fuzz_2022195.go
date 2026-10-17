package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{("foo bar" + (("a" + "foo bar") + "\n\t"))}
    var v3 uint = 9
    var v4 float64 = (1.5 * 3.14159)
    var v5 float64 = (0.1 - v4)
    var v6 uint8 = 0
    var v7 int = 42
    v7 += int(v3)
    var v8 int16 = 42
    v9 := GenS1{F0: v8, F1: 2}
    v10 := GenS2{F0: v3, F1: 611, F2: -1}
    v4 -= float64((float32((float64(0.1) + 1e10)) - (((float32(v4) * (0.0 + (-2.25 + 0.1))) + 0.1) - 0.0)))
    var v11 int8 = -128
    v12 := GenS1{F0: v8, F1: 65535}
    v6 = 255
    v13 := GenS2{F0: (^(v3)), F1: 594, F2: 42}
    if ((v8 | 32767) >= -32768) {
        v8 = (v8 ^ 1)
    } else {
        switch (v7 << (((v3 / (((^(v3))) | 1))) & 63)) {
        case 534:
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint = v3
                var v16 bool = (!true)
                _ = fi14
                _ = v15
                _ = v16
            }
        case 224:
            v3 -= 478
        default:
            for fi17 := 0; fi17 < 12; fi17++ {
                if (1 != -1) {
                    fmt.Println(v6, v3, v3)
                    delete(v1, "hello")
                    var v18 uint16 = 65535
                    _ = v18
                }
                _ = fi17
            }
        }
        fmt.Println(v6, v8, v6)
    }
    v19 := v1["a"]
    v20 := GenS1{F0: -32768, F1: 255}
    v21 := v1["hello"]
    if (!((-737 <= 1) && (2 != 0))) {
        var v22 bool = (false && true)
        _ = v22
    } else {
        if true {
            v6 += v6
            v23 := GenS1{F0: v8, F1: 1}
            _ = v23
        }
        v24 := GenS2{F0: 1, F1: 42, F2: 248}
        _ = v24
    }
    v11 += v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
