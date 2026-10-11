package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int8, p1 uint32) string {
    fmt.Println(p1)
    p0 += p0
    p1 = 482
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 int16) int {
    v5 := GenS1{F0: 128, F1: 42}
    v6 := GenS1{F0: 2, F1: 900}
    switch "\n\t" {
    case "foo bar":
        if (2147483648 != 742) {
            p0 += (p0 - p0)
        } else {
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := GenS1{F0: 128, F1: 1}
                var v9 uint = 4294967295
                v10 := GenS1{F0: 42, F1: 0}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
    default:
        switch 1 {
        case 128:
            if ((((0.1 + 0.1) / 1e10) - 0.0) != 0.1) {
                switch "" {
                case "":
                    for fi11 := 0; fi11 < 4; fi11++ {
                        p0 = (^(p0))
                        v12 := GenS2{F0: -67, F1: 42, F2: 42}
                        _ = fi11
                        _ = v12
                    }
                case "foo bar":
                    if true {
                        p0 = (p0 << ((577) & 63))
                        v13 := GenS1{F0: 255, F1: 42}
                        _ = v13
                    }
                case "\n\t":
                    if (0.1 <= float32(1.5)) {
                        var v14 int = 9223372036854775807
                        p0 = (-(p0))
                        _ = v14
                    }
                }
                switch ("foo bar" + "hello") {
                case "":
                    if true {
                        fmt.Println(p0)
                    } else {
                        if (87 == 42) {
                            if (!true) {
                                p0 -= 676
                                v15 := []int{}
                                for fi16 := 0; fi16 < 8; fi16++ {
                                    v17 := GenS1{F0: 127, F1: 542}
                                    switch "hello" {
                                    case "foo bar":
                                        var v18 string = ""
                                        _ = v18
                                    case "a":
                                        p0 = (^(p0))
                                    case "":
                                        var v19 string = "\n\t"
                                        _ = v19
                                    }
                                    for fi20 := 0; fi20 < 20; fi20++ {
                                        switch "foo bar" {
                                        case "a":
                                            p0 -= -778
                                        case "hello":
                                            v21 := GenS2{F0: -66, F1: 1, F2: 65535}
                                            _ = v21
                                        case "foo bar":
                                            v22 := GenS1{F0: 1, F1: 2}
                                            fmt.Println(fi16, p0)
                                            switch 42 {
                                            case 479:
                                                var v23 uint = 65535
                                                _ = v23
                                            case 2147483647:
                                                fmt.Println(fi20, fi20, fi20)
                                            case -9223372036854775808:
                                                v15 = append(v15, int(p0))
                                            }
                                            _ = v22
                                        default:
                                            p0 = -120
                                        }
                                        _ = fi20
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                                _ = v15
                            }
                            p0 = p0
                            p0 = p0
                        }
                    }
                case "a":
                    switch -684 {
                    case 9223372036854775807:
                        var v24 uint16 = 169
                        _ = v24
                    default:
                        v25 := GenS3{F0: 42, F1: -450}
                        _ = v25
                    }
                default:
                    switch "hello" {
                    case "":
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    }
                }
            } else {
                if (!(32767 <= 32767)) {
                    var v27 uint8 = 127
                    v27 = (v27 * (^(v27)))
                    if false {
                        if (true || true) {
                            p0 = p0
                        } else {
                            var v28 float64 = float64(1e10)
                            v29 := []int{}
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = v27
                }
            }
        case 127:
            if true {
                if ((true && false) && (!(float64((0.0 * float32((0.0 / 1.5)))) == 3.14159))) {
                    v30 := []int{2147483647, -401}
                    switch 9223372036854775807 {
                    case 32768:
                        var v31 uint = 18446744073709551615
                        _ = v31
                    case 32767:
                        for fi32 := 0; fi32 < 13; fi32++ {
                            p0 += (p0 << ((4294967295) & 63))
                            _ = fi32
                        }
                    }
                    if false {
                        p0 = p0
                        var v33 int32 = 592
                        for fi34 := 0; fi34 < 18; fi34++ {
                            var v35 float64 = 1e10
                            fmt.Println(fi34, fi34)
                            _ = fi34
                            _ = v35
                        }
                        _ = v33
                    }
                    _ = v30
                }
                switch 32767 {
                case 42:
                    v36 := GenS4{F0: 18446744073709551615, F1: 4294967295}
                    _ = v36
                case 127:
                    fmt.Println(p0)
                }
            } else {
                v37 := GenS1{F0: 2, F1: 4294967295}
                var v38 bool = ((39 > 12) && ((42 == 0) && (!((157 <= -2147483648) || (371 == 767)))))
                _ = v37
                _ = v38
            }
        case 214:
            fmt.Println(p0)
        }
    }
    v39 := GenS4{F0: 1, F1: 368}
    _ = p0
    _ = v39
    _ = v5
    _ = v6
    return 32768
}

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{2147483648, 2147483647, 315}
    var v3 uint16 = 305
    var v4 float32 = float32((3.14159 / ((3.14159 / (-2.25 - 1e10)) * 1.5)))
    var v40 int = 9223372036854775807
    var v41 uint = 0
    var v42 int32 = -1
    var v43 int32 = 253
    v4 = float32(0.1)
    var v44 bool = true
    _ = len(v2)
    fmt.Println(v3)
    fmt.Println(v4, v43)
    v1 = float32(((float64(v1) / 1e10) / float64(v4)))
    _ = len(v2)
    switch "foo bar" {
    case "":
        var v45 float32 = (v4 - -2.25)
        _ = v45
    case "\n\t":
        for fi46 := 0; fi46 < 18; fi46++ {
            v47 := GenS1{F0: 192, F1: 0}
            _ = fi46
            _ = v47
        }
    case "a":
        v40 = 123
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
