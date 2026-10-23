package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0(p0 uint64, p1 int) int64 {
    for fi7 := 0; fi7 < 17; fi7++ {
        switch -536 {
        case 42:
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 float64 = (0.1 * float64((0.0 / 1e10)))
                var v10 uint = 18446744073709551615
                _ = fi8
                _ = v10
                _ = v9
            }
        case 775:
            var v11 uint32 = 987
            _ = v11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint8 = 249
            _ = fi12
            _ = v13
        }
        _ = fi7
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 bool = ((0.0 * 0.1) > 1e10)
        switch ("foo bar" + (("\n\t" + "a") + "foo bar")) {
        case "a":
            switch 0 {
            case 609:
                v16 := GenS1{F0: (p0 + p0), F1: 42}
                _ = v16
            }
        case "foo bar":
            fmt.Println(v15, p1, fi14)
        case "\n\t":
            var v17 bool = v15
            _ = v17
        default:
            v15 = (108 != 637)
        }
        v18 := []int{fi14, fi14}
        _ = fi14
        _ = v15
        _ = v18
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []string{"foo bar", ""}
    v2 := make(map[uint32]bool)
    var v3 string = ""
    var v4 uint = 4294967295
    var v5 string = v3
    v6 := []string{v3, v5}
    var v19 float32 = -2.25
    switch 824 {
    case 0:
        var v20 uint64 = 42
        _ = v20
    default:
        v21 := GenS1{F0: 18446744073709551615, F1: genFn0(18446744073709551615, 2147483647)}
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        if ((-2.25 / float64(v19)) > 3.14159) {
            var v23 int32 = 0
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 uint = (v4 % ((v4) | 1))
                fmt.Println(fi24, v25, v23)
                fmt.Println(v23, v23, v19)
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        var v26 int8 = 71
        _ = fi22
        _ = v26
    }
    fmt.Println(v3)
    v3 = "foo bar"
    v27 := GenS2{F0: 0, F1: 2, F2: (^(v4))}
    var v28 uint = 0
    fmt.Println(v28)
    for fi29 := 0; fi29 < 19; fi29++ {
        v28 -= (^(v28))
        _ = fi29
    }
    var v30 uint64 = 25
    v30 += v30
    var v31 uint64 = (^(v30))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
