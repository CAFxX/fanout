package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    v4 := make(map[int]uint32)
    v5 := []string{((("a" + "foo bar") + ("foo bar" + ("foo bar" + "hello"))) + ""), "a", ("" + ("" + "foo bar")), "foo bar"}
    var v6 float32 = (0.0 / (1.5 + ((0.1 + -2.25) / 1.0)))
    switch -1 {
    case 32768:
        v7 := GenS1{F0: 0, F1: -1}
        _ = v7
    case 32767:
        if len(v5) > 0 {
            v5[0] = ("\n\t" + ((("\n\t" + (("\n\t" + "a") + (("" + "a") + "foo bar"))) + (("foo bar" + ((("a" + "foo bar") + "a") + "foo bar")) + "\n\t")) + "hello"))
        }
    case 42:
        var v8 uint8 = 128
        _ = v8
    default:
        v6 += float32(0.1)
    }
    v6 = v6
    for fi9 := 0; fi9 < 17; fi9++ {
        _ = len(v5)
        _ = fi9
    }
    v10 := GenS1{F0: 751, F1: 32767}
    if (0 >= v3) {
        var v11 uint16 = 747
        _ = v11
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v2[1] = 42
            var v14 int8 = 127
            var v15 int = 1
            v5 = append(v5, "")
            if true {
                var v16 int32 = (v3 % (((v3 & (v3 * v3))) | 1))
                _ = v16
            } else {
                fmt.Println(v15, fi13, v3)
                var v17 int = (fi12 % (((fi12 / ((fi12) | 1))) | 1))
                switch -325 {
                case 4294967295:
                    v18 := GenS2{F0: 18446744073709551615, F1: 2}
                    _ = v18
                }
                v3 = (v3 & 0)
                _ = v17
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    v5 = append(v5, ("\n\t" + (("" + "a") + "hello")))
    fmt.Println(v6)
    var v19 int16 = -283
    var v20 float64 = float64((v6 / (0.0 / v6)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
