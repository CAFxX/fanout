package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int }
func genFn0(p0 string, p1 int, p2 uint) string {
    v5 := []int{1, -110}
    var v6 uint64 = 2
    v7 := GenS1{F0: -104}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1() int {
    switch (("foo bar" + ("\n\t" + "hello")) + (("hello" + "foo bar") + ("foo bar" + ((("" + (("a" + "\n\t") + "")) + "foo bar") + "a")))) {
    case "hello":
        if true {
            v8 := []int{0}
            var v9 int = 9223372036854775807
            _ = v8
            _ = v9
        }
    case "":
        var v10 uint64 = 1
        _ = v10
    default:
        var v11 uint = 4294967295
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float32 = (3.14159 - (0.0 * (0.0 + 3.14159)))
        _ = fi12
        _ = v13
    }
    switch "hello" {
    case "":
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := []int{598, fi14, -2147483648, (fi14 ^ 42)}
            if true {
                fmt.Println(fi14)
                v16 := GenS1{F0: 42}
                var v17 float64 = ((-2.25 - 1e10) + 1.5)
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 9; fi18++ {
                switch fi14 {
                case 4294967295:
                    var v19 uint64 = 1
                    _ = v19
                case 32767:
                    var v20 string = "a"
                    _ = v20
                case 2147483647:
                    if true {
                        var v21 uint = 1
                        v21 += v21
                        _ = v21
                    }
                default:
                    v15 = append(v15, 834)
                }
                switch "hello" {
                case "hello":
                    fmt.Println(fi14, fi14)
                case "a":
                    var v22 int = (fi14 & fi18)
                    _ = v22
                default:
                    if len(v15) > 0 {
                        v15[0] = fi18
                    }
                }
                if false {
                    if ("hello" > "a") {
                        switch "" {
                        case "foo bar":
                            var v23 string = ""
                            _ = v23
                        case "":
                            if true {
                                var v24 uint32 = 800
                                if (960 == 517) {
                                    var v25 float32 = (0.1 / 0.1)
                                    v24 += (v24 & (v24 | 1))
                                    _ = v25
                                } else {
                                    var v26 uint16 = 925
                                    v27 := GenS2{F0: true, F1: (!(4294967295 != 2))}
                                    _ = v26
                                    _ = v27
                                }
                                if false {
                                    var v28 float64 = float64(-2.25)
                                    v29 := GenS2{F0: true, F1: false}
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    var v30 uint64 = 42
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                _ = v24
                            } else {
                                var v31 string = ""
                                v32 := GenS1{F0: -1}
                                _ = v31
                                _ = v32
                            }
                        case "hello":
                            var v33 int64 = -1
                            _ = v33
                        }
                    } else {
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            v34 := GenS3{F0: -1}
                            _ = v34
                        default:
                            v15 = append(v15, (^(fi18)))
                        }
                    }
                    var v35 uint = 18446744073709551615
                    fmt.Println(fi18, fi18)
                    _ = v35
                } else {
                    var v36 float64 = (1e10 / 1.0)
                    _ = v36
                }
                _ = fi18
            }
            _ = fi14
            _ = v15
        }
    default:
        v37 := GenS4{F0: 42, F1: 32767}
        _ = v37
    }
    var v38 bool = (3.14159 != 1.5)
    v39 := []int{776, 32767, 42, -5}
    _ = v38
    _ = v39
    return -2147483648
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"\n\t", ("a" + "hello")}
    var v3 uint16 = 2
    var v4 uint = 666
    v3 += (v3 | 1)
    v4 = 2
    switch ("" + ("foo bar" + "foo bar")) {
    case "":
        v40 := v1[(((("" + "hello") + "a") + (("hello" + ("a" + "\n\t")) + "foo bar")) + ("\n\t" + "foo bar"))]
        _ = v40
    case "a":
        v41 := GenS2{F0: (false || true), F1: true}
        _ = v41
    }
    if false {
        _ = len(v1)
    } else {
        fmt.Println(v3)
        if (true && (823 <= 18446744073709551615)) {
            if (((2 < 18446744073709551615) || true) || (-130 < genFn1())) {
                v3 = v3
            }
            fmt.Println(v3, v3)
            var v42 string = "a"
            _ = v42
        } else {
            v3 = v3
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        v4 = v4
        var v43 int16 = 0
        v44 := GenS2{F0: true, F1: ((^(v43)) < v43)}
        _ = v43
        _ = v44
    }
    if (!(-251 <= -2147483648)) {
        var v45 int32 = -2147483648
        v46 := GenS5{F0: 127}
        _ = v45
        _ = v46
    }
    var v47 int64 = 0
    for fi48 := 0; fi48 < 7; fi48++ {
        fmt.Println(v47, v3)
        var v49 string = ""
        if (901 < 42) {
            v50 := GenS6{F0: fi48}
            _ = v50
        }
        _ = fi48
        _ = v49
    }
    v3 = (v3 << (((v4 * (^(v4)))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
}
