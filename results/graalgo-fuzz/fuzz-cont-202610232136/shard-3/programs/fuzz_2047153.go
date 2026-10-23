package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }

func main() {
    var v1 uint16 = 0
    var v2 uint16 = v1
    var v3 float64 = 1.5
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    if (false || (("a" + "foo bar") < "\n\t")) {
        var v6 int32 = 0
        for fi7 := 0; fi7 < 20; fi7++ {
            v3 = (((3.14159 - v3) * (((v3 + float64(1e10)) - (((v3 * float64(1.5)) - (1.5 + ((3.14159 / 0.1) - v3))) - v3)) * (v3 + float64(float32(v3))))) + 0.0)
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 uint16 = (v2 * 42)
                var v10 uint32 = 65535
                v9 -= 2
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        var v11 bool = (!false)
        _ = v11
        _ = v6
    }
    fmt.Println(v1, v1)
    for fi12 := 0; fi12 < 12; fi12++ {
        switch fi12 {
        case 42:
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint64 = 42
                _ = fi13
                _ = v14
            }
        case -4:
            var v15 float64 = 1e10
            _ = v15
        case 4294967295:
            v16 := GenS1{F0: (("\n\t" + "foo bar") + ""), F1: 963, F2: -9223372036854775808}
            _ = v16
        default:
            var v17 int = (fi12 - -1)
            _ = v17
        }
        _ = fi12
    }
    v18 := []int{128, -980, 0}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v19 string = "foo bar"
    v3 -= 1e10
    if true {
        v20 := GenS1{F0: v19, F1: 4294967295, F2: 0}
        v2 = (v1 ^ (v1 << ((4294967295) & 63)))
        _ = v20
    } else {
        switch "hello" {
        case "a":
            v1 = v1
        }
        _ = len(v4)
        v3 += v3
    }
    var v21 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
