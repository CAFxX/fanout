package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0(p0 uint8, p1 uint8) uint32 {
    if (540 <= 255) {
        p0 -= p0
    }
    var v4 float32 = 3.14159
    if ((715 >= 2147483648) || (94 > 544)) {
        v5 := GenS1{F0: 65535}
        v4 = v4
        v4 -= ((v4 + v4) - float32(0.1))
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 323
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", ("\n\t" + "hello")}
    var v3 int = 753
    _ = len(v1)
    var v6 uint = 2
    _ = len(v2)
    var v7 uint32 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(fi8, v3)
        v9 := GenS2{F0: ""}
        v10 := GenS2{F0: (("a" + "a") + ("foo bar" + "hello"))}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(fi11, v3)
        var v12 float32 = -2.25
        fmt.Println(v12, v6)
        _ = fi11
        _ = v12
    }
    if true {
        if len(v2) > 0 {
            v2[0] = ("a" + (("a" + ("a" + ("hello" + "hello"))) + ""))
        }
        v6 += v6
    } else {
        v3 += 2147483647
        if false {
            v13 := v1[(v3 - v3)]
            switch 9223372036854775807 {
            case 32767:
                var v14 int64 = 42
                _ = v14
            case 0:
                if v13 {
                    fmt.Println(v6, v3)
                    v6 = (v6 * v6)
                    v3 = v3
                }
            }
            if true {
                if true {
                    fmt.Println(v7, v3)
                    v15 := GenS2{F0: "\n\t"}
                    _ = v15
                }
            } else {
                v1[32768] = (42 < 0)
                if true {
                    if false {
                        switch (("\n\t" + "hello") + "foo bar") {
                        case "":
                            for fi16 := 0; fi16 < 2; fi16++ {
                                _ = len(v1)
                                for fi17 := 0; fi17 < 13; fi17++ {
                                    v18 := GenS3{F0: 681}
                                    v19 := GenS1{F0: 1}
                                    v6 = (v6 * v6)
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                                switch 0 {
                                case -1:
                                    for fi20 := 0; fi20 < 11; fi20++ {
                                        var v21 uint16 = 32767
                                        var v22 int64 = 9223372036854775807
                                        v23 := GenS3{F0: 1000}
                                        _ = fi20
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = (("" + ("" + "")) + ((("hello" + "\n\t") + "foo bar") + ("foo bar" + "foo bar")))
                                    }
                                }
                                _ = fi16
                            }
                        case "hello":
                            v24 := GenS1{F0: 255}
                            _ = v24
                        case "\n\t":
                            switch v3 {
                            case 32768:
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    v26 := GenS4{F0: -1}
                                    _ = fi25
                                    _ = v26
                                }
                            case 0:
                                if true {
                                    var v27 int = (-(v3))
                                    if false {
                                        v27 = 9223372036854775807
                                        v28 := GenS1{F0: 65535}
                                        v3 = 32768
                                        _ = v28
                                    }
                                    v1[(^(v3))] = false
                                    _ = v27
                                }
                            case -1:
                                var v29 string = ("hello" + "hello")
                                _ = v29
                            default:
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    v6 -= 2
                                    var v31 uint = v6
                                    _ = fi30
                                    _ = v31
                                }
                            }
                        }
                        fmt.Println(v13, v3, v6)
                    } else {
                        var v32 float64 = (-2.25 + (((-2.25 + (0.1 / 0.1)) * 0.1) + 3.14159))
                        _ = v32
                    }
                    fmt.Println(v6)
                }
                fmt.Println(v7, v3)
            }
            _ = v13
        } else {
            var v33 bool = true
            fmt.Println(v7)
            for fi34 := 0; fi34 < 8; fi34++ {
                if (-303 < int16(v6)) {
                    var v35 string = (("hello" + "foo bar") + "")
                    switch ("" + "\n\t") {
                    case "a":
                        for fi36 := 0; fi36 < 1; fi36++ {
                            var v37 uint32 = (v7 ^ (v7 | 2147483648))
                            var v38 uint32 = (^(v37))
                            delete(v1, v3)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    case "\n\t":
                        for fi39 := 0; fi39 < 12; fi39++ {
                            delete(v1, fi34)
                            if v33 {
                                v40 := GenS1{F0: 699}
                                v41 := GenS1{F0: 42}
                                _ = v40
                                _ = v41
                            } else {
                                v42 := GenS5{F0: 0}
                                _ = v42
                            }
                            _ = fi39
                        }
                    }
                    _ = v35
                }
                _ = fi34
            }
            _ = v33
        }
        fmt.Println(v6)
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        v3 += 0
        if ((v7 ^ (v7 >> (((v6 - 42)) & 63))) != v7) {
            for fi44 := 0; fi44 < 1; fi44++ {
                if (18446744073709551615 == 222) {
                    v7 -= (v7 | v7)
                } else {
                    fmt.Println(v3, fi43, v3)
                    var v45 float64 = (-2.25 + 0.1)
                    _ = v45
                }
                if true {
                    v46 := GenS1{F0: 65535}
                    var v47 int64 = 9223372036854775807
                    if false {
                        fmt.Println(v7, v6)
                        v48 := GenS4{F0: (v47 << (((v6 | (^(v6)))) & 63))}
                        _ = v48
                    } else {
                        var v49 float32 = (float32(1.5) + 3.14159)
                        if (v6 <= 42) {
                            if (197 != 18446744073709551615) {
                                v49 = v49
                                v50 := GenS4{F0: 42}
                                var v51 uint16 = 65535
                                _ = v50
                                _ = v51
                            } else {
                                var v52 int8 = 0
                                _ = v52
                            }
                            fmt.Println(v47, v6, fi44)
                            if true {
                                v47 = v47
                            } else {
                                var v53 string = "a"
                                _ = v53
                            }
                        }
                        v54 := GenS6{F0: 42, F1: -1}
                        var v55 string = ("foo bar" + ("" + ("foo bar" + ("" + "a"))))
                        if false {
                            fmt.Println(v49, v7)
                        }
                        _ = v49
                        _ = v54
                        _ = v55
                    }
                    _ = v46
                    _ = v47
                }
                v56 := GenS1{F0: 0}
                var v57 int = fi44
                switch "\n\t" {
                case "foo bar":
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v6, fi43)
                    case "hello":
                        v58 := GenS7{F0: 42, F1: 225}
                        _ = v58
                    default:
                        var v59 float64 = (float64(1.5) / -2.25)
                        _ = v59
                    }
                }
                _ = fi44
                _ = v56
                _ = v57
            }
            for fi60 := 0; fi60 < 16; fi60++ {
                v61 := GenS4{F0: 0}
                _ = len(v2)
                delete(v1, fi60)
                v3 -= fi60
                _ = fi60
                _ = v61
            }
        } else {
            v62 := GenS3{F0: -1}
            for fi63 := 0; fi63 < 12; fi63++ {
                if true {
                    var v64 string = "a"
                    var v65 float32 = ((0.1 + (1e10 - 1.5)) / 0.1)
                    var v66 float64 = -2.25
                    _ = v64
                    _ = v65
                    _ = v66
                }
                v6 = 455
                _ = fi63
            }
            _ = v62
        }
        for fi67 := 0; fi67 < 17; fi67++ {
            v68 := GenS1{F0: 255}
            if (!true) {
                v3 = (fi67 + (fi43 << (((v6 | (^(v6)))) & 63)))
                v7 = (^(v7))
                v2 = append(v2, "a")
            } else {
                v69 := GenS8{F0: 25, F1: 248, F2: (v7 * v7)}
                fmt.Println(v6, v7)
                var v70 string = "a"
                _ = v69
                _ = v70
            }
            _ = fi67
            _ = v68
        }
        _ = fi43
    }
    v3 += -9223372036854775808
    var v71 int = 290
    fmt.Println(v7)
    v1[v3] = ((!true) && (0 <= -23))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v71
}
