package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8 }
func genFn0() int {
    var v7 int8 = 56
    v7 = v7
    v7 += 0
    v7 = (-(v7))
    v7 = (v7 + (-(v7)))
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 uint8) bool {
    switch "" {
    case "\n\t":
        p1 = (^(p1))
    case "hello":
        p1 += (^(p1))
    case "":
        switch ("" + ("" + ("hello" + (("hello" + "\n\t") + "\n\t")))) {
        case "":
            var v8 int16 = 0
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: 32767}
            if (!true) {
                p0 += (p0 << ((996) & 63))
            }
            _ = fi9
            _ = v10
        }
    }
    p0 += 2
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 uint32 = 2
    v2 := []string{(("a" + "") + "foo bar"), ("a" + "foo bar")}
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 float32 = float32(0.1)
    var v11 float32 = v6
    for fi12 := 0; fi12 < 5; fi12++ {
        v6 = 3.14159
        if false {
            _ = len(v4)
        } else {
            for fi13 := 0; fi13 < 9; fi13++ {
                switch "foo bar" {
                case "a":
                    if ((42 <= 124) && (!true)) {
                        if true {
                            switch 4294967295 {
                            case 1:
                                _ = len(v4)
                            default:
                                v14 := GenS2{F0: 0, F1: 18446744073709551615}
                                _ = v14
                            }
                        }
                        v6 += 1.5
                    }
                case "":
                    v6 += (v11 * float32(0.0))
                case "\n\t":
                    v15 := GenS2{F0: 169, F1: 65535}
                    _ = v15
                }
                var v16 string = "foo bar"
                v17 := GenS1{F0: 1000}
                _ = fi13
                _ = v16
                _ = v17
            }
        }
        _ = fi12
    }
    var v18 float64 = -2.25
    switch ("a" + "\n\t") {
    case "\n\t":
        switch "a" {
        case "":
            v6 += v11
        case "foo bar":
            if (true && true) {
                if (3 > 42) {
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    for fi19 := 0; fi19 < 11; fi19++ {
                        if false {
                            for fi20 := 0; fi20 < 12; fi20++ {
                                switch fi19 {
                                case 525:
                                    v1 = v1
                                case 42:
                                    var v21 uint16 = 722
                                    _ = v21
                                default:
                                    v2 = append(v2, ("a" + "\n\t"))
                                }
                                _ = fi20
                            }
                            v18 += float64(float32(float64(v6)))
                            if (815 == 1000) {
                                if (-1 != 972) {
                                    v2 = append(v2, "foo bar")
                                }
                            } else {
                                var v22 bool = true
                                v6 -= v11
                                var v23 int32 = 42
                                _ = v22
                                _ = v23
                            }
                        } else {
                            switch "hello" {
                            case "a":
                                var v24 float64 = v18
                                _ = v24
                            case "\n\t":
                                var v25 bool = false
                                _ = v25
                            case "foo bar":
                                var v26 uint32 = 42
                                _ = v26
                            default:
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v1, v1, v1)
                                case "":
                                    for fi27 := 0; fi27 < 20; fi27++ {
                                        v28 := GenS3{F0: -1, F1: 0}
                                        var v29 int32 = 1
                                        fmt.Println(v18)
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                case "hello":
                                    v18 = v18
                                }
                            }
                        }
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v31 := GenS3{F0: -1, F1: 131}
                            fmt.Println(fi30, v11, v6)
                            _ = fi30
                            _ = v31
                        }
                        _ = fi19
                    }
                }
                switch -333 {
                case -611:
                    v1 -= (v1 | 2147483647)
                default:
                    for fi32 := 0; fi32 < 4; fi32++ {
                        var v33 float64 = (float64(((0.0 / v11) + v11)) + (float64(1e10) * v18))
                        var v34 int64 = 795
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                }
                var v35 int32 = 0
                _ = v35
            }
        }
    case "":
        v11 -= 1.5
    case "hello":
        for fi36 := 0; fi36 < 16; fi36++ {
            if (((float64(v11) / 1.5) + (1.5 * (-2.25 / (v18 + (float64(0.0) + v18))))) < 0.0) {
                if ((!true) || ((true && (false || true)) && (403 > 1))) {
                    fmt.Println(v1)
                } else {
                    var v37 int32 = 0
                    _ = v37
                }
            } else {
                v38 := v3[("hello" + "a")]
                _ = v38
            }
            _ = fi36
        }
    }
    delete(v3, (("hello" + "foo bar") + ""))
    var v39 int8 = 0
    v40 := GenS3{F0: -1, F1: 1}
    v6 += -2.25
    fmt.Println(v6, v18, v18)
    v6 -= float32((v18 / v18))
    for fi41 := 0; fi41 < 6; fi41++ {
        v5[1] = v1
        fmt.Println(v1, fi41)
        for fi42 := 0; fi42 < 18; fi42++ {
            v1 = v1
            v43 := GenS1{F0: -1}
            var v44 uint32 = 4294967295
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi41
    }
    v45 := GenS4{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v6
}
