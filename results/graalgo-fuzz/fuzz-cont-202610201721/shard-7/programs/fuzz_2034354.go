package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 int16 }
func genFn0() int {
    var v5 int8 = 100
    v6 := GenS1{F0: -242, F1: -9223372036854775808, F2: v5}
    _ = v5
    _ = v6
    return 576
}

func main() {
    var v1 int16 = 1000
    v2 := []string{("hello" + ""), "hello", "", "hello"}
    v3 := make(map[uint32]string)
    var v4 float32 = (0.1 / ((float32(1.5) / 1e10) / 1e10))
    v7 := v3[1]
    v1 -= (v1 & 1)
    var v8 bool = true
    v4 -= v4
    v8 = false
    v9 := GenS2{F0: 1}
    v4 = ((float32(float64((v4 * v4))) + v4) + 0.1)
    var v10 int = 42
    var v11 int32 = 969
    v1 -= v1
    if (v8 && false) {
        for fi12 := 0; fi12 < 6; fi12++ {
            switch (v10 >> ((2) & 63)) {
            case 608:
                v7 = "\n\t"
            case 127:
                var v13 int = -105
                _ = v13
            case 459:
                if ((^(v11)) >= 42) {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v11 += 0
                        v7 = "foo bar"
                        if (v8 && v8) {
                            if (-1 >= (fi14 << ((0) & 63))) {
                                switch (v7 + "foo bar") {
                                case "hello":
                                    var v15 int64 = -1
                                    _ = v15
                                case "":
                                    fmt.Println(v10, v10, v10)
                                default:
                                    if v8 {
                                        var v16 string = ""
                                        _ = v16
                                    }
                                }
                                var v17 int64 = -509
                                _ = v17
                            }
                            fmt.Println(fi12, v7, fi14)
                            var v18 int16 = 0
                            _ = v18
                        } else {
                            var v19 uint16 = 0
                            v20 := GenS2{F0: v1}
                            if len(v2) > 0 {
                                v2[0] = ((v7 + "hello") + v7)
                            }
                            _ = v19
                            _ = v20
                        }
                        _ = fi14
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v3[56] = (("foo bar" + ((v7 + ("hello" + (v7 + ("\n\t" + (((((v7 + v7) + v7) + "\n\t") + v7) + ("" + v7)))))) + "a")) + v7)
                        _ = fi21
                    }
                    if false {
                        _ = len(v3)
                    }
                } else {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v4 += (((((float32(3.14159) / v4) + 0.0) + (((v4 * v4) * v4) / 1.0)) - 0.0) / float32(1e10))
                        v23 := GenS1{F0: -9223372036854775808, F1: (^(fi12)), F2: 42}
                        _ = fi22
                        _ = v23
                    }
                    v11 = (v11 >> ((255) & 63))
                }
            }
            var v24 int8 = 127
            _ = fi12
            _ = v24
        }
        fmt.Println(v4, v4, v8)
    }
    v7 = (v7 + "")
    v25 := GenS1{F0: 1, F1: (v10 ^ 4294967295), F2: 0}
    var v26 int8 = 54
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
