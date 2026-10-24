package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 int64; F1 bool }
func genFn0(p0 uint16, p1 bool) bool {
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(p1, fi5, fi5)
        switch fi5 {
        case -845:
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := []int{(^(fi5)), fi6, -437, fi6}
                var v8 uint16 = (p0 - (p0 << ((269) & 63)))
                _ = fi6
                _ = v7
                _ = v8
            }
        case 1:
            if p1 {
                switch "foo bar" {
                case "":
                    fmt.Println(p1, p0, fi5)
                case "a":
                    var v9 int8 = int8(fi5)
                    _ = v9
                case "\n\t":
                    if (4 > -24) {
                        p1 = p1
                    }
                default:
                    if p1 {
                        v10 := GenS1{F0: 65535, F1: "hello", F2: 1}
                        _ = v10
                    } else {
                        v11 := GenS2{F0: 1}
                        if (182 > 2) {
                            var v12 int = fi5
                            for fi13 := 0; fi13 < 9; fi13++ {
                                v14 := GenS1{F0: 0, F1: "\n\t", F2: 162}
                                _ = fi13
                                _ = v14
                            }
                            var v15 bool = (((v12 % (((^(fi5))) | 1)) != (fi5 & fi5)) || p1)
                            _ = v12
                            _ = v15
                        }
                        if (-795 == 0) {
                            p1 = (18446744073709551615 >= 42)
                            fmt.Println(p0)
                            switch fi5 {
                            case 32767:
                                v16 := GenS1{F0: p0, F1: "\n\t", F2: 42}
                                _ = v16
                            default:
                                if (p1 && (1.5 < ((1e10 + 0.0) * (float64(((3.14159 * 0.0) - -2.25)) * -2.25)))) {
                                    p0 = 42
                                    var v17 int16 = 32767
                                    v17 = 406
                                    _ = v17
                                } else {
                                    if p1 {
                                        p1 = (1 < 39)
                                        for fi18 := 0; fi18 < 4; fi18++ {
                                            p1 = p1
                                            _ = fi18
                                        }
                                        var v19 bool = p1
                                        _ = v19
                                    } else {
                                        var v20 int8 = -1
                                        _ = v20
                                    }
                                    v21 := []int{-9223372036854775808}
                                    v22 := GenS3{F0: 42, F1: fi5}
                                    _ = v21
                                    _ = v22
                                }
                            }
                        } else {
                            var v23 int8 = -1
                            v24 := GenS2{F0: 343}
                            p0 = p0
                            _ = v23
                            _ = v24
                        }
                        _ = v11
                    }
                }
            }
        default:
            v25 := GenS4{F0: 65535, F1: p0, F2: 46}
            _ = v25
        }
        _ = fi5
    }
    if (1 <= -494) {
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 int8 = 13
            v28 := GenS5{F0: 164, F1: 0}
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 uint32 = 65535
        _ = v29
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]uint32)
    v3 := []string{"hello"}
    v4 := make(map[string]string)
    var v30 int16 = 1
    if false {
        if (true || true) {
            v31 := GenS4{F0: 65535, F1: 65535, F2: 0}
            var v32 string = ("hello" + "")
            v32 = v32
            _ = v31
            _ = v32
        } else {
            v30 = (v30 | (v30 | v30))
            var v33 float32 = (1e10 + float32(-2.25))
            _ = v33
        }
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 uint32 = 320
            fmt.Println(v1, v35)
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        switch ("" + ("foo bar" + ("a" + "\n\t"))) {
        case "foo bar":
            delete(v4, "\n\t")
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v30 -= v30
        if true {
            var v37 uint32 = 835
            _ = v37
        }
        _ = fi36
    }
    switch "a" {
    case "":
        if (0 <= -38) {
            fmt.Println(v30, v1)
        }
    case "a":
        v30 -= 577
    default:
        _ = len(v2)
    }
    v38 := GenS6{F0: -1, F1: 410}
    var v39 float64 = (-2.25 + (0.1 / 0.1))
    var v40 uint16 = 1
    var v41 uint16 = 42
    var v42 float32 = -2.25
    v41 += v40
    switch "\n\t" {
    case "\n\t":
        v43 := GenS5{F0: 50, F1: 36}
        _ = v43
    case "foo bar":
        if true {
            v44 := GenS2{F0: v30}
            _ = v44
        } else {
            v45 := GenS7{F0: 9223372036854775807, F1: ((9223372036854775807 <= 1) || false)}
            _ = v45
        }
    default:
        var v46 string = "foo bar"
        _ = v46
    }
    v47 := GenS3{F0: v30, F1: 0}
    for fi48 := 0; fi48 < 6; fi48++ {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi48
    }
    switch "a" {
    case "hello":
        var v49 int64 = (v1 % ((v1) | 1))
        _ = v49
    case "a":
        _ = len(v4)
    default:
        _ = len(v3)
    }
    fmt.Println(v30, v30, v42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v47
}
