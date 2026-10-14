package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = (-2.25 - (float32((float64(((1.5 / (0.1 / 0.1)) * (0.0 * float32((0.1 / (1e10 * (1e10 / 1.5))))))) + 0.1)) + float32(-2.25)))
    var v3 uint64 = 4294967295
    v4 := []string{("foo bar" + "foo bar"), "", "\n\t"}
    v5 := []string{"foo bar", "a", "hello"}
    v3 += v3
    fmt.Println(v2)
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 42, F1: 639}
        switch ((("" + "hello") + (("\n\t" + "\n\t") + ((("" + ("" + ("" + "foo bar"))) + "") + "\n\t"))) + ("\n\t" + ("foo bar" + "foo bar"))) {
        case "foo bar":
            if true {
                var v8 uint8 = 191
                _ = v8
            }
        }
        var v9 float64 = 1.5
        _ = fi6
        _ = v7
        _ = v9
    }
    v3 = 18446744073709551615
    for fi10 := 0; fi10 < 3; fi10++ {
        v3 += (v3 / (((v3 << ((602) & 63))) | 1))
        _ = fi10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if ((32768 < fi11) || true) {
            var v12 int = fi11
            var v13 int16 = 32767
            var v14 string = ("" + ((("\n\t" + "a") + "hello") + "hello"))
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 float64 = (1.5 - 0.1)
            v15 -= (v15 / v15)
            var v16 uint16 = 2
            _ = v15
            _ = v16
        }
        var v17 int64 = 0
        _ = fi11
        _ = v17
    }
    _ = len(v5)
    fmt.Println(v2)
    v2 -= ((float32(float64(0.1)) / 1e10) + v2)
    fmt.Println(v2)
    v3 = v3
    var v18 float32 = v2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v19 int8 = 1
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
