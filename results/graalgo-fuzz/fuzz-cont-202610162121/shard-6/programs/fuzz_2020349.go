package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int, p2 int) bool {
    var v7 float32 = (3.14159 + 1.5)
    p1 += 127
    var v8 int8 = 127
    var v9 int16 = 26
    var v10 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 uint16 = 2
    var v4 uint16 = v3
    v5 := []int{-854, 0, 0}
    v6 := make(map[uint32]bool)
    for fi11 := 0; fi11 < 4; fi11++ {
        v4 = v4
        _ = fi11
    }
    v4 += 516
    var v12 int32 = -344
    switch ("a" + (((("\n\t" + "") + "\n\t") + "\n\t") + "foo bar")) {
    case "\n\t":
        if (53 > 127) {
            v13 := GenS1{F0: false, F1: 250, F2: 812}
            _ = len(v1)
            _ = v13
        } else {
            var v14 float32 = -2.25
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v15 := v2[("" + "hello")]
            _ = v14
            _ = v15
        }
    case "a":
        v5 = append(v5, 128)
    case "foo bar":
        var v16 bool = false
        _ = v16
    default:
        if ((32767 < 1) && (0 != (v12 / ((v12) | 1)))) {
            v17 := GenS1{F0: false, F1: 2, F2: 4294967295}
            _ = v17
        } else {
            v3 = v3
            switch ("" + ("" + "")) {
            case "\n\t":
                var v18 float64 = 1e10
                _ = v18
            default:
                v12 += v12
            }
        }
    }
    v12 = (^(v12))
    v3 = v3
    v19 := GenS2{F0: 78}
    var v20 int = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
