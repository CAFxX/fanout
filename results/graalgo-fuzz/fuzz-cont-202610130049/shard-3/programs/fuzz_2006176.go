package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 string, p1 uint64, p2 int8) uint {
    var v7 string = p0
    fmt.Println(v7, p1)
    v8 := GenS1{F0: 359}
    v7 = "\n\t"
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 int64) string {
    v9 := []int{513, 4294967295, -2147483648}
    v9 = append(v9, 9223372036854775807)
    var v10 uint8 = 238
    switch (((("" + ("a" + "")) + "hello") + "foo bar") + "a") {
    case "":
        v11 := GenS1{F0: 547}
        _ = v11
    case "hello":
        v10 = 255
    }
    _ = p0
    _ = v10
    _ = v9
    return "\n\t"
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 int32 = 1
    v5 := []int{-152}
    var v6 int8 = 100
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (((((("\n\t" + "\n\t") + ("foo bar" + "")) + "foo bar") + (((("" + "foo bar") + ("\n\t" + ("foo bar" + "hello"))) + "a") + "\n\t")) + ("a" + "hello")) + "\n\t") {
        case "foo bar":
            var v13 uint32 = 4294967295
            _ = v13
        default:
            if (32767 <= 0) {
                switch 4294967295 {
                case -406:
                    v4 = int32(v6)
                case 32768:
                    if (1.5 < ((-2.25 - 1.5) * (-2.25 / 1.0))) {
                        v6 -= (-(v6))
                        if (127 >= 224) {
                            _ = len(v2)
                        } else {
                            var v14 uint16 = 2
                            var v15 uint8 = 138
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS1{F0: 146}
                        _ = v16
                    }
                default:
                    var v17 int64 = v1
                    _ = v17
                }
                v18 := v2[2147483648]
                _ = v18
            }
        }
        _ = fi12
    }
    v4 -= (^(v4))
    v19 := v2[2147483647]
    fmt.Println(v3, v1, v3)
    var v20 int16 = 923
    v3 = v3
    v21 := GenS2{F0: v19, F1: v19}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v22 bool = (!("\n\t" == "a"))
    fmt.Println(v3, v3, v22)
    v2[65535] = (v19 / (((v19 & (v19 / ((0) | 1)))) | 1))
    var v23 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
