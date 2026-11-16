package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
func genFn0(p0 int16, p1 int16) uint32 {
    fmt.Println(p1)
    fmt.Println(p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := []uint32{4294967295, 42, 543}
    var v2 uint32 = 2147483647
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    v2 -= genFn0(517, 0)
    var v5 uint64 = 18446744073709551615
    var v6 int16 = 1
    if (!false) {
        var v7 int8 = -109
        var v8 int8 = v7
        var v9 uint32 = (v2 | 2)
        _ = v7
        _ = v8
        _ = v9
    } else {
        if len(v1) > 0 {
            v1[0] = v2
        }
        if false {
            var v10 uint8 = 0
            v11 := GenS1{F0: "foo bar", F1: 163, F2: 65535}
            var v12 uint32 = genFn0(v6, v6)
            _ = v10
            _ = v11
            _ = v12
        } else {
            if (false || true) {
                for fi13 := 0; fi13 < 1; fi13++ {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        fmt.Println(fi13, fi14)
                        var v15 uint8 = 0
                        var v16 string = ("" + "")
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 := GenS1{F0: (v16 + (v16 + (((v16 + v16) + (v16 + (v16 + v16))) + (("hello" + v16) + ("" + ""))))), F1: (v5 * v5), F2: 1}
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 10; fi18++ {
                        _ = len(v4)
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                        v6 -= v6
                        v1 = append(v1, v2)
                        var v19 uint64 = 4294967295
                        _ = fi18
                        _ = v19
                    }
                    if false {
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := GenS1{F0: "foo bar", F1: v5, F2: 304}
                            v2 = v2
                            _ = fi20
                            _ = v21
                        }
                    } else {
                        v2 += 1
                        if true {
                            v2 += (v2 & 2147483648)
                            var v22 string = ""
                            _ = v22
                        }
                    }
                    _ = fi13
                }
            }
            switch "foo bar" {
            case "a":
                for fi23 := 0; fi23 < 17; fi23++ {
                    fmt.Println(v6, v5)
                    _ = fi23
                }
            case "":
                for fi24 := 0; fi24 < 17; fi24++ {
                    fmt.Println(v5)
                    _ = fi24
                }
            case "foo bar":
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 int = 45
                    _ = fi25
                    _ = v26
                }
            }
        }
        v27 := GenS1{F0: "a", F1: (^(v5)), F2: 970}
        _ = v27
    }
    v2 += (v2 | v2)
    var v28 uint64 = v5
    if true {
        if true {
            v29 := GenS1{F0: "a", F1: v28, F2: 963}
            v30 := GenS1{F0: "\n\t", F1: (v28 >> ((4294967295) & 63)), F2: 97}
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            if true {
                v32 := v4[287]
                _ = v32
            }
            _ = fi31
        }
    }
    v6 += v6
    var v33 uint64 = v28
    for fi34 := 0; fi34 < 20; fi34++ {
        v1 = append(v1, (v2 * v2))
        delete(v3, fi34)
        _ = fi34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
