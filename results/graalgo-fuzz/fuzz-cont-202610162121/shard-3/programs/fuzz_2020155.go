package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint; F2 string }
type GenS4 struct { F0 string }
func genFn0() bool {
    v4 := []int{1, -899, 19, 32767}
    var v5 string = ""
    v4 = append(v4, 32767)
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 int, p1 string, p2 uint16) int16 {
    p0 -= int(p2)
    var v6 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1000
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    v3 := []int{-9223372036854775808, 0, 849}
    var v7 int32 = 42
    switch "hello" {
    case "":
        var v8 int16 = -317
        _ = v8
    case "a":
        v9 := GenS1{F0: v1, F1: 802}
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        if (!true) {
            v11 := GenS1{F0: (^(v1)), F1: 206}
            if false {
                if true {
                    var v12 uint = 443
                    var v13 uint32 = 42
                    v1 += (^(v1))
                    _ = v12
                    _ = v13
                }
                switch (-(fi10)) {
                case 1:
                    for fi14 := 0; fi14 < 12; fi14++ {
                        var v15 string = (("a" + "") + (("" + "foo bar") + "\n\t"))
                        var v16 uint8 = (^(v1))
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                default:
                    var v17 int32 = v7
                    _ = v17
                }
            } else {
                if (1.5 >= 0.1) {
                    var v18 uint32 = uint32(v1)
                    _ = v18
                }
                _ = len(v2)
                if false {
                    var v19 float64 = (1e10 + 1e10)
                    var v20 float32 = (3.14159 * 3.14159)
                    if false {
                        for fi21 := 0; fi21 < 19; fi21++ {
                            if (9223372036854775807 > -1) {
                                v22 := v2[(-(fi21))]
                                if (432 != 255) {
                                    var v23 uint = 42
                                    v20 = (v20 + 3.14159)
                                    _ = v23
                                }
                                var v24 float64 = v19
                                v25 := GenS1{F0: 1, F1: -1}
                                _ = v22
                                _ = v24
                                _ = v25
                            } else {
                                if true {
                                    var v26 int = 42
                                    v20 = (v20 / v20)
                                    switch fi10 {
                                    case 1:
                                        var v27 string = "hello"
                                        _ = v27
                                    default:
                                        var v28 int8 = -35
                                        _ = v28
                                    }
                                    _ = v26
                                } else {
                                    fmt.Println(v19, v1)
                                }
                                v29 := GenS1{F0: 255, F1: genFn1(0, (((("" + "hello") + "hello") + ("" + (("foo bar" + ((("a" + "") + ("" + "")) + "\n\t")) + "a"))) + ""), 204)}
                                _ = v29
                            }
                            v2[fi21] = 42
                            v30 := GenS1{F0: (v1 - v1), F1: 1000}
                            v31 := GenS2{F0: -118}
                            _ = fi21
                            _ = v30
                            _ = v31
                        }
                    } else {
                        v7 = (v7 >> ((1) & 63))
                    }
                    _ = v19
                    _ = v20
                }
            }
            var v32 bool = (1.5 < 1.5)
            _ = v11
            _ = v32
        }
        var v33 float32 = 1e10
        if true {
            v33 += (-2.25 / 1.5)
            if false {
                v34 := GenS3{F0: 1, F1: 65535, F2: "hello"}
                if false {
                    var v35 uint = 65535
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 20; fi36++ {
                        var v37 bool = false
                        var v38 uint16 = 2
                        fmt.Println(v33)
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
                _ = v34
            } else {
                v1 = v1
                _ = len(v3)
                v7 += (v7 ^ v7)
            }
            var v39 bool = false
            _ = v39
        } else {
            fmt.Println(v33)
        }
        _ = fi10
        _ = v33
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := v2[(fi40 % ((fi40) | 1))]
        _ = fi40
        _ = v41
    }
    if len(v3) > 0 {
        v3[0] = int(v7)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        var v42 int64 = 0
        var v43 uint16 = 914
        v7 = v7
        _ = v42
        _ = v43
    } else {
        switch "" {
        case "a":
            if (0 > 0) {
                v7 += (v7 >> ((65535) & 63))
                switch "foo bar" {
                case "a":
                    v44 := GenS3{F0: 688, F1: 787, F2: ""}
                    _ = v44
                }
                fmt.Println(v7)
            }
        case "hello":
            var v45 int = 4294967295
            _ = v45
        }
        var v46 float64 = 1.5
        _ = v46
    }
    var v47 bool = false
    var v48 uint16 = 816
    fmt.Println(v1, v47)
    delete(v2, 68)
    if (!false) {
        var v49 uint32 = 602
        v50 := GenS4{F0: "a"}
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        _ = v49
        _ = v50
    }
    v51 := GenS2{F0: 32767}
    v7 = v7
    var v52 uint16 = v48
    var v53 float32 = 1e10
    for fi54 := 0; fi54 < 5; fi54++ {
        v7 += (v7 ^ 42)
        _ = fi54
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        var v56 int8 = -1
        v2[(^(fi55))] = 253
        switch "" {
        case "":
            v57 := GenS3{F0: v7, F1: uint(v1), F2: "foo bar"}
            _ = v57
        case "hello":
            var v58 float32 = -2.25
            _ = v58
        case "\n\t":
            fmt.Println(v7, v53, v7)
        default:
            var v59 uint8 = 0
            _ = v59
        }
        _ = fi55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v51
    _ = v52
    _ = v53
    _ = v7
}
