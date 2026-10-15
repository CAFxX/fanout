package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint }
func genFn0(p0 uint32, p1 uint32, p2 uint) uint64 {
    switch 42 {
    case -9223372036854775808:
        fmt.Println(p2)
    default:
        fmt.Println(p1, p0, p2)
    }
    p2 = p2
    v4 := GenS1{F0: (true || (!(!(110 < -1)))), F1: 9223372036854775807}
    v5 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 158
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483647}
    var v3 int8 = -12
    var v6 uint16 = 596
    v7 := GenS1{F0: false, F1: 9223372036854775807}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 128 {
    case 0:
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int = (^(fi8))
            _ = fi8
            _ = v9
        }
    case 262:
        var v10 int32 = 766
        _ = v10
    }
    var v11 int32 = -909
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 14; fi13++ {
        delete(v1, fi13)
        v14 := v1[fi13]
        var v15 uint16 = 151
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint32 = 366
    switch (v12 + ((v12 + "foo bar") + "\n\t")) {
    case "hello":
        var v17 int32 = 396
        _ = v17
    case "":
        v11 += int32(v3)
    default:
        if (false || false) {
            var v18 string = "hello"
            v12 = (v12 + "")
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 uint8 = 142
                for fi21 := 0; fi21 < 13; fi21++ {
                    if ((^(v11)) != v11) {
                        v3 = v3
                        fmt.Println(fi21, v16, v18)
                        delete(v1, fi21)
                    }
                    _ = fi21
                }
                if (!(!(false || true))) {
                    if false {
                        var v22 int = -308
                        for fi23 := 0; fi23 < 3; fi23++ {
                            if true {
                                v24 := GenS2{F0: 483}
                                var v25 uint = 4294967295
                                v25 = v25
                                _ = v24
                                _ = v25
                            } else {
                                v16 -= v16
                                v26 := GenS2{F0: (v16 * (v16 >> ((780) & 63)))}
                                fmt.Println(v18, fi19)
                                v27 := GenS2{F0: v16}
                                _ = v26
                                _ = v27
                            }
                            if (((^(v16)) > v16) && (!((false || false) || false))) {
                                var v28 int16 = 1
                                v16 = v16
                                _ = v28
                            } else {
                                v29 := v1[fi19]
                                switch v12 {
                                case "":
                                    fmt.Println(fi23, fi19)
                                case "foo bar":
                                    var v30 int64 = 42
                                    _ = v30
                                case "a":
                                    if ((858 >= 42) && v29) {
                                        v2 = append(v2, 65535)
                                        fmt.Println(v3, v12)
                                        v16 = v16
                                    } else {
                                        var v31 int64 = -9223372036854775808
                                        _ = v31
                                    }
                                default:
                                    if v29 {
                                        fmt.Println(v20)
                                    }
                                }
                                var v32 bool = (v18 > (v18 + (v18 + "a")))
                                v33 := GenS1{F0: (!v29), F1: -1}
                                var v34 string = "hello"
                                _ = v29
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v11, v18, v20)
                            _ = fi23
                        }
                        _ = v22
                    } else {
                        if ((!false) && (v20 <= 127)) {
                            switch "hello" {
                            case "hello":
                                v3 -= v3
                            default:
                                var v35 int16 = 42
                                _ = v35
                            }
                            v20 -= (v20 * 2)
                            for fi36 := 0; fi36 < 15; fi36++ {
                                if true {
                                    var v37 float32 = 0.0
                                    _ = v37
                                } else {
                                    v38 := GenS2{F0: (v16 + (v16 ^ 373))}
                                    var v39 string = ("hello" + v12)
                                    _ = v38
                                    _ = v39
                                }
                                _ = fi36
                            }
                        }
                        v40 := GenS1{F0: false, F1: -1}
                        fmt.Println(v20, v12)
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 14; fi41++ {
                        fmt.Println(fi19, v11)
                        _ = fi41
                    }
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            if true {
                var v42 int8 = (v3 / (((v3 - v3)) | 1))
                _ = v42
            }
        }
    }
    v12 = (((v12 + ("" + v12)) + "hello") + v12)
    var v43 float32 = (0.1 - 1e10)
    for fi44 := 0; fi44 < 18; fi44++ {
        v45 := GenS1{F0: false, F1: int64(v16)}
        v46 := GenS3{F0: -613, F1: 894}
        _ = fi44
        _ = v45
        _ = v46
    }
    v47 := v1[-340]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v43
    _ = v47
    _ = v6
    _ = v7
}
