package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 string }
func genFn0(p0 int64, p1 bool) bool {
    if ((p0 & p0) < p0) {
        p0 += 42
        for fi4 := 0; fi4 < 12; fi4++ {
            p1 = (((1.5 * 1.5) - 0.0) >= float64(float32(1e10)))
            var v5 float64 = float64(1.5)
            if (1e10 < v5) {
                if false {
                    if (18446744073709551615 >= 0) {
                        fmt.Println(p0)
                        v6 := []int{32767}
                        var v7 uint8 = 42
                        _ = v6
                        _ = v7
                    } else {
                        var v8 uint32 = 2
                        var v9 uint16 = 1
                        _ = v8
                        _ = v9
                    }
                    var v10 string = "\n\t"
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v10 = (v10 + (v10 + "a"))
                        _ = fi11
                    }
                    _ = v10
                }
            }
            _ = fi4
            _ = v5
        }
        p0 += p0
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS1{F0: -1, F1: 1, F2: (p0 / ((p0) | 1))}
                var v15 uint = 1
                fmt.Println(v15, fi12)
                _ = fi13
                _ = v14
                _ = v15
            }
            if true {
                switch int(p0) {
                case 1:
                    if p1 {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            var v17 uint64 = 2
                            p1 = ((836 <= 788) && (!(927 != 9223372036854775807)))
                            _ = fi16
                            _ = v17
                        }
                        var v18 float32 = 0.1
                        var v19 int8 = 46
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint64 = 132
                        v21 := GenS1{F0: 197, F1: 0, F2: (p0 % ((p0) | 1))}
                        if (p1 && (!(((fi12 * -801) <= (^(fi12))) || true))) {
                            for fi22 := 0; fi22 < 15; fi22++ {
                                v23 := GenS2{F0: 1, F1: -31, F2: ("" + "hello")}
                                v20 += (v20 ^ 4294967295)
                                _ = fi22
                                _ = v23
                            }
                        }
                        _ = v20
                        _ = v21
                    }
                case 9223372036854775807:
                    p1 = p1
                case 804:
                    v24 := GenS2{F0: 42, F1: -128, F2: "a"}
                    _ = v24
                }
                var v25 float64 = -2.25
                _ = v25
            } else {
                fmt.Println(p1, p0)
            }
            _ = fi12
        }
        v26 := []int{}
        _ = len(v26)
        _ = v26
    }
    v27 := GenS1{F0: -128, F1: 0, F2: -718}
    _ = p0
    _ = p1
    _ = v27
    return (3.14159 > (0.0 / float64(-2.25)))
}

func main() {
    var v1 uint8 = 119
    v2 := []string{}
    var v3 uint8 = 127
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(v1)
            var v30 float32 = float32(((0.1 + (-2.25 - (3.14159 + (-2.25 * 3.14159)))) + 0.1))
            v1 = 62
            _ = fi29
            _ = v30
        }
        var v31 bool = true
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = fi28
        _ = v31
    }
    _ = len(v2)
    v1 -= (v1 - v3)
    if (!(-2147483648 > 279)) {
        v3 = (^(v3))
    }
    var v32 int64 = 9223372036854775807
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v33 uint64 = 485
    switch 128 {
    case 128:
        v32 += v32
    case 127:
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 int32 = 878
            var v36 uint8 = v1
            if (113 == uint(v36)) {
                var v37 float64 = -2.25
                _ = v37
            }
            _ = fi34
            _ = v35
            _ = v36
        }
    case -877:
        fmt.Println(v1, v3, v1)
    default:
        v38 := GenS2{F0: (v3 >> ((1) & 63)), F1: 127, F2: "a"}
        _ = v38
    }
    var v39 float32 = 1e10
    var v40 uint64 = 18446744073709551615
    var v41 float64 = float64(0.1)
    for fi42 := 0; fi42 < 10; fi42++ {
        v39 = v39
        v39 += v39
        _ = fi42
    }
    v1 -= (v3 << ((42) & 63))
    v2 = append(v2, "\n\t")
    if false {
        fmt.Println(v41, v32)
        var v43 int = 2147483647
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v40
    _ = v41
}
