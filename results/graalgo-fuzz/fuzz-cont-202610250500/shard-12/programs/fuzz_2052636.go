package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 bool }
func genFn0(p0 bool) string {
    switch "\n\t" {
    case "":
        if p0 {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: 607, F1: 4294967295}
                if false {
                    v9 := GenS1{F0: 4294967295, F1: 2147483647}
                    for fi10 := 0; fi10 < 7; fi10++ {
                        p0 = p0
                        p0 = false
                        v11 := []int{}
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                } else {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        fmt.Println(fi12, fi7)
                        p0 = p0
                        if p0 {
                            if ((float32((0.1 / 3.14159)) / ((0.1 - 1.5) / 3.14159)) < 3.14159) {
                                p0 = (true && (0 <= 1))
                            }
                        } else {
                            v13 := []int{(fi12 + fi7), 135}
                            _ = v13
                        }
                        _ = fi12
                    }
                    p0 = (-1 >= -2147483648)
                    p0 = p0
                }
                _ = fi7
                _ = v8
            }
            p0 = (42 != 510)
            p0 = true
        }
    default:
        v14 := GenS2{F0: "\n\t", F1: -519, F2: 42}
        _ = v14
    }
    switch 4294967295 {
    case -522:
        var v15 float32 = (0.1 + (0.0 - float32((float64(3.14159) * 0.0))))
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: 310}
    if p0 {
        if p0 {
            p0 = (p0 && (p0 || (1e10 != 3.14159)))
            var v17 string = ""
            v18 := []int{-459, -143, -77, 42}
            _ = v17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS3{F0: "foo bar"}
            _ = fi19
            _ = v20
        }
        var v21 uint16 = 431
        _ = v21
    }
    p0 = (2 != 2)
    _ = p0
    _ = v16
    return ("a" + "a")
}
func genFn1(p0 int32) string {
    var v22 int8 = 42
    v23 := GenS4{F0: 623}
    var v24 int16 = 0
    v24 = (^(v24))
    var v25 uint = 42
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    return "a"
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("foo bar" + "hello")
    var v3 float64 = 1.5
    v4 := []string{v2, (v2 + ""), v2, (v2 + (v2 + "foo bar"))}
    v5 := make(map[uint32]uint32)
    var v6 float32 = ((0.1 + 0.0) / (3.14159 / 1.0))
    v26 := GenS4{F0: 65535}
    var v27 uint = 65535
    for fi28 := 0; fi28 < 9; fi28++ {
        switch (fi28 << (((v27 + (v27 >> ((v27) & 63)))) & 63)) {
        case 42:
            v29 := GenS3{F0: v2}
            _ = v29
        case 32767:
            v30 := GenS2{F0: (v2 + "hello"), F1: 42, F2: 100}
            _ = v30
        default:
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 uint = (v27 ^ (v27 & v27))
                _ = fi31
                _ = v32
            }
        }
        _ = fi28
    }
    var v33 int16 = -155
    delete(v1, v2)
    switch v2 {
    case "\n\t":
        for fi34 := 0; fi34 < 4; fi34++ {
            v35 := GenS3{F0: v2}
            _ = fi34
            _ = v35
        }
    case "a":
        var v36 uint16 = 680
        _ = v36
    }
    v37 := GenS5{F0: 297}
    v38 := GenS6{F0: 0, F1: (-128 > -24)}
    for fi39 := 0; fi39 < 14; fi39++ {
        if false {
            v33 += v33
            v27 += v27
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            fmt.Println(v33, v3)
            var v41 int = (fi39 | fi39)
            fmt.Println(v6, v2, v2)
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    if true {
        var v42 uint16 = 2
        v2 = (v2 + "foo bar")
        _ = v42
    } else {
        v43 := GenS5{F0: 65535}
        _ = v43
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        var v45 uint32 = 2147483647
        v3 += (float64((3.14159 / float32(1.5))) * (-2.25 - v3))
        _ = fi44
        _ = v45
    }
    var v46 float64 = v3
    v46 = (3.14159 + 0.1)
    v2 = v2
    if (-9223372036854775808 <= 42) {
        v4 = append(v4, "a")
        fmt.Println(v27, v2, v3)
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        var v48 string = ("" + "hello")
        fmt.Println(v2, fi47, v33)
        if len(v4) > 0 {
            v4[0] = ((("foo bar" + (((v48 + v48) + v2) + ((v2 + (((v48 + (v2 + v2)) + v2) + "a")) + v48))) + (((v2 + "foo bar") + "foo bar") + (("" + "hello") + v2))) + "")
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
