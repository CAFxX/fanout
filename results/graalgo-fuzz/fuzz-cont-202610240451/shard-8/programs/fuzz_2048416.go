package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int }

func main() {
    v1 := []uint32{}
    var v2 int8 = -1
    v3 := make(map[uint32]bool)
    var v4 uint32 = 4294967295
    if true {
        switch 127 {
        case -94:
            fmt.Println(v4, v2)
        case 2147483647:
            v5 := GenS1{F0: 1, F1: 255, F2: false}
            _ = v5
        case 127:
            v6 := GenS1{F0: -727, F1: 68, F2: (-643 >= 42)}
            _ = v6
        }
        v2 += (^(v2))
    }
    switch "\n\t" {
    case "a":
        if ((v4 >> ((42) & 63)) <= v4) {
            var v7 int64 = 9223372036854775807
            switch 886 {
            case -9223372036854775808:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 4294967295:
                for fi8 := 0; fi8 < 4; fi8++ {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 uint64 = 164
                        var v11 float32 = -2.25
                        delete(v3, 65535)
                        var v12 int16 = 32767
                        switch ("\n\t" + "a") {
                        case "hello":
                            v11 = 0.0
                        case "":
                            var v13 uint8 = 39
                            _ = v13
                        }
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    if false {
                        var v14 uint64 = 282
                        if (true && true) {
                            v4 = (v4 ^ v4)
                        }
                        _ = len(v1)
                        _ = v14
                    } else {
                        v4 = v4
                        v15 := GenS1{F0: 1000, F1: 0, F2: true}
                        _ = v15
                    }
                    _ = fi8
                }
            case -2147483648:
                var v16 uint = 659
                _ = v16
            default:
                fmt.Println(v2)
            }
            _ = v7
        } else {
            var v17 string = "foo bar"
            var v18 float64 = float64(0.1)
            var v19 uint = 1
            _ = v17
            _ = v18
            _ = v19
        }
    case "\n\t":
        v20 := GenS1{F0: -32768, F1: 1, F2: ((-2147483648 > 1) && (!((("a" >= (("hello" + ("foo bar" + "\n\t")) + ("hello" + ("\n\t" + "hello")))) || true) || true)))}
        _ = v20
    default:
        switch 127 {
        case 4294967295:
            v21 := GenS1{F0: -719, F1: 218, F2: (65535 >= 18446744073709551615)}
            _ = v21
        case 905:
            for fi22 := 0; fi22 < 17; fi22++ {
                fmt.Println(fi22, fi22, v4)
                v23 := GenS2{F0: 18446744073709551615}
                _ = len(v3)
                _ = fi22
                _ = v23
            }
        default:
            v4 += 65535
        }
    }
    v24 := GenS2{F0: 635}
    var v25 int32 = -2147483648
    v26 := GenS3{F0: v4, F1: true}
    v27 := GenS1{F0: 0, F1: 138, F2: true}
    v28 := GenS3{F0: v4, F1: (false && true)}
    fmt.Println(v4)
    v25 -= -1
    v29 := v3[v4]
    var v30 bool = (4294967295 != 473)
    if (0 <= 9223372036854775807) {
        if v29 {
            v30 = (false || (true || (42 == -9223372036854775808)))
        }
        v25 = (v25 * (v25 - (v25 + (-(v25)))))
    } else {
        var v31 uint = 18446744073709551615
        var v32 string = ("" + ("foo bar" + ""))
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        if v29 {
            fmt.Println(fi33)
            var v34 int16 = int16(v2)
            _ = v34
        } else {
            for fi35 := 0; fi35 < 9; fi35++ {
                switch "\n\t" {
                case "hello":
                    var v36 int = fi33
                    _ = v36
                case "a":
                    if false {
                        if v30 {
                            fmt.Println(v30)
                        }
                        var v37 int16 = 0
                        var v38 uint = 65535
                        _ = v37
                        _ = v38
                    } else {
                        var v39 uint32 = (v4 - v4)
                        _ = v39
                    }
                case "\n\t":
                    var v40 uint8 = 1
                    _ = v40
                default:
                    var v41 uint = 980
                    _ = v41
                }
                switch 42 {
                case 127:
                    var v42 int = fi33
                    _ = v42
                case -74:
                    v3[v4] = v29
                case 32768:
                    var v43 uint = 257
                    _ = v43
                default:
                    switch fi35 {
                    case -1:
                        _ = len(v3)
                    case 852:
                        v25 = (^(v25))
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v44 := GenS1{F0: 1000, F1: 201, F2: (int64(v4) == -9223372036854775808)}
                    _ = v44
                }
                v45 := GenS3{F0: v4, F1: true}
                if (255 <= 42) {
                    var v46 float32 = (1e10 + -2.25)
                    v2 = (v2 | (v2 % (((v2 << ((464) & 63))) | 1)))
                    if v30 {
                        var v47 string = (("hello" + "") + ("\n\t" + "\n\t"))
                        _ = v47
                    } else {
                        if (v46 > float32(0.1)) {
                            _ = len(v1)
                        }
                        v29 = (((!v30) && ((v30 && v29) || v30)) || ((100 < (v2 | v2)) && v29))
                        fmt.Println(v46, fi33, v2)
                    }
                    _ = v46
                }
                _ = fi35
                _ = v45
            }
        }
        var v48 uint = 0
        _ = fi33
        _ = v48
    }
    v49 := GenS4{F0: 1}
    if len(v1) > 0 {
        v1[0] = v4
    }
    v3[(v4 >> ((2) & 63))] = true
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v49
}
