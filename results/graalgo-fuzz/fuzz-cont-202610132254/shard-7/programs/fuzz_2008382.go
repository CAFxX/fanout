package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 string; F1 int32; F2 uint64 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 0}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            fmt.Println(fi9)
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            switch "" {
            case "a":
                var v11 uint16 = 32767
                _ = v11
            case "":
                v12 := []int{fi10, 127, (fi10 & -163)}
                _ = v12
            default:
                v13 := GenS2{F0: 1}
                _ = v13
            }
            var v14 int32 = 42
            _ = fi10
            _ = v14
        }
        _ = fi9
    }
    return 1000
}
func genFn1(p0 int32, p1 uint16, p2 uint) uint {
    fmt.Println(p0)
    if false {
        fmt.Println(p0, p2, p0)
        var v15 uint16 = p1
        _ = v15
    } else {
        v16 := GenS1{F0: 4294967295}
        var v17 int8 = 127
        v17 = (-(v17))
        _ = v16
        _ = v17
    }
    var v18 int32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (^(p2))
}

func main() {
    v1 := []uint32{4294967295}
    var v2 float32 = 0.0
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    var v5 int = -802
    var v6 float64 = (0.0 + 0.1)
    var v19 int8 = -60
    fmt.Println(v6, v2)
    _ = len(v4)
    switch "hello" {
    case "\n\t":
        for fi20 := 0; fi20 < 18; fi20++ {
            switch "foo bar" {
            case "a":
                v21 := GenS3{F0: "hello", F1: "a"}
                _ = v21
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 42
                }
            case "foo bar":
                fmt.Println(v6)
            default:
                v22 := GenS2{F0: 42}
                _ = v22
            }
            _ = fi20
        }
    default:
        v19 = -128
    }
    var v23 int64 = 869
    for fi24 := 0; fi24 < 1; fi24++ {
        v6 += v6
        _ = fi24
    }
    fmt.Println(v6, v5, v23)
    v2 = v2
    switch v5 {
    case 32768:
        if true {
            fmt.Println(v23, v5)
            v19 = (v19 & 0)
        }
    case 568:
        fmt.Println(v19, v6)
    case -9223372036854775808:
        if true {
            switch ("a" + "hello") {
            case "hello":
                v23 -= (v23 >> ((334) & 63))
            case "":
                for fi25 := 0; fi25 < 18; fi25++ {
                    var v26 float64 = 1.5
                    _ = fi25
                    _ = v26
                }
            default:
                v23 += -9223372036854775808
            }
            for fi27 := 0; fi27 < 14; fi27++ {
                switch "hello" {
                case "\n\t":
                    if true {
                        var v28 uint8 = 0
                        _ = v28
                    }
                case "a":
                    if true {
                        v29 := GenS4{F0: 1, F1: -1}
                        var v30 uint8 = 255
                        _ = v29
                        _ = v30
                    }
                case "hello":
                    var v31 int32 = 2147483647
                    _ = v31
                }
                switch "hello" {
                case "a":
                    v23 -= 9223372036854775807
                case "hello":
                    v32 := GenS3{F0: ("foo bar" + "foo bar"), F1: "hello"}
                    _ = v32
                default:
                    fmt.Println(v2, v5)
                }
                var v33 uint = 42
                _ = fi27
                _ = v33
            }
            v34 := GenS3{F0: "hello", F1: "hello"}
            _ = v34
        }
    default:
        var v35 uint32 = 2
        _ = v35
    }
    if false {
        var v36 bool = true
        v37 := GenS5{F0: v5, F1: genFn1(2147483647, 255, 396)}
        _ = v36
        _ = v37
    } else {
        v38 := GenS6{F0: "\n\t", F1: 1, F2: 440}
        for fi39 := 0; fi39 < 8; fi39++ {
            v19 += v19
            var v40 uint8 = 18
            if false {
                var v41 string = "\n\t"
                for fi42 := 0; fi42 < 5; fi42++ {
                    v43 := GenS3{F0: "hello", F1: (v41 + "foo bar")}
                    _ = fi42
                    _ = v43
                }
                _ = v41
            }
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
