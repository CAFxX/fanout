package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0() int {
    v6 := []int{978, -269, 170, 32768}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    var v3 bool = true
    v4 := make(map[string]bool)
    var v5 uint16 = 1
    v5 = v5
    v3 = v3
    switch 2147483647 {
    case 0:
        if (-2.25 > float64(float32((0.1 + (0.0 - -2.25))))) {
            delete(v4, "\n\t")
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 float32 = 0.1
                var v9 int32 = 154
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 uint32 = 310
            _ = v10
        }
    case 9223372036854775807:
        switch -125 {
        case 32768:
            v1[(((("a" + ("" + "hello")) + "\n\t") + "foo bar") + ("\n\t" + "foo bar"))] = "foo bar"
        case 251:
            var v11 int32 = 1
            _ = v11
        }
    default:
        v12 := GenS1{F0: 139, F1: 248}
        _ = v12
    }
    if v3 {
        v13 := []int{-123, 32768, -2147483648}
        if true {
            var v14 float32 = 3.14159
            _ = v14
        }
        var v15 uint64 = (v2 * 382)
        _ = v13
        _ = v15
    } else {
        switch "\n\t" {
        case "foo bar":
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int = fi16
                _ = fi16
                _ = v17
            }
        }
    }
    v18 := GenS1{F0: -1, F1: uint(v5)}
    v2 = (v2 | (v2 + v2))
    var v19 int16 = 1
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
