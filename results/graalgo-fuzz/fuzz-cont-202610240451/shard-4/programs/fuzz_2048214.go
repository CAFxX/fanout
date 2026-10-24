package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 bool; F1 int }
func genFn0() uint64 {
    var v7 uint64 = 42
    if (v7 <= v7) {
        if (617 > 0) {
            switch "hello" {
            case "hello":
                var v8 uint32 = 97
                _ = v8
            case "a":
                v9 := []int{}
                _ = v9
            case "":
                var v10 uint64 = (v7 % (((^(v7))) | 1))
                _ = v10
            }
            fmt.Println(v7)
        } else {
            var v11 uint = 4294967295
            var v12 uint64 = v7
            v11 = 4294967295
            _ = v11
            _ = v12
        }
    }
    _ = v7
    return 57
}

func main() {
    var v1 string = ""
    v2 := []uint32{1}
    var v3 uint = 255
    var v4 int16 = 32767
    v5 := []string{(v1 + ((v1 + ("\n\t" + v1)) + "\n\t"))}
    v6 := make(map[uint32]uint32)
    v4 = (v4 - -32768)
    var v13 float32 = ((3.14159 * 1e10) / ((0.1 * 3.14159) + 3.14159))
    if (65535 != 2) {
        var v14 int = -187
        var v15 int8 = 1
        var v16 int = (v14 + (v14 - 2147483647))
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        delete(v6, 941)
        _ = fi17
    }
    v18 := GenS1{F0: 108}
    var v19 string = "hello"
    var v20 int = -665
    if true {
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 int = fi21
            _ = fi21
            _ = v22
        }
        v13 -= v13
        v23 := GenS2{F0: true, F1: 110}
        _ = v23
    } else {
        v20 += v20
        _ = len(v2)
        var v24 uint = (^(v3))
        _ = v24
    }
    v25 := GenS1{F0: -128}
    var v26 int32 = -2147483648
    var v27 int = (v20 & (^(v20)))
    _ = len(v2)
    fmt.Println(v13)
    for fi28 := 0; fi28 < 3; fi28++ {
        switch "\n\t" {
        case "foo bar":
            switch v1 {
            case "":
                fmt.Println(v1, v1, v13)
            default:
                var v29 uint16 = 255
                _ = v29
            }
        case "\n\t":
            v30 := GenS3{F0: true, F1: (v20 / ((fi28) | 1))}
            _ = v30
        default:
            v31 := GenS2{F0: (!false), F1: 161}
            _ = v31
        }
        _ = fi28
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := GenS2{F0: false, F1: 128}
            switch "a" {
            case "a":
                v3 = (v3 << (((v3 % (((v3 >> ((v3) & 63))) | 1))) & 63))
            case "\n\t":
                v1 = v1
            case "foo bar":
                delete(v6, 2)
            default:
                v35 := v6[1]
                _ = v35
            }
            var v36 uint8 = uint8(v3)
            _ = fi33
            _ = v34
            _ = v36
        }
        v26 -= (v26 ^ (v26 << ((0) & 63)))
        _ = fi32
    }
    v13 += v13
    var v37 float64 = (0.1 * 0.1)
    if false {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        for fi38 := 0; fi38 < 16; fi38++ {
            v6[4294967295] = 1
            if false {
                v20 += (fi38 >> ((65535) & 63))
            }
            if false {
                switch 1 {
                case 2147483647:
                    v26 -= 415
                case -9223372036854775808:
                    switch v1 {
                    case "foo bar":
                        for fi39 := 0; fi39 < 3; fi39++ {
                            var v40 uint8 = 66
                            if true {
                                v6[65535] = 0
                            }
                            var v41 float32 = (-2.25 + 1e10)
                            var v42 bool = false
                            v3 -= v3
                            _ = fi39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    case "\n\t":
                        delete(v6, 283)
                    case "hello":
                        v19 = v1
                    default:
                        v26 = v26
                    }
                default:
                    var v43 uint = v3
                    _ = v43
                }
                for fi44 := 0; fi44 < 16; fi44++ {
                    var v45 uint16 = 32767
                    v2 = append(v2, 933)
                    _ = fi44
                    _ = v45
                }
                v3 -= v3
            }
            _ = fi38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
