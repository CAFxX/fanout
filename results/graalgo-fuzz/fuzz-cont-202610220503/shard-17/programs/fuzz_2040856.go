package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int8; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 string = ""
    var v3 int = -19
    var v4 uint32 = 0
    var v5 uint8 = 255
    var v6 int = v3
    v7 := []int{(v6 - v6), v3, v6}
    fmt.Println(v6, v2, v3)
    if false {
        if true {
            if (float32(((0.1 * float64(0.1)) / 1e10)) == (-2.25 - (((((float32((-2.25 + (0.0 - float64(1.5)))) / ((3.14159 / (((-2.25 + -2.25) * float32(0.1)) - ((float32(0.0) * ((0.1 - 0.1) / (1e10 - 0.0))) + float32(-2.25)))) - (0.1 - 0.1))) * 1e10) * (-2.25 - 1e10)) * 0.1) - 1.5))) {
                for fi8 := 0; fi8 < 2; fi8++ {
                    switch (v2 + "\n\t") {
                    case "a":
                        v3 = (-(v6))
                    }
                    var v9 float64 = ((float64(-2.25) + 0.1) + (1.5 - ((3.14159 - ((float64(1.5) / 0.1) - float64(1e10))) + 0.1)))
                    _ = fi8
                    _ = v9
                }
                if false {
                    v4 -= (^(v4))
                    for fi10 := 0; fi10 < 1; fi10++ {
                        var v11 int32 = 0
                        v5 += v5
                        var v12 uint32 = v4
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
            } else {
                v1[""] = true
            }
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v6 = v3
            v7 = append(v7, 32767)
            if (42 > -987) {
                var v14 int16 = int16(fi13)
                var v15 int = fi13
                _ = v14
                _ = v15
            }
            _ = fi13
        }
    }
    fmt.Println(v2)
    for fi16 := 0; fi16 < 3; fi16++ {
        v6 = v3
        _ = fi16
    }
    v17 := GenS1{F0: 42}
    v18 := GenS1{F0: v5}
    switch ("a" + (v2 + (v2 + "a"))) {
    case "a":
        if (1 < 1) {
            v1[("" + (v2 + ("foo bar" + "")))] = true
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            v19 := GenS2{F0: (^(v5)), F1: 42, F2: v2}
            v5 = 117
            _ = v19
        }
    case "foo bar":
        v6 += (v6 % ((v3) | 1))
    default:
        var v20 int64 = -9223372036854775808
        _ = v20
    }
    v21 := v1[v2]
    v3 = v6
    var v22 int = (v6 * v3)
    v6 = v6
    if (!true) {
        var v23 uint64 = 49
        v21 = v21
        var v24 uint = 871
        _ = v23
        _ = v24
    } else {
        v25 := GenS2{F0: v5, F1: -128, F2: "foo bar"}
        if v21 {
            var v26 string = (("a" + v2) + "hello")
            _ = v26
        }
        _ = len(v7)
        _ = v25
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
