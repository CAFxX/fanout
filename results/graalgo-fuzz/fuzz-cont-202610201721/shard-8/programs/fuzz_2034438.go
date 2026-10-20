package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0() int8 {
    if true {
        var v6 int = 599
        v7 := []int{32767, (v6 * 9223372036854775807), (v6 | (v6 * -2147483648)), v6}
        _ = v6
        _ = v7
    } else {
        var v8 bool = true
        _ = v8
    }
    var v9 int16 = -32768
    fmt.Println(v9)
    for fi10 := 0; fi10 < 16; fi10++ {
        v9 -= (v9 << ((4294967295) & 63))
        _ = fi10
    }
    v9 -= v9
    _ = v9
    return -128
}
func genFn1(p0 int8, p1 string) uint32 {
    v11 := []int{-2147483648, 42, -411}
    _ = len(v11)
    if (42 < 924) {
        var v12 uint64 = 18446744073709551615
        fmt.Println(v12)
        _ = v12
    }
    p0 = (p0 - 1)
    var v13 float64 = 1e10
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 bool = (-507 <= -32768)
    v3 := make(map[uint32]uint32)
    v4 := []string{"\n\t", ("" + "\n\t"), ((("hello" + "") + "foo bar") + "foo bar"), ("\n\t" + "hello")}
    var v5 int32 = 1
    if true {
        var v14 uint = 1
        var v15 int64 = 0
        _ = v14
        _ = v15
    }
    var v16 uint32 = 42
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: 734, F1: ("foo bar" + (("foo bar" + (("" + ("a" + "\n\t")) + ("" + "\n\t"))) + "foo bar"))}
        _ = fi17
        _ = v18
    }
    var v19 int8 = genFn0()
    switch 44 {
    case 1:
        v5 = (v5 ^ int32(v16))
    }
    var v20 uint = 1
    for fi21 := 0; fi21 < 13; fi21++ {
        switch ("" + (("\n\t" + "hello") + "hello")) {
        case "hello":
            v4 = append(v4, "foo bar")
        case "\n\t":
            var v22 uint = (v20 / ((2) | 1))
            _ = v22
        default:
            var v23 float64 = (3.14159 + (float64((1.5 * 3.14159)) * 0.0))
            _ = v23
        }
        v19 = genFn0()
        _ = fi21
    }
    if (950 > 717) {
        v16 = v16
    }
    switch "foo bar" {
    case "\n\t":
        switch 129 {
        case -548:
            switch -656 {
            case 42:
                var v24 float32 = (0.0 * -2.25)
                _ = v24
            case 390:
                v2 = v2
            case 32767:
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            }
        case 128:
            fmt.Println(v20, v19)
        default:
            switch 0 {
            case 32767:
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            case -970:
                fmt.Println(v20, v5, v2)
            }
        }
    case "hello":
        for fi25 := 0; fi25 < 5; fi25++ {
            v5 -= (v5 | v5)
            _ = fi25
        }
    default:
        delete(v3, (v16 * (^(v16))))
    }
    v26 := GenS2{F0: -139, F1: 4294967295}
    for fi27 := 0; fi27 < 7; fi27++ {
        switch "foo bar" {
        case "foo bar":
            var v28 int32 = 1
            _ = v28
        case "\n\t":
            if (-275 != 42) {
                v16 -= (^(v16))
                var v29 int8 = (v19 | (v19 / ((1) | 1)))
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS1{F0: 0, F1: ((("hello" + "foo bar") + "") + ("\n\t" + (("a" + "hello") + ("hello" + "\n\t"))))}
                    v32 := GenS3{F0: 2, F1: v16, F2: 9223372036854775807}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = v29
            } else {
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 uint = 387
                    v3[(v16 % ((v16) | 1))] = genFn1((v19 % (((^(v19))) | 1)), ("\n\t" + "\n\t"))
                    var v35 float32 = 1e10
                    var v36 string = "a"
                    if v2 {
                        v37 := GenS2{F0: 477, F1: (fi27 % (((fi33 / (((fi27 - fi33)) | 1))) | 1))}
                        var v38 int16 = -1
                        _ = v37
                        _ = v38
                    } else {
                        fmt.Println(v20, v20, v36)
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
                v39 := GenS4{F0: (v16 & 0), F1: "\n\t", F2: 1}
                _ = v39
            }
        case "a":
            switch "hello" {
            case "hello":
                var v40 float32 = 0.1
                _ = v40
            case "\n\t":
                v2 = v2
            default:
                fmt.Println(v16, v19, v2)
            }
        default:
            v41 := GenS3{F0: 2, F1: 2, F2: 9223372036854775807}
            _ = v41
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
