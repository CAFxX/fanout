package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
func genFn0(p0 bool, p1 int16) uint {
    p0 = ((false || (65535 < 2)) || p0)
    if p0 {
        for fi4 := 0; fi4 < 12; fi4++ {
            for fi5 := 0; fi5 < 8; fi5++ {
                var v6 uint64 = 1
                v6 = 678
                var v7 uint8 = uint8(v6)
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
        var v8 uint8 = 83
        v9 := []int{-635, -992, 975, -9223372036854775808}
        _ = v8
        _ = v9
    } else {
        var v10 int16 = p1
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        switch "foo bar" {
        case "hello":
            if false {
                switch ("" + "hello") {
                case "a":
                    v12 := GenS1{F0: 793, F1: 101}
                    _ = v12
                case "\n\t":
                    if p0 {
                        var v13 uint64 = 18446744073709551615
                        _ = v13
                    } else {
                        switch ((("hello" + "\n\t") + "") + "\n\t") {
                        case "":
                            p0 = p0
                        default:
                            var v14 int32 = 0
                            _ = v14
                        }
                    }
                case "foo bar":
                    v15 := []int{128}
                    _ = v15
                }
            }
        case "a":
            v16 := []int{fi11}
            _ = v16
        case "":
            p1 -= p1
        default:
            p1 += (p1 ^ (p1 & p1))
        }
        _ = fi11
    }
    p1 -= p1
    v17 := []int{0, -1, 1}
    _ = p0
    _ = p1
    _ = v17
    return 0
}

func main() {
    v1 := []string{"", "\n\t", "foo bar"}
    var v2 uint8 = 2
    var v3 uint32 = 2
    for fi18 := 0; fi18 < 1; fi18++ {
        v2 = (^(v2))
        _ = fi18
    }
    v19 := GenS2{F0: (!true), F1: v3, F2: 128}
    switch "" {
    case "foo bar":
        var v20 bool = false
        _ = v20
    case "\n\t":
        v21 := GenS3{F0: uint32(v2), F1: "\n\t", F2: -32768}
        _ = v21
    case "hello":
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 float64 = 3.14159
            var v24 int = (fi22 / ((fi22) | 1))
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        var v25 int32 = -47
        _ = v25
    }
    v1 = append(v1, "foo bar")
    fmt.Println(v3, v3)
    var v26 int8 = 42
    v26 += (v26 * 103)
    fmt.Println(v2, v2)
    var v27 float64 = 0.1
    switch "" {
    case "hello":
        v28 := GenS2{F0: false, F1: (v3 * v3), F2: 53}
        _ = v28
    case "a":
        if (!(!true)) {
            if false {
                var v29 int = 191
                var v30 bool = true
                for fi31 := 0; fi31 < 6; fi31++ {
                    var v32 bool = (false && (v30 || ((((3.14159 / 1.0) > 0.0) || (v2 != 2)) || (v27 >= (v27 / v27)))))
                    switch "" {
                    case "":
                        for fi33 := 0; fi33 < 18; fi33++ {
                            switch ("" + "hello") {
                            case "foo bar":
                                v1 = append(v1, ("a" + "\n\t"))
                            case "a":
                                v2 -= v2
                            case "\n\t":
                                fmt.Println(v3, v32, fi31)
                            default:
                                v30 = true
                            }
                            v34 := GenS3{F0: v3, F1: "\n\t", F2: 193}
                            v3 += v3
                            switch ("a" + "a") {
                            case "":
                                v35 := GenS2{F0: v30, F1: (v3 + v3), F2: 756}
                                _ = v35
                            case "foo bar":
                                v2 = v2
                            }
                            _ = fi33
                            _ = v34
                        }
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            if (!true) {
                                v2 -= v2
                                var v36 float64 = ((((v27 - (v27 + float64(1.5))) * v27) - ((v27 - v27) * ((0.1 / v27) * 0.1))) + (1e10 / float64(0.1)))
                                _ = v36
                            } else {
                                if v32 {
                                    v3 += v3
                                    var v37 string = (("hello" + "a") + (("\n\t" + ("foo bar" + "\n\t")) + ""))
                                    _ = v37
                                }
                                v27 += 1.5
                                v27 -= v27
                            }
                        }
                    }
                    _ = fi31
                    _ = v32
                }
                _ = v29
                _ = v30
            } else {
                v3 = v3
            }
            fmt.Println(v2, v2)
        }
    }
    v27 -= -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
