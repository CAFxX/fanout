package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 string) int64 {
    var v6 bool = (127 >= 2147483647)
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{}
        switch fi7 {
        case 912:
            switch "" {
            case "foo bar":
                var v9 float64 = (1.5 / (float64(-2.25) + float64(-2.25)))
                _ = v9
            }
        case 1:
            var v10 float64 = 3.14159
            _ = v10
        case -563:
            fmt.Println(p0, p0)
        }
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    var v11 uint64 = 2
    _ = p0
    _ = v11
    _ = v6
    return -690
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 int32 = 63
    var v5 int16 = -1
    for fi12 := 0; fi12 < 19; fi12++ {
        delete(v1, 488)
        v4 = v4
        fmt.Println(v5)
        _ = fi12
    }
    if (0 < 0) {
        v13 := v1[625]
        _ = v13
    }
    fmt.Println(v4, v4)
    v5 = 210
    v3[128] = 1
    if (!true) {
        v5 += v5
    }
    var v14 uint32 = 2147483648
    v15 := []int{-195}
    v5 = v5
    var v16 uint16 = 0
    switch 128 {
    case 32768:
        if ((!true) || (460 != 1)) {
            var v17 uint = 2
            var v18 uint16 = 32767
            v19 := GenS1{F0: -509, F1: -581, F2: 924}
            _ = v17
            _ = v18
            _ = v19
        }
    case -2147483648:
        v5 += v5
    default:
        v20 := v2[0]
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
