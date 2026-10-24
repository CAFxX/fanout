package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0() int {
    var v4 int64 = 649
    v5 := GenS1{F0: 0, F1: 66, F2: 788}
    switch -576 {
    case 2147483647:
        var v6 uint64 = 2
        _ = v6
    case 121:
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 65535, F1: 459, F2: 2}
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 float64 = -2.25
                v11 := GenS1{F0: 1, F1: 705, F2: 716}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
            _ = v8
        }
    default:
        var v12 uint = 18446744073709551615
        _ = v12
    }
    var v13 uint64 = 0
    var v14 float32 = 0.0
    _ = v13
    _ = v14
    _ = v4
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 string, p1 uint16, p2 int16) uint16 {
    v15 := GenS2{F0: -122}
    if false {
        var v16 float64 = 1e10
        p2 = p2
        _ = v16
    } else {
        v17 := []int{genFn0()}
        p0 = p0
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p1 & p1)
}

func main() {
    v1 := []uint32{42, 65535}
    v2 := []uint32{2}
    var v3 uint16 = 1
    v3 += 32767
    fmt.Println(v3)
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(v3, fi18)
        _ = fi18
    }
    var v19 uint8 = 0
    if ((1 != 100) && true) {
        if ((v3 | genFn1(("" + ("\n\t" + "\n\t")), v3, -32768)) != v3) {
            v19 = (v19 % (((v19 | 219)) | 1))
        } else {
            switch genFn0() {
            case 648:
                var v20 int32 = 42
                _ = v20
            default:
                switch ("hello" + ((("" + "") + "\n\t") + ("\n\t" + ("\n\t" + (("foo bar" + "hello") + "hello"))))) {
                case "":
                    switch -927 {
                    case 127:
                        _ = len(v2)
                    case 4294967295:
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch -420 {
                            case 32767:
                                v3 += (v3 - (v3 | uint16(v19)))
                            default:
                                if false {
                                    var v21 int8 = 1
                                    v21 += v21
                                    var v22 uint = 365
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 int64 = -1
                                    _ = v23
                                }
                            }
                            var v24 string = ("" + ("" + "foo bar"))
                            _ = v24
                        }
                    }
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v25 := GenS1{F0: v3, F1: 4294967295, F2: 65535}
        _ = v25
    } else {
        if true {
            v26 := GenS1{F0: v3, F1: 42, F2: 88}
            for fi27 := 0; fi27 < 15; fi27++ {
                v3 = genFn1("hello", v3, -119)
                _ = fi27
            }
            fmt.Println(v19, v19)
            _ = v26
        }
        var v28 float32 = 1.5
        v29 := GenS1{F0: 2, F1: 42, F2: 2}
        _ = v28
        _ = v29
    }
    switch "\n\t" {
    case "":
        var v30 uint64 = 4294967295
        _ = v30
    case "a":
        fmt.Println(v3)
    case "\n\t":
        v19 -= 1
    default:
        v3 = (v3 + (v3 << ((440) & 63)))
    }
    v2 = append(v2, 301)
    var v31 uint32 = 688
    fmt.Println(v31, v3)
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 uint16 = genFn1((((("a" + "foo bar") + "foo bar") + ("\n\t" + "\n\t")) + ("a" + ("a" + ""))), (v3 >> ((0) & 63)), 949)
        for fi34 := 0; fi34 < 7; fi34++ {
            switch ("foo bar" + "foo bar") {
            case "a":
                if len(v2) > 0 {
                    v2[0] = 2
                }
            default:
                switch 127 {
                case 224:
                    v35 := GenS3{F0: v31}
                    _ = v35
                default:
                    if false {
                        v3 -= (v3 * genFn1("hello", (v33 & 65535), -113))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v31 += (v31 - (v31 & v31))
                        var v36 string = "hello"
                        v36 = ""
                        _ = v36
                    }
                }
            }
            v37 := GenS1{F0: 0, F1: 42, F2: v31}
            _ = fi34
            _ = v37
        }
        _ = fi32
        _ = v33
    }
    switch -2147483648 {
    case -9223372036854775808:
        if ("hello" >= "a") {
            for fi38 := 0; fi38 < 6; fi38++ {
                var v39 int32 = -1
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v40 uint16 = 42
                _ = fi38
                _ = v39
                _ = v40
            }
            v31 = v31
            v3 = (v3 | 42)
        } else {
            for fi41 := 0; fi41 < 4; fi41++ {
                if len(v1) > 0 {
                    v1[0] = (v31 >> ((65535) & 63))
                }
                for fi42 := 0; fi42 < 12; fi42++ {
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    v19 = v19
                    if (974 != fi41) {
                        if false {
                            v43 := GenS3{F0: v31}
                            v44 := GenS3{F0: v31}
                            switch (fi42 - (fi41 & -2147483648)) {
                            case 0:
                                v3 = 2
                            }
                            if true {
                                if ((fi42 ^ fi41) < (fi41 / (((fi42 >> ((1) & 63))) | 1))) {
                                    switch (fi42 + fi42) {
                                    case 1:
                                        v45 := GenS3{F0: uint32(v3)}
                                        _ = v45
                                    case -2147483648:
                                        v31 -= (v31 & 65535)
                                    }
                                    var v46 string = "foo bar"
                                    _ = v46
                                } else {
                                    _ = len(v1)
                                    v47 := GenS4{F0: 9223372036854775807, F1: 32767}
                                    switch genFn0() {
                                    case -2147483648:
                                        fmt.Println(v3, v19, v19)
                                    case 32768:
                                        fmt.Println(v19, fi41)
                                    }
                                    _ = v47
                                }
                                fmt.Println(v31, fi41, v3)
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (v31 ^ 2147483647)
                                }
                                _ = len(v1)
                            }
                            _ = v43
                            _ = v44
                        }
                        if true {
                            v3 += (v3 | (^(v3)))
                        } else {
                            v48 := GenS4{F0: 0, F1: 1000}
                            fmt.Println(v19)
                            v19 = v19
                            _ = v48
                        }
                    }
                    fmt.Println(v3, fi41, v3)
                    v49 := GenS1{F0: (v3 ^ (v3 / (((v3 + (^(v3)))) | 1))), F1: 255, F2: (v31 * v31)}
                    if true {
                        var v50 uint64 = 0
                        _ = v50
                    }
                    _ = fi42
                    _ = v49
                }
                _ = fi41
            }
        }
    case -1:
        for fi51 := 0; fi51 < 2; fi51++ {
            var v52 bool = (-1 >= -826)
            _ = fi51
            _ = v52
        }
    default:
        fmt.Println(v19, v3, v31)
    }
    v1 = append(v1, 2)
    v19 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
}
