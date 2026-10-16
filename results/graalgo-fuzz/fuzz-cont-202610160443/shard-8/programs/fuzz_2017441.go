package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int32; F2 int64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 bool; F1 string; F2 uint }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            if (0 <= 2) {
                var v8 int8 = 127
                _ = v8
            }
            fmt.Println(fi6, fi6)
            _ = fi7
        }
        var v9 uint = 42
        v10 := GenS1{F0: 42}
        _ = fi6
        _ = v10
        _ = v9
    }
    switch "foo bar" {
    case "":
        switch "\n\t" {
        case "":
            for fi11 := 0; fi11 < 10; fi11++ {
                switch 9223372036854775807 {
                case -1:
                    var v12 uint = 42
                    _ = v12
                case 643:
                    var v13 float64 = (3.14159 / 1.5)
                    _ = v13
                case -403:
                    v14 := GenS2{F0: ((-943 >= 0) || true), F1: 0, F2: -1}
                    _ = v14
                }
                var v15 int8 = -128
                v15 -= (v15 & (v15 ^ (v15 | 100)))
                _ = fi11
                _ = v15
            }
        case "\n\t":
            var v16 float32 = (-2.25 - 1e10)
            _ = v16
        default:
            var v17 int8 = 1
            _ = v17
        }
    case "a":
        var v18 int8 = 127
        _ = v18
    default:
        switch ("a" + ("hello" + ("" + (((("" + "foo bar") + "") + ("foo bar" + "hello")) + "\n\t")))) {
        case "a":
            var v19 float32 = 0.0
            _ = v19
        case "foo bar":
            var v20 int = 4294967295
            _ = v20
        case "hello":
            var v21 uint8 = 127
            _ = v21
        default:
            if false {
                if (!true) {
                    switch 1 {
                    case 128:
                        var v22 string = ""
                        _ = v22
                    }
                }
                if (2147483647 < -1) {
                    var v23 uint = 0
                    var v24 int8 = -128
                    _ = v23
                    _ = v24
                }
                var v25 bool = true
                _ = v25
            }
        }
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := []int{fi26, -2147483648}
        _ = fi26
        _ = v27
    }
    if (-1 >= 127) {
        var v28 int8 = 75
        var v29 int8 = v28
        v30 := GenS2{F0: true, F1: 0, F2: 0}
        _ = v28
        _ = v29
        _ = v30
    }
    return 238
}
func genFn1(p0 bool, p1 uint32, p2 uint8) int32 {
    for fi31 := 0; fi31 < 4; fi31++ {
        v32 := GenS1{F0: 609}
        var v33 uint = 0
        var v34 uint8 = 0
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 uint16 = 373
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    return -2147483648
}

func main() {
    v1 := []string{("foo bar" + "foo bar"), "", ("foo bar" + ""), "\n\t"}
    var v2 int8 = 55
    v3 := []uint32{0}
    var v4 uint8 = 138
    v5 := make(map[int]bool)
    v2 -= 42
    fmt.Println(v2)
    var v36 uint32 = 1
    var v37 float32 = (0.1 + float32((((-2.25 / ((1.5 - -2.25) / 1.5)) * (float64(0.1) * (1e10 * (1.5 / 3.14159)))) - 3.14159)))
    v38 := GenS1{F0: 265}
    for fi39 := 0; fi39 < 9; fi39++ {
        if false {
            v40 := GenS3{F0: 980, F1: (fi39 | 1)}
            if (!false) {
                for fi41 := 0; fi41 < 7; fi41++ {
                    delete(v5, 1)
                    if true {
                        if false {
                            var v42 uint16 = 32767
                            switch "a" {
                            case "hello":
                                var v43 int = 128
                                _ = v43
                            }
                            _ = v42
                        } else {
                            v2 -= v2
                            v37 += (((v37 + v37) - v37) / v37)
                        }
                        switch ("" + "hello") {
                        case "a":
                            fmt.Println(v37)
                        case "":
                            v2 += v2
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v36 += (v36 ^ (v36 | 42))
                        default:
                            v4 -= 42
                        }
                        fmt.Println(v36, fi39, v4)
                    }
                    var v44 uint8 = v4
                    _ = fi41
                    _ = v44
                }
                switch "hello" {
                case "":
                    v45 := GenS4{F0: true, F1: "hello", F2: 65535}
                    _ = v45
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    v1 = append(v1, "foo bar")
                }
                v36 = 0
            }
            _ = v40
        } else {
            _ = len(v5)
        }
        var v46 uint8 = (v4 + 2)
        _ = fi39
        _ = v46
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v36 = 391
        _ = fi47
    }
    v48 := GenS2{F0: false, F1: -2147483648, F2: -1}
    if false {
        v4 -= v4
        _ = len(v3)
    } else {
        var v49 uint8 = v4
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v48
    _ = v5
}
