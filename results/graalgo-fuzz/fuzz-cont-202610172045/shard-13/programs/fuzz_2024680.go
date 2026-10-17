package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32; F1 int64; F2 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 string; F1 int }
func genFn0(p0 uint8, p1 uint64, p2 uint32) uint8 {
    v4 := GenS1{F0: 145, F1: 1}
    var v5 float64 = float64(-2.25)
    p2 -= p2
    fmt.Println(p0, p1)
    p1 = (p1 * (p1 >> ((791) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (^(p0))
}
func genFn1(p0 bool) int {
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(fi6, fi6)
        p0 = p0
        _ = fi6
    }
    p0 = (!(2 >= 4294967295))
    p0 = (false && false)
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(p0, fi7)
        _ = fi7
    }
    p0 = p0
    _ = p0
    return -309
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-9223372036854775808, 127}
    v3 := make(map[string]string)
    var v8 float64 = (-2.25 + 0.1)
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 string = ""
        if true {
            switch 32768 {
            case 9223372036854775807:
                var v11 uint32 = 42
                _ = v11
            }
            _ = len(v2)
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 uint32 = 2147483647
                _ = fi12
                _ = v13
            }
        } else {
            fmt.Println(v10, v8, fi9)
        }
        if (true && ((true && false) || ((v10 + v10) > ("a" + "foo bar")))) {
            switch fi9 {
            case 9223372036854775807:
                var v14 int = (fi9 << ((42) & 63))
                _ = v14
            case 178:
                switch fi9 {
                case 127:
                    if (false && ((353 == 0) && false)) {
                        switch v10 {
                        case "foo bar":
                            var v15 bool = (1 < 127)
                            _ = v15
                        default:
                            switch v10 {
                            case "hello":
                                var v16 uint8 = 31
                                _ = v16
                            default:
                                for fi17 := 0; fi17 < 16; fi17++ {
                                    v8 += (1.5 - 1.5)
                                    _ = fi17
                                }
                            }
                        }
                        switch ("a" + (v10 + v10)) {
                        case "a":
                            fmt.Println(v8, v8)
                        }
                    } else {
                        var v18 int16 = 0
                        if false {
                            if false {
                                fmt.Println(v18, v8)
                                var v19 int16 = (^(v18))
                                var v20 float64 = v8
                                _ = v19
                                _ = v20
                            }
                            var v21 uint32 = 4294967295
                            _ = v21
                        }
                        _ = v18
                    }
                case 954:
                    switch "foo bar" {
                    case "foo bar":
                        var v22 int16 = 32767
                        _ = v22
                    case "\n\t":
                        switch 32768 {
                        case 17:
                            var v23 float32 = float32(v8)
                            _ = v23
                        case 684:
                            v8 += v8
                        }
                    case "hello":
                        switch ((("a" + v10) + (v10 + "")) + (v10 + "")) {
                        case "":
                            var v24 int8 = -1
                            _ = v24
                        case "foo bar":
                            v10 = v10
                        case "hello":
                            var v25 uint32 = 291
                            _ = v25
                        }
                    default:
                        for fi26 := 0; fi26 < 16; fi26++ {
                            switch v10 {
                            case "":
                                var v27 float32 = 3.14159
                                v8 = 0.1
                                v28 := GenS2{F0: -987, F1: 42, F2: 178}
                                _ = v27
                                _ = v28
                            case "\n\t":
                                if (473 == 975) {
                                    var v29 int32 = -2147483648
                                    v30 := GenS3{F0: -125, F1: 2}
                                    _ = v29
                                    _ = v30
                                }
                            case "hello":
                                v10 = (((v10 + (v10 + v10)) + v10) + ("\n\t" + ("hello" + "\n\t")))
                            }
                            switch fi26 {
                            case 307:
                                var v31 int64 = -1
                                _ = v31
                            case 128:
                                v10 = ((v10 + "foo bar") + (v10 + ((("\n\t" + "a") + ("a" + (v10 + "hello"))) + v10)))
                            default:
                                v32 := GenS4{F0: v10, F1: 108}
                                _ = v32
                            }
                            _ = fi26
                        }
                    }
                case 4294967295:
                    v8 += 3.14159
                }
            default:
                v8 = float64(0.1)
            }
            v10 = (v10 + v10)
            v10 = "\n\t"
        } else {
            fmt.Println(fi9)
            var v33 float64 = v8
            _ = v33
        }
        _ = fi9
        _ = v10
    }
    var v34 int16 = 0
    v8 -= v8
    v35 := v1["a"]
    v8 += 1e10
    fmt.Println(v8, v35, v8)
    v36 := GenS2{F0: 2147483647, F1: 9223372036854775807, F2: 2}
    var v37 string = "foo bar"
    switch 753 {
    case 127:
        v35 -= (^(v35))
    case -71:
        fmt.Println(v8, v35)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v8
}
