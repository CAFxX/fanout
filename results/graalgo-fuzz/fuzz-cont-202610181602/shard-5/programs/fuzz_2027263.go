package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0(p0 uint8) int64 {
    var v6 uint64 = 0
    var v7 uint32 = 566
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    v3 := []uint32{}
    var v4 uint8 = 127
    var v5 bool = true
    if true {
        v4 -= v4
        v5 = ("\n\t" >= v1)
    } else {
        fmt.Println(v5)
    }
    v8 := GenS1{F0: ("foo bar" + (v1 + v1)), F1: (250 >= 4294967295), F2: 2}
    var v9 int32 = -852
    _ = len(v2)
    fmt.Println(v4)
    if (0.1 <= (3.14159 * (3.14159 / (float64(1.5) + -2.25)))) {
        v9 += 543
    }
    var v10 bool = true
    if v10 {
        v11 := GenS1{F0: "\n\t", F1: ((v5 && true) || v10), F2: 0}
        v12 := GenS1{F0: "foo bar", F1: (42 >= 32767), F2: 1}
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int16 = -477
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi13
            _ = v14
        }
        _ = v11
        _ = v12
    }
    v5 = (42 < 2)
    v4 = 42
    _ = len(v3)
    v15 := GenS1{F0: ((v1 + "hello") + "\n\t"), F1: false, F2: 4294967295}
    var v16 int64 = 0
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 float32 = ((0.1 - 0.0) * float32(-2.25))
        _ = fi17
        _ = v18
    }
    var v19 uint32 = 4294967295
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := GenS2{F0: -9223372036854775808, F1: (-(v9))}
        fmt.Println(v16, v1)
        _ = fi20
        _ = v21
    }
    v19 += v19
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
