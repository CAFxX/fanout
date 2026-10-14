package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint16; F2 bool }
func genFn0(p0 uint, p1 uint64, p2 uint8) int8 {
    v7 := []int{32767}
    v8 := GenS1{F0: 710, F1: 42}
    var v9 string = ("\n\t" + "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint32, p1 uint8, p2 bool) bool {
    v10 := []int{1, 32767, -372}
    _ = len(v10)
    p0 = (p0 ^ 2147483647)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (false || ((("" + (("hello" + ("\n\t" + "\n\t")) + "hello")) + "") != (("foo bar" + ("foo bar" + ("foo bar" + ""))) + "foo bar")))
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]uint32)
    var v3 uint8 = 223
    v4 := make(map[int]string)
    var v5 int32 = 223
    v6 := make(map[uint32]string)
    if true {
        var v11 string = ("\n\t" + "hello")
        switch 42 {
        case 869:
            fmt.Println(v5)
        default:
            for fi12 := 0; fi12 < 8; fi12++ {
                v1 = -27
                switch int(v3) {
                case 43:
                    v11 = v11
                default:
                    v11 = "a"
                }
                _ = fi12
            }
        }
        _ = v11
    }
    var v13 int64 = v1
    var v14 int64 = 0
    _ = len(v4)
    if (false || (((!false) && (42 != -88)) || true)) {
        fmt.Println(v1)
    } else {
        for fi15 := 0; fi15 < 11; fi15++ {
            v13 += (v13 << ((1) & 63))
            var v16 uint8 = v3
            _ = fi15
            _ = v16
        }
        switch -1 {
        case -700:
            if false {
                if (float64((1e10 - ((1e10 / (3.14159 * 1e10)) + float32(0.0)))) == ((float64(0.1) / float64(-2.25)) / 0.1)) {
                    switch "a" {
                    case "hello":
                        switch 128 {
                        case 9223372036854775807:
                            var v17 uint = 42
                            _ = v17
                        case 742:
                            v5 = 42
                        }
                    case "":
                        var v18 float64 = (1.5 - float64(((float32(1.5) + (((1e10 - 3.14159) * 1.5) * 3.14159)) * -2.25)))
                        _ = v18
                    case "foo bar":
                        var v19 int8 = 1
                        _ = v19
                    default:
                        var v20 int32 = v5
                        _ = v20
                    }
                }
                if true {
                    if (v5 == (v5 - 239)) {
                        switch "hello" {
                        case "foo bar":
                            v21 := []int{}
                            _ = v21
                        case "a":
                            v13 = (v14 - (v1 | v13))
                        case "\n\t":
                            var v22 int64 = (v1 + v14)
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 9; fi23++ {
                            if false {
                                v1 = v1
                            } else {
                                v3 -= 1
                            }
                            v3 = v3
                            v14 -= -9223372036854775808
                            _ = fi23
                        }
                        var v24 int8 = genFn0(2, 42, (^(v3)))
                        _ = v24
                    }
                } else {
                    v25 := GenS1{F0: 2147483648, F1: 65535}
                    var v26 int16 = 1
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 uint64 = 18446744073709551615
                if (0 != 32767) {
                    if true {
                        var v28 uint8 = v3
                        _ = v28
                    } else {
                        if (-958 == -133) {
                            v14 -= 9223372036854775807
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v29 uint8 = 42
                            _ = v29
                        case "a":
                            switch ("" + "") {
                            case "":
                                fmt.Println(v13, v1, v5)
                            case "hello":
                                var v30 uint16 = uint16(v27)
                                _ = v30
                            case "a":
                                v31 := GenS2{F0: 1, F1: int(v5), F2: v27}
                                _ = v31
                            default:
                                var v32 int64 = 9223372036854775807
                                _ = v32
                            }
                        }
                    }
                }
                _ = v27
            }
        }
        var v33 string = "\n\t"
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v14)
        var v35 uint16 = 255
        _ = fi34
        _ = v35
    }
    if (!false) {
        if (65535 >= 65535) {
            var v36 int = 127
            _ = v36
        } else {
            delete(v2, "\n\t")
        }
        for fi37 := 0; fi37 < 7; fi37++ {
            var v38 int8 = genFn0(249, 78, (^(v3)))
            fmt.Println(v38, v3, v3)
            for fi39 := 0; fi39 < 10; fi39++ {
                v14 = v1
                var v40 uint64 = 679
                _ = fi39
                _ = v40
            }
            _ = fi37
            _ = v38
        }
        var v41 int64 = int64(v5)
        _ = v41
    }
    var v42 string = ("" + "a")
    v43 := GenS3{F0: 32767}
    for fi44 := 0; fi44 < 9; fi44++ {
        v1 += 42
        var v45 float32 = 1.5
        for fi46 := 0; fi46 < 18; fi46++ {
            v3 = (^(v3))
            v14 += v1
            _ = fi46
        }
        _ = fi44
        _ = v45
    }
    if (true && true) {
        delete(v4, 436)
        fmt.Println(v14)
        if true {
            if false {
                fmt.Println(v42, v14, v42)
            } else {
                fmt.Println(v3)
                var v47 int32 = v5
                _ = v47
            }
            for fi48 := 0; fi48 < 10; fi48++ {
                _ = len(v4)
                v49 := GenS1{F0: 960, F1: 1}
                var v50 int = (fi48 - fi48)
                _ = fi48
                _ = v49
                _ = v50
            }
        }
    } else {
        v51 := GenS4{F0: false, F1: 65535, F2: (!true)}
        _ = len(v2)
        _ = v51
    }
    v52 := []int{int(v5), -1, -1}
    v13 -= (v13 << ((4294967295) & 63))
    fmt.Println(v14, v13)
    v13 -= (v14 ^ (v14 + 42))
    for fi53 := 0; fi53 < 11; fi53++ {
        var v54 int = fi53
        v14 = v1
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v52
    _ = v6
}
