package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 uint32, p1 bool, p2 int) int {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{32768, p2, fi7, (^(p2))}
        v9 := GenS1{F0: -22}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = ("a" != "\n\t")
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 float64 = 1.5
        for fi13 := 0; fi13 < 18; fi13++ {
            fmt.Println(v10, fi13, v12)
            var v14 uint8 = 42
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 uint32 = 42
    var v16 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v15
    _ = v16
    return p2
}

func main() {
    var v1 float64 = (0.1 / (1.5 - ((0.0 - float64(3.14159)) / 1.0)))
    v2 := []string{(("" + "foo bar") + "")}
    var v3 uint16 = 484
    v4 := make(map[int]string)
    var v5 string = "hello"
    var v6 uint32 = 818
    var v17 int64 = 42
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint16 = v3
        _ = fi18
        _ = v19
    }
    var v20 uint16 = v3
    var v21 uint16 = v20
    var v22 uint8 = 103
    v6 = v6
    var v23 int = -251
    if false {
        for fi24 := 0; fi24 < 15; fi24++ {
            if (false || (2 <= (v6 % ((v6) | 1)))) {
                var v25 uint64 = 18446744073709551615
                v5 = "a"
                _ = v25
            }
            if (-1 <= (v17 ^ 775)) {
                var v26 int64 = (-(v17))
                _ = v26
            } else {
                delete(v4, (v23 / ((42) | 1)))
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 string = (("hello" + "hello") + v5)
                    var v29 int8 = int8(v22)
                    v22 -= 42
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                switch (v5 + (v5 + ((v5 + v5) + (v5 + "a")))) {
                case "\n\t":
                    _ = len(v4)
                case "a":
                    var v31 uint8 = (v22 ^ (v22 << ((967) & 63)))
                    _ = v31
                default:
                    v2 = append(v2, v5)
                }
                var v32 int = (^(v23))
                if (217 > 42) {
                    var v33 bool = true
                    var v34 int64 = v17
                    if (583 < 4294967295) {
                        switch "hello" {
                        case "hello":
                            v20 -= v3
                        default:
                            fmt.Println(v5, v3)
                        }
                        v35 := GenS2{F0: -32768, F1: -69}
                        if v33 {
                            v22 = (v22 / (((^(v22))) | 1))
                        }
                        var v36 bool = (!(992 != 260))
                        delete(v4, v32)
                        _ = v35
                        _ = v36
                    } else {
                        fmt.Println(v22)
                        fmt.Println(v6, v32)
                        var v37 float64 = 1.5
                        if ((^(v22)) < (v22 * 2)) {
                            v38 := GenS1{F0: 81}
                            _ = v38
                        }
                        var v39 uint64 = 0
                        _ = v37
                        _ = v39
                    }
                    _ = v33
                    _ = v34
                } else {
                    var v40 uint = 972
                    v5 = v5
                    _ = v40
                }
                var v41 uint32 = 1
                _ = fi30
                _ = v32
                _ = v41
            }
            _ = fi24
        }
        var v42 int32 = 2147483647
        _ = v42
    } else {
        v1 = 0.0
        v5 = "hello"
        switch (^(v23)) {
        case -9223372036854775808:
            _ = len(v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
