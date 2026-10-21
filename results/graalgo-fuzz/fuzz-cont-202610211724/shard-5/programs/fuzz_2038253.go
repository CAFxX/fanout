package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = -2.25
    v3 := []int{42, 128, 42, 127}
    var v4 uint32 = 842
    v5 := []string{}
    var v6 uint64 = 942
    fmt.Println(v2, v6, v2)
    v3 = append(v3, 1)
    for fi7 := 0; fi7 < 19; fi7++ {
        v4 -= (v4 / ((v4) | 1))
        var v8 int = -1
        if false {
            var v9 uint8 = 128
            var v10 float32 = 3.14159
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v11 int16 = 150
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: (^(v6))}
        _ = fi12
        _ = v13
    }
    fmt.Println(v2)
    v4 = 106
    fmt.Println(v4)
    v11 += v11
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v2, fi14, v4)
        v6 = v6
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi14
    }
    var v15 bool = (true || true)
    v11 -= -32768
    v4 -= v4
    v11 = (v11 | v11)
    if ((v15 || (((127 == 1) && (!(-9223372036854775808 < -1))) || (0.0 != (-2.25 / 0.1)))) || (!v15)) {
        v16 := GenS1{F0: v6}
        _ = v16
    }
    switch int(v4) {
    case -15:
        v17 := GenS1{F0: (v6 >> ((42) & 63))}
        _ = v17
    default:
        v18 := v1[(v4 + (v4 & 0))]
        _ = v18
    }
    fmt.Println(v2, v11, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
