package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int }
func genFn0(p0 int8, p1 int16, p2 string) int8 {
    v4 := GenS1{F0: 128, F1: 114, F2: 42}
    if (false && ((0.1 + (0.1 / 1.0)) != float32(float64(0.0)))) {
        var v5 float64 = 0.1
        v5 += 1e10
        _ = v5
    } else {
        v6 := GenS1{F0: -505, F1: 128, F2: 0}
        switch int(p0) {
        case 4294967295:
            p0 = 54
        }
        v7 := []int{127, 239, 1}
        _ = v6
        _ = v7
    }
    if (!(1e10 == float64(3.14159))) {
        var v8 uint64 = 0
        var v9 bool = (!true)
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int32 = 414
            if v9 {
                switch (fi10 >> ((1) & 63)) {
                case 128:
                    var v12 uint32 = 234
                    _ = v12
                default:
                    fmt.Println(fi10, v8)
                }
                v8 = (v8 / ((v8) | 1))
                v13 := []int{fi10, (fi10 ^ 2147483647), 32768}
                _ = v13
            } else {
                if (!true) {
                    fmt.Println(fi10, v9, v11)
                } else {
                    v14 := GenS1{F0: fi10, F1: 2, F2: fi10}
                    var v15 uint8 = 2
                    _ = v14
                    _ = v15
                }
            }
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    var v16 uint8 = 56
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v4
    return int8(v16)
}

func main() {
    var v1 float64 = ((float64(-2.25) + 3.14159) / (0.1 - ((0.0 * -2.25) + -2.25)))
    var v2 uint64 = 0
    v3 := make(map[int]int)
    if (!(((("a" + "foo bar") + "") + (((("foo bar" + ("a" + "hello")) + ("" + ("\n\t" + "\n\t"))) + "hello") + "hello")) <= "")) {
        v2 = 516
    }
    v3[-67] = 348
    for fi17 := 0; fi17 < 4; fi17++ {
        v2 -= 276
        switch (("a" + "a") + "foo bar") {
        case "hello":
            if (float32((v1 - 0.0)) != 3.14159) {
                var v18 int = fi17
                _ = v18
            }
        default:
            var v19 uint64 = 42
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: fi20, F1: 118, F2: fi20}
        v3[-499] = fi20
        _ = fi20
        _ = v21
    }
    var v22 uint16 = 1
    v23 := GenS1{F0: 2147483647, F1: 0, F2: 572}
    var v24 uint32 = 4294967295
    v22 = 32767
    v22 = 0
    v22 -= (v22 ^ 42)
    v25 := v3[-842]
    var v26 uint8 = 169
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
