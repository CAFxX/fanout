package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = -1
    var v2 float32 = (float32(0.1) - ((((float32((1.5 - float64(0.0))) * (-2.25 / 1.5)) - 0.0) - 3.14159) + 1e10))
    v3 := make(map[int]string)
    var v4 int = v1
    var v5 uint = 650
    fmt.Println(v2, v5, v2)
    v1 = 127
    var v6 uint = v5
    v4 = v1
    for fi7 := 0; fi7 < 6; fi7++ {
        if (!true) {
            if (true || false) {
                v5 += v5
            }
            var v8 uint64 = 18446744073709551615
            var v9 uint32 = 2147483648
            _ = v8
            _ = v9
        } else {
            v1 -= fi7
            var v10 uint32 = 159
            _ = v10
        }
        v11 := GenS1{F0: 2147483647}
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            switch 765 {
            case -9223372036854775808:
                var v14 int8 = 42
                _ = v14
            case -351:
                var v15 int16 = -1
                _ = v15
            }
            var v16 int32 = 61
            _ = fi13
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            v5 = v5
            _ = fi17
        }
        switch (-(v1)) {
        case 48:
            v18 := []int{fi12, v4, v1, fi12}
            _ = v18
        case 128:
            var v19 int = v1
            _ = v19
        }
        _ = fi12
    }
    var v20 float64 = (float64((0.0 * (v2 / v2))) / 1.0)
    v21 := []int{v1, v4, v4, (v1 << (((^(v6))) & 63))}
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(v5, v5, v6)
        _ = fi22
    }
    var v23 uint8 = 255
    var v24 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
