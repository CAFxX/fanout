package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int }
func genFn0(p0 uint, p1 uint) uint {
    v7 := []int{-9223372036854775808}
    v8 := GenS1{F0: 0, F1: 494, F2: 80}
    p1 += (p0 | (p1 | 2))
    if (0 < 226) {
        v9 := GenS2{F0: ((false && (55 > 42)) && (!true))}
        p1 = p1
        fmt.Println(p1)
        _ = v9
    } else {
        switch "\n\t" {
        case "hello":
            var v10 uint = 433
            _ = v10
        case "\n\t":
            v11 := GenS1{F0: 239, F1: 333, F2: 42}
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 + p0)
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    v3 := []int{128, 2147483647}
    v4 := []string{}
    var v5 uint32 = 42
    var v6 uint64 = 2
    if false {
        v5 -= (v5 + uint32(v6))
        v12 := GenS3{F0: 65535}
        _ = v12
    } else {
        var v13 uint = 0
        if len(v4) > 0 {
            v4[0] = ("a" + "")
        }
        if (-2.25 < 0.0) {
            var v14 uint8 = 128
            _ = v14
        }
        _ = v13
    }
    v15 := GenS4{F0: 0, F1: 100, F2: 0}
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int8 = -128
        var v18 uint32 = (v5 >> ((1) & 63))
        if (!false) {
            v19 := GenS3{F0: 0}
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS3{F0: 2}
                if false {
                    if (!(false || ((9223372036854775807 == int64(fi16)) || false))) {
                        v3 = append(v3, (fi20 + fi16))
                        var v22 uint32 = (^(v18))
                        _ = len(v1)
                        _ = v22
                    } else {
                        switch "foo bar" {
                        case "hello":
                            var v23 uint32 = v18
                            _ = v23
                        }
                    }
                    var v24 int64 = 310
                    if (v5 < (v5 >> ((436) & 63))) {
                        v5 = (v18 & (v5 ^ v18))
                        v25 := GenS4{F0: v2, F1: v17, F2: 1}
                        var v26 uint16 = v2
                        _ = v25
                        _ = v26
                    }
                    _ = v24
                }
                if (("\n\t" + (("hello" + ("a" + "foo bar")) + "\n\t")) < "foo bar") {
                    v27 := v1[(fi16 * fi20)]
                    v28 := GenS5{F0: 128}
                    _ = v27
                    _ = v28
                }
                _ = fi20
                _ = v21
            }
            v29 := GenS3{F0: 255}
            _ = v19
            _ = v29
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    v30 := GenS4{F0: (^(v2)), F1: -43, F2: 0}
    v31 := GenS3{F0: 1}
    for fi32 := 0; fi32 < 4; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            v5 = v5
            _ = fi33
        }
        var v34 float32 = (1e10 / float32(1.5))
        _ = fi32
        _ = v34
    }
    switch "hello" {
    case "foo bar":
        switch "foo bar" {
        case "hello":
            if (!false) {
                v5 += v5
            }
        }
    default:
        v5 = 0
    }
    v2 = (v2 & v2)
    for fi35 := 0; fi35 < 12; fi35++ {
        v4 = append(v4, "hello")
        _ = fi35
    }
    v3 = append(v3, 127)
    for fi36 := 0; fi36 < 8; fi36++ {
        v2 = (v2 & (v2 % ((v2) | 1)))
        _ = fi36
    }
    switch 42 {
    case 127:
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 bool = ((1e10 + 0.0) != 1e10)
            _ = fi37
            _ = v38
        }
    case 0:
        v2 = 2
    default:
        var v39 string = (("" + "") + ("\n\t" + ("" + (("foo bar" + "hello") + "hello"))))
        _ = v39
    }
    switch -1 {
    case 9223372036854775807:
        var v40 bool = ((!(!(true || false))) || (false || true))
        _ = v40
    default:
        for fi41 := 0; fi41 < 11; fi41++ {
            v42 := GenS1{F0: 2, F1: -1, F2: 42}
            switch fi41 {
            case 4294967295:
                if false {
                    var v43 string = ("a" + ((("a" + "") + "") + "foo bar"))
                    _ = v43
                }
            default:
                if true {
                    for fi44 := 0; fi44 < 10; fi44++ {
                        v6 = v6
                        var v45 int64 = 0
                        switch fi41 {
                        case -9223372036854775808:
                            fmt.Println(fi44)
                        }
                        if true {
                            v5 -= (v5 * uint32(v2))
                            v46 := GenS6{F0: (fi44 & int(v2))}
                            _ = v46
                        }
                        _ = fi44
                        _ = v45
                    }
                    var v47 int = (fi41 | (fi41 | (fi41 & 128)))
                    switch "\n\t" {
                    case "":
                        var v48 bool = true
                        _ = v48
                    case "\n\t":
                        v4 = append(v4, ("hello" + "a"))
                    }
                    _ = v47
                }
            }
            _ = fi41
            _ = v42
        }
    }
    fmt.Println(v2, v6)
    fmt.Println(v5, v5, v6)
    var v49 int32 = 42
    var v50 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
