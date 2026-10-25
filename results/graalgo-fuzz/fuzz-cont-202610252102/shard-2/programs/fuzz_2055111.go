package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int8; F1 bool }
func genFn0() string {
    for fi4 := 0; fi4 < 10; fi4++ {
        fmt.Println(fi4)
        var v5 float32 = float32(1e10)
        _ = fi4
        _ = v5
    }
    var v6 int16 = 0
    _ = v6
    return "a"
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]string)
    var v3 uint8 = 1
    v3 += 157
    if false {
        fmt.Println(v1)
        var v7 string = "hello"
        _ = v7
    }
    if false {
        v1 = (((float32((1.5 - 1.5)) * (v1 + ((float32((1e10 * (0.1 - 0.1))) * v1) + (((v1 / 1.0) - -2.25) + ((v1 - v1) - ((3.14159 * v1) + v1)))))) + v1) + v1)
        v2[0] = ""
        v8 := GenS1{F0: -32768}
        _ = v8
    } else {
        if ((127 > 127) || false) {
            switch 0 {
            case 9223372036854775807:
                fmt.Println(v3)
            case -287:
                var v9 int = 0
                _ = v9
            }
            v1 -= ((float32(0.1) - v1) - float32(0.0))
            var v10 float64 = 3.14159
            _ = v10
        }
        if true {
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := []int{}
                var v13 uint16 = 819
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    fmt.Println(v1, v1)
    for fi14 := 0; fi14 < 9; fi14++ {
        v3 = 128
        _ = fi14
    }
    v15 := []int{-9223372036854775808}
    fmt.Println(v3, v1)
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    var v17 uint = 1
    for fi18 := 0; fi18 < 20; fi18++ {
        switch "\n\t" {
        case "hello":
            if (!false) {
                fmt.Println(v17, v17)
                var v19 int64 = -1
                _ = v19
            }
        default:
            v20 := GenS1{F0: -932}
            _ = v20
        }
        v15 = append(v15, 4294967295)
        for fi21 := 0; fi21 < 20; fi21++ {
            fmt.Println(fi18, fi21)
            v22 := GenS2{F0: -661, F1: (v17 & (^(v17)))}
            _ = fi21
            _ = v22
        }
        _ = fi18
    }
    v1 += float32(0.1)
    switch "\n\t" {
    case "hello":
        if false {
            if false {
                if false {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        fmt.Println(v17)
                        _ = fi23
                    }
                    delete(v2, 42)
                    v3 -= 128
                }
            }
            v15 = append(v15, -9223372036854775808)
            fmt.Println(v17, v1, v3)
        } else {
            v1 += (1e10 + 3.14159)
            v3 = v3
        }
    default:
        switch -263 {
        case -509:
            v24 := GenS1{F0: 0}
            _ = v24
        case 4294967295:
            for fi25 := 0; fi25 < 10; fi25++ {
                switch 32767 {
                case 1:
                    for fi26 := 0; fi26 < 20; fi26++ {
                        if (true || (!((106 == 100) && false))) {
                            var v27 uint16 = 42
                            _ = v27
                        } else {
                            var v28 int = (fi25 & (fi25 + (fi26 >> ((v17) & 63))))
                            v1 -= v1
                            var v29 uint8 = v3
                            var v30 uint16 = 1
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        switch fi26 {
                        case 32768:
                            if true {
                                var v31 int32 = -217
                                var v32 uint16 = 255
                                _ = v31
                                _ = v32
                            }
                        case 32767:
                            _ = len(v2)
                            v33 := GenS3{F0: 2, F1: ("" + (("hello" + "") + "a"))}
                            _ = v33
                        }
                        switch fi25 {
                        case 2147483647:
                            var v34 uint64 = 785
                            _ = v34
                        case -614:
                            delete(v2, 2147483647)
                        }
                        _ = fi26
                    }
                case -9223372036854775808:
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v3 += (v3 * (v3 / ((v3) | 1)))
                        v1 = 1e10
                        if (!(false || true)) {
                            var v36 bool = ((false || (false && (9223372036854775807 > fi35))) || (!(!(v3 < 127))))
                            if true {
                                v37 := GenS2{F0: 42, F1: 18446744073709551615}
                                v17 = 18446744073709551615
                                fmt.Println(fi35, v1, fi35)
                                _ = v37
                            }
                            v38 := GenS2{F0: 42, F1: (^(v17))}
                            _ = v36
                            _ = v38
                        }
                        _ = fi35
                    }
                case -718:
                    if (255 > 1) {
                        var v39 int32 = 1
                        if true {
                            for fi40 := 0; fi40 < 11; fi40++ {
                                v1 = float32(1e10)
                                v41 := GenS4{F0: 0, F1: 42}
                                v3 = (v3 ^ 1)
                                _ = fi40
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 12; fi42++ {
                                v43 := GenS3{F0: 4294967295, F1: "hello"}
                                v3 = (v3 << (((v17 % ((uint(v3)) | 1))) & 63))
                                var v44 int = 0
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                        } else {
                            v1 = v1
                        }
                        var v45 uint8 = v3
                        _ = v39
                        _ = v45
                    }
                }
                v1 -= v1
                _ = fi25
            }
        case -242:
            v17 -= (v17 & v17)
        default:
            var v46 int32 = 0
            _ = v46
        }
    }
    var v47 int8 = 100
    v1 = (v1 / (v1 * float32(3.14159)))
    fmt.Println(v47, v3, v47)
    switch -1 {
    case 32767:
        if (42 >= 18446744073709551615) {
            v17 -= uint(v3)
            v3 = (v3 * v3)
        } else {
            for fi48 := 0; fi48 < 9; fi48++ {
                v49 := v2[uint32(v3)]
                var v50 int8 = (v47 & v47)
                _ = fi48
                _ = v49
                _ = v50
            }
        }
    default:
        v51 := GenS5{F0: (v47 % ((v47) | 1)), F1: (("foo bar" + "hello") < "\n\t")}
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v47
}
