package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }

func main() {
    var v1 float32 = ((1e10 * -2.25) * float32((float64(((-2.25 - 1e10) * 1.5)) - 0.1)))
    v2 := make(map[string]bool)
    var v3 uint = 1
    v1 = v1
    v4 := GenS1{F0: -594, F1: -9223372036854775808, F2: ""}
    v3 = 594
    for fi5 := 0; fi5 < 18; fi5++ {
        if (255 <= 32767) {
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 uint64 = 0
                var v8 string = "foo bar"
                _ = fi6
                _ = v7
                _ = v8
            }
            switch -461 {
            case 32768:
                v9 := GenS1{F0: -9223372036854775808, F1: 0, F2: ("\n\t" + ("hello" + "hello"))}
                _ = v9
            case 128:
                var v10 uint8 = 255
                _ = v10
            case 2147483647:
                if false {
                    fmt.Println(v3, v1, v1)
                } else {
                    fmt.Println(fi5)
                }
            default:
                v1 = (float32(-2.25) + float32(1e10))
            }
        }
        var v11 uint16 = 104
        _ = fi5
        _ = v11
    }
    fmt.Println(v3, v3)
    _ = len(v2)
    var v12 uint8 = 161
    v13 := []int{2147483647}
    v1 -= v1
    v3 -= (v3 + 4294967295)
    switch ("hello" + "hello") {
    case "hello":
        var v14 int = 2147483647
        _ = v14
    case "\n\t":
        v15 := GenS1{F0: 0, F1: 0, F2: ("foo bar" + "")}
        _ = v15
    default:
        var v16 uint64 = 2
        _ = v16
    }
    v1 -= v1
    for fi17 := 0; fi17 < 14; fi17++ {
        v2[("a" + ("a" + ("" + "a")))] = false
        v12 += (v12 - v12)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
