package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 uint32, p1 int64, p2 uint16) uint16 {
    fmt.Println(p2)
    v4 := GenS1{F0: 125}
    for fi5 := 0; fi5 < 4; fi5++ {
        if true {
            v6 := GenS1{F0: uint8(p2)}
            if (fi5 <= 32767) {
                var v7 uint32 = p0
                var v8 float64 = float64(-2.25)
                if false {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        switch (^(fi5)) {
                        case 4294967295:
                            var v10 uint64 = 909
                            _ = v10
                        }
                        switch "foo bar" {
                        case "foo bar":
                            p2 += 65535
                        case "\n\t":
                            var v11 int8 = -1
                            _ = v11
                        }
                        if (!(1 == 0)) {
                            p2 = (^(p2))
                        } else {
                            v12 := GenS2{F0: (^(fi5)), F1: 4294967295, F2: false}
                            _ = v12
                        }
                        _ = fi9
                    }
                    var v13 float32 = 0.1
                    _ = v13
                }
                _ = v7
                _ = v8
            } else {
                switch ("foo bar" + "hello") {
                case "\n\t":
                    var v14 int = (fi5 | fi5)
                    _ = v14
                }
                v15 := []int{fi5, fi5, (fi5 >> ((65535) & 63))}
                _ = v15
            }
            p2 = p2
            _ = v6
        }
        var v16 float64 = float64(((-2.25 + 1e10) - 1e10))
        var v17 int64 = -404
        _ = fi5
        _ = v16
        _ = v17
    }
    v18 := []int{127, -1, -9223372036854775808}
    var v19 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v4
    return p2
}
func genFn1(p0 bool, p1 int8) uint64 {
    switch ("a" + "hello") {
    case "foo bar":
        var v20 uint32 = 65535
        _ = v20
    case "a":
        v21 := []int{}
        _ = v21
    case "\n\t":
        switch 32768 {
        case -9223372036854775808:
            var v22 int16 = 143
            _ = v22
        default:
            v23 := GenS3{F0: 42, F1: -718}
            _ = v23
        }
    }
    v24 := []int{-110}
    if true {
        fmt.Println(p0, p1)
        var v25 int = 32768
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        _ = v25
    }
    var v26 uint8 = 127
    _ = p0
    _ = p1
    _ = v24
    _ = v26
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]uint32)
    var v3 bool = true
    for fi27 := 0; fi27 < 2; fi27++ {
        if v3 {
            v1 -= (v1 / ((v1) | 1))
            v28 := GenS1{F0: 198}
            _ = v28
        }
        if (!(!v3)) {
            if ((0 == 128) && (0.0 >= 3.14159)) {
                _ = len(v2)
                var v29 float32 = 0.0
                if (fi27 != fi27) {
                    v29 = (float32(-2.25) - (3.14159 - 1e10))
                    v2["hello"] = (v1 & (^(v1)))
                    fmt.Println(v1)
                }
                _ = v29
            } else {
                if (337 != 2) {
                    var v30 uint16 = 255
                    for fi31 := 0; fi31 < 11; fi31++ {
                        var v32 uint8 = 205
                        _ = fi31
                        _ = v32
                    }
                    v30 = uint16(v1)
                    _ = v30
                }
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                if (v3 && (!(-1 >= int64(v1)))) {
                    if (0 > 2147483647) {
                        var v34 uint64 = 18446744073709551615
                        v34 -= v34
                        switch "\n\t" {
                        case "foo bar":
                            if (!v3) {
                                if v3 {
                                    v1 = (v1 ^ v1)
                                }
                                v35 := GenS2{F0: (fi27 + 32768), F1: (v34 & v34), F2: (!v3)}
                                _ = v35
                            }
                        case "":
                            var v36 int64 = 9223372036854775807
                            _ = v36
                        case "a":
                            v1 -= v1
                        }
                        _ = v34
                    } else {
                        v1 = (^(v1))
                        v1 = 42
                        var v37 uint16 = 42
                        _ = v37
                    }
                }
                _ = fi33
            }
            fmt.Println(v1)
        }
        _ = fi27
    }
    v1 += v1
    if v3 {
        if ((v3 && v3) || v3) {
            for fi38 := 0; fi38 < 18; fi38++ {
                if (127 > 0) {
                    v39 := v2["foo bar"]
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 19; fi40++ {
                        fmt.Println(fi38, v3, v3)
                        var v41 uint = 561
                        if (128 > 246) {
                            fmt.Println(fi40)
                            fmt.Println(v3)
                            switch "foo bar" {
                            case "\n\t":
                                v41 += (v41 | 572)
                            default:
                                if (((15 != 127) || v3) && (0.1 < (0.1 / (float64((-2.25 / 1.0)) / 1e10)))) {
                                    if true {
                                        if (!(-714 >= -9223372036854775808)) {
                                            v1 = 42
                                            var v42 int16 = -32768
                                            _ = v42
                                        }
                                        fmt.Println(fi38, v41)
                                    } else {
                                        var v43 string = "\n\t"
                                        var v44 int64 = -1
                                        var v45 int32 = int32(v41)
                                        v46 := GenS4{F0: 42}
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                    fmt.Println(v1)
                                } else {
                                    v3 = ((^(fi38)) == (fi40 >> ((v41) & 63)))
                                }
                            }
                        } else {
                            if (42 < 134) {
                                var v47 uint16 = genFn0(v1, 689, 42)
                                var v48 int = fi40
                                v49 := []int{(fi38 | fi40)}
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            var v50 float64 = (3.14159 + ((float64(3.14159) + (1.5 / 0.1)) + 1.5))
                            var v51 int16 = 32767
                            _ = v50
                            _ = v51
                        }
                        _ = fi40
                        _ = v41
                    }
                    for fi52 := 0; fi52 < 10; fi52++ {
                        v53 := []int{fi52, fi52}
                        if ("" < "hello") {
                            if v3 {
                                switch "hello" {
                                case "foo bar":
                                    switch (fi38 % ((32767) | 1)) {
                                    case 127:
                                        var v54 uint8 = 42
                                        _ = v54
                                    }
                                    v55 := GenS1{F0: 42}
                                    fmt.Println(v1, fi52)
                                    _ = v55
                                default:
                                    v56 := GenS4{F0: 42}
                                    _ = v56
                                }
                                v57 := GenS5{F0: (((v3 && v3) || v3) && v3)}
                                _ = v57
                            } else {
                                if false {
                                    switch "a" {
                                    case "hello":
                                        var v58 int16 = 0
                                        _ = v58
                                    default:
                                        delete(v2, ("hello" + "hello"))
                                    }
                                } else {
                                    var v59 uint32 = (v1 + (v1 / (((v1 + 0)) | 1)))
                                    _ = v59
                                }
                            }
                        } else {
                            fmt.Println(fi52)
                            if true {
                                v60 := GenS6{F0: 1, F1: ""}
                                delete(v2, ("hello" + "hello"))
                                v61 := GenS5{F0: (558 <= -32768)}
                                _ = v60
                                _ = v61
                            }
                            v1 += v1
                        }
                        var v62 int64 = -1
                        _ = fi52
                        _ = v53
                        _ = v62
                    }
                }
                _ = fi38
            }
        } else {
            v63 := GenS4{F0: 255}
            _ = v63
        }
    }
    v1 += 0
    for fi64 := 0; fi64 < 20; fi64++ {
        v3 = false
        _ = fi64
    }
    _ = len(v2)
    for fi65 := 0; fi65 < 10; fi65++ {
        switch "hello" {
        case "foo bar":
            var v66 int = 2147483647
            _ = v66
        }
        _ = fi65
    }
    v67 := GenS7{F0: 2147483647, F1: false, F2: 42}
    switch "" {
    case "\n\t":
        v2[""] = v1
    default:
        fmt.Println(v1)
    }
    v68 := v2[""]
    var v69 int8 = -128
    fmt.Println(v1, v68, v69)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v67
    _ = v68
    _ = v69
}
