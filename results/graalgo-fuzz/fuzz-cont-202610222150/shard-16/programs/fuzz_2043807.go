package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    var v3 uint = 4294967295
    var v4 float32 = float32(3.14159)
    v5 := make(map[string]uint32)
    delete(v2, "\n\t")
    switch 438 {
    case -2147483648:
        v6 := GenS1{F0: 0, F1: v3, F2: v1}
        _ = v6
    default:
        if (-1 != 104) {
            var v7 uint32 = 432
            _ = v7
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    case "a":
        for fi10 := 0; fi10 < 12; fi10++ {
            if (0 >= 852) {
                if false {
                    v5[("\n\t" + (("foo bar" + ("hello" + ("a" + "\n\t"))) + ("foo bar" + "foo bar")))] = 870
                    switch ("" + "") {
                    case "":
                        _ = len(v5)
                    case "\n\t":
                        if ((false || (false || false)) || false) {
                            v4 = 0.1
                        }
                    }
                } else {
                    if (244 < -216) {
                        for fi11 := 0; fi11 < 12; fi11++ {
                            v12 := []int{fi10, (fi11 ^ (fi10 ^ fi10)), (fi11 / (((fi10 / ((fi10) | 1))) | 1)), fi11}
                            _ = fi11
                            _ = v12
                        }
                    }
                    v13 := GenS1{F0: 4294967295, F1: 4294967295, F2: (^(v1))}
                    v1 += (v1 | (v1 >> (((v3 >> ((v3) & 63))) & 63)))
                    _ = v13
                }
                fmt.Println(fi10, fi10)
            }
            _ = fi10
        }
    case "":
        for fi14 := 0; fi14 < 7; fi14++ {
            if (!false) {
                _ = len(v2)
            }
            if true {
                if (!((v4 + float32(float64((v4 / 3.14159)))) == 1.5)) {
                    v15 := []int{(fi14 * (fi14 ^ fi14)), fi14, fi14}
                    v16 := GenS1{F0: 2147483647, F1: (^(v3)), F2: v1}
                    _ = v15
                    _ = v16
                }
            } else {
                v17 := v5[("a" + ("foo bar" + "a"))]
                var v18 uint64 = uint64(v17)
                fmt.Println(v4, v18, v3)
                _ = v17
                _ = v18
            }
            var v19 uint64 = 1
            _ = fi14
            _ = v19
        }
    }
    v3 -= 1
    for fi20 := 0; fi20 < 17; fi20++ {
        if ((v4 * 0.1) > v4) {
            var v21 int8 = 1
            _ = v21
        }
        switch fi20 {
        case -9223372036854775808:
            var v22 bool = false
            _ = v22
        default:
            var v23 int32 = 512
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int16 = 433
        fmt.Println(v25, v4, v3)
        v26 := GenS1{F0: 65535, F1: 1, F2: v1}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint64 = 4294967295
    if (!(!(377 != 2147483648))) {
        var v28 uint = v3
        _ = v28
    }
    var v29 uint64 = 4294967295
    var v30 uint64 = (v27 / ((766) | 1))
    v1 += v1
    fmt.Println(v3)
    v3 += (v3 | 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
