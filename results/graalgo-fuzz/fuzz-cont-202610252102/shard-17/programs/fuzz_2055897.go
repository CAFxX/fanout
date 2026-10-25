package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 bool; F2 int8 }
type GenS6 struct { F0 int16; F1 uint64; F2 string }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 bool; F1 uint32; F2 int8 }
func genFn0(p0 int, p1 uint64, p2 uint8) int32 {
    var v6 int64 = 1
    v7 := GenS1{F0: (p0 + p0)}
    var v8 string = "\n\t"
    var v9 int64 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 946
    var v2 uint16 = 453
    var v3 int = 1
    v4 := []string{}
    var v5 int8 = 1
    v1 = (v1 % (((v1 / ((63) | 1))) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v1 -= v2
        if (2 == 73) {
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 int16 = 1
                v4 = append(v4, ("\n\t" + "hello"))
                if true {
                    switch 128 {
                    case 127:
                        v5 -= (v5 - v5)
                    }
                    v3 = fi10
                } else {
                    v5 += v5
                    v3 -= fi11
                    var v13 int8 = 42
                    switch "hello" {
                    case "":
                        switch "\n\t" {
                        case "a":
                            v13 += (-(v13))
                            v5 += v5
                            if true {
                                if (!(true || ((v1 | v1) >= (v2 >> ((18446744073709551615) & 63))))) {
                                    if (4294967295 <= 65535) {
                                        _ = len(v4)
                                    } else {
                                        v4 = append(v4, "\n\t")
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = "\n\t"
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = ("\n\t" + "foo bar")
                                            }
                                            var v14 int32 = 492
                                            var v15 uint8 = 0
                                            var v16 float32 = 1e10
                                            _ = v14
                                            _ = v15
                                            _ = v16
                                        }
                                    }
                                }
                                fmt.Println(v3)
                            } else {
                                v2 -= 32767
                                var v17 int = (fi11 | 0)
                                v18 := GenS1{F0: (^(fi10))}
                                _ = v17
                                _ = v18
                            }
                        case "foo bar":
                            v13 += v5
                        case "":
                            v3 -= fi10
                            v5 = 0
                            switch fi10 {
                            case -733:
                                var v19 uint = 800
                                _ = v19
                            default:
                                var v20 uint32 = uint32(v1)
                                _ = v20
                            }
                            v21 := GenS1{F0: fi11}
                            _ = v21
                        }
                    case "foo bar":
                        if false {
                            var v22 uint64 = 4294967295
                            v23 := GenS1{F0: (v3 << ((255) & 63))}
                            var v24 int32 = int32(v12)
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            switch v3 {
                            case 9223372036854775807:
                                var v25 int8 = 0
                                _ = v25
                            case -9223372036854775808:
                                var v26 float64 = -2.25
                                _ = v26
                            default:
                                v13 = v5
                                if false {
                                    var v27 uint32 = 1
                                    v28 := GenS2{F0: fi10, F1: 9223372036854775807}
                                    if ((v12 / ((42) | 1)) < (v12 * v12)) {
                                        fmt.Println(v3)
                                        v4 = append(v4, "")
                                    } else {
                                        var v29 string = "\n\t"
                                        v30 := GenS3{F0: false, F1: 65535}
                                        v31 := GenS4{F0: -88}
                                        var v32 uint16 = 778
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 uint64 = 18446744073709551615
                                    if (!true) {
                                        v4 = append(v4, "foo bar")
                                    }
                                    _ = v27
                                    _ = v28
                                    _ = v33
                                } else {
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v5, fi10, v1)
                                    case "hello":
                                        v34 := GenS2{F0: (-(fi11)), F1: fi10}
                                        _ = v34
                                    default:
                                        v35 := GenS5{F0: 2147483647, F1: false, F2: (v13 << ((4294967295) & 63))}
                                        _ = v35
                                    }
                                    v3 = -2147483648
                                    var v36 int16 = (v12 * (v12 / ((v12) | 1)))
                                    v12 -= v36
                                    _ = v36
                                }
                                v4 = append(v4, "foo bar")
                            }
                            var v37 int64 = -9223372036854775808
                            _ = v37
                        }
                    case "hello":
                        _ = len(v4)
                    }
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            v3 += fi10
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi10
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        switch fi38 {
        case 4294967295:
            fmt.Println(v3, v3)
        case -2147483648:
            if (3.14159 > (3.14159 - 0.0)) {
                v4 = append(v4, ("foo bar" + ""))
                v1 = v2
            } else {
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 uint = 720
                    if true {
                        v1 += (v1 ^ v1)
                        var v41 uint16 = (v2 ^ 2)
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
                v42 := GenS5{F0: 709, F1: (0.1 == 0.1), F2: (v5 * v5)}
                _ = v42
            }
        case 236:
            var v43 uint = 255
            _ = v43
        }
        var v44 bool = (-341 != -235)
        fmt.Println(fi38, fi38, v44)
        _ = fi38
        _ = v44
    }
    v45 := GenS3{F0: false, F1: 255}
    v46 := GenS5{F0: 908, F1: false, F2: (v5 % (((v5 - 0)) | 1))}
    v47 := GenS5{F0: 0, F1: ((-2.25 + 1.5) >= 3.14159), F2: v5}
    switch (^(v3)) {
    case -9223372036854775808:
        switch 128 {
        case 2147483647:
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
    case 222:
        if true {
            v48 := GenS3{F0: (("" + ("foo bar" + "hello")) > ("foo bar" + "hello")), F1: 42}
            _ = len(v4)
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 20; fi49++ {
        v5 += v5
        _ = fi49
    }
    v5 -= 1
    v2 -= (v2 & v1)
    v2 -= (v2 & v1)
    v50 := GenS6{F0: -194, F1: 18446744073709551615, F2: "hello"}
    switch ("" + "") {
    case "":
        v1 = 360
    case "hello":
        v51 := GenS7{F0: (!((4294967295 < 954) || (1e10 != 1.5))), F1: 75}
        _ = v51
    }
    v1 -= 816
    v1 = (^(v2))
    if ((^(v5)) == 0) {
        v52 := GenS5{F0: 771, F1: true, F2: (v5 % ((42) | 1))}
        var v53 int = (^(v3))
        var v54 int64 = 42
        _ = v52
        _ = v53
        _ = v54
    } else {
        var v55 uint32 = 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi56 := 0; fi56 < 3; fi56++ {
            var v57 uint64 = 2
            v58 := GenS8{F0: (!(uint(v1) > 42)), F1: (^(v55)), F2: (v5 << ((568) & 63))}
            _ = fi56
            _ = v57
            _ = v58
        }
        _ = v55
    }
    if true {
        v59 := GenS8{F0: (!(-343 < 419)), F1: 2, F2: v5}
        var v60 float32 = 0.1
        _ = v59
        _ = v60
    } else {
        var v61 float64 = (0.1 + (float64(3.14159) + 1e10))
        var v62 string = ""
        v63 := GenS8{F0: false, F1: 2147483648, F2: v5}
        _ = v61
        _ = v62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}
