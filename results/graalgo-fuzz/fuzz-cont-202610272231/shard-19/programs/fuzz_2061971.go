package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int64; F1 int32; F2 uint16 }
func genFn0(p0 int32, p1 int, p2 int32) uint8 {
    fmt.Println(p1, p1)
    switch "" {
    case "foo bar":
        if (false && false) {
            v5 := GenS1{F0: (31 >= p1), F1: 0}
            _ = v5
        }
    case "a":
        if (!true) {
            p0 += (p0 ^ (p0 ^ (p0 ^ 42)))
        }
    default:
        if (100 == -128) {
            p2 += 957
        }
    }
    p2 += (p2 * p2)
    var v6 int64 = 716
    v7 := GenS2{F0: (4294967295 > 478)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint64 = 42
    var v2 int8 = -128
    var v3 int16 = 42
    v4 := make(map[int]int)
    var v8 uint8 = 127
    var v9 int16 = (v3 << ((0) & 63))
    var v10 int64 = 1
    var v11 uint32 = 2147483648
    v12 := GenS2{F0: false}
    v9 += 1000
    v11 = v11
    var v13 float64 = (-2.25 + 1e10)
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v11, fi14, v1)
        v11 -= (^(v11))
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    switch (("\n\t" + "foo bar") + ((("hello" + "\n\t") + "a") + "a")) {
    case "a":
        var v16 string = ("\n\t" + (("\n\t" + ((("hello" + (("" + "a") + "foo bar")) + "") + "\n\t")) + "foo bar"))
        _ = v16
    case "hello":
        if false {
            var v17 int64 = v10
            switch (("a" + "foo bar") + (("a" + "") + ("a" + ""))) {
            case "\n\t":
                _ = len(v4)
            default:
                if false {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        _ = len(v4)
                        v19 := GenS3{F0: (v1 + v1), F1: -300}
                        _ = fi18
                        _ = v19
                    }
                }
            }
            _ = v17
        } else {
            v20 := GenS4{F0: v10, F1: -666, F2: 506}
            var v21 float64 = ((v13 / 1e10) * (((1e10 + v13) / (v13 - ((v13 * (0.1 / (v13 * (v13 + v13)))) * v13))) + ((1.5 * v13) - 1.5)))
            _ = v20
            _ = v21
        }
    case "":
        v2 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
