package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 int64) bool {
    if (!("\n\t" >= "")) {
        p0 = (p0 ^ (p0 >> ((710) & 63)))
        if true {
            var v6 int = 32768
            v6 -= v6
            var v7 uint32 = 142
            _ = v6
            _ = v7
        } else {
            var v8 int = -9223372036854775808
            _ = v8
        }
        var v9 int16 = 1
        _ = v9
    } else {
        v10 := []int{}
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        p0 = p0
        p0 = p0
        _ = fi11
    }
    p0 = p0
    var v12 float32 = 1e10
    v13 := []int{9223372036854775807, -2147483648, 0, 127}
    _ = p0
    _ = v12
    _ = v13
    return (-32768 >= 936)
}
func genFn1(p0 string, p1 uint8) int {
    v14 := []int{42, 9223372036854775807, -2147483648, 455}
    var v15 uint64 = 4294967295
    if true {
        var v16 int64 = -1
        _ = v16
    }
    if true {
        if true {
            v17 := GenS1{F0: -1}
            _ = v17
        } else {
            var v18 int32 = 116
            var v19 int64 = -330
            switch p0 {
            case "":
                v15 = 1
            }
            _ = v18
            _ = v19
        }
    } else {
        var v20 uint64 = (v15 / ((1) | 1))
        switch 349 {
        case 0:
            v15 = 4294967295
        case 84:
            v21 := GenS1{F0: 0}
            _ = v21
        case 165:
            if (!((v20 & (v15 | (v15 % ((v20) | 1)))) != (v15 << ((1) & 63)))) {
                var v22 float64 = float64((1e10 - -2.25))
                switch p0 {
                case "":
                    var v23 int64 = 9223372036854775807
                    _ = v23
                case "hello":
                    _ = len(v14)
                case "foo bar":
                    v20 -= (v20 | v15)
                }
                _ = v22
            } else {
                var v24 uint8 = (p1 | (^(p1)))
                fmt.Println(v15)
                _ = v24
            }
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return -2147483648
}

func main() {
    var v1 uint = 939
    var v2 uint64 = 18446744073709551615
    v3 := []string{"\n\t", ("hello" + "foo bar")}
    v4 := make(map[int]string)
    v5 := []string{"a", "\n\t", (((("a" + "\n\t") + "\n\t") + "hello") + "hello")}
    var v25 uint32 = 996
    switch "foo bar" {
    case "foo bar":
        v1 = (^(v1))
    case "\n\t":
        _ = len(v3)
    case "a":
        var v26 uint64 = v2
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v1 += 18446744073709551615
        _ = fi27
    }
    var v28 float32 = 1.5
    if (1.5 < v28) {
        v2 = (v2 % ((2) | 1))
    } else {
        var v29 float64 = (1e10 - -2.25)
        v25 = (v25 | 2147483648)
        if (!(true || ((v29 >= 0.0) && (false && false)))) {
            switch 9223372036854775807 {
            case 2147483647:
                v2 += (v2 + (v2 * v2))
            case 127:
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            }
        }
        _ = v29
    }
    var v30 uint64 = v2
    _ = len(v3)
    if false {
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := GenS2{F0: "hello", F1: 2}
            _ = fi31
            _ = v32
        }
        switch "\n\t" {
        case "":
            switch "a" {
            case "hello":
                if false {
                    for fi33 := 0; fi33 < 20; fi33++ {
                        v25 -= v25
                        if ("hello" > "\n\t") {
                            v34 := v4[(fi33 >> ((v1) & 63))]
                            var v35 int = fi33
                            v4[fi33] = ""
                            _ = v34
                            _ = v35
                        } else {
                            v36 := GenS1{F0: 484}
                            _ = v36
                        }
                        if (!true) {
                            v1 = uint(v2)
                        }
                        _ = fi33
                    }
                    v1 = v1
                    v37 := GenS2{F0: "hello", F1: 62}
                    _ = v37
                } else {
                    if len(v5) > 0 {
                        v5[0] = ("\n\t" + "foo bar")
                    }
                    fmt.Println(v28, v28)
                    fmt.Println(v25)
                }
            default:
                var v38 bool = (0.1 > v28)
                _ = v38
            }
        default:
            var v39 int32 = 0
            _ = v39
        }
    } else {
        if true {
            switch ((("hello" + "foo bar") + "a") + "foo bar") {
            case "\n\t":
                v2 += v30
            case "foo bar":
                var v40 int = 2147483647
                _ = v40
            }
            fmt.Println(v1)
        }
        v30 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
