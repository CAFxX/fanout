package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS8 struct { F0 int }
func genFn0(p0 uint8, p1 uint32, p2 int64) int32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        switch "foo bar" {
        case "foo bar":
            p1 = 4294967295
        case "\n\t":
            if false {
                p0 = (^(p0))
            } else {
                switch (fi6 % (((fi6 | fi6)) | 1)) {
                case 32768:
                    p1 += 598
                case 211:
                    if ((true || false) && true) {
                        v7 := GenS1{F0: -907}
                        _ = v7
                    } else {
                        p1 += 490
                        p1 += p1
                        fmt.Println(p1, fi6)
                    }
                default:
                    p0 -= p0
                }
            }
        case "hello":
            if (-2147483648 < 493) {
                var v8 float64 = float64((1.5 / 1.0))
                v8 -= v8
                if false {
                    fmt.Println(p2, p0, fi6)
                    p2 = p2
                    for fi9 := 0; fi9 < 13; fi9++ {
                        var v10 uint = 65535
                        if ((p1 | p1) > 1) {
                            v11 := GenS2{F0: fi6, F1: p0, F2: p2}
                            v8 -= -2.25
                            var v12 uint8 = 255
                            _ = v11
                            _ = v12
                        }
                        var v13 string = "foo bar"
                        _ = fi9
                        _ = v10
                        _ = v13
                    }
                }
                _ = v8
            } else {
                v14 := GenS3{F0: (2 > 4294967295)}
                var v15 int8 = -109
                switch int(v15) {
                case 521:
                    v16 := GenS4{F0: 0}
                    _ = v16
                case 42:
                    if false {
                        var v17 int32 = 42
                        if ("a" != "hello") {
                            v18 := GenS1{F0: 1}
                            for fi19 := 0; fi19 < 18; fi19++ {
                                v20 := []int{}
                                v21 := GenS1{F0: -755}
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            var v22 uint32 = uint32(p0)
                            _ = v18
                            _ = v22
                        }
                        _ = v17
                    }
                default:
                    var v23 uint8 = p0
                    _ = v23
                }
                _ = v14
                _ = v15
            }
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 int8 = 1
            fmt.Println(v25)
            var v26 uint32 = (p1 ^ p1)
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi6
    }
    switch int(p0) {
    case 128:
        v27 := GenS3{F0: (318 > (p1 / ((733) | 1)))}
        _ = v27
    default:
        for fi28 := 0; fi28 < 1; fi28++ {
            v29 := GenS5{F0: (^(p1)), F1: (p1 | p1)}
            v30 := GenS6{F0: p1, F1: 2, F2: 2}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    var v31 uint16 = 808
    var v32 uint = uint(p1)
    switch ("\n\t" + "a") {
    case "":
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 6; fi34++ {
                if true {
                    v35 := GenS3{F0: (false || (3.14159 < -2.25))}
                    var v36 int16 = -32768
                    fmt.Println(v31)
                    _ = v35
                    _ = v36
                } else {
                    var v37 int64 = (p2 / ((p2) | 1))
                    var v38 uint = (v32 + (v32 - v32))
                    v31 += v31
                    p0 -= 1
                    _ = v37
                    _ = v38
                }
                var v39 uint16 = (v31 & 1)
                _ = fi34
                _ = v39
            }
            switch "foo bar" {
            case "hello":
                p2 += 42
            case "a":
                switch fi33 {
                case -1:
                    for fi40 := 0; fi40 < 4; fi40++ {
                        p1 = p1
                        v41 := GenS2{F0: (fi40 ^ (fi40 >> ((uint(p1)) & 63))), F1: (^(p0)), F2: p2}
                        _ = fi40
                        _ = v41
                    }
                default:
                    v42 := []int{32768, fi33, (^(fi33))}
                    _ = v42
                }
            case "":
                v43 := []int{fi33, (fi33 - -168), int(p0)}
                _ = v43
            }
            _ = fi33
        }
    case "foo bar":
        var v44 uint16 = v31
        _ = v44
    default:
        v32 += v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v32
    return -504
}
func genFn1() string {
    for fi45 := 0; fi45 < 16; fi45++ {
        if true {
            v46 := []int{9223372036854775807}
            for fi47 := 0; fi47 < 7; fi47++ {
                var v48 int8 = 1
                _ = fi47
                _ = v48
            }
            _ = v46
        }
        _ = fi45
    }
    var v49 float64 = float64(0.0)
    v49 += 3.14159
    _ = v49
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 108
    var v3 bool = ("" <= "foo bar")
    var v4 bool = v3
    v5 := []string{("a" + ("" + "foo bar")), "foo bar", "hello", ("a" + ("a" + ""))}
    var v50 float32 = (float32(-2.25) - float32((1e10 - (1e10 * 0.1))))
    switch 128 {
    case 127:
        v4 = v4
    case -387:
        var v51 int = 127
        _ = v51
    }
    _ = len(v1)
    if v3 {
        v52 := GenS7{F0: v2, F1: 0, F2: 2}
        _ = v52
    }
    var v53 uint32 = 0
    for fi54 := 0; fi54 < 3; fi54++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi55 := 0; fi55 < 3; fi55++ {
            var v56 uint64 = 42
            v56 = 0
            _ = fi55
            _ = v56
        }
        if v4 {
            if v3 {
                if true {
                    switch "a" {
                    case "hello":
                        fmt.Println(fi54, v2)
                    case "":
                        v53 -= v53
                    default:
                        var v57 uint = 2
                        _ = v57
                    }
                }
            } else {
                if v3 {
                    _ = len(v5)
                    var v58 float64 = 3.14159
                    _ = v58
                }
            }
            delete(v1, v53)
        } else {
            v4 = (fi54 > fi54)
        }
        _ = fi54
    }
    var v59 string = "\n\t"
    if true {
        v3 = (42 > 214)
    }
    switch int(v2) {
    case 0:
        if (0 != 127) {
            fmt.Println(v50)
            if v3 {
                if (v53 != v53) {
                    fmt.Println(v50, v4, v50)
                    switch -603 {
                    case 128:
                        v4 = ((!v3) && false)
                    default:
                        if false {
                            var v60 float64 = ((-2.25 / (0.0 / ((-2.25 * float64(v50)) + (0.0 / float64(v50))))) / (1e10 + ((3.14159 - -2.25) - (3.14159 / (0.1 - ((0.1 / 1.0) / (1e10 * 0.1)))))))
                            v53 += v53
                            _ = v60
                        }
                    }
                } else {
                    v2 = v2
                }
                for fi61 := 0; fi61 < 16; fi61++ {
                    v62 := GenS2{F0: fi61, F1: (^(v2)), F2: 351}
                    switch (fi61 - 1) {
                    case 4294967295:
                        v63 := v1[v53]
                        _ = v63
                    default:
                        v3 = v3
                    }
                    _ = fi61
                    _ = v62
                }
            }
        } else {
            v2 -= (v2 << ((4294967295) & 63))
        }
    case 32768:
        v64 := GenS8{F0: 42}
        _ = v64
    default:
        var v65 uint = 440
        _ = v65
    }
    if v4 {
        v2 += 1
        v2 = v2
        var v66 float64 = 0.0
        _ = v66
    }
    for fi67 := 0; fi67 < 1; fi67++ {
        v50 = v50
        _ = fi67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v59
}
