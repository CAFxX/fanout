package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint8; F2 string }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 string; F1 uint16 }
func genFn0(p0 string, p1 int64) uint16 {
    if false {
        v7 := []int{32767, 32768, 1}
        var v8 bool = (143 == 255)
        v9 := GenS1{F0: 2}
        _ = v7
        _ = v8
        _ = v9
    }
    p1 += p1
    v10 := []int{-9223372036854775808}
    p0 = p0
    if len(v10) > 0 {
        v10[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = v10
    return 65535
}
func genFn1(p0 uint16, p1 int16, p2 uint64) bool {
    var v11 int8 = 42
    switch (("a" + "a") + "a") {
    case "foo bar":
        v11 = v11
    case "":
        v12 := []int{-639}
        _ = v12
    default:
        p0 += genFn0("hello", int64(p1))
    }
    v13 := []int{1, int(v11), -1, 923}
    var v14 uint8 = 1
    var v15 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    return false
}

func main() {
    v1 := []int{4294967295}
    v2 := make(map[string]bool)
    v3 := []uint32{1}
    var v4 uint32 = 251
    var v5 uint8 = 89
    v6 := []string{"hello", "foo bar", ((("a" + "") + "a") + ("a" + "hello")), (((("" + ("foo bar" + "")) + "\n\t") + "foo bar") + "a")}
    _ = len(v6)
    v16 := GenS2{F0: 759, F1: 4294967295}
    switch "a" {
    case "foo bar":
        var v17 string = "hello"
        _ = v17
    }
    fmt.Println(v5)
    v18 := GenS2{F0: 574, F1: 292}
    v5 += (^(v5))
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 float64 = 1e10
        fmt.Println(fi19)
        var v21 int = 127
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: -144}
    fmt.Println(v5, v4)
    v23 := GenS4{F0: -2147483648, F1: 1}
    fmt.Println(v4, v4)
    fmt.Println(v5, v4)
    switch -87 {
    case -9223372036854775808:
        v24 := GenS4{F0: -524, F1: (v4 / ((1) | 1))}
        _ = v24
    }
    var v25 uint64 = 334
    var v26 int8 = 0
    v27 := GenS3{F0: 42}
    switch ("a" + (("a" + "a") + "\n\t")) {
    case "hello":
        if false {
            if false {
                v28 := GenS5{F0: v26}
                v29 := GenS6{F0: int(v26), F1: 0, F2: ""}
                var v30 uint16 = 65535
                _ = v28
                _ = v29
                _ = v30
            }
            var v31 float32 = 0.0
            _ = v31
        } else {
            for fi32 := 0; fi32 < 11; fi32++ {
                v1 = append(v1, fi32)
                _ = fi32
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                switch "\n\t" {
                case "a":
                    fmt.Println(v4, v5)
                case "\n\t":
                    switch "foo bar" {
                    case "foo bar":
                        if ((("" + "hello") + "") <= ("hello" + ("\n\t" + "foo bar"))) {
                            for fi34 := 0; fi34 < 9; fi34++ {
                                switch "" {
                                case "\n\t":
                                    switch -358 {
                                    case -2147483648:
                                        var v35 int64 = 513
                                        _ = v35
                                    case -258:
                                        v25 -= (^(v25))
                                    case 42:
                                        delete(v2, ("" + ("\n\t" + "foo bar")))
                                    }
                                case "hello":
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = 1
                                        }
                                        v26 -= -1
                                        var v36 float64 = 1.5
                                        _ = v36
                                    }
                                case "a":
                                    var v37 int32 = -2147483648
                                    _ = v37
                                default:
                                    var v38 bool = (-1 != 0)
                                    _ = v38
                                }
                                fmt.Println(v5)
                                _ = fi34
                            }
                        }
                    case "hello":
                        fmt.Println(v26, v26)
                    default:
                        if true {
                            v25 = 453
                            for fi39 := 0; fi39 < 19; fi39++ {
                                v40 := GenS5{F0: (-(v26))}
                                v41 := GenS7{F0: 9223372036854775807, F1: fi33}
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            v42 := GenS8{F0: ("\n\t" + "foo bar"), F1: 2}
                            _ = v42
                        }
                    }
                }
                var v43 float32 = 0.0
                _ = fi33
                _ = v43
            }
        }
    case "":
        v6 = append(v6, ("hello" + ("hello" + "\n\t")))
    default:
        var v44 float64 = (((3.14159 / 1.0) - -2.25) + 3.14159)
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
