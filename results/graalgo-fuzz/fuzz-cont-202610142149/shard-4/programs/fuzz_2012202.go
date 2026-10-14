package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16, p1 int64, p2 uint8) uint64 {
    v6 := []int{}
    var v7 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 22
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]int)
    v3 := []string{}
    var v4 uint = 300
    var v5 int32 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(v1, v5, fi8)
        fmt.Println(v1, v4, v5)
        _ = fi8
    }
    v5 += 42
    var v9 uint16 = 42
    switch "\n\t" {
    case "hello":
        if true {
            if false {
                v2[-2147483648] = 2147483647
            } else {
                if (-32768 >= 371) {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        var v11 uint32 = 2147483647
                        _ = fi10
                        _ = v11
                    }
                }
                var v12 int16 = 975
                _ = v12
            }
            var v13 string = ""
            fmt.Println(v5, v1)
            _ = v13
        } else {
            v1 += v1
            for fi14 := 0; fi14 < 19; fi14++ {
                if (!(false || ((v1 * v1) != (v1 % ((-9223372036854775808) | 1))))) {
                    if (42 <= (v4 * v4)) {
                        fmt.Println(v1)
                        v9 = v9
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if (!(genFn0((v9 ^ v9), v1, 127) <= 369)) {
                            v15 := GenS1{F0: 1000, F1: (^(v1)), F2: -128}
                            _ = v15
                        }
                    default:
                        if true {
                            var v16 int8 = 1
                            v1 = int64(v5)
                            v17 := GenS2{F0: v4}
                            _ = v16
                            _ = v17
                        } else {
                            v18 := GenS1{F0: -502, F1: (^(v1)), F2: 93}
                            for fi19 := 0; fi19 < 18; fi19++ {
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    switch 42 {
                                    case 32768:
                                        if true {
                                            switch (^(fi14)) {
                                            case 42:
                                                v4 = (v4 | (v4 & (v4 << (((v4 % ((v4) | 1))) & 63))))
                                            }
                                        }
                                    case -2147483648:
                                        v9 -= v9
                                    default:
                                        v20 := GenS2{F0: 42}
                                        _ = v20
                                    }
                                case "":
                                    fmt.Println(fi14)
                                case "hello":
                                    var v21 uint = uint(v9)
                                    _ = v21
                                }
                                _ = fi19
                            }
                            var v22 uint16 = v9
                            _ = v18
                            _ = v22
                        }
                    }
                }
                var v23 uint = v4
                _ = fi14
                _ = v23
            }
            var v24 int32 = -296
            _ = v24
        }
    }
    var v25 uint = 4294967295
    delete(v2, -9223372036854775808)
    v26 := GenS1{F0: -32768, F1: v1, F2: -62}
    fmt.Println(v5)
    var v27 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
