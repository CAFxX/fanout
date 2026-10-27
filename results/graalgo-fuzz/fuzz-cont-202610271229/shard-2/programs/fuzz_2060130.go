package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint }
func genFn0(p0 uint32, p1 bool, p2 int32) uint16 {
    var v7 string = ("" + "\n\t")
    if false {
        p1 = p1
        v7 = "a"
        if p1 {
            v8 := GenS1{F0: v7}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 2; fi9++ {
                v7 = v7
                v10 := []int{-795, fi9, 0, -1}
                if false {
                    if false {
                        v11 := GenS1{F0: (v7 + v7)}
                        _ = v11
                    }
                    v12 := GenS1{F0: ""}
                    if (!(float64(0.0) <= ((3.14159 * (float64(-2.25) / 0.1)) * 3.14159))) {
                        v13 := GenS1{F0: ""}
                        if false {
                            switch fi9 {
                            case 42:
                                if p1 {
                                    v14 := GenS1{F0: (v7 + v7)}
                                    _ = len(v10)
                                    switch (v7 + "foo bar") {
                                    case "a":
                                        if len(v10) > 0 {
                                            v10[0] = int(p2)
                                        }
                                    default:
                                        v15 := GenS2{F0: 42, F1: p1}
                                        _ = v15
                                    }
                                    _ = v14
                                }
                            default:
                                if (p1 || true) {
                                    switch (^(fi9)) {
                                    case 169:
                                        for fi16 := 0; fi16 < 8; fi16++ {
                                            p0 -= (p0 << ((198) & 63))
                                            var v17 uint16 = 653
                                            _ = fi16
                                            _ = v17
                                        }
                                    }
                                }
                            }
                        }
                        _ = v13
                    } else {
                        var v18 int64 = 0
                        _ = v18
                    }
                    _ = v12
                } else {
                    p0 = (p0 & p0)
                    var v19 uint8 = 42
                    p1 = false
                    _ = v19
                }
                _ = fi9
                _ = v10
            }
        }
    } else {
        var v20 float64 = float64((float32(-2.25) + 0.0))
        var v21 string = "hello"
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 774
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = v1
    var v3 uint8 = 0
    v4 := []string{"a", (("a" + "hello") + ("foo bar" + "foo bar")), "a"}
    v5 := make(map[string]bool)
    var v6 uint8 = 127
    if true {
        v6 -= (^(v3))
    }
    v22 := GenS2{F0: 9223372036854775807, F1: ((!true) && (297 == 218))}
    delete(v5, "\n\t")
    var v23 int64 = -9223372036854775808
    switch -216 {
    case -1:
        v2 = v2
    case 970:
        if (!false) {
            var v24 int8 = 1
            _ = v24
        } else {
            v25 := GenS3{F0: 0}
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 uint64 = (v2 >> ((0) & 63))
                v28 := GenS4{F0: 1, F1: 58}
                var v29 uint64 = (v27 | 171)
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            switch -1 {
            case 32768:
                v2 = 18446744073709551615
            case 0:
                _ = len(v5)
            case 240:
                var v30 float64 = (0.1 - 0.1)
                _ = v30
            default:
                if (false && ((v2 * (^(v2))) < 749)) {
                    v31 := GenS4{F0: -612, F1: -128}
                    if len(v4) > 0 {
                        v4[0] = ("a" + "foo bar")
                    }
                    _ = v31
                } else {
                    fmt.Println(v1)
                }
            }
            _ = v25
        }
    default:
        v32 := GenS5{F0: 370, F1: v23, F2: uint(v6)}
        _ = v32
    }
    v6 -= (v6 | v3)
    _ = len(v4)
    if (!false) {
        v3 += (v3 * 127)
    } else {
        var v33 string = ""
        for fi34 := 0; fi34 < 15; fi34++ {
            if (0.0 == (-2.25 + 1e10)) {
                var v35 uint = 2
                _ = v35
            }
            v4 = append(v4, ("hello" + v33))
            _ = fi34
        }
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
