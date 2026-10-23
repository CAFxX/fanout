package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0(p0 bool) uint64 {
    v4 := GenS1{F0: 882}
    v5 := []int{-2147483648, -9223372036854775808, -2147483648, -173}
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := []uint32{2, 337, 780, 42}
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    v6 := GenS1{F0: -2147483648}
    if (1 != 65535) {
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(fi7, v2)
            v2 = 0
            _ = fi7
        }
    } else {
        fmt.Println(v2)
        if false {
            var v8 int = 275
            v8 -= int(v2)
            _ = v8
        } else {
            v2 += (v2 / (((^(v2))) | 1))
        }
    }
    v9 := GenS1{F0: 2147483647}
    switch "" {
    case "\n\t":
        v2 += v2
    case "foo bar":
        fmt.Println(v2)
    default:
        _ = len(v1)
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
        _ = fi10
    }
    v11 := GenS2{F0: (32767 < -1), F1: 32767, F2: 128}
    if len(v1) > 0 {
        v1[0] = 42
    }
    switch "\n\t" {
    case "":
        var v12 int16 = -1
        _ = v12
    case "\n\t":
        v1 = append(v1, 2)
    case "a":
        v2 = v2
    }
    fmt.Println(v2)
    if (0 != 530) {
        v13 := GenS1{F0: 618}
        _ = v13
    }
    if ((-218 == -2147483648) || true) {
        v14 := GenS2{F0: (true || (float32(-2.25) >= (1.5 + (((-2.25 * 1e10) / ((1e10 / 1.0) - float32((0.0 + 1.5)))) * -2.25)))), F1: 454, F2: 410}
        var v15 int64 = 9223372036854775807
        if (false && false) {
            var v16 uint = 0
            var v17 uint32 = 4294967295
            v17 += v17
            _ = v16
            _ = v17
        }
        _ = v14
        _ = v15
    } else {
        var v18 int32 = 699
        _ = v18
    }
    v19 := v3[("" + (("hello" + ("" + "\n\t")) + "a"))]
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint64 = 1
        v21 = 4294967295
        for fi22 := 0; fi22 < 10; fi22++ {
            v21 += genFn0(true)
            switch "\n\t" {
            case "hello":
                v3["a"] = v19
            case "foo bar":
                if (false || (v19 || ("" == ""))) {
                    var v23 uint8 = 42
                    if (true || false) {
                        switch fi20 {
                        case -1:
                            if (-979 == 1000) {
                                if (v2 == v2) {
                                    v24 := GenS2{F0: (27 == 2147483647), F1: 2, F2: (-(fi20))}
                                    _ = v24
                                }
                                var v25 int = 42
                                _ = len(v3)
                                switch "foo bar" {
                                case "a":
                                    var v26 float64 = -2.25
                                    _ = v26
                                case "hello":
                                    if (1 > 1) {
                                        switch (v25 & 32768) {
                                        case 42:
                                            fmt.Println(v21)
                                        case 128:
                                            v27 := GenS1{F0: 2147483647}
                                            _ = v27
                                        }
                                    } else {
                                        v2 = uint(v21)
                                        if v19 {
                                            var v28 int8 = 127
                                            if ((42 == 32767) && (int32(v2) < -2147483648)) {
                                                _ = len(v3)
                                                v28 = (v28 * (v28 | v28))
                                                var v29 uint = v2
                                                var v30 int16 = 42
                                                v29 = (v2 ^ v2)
                                                var v31 uint32 = 0
                                                var v32 uint = (^(v2))
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            } else {
                                                if v19 {
                                                    fmt.Println(v28, v23, v28)
                                                } else {
                                                    v25 = (fi22 ^ fi20)
                                                }
                                                v25 += 42
                                            }
                                            _ = v28
                                        } else {
                                            _ = len(v1)
                                        }
                                    }
                                }
                                _ = v25
                            }
                            var v33 uint64 = (v21 << ((v2) & 63))
                            var v34 int = (fi20 ^ (fi20 << (((^(v2))) & 63)))
                            v19 = v19
                            var v35 uint32 = 2147483647
                            _ = v33
                            _ = v34
                            _ = v35
                        case 2147483647:
                            v36 := v3[("foo bar" + ("\n\t" + "hello"))]
                            _ = v36
                        default:
                            v21 += 509
                        }
                    } else {
                        if true {
                            v37 := GenS3{F0: "\n\t", F1: (("foo bar" + "") + (("hello" + "a") + "a")), F2: -319}
                            v38 := GenS3{F0: "", F1: "a", F2: 32767}
                            delete(v3, "foo bar")
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS2{F0: v19, F1: 203, F2: fi20}
                            delete(v3, ("foo bar" + "\n\t"))
                            _ = v39
                        }
                        v40 := GenS4{F0: true, F1: 96, F2: (v23 & v23)}
                        _ = v40
                    }
                    v41 := GenS3{F0: "a", F1: ("a" + "hello"), F2: fi20}
                    _ = v23
                    _ = v41
                }
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    v2 -= 2
    var v42 uint8 = 108
    v2 = v2
    v19 = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v42
    _ = v6
    _ = v9
}
