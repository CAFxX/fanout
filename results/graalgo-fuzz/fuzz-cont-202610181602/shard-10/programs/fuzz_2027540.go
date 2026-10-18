package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 string) uint {
    v5 := GenS1{F0: 32767, F1: 4294967295, F2: 127}
    if true {
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint8 = 83
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    }
    v9 := []int{128, -2147483648, -2147483648}
    _ = p0
    _ = v5
    _ = v9
    return 536
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -194
    var v3 uint32 = 2147483647
    var v4 uint8 = 42
    v4 = 86
    switch 9223372036854775807 {
    case -356:
        for fi10 := 0; fi10 < 5; fi10++ {
            if ((674 >= -600) && true) {
                v11 := GenS1{F0: 1, F1: v3, F2: uint8(v3)}
                var v12 uint = 498
                for fi13 := 0; fi13 < 1; fi13++ {
                    v3 -= 514
                    v12 = (v12 << ((v12) & 63))
                    _ = fi13
                }
                _ = v11
                _ = v12
            }
            _ = fi10
        }
    }
    fmt.Println(v3, v3)
    var v14 int16 = 1
    if true {
        v15 := GenS2{F0: 127}
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v2)
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint8 = 96
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v2 += v2
    v19 := GenS1{F0: 0, F1: v3, F2: (v4 | (v4 | (v4 << ((genFn0(("foo bar" + ("a" + "hello")))) & 63))))}
    var v20 uint16 = 613
    v20 = 931
    var v21 float64 = 0.0
    var v22 uint = 18446744073709551615
    var v23 uint16 = v20
    fmt.Println(v4, v20)
    v14 = (v14 >> ((1) & 63))
    fmt.Println(v14, v14)
    v24 := []int{int(v2), 127}
    fmt.Println(v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
