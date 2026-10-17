package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0(p0 int64, p1 int, p2 uint8) int8 {
    v5 := []int{p1, (p1 * 32768)}
    p0 = (p0 << ((18446744073709551615) & 63))
    v6 := GenS1{F0: 1, F1: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 10
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{4294967295, 50, 107, 42}
    var v3 string = ("" + "\n\t")
    var v4 uint16 = 0
    if (!false) {
        var v7 int = -598
        fmt.Println(v4, v7)
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 bool = true
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v3 = ""
    var v11 float32 = 1.5
    v3 = (v3 + v3)
    fmt.Println(v4, v1)
    var v12 float64 = (-2.25 / (-2.25 - 1.5))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v13 bool = (37 <= genFn0(v1, -659, 2))
    switch v3 {
    case "foo bar":
        var v14 float32 = v11
        _ = v14
    case "":
        switch 2147483647 {
        case 9223372036854775807:
            v3 = "a"
        case 906:
            v11 -= ((0.1 * 0.1) * v11)
        default:
            fmt.Println(v1, v13, v3)
        }
    }
    if (!v13) {
        if v13 {
            var v15 int64 = v1
            for fi16 := 0; fi16 < 11; fi16++ {
                if false {
                    var v17 float32 = (v11 * v11)
                    _ = v17
                }
                _ = fi16
            }
            _ = v15
        } else {
            if true {
                _ = len(v2)
                var v18 uint16 = (v4 ^ v4)
                v4 += v4
                _ = v18
            } else {
                if v13 {
                    switch v3 {
                    case "foo bar":
                        v19 := GenS1{F0: 32767, F1: 2}
                        _ = v19
                    case "":
                        v20 := GenS1{F0: 925, F1: 42}
                        _ = v20
                    case "a":
                        fmt.Println(v4)
                    }
                    var v21 uint32 = 0
                    _ = v21
                }
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v1 -= (v1 | v1)
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            v12 = v12
            fmt.Println(fi23, v3)
            _ = fi23
        }
        if false {
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int16 = -1
                if v13 {
                    switch v3 {
                    case "hello":
                        var v26 int32 = 1
                        _ = v26
                    case "a":
                        if v13 {
                            var v27 uint64 = 415
                            if v13 {
                                fmt.Println(v25)
                                v13 = false
                                var v28 uint64 = (v27 - 1)
                                _ = v28
                            }
                            _ = v27
                        } else {
                            v1 = (v1 % ((v1) | 1))
                            v4 -= (v4 / ((65535) | 1))
                        }
                    }
                    var v29 float64 = v12
                    _ = v29
                } else {
                    if true {
                        var v30 bool = true
                        v2 = append(v2, 964)
                        _ = v30
                    } else {
                        v13 = ((!(((v4 % ((42) | 1)) > 640) || true)) && v13)
                        fmt.Println(fi24, fi22)
                        v2 = append(v2, 42)
                    }
                }
                _ = fi24
                _ = v25
            }
            _ = len(v2)
        } else {
            v13 = v13
        }
        if (float32((v12 + v12)) > float32((-2.25 - (1e10 * (v12 / (float64(((float32(v12) - v11) / 3.14159)) - v12)))))) {
            switch (("a" + v3) + "\n\t") {
            case "a":
                v3 = v3
            case "\n\t":
                var v31 int32 = 42
                _ = v31
            default:
                fmt.Println(v4, v4)
            }
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
