package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint8; F1 bool; F2 int16 }
type GenS5 struct { F0 string; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 string; F1 uint8; F2 int16 }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 uint32 = 0
    var v2 int16 = -32768
    v3 := make(map[int]uint32)
    var v4 int64 = -1
    var v5 uint16 = 0
    v6 := []int{}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int32 = -1
        var v9 uint = 519
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v5)
    v4 = v4
    v1 += (^(v1))
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: -48}
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 += 2
            v13 := GenS2{F0: 42, F1: int64(v2)}
            var v14 uint = 285
            var v15 string = ""
            v16 := GenS1{F0: 22}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
        }
        v1 += 2147483647
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v5)
    if true {
        var v17 int = 32767
        _ = v17
    }
    var v18 int32 = 1
    if (9223372036854775807 == -9223372036854775808) {
        v5 = v5
    } else {
        if (249 < 255) {
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(v5)
                var v20 float64 = 0.0
                _ = fi19
                _ = v20
            }
            fmt.Println(v4, v4, v5)
            for fi21 := 0; fi21 < 5; fi21++ {
                delete(v3, fi21)
                fmt.Println(v2, v2, v5)
                v22 := GenS2{F0: 61, F1: (v4 & -9223372036854775808)}
                _ = fi21
                _ = v22
            }
        }
        switch (("hello" + "a") + "foo bar") {
        case "\n\t":
            switch 4294967295 {
            case 128:
                v2 += (v2 + (v2 - -32768))
            default:
                if (1 < 42) {
                    var v23 int8 = 1
                    v24 := GenS3{F0: 973, F1: v1}
                    var v25 uint32 = v1
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    v4 = (v4 + -9223372036854775808)
                    if true {
                        for fi26 := 0; fi26 < 13; fi26++ {
                            for fi27 := 0; fi27 < 6; fi27++ {
                                _ = len(v3)
                                _ = len(v3)
                                _ = fi27
                            }
                            switch ("hello" + "hello") {
                            case "a":
                                var v28 uint64 = 0
                                _ = v28
                            case "\n\t":
                                fmt.Println(fi26, v2, v18)
                            case "hello":
                                fmt.Println(v1, fi26)
                            default:
                                v29 := GenS4{F0: 255, F1: (((!true) && (!((int32(v5) >= (^(v18))) || true))) || (true && false)), F2: (v2 % (((v2 ^ v2)) | 1))}
                                _ = v29
                            }
                            var v30 uint64 = 355
                            _ = fi26
                            _ = v30
                        }
                        v31 := GenS2{F0: 1, F1: 0}
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v31
                    } else {
                        if (42 >= v4) {
                            if false {
                                v1 = 0
                            } else {
                                switch "hello" {
                                case "a":
                                    v4 = (v4 - (v4 * v4))
                                case "\n\t":
                                    for fi32 := 0; fi32 < 4; fi32++ {
                                        for fi33 := 0; fi33 < 13; fi33++ {
                                            v18 += 42
                                            v1 += v1
                                            var v34 uint8 = uint8(v1)
                                            if false {
                                                switch fi33 {
                                                case -9223372036854775808:
                                                    if len(v6) > 0 {
                                                        v6[0] = (fi33 - fi33)
                                                    }
                                                case 0:
                                                    switch (fi32 / ((32768) | 1)) {
                                                    case 2147483647:
                                                        v35 := GenS5{F0: "\n\t", F1: 986, F2: v5}
                                                        var v36 int32 = (v18 ^ v18)
                                                        _ = v35
                                                        _ = v36
                                                    case 4294967295:
                                                        v37 := GenS1{F0: 120}
                                                        _ = v37
                                                    case 32767:
                                                        v18 = v18
                                                    default:
                                                        fmt.Println(v34)
                                                    }
                                                default:
                                                    fmt.Println(v5, v5)
                                                }
                                            } else {
                                                fmt.Println(fi33, v2)
                                            }
                                            _ = fi33
                                            _ = v34
                                        }
                                        _ = fi32
                                    }
                                case "hello":
                                    v1 += 2147483648
                                }
                            }
                            for fi38 := 0; fi38 < 14; fi38++ {
                                delete(v3, 42)
                                v5 = (^(v5))
                                _ = fi38
                            }
                        }
                    }
                }
            }
        case "":
            if true {
                v39 := GenS4{F0: 127, F1: (v1 == 2147483647), F2: v2}
                v40 := GenS6{F0: v2, F1: "a"}
                switch 128 {
                case 32768:
                    if (4294967295 < 1) {
                        v41 := GenS7{F0: "a", F1: 127, F2: -1}
                        var v42 int = 1
                        _ = len(v3)
                        _ = v41
                        _ = v42
                    } else {
                        var v43 bool = true
                        v18 -= 1
                        _ = v43
                    }
                case -116:
                    v6 = append(v6, 541)
                default:
                    _ = len(v3)
                }
                _ = v39
                _ = v40
            } else {
                if (2 < 0) {
                    if false {
                        for fi44 := 0; fi44 < 3; fi44++ {
                            v2 -= (v2 * v2)
                            _ = fi44
                        }
                        fmt.Println(v18, v5)
                    }
                } else {
                    switch -970 {
                    case 475:
                        var v45 float64 = 0.1
                        _ = v45
                    default:
                        for fi46 := 0; fi46 < 15; fi46++ {
                            v5 = 0
                            v4 = v4
                            _ = fi46
                        }
                    }
                    fmt.Println(v4, v1)
                    if true {
                        var v47 int16 = v2
                        switch "" {
                        case "":
                            for fi48 := 0; fi48 < 9; fi48++ {
                                var v49 bool = (!false)
                                v50 := GenS2{F0: 42, F1: v4}
                                if v49 {
                                    v51 := GenS8{F0: 2}
                                    var v52 int16 = (v47 + 323)
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 float64 = (-2.25 / ((float64(1e10) + 0.1) / 3.14159))
                                    _ = v53
                                }
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                        case "foo bar":
                            v47 -= v2
                        case "\n\t":
                            var v54 bool = true
                            _ = v54
                        }
                        _ = v47
                    } else {
                        v55 := GenS6{F0: (v2 / ((v2) | 1)), F1: ""}
                        v5 = (^(v5))
                        _ = v55
                    }
                }
            }
        }
    }
    var v56 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
