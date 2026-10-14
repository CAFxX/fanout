package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 int32, p1 int64) uint32 {
    if (42 > 1) {
        v7 := GenS1{F0: 4294967295, F1: "", F2: 1}
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        p0 -= (p0 | p0)
        p1 = 1
        var v9 int32 = -1
        _ = fi8
        _ = v9
    }
    p1 -= (p1 ^ p1)
    _ = p0
    _ = p1
    return 2147483648
}
func genFn1(p0 uint32, p1 uint) int32 {
    var v10 uint = 0
    var v11 int8 = -71
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 string = (("hello" + "") + "foo bar")
    v2 := []uint32{65535, 0, 42, 0}
    v3 := make(map[int]uint32)
    v4 := []uint32{329, 2147483648}
    v5 := []string{"\n\t", ((v1 + (v1 + (v1 + ("" + v1)))) + "a")}
    var v6 string = (v1 + v1)
    v12 := v3[32767]
    v4 = append(v4, (v12 / ((315) | 1)))
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v1)
        if false {
            if (0 != 11) {
                switch v6 {
                case "a":
                    if false {
                        var v14 uint8 = 242
                        _ = v14
                    } else {
                        v1 = ((("hello" + v6) + v6) + v6)
                    }
                case "\n\t":
                    var v15 string = "a"
                    _ = v15
                case "":
                    var v16 uint = 567
                    _ = v16
                }
                switch 2147483647 {
                case 128:
                    v17 := GenS1{F0: 571, F1: "hello", F2: v12}
                    _ = v17
                default:
                    var v18 bool = ((v12 + genFn0(42, 1)) == v12)
                    _ = v18
                }
                switch ((v1 + v1) + (((v1 + v6) + "foo bar") + v6)) {
                case "\n\t":
                    v6 = (v1 + v1)
                }
            }
        }
        _ = fi13
    }
    var v19 float64 = (1e10 - 0.1)
    v6 = "foo bar"
    var v20 uint8 = 212
    var v21 int16 = 32767
    var v22 int8 = -3
    for fi23 := 0; fi23 < 10; fi23++ {
        if (float64(((0.1 * (0.0 + float32(1e10))) - float32(v19))) < v19) {
            var v24 uint16 = 1
            if false {
                v3[fi23] = (v12 / (((v12 % (((^(v12))) | 1))) | 1))
            } else {
                if false {
                    if false {
                        for fi25 := 0; fi25 < 13; fi25++ {
                            fmt.Println(v21, v22, v12)
                            fmt.Println(v19, v19, fi23)
                            var v26 uint = 87
                            _ = fi25
                            _ = v26
                        }
                        if false {
                            v22 += (v22 & -38)
                            for fi27 := 0; fi27 < 16; fi27++ {
                                fmt.Println(v21, v6)
                                _ = fi27
                            }
                        } else {
                            var v28 int8 = (v22 + -1)
                            _ = v28
                        }
                    } else {
                        v29 := GenS2{F0: (v24 | (v24 - (v24 % (((v24 % (((v24 + (v24 | v24))) | 1))) | 1)))), F1: -9223372036854775808}
                        var v30 uint8 = v20
                        _ = v29
                        _ = v30
                    }
                }
                fmt.Println(v1, v24)
            }
            _ = v24
        }
        _ = fi23
    }
    v6 = v1
    if ((v20 * v20) <= v20) {
        if true {
            if (((v22 == v22) && (v12 > v12)) && true) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v31 uint8 = (v20 ^ v20)
                _ = v31
            } else {
                for fi32 := 0; fi32 < 4; fi32++ {
                    v33 := v3[(fi32 >> ((42) & 63))]
                    if false {
                        var v34 uint32 = (^(v33))
                        v35 := GenS2{F0: 0, F1: 1}
                        _ = v34
                        _ = v35
                    }
                    _ = fi32
                    _ = v33
                }
                var v36 uint8 = 196
                _ = v36
            }
            for fi37 := 0; fi37 < 4; fi37++ {
                v1 = v1
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 float64 = v19
                    v40 := GenS2{F0: 603, F1: -705}
                    switch fi38 {
                    case -1:
                        v39 = ((v39 * v39) + 0.1)
                    case 127:
                        v41 := GenS2{F0: 65535, F1: 9223372036854775807}
                        _ = v41
                    case 128:
                        var v42 int = (fi37 ^ fi37)
                        _ = v42
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                for fi43 := 0; fi43 < 8; fi43++ {
                    var v44 float64 = v19
                    v20 = v20
                    _ = fi43
                    _ = v44
                }
                _ = fi37
            }
        }
        v6 = v6
        var v45 float64 = 1e10
        _ = v45
    }
    v1 = ""
    v46 := GenS1{F0: 969, F1: (v6 + (v6 + ((v1 + (v1 + ("foo bar" + "foo bar"))) + ""))), F2: 356}
    v4 = append(v4, (v12 + (v12 % ((v12) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
