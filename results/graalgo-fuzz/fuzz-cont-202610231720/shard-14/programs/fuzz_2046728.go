package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 bool }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint16, p2 int) uint64 {
    switch -10 {
    case -2147483648:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(p1)
            fmt.Println(fi6, fi6)
            _ = fi6
        }
    case 1:
        v7 := GenS1{F0: 0}
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        p2 = (p2 - -1)
        _ = fi8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        p1 -= 65535
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 bool, p1 bool) int8 {
    fmt.Println(p0, p1)
    p1 = (-2.25 != 1.5)
    fmt.Println(p1)
    if false {
        var v10 int8 = -118
        p1 = (0 >= 2147483647)
        var v11 uint = 293
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 int8 = 1
    v5 := make(map[string]string)
    v12 := GenS1{F0: v4}
    switch ("hello" + "") {
    case "a":
        if (((128 == 45) && (!false)) && (!(false || false))) {
            v1 = (v1 ^ (v1 >> ((255) & 63)))
            v13 := GenS2{F0: 1, F1: 128}
            _ = v13
        }
    case "":
        if false {
            v14 := []int{int(v4), 1, 597, 4294967295}
            switch 1 {
            case 127:
                var v15 string = "\n\t"
                _ = v15
            default:
                v16 := GenS1{F0: (^(v4))}
                _ = v16
            }
            for fi17 := 0; fi17 < 16; fi17++ {
                switch 127 {
                case -455:
                    v5["\n\t"] = "hello"
                }
                switch -108 {
                case -1:
                    delete(v3, 2)
                case 9223372036854775807:
                    fmt.Println(v1, fi17, fi17)
                default:
                    v4 += v4
                }
                v18 := v5[("foo bar" + "a")]
                _ = fi17
                _ = v18
            }
            _ = v14
        } else {
            fmt.Println(v1, v4)
            switch 0 {
            case -667:
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := []int{(fi19 ^ (fi19 / (((fi19 % ((128) | 1))) | 1)))}
                    _ = len(v5)
                    _ = fi19
                    _ = v20
                }
            }
        }
    case "foo bar":
        var v21 float32 = float32(0.1)
        _ = v21
    default:
        var v22 int64 = 470
        _ = v22
    }
    if ((0 < 1) || false) {
        var v23 float64 = (((1e10 + -2.25) - 1.5) * (float64(1e10) * float64(1e10)))
        var v24 int = 32767
        _ = v23
        _ = v24
    } else {
        var v25 int8 = v4
        _ = v25
    }
    if false {
        v26 := []int{2147483647}
        fmt.Println(v1, v4)
        _ = v26
    }
    var v27 int8 = (v4 * (v4 >> ((18446744073709551615) & 63)))
    var v28 int64 = 1
    var v29 uint = 819
    if (((!((42 > 1) || true)) && (v27 >= (v27 >> (((v29 + 2)) & 63)))) || true) {
        var v30 int32 = 2147483647
        if true {
            var v31 float64 = -2.25
            v32 := GenS3{F0: 235, F1: 128, F2: 1}
            _ = v31
            _ = v32
        }
        _ = v30
    } else {
        switch "hello" {
        case "foo bar":
            v27 -= (v27 - genFn1((((v29 ^ 18446744073709551615) != (v29 / ((v29) | 1))) || (!false)), (true || ((v28 >> (((v29 * (v29 / ((2) | 1)))) & 63)) != -224))))
        case "a":
            v29 -= 0
        case "hello":
            for fi33 := 0; fi33 < 8; fi33++ {
                v34 := GenS4{F0: 0}
                _ = fi33
                _ = v34
            }
        default:
            var v35 int16 = -5
            _ = v35
        }
    }
    if false {
        v36 := GenS4{F0: -32768}
        v4 = (v4 + (-(v4)))
        _ = v36
    } else {
        v37 := GenS5{F0: uint32(v29), F1: uint32(v29), F2: false}
        _ = v37
    }
    v1 = v1
    var v38 int32 = 0
    v1 = genFn0(2, 988, int(v38))
    v39 := []int{42, 0}
    for fi40 := 0; fi40 < 13; fi40++ {
        v5[("a" + "foo bar")] = ("\n\t" + "\n\t")
        switch fi40 {
        case -9223372036854775808:
            v41 := GenS3{F0: 895, F1: 255, F2: -829}
            _ = v41
        case 4294967295:
            v5[(((("hello" + "\n\t") + "a") + ("foo bar" + ("" + ("" + "a")))) + "a")] = ""
        case 2147483647:
            switch ("a" + "foo bar") {
            case "":
                v42 := GenS6{F0: 1}
                _ = v42
            case "foo bar":
                v43 := GenS4{F0: 0}
                _ = v43
            }
        default:
            _ = len(v39)
        }
        if (uint32(v29) != 200) {
            var v44 bool = false
            fmt.Println(v29)
            _ = v44
        }
        _ = fi40
    }
    fmt.Println(v1, v27)
    var v45 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v5
}
