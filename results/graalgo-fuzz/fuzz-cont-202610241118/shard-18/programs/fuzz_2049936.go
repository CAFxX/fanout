package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 string }
func genFn0(p0 uint16, p1 uint) uint64 {
    p1 += (p1 & p1)
    for fi4 := 0; fi4 < 14; fi4++ {
        fmt.Println(p0, fi4, fi4)
        _ = fi4
    }
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(p0, p1, fi5)
            _ = fi6
        }
        p1 = p1
        _ = fi5
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        if true {
            if ((!false) || false) {
                switch 128 {
                case 32768:
                    fmt.Println(p1, fi7)
                case 9223372036854775807:
                    switch (("foo bar" + "") + ("foo bar" + ("a" + ("hello" + ("a" + "hello"))))) {
                    case "a":
                        if (-504 != -813) {
                            if ((0.1 < (-2.25 / (0.1 - (1.5 * 0.0)))) && false) {
                                v8 := []int{(fi7 - fi7), -45}
                                _ = v8
                            } else {
                                v9 := []int{fi7, -1}
                                v10 := GenS1{F0: 32767}
                                switch "a" {
                                case "foo bar":
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                default:
                                    p0 -= (p0 & p0)
                                }
                                _ = v10
                                _ = v9
                            }
                        } else {
                            if (true || true) {
                                for fi11 := 0; fi11 < 12; fi11++ {
                                    v12 := GenS2{F0: 2147483647}
                                    _ = fi11
                                    _ = v12
                                }
                                p0 += (^(p0))
                            } else {
                                var v13 uint = uint(p0)
                                var v14 uint16 = p0
                                v15 := []int{fi7, fi7, (fi7 | (fi7 & fi7)), (fi7 * (fi7 / (((fi7 % ((fi7) | 1))) | 1)))}
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            v16 := []int{fi7}
                            _ = v16
                        }
                    case "\n\t":
                        p1 = p1
                    case "":
                        var v17 uint64 = 165
                        _ = v17
                    default:
                        v18 := []int{(fi7 + fi7), 0, int(p0), fi7}
                        _ = v18
                    }
                }
                var v19 uint16 = p0
                _ = v19
            } else {
                v20 := []int{}
                _ = v20
            }
            fmt.Println(p0, fi7, p0)
        }
        _ = fi7
    }
    if ((!true) || (!true)) {
        v21 := GenS2{F0: 1}
        _ = v21
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]uint32)
    v3 := []int{-2147483648, 2147483647, -9223372036854775808}
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 uint32 = 1
        _ = fi22
        _ = v23
    }
    var v24 int = 787
    for fi25 := 0; fi25 < 2; fi25++ {
        if false {
            fmt.Println(v24, fi25)
            v1 = (v1 + 81)
        }
        _ = fi25
    }
    v1 += v1
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint = 18446744073709551615
        _ = fi26
        _ = v27
    }
    v24 = (v24 | v24)
    fmt.Println(v24, v1)
    var v28 uint64 = 1
    v1 = (v1 * (v1 | 2))
    var v29 uint32 = 0
    if (float64(3.14159) <= -2.25) {
        var v30 bool = false
        v31 := GenS1{F0: -32768}
        for fi32 := 0; fi32 < 5; fi32++ {
            v33 := GenS3{F0: 255, F1: (v28 ^ v28), F2: "hello"}
            var v34 int32 = -983
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v30
        _ = v31
    }
    var v35 uint32 = (v29 | v29)
    v36 := v2[("a" + "a")]
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
}
