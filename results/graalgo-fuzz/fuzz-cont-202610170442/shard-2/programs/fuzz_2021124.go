package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
func genFn0(p0 uint64, p1 uint8, p2 uint32) string {
    fmt.Println(p1, p2, p2)
    if false {
        switch 0 {
        case -84:
            p0 = 0
        case -65:
            v4 := []int{0, 1, -565}
            _ = v4
        case -66:
            if true {
                fmt.Println(p1, p1, p0)
                v5 := []int{4294967295, 390, -966, -2147483648}
                _ = v5
            }
        }
        if (p1 >= 209) {
            fmt.Println(p0, p0, p0)
            p1 -= 127
        }
        v6 := GenS1{F0: 1}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(p2)
            _ = fi7
        }
        var v8 uint64 = p0
        _ = v8
    }
    p0 = 42
    p0 = (^(p0))
    v9 := GenS1{F0: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (("\n\t" + ("hello" + ("foo bar" + (("\n\t" + "\n\t") + "")))) + "\n\t")
}
func genFn1(p0 bool, p1 int32) uint16 {
    var v10 int = 1
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v10
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    v3 := make(map[string]int)
    switch -9223372036854775808 {
    case -610:
        if (false && true) {
            switch 997 {
            case 2147483647:
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 float32 = v2
                    _ = fi11
                    _ = v12
                }
            case -1:
                v2 -= 0.0
            default:
                var v13 int32 = 0
                _ = v13
            }
            var v14 int16 = 0
            fmt.Println(v2, v2)
            _ = v14
        }
    case 95:
        var v15 float64 = 0.0
        _ = v15
    }
    var v16 uint32 = 666
    var v17 uint64 = 0
    if (true || true) {
        v2 = 3.14159
    } else {
        var v18 float64 = (((1.5 + ((0.0 - (0.1 / 0.1)) / (0.0 - 0.1))) - -2.25) - float64((0.0 * (0.1 / 1.0))))
        _ = v18
    }
    fmt.Println(v2)
    v16 += (v16 + v16)
    v19 := GenS2{F0: 255, F1: (false && false), F2: 0}
    v20 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
