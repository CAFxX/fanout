package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint64 = 471
    var v2 uint8 = 255
    var v3 bool = true
    var v4 int8 = -1
    var v5 uint64 = v1
    var v6 int16 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint16 = 32767
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 402
    var v10 uint32 = (v9 * (v9 << ((2) & 63)))
    var v11 int64 = 42
    v12 := GenS1{F0: 354}
    v2 -= (^(v2))
    fmt.Println(v6, v9, v6)
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int32 = -884
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: (fi13 | fi13)}
        v9 = 42
        _ = fi13
        _ = v16
    }
    v5 = v5
    v4 += int8(v6)
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 float64 = 3.14159
        _ = fi17
        _ = v18
    }
    if (v3 && (v3 || (!v3))) {
        v19 := []int{int(v1), 1}
        _ = v19
    }
    switch 4294967295 {
    case -2147483648:
        v9 = v9
    }
    var v20 float32 = ((-2.25 * 0.0) * (-2.25 - float32((0.1 * 0.1))))
    var v21 uint64 = v1
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint16 = 331
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
