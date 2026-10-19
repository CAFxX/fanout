package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0(p0 uint8, p1 int8, p2 uint) int32 {
    p2 = 4294967295
    v4 := GenS1{F0: p2, F1: -251}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}
func genFn1() int8 {
    v5 := []int{-511, -1}
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = 178
        if (-432 >= 1000) {
            v7 = v7
            if false {
                var v8 int8 = 42
                v8 -= v8
                var v9 bool = false
                _ = v8
                _ = v9
            }
            _ = len(v5)
        }
        v10 := GenS1{F0: 1, F1: -2147483648}
        _ = fi6
        _ = v10
        _ = v7
    }
    var v11 uint32 = 2
    for fi12 := 0; fi12 < 10; fi12++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v11 -= v11
        _ = fi12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 float32 = 3.14159
        _ = fi13
        _ = v14
    }
    _ = v11
    _ = v5
    return -128
}

func main() {
    var v1 bool = true
    var v2 uint32 = 211
    v3 := make(map[uint32]bool)
    for fi15 := 0; fi15 < 13; fi15++ {
        v1 = ((1e10 * 1e10) <= 1.5)
        var v16 int16 = 0
        _ = fi15
        _ = v16
    }
    var v17 uint32 = v2
    fmt.Println(v1, v17, v1)
    fmt.Println(v1, v2, v2)
    if (!true) {
        v2 = v17
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 int64 = -905
            if ("foo bar" <= ("\n\t" + ("" + "hello"))) {
                v1 = ((102 != 4294967295) && false)
                var v20 uint16 = 662
                _ = v20
            } else {
                v2 += (v2 % ((1) | 1))
                var v21 int = fi18
                _ = v21
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                if (42 > 1) {
                    var v23 uint16 = 159
                    switch (fi18 >> ((65535) & 63)) {
                    case 762:
                        var v24 bool = (!v1)
                        _ = v24
                    case 561:
                        v25 := GenS2{F0: 1, F1: (1e10 != (((float32(float64(0.0)) + (0.1 / -2.25)) + 3.14159) + 3.14159))}
                        _ = v25
                    case -602:
                        v26 := GenS2{F0: 257, F1: v1}
                        _ = v26
                    }
                    _ = v23
                }
                if (1 >= 1) {
                    v27 := []int{(fi22 | (fi18 ^ (fi22 | (fi22 | fi18)))), (fi22 - (fi18 * (fi18 << ((628) & 63)))), (fi22 + 2147483647)}
                    v28 := v3[(^(v2))]
                    switch int(v2) {
                    case 42:
                        switch (("" + ("a" + "")) + ("hello" + "hello")) {
                        case "":
                            var v29 bool = ((("\n\t" + "\n\t") + "hello") > "foo bar")
                            _ = v29
                        default:
                            var v30 uint = 0
                            _ = v30
                        }
                    default:
                        var v31 int = (fi22 + -361)
                        _ = v31
                    }
                    v17 -= (v2 + 65535)
                    _ = v27
                    _ = v28
                } else {
                    v17 -= (^(v2))
                }
                _ = fi22
            }
            _ = fi18
            _ = v19
        }
    } else {
        v1 = v1
        v17 = (v2 / ((v17) | 1))
        v3[(^(v17))] = (1 > 0)
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 int32 = 249
        fmt.Println(v1)
        var v34 bool = v1
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS2{F0: 308, F1: v1}
    var v36 uint32 = (v2 ^ v17)
    var v37 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
