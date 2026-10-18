package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64, p1 int, p2 int8) uint64 {
    if (!true) {
        var v6 uint8 = 127
        _ = v6
    } else {
        if ((^(p0)) >= (^(p0))) {
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := []int{p1, (p1 / ((fi7) | 1))}
                if (p0 >= p0) {
                    p1 -= 4294967295
                }
                _ = fi7
                _ = v8
            }
        }
        p1 += (^(p1))
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{128, (fi9 | p1), 0, (fi10 / (((p1 + (-(fi9)))) | 1))}
            p0 += p0
            switch (^(fi9)) {
            case 42:
                p0 = p0
                var v12 float64 = float64(0.0)
                _ = v12
            default:
                p1 = (p1 % ((fi10) | 1))
            }
            p1 -= (p1 / ((p1) | 1))
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 * (p0 * p0))
}

func main() {
    v1 := []string{("a" + "a"), "foo bar"}
    var v2 float64 = float64(0.0)
    var v3 bool = ((!true) && true)
    v4 := make(map[uint32]int)
    var v5 string = ""
    var v13 uint32 = 189
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            v13 += v13
            var v16 uint16 = 642
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 uint32 = (v13 ^ 799)
    if false {
        v4[(v13 * v13)] = 377
        v2 -= v2
        switch v5 {
        case "hello":
            if true {
                var v18 uint64 = 619
                if (32767 > 255) {
                    _ = len(v1)
                }
                v18 += (v18 & 4294967295)
                _ = v18
            }
        }
    }
    var v19 uint16 = 987
    v20 := GenS1{F0: int32(v19)}
    v2 -= v2
    if v3 {
        v3 = (-1 >= -65)
        v13 -= 523
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
