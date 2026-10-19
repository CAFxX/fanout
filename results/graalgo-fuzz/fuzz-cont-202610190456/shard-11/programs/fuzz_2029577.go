package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0() int32 {
    var v7 int = -657
    if false {
        v8 := GenS1{F0: 436}
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(fi9, v7)
            v7 -= (v7 << ((1) & 63))
            v10 := []int{(fi9 ^ 42), (fi9 | (^(fi9))), -9223372036854775808, (v7 / ((v7) | 1))}
            _ = fi9
            _ = v10
        }
        fmt.Println(v7)
        _ = v8
    }
    v11 := GenS1{F0: 4294967295}
    for fi12 := 0; fi12 < 8; fi12++ {
        if (65535 != 698) {
            v13 := []int{(fi12 / ((fi12) | 1)), (v7 - (fi12 % (((^(v7))) | 1))), 32767}
            _ = v13
        } else {
            fmt.Println(fi12)
            if true {
                var v14 int16 = 0
                fmt.Println(v14, v7)
                _ = v14
            } else {
                fmt.Println(v7, fi12)
                var v15 uint16 = 255
                v7 = (v7 ^ int(v15))
                _ = v15
            }
        }
        var v16 float32 = 0.1
        _ = fi12
        _ = v16
    }
    _ = v11
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (0.1 / (1e10 + 3.14159))
    v3 := make(map[int]int)
    v4 := []string{"hello", "", "\n\t", ("\n\t" + "")}
    var v5 float64 = v2
    var v6 uint8 = 1
    if ("hello" >= "") {
        fmt.Println(v5)
    }
    if (!true) {
        if true {
            fmt.Println(v2, v5, v2)
        }
        fmt.Println(v5, v6, v5)
    } else {
        v6 = v6
        v6 = v6
        v6 = v6
    }
    var v17 uint8 = 62
    v6 = v6
    v17 -= v17
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS1{F0: 71}
        v20 := GenS1{F0: 669}
        _ = fi18
        _ = v19
        _ = v20
    }
    if false {
        v21 := GenS1{F0: 1}
        _ = v21
    }
    var v22 int64 = 9223372036854775807
    var v23 int = -2147483648
    if (v23 < v23) {
        v4 = append(v4, "a")
    } else {
        v17 -= 87
        for fi24 := 0; fi24 < 1; fi24++ {
            if true {
                v25 := GenS2{F0: 42, F1: 127}
                v5 += ((1e10 / float64(1e10)) - v5)
                _ = v25
            } else {
                if false {
                    var v26 bool = ((((v2 * v5) / 1e10) - v2) != 3.14159)
                    if (!(!false)) {
                        v27 := GenS2{F0: 32767, F1: 1}
                        if ((!(!(-2.25 <= float32(v2)))) || true) {
                            v2 = (v5 + (1.5 + v5))
                            delete(v1, (("a" + "hello") + "foo bar"))
                            if ((3.14159 * v5) >= v5) {
                                var v28 uint64 = 42
                                fmt.Println(v6)
                                _ = v28
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v27
                    }
                    for fi29 := 0; fi29 < 8; fi29++ {
                        v6 += (^(v17))
                        _ = fi29
                    }
                    _ = v26
                } else {
                    _ = len(v3)
                }
                v30 := GenS2{F0: 804, F1: (v6 >> ((255) & 63))}
                v31 := v1["foo bar"]
                _ = v30
                _ = v31
            }
            _ = fi24
        }
    }
    fmt.Println(v22, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
