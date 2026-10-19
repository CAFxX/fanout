package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 uint16, p1 int8, p2 uint32) uint64 {
    var v5 int64 = 0
    p1 += -4
    var v6 int32 = 991
    var v7 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 uint32, p2 bool) int64 {
    var v8 int = -9223372036854775808
    switch -283 {
    case 9223372036854775807:
        if (1 <= 127) {
            var v9 uint64 = 627
            if (1.5 != 0.1) {
                var v10 int64 = 206
                _ = v10
            }
            _ = v9
        } else {
            switch ("a" + "hello") {
            case "foo bar":
                p1 += 0
            case "a":
                var v11 uint16 = 129
                _ = v11
            case "hello":
                if true {
                    v12 := GenS1{F0: 255, F1: "hello", F2: ""}
                    if p2 {
                        p2 = (1000 < 1)
                    } else {
                        if (("foo bar" < ("" + (((("a" + "a") + ("a" + "a")) + "\n\t") + "hello"))) || p2) {
                            p2 = p2
                            p1 -= p1
                            var v13 int8 = 0
                            _ = v13
                        } else {
                            p2 = p2
                            v14 := []int{4294967295}
                            _ = v14
                        }
                        if p2 {
                            var v15 uint = 2
                            switch "hello" {
                            case "\n\t":
                                var v16 float64 = 1.5
                                _ = v16
                            }
                            _ = v15
                        }
                        fmt.Println(p1)
                    }
                    _ = v12
                } else {
                    if (p2 || (!(-9223372036854775808 >= 0))) {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            p1 = p1
                            v18 := GenS1{F0: 42, F1: "a", F2: "\n\t"}
                            v19 := GenS1{F0: 1, F1: ("hello" + "\n\t"), F2: "a"}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        fmt.Println(v8, v8, v8)
                        switch v8 {
                        case 128:
                            var v20 uint32 = 189
                            _ = v20
                        case 32768:
                            v21 := []int{-283}
                            _ = v21
                        case 642:
                            var v22 string = "a"
                            _ = v22
                        default:
                            v23 := GenS1{F0: 50, F1: "\n\t", F2: ""}
                            _ = v23
                        }
                    } else {
                        v24 := []int{4294967295, (v8 - 2147483647), (-(v8))}
                        v8 = int(p1)
                        _ = v24
                    }
                    var v25 int32 = -370
                    p2 = (!p2)
                    _ = v25
                }
            default:
                if p2 {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        fmt.Println(p0)
                        _ = fi26
                    }
                    p2 = (9223372036854775807 != int(p0))
                    if true {
                        var v27 float32 = 1.5
                        _ = v27
                    } else {
                        switch (^(v8)) {
                        case -465:
                            p2 = (p2 && p2)
                        case -2147483648:
                            var v28 uint32 = 1
                            _ = v28
                        }
                        if p2 {
                            var v29 string = ("hello" + "foo bar")
                            if (32767 != 32767) {
                                var v30 uint16 = uint16(p0)
                                if ((-154 != -2147483648) && ((0.0 * 1.5) == float64(1e10))) {
                                    v31 := GenS1{F0: uint8(p1), F1: (v29 + (v29 + v29)), F2: "a"}
                                    if p2 {
                                        p1 = (p1 - p0)
                                    }
                                    _ = v31
                                }
                                switch (v29 + v29) {
                                case "":
                                    p1 = p0
                                }
                                _ = v30
                            } else {
                                v32 := []int{32768, 4294967295}
                                v33 := GenS1{F0: 255, F1: v29, F2: "foo bar"}
                                if len(v32) > 0 {
                                    v32[0] = v8
                                }
                                _ = v32
                                _ = v33
                            }
                            v34 := []int{127}
                            _ = v29
                            _ = v34
                        } else {
                            p1 += p0
                        }
                    }
                } else {
                    var v35 int8 = 127
                    _ = v35
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 597
    v3 := make(map[string]uint32)
    var v4 string = "foo bar"
    v36 := []int{32768, 0, 773}
    _ = len(v36)
    v2 -= v2
    v37 := GenS1{F0: 127, F1: v4, F2: v4}
    if true {
        v38 := GenS2{F0: 65535, F1: "a"}
        _ = v38
    }
    v2 = v2
    v2 += v2
    v2 += 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
