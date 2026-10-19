package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 uint }
func genFn0(p0 int64, p1 uint8) string {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int16 = -209
        _ = fi7
        _ = v8
    }
    if true {
        v9 := GenS1{F0: p1, F1: 0}
        var v10 bool = (-1 >= 127)
        _ = v10
        _ = v9
    } else {
        p1 = p1
        p0 += p0
        for fi11 := 0; fi11 < 3; fi11++ {
            p0 -= (p0 / ((p0) | 1))
            if true {
                fmt.Println(p1)
            }
            switch "a" {
            case "":
                v12 := []int{fi11, fi11}
                _ = v12
            case "a":
                var v13 uint8 = (p1 - (^(p1)))
                _ = v13
            }
            _ = fi11
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    return "hello"
}
func genFn1(p0 int8, p1 uint, p2 uint32) bool {
    fmt.Println(p2)
    v14 := []int{-789, 42, 1, 42}
    switch -1 {
    case 9223372036854775807:
        if true {
            v15 := GenS2{F0: -9223372036854775808}
            _ = v15
        }
    case -9223372036854775808:
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS2{F0: fi16}
            _ = fi16
            _ = v17
        }
    case 2147483647:
        p1 -= p1
    }
    _ = len(v14)
    v18 := GenS1{F0: 1, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v18
    return ((428 >= 2147483647) && false)
}

func main() {
    var v1 int32 = -168
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := []string{"\n\t", "", ""}
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int32 = (v1 / (((^(v1))) | 1))
        if true {
            switch fi19 {
            case 9223372036854775807:
                for fi21 := 0; fi21 < 7; fi21++ {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(fi21)
                    default:
                        if true {
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                            var v22 string = ""
                            _ = v22
                        } else {
                            var v23 uint8 = 127
                            fmt.Println(v1, v20, fi19)
                            _ = v23
                        }
                    }
                    _ = len(v2)
                    v20 -= 0
                    _ = fi21
                }
            case 285:
                v1 = 166
            }
        } else {
            v1 += (^(v1))
            if ("hello" <= "") {
                var v24 int = (fi19 * fi19)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v24
            }
            v1 = v1
        }
        _ = fi19
        _ = v20
    }
    v1 = (v1 - v1)
    for fi25 := 0; fi25 < 11; fi25++ {
        v1 -= v1
        var v26 int = 537
        _ = fi25
        _ = v26
    }
    var v27 int32 = (-(v1))
    if (!true) {
        if false {
            v27 -= -2147483648
        } else {
            var v28 bool = (false || false)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v28
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "\n\t" {
        case "\n\t":
            v27 -= v27
        case "hello":
            switch "a" {
            case "\n\t":
                v1 = v27
            default:
                var v29 string = "\n\t"
                _ = v29
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ("" + ("\n\t" + ("foo bar" + "a"))) {
    case "":
        for fi30 := 0; fi30 < 2; fi30++ {
            v31 := GenS3{F0: 245, F1: 42}
            _ = fi30
            _ = v31
        }
    case "foo bar":
        var v32 uint32 = 0
        _ = v32
    }
    _ = len(v5)
    var v33 int16 = 1000
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := GenS4{F0: -1, F1: -1, F2: 147}
        v27 = (^(v1))
        fmt.Println(v1)
        _ = fi34
        _ = v35
    }
    if (!false) {
        for fi36 := 0; fi36 < 18; fi36++ {
            v4 = append(v4, "")
            _ = fi36
        }
        v3["\n\t"] = true
    }
    v37 := GenS2{F0: 1}
    switch (("a" + "hello") + "foo bar") {
    case "foo bar":
        fmt.Println(v33)
    case "":
        v38 := GenS1{F0: 42, F1: 255}
        _ = v38
    default:
        v27 = v1
    }
    var v39 float64 = 1.5
    v2[42] = (v39 >= v39)
    for fi40 := 0; fi40 < 19; fi40++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v39 += v39
        _ = fi40
    }
    v41 := GenS4{F0: -1, F1: -89, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v37
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
