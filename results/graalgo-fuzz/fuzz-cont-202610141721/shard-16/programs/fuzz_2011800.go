package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8) uint {
    v4 := GenS1{F0: 2, F1: 42, F2: 973}
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 float32 = 3.14159
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: (p0 + (p0 - 2))}
    v8 := []int{2147483647, 127}
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int = -1
    var v2 uint32 = 1
    var v3 float32 = (1.5 / 1.5)
    var v9 float32 = -2.25
    var v10 int = v1
    fmt.Println(v3)
    v9 = (v3 * (v3 / v3))
    v11 := []int{}
    for fi12 := 0; fi12 < 14; fi12++ {
        v3 = 1.5
        if false {
            if ((!(false || (-923 >= 1))) || true) {
                v1 = fi12
            }
            var v13 uint64 = 1
            var v14 bool = false
            _ = v13
            _ = v14
        }
        if (!false) {
            v15 := GenS1{F0: -294, F1: 0, F2: -767}
            _ = len(v11)
            var v16 uint16 = 0
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch "" {
        case "foo bar":
            v2 += v2
        case "\n\t":
            v18 := GenS1{F0: v1, F1: -601, F2: 32767}
            _ = v18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 float32 = (-2.25 / (v3 / 0.1))
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi19
        _ = v20
    }
    var v21 int = v1
    v22 := GenS1{F0: v1, F1: 42, F2: 1000}
    var v23 int = v10
    var v24 float32 = v9
    fmt.Println(v10, v9, v23)
    v24 -= 1.5
    v2 -= 191
    v9 += 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v9
}
