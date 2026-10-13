package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int8; F2 int32 }
func genFn0(p0 int16, p1 uint64, p2 int16) uint64 {
    p2 += -269
    var v4 uint64 = (p1 * (p1 ^ (p1 & 591)))
    switch int(v4) {
    case 948:
        p0 -= (p0 | 32767)
    case -468:
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := GenS1{F0: (p1 * (p1 | (^(v4)))), F1: 18446744073709551615}
            var v7 float32 = float32(((0.0 - (0.0 + -2.25)) / 1e10))
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        p0 = -170
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return v4
}
func genFn1(p0 int, p1 uint8, p2 string) bool {
    var v8 string = "a"
    switch p0 {
    case 42:
        p2 = ""
    }
    switch p0 {
    case 127:
        v9 := GenS2{F0: 100, F1: "foo bar"}
        _ = v9
    default:
        var v10 uint8 = (p1 + (p1 | p1))
        _ = v10
    }
    if false {
        v11 := GenS3{F0: v8}
        p2 = v8
        var v12 int16 = -1
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 0, F1: 518}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v8
    return true
}

func main() {
    var v1 int = 32767
    var v2 uint64 = 42
    v3 := make(map[uint32]string)
    v2 = genFn0(0, (v2 / ((42) | 1)), 32767)
    var v14 int64 = 42
    v2 = (v2 >> ((255) & 63))
    _ = len(v3)
    var v15 bool = true
    v16 := []int{}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    fmt.Println(v15, v14, v14)
    switch "" {
    case "":
        v2 += (v2 - 0)
    default:
        switch v1 {
        case 32768:
            v15 = v15
        case -928:
            switch v1 {
            case 1:
                if (v15 || v15) {
                    v1 += -9223372036854775808
                }
            case -703:
                var v17 uint16 = 32767
                _ = v17
            default:
                var v18 string = "\n\t"
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 12; fi19++ {
                v15 = true
                if (!true) {
                    if true {
                        v20 := GenS3{F0: "hello"}
                        _ = v20
                    }
                    v21 := GenS3{F0: "hello"}
                    _ = v21
                } else {
                    var v22 uint64 = v2
                    if false {
                        var v23 int16 = 0
                        for fi24 := 0; fi24 < 9; fi24++ {
                            var v25 uint8 = 197
                            v14 = v14
                            v26 := GenS3{F0: ""}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        v1 -= (fi19 | v1)
                        _ = v23
                    }
                    _ = v22
                }
                var v27 uint = 18446744073709551615
                _ = fi19
                _ = v27
            }
        }
    }
    v28 := GenS3{F0: "foo bar"}
    var v29 float64 = (-2.25 * (-2.25 / 1.0))
    var v30 int = v1
    var v31 uint64 = genFn0(42, 499, -32768)
    if (v14 >= -279) {
        var v32 uint32 = 42
        v33 := GenS3{F0: ("" + ("" + ("foo bar" + ("\n\t" + (("a" + "") + "foo bar")))))}
        if true {
            for fi34 := 0; fi34 < 16; fi34++ {
                v30 = fi34
                v29 += 0.1
                _ = fi34
            }
            for fi35 := 0; fi35 < 7; fi35++ {
                v36 := GenS4{F0: "hello", F1: 100, F2: 780}
                _ = fi35
                _ = v36
            }
            v29 = 0.1
        } else {
            var v37 int8 = 0
            fmt.Println(v14)
            _ = v37
        }
        _ = v32
        _ = v33
    }
    if false {
        var v38 uint64 = (v2 / ((v31) | 1))
        _ = v38
    }
    v39 := GenS1{F0: 80, F1: 948}
    switch "a" {
    case "a":
        delete(v3, 443)
    }
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v39
}
