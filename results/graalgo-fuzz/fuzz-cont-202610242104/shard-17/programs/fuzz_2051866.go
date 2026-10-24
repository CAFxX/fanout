package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 int32, p2 int32) int8 {
    v6 := GenS1{F0: 255, F1: true}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 uint64 = 791
    var v2 int32 = 0
    var v3 uint8 = 42
    v4 := make(map[uint32]uint32)
    v5 := []string{"a", "", "a", "\n\t"}
    var v8 float64 = 0.0
    fmt.Println(v8, v3, v3)
    var v9 int16 = 1
    if true {
        var v10 float32 = float32(0.0)
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint32 = 65535
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v13 uint32 = 732
    v13 = (v13 & (^(v13)))
    v14 := GenS2{F0: v13, F1: ""}
    v5 = append(v5, "foo bar")
    var v15 int8 = genFn0(42, 1, -2147483648)
    fmt.Println(v9, v13, v2)
    if false {
        for fi16 := 0; fi16 < 16; fi16++ {
            if ((float32((float64((float32(1.5) - 0.0)) + v8)) + (-2.25 / 1.5)) <= -2.25) {
                switch "foo bar" {
                case "a":
                    v17 := GenS2{F0: v13, F1: ("foo bar" + "foo bar")}
                    _ = v17
                case "foo bar":
                    if false {
                        for fi18 := 0; fi18 < 2; fi18++ {
                            switch 472 {
                            case -1:
                                if (((0.0 + (v8 - v8)) / (0.0 * v8)) >= 1e10) {
                                    var v19 uint16 = 255
                                    _ = v19
                                } else {
                                    fmt.Println(v3, v9, v9)
                                    v15 += (v15 - int8(v1))
                                    v20 := GenS3{F0: (v13 & 911), F1: 1}
                                    _ = v20
                                }
                            case 4294967295:
                                switch int(v15) {
                                case -2147483648:
                                    v21 := GenS1{F0: 578, F1: true}
                                    _ = v21
                                case 450:
                                    v22 := GenS4{F0: 0, F1: (true || (genFn0(v15, 42, int32(v15)) <= 42))}
                                    _ = v22
                                case 452:
                                    var v23 int16 = (v9 % ((v9) | 1))
                                    _ = v23
                                }
                                var v24 string = "foo bar"
                                fmt.Println(v2, v8)
                                _ = v24
                            }
                            _ = fi18
                        }
                        _ = len(v4)
                        var v25 int = 0
                        _ = v25
                    } else {
                        var v26 uint64 = v1
                        if ((v2 - v2) > v2) {
                            v27 := GenS2{F0: (v13 % ((v13) | 1)), F1: "hello"}
                            var v28 float32 = (((1e10 * (0.1 - float32(((float64(1.5) + v8) + (0.0 * v8))))) * (float32(1.5) + 1.5)) / (0.0 + 3.14159))
                            if ((!true) && false) {
                                var v29 int = int(v9)
                                var v30 string = "a"
                                for fi31 := 0; fi31 < 12; fi31++ {
                                    v2 = v2
                                    _ = fi31
                                }
                                _ = v29
                                _ = v30
                            } else {
                                var v32 uint32 = v13
                                fmt.Println(v32, v1, v1)
                                _ = v32
                            }
                            _ = v27
                            _ = v28
                        } else {
                            switch fi16 {
                            case -2147483648:
                                fmt.Println(v8, v3)
                            case 289:
                                var v33 uint32 = (v13 % ((2147483648) | 1))
                                _ = v33
                            case 692:
                                if true {
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        fmt.Println(v8, v15, fi34)
                                        _ = fi34
                                    }
                                }
                            default:
                                var v35 float32 = 1e10
                                _ = v35
                            }
                            var v36 int32 = v2
                            var v37 uint64 = (v26 >> ((42) & 63))
                            _ = v36
                            _ = v37
                        }
                        v8 -= float64((1.5 * 0.0))
                        _ = v26
                    }
                default:
                    var v38 int = fi16
                    _ = v38
                }
            }
            for fi39 := 0; fi39 < 18; fi39++ {
                switch 42 {
                case 9223372036854775807:
                    v40 := GenS5{F0: ((("foo bar" + ("foo bar" + ("a" + "hello"))) + "a") + "a")}
                    _ = v40
                case 4294967295:
                    v41 := GenS2{F0: (^(v13)), F1: "foo bar"}
                    _ = v41
                default:
                    switch -1 {
                    case 466:
                        if (((false && (!(725 >= int64(v1)))) || false) || (v3 != (v3 / (((^(v3))) | 1)))) {
                            if false {
                                _ = len(v4)
                                v1 -= v1
                                switch 0 {
                                case 315:
                                    v1 += v1
                                case 334:
                                    var v42 int8 = genFn0(29, 2147483647, (v2 - v2))
                                    _ = v42
                                default:
                                    switch (fi39 / (((fi39 + (fi16 | 694))) | 1)) {
                                    case 9223372036854775807:
                                        v43 := v4[v13]
                                        _ = v43
                                    }
                                }
                            }
                        }
                    default:
                        _ = len(v5)
                    }
                }
                fmt.Println(fi39)
                _ = fi39
            }
            _ = fi16
        }
        for fi44 := 0; fi44 < 7; fi44++ {
            _ = len(v5)
            v3 += v3
            _ = len(v4)
            _ = fi44
        }
        for fi45 := 0; fi45 < 2; fi45++ {
            delete(v4, v13)
            _ = fi45
        }
    } else {
        var v46 float64 = (float64(float32(v8)) - v8)
        var v47 int32 = v2
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v46
        _ = v47
    }
    v1 = 1
    var v48 float32 = 0.1
    v49 := GenS4{F0: 42, F1: (((!(int64(v2) != 9223372036854775807)) && false) && ((false && ((1.5 + v8) > (v8 * ((-2.25 / 0.1) - (v8 + 1e10))))) || true))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v8
    _ = v9
}
