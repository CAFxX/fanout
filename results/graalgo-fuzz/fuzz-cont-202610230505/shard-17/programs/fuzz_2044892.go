package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
type GenS5 struct { F0 string; F1 uint64; F2 bool }
type GenS6 struct { F0 string; F1 uint64 }
func genFn0(p0 uint64, p1 int) string {
    var v5 string = ""
    switch p1 {
    case 32767:
        if false {
            var v6 float32 = 3.14159
            v7 := GenS1{F0: 47}
            fmt.Println(p1, p0)
            _ = v6
            _ = v7
        }
    case 9223372036854775807:
        fmt.Println(v5, v5)
    case 4294967295:
        var v8 uint64 = p0
        _ = v8
    }
    v9 := GenS1{F0: 42}
    if false {
        p0 -= p0
    } else {
        fmt.Println(p0)
        p1 = p1
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return v5
}
func genFn1(p0 uint64) uint {
    if (42 < 0) {
        v10 := []int{}
        v11 := GenS1{F0: 42}
        _ = v10
        _ = v11
    } else {
        switch "\n\t" {
        case "foo bar":
            var v12 uint16 = 162
            _ = v12
        }
        var v13 float32 = 0.1
        if ((((!(-53 >= 1)) && (1 < 9223372036854775807)) && false) || (v13 == float32(((3.14159 * 3.14159) - (((-2.25 * 1e10) / 1.5) / 0.1))))) {
            p0 -= p0
        } else {
            v13 = (v13 + float32(1e10))
            v14 := GenS2{F0: 2147483647, F1: 630, F2: -2147483648}
            v15 := []int{2147483647}
            _ = v14
            _ = v15
        }
        _ = v13
    }
    switch 0 {
    case -334:
        var v16 int8 = 1
        _ = v16
    case 32768:
        switch "foo bar" {
        case "foo bar":
            v17 := []int{909}
            _ = v17
        default:
            if false {
                if (true || (true && false)) {
                    switch 128 {
                    case -862:
                        v18 := []int{42, 128, -9223372036854775808, 9223372036854775807}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v20 := []int{(fi19 * (^(fi19)))}
                            _ = fi19
                            _ = v20
                        }
                    }
                    var v21 int = 561
                    _ = v21
                }
                var v22 int16 = 0
                v23 := []int{-2147483648, 9223372036854775807, -2147483648, -931}
                _ = v22
                _ = v23
            }
        }
    }
    v24 := GenS3{F0: 42, F1: p0, F2: 55}
    var v25 string = ""
    var v26 float32 = 1.5
    _ = p0
    _ = v24
    _ = v25
    _ = v26
    return 935
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]bool)
    var v3 bool = true
    v4 := make(map[string]string)
    if (!v3) {
        v4["a"] = (((("" + ("a" + "a")) + "a") + "foo bar") + "hello")
    }
    v27 := GenS4{F0: false, F1: 524, F2: 42}
    v28 := []int{912}
    _ = len(v4)
    v3 = (255 != 65535)
    switch 2147483647 {
    case 4294967295:
        v2["foo bar"] = (0 <= 110)
    case 2147483647:
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := GenS3{F0: 33, F1: 18446744073709551615, F2: 0}
            fmt.Println(fi29, v1, fi29)
            _ = fi29
            _ = v30
        }
    }
    var v31 uint = 400
    var v32 int = 32767
    switch ("a" + "a") {
    case "foo bar":
        v33 := GenS3{F0: 42, F1: 1, F2: 104}
        _ = v33
    case "":
        for fi34 := 0; fi34 < 10; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                v3 = false
                _ = fi35
            }
            _ = fi34
        }
    default:
        var v36 uint64 = 42
        _ = v36
    }
    if (-780 > 0) {
        v37 := GenS1{F0: -1}
        for fi38 := 0; fi38 < 10; fi38++ {
            if false {
                fmt.Println(fi38, v3)
                fmt.Println(v32)
            } else {
                var v39 uint = (^(v31))
                if v3 {
                    v32 -= (v32 * (^(fi38)))
                } else {
                    var v40 int64 = 0
                    for fi41 := 0; fi41 < 14; fi41++ {
                        var v42 uint64 = 0
                        var v43 uint = v31
                        var v44 float64 = ((float64(1e10) / 0.1) + float64(1.5))
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    _ = v40
                }
                _ = v39
            }
            fmt.Println(v31, v32)
            _ = fi38
        }
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        _ = v37
    } else {
        v45 := GenS2{F0: 42, F1: 332, F2: int32(v32)}
        if true {
            v31 += v31
            if (4294967295 >= 42) {
                v32 -= v32
            } else {
                if v3 {
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    for fi46 := 0; fi46 < 14; fi46++ {
                        fmt.Println(v31, v1)
                        _ = fi46
                    }
                    fmt.Println(v31)
                } else {
                    if (!((1 > 32767) || false)) {
                        v47 := GenS4{F0: (0 >= 42), F1: -660, F2: 42}
                        _ = v47
                    } else {
                        switch "a" {
                        case "a":
                            if v3 {
                                v48 := GenS1{F0: 127}
                                var v49 string = ("" + (("\n\t" + "foo bar") + "\n\t"))
                                v31 = v31
                                _ = v48
                                _ = v49
                            }
                        default:
                            fmt.Println(v3, v32, v3)
                        }
                        v3 = (1 > 2)
                        if (-354 < 0) {
                            var v50 int8 = 0
                            if false {
                                if false {
                                    if len(v28) > 0 {
                                        v28[0] = (v32 >> (((v31 % ((genFn1(1)) | 1))) & 63))
                                    }
                                    v31 = 42
                                    if (v3 || v3) {
                                        v32 -= v32
                                        switch "a" {
                                        case "hello":
                                            var v51 uint8 = 127
                                            _ = v51
                                        case "\n\t":
                                            v52 := GenS5{F0: "a", F1: 319, F2: (1 <= -409)}
                                            _ = v52
                                        }
                                        fmt.Println(v50)
                                    } else {
                                        if v3 {
                                            var v53 float64 = float64((v1 + (-2.25 + (v1 + v1))))
                                            _ = v53
                                        } else {
                                            switch (((("a" + ("" + "foo bar")) + "foo bar") + "\n\t") + ("foo bar" + "a")) {
                                            case "":
                                                v54 := v4[(("foo bar" + "a") + (("" + ("hello" + ("hello" + "foo bar"))) + "a"))]
                                                _ = v54
                                            default:
                                                fmt.Println(v50, v32)
                                            }
                                        }
                                        switch -219 {
                                        case 32767:
                                            v3 = (3.14159 < (-2.25 * 1e10))
                                        }
                                        fmt.Println(v32)
                                    }
                                }
                                v50 += (v50 + v50)
                                v55 := GenS6{F0: ("\n\t" + "hello"), F1: 2}
                                _ = v55
                            } else {
                                v32 = v32
                            }
                            var v56 uint32 = 65535
                            _ = v50
                            _ = v56
                        }
                    }
                    switch "hello" {
                    case "a":
                        v31 = genFn1(901)
                    }
                }
                var v57 bool = ((!false) && v3)
                v31 = 255
                _ = v57
            }
            v28 = append(v28, (v32 - (-(v32))))
        } else {
            for fi58 := 0; fi58 < 6; fi58++ {
                if (!v3) {
                    v31 -= (v31 - v31)
                }
                v59 := GenS2{F0: -308, F1: 632, F2: 1}
                _ = fi58
                _ = v59
            }
        }
        v28 = append(v28, (v32 ^ v32))
        _ = v45
    }
    v31 = v31
    _ = len(v28)
    var v60 float32 = (v1 - ((-2.25 * (float32(3.14159) * v1)) * v1))
    var v61 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v60
    _ = v61
}
