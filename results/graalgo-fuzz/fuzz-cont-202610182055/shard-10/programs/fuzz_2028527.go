package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint32; F1 int16; F2 uint16 }
func genFn0(p0 int, p1 int64) uint16 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: (true && false), F1: 42}
            v8 := []int{fi6}
            if false {
                p0 += (fi5 * 0)
                p0 -= 32767
                switch fi6 {
                case 9223372036854775807:
                    p1 = -1
                case -1:
                    var v9 float32 = 0.0
                    _ = v9
                case 1:
                    v10 := GenS2{F0: 42}
                    p0 = p0
                    var v11 int16 = 0
                    var v12 int16 = v11
                    switch "foo bar" {
                    case "":
                        var v13 uint32 = 2147483647
                        switch p0 {
                        case -1:
                            v8 = append(v8, (fi6 & -284))
                        default:
                            var v14 int = (^(p0))
                            _ = v14
                        }
                        _ = v13
                    case "\n\t":
                        p0 = fi5
                    default:
                        switch fi5 {
                        case 0:
                            if len(v8) > 0 {
                                v8[0] = fi6
                            }
                        case 128:
                            fmt.Println(fi5, p0, fi5)
                        default:
                            _ = len(v8)
                        }
                        v15 := GenS2{F0: 65535}
                        fmt.Println(p1)
                        if (65535 >= 925) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        } else {
                            v12 -= -320
                            var v16 uint64 = 1
                            v17 := GenS2{F0: 2}
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    }
                    _ = v10
                    _ = v11
                    _ = v12
                }
                var v18 uint64 = 1
                _ = v18
            } else {
                p1 = p1
            }
            var v19 uint8 = 255
            _ = fi6
            _ = v19
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    if (128 < 255) {
        p0 += (p0 - 32768)
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int = (fi20 - p0)
            _ = fi20
            _ = v21
        }
        p0 = 128
        var v22 uint64 = 42
        _ = v22
    }
    p0 = (-(p0))
    if (127 < 42) {
        var v23 uint16 = 2
        p0 = (-(p0))
        _ = v23
    } else {
        switch p0 {
        case 42:
            switch "foo bar" {
            case "":
                fmt.Println(p1)
            case "\n\t":
                p0 -= (p0 >> ((283) & 63))
            default:
                v24 := []int{p0, 9223372036854775807, p0}
                _ = v24
            }
        }
    }
    p1 += (p1 & p1)
    _ = p0
    _ = p1
    return 237
}

func main() {
    var v1 int64 = -1
    var v2 int16 = 121
    var v3 int8 = 40
    v4 := make(map[int]string)
    v2 = (v2 + (v2 + -1))
    var v25 uint64 = 1
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS1{F0: false, F1: 2}
        if true {
            for fi28 := 0; fi28 < 2; fi28++ {
                v25 += (^(v25))
                _ = fi28
            }
        }
        _ = fi26
        _ = v27
    }
    var v29 uint = 615
    v4[-9223372036854775808] = "\n\t"
    var v30 int64 = -1
    v31 := GenS3{F0: -949, F1: 1, F2: (v29 > 255)}
    if true {
        v4[127] = ""
        var v32 int64 = int64(v3)
        v33 := GenS4{F0: false}
        _ = v32
        _ = v33
    }
    if false {
        fmt.Println(v29, v2)
        switch ((((("" + "foo bar") + (("foo bar" + "a") + "hello")) + "\n\t") + ("hello" + "")) + "foo bar") {
        case "":
            delete(v4, -2147483648)
        case "hello":
            var v34 bool = (genFn0(4294967295, (v1 - 0)) != 42)
            _ = v34
        case "foo bar":
            switch "\n\t" {
            case "":
                var v35 int16 = (^(v2))
                _ = v35
            }
        }
    } else {
        var v36 uint32 = 253
        delete(v4, 2147483647)
        var v37 float32 = 3.14159
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        v39 := GenS2{F0: 0}
        fmt.Println(v1)
        if false {
            _ = len(v4)
            var v40 int = fi38
            v41 := GenS5{F0: -1, F1: 1}
            _ = v40
            _ = v41
        } else {
            v42 := v4[fi38]
            _ = v42
        }
        _ = fi38
        _ = v39
    }
    var v43 float32 = (-2.25 + -2.25)
    v44 := GenS6{F0: 4294967295, F1: v2, F2: genFn0(128, v1)}
    v29 += 4294967295
    v45 := v4[975]
    switch (v45 + v45) {
    case "foo bar":
        fmt.Println(v25, v43, v45)
    case "":
        v46 := GenS4{F0: ("" >= (((v45 + (v45 + v45)) + (v45 + "")) + v45))}
        _ = v46
    case "a":
        var v47 uint8 = 255
        _ = v47
    default:
        for fi48 := 0; fi48 < 11; fi48++ {
            switch fi48 {
            case 42:
                v1 = v30
            case 9223372036854775807:
                v4[fi48] = (v45 + (v45 + v45))
            }
            if true {
                var v49 uint64 = v25
                _ = v49
            }
            v1 = (v30 ^ (v1 + v1))
            _ = fi48
        }
    }
    if false {
        var v50 uint32 = 491
        var v51 uint8 = 2
        _ = v50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
