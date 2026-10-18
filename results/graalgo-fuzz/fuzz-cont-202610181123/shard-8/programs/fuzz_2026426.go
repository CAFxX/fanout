package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }
func genFn0(p0 int) bool {
    var v4 int32 = 0
    v5 := []int{}
    if len(v5) > 0 {
        v5[0] = (p0 & (p0 - p0))
    }
    if (false || true) {
        v4 = (v4 << ((18446744073709551615) & 63))
    }
    _ = p0
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 string = "a"
    var v2 float64 = (0.0 - float64((((1e10 - 0.1) - 3.14159) * (0.0 / float32((1.5 / 1.0))))))
    v3 := make(map[string]uint32)
    var v6 float32 = 1e10
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v1 = ""
            var v9 int16 = 42
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 = ""
            v3[v1] = 853
            delete(v3, ((v1 + (("a" + "\n\t") + v1)) + "a"))
            _ = fi10
        }
        fmt.Println(fi7, v6)
        _ = fi7
    }
    v11 := GenS1{F0: 462, F1: true, F2: 2}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v2, v2, v6)
    case "a":
        switch v1 {
        case "hello":
            v2 = v2
        case "foo bar":
            fmt.Println(v2)
        default:
            switch v1 {
            case "foo bar":
                v1 = v1
            case "hello":
                v6 = 1e10
            case "a":
                v6 -= v6
            }
        }
    }
    var v12 uint32 = 154
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 float32 = float32((((v2 * v2) - float64(3.14159)) + 1e10))
        v14 -= (v14 * v6)
        _ = fi13
        _ = v14
    }
    if (!false) {
        var v15 int8 = -67
        fmt.Println(v12)
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v6)
        v17 := []int{499, 9223372036854775807, 712, 42}
        v1 = v1
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: v12, F1: false, F2: 89}
    var v19 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
