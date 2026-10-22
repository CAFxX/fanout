package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 bool }
func genFn0(p0 uint32, p1 int8, p2 uint8) int32 {
    v4 := []int{4294967295, -2147483648}
    for fi5 := 0; fi5 < 18; fi5++ {
        p1 += p1
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 42, F1: "foo bar"}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if false {
        p0 = (p0 ^ p0)
        v4 = append(v4, 9223372036854775807)
    }
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]bool)
    v3 := []string{"foo bar", "hello", "a"}
    var v8 string = "\n\t"
    v1 -= v1
    for fi9 := 0; fi9 < 19; fi9++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 uint64 = 4294967295
            if false {
                v12 := GenS2{F0: 128}
                _ = len(v3)
                v11 = v11
                v3 = append(v3, (v8 + ((((v8 + ("\n\t" + ("" + v8))) + v8) + (((("hello" + v8) + v8) + v8) + ("hello" + "hello"))) + v8)))
                _ = v12
            } else {
                delete(v2, fi9)
                switch fi10 {
                case 32767:
                    var v13 int64 = v1
                    _ = v13
                }
            }
            v14 := GenS2{F0: 1}
            _ = fi10
            _ = v11
            _ = v14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 string = v8
            if (1 < -54) {
                var v17 uint64 = 1
                v2[fi9] = (true && (110 <= -128))
                var v18 int16 = -512
                _ = v17
                _ = v18
            }
            v1 = (v1 >> ((4294967295) & 63))
            _ = fi15
            _ = v16
        }
        _ = fi9
    }
    if false {
        v8 = ((("hello" + v8) + "foo bar") + v8)
        v19 := GenS3{F0: 2147483647, F1: 485, F2: (!(42 < 1))}
        for fi20 := 0; fi20 < 10; fi20++ {
            v2[9223372036854775807] = ((0.0 == -2.25) || true)
            _ = len(v2)
            _ = fi20
        }
        _ = v19
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            if (true && false) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    _ = len(v2)
                    _ = fi22
                }
                if false {
                    v23 := GenS2{F0: 1}
                    _ = v23
                }
                var v24 bool = true
                _ = v24
            }
            v25 := GenS1{F0: genFn0(2147483647, -56, 255), F1: v8}
            v1 -= v1
            _ = fi21
            _ = v25
        }
        v8 = "\n\t"
    }
    v8 = v8
    v8 = ((v8 + ("hello" + "\n\t")) + (v8 + ((v8 + v8) + v8)))
    _ = len(v2)
    var v26 int32 = 1
    if false {
        fmt.Println(v26, v8)
        var v27 uint32 = 2
        _ = v27
    } else {
        _ = len(v3)
    }
    switch v8 {
    case "":
        var v28 uint8 = 255
        _ = v28
    }
    if (!(true && true)) {
        _ = len(v2)
    } else {
        _ = len(v3)
        v29 := GenS1{F0: v26, F1: v8}
        if (-709 == 1000) {
            v26 = v26
            if true {
                var v30 bool = true
                fmt.Println(v8, v1, v8)
                switch 505 {
                case -2147483648:
                    v1 = -9223372036854775808
                default:
                    switch "a" {
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "":
                        for fi31 := 0; fi31 < 4; fi31++ {
                            switch fi31 {
                            case 2147483647:
                                fmt.Println(v30, fi31)
                            case 9223372036854775807:
                                v8 = (v8 + v8)
                            default:
                                switch "a" {
                                case "\n\t":
                                    v30 = v30
                                default:
                                    fmt.Println(v8, v8, v26)
                                }
                            }
                            _ = fi31
                        }
                    case "hello":
                        var v32 int32 = -283
                        _ = v32
                    }
                }
                _ = v30
            } else {
                switch -1 {
                case 128:
                    if (!(!(42 > 1))) {
                        var v33 bool = true
                        fmt.Println(v1)
                        var v34 int32 = v26
                        _ = v33
                        _ = v34
                    }
                case -2147483648:
                    fmt.Println(v8)
                default:
                    switch v8 {
                    case "hello":
                        if false {
                            var v35 int32 = (v26 >> ((157) & 63))
                            v35 -= (-(v26))
                            var v36 int32 = (v35 & genFn0(587, 127, 0))
                            _ = v35
                            _ = v36
                        }
                    case "\n\t":
                        var v37 uint16 = 65535
                        _ = v37
                    default:
                        var v38 float32 = ((float32((0.0 * 1e10)) / 3.14159) / 1.0)
                        _ = v38
                    }
                }
                v39 := GenS1{F0: (^(v26)), F1: v8}
                _ = v39
            }
        }
        _ = v29
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        for fi41 := 0; fi41 < 18; fi41++ {
            var v42 int = 0
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    delete(v2, 42)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v43 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v43
    _ = v8
}
