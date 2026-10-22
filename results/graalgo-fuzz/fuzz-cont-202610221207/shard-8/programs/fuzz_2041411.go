package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 int64, p2 int64) int16 {
    v7 := GenS1{F0: 2, F1: 42, F2: 951}
    fmt.Println(p2)
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS2{F0: -619, F1: 387, F2: 65535}
            p2 -= p2
            _ = fi9
            _ = v10
        }
        fmt.Println(fi8, p2, p1)
        p1 -= p1
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}
func genFn1() uint64 {
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := []int{fi11, 4294967295}
        _ = fi11
        _ = v12
    }
    var v13 uint64 = 42
    var v14 int32 = 216
    _ = v13
    _ = v14
    return 620
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar"}
    var v3 int32 = 2147483647
    v4 := []int{559}
    v5 := []uint32{2147483648}
    var v6 int8 = 100
    switch 4294967295 {
    case 32768:
        v3 = (v3 + 459)
    default:
        v3 -= v3
    }
    var v15 int16 = 219
    if (((0.1 / ((3.14159 * float32(1e10)) / 0.1)) + (0.0 * ((0.1 * (1e10 - 3.14159)) - (1e10 + -2.25)))) >= (-2.25 * -2.25)) {
        _ = len(v1)
    } else {
        switch 42 {
        case 42:
            fmt.Println(v6, v3)
        default:
            var v16 int8 = v6
            _ = v16
        }
        delete(v1, (("\n\t" + "\n\t") + ""))
        delete(v1, "\n\t")
    }
    var v17 uint8 = 190
    v3 += (v3 << ((65535) & 63))
    v18 := GenS1{F0: 1, F1: (v17 % ((v17) | 1)), F2: 42}
    fmt.Println(v15, v3, v3)
    for fi19 := 0; fi19 < 7; fi19++ {
        v15 = (v15 * (^(v15)))
        for fi20 := 0; fi20 < 1; fi20++ {
            v15 = -507
            _ = fi20
        }
        _ = fi19
    }
    switch ("" + ("hello" + ((("hello" + "hello") + "hello") + ""))) {
    case "hello":
        v17 -= 0
    case "a":
        for fi21 := 0; fi21 < 13; fi21++ {
            switch ("a" + ("" + (("a" + "") + (("a" + "") + ("foo bar" + ("a" + (("hello" + ("hello" + "hello")) + "a"))))))) {
            case "\n\t":
                v17 -= (v17 & 128)
            }
            v6 = 74
            _ = fi21
        }
    }
    v17 = v17
    v6 += v6
    for fi22 := 0; fi22 < 4; fi22++ {
        v15 = 42
        v23 := GenS1{F0: 0, F1: 255, F2: (v15 >> ((4294967295) & 63))}
        if (!(false && (((^(v15)) >= (-(v15))) || (0.0 == (-2.25 / 0.1))))) {
            v24 := GenS3{F0: (false && true), F1: 0}
            var v25 int64 = 1
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 19; fi26++ {
                if false {
                    v6 = (v6 & (v6 & (-(v6))))
                    switch "hello" {
                    case "":
                        if false {
                            var v27 float32 = 0.1
                            var v28 uint = 42
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int8 = (v6 % (((v6 ^ v6)) | 1))
                            _ = v29
                        }
                    case "foo bar":
                        switch ("" + "a") {
                        case "a":
                            v15 = v15
                        case "foo bar":
                            v15 += v15
                        case "":
                            if false {
                                v17 = (^(v17))
                            } else {
                                delete(v1, (((((("hello" + "hello") + "\n\t") + "a") + "") + ("a" + ((("foo bar" + "") + "") + "a"))) + "hello"))
                                v30 := GenS2{F0: -2147483648, F1: 0, F2: 65535}
                                _ = v30
                            }
                        }
                        var v31 float64 = float64(0.1)
                        var v32 uint32 = 141
                        _ = v31
                        _ = v32
                    case "\n\t":
                        if true {
                            v33 := GenS4{F0: 255}
                            _ = v33
                        }
                    default:
                        var v34 uint32 = 4294967295
                        _ = v34
                    }
                    if true {
                        var v35 uint8 = (v17 * (v17 * v17))
                        var v36 int64 = -589
                        fmt.Println(v6, v3)
                        _ = v35
                        _ = v36
                    } else {
                        var v37 uint32 = 222
                        _ = v37
                    }
                }
                switch 752 {
                case 32768:
                    v15 = int16(fi26)
                default:
                    v3 -= 42
                }
                _ = fi26
            }
            for fi38 := 0; fi38 < 8; fi38++ {
                var v39 int8 = v6
                fmt.Println(v17, fi22, v6)
                _ = fi38
                _ = v39
            }
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
