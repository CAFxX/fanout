package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0() int16 {
    v4 := []int{}
    var v5 uint32 = 2147483647
    if false {
        if true {
            v5 = v5
            var v6 bool = (1 >= uint64(v5))
            v6 = v6
            _ = v6
        } else {
            for fi7 := 0; fi7 < 12; fi7++ {
                v5 += 65535
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 string = "a"
                    fmt.Println(v9)
                    _ = fi8
                    _ = v9
                }
                v5 = 2
                _ = fi7
            }
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 1
    var v3 int32 = 0
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int64 = 42
            _ = fi11
            _ = v12
        }
        if ((98 != 14) || (!(387 < 18446744073709551615))) {
            v3 -= v3
            _ = len(v1)
        }
        v2 = 1
        _ = fi10
    }
    var v13 float64 = (float64(-2.25) * 3.14159)
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint = 4294967295
        var v16 int8 = -128
        v2 = 42
        _ = fi14
        _ = v15
        _ = v16
    }
    delete(v1, "foo bar")
    var v17 float32 = -2.25
    var v18 uint32 = 1
    _ = len(v1)
    v19 := []int{-571}
    v13 += (1e10 / -2.25)
    var v20 bool = true
    v17 = v17
    var v21 float64 = (v13 + float64(((v17 / v17) / (v17 / 3.14159))))
    for fi22 := 0; fi22 < 13; fi22++ {
        switch "a" {
        case "":
            if true {
                if (128 <= 63) {
                    v3 += -2147483648
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if len(v19) > 0 {
                            v19[0] = (fi23 - (fi22 | fi22))
                        }
                        switch fi22 {
                        case 1:
                            v24 := GenS1{F0: 255, F1: 87}
                            _ = v24
                        case 32768:
                            var v25 uint32 = 0
                            _ = v25
                        default:
                            var v26 int8 = -128
                            _ = v26
                        }
                        if (true && v20) {
                            v19 = append(v19, (fi23 | fi23))
                            var v27 uint64 = 885
                            _ = v27
                        }
                        _ = fi23
                    }
                }
            }
        case "foo bar":
            var v28 string = (("a" + "") + (("\n\t" + "hello") + "\n\t"))
            _ = v28
        default:
            for fi29 := 0; fi29 < 8; fi29++ {
                switch (-(fi29)) {
                case 42:
                    v3 -= int32(fi22)
                default:
                    if (-34 >= 131) {
                        v13 -= (3.14159 - (v13 - float64(v17)))
                        fmt.Println(v3, v20)
                    } else {
                        v2 = 25
                    }
                }
                v2 += (v2 << ((1) & 63))
                _ = fi29
            }
        }
        _ = fi22
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        fmt.Println(fi30, v3)
        _ = fi30
    }
    fmt.Println(v18, v3)
    v31 := GenS1{F0: 49, F1: 0}
    var v32 string = ((("\n\t" + (("a" + "foo bar") + ("" + "hello"))) + "foo bar") + "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
}
