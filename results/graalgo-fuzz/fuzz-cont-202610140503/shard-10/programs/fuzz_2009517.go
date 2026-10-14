package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 uint32) bool {
    v6 := GenS1{F0: ("\n\t" + ("hello" + "")), F1: -699}
    switch "\n\t" {
    case "a":
        p0 = p0
    default:
        switch 877 {
        case 852:
            v7 := GenS1{F0: ("" + ("\n\t" + ("\n\t" + "\n\t"))), F1: -308}
            _ = v7
        case 32768:
            var v8 uint32 = p0
            _ = v8
        case -2147483648:
            if (-2.25 <= ((0.0 - (float32((1.5 - ((((1.5 + (3.14159 - 1e10)) - 0.1) * (1e10 + float64(float32(3.14159)))) * ((-2.25 - 0.0) - 3.14159)))) - 3.14159)) - ((((((float32(((-2.25 / -2.25) - 3.14159)) / (1e10 + 3.14159)) - 1.5) / 1.0) + -2.25) + 3.14159) + 1e10))) {
                var v9 float32 = 0.1
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 14; fi10++ {
                switch "hello" {
                case "":
                    var v11 int32 = -2147483648
                    _ = v11
                case "hello":
                    v12 := GenS1{F0: "foo bar", F1: -32768}
                    _ = v12
                }
                var v13 uint = 988
                _ = fi10
                _ = v13
            }
        }
    }
    fmt.Println(p0)
    var v14 int16 = -354
    _ = p0
    _ = v14
    _ = v6
    return false
}

func main() {
    var v1 uint = 194
    var v2 int16 = 91
    var v3 int32 = -1
    var v4 int32 = 42
    v5 := make(map[string]string)
    var v15 string = ""
    v16 := []int{570, -9223372036854775808, -161}
    v2 -= 1000
    var v17 int16 = v2
    _ = len(v5)
    v3 += (v3 ^ v4)
    var v18 int32 = -911
    var v19 uint16 = 2
    if (!true) {
        var v20 uint32 = 42
        _ = v20
    } else {
        v21 := GenS1{F0: "hello", F1: v2}
        _ = v21
    }
    v18 = v4
    if (42 >= uint64(v19)) {
        switch v15 {
        case "":
            if (-803 >= 232) {
                for fi22 := 0; fi22 < 2; fi22++ {
                    v23 := GenS1{F0: ("\n\t" + (v15 + "hello")), F1: v2}
                    v24 := GenS1{F0: ("a" + ("hello" + v15)), F1: (v17 * (-(v17)))}
                    v25 := GenS1{F0: "hello", F1: 32767}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 uint32 = 65535
                    var v28 float64 = 1e10
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                v17 += int16(v3)
            }
        case "foo bar":
            v29 := GenS2{F0: 2, F1: false}
            _ = v29
        default:
            v2 = v17
        }
        var v30 bool = ((1e10 / float64(-2.25)) > ((1.5 / 1.5) - (3.14159 * 3.14159)))
        _ = v30
    }
    if ((2147483648 != 179) && false) {
        v18 = 808
        var v31 int = 128
        _ = v31
    } else {
        var v32 string = v15
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
