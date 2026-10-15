package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 uint }
func genFn0() int32 {
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: "hello", F1: "\n\t", F2: "hello"}
        _ = fi5
        _ = v6
    }
    _ = v4
    return -574
}
func genFn1() uint {
    var v7 uint32 = 0
    v8 := []int{285, 9223372036854775807, 423, 325}
    v7 = v7
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(v7, v7)
        _ = fi9
    }
    _ = v7
    _ = v8
    return 273
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = float32(-2.25)
    v3 := make(map[string]int)
    for fi10 := 0; fi10 < 15; fi10++ {
        v2 -= (v2 * (v2 * ((float32(3.14159) / v2) + (float32((float64(-2.25) * ((3.14159 - 1e10) * (float64(0.1) / 3.14159)))) + v2))))
        fmt.Println(fi10)
        _ = fi10
    }
    v2 += v2
    switch "a" {
    case "a":
        v2 += float32(-2.25)
    case "\n\t":
        v11 := GenS2{F0: 255}
        _ = v11
    default:
        switch "\n\t" {
        case "foo bar":
            v2 = ((v2 + (v2 + (1.5 * v2))) * float32((0.1 * 1e10)))
        case "a":
            var v12 int32 = 1
            _ = v12
        case "\n\t":
            switch "a" {
            case "hello":
                var v13 uint32 = 2147483648
                _ = v13
            default:
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 bool = (false || true)
                    v15 = false
                    var v16 int32 = -2147483648
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
        default:
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 uint32 = 2
                if (253 >= 65535) {
                    var v19 uint32 = v18
                    v19 = (v18 | (v18 ^ (v18 >> ((363) & 63))))
                    _ = v19
                } else {
                    v2 = ((1e10 / v2) - 0.1)
                }
                v20 := []int{(fi17 >> ((65535) & 63)), -983, (fi17 * (fi17 ^ (fi17 * (fi17 << ((275) & 63))))), 9223372036854775807}
                _ = fi17
                _ = v18
                _ = v20
            }
        }
    }
    if (255 == 1) {
        v21 := GenS1{F0: "\n\t", F1: ("\n\t" + ""), F2: (("foo bar" + (("\n\t" + "foo bar") + "a")) + "")}
        switch "\n\t" {
        case "foo bar":
            var v22 string = ""
            _ = v22
        }
        _ = v21
    }
    v3[""] = -2147483648
    v23 := GenS2{F0: 0}
    v24 := []int{2147483647, 1, 2147483647, -9223372036854775808}
    var v25 int = 9223372036854775807
    var v26 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
