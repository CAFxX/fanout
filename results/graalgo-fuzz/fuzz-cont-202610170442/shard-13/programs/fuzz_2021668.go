package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
func genFn0(p0 int16) uint16 {
    p0 -= 0
    var v7 bool = true
    switch "\n\t" {
    case "\n\t":
        v8 := GenS1{F0: ((!true) && ((1.5 / (1e10 / (1e10 + (0.0 + 3.14159)))) >= 3.14159)), F1: 67, F2: 1}
        _ = v8
    default:
        v9 := []int{1}
        _ = v9
    }
    _ = p0
    _ = v7
    return 607
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -820
    v3 := []uint32{670}
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    v6 := make(map[int]bool)
    var v10 bool = ((65535 == 65535) && ((false && (!((0.1 <= float32((1.5 - (-2.25 * float64(-2.25))))) || true))) || false))
    v10 = (0 == 32767)
    v2 -= (v2 / ((0) | 1))
    v10 = (float32(((3.14159 * (1e10 / 3.14159)) * (0.1 * 1e10))) != 1e10)
    fmt.Println(v10, v10)
    for fi11 := 0; fi11 < 12; fi11++ {
        switch ("\n\t" + "a") {
        case "hello":
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 int32 = -1
                var v14 uint = 809
                fmt.Println(v13, fi12)
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            if v10 {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi11, v10, fi11)
                }
            } else {
                fmt.Println(v2, fi11)
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi11
    }
    var v15 string = ("a" + ("\n\t" + ""))
    fmt.Println(v15, v2, v15)
    switch v15 {
    case "a":
        v16 := v5[42]
        _ = v16
    case "foo bar":
        v15 = v15
    default:
        _ = len(v6)
    }
    fmt.Println(v2, v15)
    v17 := GenS1{F0: v10, F1: 106, F2: 723}
    var v18 float32 = ((((1e10 / (0.0 - (-2.25 - (0.0 + 0.1)))) - (0.1 - (-2.25 - -2.25))) - -2.25) / 1.5)
    v2 += v2
    v19 := GenS1{F0: true, F1: 253, F2: 531}
    v3 = append(v3, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
