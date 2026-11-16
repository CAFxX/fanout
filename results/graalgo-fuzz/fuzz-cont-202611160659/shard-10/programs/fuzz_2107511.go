package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) bool {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{(-(fi7)), fi7, fi7, (fi7 ^ (^(fi7)))}
        fmt.Println(fi7, fi7)
        _ = fi7
        _ = v8
    }
    var v9 int32 = -2147483648
    p0 = (^(p0))
    switch 4294967295 {
    case 32767:
        for fi10 := 0; fi10 < 11; fi10++ {
            if (!true) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := []int{(fi11 ^ (fi11 + (fi11 | (fi10 / ((999) | 1))))), int(v9), (fi11 % (((fi10 / (((fi11 >> ((255) & 63))) | 1))) | 1)), fi10}
                    fmt.Println(p0, fi10)
                    v9 = v9
                    _ = fi11
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int16 = 678
                _ = fi13
                _ = v14
            }
            _ = fi10
        }
    case 1:
        v15 := []int{4294967295}
        _ = v15
    default:
        v16 := []int{1, 0, 32767}
        _ = v16
    }
    var v17 uint = 2
    _ = p0
    _ = v17
    _ = v9
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 uint = 1
    v6 := []uint32{405, 538, 2147483647}
    v2[-2147483648] = -1
    if (!false) {
        v6 = append(v6, 1)
    }
    v18 := GenS1{F0: 466}
    _ = len(v6)
    if true {
        for fi19 := 0; fi19 < 5; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := GenS2{F0: 675}
                _ = fi20
                _ = v21
            }
            fmt.Println(fi19, v5)
            var v22 float64 = -2.25
            _ = fi19
            _ = v22
        }
        switch 687 {
        case 4294967295:
            v5 += v5
        case 377:
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                for fi23 := 0; fi23 < 11; fi23++ {
                    if (!((3.14159 >= 1.5) && true)) {
                        v5 = v5
                        _ = len(v2)
                    }
                    if ((9223372036854775807 >= 0) || true) {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v3[756] = (^(fi23))
                            var v25 float64 = float64(1.5)
                            v25 += (v25 * 1.5)
                            _ = fi24
                            _ = v25
                        }
                        v5 = 18446744073709551615
                    }
                    if true {
                        if false {
                            v5 = v5
                        } else {
                            v5 += (v5 | v5)
                            for fi26 := 0; fi26 < 2; fi26++ {
                                if true {
                                    var v27 int8 = 100
                                    fmt.Println(v27)
                                    _ = v27
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi26
                            }
                            var v28 int = 876
                            _ = v28
                        }
                        var v29 int32 = int32(v5)
                        _ = v29
                    } else {
                        fmt.Println(fi23, v5)
                    }
                    _ = fi23
                }
            case "":
                var v30 float64 = ((((1.5 / 1.0) * float64((((-2.25 / 1e10) + -2.25) + 3.14159))) - -2.25) + ((0.1 + 0.0) / 1.0))
                _ = v30
            default:
                var v31 int64 = 9223372036854775807
                _ = v31
            }
        }
        var v32 uint8 = 2
        _ = v32
    }
    var v33 float64 = -2.25
    var v34 int16 = 0
    for fi35 := 0; fi35 < 14; fi35++ {
        v34 += v34
        v5 -= (^(v5))
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
