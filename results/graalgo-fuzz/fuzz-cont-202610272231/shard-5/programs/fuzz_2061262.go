package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 uint16, p1 int) uint8 {
    v7 := []int{p1, p1}
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 4294967295
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    var v6 string = (("hello" + "a") + ("" + (("a" + "foo bar") + ("\n\t" + ("hello" + ("hello" + ""))))))
    var v8 int = (-(v2))
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            _ = len(v4)
            v11 := []int{v8, fi10, 434, -324}
            v2 += (fi9 * -65)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 bool = false
    v13 := GenS1{F0: -1, F1: -1, F2: -689}
    var v14 float32 = 1.5
    var v15 int8 = 100
    v16 := GenS1{F0: 1, F1: 1, F2: -2147483648}
    for fi17 := 0; fi17 < 3; fi17++ {
        v14 = v14
        _ = fi17
    }
    switch (^(v8)) {
    case 32768:
        for fi18 := 0; fi18 < 2; fi18++ {
            if (int64(v3) != 1) {
                v8 -= v8
                fmt.Println(v12, v12)
                if (!false) {
                    if v12 {
                        if (42 != 128) {
                            v14 += 3.14159
                        }
                        fmt.Println(v2, v6, v6)
                        v19 := GenS2{F0: ((v6 + v6) + (v6 + ("\n\t" + "foo bar"))), F1: -852}
                        _ = v19
                    } else {
                        fmt.Println(v12, v15, v3)
                    }
                }
            }
            v20 := []int{}
            _ = fi18
            _ = v20
        }
    case -9223372036854775808:
        v21 := []int{(v2 * (v2 & v8))}
        _ = v21
    case 128:
        v4[v8] = 80
    default:
        var v22 int8 = (v15 * v15)
        _ = v22
    }
    if (-2.25 != float64(1e10)) {
        v3 = 197
        var v23 bool = (1e10 != 1.5)
        for fi24 := 0; fi24 < 15; fi24++ {
            v25 := []int{128, v2, (-(v8)), v2}
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
