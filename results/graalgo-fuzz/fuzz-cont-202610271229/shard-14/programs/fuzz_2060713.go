package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
func genFn0(p0 int8) int8 {
    var v5 int16 = 42
    fmt.Println(v5)
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 int = 127
    v2 := []uint32{2147483648, 0}
    var v3 int8 = -1
    v4 := make(map[string]bool)
    var v6 string = ("foo bar" + "foo bar")
    var v7 float64 = (0.1 * 1.5)
    if (((0.1 - ((0.0 / 1e10) / -2.25)) / (0.1 - 1e10)) > 0.1) {
        if (!(547 >= 0)) {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        } else {
            var v8 uint32 = 2147483648
            _ = v8
        }
        v7 -= (3.14159 / (v7 + ((float64(3.14159) * 0.1) * (0.1 + (v7 - ((1e10 * ((0.0 - float64(float32(v7))) - (-2.25 - v7))) + (v7 - (float64(0.1) * v7))))))))
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(v6, fi9)
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 bool = true
                v12 := GenS1{F0: (-(v3)), F1: 32767, F2: "hello"}
                if len(v2) > 0 {
                    v2[0] = 988
                }
                v3 += -128
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = len(v4)
            _ = fi9
        }
    }
    v13 := GenS2{F0: -255, F1: 465, F2: v3}
    fmt.Println(v6, v1)
    if false {
        var v14 int16 = 1000
        v15 := GenS1{F0: v3, F1: (-(v14)), F2: v6}
        if ((!true) || true) {
            switch (v6 + (v6 + (v6 + v6))) {
            case "hello":
                if (!(((32767 >= 65535) && ((-87 > -210) || true)) && ((1 == int32(v14)) && false))) {
                    var v16 uint16 = 191
                    _ = v16
                } else {
                    var v17 uint32 = 65535
                    _ = v17
                }
            }
            v14 += v14
        } else {
            v18 := v4[v6]
            _ = v18
        }
        _ = v14
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = 135
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        if false {
            v6 = (v6 + v6)
        } else {
            v6 = (v6 + (v6 + "\n\t"))
            v3 -= int8(v1)
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
