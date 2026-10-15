package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int8; F1 uint; F2 uint16 }
type GenS4 struct { F0 int8; F1 string; F2 uint64 }
type GenS5 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 int8, p1 int64, p2 bool) uint8 {
    v4 := []int{2147483647}
    if (2 > 78) {
        switch "\n\t" {
        case "foo bar":
            v5 := GenS1{F0: p1, F1: 2147483647, F2: 81}
            _ = v5
        case "hello":
            var v6 float32 = 0.1
            _ = v6
        default:
            var v7 float64 = float64((-2.25 * 1.5))
            _ = v7
        }
        if len(v4) > 0 {
            v4[0] = 42
        }
        if (true && p2) {
            var v8 uint = 255
            var v9 int = -532
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS1{F0: p1, F1: 1, F2: 42}
                _ = fi10
                _ = v11
            }
            _ = v8
            _ = v9
        } else {
            if (42 > 976) {
                if (1 >= 42) {
                    var v12 int16 = 989
                    v12 = v12
                    v12 = v12
                    _ = v12
                }
                v13 := GenS1{F0: -120, F1: -830, F2: 2}
                _ = v13
            } else {
                var v14 uint16 = 360
                _ = v14
            }
            p2 = p2
            if true {
                if (1.5 >= 1e10) {
                    v15 := GenS1{F0: p1, F1: -1, F2: 0}
                    _ = v15
                }
                p0 = p0
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v16 := GenS1{F0: p1, F1: 1, F2: 0}
                _ = v16
            }
        }
    }
    var v17 uint8 = 30
    fmt.Println(v17)
    var v18 uint64 = 57
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v4
    return (v17 / ((uint8(v18)) | 1))
}
func genFn1(p0 uint64, p1 uint, p2 bool) uint8 {
    fmt.Println(p1, p0, p2)
    fmt.Println(p0, p2)
    var v19 uint64 = p0
    v20 := []int{-414, -1, -1}
    if p2 {
        for fi21 := 0; fi21 < 11; fi21++ {
            p1 += uint(p0)
            _ = fi21
        }
        if p2 {
            if p2 {
                v22 := GenS2{F0: "hello", F1: 127}
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := GenS1{F0: 0, F1: 42, F2: 2}
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = fi23
                    _ = v24
                }
                var v25 int = 32767
                _ = v22
                _ = v25
            } else {
                p2 = (p2 && p2)
            }
        } else {
            var v26 uint8 = 0
            _ = v26
        }
        _ = len(v20)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return uint8(p0)
}

func main() {
    v1 := []int{}
    v2 := []string{"\n\t", ""}
    v3 := make(map[string]uint32)
    switch ("foo bar" + (("foo bar" + "\n\t") + ("" + ""))) {
    case "foo bar":
        switch "a" {
        case "hello":
            for fi27 := 0; fi27 < 5; fi27++ {
                fmt.Println(fi27)
                var v28 int16 = -1
                fmt.Println(v28, fi27)
                _ = fi27
                _ = v28
            }
        case "foo bar":
            v3[""] = 1
        }
    case "a":
        var v29 int8 = 127
        _ = v29
    case "\n\t":
        v3["\n\t"] = 2
    default:
        _ = len(v1)
    }
    var v30 uint = 18446744073709551615
    v30 = 1
    fmt.Println(v30)
    v30 -= (^(v30))
    var v31 string = "\n\t"
    var v32 float64 = (0.0 * (-2.25 + (float64((3.14159 - (((0.0 * 1e10) * (((0.0 / 1.0) * float32(1.5)) - 1e10)) + (1e10 * 1e10)))) * 1e10)))
    if ((0.1 + 1.5) < (-2.25 - 0.0)) {
        if false {
            v33 := v3[("foo bar" + "a")]
            for fi34 := 0; fi34 < 20; fi34++ {
                switch "a" {
                case "\n\t":
                    v35 := GenS3{F0: 100, F1: v30, F2: 255}
                    _ = v35
                case "hello":
                    if (-70 != -32768) {
                        var v36 float32 = float32(v32)
                        _ = len(v2)
                        var v37 uint16 = 65535
                        _ = v36
                        _ = v37
                    } else {
                        var v38 uint16 = 42
                        if len(v2) > 0 {
                            v2[0] = (("" + v31) + v31)
                        }
                        _ = v38
                    }
                case "":
                    delete(v3, "a")
                default:
                    fmt.Println(v31, v30, v31)
                }
                _ = fi34
            }
            var v39 int32 = -1
            _ = v33
            _ = v39
        }
        v32 += (((1.5 * -2.25) * v32) * ((v32 * 0.0) / 3.14159))
        v40 := GenS4{F0: 42, F1: v31, F2: 0}
        _ = v40
    } else {
        for fi41 := 0; fi41 < 2; fi41++ {
            fmt.Println(v32, v32)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi41
        }
        fmt.Println(v31, v31)
        switch 0 {
        case 602:
            for fi42 := 0; fi42 < 1; fi42++ {
                v32 += (v32 * v32)
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                _ = fi42
            }
        case -838:
            _ = len(v3)
        case 613:
            var v43 float32 = ((float32(0.0) + 1e10) - 1e10)
            _ = v43
        }
    }
    v44 := GenS5{F0: 32767, F1: "hello", F2: 127}
    for fi45 := 0; fi45 < 15; fi45++ {
        if (42 == 645) {
            v32 = (v32 / 0.1)
            fmt.Println(v30, v32, fi45)
            fmt.Println(fi45, fi45, v32)
        }
        switch 0 {
        case -509:
            var v46 int8 = 0
            _ = v46
        case 1:
            var v47 string = ((v31 + "") + v31)
            _ = v47
        default:
            for fi48 := 0; fi48 < 15; fi48++ {
                var v49 string = "\n\t"
                _ = fi48
                _ = v49
            }
        }
        _ = fi45
    }
    v30 = v30
    v30 = 255
    var v50 uint = v30
    var v51 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v44
    _ = v50
    _ = v51
}
