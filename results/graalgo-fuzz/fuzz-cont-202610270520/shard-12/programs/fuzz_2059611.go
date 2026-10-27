package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 bool; F2 bool }
type GenS7 struct { F0 uint8; F1 uint; F2 string }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{-2147483648, 42, 695, -76}
    var v3 bool = ((212 < -881) || false)
    var v4 float64 = 1e10
    var v5 uint32 = 2147483648
    v6 := make(map[string]uint32)
    switch "" {
    case "\n\t":
        var v7 uint = 1
        _ = v7
    case "a":
        if ((!v3) && false) {
            fmt.Println(v5)
            var v8 uint64 = v1
            _ = v8
        }
    case "":
        v9 := v6[("hello" + ("a" + "a"))]
        _ = v9
    }
    fmt.Println(v1, v4, v5)
    v10 := GenS1{F0: "\n\t", F1: 18}
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: "hello", F1: -1}
        v13 := GenS1{F0: "a", F1: 0}
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = (^(v5))
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int8 = -128
            _ = fi15
            _ = v16
        }
        if true {
            if v3 {
                v17 := GenS1{F0: ((("" + ((("\n\t" + ("" + ("hello" + ((("foo bar" + "foo bar") + ("" + ("" + ""))) + ("" + "hello"))))) + "") + "hello")) + "foo bar") + ("\n\t" + "hello")), F1: 100}
                if (4294967295 > 2) {
                    v18 := GenS1{F0: ((("" + "\n\t") + "a") + ("foo bar" + "")), F1: 60}
                    v2 = append(v2, fi14)
                    if (v1 != (v1 + (v1 >> ((353) & 63)))) {
                        if v3 {
                            _ = len(v6)
                            _ = len(v2)
                        } else {
                            if ((991 < 1) && (true && v3)) {
                                for fi19 := 0; fi19 < 2; fi19++ {
                                    v20 := GenS2{F0: (^(v5)), F1: (v1 << ((65535) & 63))}
                                    switch ("\n\t" + ((((("hello" + "hello") + ("a" + "a")) + (("hello" + "a") + "hello")) + ((("" + ("hello" + "foo bar")) + "hello") + "hello")) + "\n\t")) {
                                    case "\n\t":
                                        if v3 {
                                            v4 = 0.0
                                            v1 = 949
                                        }
                                    default:
                                        if (!(!v3)) {
                                            v5 += (^(v5))
                                            v21 := GenS2{F0: (^(v5)), F1: v1}
                                            _ = v21
                                        } else {
                                            if v3 {
                                                var v22 uint16 = 0
                                                _ = v22
                                            } else {
                                                if (42 <= 1) {
                                                    v4 += 1.5
                                                    _ = len(v2)
                                                }
                                            }
                                            if (!(v3 || true)) {
                                                v23 := GenS1{F0: "\n\t", F1: 103}
                                                var v24 bool = false
                                                var v25 float32 = 1.5
                                                var v26 int16 = 776
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v27 := GenS2{F0: v5, F1: (v1 << ((3) & 63))}
                                                v28 := GenS3{F0: ("\n\t" + "a")}
                                                switch ("\n\t" + "") {
                                                case "a":
                                                    var v29 int64 = 9223372036854775807
                                                    _ = v29
                                                case "":
                                                    v30 := GenS4{F0: 0}
                                                    _ = v30
                                                default:
                                                    if (v5 > (v5 >> ((0) & 63))) {
                                                        v31 := GenS4{F0: 366}
                                                        var v32 float64 = (-2.25 / v4)
                                                        fmt.Println(v5, fi19, v3)
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                                switch fi14 {
                                                case 2147483647:
                                                    fmt.Println(v4)
                                                case 891:
                                                    var v33 int = fi14
                                                    _ = v33
                                                default:
                                                    v34 := GenS2{F0: 0, F1: v1}
                                                    _ = v34
                                                }
                                                _ = v27
                                                _ = v28
                                            }
                                            v35 := GenS5{F0: (fi19 - (fi19 - 42))}
                                            _ = v35
                                        }
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                                v36 := GenS6{F0: false, F1: (v3 || true), F2: (!v3)}
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 1; fi37++ {
                                fmt.Println(fi37, v4)
                                _ = fi37
                            }
                        }
                        var v38 uint = 695
                        v39 := GenS7{F0: 128, F1: v38, F2: ("" + "foo bar")}
                        _ = v38
                        _ = v39
                    }
                    _ = v18
                }
                for fi40 := 0; fi40 < 12; fi40++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v41 float32 = 3.14159
                    var v42 int = (fi40 | fi40)
                    var v43 uint32 = v5
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = v17
            }
            v2 = append(v2, (^(fi14)))
            for fi44 := 0; fi44 < 12; fi44++ {
                fmt.Println(v5)
                _ = fi44
            }
        } else {
            v1 -= (v1 << ((0) & 63))
        }
        _ = fi14
    }
    v5 -= (v5 ^ (^(v5)))
    var v45 int64 = 869
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v46 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
