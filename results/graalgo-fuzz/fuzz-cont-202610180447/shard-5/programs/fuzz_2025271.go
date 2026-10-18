package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 int }
func genFn0(p0 uint16) uint64 {
    p0 = 255
    var v7 string = ""
    if true {
        v7 = (((((v7 + v7) + "a") + (((((((v7 + v7) + v7) + "") + v7) + (v7 + v7)) + "hello") + (v7 + (v7 + v7)))) + "hello") + "a")
        switch "a" {
        case "a":
            var v8 uint16 = p0
            _ = v8
        case "foo bar":
            v7 = v7
        case "\n\t":
            v9 := GenS1{F0: "foo bar", F1: -797}
            _ = v9
        }
        var v10 float64 = float64(0.1)
        _ = v10
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    return 121
}

func main() {
    var v1 int64 = 0
    var v2 float64 = float64(3.14159)
    v3 := []uint32{0, 0}
    v4 := []string{(("hello" + (("" + "foo bar") + (("a" + ("\n\t" + "foo bar")) + ("\n\t" + (("\n\t" + (("a" + "\n\t") + ("" + "foo bar"))) + "hello"))))) + "hello"), ("foo bar" + "")}
    var v5 int16 = 445
    var v6 int = 32767
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(v6, v5, v1)
        _ = len(v4)
        _ = fi11
    }
    v1 = v1
    switch 127 {
    case 9223372036854775807:
        switch v6 {
        case 4294967295:
            switch v6 {
            case 32768:
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 bool = true
                    var v14 uint8 = 2
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case 802:
                if (v1 <= v1) {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        if true {
                            v16 := GenS2{F0: 294}
                            var v17 bool = (!true)
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint16 = 0
                            var v19 float32 = float32(float64(-2.25))
                            _ = v18
                            _ = v19
                        }
                        _ = fi15
                    }
                    v20 := GenS2{F0: 25}
                    switch "\n\t" {
                    case "a":
                        v6 = v6
                    case "\n\t":
                        var v21 uint32 = 2147483648
                        _ = v21
                    }
                    _ = v20
                }
            case 536:
                for fi22 := 0; fi22 < 11; fi22++ {
                    v6 += (fi22 - v6)
                    if len(v4) > 0 {
                        v4[0] = ("a" + "\n\t")
                    }
                    v23 := GenS3{F0: 0, F1: true}
                    _ = fi22
                    _ = v23
                }
            default:
                fmt.Println(v6)
            }
        case -1:
            v24 := GenS4{F0: v6}
            _ = v24
        case 508:
            if (75 != 439) {
                var v25 uint8 = 72
                var v26 float32 = float32((v2 - (float64(0.0) + (v2 - 1e10))))
                _ = v25
                _ = v26
            } else {
                v27 := GenS5{F0: 1}
                _ = v27
            }
        default:
            var v28 int32 = 1
            _ = v28
        }
    case 0:
        var v29 uint = 1
        _ = v29
    }
    v30 := GenS1{F0: ("" + "foo bar"), F1: (v5 + (^(v5)))}
    var v31 int16 = v5
    var v32 int = 4294967295
    v1 = (v1 % (((v1 << ((20) & 63))) | 1))
    var v33 uint = 0
    v33 += (^(v33))
    fmt.Println(v32)
    _ = len(v3)
    if (!(!false)) {
        var v34 int64 = v1
        v1 = -587
        fmt.Println(v33)
        _ = v34
    } else {
        v6 = v6
        switch "" {
        case "":
            if false {
                v35 := GenS3{F0: v33, F1: false}
                v3 = append(v3, 2147483648)
                _ = v35
            } else {
                v36 := GenS6{F0: "foo bar", F1: v32}
                v4 = append(v4, ("foo bar" + "hello"))
                _ = v36
            }
        case "\n\t":
            v2 = (3.14159 - ((float64(float32((((1.5 + (v2 / -2.25)) - 1e10) + v2))) * v2) - v2))
        case "a":
            v5 = (v31 ^ (v5 + v31))
        default:
            if (((^(v1)) <= (v1 + (v1 / (((v1 << ((v33) & 63))) | 1)))) && (true || true)) {
                v6 -= (v32 - (^(v6)))
            }
        }
        var v37 int16 = v31
        _ = v37
    }
    v38 := GenS4{F0: (v6 | -2147483648)}
    if true {
        if len(v3) > 0 {
            v3[0] = 572
        }
        var v39 int = (v32 << (((v33 - 78)) & 63))
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
