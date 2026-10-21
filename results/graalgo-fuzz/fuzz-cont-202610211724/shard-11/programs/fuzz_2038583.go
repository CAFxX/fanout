package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 int64) int {
    v6 := []int{127, 34}
    p0 = p0
    _ = p0
    _ = v6
    return 980
}

func main() {
    var v1 int16 = -1
    var v2 int8 = 127
    v3 := make(map[uint32]bool)
    var v4 uint64 = 953
    var v5 uint = 1
    v7 := []int{-307, 9223372036854775807, -498}
    v1 = (-(v1))
    var v8 int32 = 0
    if (((v5 << (((v5 | uint(v4))) & 63)) == (v5 + v5)) && (v8 <= v8)) {
        switch 32767 {
        case 32768:
            if true {
                var v9 int16 = v1
                var v10 float32 = 0.0
                var v11 float64 = 0.0
                _ = v10
                _ = v11
                _ = v9
            }
        case -1:
            for fi12 := 0; fi12 < 8; fi12++ {
                v8 = 2147483647
                v4 = (^(v4))
                fmt.Println(v8, v1, v8)
                _ = fi12
            }
        case 128:
            _ = len(v7)
        default:
            if true {
                if ((v4 * v4) != v4) {
                    var v13 int64 = 0
                    for fi14 := 0; fi14 < 6; fi14++ {
                        if (!("\n\t" == ("hello" + ("a" + "foo bar")))) {
                            var v15 string = ("" + "")
                            fmt.Println(v13, v2, v5)
                            _ = len(v7)
                            _ = v15
                        } else {
                            v2 = 1
                        }
                        switch "" {
                        case "hello":
                            switch ("hello" + (("\n\t" + ("\n\t" + "hello")) + (("hello" + "\n\t") + (("hello" + "foo bar") + "")))) {
                            case "\n\t":
                                var v16 float64 = ((1.5 + float64(0.1)) - 1e10)
                                _ = v16
                            default:
                                fmt.Println(v4)
                            }
                        default:
                            var v17 uint16 = 1
                            _ = v17
                        }
                        _ = fi14
                    }
                    delete(v3, 966)
                    _ = v13
                } else {
                    v5 = (v5 >> ((v5) & 63))
                    var v18 int64 = 525
                    _ = v18
                }
                var v19 int32 = 1
                _ = len(v3)
                _ = v19
            }
        }
        v20 := GenS1{F0: (-(v2))}
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 6; fi21++ {
            fmt.Println(v5)
            _ = fi21
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            v8 -= v8
            _ = fi22
        }
    } else {
        var v23 uint = 1
        _ = v23
    }
    v5 = 42
    v24 := GenS2{F0: -488, F1: 1}
    delete(v3, 1)
    v2 = (^(v2))
    v4 -= (v4 / ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
