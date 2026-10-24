package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 uint8) int16 {
    p0 = (p0 / ((p0) | 1))
    var v4 string = ("a" + ("" + ("hello" + ("\n\t" + "a"))))
    _ = p0
    _ = v4
    return -1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint = 675
    var v3 float32 = 1.5
    v5 := []int{491, 32767}
    if false {
        v2 += (v2 ^ v2)
    } else {
        v6 := GenS1{F0: -118, F1: 136, F2: true}
        v2 = (v2 + v2)
        _ = v6
    }
    fmt.Println(v2)
    if ((float64(v3) >= ((0.0 - 0.0) - 1e10)) || (127 != -534)) {
        fmt.Println(v1, v1)
    }
    v7 := GenS1{F0: -18, F1: 255, F2: (true && false)}
    if len(v5) > 0 {
        v5[0] = -360
    }
    switch "hello" {
    case "\n\t":
        v8 := GenS1{F0: 32, F1: 1, F2: (false || false)}
        _ = v8
    case "":
        var v9 string = "foo bar"
        _ = v9
    case "a":
        if true {
            if false {
                v10 := GenS1{F0: 127, F1: 88, F2: (18446744073709551615 != v1)}
                v11 := GenS1{F0: -1, F1: 74, F2: false}
                _ = v10
                _ = v11
            }
            var v12 bool = false
            var v13 uint32 = 682
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                v1 = (v1 ^ v1)
                _ = fi14
            }
        }
    default:
        v1 = (v1 + v1)
    }
    v15 := GenS2{F0: 42, F1: false}
    fmt.Println(v3, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v7
}
