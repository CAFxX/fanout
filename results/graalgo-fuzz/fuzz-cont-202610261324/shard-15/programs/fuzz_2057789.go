package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0() int8 {
    v6 := GenS1{F0: 32767, F1: ""}
    var v7 int = -505
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 uint64) int16 {
    switch 2147483647 {
    case 4294967295:
        v8 := GenS1{F0: 349, F1: "a"}
        _ = v8
    default:
        p1 -= p1
    }
    v9 := []int{2147483647, 0, 42}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int = fi10
        p0 = p1
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return -32768
}

func main() {
    var v1 uint = 1
    v2 := []uint32{655, 2147483647, 65535, 2}
    v3 := []uint32{65535, 4294967295}
    var v4 uint16 = 0
    var v5 int16 = 693
    var v12 uint32 = 65535
    if (716 <= 2147483647) {
        v13 := GenS2{F0: 127}
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 bool = true
            if (0 != 1) {
                fmt.Println(v4)
            } else {
                v16 := GenS2{F0: 21}
                _ = v16
            }
            if true {
                switch fi14 {
                case 508:
                    if v15 {
                        switch ("foo bar" + ((("a" + ("hello" + "a")) + (("" + "foo bar") + (((("hello" + "hello") + ("foo bar" + "\n\t")) + "") + "a"))) + ("" + "a"))) {
                        case "\n\t":
                            var v17 uint32 = 2147483647
                            _ = v17
                        case "hello":
                            for fi18 := 0; fi18 < 2; fi18++ {
                                if len(v3) > 0 {
                                    v3[0] = v12
                                }
                                v1 += (v1 * 798)
                                v1 = (v1 & (v1 | (v1 >> ((v1) & 63))))
                                _ = fi18
                            }
                        case "a":
                            fmt.Println(fi14, v12)
                        }
                        for fi19 := 0; fi19 < 20; fi19++ {
                            if (-2147483648 != 0) {
                                var v20 int = (fi19 & fi14)
                                var v21 uint8 = 255
                                if false {
                                    v15 = v15
                                    if (!true) {
                                        var v22 uint64 = 18446744073709551615
                                        if ((3.14159 * float64(((0.1 + float32(3.14159)) / float32(-2.25)))) <= (float64((((0.1 + float32(0.1)) + 1e10) * float32(((1e10 + 1e10) - 0.0)))) + (0.1 + 1e10))) {
                                            v22 += (v22 << ((uint(v21)) & 63))
                                        }
                                        var v23 int = 855
                                        _ = v22
                                        _ = v23
                                    }
                                } else {
                                    var v24 int64 = 1
                                    _ = v24
                                }
                                _ = v20
                                _ = v21
                            } else {
                                v15 = (42 > 0)
                                if (v15 && v15) {
                                    var v25 int32 = -462
                                    _ = v25
                                }
                                v12 -= (v12 << (((v1 ^ (^(v1)))) & 63))
                                v12 = 0
                            }
                            _ = fi19
                        }
                        var v26 string = ""
                        _ = v26
                    } else {
                        v5 -= genFn1(530, 0)
                        if (0 < 0) {
                            if v15 {
                                switch (fi14 - (fi14 * fi14)) {
                                case -2147483648:
                                    if v15 {
                                        fmt.Println(fi14)
                                        for fi27 := 0; fi27 < 4; fi27++ {
                                            if true {
                                                v15 = (!((1 <= 42) || (!v15)))
                                                var v28 uint64 = 1
                                                var v29 uint32 = v12
                                                _ = v28
                                                _ = v29
                                            } else {
                                                var v30 int = -872
                                                v3 = append(v3, v12)
                                                _ = v30
                                            }
                                            var v31 bool = (!v15)
                                            v31 = v31
                                            _ = fi27
                                            _ = v31
                                        }
                                    } else {
                                        if false {
                                            v32 := GenS2{F0: 1}
                                            var v33 int64 = 0
                                            _ = v32
                                            _ = v33
                                        }
                                        fmt.Println(v15, v5, v15)
                                    }
                                case 42:
                                    v1 = 2
                                default:
                                    for fi34 := 0; fi34 < 7; fi34++ {
                                        var v35 uint64 = 4294967295
                                        v36 := GenS2{F0: -61}
                                        switch "\n\t" {
                                        case "hello":
                                            var v37 uint64 = 42
                                            _ = v37
                                        case "a":
                                            v38 := GenS3{F0: v1}
                                            _ = v38
                                        default:
                                            switch "foo bar" {
                                            case "\n\t":
                                                v39 := GenS4{F0: (("foo bar" + "a") + ("hello" + ("\n\t" + (("a" + "\n\t") + "foo bar")))), F1: genFn1(30, (v35 + (v35 ^ v35)))}
                                                _ = v39
                                            case "hello":
                                                var v40 int16 = v5
                                                _ = v40
                                            }
                                        }
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            }
                            var v41 int16 = (v5 & genFn1(997, 4294967295))
                            fmt.Println(v15, v4)
                            _ = v41
                        }
                        fmt.Println(v12)
                    }
                case 42:
                    v42 := GenS5{F0: -399, F1: genFn0()}
                    _ = v42
                case 2147483647:
                    switch 1 {
                    case 128:
                        var v43 float64 = 1e10
                        _ = v43
                    default:
                        for fi44 := 0; fi44 < 8; fi44++ {
                            fmt.Println(v5)
                            if (!((fi14 << ((333) & 63)) < fi14)) {
                                var v45 int = fi14
                                _ = v45
                            }
                            _ = fi44
                        }
                    }
                }
            }
            _ = fi14
            _ = v15
        }
        if false {
            v4 = (v4 & 392)
            v5 += 32767
            var v46 int64 = int64(v4)
            _ = v46
        } else {
            v47 := GenS2{F0: -99}
            if (true && true) {
                for fi48 := 0; fi48 < 15; fi48++ {
                    v1 = (v1 << ((v1) & 63))
                    v12 -= (v12 >> ((v1) & 63))
                    _ = fi48
                }
                var v49 uint = 42
                _ = v49
            }
            var v50 bool = (!(-2147483648 <= 1))
            _ = v47
            _ = v50
        }
        _ = v13
    }
    var v51 uint64 = 708
    v52 := GenS1{F0: (-(v5)), F1: "a"}
    var v53 int = 9223372036854775807
    var v54 float32 = (0.0 - float32(0.0))
    var v55 int64 = 9223372036854775807
    v54 -= float32(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
