package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
func genFn0(p0 int8, p1 int) uint {
    p1 = p1
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(p1, p1, p0)
        p0 = 42
        p1 -= fi4
        _ = fi4
    }
    p1 = (^(p1))
    var v5 int8 = p0
    _ = p0
    _ = p1
    _ = v5
    return 82
}

func main() {
    v1 := []int{-700, 0, 42}
    var v2 float64 = float64((-2.25 - 0.1))
    var v3 float64 = (float64(0.1) + -2.25)
    if false {
        var v6 float64 = (-2.25 * 3.14159)
        _ = v6
    }
    fmt.Println(v2, v3)
    var v7 int32 = -861
    for fi8 := 0; fi8 < 19; fi8++ {
        v1 = append(v1, (fi8 | 0))
        var v9 uint = 2
        var v10 uint8 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: "a", F1: 91}
    for fi12 := 0; fi12 < 13; fi12++ {
        if len(v1) > 0 {
            v1[0] = fi12
        }
        var v13 string = ("" + ("\n\t" + "hello"))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v2 = 1e10
        v2 += 1e10
        v3 += (1e10 * v3)
        _ = fi14
    }
    if true {
        var v15 uint32 = 65535
        _ = v15
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 float32 = ((((-2.25 * -2.25) * 0.1) + (1e10 * 0.1)) / 3.14159)
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                switch (fi19 ^ 930) {
                case 642:
                    v2 = float64((float32(1e10) / (((-2.25 - 1.5) + 1e10) - 3.14159)))
                default:
                    v20 := GenS1{F0: "a", F1: 33}
                    _ = v20
                }
                v21 := GenS1{F0: "a", F1: 255}
                _ = fi19
                _ = v21
            }
            _ = fi16
        }
    }
    if (!(127 == -1)) {
        v2 -= (v3 * (1.5 * -2.25))
        v3 -= (1e10 * float64(float32((v3 * v3))))
    }
    var v22 uint64 = 973
    switch 42 {
    case 4294967295:
        var v23 uint32 = 2
        _ = v23
    }
    var v24 uint16 = 65535
    switch ("a" + "\n\t") {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if (true || false) {
            switch ("" + "foo bar") {
            case "foo bar":
                var v25 float64 = 1e10
                _ = v25
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            fmt.Println(v2, v22)
            if len(v1) > 0 {
                v1[0] = 42
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v7
}
