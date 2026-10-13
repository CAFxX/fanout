package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint }
type GenS4 struct { F0 uint32; F1 int }
func genFn0(p0 string, p1 uint32, p2 uint64) int8 {
    v5 := []int{-664, -62, -901}
    var v6 float32 = 3.14159
    var v7 float32 = -2.25
    switch p0 {
    case "hello":
        var v8 uint = 2
        _ = v8
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "a":
        switch 42 {
        case 0:
            switch -1 {
            case 9223372036854775807:
                v9 := GenS1{F0: 856, F1: true, F2: 42}
                _ = v9
            case -953:
                fmt.Println(v6)
            case 32767:
                _ = len(v5)
            default:
                var v10 uint64 = (p2 | (^(p2)))
                _ = v10
            }
        case 32768:
            var v11 float32 = (v6 - v7)
            _ = v11
        case 690:
            var v12 uint16 = 255
            _ = v12
        }
    }
    var v13 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    _ = v6
    _ = v7
    return 79
}
func genFn1(p0 int32, p1 uint16) int8 {
    fmt.Println(p1, p0)
    var v14 uint = 65535
    v15 := []int{}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return -128
}

func main() {
    var v1 int32 = -705
    v2 := make(map[int]bool)
    v3 := []int{957, -48}
    var v4 int8 = -100
    switch "" {
    case "":
        v1 -= 243
    case "\n\t":
        if false {
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(v4, v1, v1)
                var v17 uint16 = 42
                v18 := GenS1{F0: 1, F1: true, F2: -32768}
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            var v19 float64 = ((1.5 + float64(0.0)) / 1.0)
            _ = v19
        }
    case "hello":
        v2[-295] = (false && false)
    }
    switch 1 {
    case 32767:
        v4 = (v4 & (v4 / ((v4) | 1)))
    }
    var v20 int32 = (v1 * (v1 % ((v1) | 1)))
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS2{F0: 157}
        v23 := GenS1{F0: 4294967295, F1: true, F2: 32767}
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int16 = 536
    var v25 uint64 = 1
    fmt.Println(v24, v25, v20)
    switch -9223372036854775808 {
    case 134:
        fmt.Println(v24)
    case 42:
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 uint32 = 4294967295
            var v28 int64 = -9223372036854775808
            v28 -= (v28 >> ((395) & 63))
            _ = fi26
            _ = v27
            _ = v28
        }
    case -9223372036854775808:
        v20 = 153
    }
    switch 1 {
    case -53:
        if true {
            switch 127 {
            case -1:
                for fi29 := 0; fi29 < 8; fi29++ {
                    v30 := GenS3{F0: 65535, F1: 1, F2: 65535}
                    v25 -= v25
                    _ = fi29
                    _ = v30
                }
            case 9223372036854775807:
                var v31 uint = 416
                _ = v31
            default:
                for fi32 := 0; fi32 < 19; fi32++ {
                    for fi33 := 0; fi33 < 16; fi33++ {
                        _ = len(v2)
                        _ = fi33
                    }
                    for fi34 := 0; fi34 < 17; fi34++ {
                        switch 128 {
                        case 1:
                            switch (fi32 + fi34) {
                            case 9223372036854775807:
                                if (v24 > v24) {
                                    v35 := GenS4{F0: 661, F1: fi34}
                                    _ = v35
                                }
                            case 2147483647:
                                v36 := GenS1{F0: v25, F1: (!false), F2: 884}
                                _ = v36
                            }
                        default:
                            var v37 float32 = -2.25
                            _ = v37
                        }
                        _ = fi34
                    }
                    v1 = v1
                    _ = fi32
                }
            }
        } else {
            v20 += v1
            v38 := GenS3{F0: 18, F1: 42, F2: 255}
            v20 = (v1 - (^(v1)))
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            v4 = (-(v4))
            var v41 bool = (766 > 9223372036854775807)
            var v42 int8 = 0
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS4{F0: 4294967295, F1: fi39}
        _ = fi39
        _ = v43
    }
    fmt.Println(v25, v24, v20)
    if (((3.14159 + -2.25) * 3.14159) <= (1.5 / 1.0)) {
        switch 9223372036854775807 {
        case 1:
            if (v1 != 316) {
                for fi44 := 0; fi44 < 6; fi44++ {
                    for fi45 := 0; fi45 < 19; fi45++ {
                        if true {
                            v4 -= (v4 ^ (v4 / ((-128) | 1)))
                        }
                        _ = fi45
                    }
                    var v46 int32 = -1
                    if false {
                        v25 = v25
                        v47 := GenS2{F0: 2}
                        v20 = -1
                        _ = v47
                    } else {
                        switch (^(fi44)) {
                        case 599:
                            v20 = v46
                        case -9223372036854775808:
                            delete(v2, (fi44 >> ((0) & 63)))
                        case 42:
                            v24 = v24
                        }
                    }
                    _ = fi44
                    _ = v46
                }
            }
        case 32768:
            var v48 int = 128
            _ = v48
        case 652:
            v49 := v2[889]
            _ = v49
        default:
            _ = len(v2)
        }
        v24 += 32767
        var v50 int = 32767
        _ = v50
    } else {
        v25 = (v25 ^ v25)
        var v51 int = 42
        _ = v51
    }
    v20 += (v20 - (v1 ^ (v1 / ((int32(v25)) | 1))))
    fmt.Println(v20)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
