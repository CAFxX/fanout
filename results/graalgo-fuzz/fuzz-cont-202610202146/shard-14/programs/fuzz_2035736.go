package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int64; F1 int8; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 int32; F2 string }
type GenS6 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 int, p1 int64, p2 uint32) uint16 {
    var v4 int32 = 0
    p0 = p0
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(p1, fi5, v4)
        _ = fi5
    }
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}
func genFn1(p0 bool, p1 int64) uint64 {
    switch -2147483648 {
    case 42:
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := []int{-9223372036854775808, 4294967295, (fi6 * 0)}
            fmt.Println(p0, p1)
            p1 = p1
            _ = fi6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(p1)
            v9 := GenS1{F0: 65535, F1: 65535, F2: genFn0(fi8, (^(p1)), 42)}
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: 772, F1: 817, F2: 255}
    _ = p0
    _ = p1
    _ = v10
    return 325
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 float32 = float32(-2.25)
    v1 = v1
    if v1 {
        var v11 float64 = 0.0
        v12 := []int{42}
        var v13 int64 = -1
        _ = v11
        _ = v12
        _ = v13
    } else {
        if (!v1) {
            for fi14 := 0; fi14 < 1; fi14++ {
                v3 = 3.14159
                v1 = v1
                v1 = (843 >= 436)
                _ = fi14
            }
            var v15 int32 = -1
            if v1 {
                var v16 uint32 = 0
                v3 = (v3 - 3.14159)
                _ = v16
            }
            _ = v15
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                if ("hello" == ("foo bar" + "hello")) {
                    v18 := GenS2{F0: -821, F1: 255}
                    v1 = v1
                    _ = v18
                }
                _ = fi17
            }
            switch 2147483647 {
            case -1:
                delete(v2, "a")
            }
        }
    }
    if (65535 < 2) {
        var v19 float64 = 1.5
        _ = v19
    } else {
        var v20 int8 = 42
        _ = v20
    }
    v3 = (0.1 - ((-2.25 + (v3 - 1e10)) / 0.1))
    var v21 float64 = -2.25
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS2{F0: -47, F1: 255}
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 uint = 255
            v26 := GenS1{F0: (v25 + (v25 + (v25 & (v25 >> ((v25) & 63))))), F1: 42, F2: 2}
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    v3 += v3
    v27 := GenS3{F0: -834, F1: 0, F2: 42}
    fmt.Println(v1, v3, v21)
    switch ("" + "") {
    case "a":
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 float32 = v3
            var v30 uint16 = 42
            v31 := GenS4{F0: -810}
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    case "foo bar":
        if v1 {
            v1 = (!v1)
            if v1 {
                v3 -= (0.0 - v3)
            } else {
                switch "\n\t" {
                case "foo bar":
                    v3 += v3
                case "hello":
                    fmt.Println(v1)
                case "\n\t":
                    var v32 bool = false
                    _ = v32
                default:
                    switch "a" {
                    case "hello":
                        fmt.Println(v21, v3)
                    default:
                        var v33 uint = 255
                        _ = v33
                    }
                }
                switch 1 {
                case 1:
                    v3 = v3
                }
                for fi34 := 0; fi34 < 20; fi34++ {
                    v35 := GenS5{F0: 255, F1: -707, F2: ("foo bar" + "foo bar")}
                    var v36 uint64 = 2
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        } else {
            fmt.Println(v1, v1)
        }
    case "\n\t":
        v21 += float64(v3)
    default:
        v37 := GenS6{F0: 255, F1: 2147483647, F2: -1}
        _ = v37
    }
    var v38 int16 = 32767
    fmt.Println(v1, v21)
    for fi39 := 0; fi39 < 14; fi39++ {
        if (2147483647 > 2147483648) {
            fmt.Println(v38, v1, v3)
            fmt.Println(v21)
        } else {
            switch fi39 {
            case 826:
                v2[(("hello" + "\n\t") + (("foo bar" + "\n\t") + "a"))] = fi39
            case 1:
                for fi40 := 0; fi40 < 8; fi40++ {
                    if (2147483648 >= 0) {
                        v21 = v21
                        var v41 float32 = (3.14159 + v3)
                        _ = v41
                    }
                    if v1 {
                        v42 := []int{(fi39 - (fi39 - fi40)), (fi39 ^ fi40)}
                        _ = v42
                    }
                    _ = fi40
                }
            default:
                var v43 int64 = 419
                _ = v43
            }
        }
        _ = fi39
    }
    fmt.Println(v3, v1, v1)
    if ("" != "") {
        v38 += (v38 * (v38 << ((967) & 63)))
        if false {
            var v44 uint8 = 1
            fmt.Println(v1)
            var v45 uint64 = 2
            _ = v44
            _ = v45
        }
    }
    v46 := GenS5{F0: 62, F1: 0, F2: "a"}
    v47 := GenS5{F0: 42, F1: 2147483647, F2: "a"}
    fmt.Println(v21, v21, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v38
    _ = v46
    _ = v47
}
